package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"
)

type Validator = exported.ValidatorI

type DistributionHooks interface {
	AfterAllocateTokensToValidator(ctx sdk.Context, val Validator, commission sdk.DecCoins, shared sdk.DecCoins)
}
