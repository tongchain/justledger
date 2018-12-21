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
	"encoding/json"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//The query json object used for paging/simple query
//Treat as paging query when PagingInfo is not null
//Treat as simple query when PagingInfo is null
type PagingOrQuery struct {
	PagingInfo *PagingInfo `json:"pagingInfo"`
	Query      interface{} `json:"query"`
}

//Paging info used for paging query
//The result of query will not change with same query conditions(even updated data)
//TotalCount: the total count of query result
//TotalPage: the amount of page of query result
//LastPageRecordCount: the amount of record of the last page of query result
//CurrentPageNum: the current page number you want to query
//PageSize: the size of a page contains
//LastQueryPageNum: the page number of last time query
//LastQueryObjectId: the objectid of the record at last of result of last time query
//LastRecordObjectId: the objectid of the record of the query
//SortBy: the item sorted by(set to _id and will never change)
type PagingInfo struct {
	TotalCount          int           `json:"totalCount"`
	TotalPage           int           `json:"totalPage"`
	LastPageRecordCount int           `json:"lastPageRecordCount"`
	CurrentPageNum      int           `json:"currentPageNum"`
	PageSize            int           `json:"pageSize"`
	LastQueryPageNum    int           `json:"lastQueryPageNum"`
	LastQueryObjectId   bson.ObjectId `json:"lastQueryObjectId"`
	LastRecordObjectId  bson.ObjectId `json:"lastRecordObjectId"`
	SortBy              string        `json:"sortBy"`
}

type PageResult struct {
	TotalCount int `json:"totalCount"`
	TotalPage  int `json:"totalPage"`
	//the count of data of last page
	LastPageRecordCount int `json:"lastPageRecordCount"`
	PageSize            int `json:"pageSize"`
	//the objectId of last record of last query
	LastQueryObjectId bson.ObjectId `json:"lastQueryObjectId"`
	//the objectId of lastRecord of the query
	LastRecordObjectId bson.ObjectId `json:"lastRecordObjectId"`
	//the last page num of the query
	LastQueryPageNum int `json:"lastQueryPageNum"`
}

//the result of paging query return to chaincode
type PagingDoc struct {
	ReturnValue      interface{} `json:"returnValue"`
	ReturnPageResult *PageResult `json:"returnPageResult"`
}

// get reverse result sorted by item
const REVERSE_PREFIX = "-"

const (
	BEGIN_AT_FIRST_RECORD             = "beginAtFirstRecord"
	BEGIN_AT_LASTQUERY_ORDER          = "beginAtLastqueryOrder"
	BEGIN_AT_LASTQUERY_REVERSE_ORDER  = "beginAtLastqueryReveseOrderer"
	BEGIN_AT_LASTRECORD_REVERSE_ORDER = "beginAtLastrecordReverseOrder"
)

//Entry for paging query
func (pagingOrQuery *PagingOrQuery) QueryDocument(collection *mgo.Collection) (*PageResult, []*MongodbResultDoc, error) {

	pagingOrQuery.PagingInfo.SortBy = ID
	useFirstPagingFlag := pagingOrQuery.checkFirstPagingFlag()
	if useFirstPagingFlag {
		return pagingOrQuery.firstQueryDocument(collection)
	} else {
		return pagingOrQuery.queryDocumentComplex(collection)

	}
}

// execute first paging query when TotalCount or TotalPage or LastQueryObjectId
// or LastQueryPageNum is 0 or nil
func (pagingOrQuery *PagingOrQuery) firstQueryDocument(collection *mgo.Collection) (*PageResult, []*MongodbResultDoc, error) {
	pageResult := PageResult{}
	pageResult.PageSize = pagingOrQuery.PagingInfo.PageSize
	cursor := collection.Find(pagingOrQuery.Query)

	if pagingOrQuery.PagingInfo.CurrentPageNum != 1 {
		return nil, nil, fmt.Errorf("the pageNum must be 1 at the first time of paging")
	}

	count, err := cursor.Count()
	if err != nil {
		return nil, nil, err
	}

	if count <= 0 {
		pageResult.TotalCount = 0
		pageResult.TotalPage = 0
		pageResult.LastQueryObjectId = ""
		return &pageResult, nil, nil
	}

	//set the paging info
	pageResult.TotalCount = count
	pageResult.TotalPage, pageResult.LastPageRecordCount = GetPage(count, pagingOrQuery.PagingInfo.PageSize)
	pageResult.LastQueryPageNum = pagingOrQuery.PagingInfo.CurrentPageNum

	result := cursor.Sort(pagingOrQuery.PagingInfo.SortBy).Limit(pagingOrQuery.PagingInfo.PageSize)
	resIter := result.Iter()
	docs := GetDocs(resIter)

	recordCount := len(docs)
	if recordCount <= 0 {
		return nil, nil, nil
	}
	pageResult.LastQueryObjectId = docs[recordCount-1].Id

	//get objectid of last record of this query result
	lastDoc := &MongodbResultDoc{}
	tmpLastdoc := make(map[string]interface{})
	cursor.Sort(REVERSE_PREFIX + pagingOrQuery.PagingInfo.SortBy).One(tmpLastdoc)
	tmpJson, _ := json.Marshal(tmpLastdoc)
	json.Unmarshal(tmpJson, lastDoc)
	pageResult.LastRecordObjectId = lastDoc.Id
	resIter.Close()

	return &pageResult, docs, nil
}

