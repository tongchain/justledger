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

package commontests

import (
	"strconv"
	"strings"
	"testing"

	"encoding/json"

	"justledger/common/ledger/testutil"
	"justledger/common/ledger/util/mongodbhelper"
	"justledger/core/ledger/kvledger/txmgmt/statedb"
	"justledger/core/ledger/kvledger/txmgmt/version"
)

func TestMongoQuery(t *testing.T, dbProvider statedb.VersionedDBProvider) {
	db, err := dbProvider.GetDBHandle("testquery")
	testutil.AssertNoError(t, err, "")
	db.Open()
	defer db.Close()
	batch := statedb.NewUpdateBatch()
	jsonValue1 := "{\"asset_name\": \"marble1\",\"color\": \"blue\",\"size\": 1,\"owner\": \"tom\"}"
	batch.Put("ns1", "key1", []byte(jsonValue1), version.NewHeight(1, 1))
	jsonValue2 := "{\"asset_name\": \"marble2\",\"color\": \"blue\",\"size\": 2,\"owner\": \"jerry\"}"
	batch.Put("ns1", "key2", []byte(jsonValue2), version.NewHeight(1, 2))
	jsonValue3 := "{\"asset_name\": \"marble3\",\"color\": \"blue\",\"size\": 3,\"owner\": \"fred\"}"
	batch.Put("ns1", "key3", []byte(jsonValue3), version.NewHeight(1, 3))
	jsonValue4 := "{\"asset_name\": \"marble4\",\"color\": \"blue\",\"size\": 4,\"owner\": \"martha\"}"
	batch.Put("ns1", "key4", []byte(jsonValue4), version.NewHeight(1, 4))
	jsonValue5 := "{\"asset_name\": \"marble5\",\"color\": \"blue\",\"size\": 5,\"owner\": \"fred\"}"
	batch.Put("ns1", "key5", []byte(jsonValue5), version.NewHeight(1, 5))
	jsonValue6 := "{\"asset_name\": \"marble6\",\"color\": \"blue\",\"size\": 6,\"owner\": \"elaine\"}"
	batch.Put("ns1", "key6", []byte(jsonValue6), version.NewHeight(1, 6))
	jsonValue7 := "{\"asset_name\": \"marble7\",\"color\": \"blue\",\"size\": 7,\"owner\": \"fred\"}"
	batch.Put("ns1", "key7", []byte(jsonValue7), version.NewHeight(1, 7))
	jsonValue8 := "{\"asset_name\": \"marble8\",\"color\": \"blue\",\"size\": 8,\"owner\": \"elaine\"}"
	batch.Put("ns1", "key8", []byte(jsonValue8), version.NewHeight(1, 8))
	jsonValue9 := "{\"asset_name\": \"marble9\",\"color\": \"green\",\"size\": 9,\"owner\": \"fred\"}"
	batch.Put("ns1", "key9", []byte(jsonValue9), version.NewHeight(1, 9))
	jsonValue10 := "{\"asset_name\": \"marble10\",\"color\": \"green\",\"size\": 10,\"owner\": \"mary\"}"
	batch.Put("ns1", "key10", []byte(jsonValue10), version.NewHeight(1, 10))
	jsonValue11 := "{\"asset_name\": \"marble11\",\"color\": \"cyan\",\"size\": 1000007,\"owner\": \"joe\"}"
	batch.Put("ns1", "key11", []byte(jsonValue11), version.NewHeight(1, 11))

	//add keys for a separate namespace
	batch.Put("ns2", "key1", []byte(jsonValue1), version.NewHeight(1, 12))
	batch.Put("ns2", "key2", []byte(jsonValue2), version.NewHeight(1, 13))
	batch.Put("ns2", "key3", []byte(jsonValue3), version.NewHeight(1, 14))
	batch.Put("ns2", "key4", []byte(jsonValue4), version.NewHeight(1, 15))
	batch.Put("ns2", "key5", []byte(jsonValue5), version.NewHeight(1, 16))
	batch.Put("ns2", "key6", []byte(jsonValue6), version.NewHeight(1, 17))
	batch.Put("ns2", "key7", []byte(jsonValue7), version.NewHeight(1, 18))
	batch.Put("ns2", "key8", []byte(jsonValue8), version.NewHeight(1, 19))
	batch.Put("ns2", "key9", []byte(jsonValue9), version.NewHeight(1, 20))
	batch.Put("ns2", "key10", []byte(jsonValue10), version.NewHeight(1, 21))

	savePoint := version.NewHeight(2, 22)
	db.ApplyUpdates(batch, savePoint)

	// query for owner=jerry, use namespace "ns1"
	itr, err := db.ExecuteQuery("ns1", "{\"query\":{\"owner\":\"jerry\"}}")
	testutil.AssertNoError(t, err, "")

	// verify one jerry result
	queryResult1, err := itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult1)
	versionedQueryRecord := queryResult1.(*statedb.VersionedKV)
	stringRecord := string(versionedQueryRecord.Value)
	bFoundRecord := strings.Contains(stringRecord, "jerry")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify no more results
	queryResult2, err := itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult2)

	// query for owner=jerry, use namespace "ns2"
	itr, err = db.ExecuteQuery("ns2", "{\"query\":{\"owner\":\"jerry\"}}")
	testutil.AssertNoError(t, err, "")

	// verify one jerry result
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult1)
	versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "jerry")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify no more results
	queryResult2, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult2)

	// query for owner=jerry, use namespace "ns3"
	itr, err = db.ExecuteQuery("ns3", "{\"query\":{\"owner\":\"jerry\"}}")
	testutil.AssertNoError(t, err, "")

	// verify results - should be no records
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult1)

	// query using bad query string
	itr, err = db.ExecuteQuery("ns1", "this is an invalid query string")
	testutil.AssertError(t, err, "Should have received an error for invalid query string")

	// query returns 0 records
	itr, err = db.ExecuteQuery("ns1", "{\"query\":{\"owner\":\"not_a_valid_name\"}}")
	testutil.AssertNoError(t, err, "")

	// verify no results
	queryResult3, err := itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult3)

	// query with complex selector, namespace "ns1"
	itr, err = db.ExecuteQuery("ns1", "{\"query\":{\"$and\":[{\"size\":{\"$gt\": 5}},{\"size\":{\"$lt\":8}},{\"size\":{\"$not\":{\"$eq\":6}}}]}}")
	testutil.AssertNoError(t, err, "")

	// verify one fred result
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult1)
	versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "fred")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify no more results
	queryResult2, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult2)

	// query with complex selector, namespace "ns2"
	itr, err = db.ExecuteQuery("ns2", "{\"query\":{\"$and\":[{\"size\":{\"$gt\": 5}},{\"size\":{\"$lt\":8}},{\"size\":{\"$not\":{\"$eq\":6}}}]}}")
	testutil.AssertNoError(t, err, "")

	// verify one fred result
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult1)
	versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "fred")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify no more results
	queryResult2, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult2)

	// query with complex selector, namespace "ns3"
	itr, err = db.ExecuteQuery("ns3", "{\"query\":{\"$and\":[{\"size\":{\"$gt\": 5}},{\"size\":{\"$lt\":8}},{\"size\":{\"$not\":{\"$eq\":6}}}]}}")
	testutil.AssertNoError(t, err, "")

	// verify no more results
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult1)

	// query with embedded implicit "AND" and explicit "OR", namespace "ns1"
	itr, err = db.ExecuteQuery("ns1", "{\"query\":{\"color\":\"green\",\"$or\":[{\"owner\":\"fred\"},{\"owner\":\"mary\"}]}}")
	testutil.AssertNoError(t, err, "")

	// verify one green result
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult1)
	versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "green")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify another green result
	queryResult2, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult2)
	versionedQueryRecord = queryResult2.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "green")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify no more results
	queryResult3, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult3)

	// query with embedded implicit "AND" and explicit "OR", namespace "ns2"
	itr, err = db.ExecuteQuery("ns2", "{\"query\":{\"color\":\"green\",\"$or\":[{\"owner\":\"fred\"},{\"owner\":\"mary\"}]}}")
	testutil.AssertNoError(t, err, "")

	// verify one green result
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult1)
	versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "green")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify another green result
	queryResult2, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult2)
	versionedQueryRecord = queryResult2.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "green")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify no more results
	queryResult3, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult3)

	// query with embedded implicit "AND" and explicit "OR", namespace "ns3"
	itr, err = db.ExecuteQuery("ns3", "{\"query\":{\"color\":\"green\",\"$or\":[{\"owner\":\"fred\"},{\"owner\":\"mary\"}]}}")
	testutil.AssertNoError(t, err, "")

	// verify no results
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult1)

	// query with integer with digit-count equals 7 and response received is also received
	// with same digit-count and there is no float transformation
	itr, err = db.ExecuteQuery("ns1", "{\"query\":{\"$and\":[{\"size\":{\"$eq\": 1000007}}]}}")
	testutil.AssertNoError(t, err, "")

	// verify one jerry result
	queryResult1, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNotNil(t, queryResult1)
	versionedQueryRecord = queryResult1.(*statedb.VersionedKV)
	stringRecord = string(versionedQueryRecord.Value)
	bFoundRecord = strings.Contains(stringRecord, "joe")
	testutil.AssertEquals(t, bFoundRecord, true)
	bFoundRecord = strings.Contains(stringRecord, "1000007")
	testutil.AssertEquals(t, bFoundRecord, true)

	// verify no more results
	queryResult2, err = itr.Next()
	testutil.AssertNoError(t, err, "")
	testutil.AssertNil(t, queryResult2)
}

