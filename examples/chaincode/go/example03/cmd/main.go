/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"

	"justledger/core/chaincode/shim"
	"justledger/examples/chaincode/go/example03"
)

func main() {
	err := shim.Start(new(example03.SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}
