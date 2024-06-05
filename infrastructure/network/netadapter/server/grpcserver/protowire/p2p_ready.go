package protowire

import (
	"github.com/rustweave-network/rustweaved/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KaspadMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RustweavedMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *KaspadMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
