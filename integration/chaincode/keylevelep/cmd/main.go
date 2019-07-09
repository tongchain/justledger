/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"os"

	"justledger/fabric/core/chaincode/shim"
	"justledger/fabric/integration/chaincode/keylevelep"
)

func main() {
	err := shim.Start(&keylevelep.EndorsementCC{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exiting SBE chaincode: %s", err)
		os.Exit(2)
	}
}
