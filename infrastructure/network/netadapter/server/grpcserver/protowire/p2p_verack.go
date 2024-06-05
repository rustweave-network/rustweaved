package protowire

import (
	"github.com/rustweave-network/rustweaved/app/appmessage"
	"github.com/pkg/errors"
)

func (x *RustweavedMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RustweavedMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *RustweavedMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