func (pagingOrQuery *PagingOrQuery) queryDocumentSimple(collection *mgo.Collection, queryFlag string) (*PageResult, []*MongodbResultDoc, error) {

	pageResult := PageResult{}
	pageResult.PageSize = pagingOrQuery.PagingInfo.PageSize

	var skipRecord, limitRecords int
	query := pagingOrQuery.Query
	queryMap := query.(map[string]interface{})

	var reverseFlag bool
	reverseFlag = false

	switch queryFlag {
	// query start from the first record
	case BEGIN_AT_FIRST_RECORD:
		skipRecord = (pagingOrQuery.PagingInfo.CurrentPageNum - 1) * pagingOrQuery.PagingInfo.PageSize
		limitRecords = pagingOrQuery.PagingInfo.PageSize

		// query start from the record of last time query
		// end with the last record id to promise the total amount will not change if the data of mongodb updated
	case BEGIN_AT_LASTQUERY_ORDER:
		left := bson.M{ID: bson.M{"$gt": pagingOrQuery.PagingInfo.LastQueryObjectId}}
		right := bson.M{ID: bson.M{"$lt": pagingOrQuery.PagingInfo.LastRecordObjectId}}
		a := []bson.M{left, right}
		queryMap["$and"] = a
		pagingOrQuery.Query = queryMap

		// skip from last query page to page you want to query.
		skipRecord = (pagingOrQuery.PagingInfo.CurrentPageNum - pagingOrQuery.PagingInfo.LastQueryPageNum - 1) * pagingOrQuery.PagingInfo.PageSize
		limitRecords = pagingOrQuery.PagingInfo.PageSize

	case BEGIN_AT_LASTQUERY_REVERSE_ORDER:
		queryMap[ID] = bson.M{"$lte": pagingOrQuery.PagingInfo.LastQueryObjectId}
		pagingOrQuery.Query = queryMap
		pagingOrQuery.PagingInfo.SortBy = REVERSE_PREFIX + pagingOrQuery.PagingInfo.SortBy

		// skip to the page you want to query reverse from record of last time query
		skipRecord = (pagingOrQuery.PagingInfo.LastQueryPageNum - pagingOrQuery.PagingInfo.CurrentPageNum) * pagingOrQuery.PagingInfo.PageSize
		limitRecords = pagingOrQuery.PagingInfo.PageSize
		reverseFlag = true

		// query end with the last record of query
	case BEGIN_AT_LASTRECORD_REVERSE_ORDER:
		queryMap[ID] = bson.M{"$lte": pagingOrQuery.PagingInfo.LastRecordObjectId}
		pagingOrQuery.Query = queryMap
		pagingOrQuery.PagingInfo.SortBy = REVERSE_PREFIX + pagingOrQuery.PagingInfo.SortBy

		// get the distance between the page num you want to query and last page
		// get the record number needed to skip to reach page you want to query deal with the distance
		if pagingOrQuery.PagingInfo.TotalPage == pagingOrQuery.PagingInfo.CurrentPageNum {
			skipRecord = 0
			limitRecords = pagingOrQuery.PagingInfo.LastPageRecordCount
		} else if pagingOrQuery.PagingInfo.TotalPage-pagingOrQuery.PagingInfo.CurrentPageNum == 1 {
			skipRecord = pagingOrQuery.PagingInfo.LastPageRecordCount
			limitRecords = pagingOrQuery.PagingInfo.PageSize
		} else {
			skipRecord = (pagingOrQuery.PagingInfo.TotalPage-pagingOrQuery.PagingInfo.CurrentPageNum-1)*pagingOrQuery.PagingInfo.PageSize + pagingOrQuery.PagingInfo.LastPageRecordCount
			limitRecords = pagingOrQuery.PagingInfo.PageSize
		}
		reverseFlag = true

	default:
		skipRecord = (pagingOrQuery.PagingInfo.CurrentPageNum - 1) * pagingOrQuery.PagingInfo.PageSize
	}

	// get result of query
	cursor := collection.Find(pagingOrQuery.Query)
	result := cursor.Sort(pagingOrQuery.PagingInfo.SortBy).Skip(skipRecord).Limit(limitRecords)
	resIter := result.Iter()
	docs := GetDocs(resIter)

	if len(docs) <= 0 {
		return nil, nil, nil
	}

	if reverseFlag {
		docs = ReverseDocs(docs)
	}

	//Set the pageResult Info of query
	pageResult.LastQueryPageNum = pagingOrQuery.PagingInfo.CurrentPageNum
	pageResult.TotalCount = pagingOrQuery.PagingInfo.TotalCount
	pageResult.TotalPage = pagingOrQuery.PagingInfo.TotalPage
	pageResult.LastPageRecordCount = pagingOrQuery.PagingInfo.LastPageRecordCount
	recordCount := len(docs)
	pageResult.LastQueryObjectId = docs[recordCount-1].Id
	pageResult.LastRecordObjectId = pagingOrQuery.PagingInfo.LastRecordObjectId

	resIter.Close()
	return &pageResult, docs, nil
}

