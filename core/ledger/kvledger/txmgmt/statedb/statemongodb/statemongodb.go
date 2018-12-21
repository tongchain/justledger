/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package statemongodb

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"sync"
	"unicode/utf8"

	"github.com/pkg/errors"
	"justledger/common/flogging"
	"justledger/common/ledger/util/mongodbhelper"
	"justledger/core/common/ccprovider"
	"justledger/core/ledger/kvledger/txmgmt/statedb"
	"justledger/core/ledger/kvledger/txmgmt/version"
)

const dataWrapper = "value"

var logger = flogging.MustGetLogger("statemongodb")

var savePointKey = "statedb_savepoint"
var savePointNs = "savepoint"

var queryskip = 0

type VersionedDBProvider struct {
	session   *mgo.Session
	databases map[string]*VersionedDB
	mux       sync.Mutex
}

type VersionedDB struct {
	mongoDB *mongodbhelper.MongoDB
	dbName  string
}

func NewVersionedDBProvider() (*VersionedDBProvider, error) {

	logger.Debugf("constructing MongoDB VersionedDBProvider")

	mongodbConf := mongodbhelper.GetMongoDBConf()
	dialInfo, err := mgo.ParseURL(mongodbConf.Url)
	if err != nil {
		return nil, err
	}
	dialInfo.Timeout = mongodbConf.RequestTimeout
	mgoSession, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		logger.Errorf(err.Error())
		return nil, err
	}

	return &VersionedDBProvider{
		session:   mgoSession,
		databases: make(map[string]*VersionedDB),
		mux:       sync.Mutex{},
	}, nil
}

//TODO retry while get error
func (provider *VersionedDBProvider) GetDBHandle(dbName string) (statedb.VersionedDB, error) {
	provider.mux.Lock()
	defer provider.mux.Unlock()

	db, ok := provider.databases[dbName]
	if ok {
		return db, nil
	}

	vdr, err := newVersionDB(provider.session, dbName)
	if err != nil {
		return nil, err
	}

	return vdr, nil
}

func (provider *VersionedDBProvider) Close() {
	provider.session.Close()
}

// GetState implements method in VersionedDB interface
func (vdb *VersionedDB) GetState(namespace string, key string) (*statedb.VersionedValue, error) {
	logger.Debugf("GetState(). ns=%s, key=%s", namespace, key)
	doc, err := vdb.mongoDB.GetDoc(namespace, key)
	if err != nil {
		return nil, err
	}
	if doc == nil {
		return nil, nil
	}

	versionedV := statedb.VersionedValue{Value: nil, Version: &doc.Version}
	if doc.Value == nil && doc.Attachments != nil {
		versionedV.Value = doc.Attachments.AttachmentBytes
	} else {
		tempbyte, _ := json.Marshal(doc.Value)
		versionedV.Value = tempbyte
	}

	return &versionedV, nil
}

// GetVersion implements method in VersionedDB interface
//TODO support GetVersion
func (vdb *VersionedDB) GetVersion(namespace string, key string) (*version.Height, error) {
	versionedValue, err := vdb.GetState(namespace, key)
	if err != nil {
		return nil, err
	}
	if versionedValue == nil {
		return nil, nil
	}
	return versionedValue.Version, nil
}

// Add "value." to all keys
func processKey(key string) string {
	return fmt.Sprintf("%v.%v", dataWrapper, key)
}

// traverse through the value
func wrapIndex(index mgo.Index) {
	for i, key := range index.Key {
		wrapKey := processKey(key)
		index.Key[i] = wrapKey
	}
}

// ProcessIndexesForChaincodeDeploy creates indexes for a specified namespace
func (vdb *VersionedDB) ProcessIndexesForChaincodeDeploy(namespace string, fileEntries []*ccprovider.TarFileEntry) error {
	c := vdb.mongoDB.GetDefaultCollection()
	for _, fileEntry := range fileEntries {
		indexData := fileEntry.FileContent
		index := mgo.Index{}
		json.Unmarshal(indexData, &index)
		//TODO warpper
		wrapIndex(index)
		filename := fileEntry.FileHeader.Name
		err := vdb.mongoDB.BuildIndex(index, *c)
		if err != nil {
			return fmt.Errorf("error during creation of index from file=[%s] for chain=[%s]. Error=%s",
				filename, namespace, err)
		}
	}
	return nil
}

