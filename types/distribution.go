package types

 type DistributionHooks interface {
	AfterAllocateTokensToValidator(ctx Context, val Validator, commission DecCoins, shared DecCoins)
	TruncateDelegatorRewards(ctx Context, val Validator, del Delegation, remainder DecCoins)
	TruncateValidatorCommission(ctx Context, valAddr ValAddress, remainder DecCoins)
}