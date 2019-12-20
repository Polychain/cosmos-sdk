package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

type Validator = stakingtypes.ValidatorI

type DistributionHooks interface {
	AfterAllocateTokensToValidator(ctx sdk.Context, val Validator, commission sdk.DecCoins, shared sdk.DecCoins)
}