func (vdb *VersionedDB) GetDBType() string {
	return "mongodb"
}

// GetStateMultipleKeys implements method in VersionedDB interface
func (vdb *VersionedDB) GetStateMultipleKeys(namespace string, keys []string) ([]*statedb.VersionedValue, error) {
	vals := make([]*statedb.VersionedValue, len(keys))
	for i, key := range keys {
		val, err := vdb.GetState(namespace, key)
		if err != nil {
			return nil, err
		}
		vals[i] = val
	}
	return vals, nil
}

// GetStateRangeScanIterator implements method in VersionedDB interface
// startKey is inclusive
// endKey is exclusive
func (vdb *VersionedDB) GetStateRangeScanIterator(namespace string, startKey string, endKey string) (statedb.ResultsIterator, error) {
	return vdb.GetStateRangeScanIteratorWithMetadata(namespace, startKey, endKey, nil)
}

// TODO GetStateRangeScanIteratorWithMetadata implements method in VersionedDB interface
func (vdb *VersionedDB) GetStateRangeScanIteratorWithMetadata(namespace string, startKey string, endKey string, metadata map[string]interface{}) (statedb.QueryResultsIterator, error) {
	querylimit := vdb.mongoDB.Conf.QueryLimit
	dbItr := vdb.mongoDB.GetIterator(namespace, startKey, endKey, querylimit, queryskip)
	return newKVScanner(dbItr, namespace), nil
}

// ExecuteQuery implements method in VersionedDB interface
func (vdb *VersionedDB) ExecuteQuery(namespace, queryOrPaingStr string) (statedb.ResultsIterator, error) {

	queryByte := []byte(queryOrPaingStr)
	if !isJson(queryByte) {
		return nil, fmt.Errorf("the queryOrPaingStr is no" +
			"t a json : %s", queryOrPaingStr)
	}

	paingOrQuery := &mongodbhelper.PagingOrQuery{}
	err := json.Unmarshal(queryByte, paingOrQuery)
	if err != nil {
		return nil, fmt.Errorf("the queryOrPaingStr string is not a pagingOrQuery json string:" + err.Error())
	}

	pagingInfo := paingOrQuery.PagingInfo
	//execute normal json query when paginginfo is nil
	//execute paging query when paginginfo is not nil
	if pagingInfo == nil {
		queryInterface := paingOrQuery.Query
		queryStr, err := json.Marshal(queryInterface)
		if err != nil {
			return nil, err
		}
		return vdb.JsonQuery(namespace, string(queryStr))
	} else {
		return vdb.PagingQuery(namespace, paingOrQuery)
	}
}

// TODO ExecuteQueryWithMetadata implements method in VersionedDB interface
func (vdb *VersionedDB) ExecuteQueryWithMetadata(namespace, query string, metadata map[string]interface{}) (statedb.QueryResultsIterator, error) {
	return nil, errors.New("ExecuteQueryWithMetadata not supported for leveldb")
}

