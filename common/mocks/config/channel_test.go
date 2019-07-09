/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"testing"

	"justledger/common/channelconfig"
)

func TestChannelConfigInterface(t *testing.T) {
	_ = channelconfig.Channel(&Channel{})
}
