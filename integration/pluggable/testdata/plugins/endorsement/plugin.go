/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"justledger/fabric/core/handlers/endorsement/api"
	"justledger/fabric/core/handlers/endorsement/builtin"
	"justledger/fabric/integration/pluggable"
)

// go build -buildmode=plugin -o plugin.so

// NewPluginFactory is the function ran by the plugin infrastructure to create an endorsement plugin factory.
func NewPluginFactory() endorsement.PluginFactory {
	pluggable.PublishEndorsementPluginActivation()
	return &builtin.DefaultEndorsementFactory{}
}