// ApplyUpdates implements method in VersionedDB interface
func (vdb VersionedDB) ApplyUpdates(batch *statedb.UpdateBatch, height *version.Height) error {
	namespaces := batch.GetUpdatedNamespaces()
	var out interface{}
	var err error
	for _, ns := range namespaces {
		updates := batch.GetUpdates(ns)
		for k, vv := range updates {
			logger.Debugf("Channel [%s]: Applying key=[%#v]", vdb.dbName, k)
			if vv.Value == nil {
				vdb.mongoDB.Delete(ns, k)
			} else {
				doc := mongodbhelper.MongodbDoc{
					Key:         k,
					Version:     *vv.Version,
					ChaincodeId: ns,
				}

				if !isJson(vv.Value) {
					logger.Debugf("Not a json, write it to the attachment ")
					attachmentByte := mongodbhelper.Attachment{
						AttachmentBytes: vv.Value,
					}
					doc.Attachments = &attachmentByte
					doc.Value = nil
				} else {
					err = json.Unmarshal(vv.Value, &out)
					if err != nil {
						logger.Errorf("Error rises while unmarshal the vv.value, error :%s", err.Error())
					}
					doc.Value = out
				}

				err = vdb.mongoDB.SaveDoc(doc)
				if err != nil {
					logger.Errorf("Error during Commit(): %s\n", err.Error())
				}
			}
		}
	}

	err = vdb.recordSavepoint(height)
	if err != nil {
		logger.Errorf("Error during recordSavepoint : %s\n", err.Error())
		return err
	}

	return nil
}

// GetLatestSavePoint implements method in VersionedDB interface
func (vdb *VersionedDB) GetLatestSavePoint() (*version.Height, error) {

	doc, err := vdb.mongoDB.GetDoc(savePointNs, savePointKey)
	if err != nil {
		logger.Errorf("Error during get latest save point key, error : %s", err.Error())
		return nil, err
	}
	if doc == nil {
		return nil, nil
	}

	return &doc.Version, nil
}

// ValidateKey implements method in VersionedDB interface
func (vdb *VersionedDB) ValidateKey(key string) error {
	if !utf8.ValidString(key) {
		return fmt.Errorf("Key should be a valid utf8 string: [%x]", key)
	}
	return nil
}

// BytesKeySuppoted implements method in VersionedDB interface
func (vdb *VersionedDB) BytesKeySuppoted() bool {
	return false
}

// Open implements method in VersionedDB interface
func (vdb *VersionedDB) Open() error {
	return nil
}

// Close implements method in VersionedDB interface
func (vdb *VersionedDB) Close() {

}

// ValidateKeyValue implements method in VersionedDB interface
func (vdb *VersionedDB) ValidateKeyValue(key string, value []byte) error {
	if !utf8.ValidString(key) {
		return fmt.Errorf("Key should be a valid utf8 string: [%x]", key)
	}

	// MongoDB supports any bytes for the value
	return nil
}

func (vdb *VersionedDB) recordSavepoint(height *version.Height) error {
	err := vdb.mongoDB.Delete(savePointNs, savePointKey)
	if err != nil {
		logger.Debugf(err.Error())
	}

	doc := mongodbhelper.MongodbDoc{}
	doc.Version.BlockNum = height.BlockNum
	doc.Version.TxNum = height.TxNum
	doc.Value = nil
	doc.ChaincodeId = savePointNs
	doc.Key = savePointKey

	err = vdb.mongoDB.SaveDoc(doc)
	if err != nil {
		logger.Errorf("Error during update savepoint , error : %s", err.Error())
		return err
	}

	return err
}

// return paging result of json query
func (vdb *VersionedDB) PagingQuery(namespace string, pagingOrQuery *mongodbhelper.PagingOrQuery) (statedb.ResultsIterator, error) {
	queryInterface := pagingOrQuery.Query
	queryStr, _ := json.Marshal(queryInterface)
	queryBson, err := mongodbhelper.GetQueryBson(namespace, string(queryStr))
	pagingOrQuery.Query = queryBson

	pagingRes, docs, err := vdb.mongoDB.QueryDocumentPagingComplex(pagingOrQuery)
	if err != nil {
		return nil, err
	}

	return newPagingScanner(docs, pagingRes), nil
}

func (vdb *VersionedDB) JsonQuery(namespace, query string) (statedb.ResultsIterator, error) {
	queryBson, err := mongodbhelper.GetQueryBson(namespace, query)
	if err != nil {
		return nil, err
	}

	var result *mgo.Iter
	result, err = vdb.mongoDB.QueryDocuments(queryBson)
	if err != nil {
		return nil, err
	}
	return newKVScanner(result, namespace), nil
}

