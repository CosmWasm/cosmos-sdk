package types

// Staking module event types
var (
	EvenTypeCompleteUnbonding     = "complete_unbonding"
	EventTypeCompleteRedelegation = "complete_redelegation"
	EventTypeCreateValidator      = "create_validator"
	EventTypeEditValidator        = "edit_validator"
	EventTypeDelegate             = "delegate"
	EventTypeUnbond               = "unbond"
	EventTypeRedelegate           = "redelegate"

	AttributeKeyValidator      = "validator"
	AttributeKeySrcValidator   = "source_validator"
	AttributeKeyDstValidator   = "destination_validator"
	AttributeKeyDelegator      = "delegator"
	AttributeKeyAmount         = "amount"
	AttributeKeyCompletionTime = "completion_time"
	AttributeValueCategory     = "staking"
)
