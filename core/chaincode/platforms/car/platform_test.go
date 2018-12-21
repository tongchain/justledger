/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package car_test

import (
	"os"
	"path/filepath"
	"testing"

	"justledger/common/util"
	"justledger/core/chaincode/platforms"
	"justledger/core/chaincode/platforms/car"
	"justledger/core/testutil"
	pb "justledger/protos/peer"
)

var _ = platforms.Platform(&car.Platform{})

func TestMain(m *testing.M) {
	testutil.SetupTestConfig()
	os.Exit(m.Run())
}

func TestCar_BuildImage(t *testing.T) {
	vm, err := NewVM()
	if err != nil {
		t.Errorf("Error getting VM: %s", err)
		return
	}

	chaincodePath := filepath.Join("testdata", "/org.hyperledger.chaincode.example02-0.1-SNAPSHOT.car")
	spec := &pb.ChaincodeSpec{
		Type: pb.ChaincodeSpec_CAR,
		ChaincodeId: &pb.ChaincodeID{
			Name: "cartest",
			Path: chaincodePath,
		},
		Input: &pb.ChaincodeInput{
			Args: util.ToChaincodeArgs("f"),
		},
	}
	if err := vm.BuildChaincodeContainer(spec); err != nil {
		t.Error(err)
	}
}
