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
	"os"
	"testing"

	"github.com/spf13/viper"
	"justledger/core/ledger/kvledger/txmgmt/statedb/commontests"
	"justledger/core/ledger/kvledger/txmgmt/version"
	ledgertestutil "justledger/core/ledger/testutil"
)

func TestMain(m *testing.M) {
	// Read the core.yaml file for default config.
	ledgertestutil.SetupCoreYAMLConfig()

	viper.Set("ledger.state.mongoDBConfig.url", "mongodb://mongodb:27017")
	viper.Set("ledger.state.mongoDBConfig.username", "")
	viper.Set("ledger.state.mongoDBConfig.password", "")
	result := m.Run()
	os.Exit(result)
}

func TestBasicRW(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testbasicrw")
	commontests.TestBasicRW(t, env.DBProvider)
}

func TestGetStateMultipleKeys(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testgetmultiplekeys")
	commontests.TestGetStateMultipleKeys(t, env.DBProvider)
}

func TestMultiDBBasicRW(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testmultidbbasicrw")
	defer env.Cleanup("testmultidbbasicrw2")
	commontests.TestMultiDBBasicRW(t, env.DBProvider)
}

func TestDeletes(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testdeletes")
	commontests.TestDeletes(t, env.DBProvider)
}

func TestEncodeDecodeValueAndVersion(t *testing.T) {
	testValueAndVersionEncodeing(t, []byte("value1"), version.NewHeight(1, 2))
	testValueAndVersionEncodeing(t, []byte{}, version.NewHeight(50, 50))
}

func testValueAndVersionEncodeing(t *testing.T, value []byte, version *version.Height) {
	//encodedValue := statedb.EncodeValue(value, version)
	//val, ver := statedb.DecodeValue(encodedValue)
	//testutil.AssertEquals(t, val, value)
	//testutil.AssertEquals(t, ver, version)
}

func TestIterator(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testiterator")
	commontests.TestIterator(t, env.DBProvider)
}

func TestJsonQuery(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testquery")
	commontests.TestMongoQuery(t, env.DBProvider)
}

func TestPaging(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("")
	commontests.TestPaingQuery(t, env.DBProvider)
}

func TestQueryOrPaging(t *testing.T) {
	env := NewTestDBEnv(t)
	defer env.Cleanup("testpaging")
	commontests.TestExecuteQueryPaging(t, env.DBProvider)
}