// Calcute the distance between of firstPageNum, lastPageNum and the pageNum of lastQuery
// query start or reverse start from the page which is nearest of current query page
func (pagingOrQuery *PagingOrQuery) queryDocumentComplex(collection *mgo.Collection) (*PageResult, []*MongodbResultDoc, error) {

	firstDistance := pagingOrQuery.PagingInfo.CurrentPageNum
	lastDistance := pagingOrQuery.PagingInfo.TotalPage - pagingOrQuery.PagingInfo.CurrentPageNum
	midDistance := pagingOrQuery.PagingInfo.CurrentPageNum - pagingOrQuery.PagingInfo.LastQueryPageNum
	midDistanceAbs := AbsInt(midDistance)

	//query start from first record
	if firstDistance <= lastDistance && firstDistance <= midDistanceAbs {
		return pagingOrQuery.queryDocumentSimple(collection, BEGIN_AT_FIRST_RECORD)
	}

	//query reverse start from last record
	if lastDistance <= firstDistance && lastDistance <= midDistanceAbs {
		return pagingOrQuery.queryDocumentSimple(collection, BEGIN_AT_LASTRECORD_REVERSE_ORDER)
	}

	//query start or reverse start from the query num of last time
	if midDistanceAbs <= firstDistance && midDistanceAbs <= lastDistance {
		if midDistance > 0 {
			return pagingOrQuery.queryDocumentSimple(collection, BEGIN_AT_LASTQUERY_ORDER)
		} else {
			return pagingOrQuery.queryDocumentSimple(collection, BEGIN_AT_LASTQUERY_REVERSE_ORDER)
		}
	}

	return nil, nil, nil
}

// check if this is the first time of query
func (pagingOrQuery *PagingOrQuery) checkFirstPagingFlag() bool {

	if pagingOrQuery.PagingInfo.TotalCount == 0 {
		return true
	}

	if pagingOrQuery.PagingInfo.TotalPage == 0 {
		return true
	}

	if pagingOrQuery.PagingInfo.LastQueryObjectId == "" {
		return true
	}

	if pagingOrQuery.PagingInfo.LastQueryPageNum == 0 {
		return true
	}

	return false
}

// get (totalCount, lastPageRecordCount)
func GetPage(count, pageSize int) (int, int) {
	var totalPage int
	totalPage = count / pageSize
	if count%pageSize == 0 {
		return totalPage, pageSize
	} else {
		return totalPage + 1, count % pageSize
	}
}

// get docs by iterator
func GetDocs(res *mgo.Iter) []*MongodbResultDoc {
	tmpDoc := make(map[string]interface{})
	docRes := make([]*MongodbResultDoc, 0)

	for ok := res.Next(tmpDoc); ok; ok = res.Next(tmpDoc) {
		doc := &MongodbResultDoc{}
		docJson, _ := json.Marshal(tmpDoc)
		json.Unmarshal(docJson, doc)
		docRes = append(docRes, doc)
	}
	return docRes
}

func ReverseDocs(docs []*MongodbResultDoc) []*MongodbResultDoc {
	index := len(docs) - 1
	for i := index; i >= index/2; i-- {
		tmp := docs[index-i]
		docs[index-i] = docs[i]
		docs[i] = tmp
	}
	return docs
}

// return abs(int)
func AbsInt(a int) int {
	if a > 0 {
		return a
	} else {
		return 0 - a
	}
}
