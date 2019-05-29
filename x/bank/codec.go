package bank

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSend{}, "cosmos-sdk/MsgSend", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "cosmos-sdk/MsgMultiSend", nil)
}

var msgCdc = codec.New()

func init() {
	RegisterCodec(msgCdc)
}

// SetMsgCodec allows sdk users use custom codex at GetSignBytes
func SetMsgCodec(cdc *codec.Codec) {
	msgCdc = cdc
}
