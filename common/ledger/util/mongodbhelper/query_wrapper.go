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
	"bytes"
	"encoding/json"
	"fmt"
)

/*
KEY is used as the item referring to key stored in mongodb.
KEY and NS will be a unique index build by default to query.
ID used for paging query function whose result is sorted by "_id"(can't be changed).
The detial about paging query will be found in mongodb_page.go
dataWrapper
*/
const KEY = "key"
const dataWrapper = "value"
const NS = "chaincodeid"
const ID = "_id"

// Vaild operations in mongodb query
var VALID_OPERATORS = []string{
	"$eq", "$gt", "$gte", "$in", "$lt", "$lte", "$ne", "$nin", "$and", "$not", "$nor", "$or", "$exists",
	"$type", "$mod", "$regex", "$text", "$where", "$geoIntersects", "$geoWithin", "$near", "$nearSphere",
	"$all", "$elemMatch", "$size", "$bitsAllClear", "$bitsAnyClear", "$bitsAnySet", "$comment", "$", "$elemMatch",
	"$meta", "$slice",
}

/*
GetQueryBson parses the query string passed to MongoDB
the wrapper prepends the wrapper "value." to all fields(except valid operations) specified in the query

Example:
Source Query:
{"owner":{"$eq": "tom"},
{"$and":[{"size":{"$gt": 5}}, {"size":{"$lt":8}}]},
}

Result Wrapped Query(chaincodeid is mycc):
{"value.owner":{"$eq": "tom"},
{"$and":[{"value.size":{"$gt": 5}}, {"value.size":{"$lt":8}}}]},
{"chaincodeid":"mycc"},
}
*/
func GetQueryBson(namespace, query string) (interface{}, error) {
	//create a generic map for the query json
	if query == "null" {
		return "", fmt.Errorf("null of query")
	}
	jsonQueryMap := make(map[string]interface{})
	//unmarshal the selected json into the generic map
	decoder := json.NewDecoder(bytes.NewBuffer([]byte(query)))
	decoder.UseNumber()
	err := decoder.Decode(&jsonQueryMap)
	if err != nil {
		return "", err
	}

	//traverse through the json query and wrap any field names
	processAndWrapQuery(jsonQueryMap)
	jsonQueryMap[NS] = namespace
	return jsonQueryMap, nil
}

// traverse through the value and process traverse through the fileds
func processAndWrapQuery(jsonQueryMap map[string]interface{}) {
	for jsonKey, jsonValue := range jsonQueryMap {
		wraperField := processField(jsonKey)
		delete(jsonQueryMap, jsonKey)
		wraperValue := processValue(jsonValue)
		jsonQueryMap[wraperField] = wraperValue
	}
}

// Add "value." to all fields except valid operations
func processField(field string) string {
	if isValidOperator(field) {
		return field
	} else {
		return fmt.Sprintf("%v.%v", dataWrapper, field)
	}
}

func processValue(value interface{}) interface{} {
	switch value.(type) {

	// return value rightly
	case string:
		return value

		// return value rightly
	case json.Number:
		return value

		// if it is a array, then iterate through the items
	case []interface{}:
		arrayValue := value.([]interface{})
		resultValue := make([]interface{}, 0)
		// iterate through the items
		for _, itemValue := range arrayValue {
			resultValue = append(resultValue, processValue(itemValue.(interface{})))
		}

		return resultValue

		// process this as a map
	case interface{}:
		tmp, ok := value.(map[string]interface{})
		if ok {
			processAndWrapQuery(tmp)
		} else {
			return value
		}
		return value

		// this never reach
	default:
		return value
	}
}

// Check the field is a valid operation or not
func isValidOperator(field string) bool {
	return arrayContains(VALID_OPERATORS, field)
}

// detect if a source array contains times
// used for check if a field is a valid operation
func arrayContains(source []string, item string) bool {
	for _, s := range source {
		if s == item {
			return true
		}
	}
	return false
}
