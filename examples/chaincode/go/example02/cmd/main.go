/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	"github.com/justledger/fabric/core/chaincode/shim"
	"github.com/justledger/fabric/examples/chaincode/go/example02"
)

func main() {
	err := shim.Start(new(example02.SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
