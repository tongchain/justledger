/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"justledger/fabric/core/chaincode/shim"
	"justledger/fabric/integration/chaincode/marbles_private"
)

func main() {
	err := shim.Start(&marbles_private.MarblesPrivateChaincode{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting Simple chaincode: %s", err)
		os.Exit(2)
	}
}
