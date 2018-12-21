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
	"time"

	"github.com/spf13/viper"
)

type MongoDBConf struct {
	Url            string
	UserName       string
	Password       string
	DBName         string
	CollectionName string
	QueryLimit     int
	RequestTimeout time.Duration
}

func GetMongoDBConf() *MongoDBConf {
	url := viper.GetString("ledger.state.mongoDBConfig.url")
	userName := viper.GetString("ledger.state.mongoDBConfig.username")
	password := viper.GetString("ledger.state.mongoDBConfig.password")
	collectionName := viper.GetString("ledger.state.mongoDBConfig.collectionName")
	timeout := viper.GetDuration("ledger.state.mongoDBConfig.requestTimeout")

	if collectionName == "" {
		collectionName = "test"
	}

	dbName := viper.GetString("ledger.state.mongoDBConfig.databaseName")
	if dbName == "" {
		dbName = "mongotest"
	}

	queryLimit := viper.GetInt("ledger.state.mongoDBConfig.queryLimit")
	if queryLimit <= 0 {
		queryLimit = 1000
	}

	if timeout <= 0 {
		timeout, _ = time.ParseDuration("35s")
	}

	return &MongoDBConf{
		Url:            url,
		UserName:       userName,
		Password:       password,
		DBName:         dbName,
		CollectionName: collectionName,
		QueryLimit:     queryLimit,
		RequestTimeout: timeout,
	}
}