type kvScanner struct {
	namespace string
	result    *mgo.Iter
}

func (scanner *kvScanner) Next() (statedb.QueryResult, error) {
	doc := mongodbhelper.MongodbDoc{}
	if !scanner.result.Next(&doc) {
		return nil, nil
	}

	blockNum := doc.Version.BlockNum
	txNum := doc.Version.TxNum
	height := version.NewHeight(blockNum, txNum)
	key := doc.Key
	value := doc.Value
	valueContent := []byte{}
	if doc.Value != nil {
		valueContent, _ = json.Marshal(value)
	} else {
		valueContent = doc.Attachments.AttachmentBytes
	}

	return &statedb.VersionedKV{
		CompositeKey:   statedb.CompositeKey{Namespace: scanner.namespace, Key: key},
		VersionedValue: statedb.VersionedValue{Value: valueContent, Version: height},
	}, nil
}

func (scanner *kvScanner) Close() {
	err := scanner.result.Close()
	if err != nil {
		logger.Errorf("Error during close the iterator of scanner error : %s", err.Error())
	}
}

func (scanner *kvScanner) GetBookmarkAndClose() string {
	retval := ""
	scanner.Close()
	return retval
}

func newKVScanner(iter *mgo.Iter, namespace string) *kvScanner {
	return &kvScanner{namespace: namespace, result: iter}
}

type pagingScanner struct {
	cursor       int
	docs         []*mongodbhelper.MongodbResultDoc
	pagingResult *mongodbhelper.PageResult
}

func (scanner *pagingScanner) Next() (statedb.QueryResult, error) {

	scanner.cursor++

	if scanner.cursor >= len(scanner.docs) {
		return nil, nil
	}

	doc := scanner.docs[scanner.cursor]
	returnVersion := doc.Version
	var returnValue []byte
	value := doc.Value

	// return paging query info in first result
	if scanner.cursor == 0 {
		tmpValue := mongodbhelper.PagingDoc{
			ReturnValue:      value,
			ReturnPageResult: scanner.pagingResult,
		}
		returnValue, _ = json.Marshal(tmpValue)
	} else {
		tmpValue, _ := json.Marshal(value)
		returnValue = tmpValue
	}

	return &statedb.VersionedKV{
		CompositeKey:   statedb.CompositeKey{Namespace: doc.ChaincodeId, Key: doc.Key},
		VersionedValue: statedb.VersionedValue{Value: returnValue, Version: &returnVersion},
	}, nil
}

func (scanner *pagingScanner) Close() {
	scanner = nil
}

func newPagingScanner(docs []*mongodbhelper.MongodbResultDoc, pagingResult *mongodbhelper.PageResult) *pagingScanner {
	return &pagingScanner{
		cursor:       -1,
		docs:         docs,
		pagingResult: pagingResult,
	}
}

//get or create newVersionDB
func newVersionDB(mgoSession *mgo.Session, dbName string) (*VersionedDB, error) {
	db := mgoSession.DB(dbName)
	conf := mongodbhelper.GetMongoDBConf()
	conf.DBName = dbName
	MongoDB := &mongodbhelper.MongoDB{db, conf}

	collectionsName, err := db.CollectionNames()
	if err != nil {
		return nil, err
	}

	//create default collection when it no exists
	if !contains(collectionsName, conf.CollectionName) {
		err := MongoDB.CreateCollection()
		if err != nil {
			return nil, err
		}
	}

	//build default index of {KEY,NS} when it not exists
	err = MongoDB.BuildDefaultIndexIfNotExisted()
	if err != nil {
		return nil, err
	}

	return &VersionedDB{MongoDB, dbName}, nil
}

func isJson(value []byte) bool {
	var result interface{}
	err := json.Unmarshal(value, &result)
	return err == nil
}

func contains(src []string, value string) bool {

	if src == nil {
		return false
	}

	for _, v := range src {
		if v == value {
			return true
		}
	}

	return false
}
