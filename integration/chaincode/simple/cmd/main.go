/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"justledger/fabric/core/chaincode/shim"
	"justledger/fabric/integration/chaincode/simple"
)

func main() {
	err := shim.Start(&simple.SimpleChaincode{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting Simple chaincode: %s", err)
		os.Exit(2)
	}
}
