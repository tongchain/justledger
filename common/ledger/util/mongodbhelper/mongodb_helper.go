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

package mongodbhelper

import (
	"fmt"

	"justledger/common/flogging"
	"justledger/core/ledger/kvledger/txmgmt/version"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Doc saved in mongodb
type MongodbDoc struct {
	Key         string         `json:"key"`
	Value       interface{}    `json:"value"`
	ChaincodeId string         `json:"chaincodeId"`
	Attachments *Attachment    `json:"attachments"`
	Version     version.Height `json:"version"`
}

//Doc getted from mongodb MongodbDoc
//Add id item compare to MongodbDoc
type MongodbResultDoc struct {
	Id          bson.ObjectId  `json:"_id"`
	Key         string         `json:"key"`
	Value       interface{}    `json:"value"`
	ChaincodeId string         `json:"chaincodeId"`
	Attachments *Attachment    `json:"attachments"`
	Version     version.Height `json:"version"`
}

//Value saved in mongodb when value is not a json
type Attachment struct {
	Name            string
	ContentType     string
	Length          uint64
	AttachmentBytes []byte
}

var logger = flogging.MustGetLogger("mongodbhelper")

type MongoDB struct {
	Db   *mgo.Database
	Conf *MongoDBConf
}

func (mongoDB *MongoDB) GetDefaultCollection() *mgo.Collection {
	return mongoDB.getCollection(mongoDB.Conf.CollectionName)
}

func (mongoDB *MongoDB) getCollection(collectionName string) *mgo.Collection {
	return mongoDB.Db.C(collectionName)
}

//Build defaultIndex with key and chaincode
func (mongoDB *MongoDB) BuildDefaultIndexIfNotExisted() error {
	c := mongoDB.GetDefaultCollection()
	indexs, err := c.Indexes()
	if err != nil {
		return err
	}

	//check the existence of index of {KEY,NS}
	hasIndexKEY := false
	hasIndexChaincodeId := false
	for _, index := range indexs {
		for _, key := range index.Key {
			if key == KEY {
				hasIndexKEY = true
			}
			if key == NS {
				hasIndexChaincodeId = true
			}
		}
	}

	//build the index of {KEY,NS} when it not exists
	if !(hasIndexKEY && hasIndexChaincodeId) {
		index := mgo.Index{Key: []string{KEY, NS}, Unique: true, DropDups: false, Background: false}
		err = c.EnsureIndex(index)
		if err != nil {
			logger.Errorf("Error during build index, error : %s", err.Error())
			return err
		}
	}

	return nil
}

func (mongoDB *MongoDB) BuildIndex(index mgo.Index, c mgo.Collection) error {
	err := c.EnsureIndex(index)
	if err != nil {
		logger.Warningf("build index failed: %s", err.Error())
		return err
	}
	return nil
}

func (mongoDB *MongoDB) Open() error {
	return nil
}

//Create the default collection
func (mongoDB *MongoDB) CreateCollection() error {

	c := mongoDB.GetDefaultCollection()

	err := c.Create(&mgo.CollectionInfo{})
	if err != nil {
		return err
	}
	return nil
}

func (mongoDB *MongoDB) Close() {

}

func (mongoDB *MongoDB) GetDoc(ns, key string) (*MongodbDoc, error) {
	collection := mongoDB.GetDefaultCollection()
	queryResult := collection.Find(bson.M{KEY: key, NS: ns})
	num, err := queryResult.Count()
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	//check the num of the query result
	//the num is valid when it is only 1
	var resultDoc MongodbDoc
	if num == 0 {
		logger.Debugf("The corresponding value of this key: %s doesn't exist", key)
		return nil, nil
	} else if num != 1 {
		logger.Errorf("This key:%s is repeated in the current collection", key)
		return nil, fmt.Errorf("The key %s is repeated", key)
	}
	queryResult.One(&resultDoc)
	return &resultDoc, nil
}

//Simple query method which don't use paging
//Get the limit number result of query
func (mongoDB *MongoDB) QueryDocuments(query interface{}) (*mgo.Iter, error) {
	collection := mongoDB.GetDefaultCollection()
	queryLimit := mongoDB.Conf.QueryLimit
	result := collection.Find(query).Sort(ID).Limit(queryLimit)
	return result.Iter(), nil
}

func (mongoDB *MongoDB) QueryDocumentPagingSample(pageNum, pageSize int, sortBy string, query interface{}) (*mgo.Iter, error) {
	collection := mongoDB.GetDefaultCollection()
	skipRecord := (pageNum - 1) * pageSize
	queryRes := collection.Find(query)
	queryRes.Count()

	result := queryRes.Sort(sortBy).Skip(skipRecord).Limit(pageSize)
	return result.Iter(), nil
}

//Paging query method
func (mongoDB *MongoDB) QueryDocumentPagingComplex(pageInfo *PagingOrQuery) (*PageResult, []*MongodbResultDoc, error) {
	collection := mongoDB.GetDefaultCollection()
	return pageInfo.QueryDocument(collection)
}

func (mongoDB *MongoDB) SaveDoc(doc MongodbDoc) error {
	collection := mongoDB.GetDefaultCollection()

	//remove the origin doc before save it
	collection.Remove(bson.M{KEY: doc.Key, NS: doc.ChaincodeId})
	err := collection.Insert(&doc)
	if err != nil {
		logger.Errorf("Error in insert the content of key : %s, error : %s", doc.Key, err.Error())
		return err
	}

	return nil
}

func (mongoDB *MongoDB) Delete(ns, key string) error {
	collection := mongoDB.GetDefaultCollection()
	err := collection.Remove(bson.M{KEY: key, NS: ns})
	if err != nil {
		logger.Errorf("Error %s happened while delete key %s", err.Error(), key)
		return err
	}

	return nil
}

func (mongoDB *MongoDB) GetIterator(ns, startkey string, endkey string, querylimit int, queryskip int) *mgo.Iter {
	collection := mongoDB.GetDefaultCollection()
	var queryResult *mgo.Query
	if endkey == "" {
		queryResult = collection.Find(bson.M{KEY: bson.M{"$gte": startkey}, NS: ns})
	} else {
		queryResult = collection.Find(bson.M{KEY: bson.M{"$gte": startkey, "$lt": endkey}, NS: ns})
	}

	queryResult = queryResult.Skip(queryskip)
	queryResult = queryResult.Limit(querylimit)
	queryResult.Sort(KEY)
	return queryResult.Iter()
}
