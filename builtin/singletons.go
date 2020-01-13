package builtin

import (
	addr "github.com/filecoin-project/go-address"
	util "github.com/filecoin-project/specs/util"
)

// Addresses for singleton system actors.
var (
	// Distinguished AccountActor that is the source of system implicit messages.
	SystemActorAddr        = mustMakeAddress(0)
	InitActorAddr          = mustMakeAddress(1)
	RewardActorAddr        = mustMakeAddress(2)
	CronActorAddr          = mustMakeAddress(3)
	StoragePowerActorAddr  = mustMakeAddress(4)
	StorageMarketActorAddr = mustMakeAddress(5)
	// Distinguished AccountActor that is the destination of all burnt funds.
	BurntFundsActorAddr = mustMakeAddress(99)
)

const FirstNonSingletonActorId = 100

func mustMakeAddress(id uint64) addr.Address {
	address, err := addr.NewIDAddress(id)
	util.Assert(err == nil)
	return address
}