func testInsert(t *testing.T, dbProvider statedb.VersionedDBProvider) {
	db, err := dbProvider.GetDBHandle("testpaging")
	testutil.AssertNoError(t, err, "")

	db.Open()
	defer db.Close()

	err = insertMultiData(db)
	testutil.AssertNoError(t, err, "")
}

func insertMultiData(db statedb.VersionedDB) error {
	batch := statedb.NewUpdateBatch()
	jsonValue1 := "{\"asset_name\": \"marble1\",\"color\": \"blue\",\"size\": 1,\"owner\": \"tom\"}"
	for i := 1; i <= 12001; i++ {
		batch.Put("ns1", "key"+strconv.Itoa(i), []byte(jsonValue1), version.NewHeight(1, uint64(i)))
	}
	savePoint := version.NewHeight(2, 1)
	err := db.ApplyUpdates(batch, savePoint)
	if err != nil {
		return err
	}
	return nil
}

func TestPaingQuery(t *testing.T, dbProvider statedb.VersionedDBProvider) {
	//insert the data before test paingquery
	testInsert(t, dbProvider)

	db, err := dbProvider.GetDBHandle("testpaging")
	testutil.AssertNoError(t, err, "")

	db.Open()
	defer db.Close()

	firstQuery := "{\"query\":{\"color\":\"blue\"},\"pagingInfo\":{\"currentPageNum\":1,\"pageSize\":30}}"
	resItr, err := db.ExecuteQuery("ns1", firstQuery)
	testutil.AssertNoError(t, err, "")

	paging, err := resItr.Next()
	testutil.AssertNoError(t, err, "")
	pagingV := paging.(*statedb.VersionedKV)

	var pagingDoc mongodbhelper.PagingDoc
	err = json.Unmarshal(pagingV.Value, &pagingDoc)
	testutil.AssertNoError(t, err, "")
	testutil.AssertEquals(t, pagingDoc.ReturnPageResult.TotalPage, 401)
	testutil.AssertEquals(t, pagingDoc.ReturnPageResult.TotalCount, 12001)

}

func TestExecuteQueryPaging(t *testing.T, dbProvider statedb.VersionedDBProvider) {
	//insert the data before test paingquery
	testInsert(t, dbProvider)

	db, err := dbProvider.GetDBHandle("testpaging")
	testutil.AssertNoError(t, err, "")

	db.Open()
	defer db.Close()

	queryOrPaging := "{\"query\":{\"age\":\"12\"}}"
	resItr, err := db.ExecuteQuery("ns1", queryOrPaging)
	testutil.AssertNoError(t, err, "")
	_, err = resItr.Next()
}
