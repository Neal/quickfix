//Package settlementinstructions msg type = T.
package settlementinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/settlinstgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SettlementInstructions FIX Message
type Message struct {
	FIXMsgType string `fix:"T"`
	Header     fixt11.Header
	//SettlInstMsgID is a required field for SettlementInstructions.
	SettlInstMsgID string `fix:"777"`
	//SettlInstReqID is a non-required field for SettlementInstructions.
	SettlInstReqID *string `fix:"791"`
	//SettlInstMode is a required field for SettlementInstructions.
	SettlInstMode string `fix:"160"`
	//SettlInstReqRejCode is a non-required field for SettlementInstructions.
	SettlInstReqRejCode *int `fix:"792"`
	//Text is a non-required field for SettlementInstructions.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SettlementInstructions.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SettlementInstructions.
	EncodedText *string `fix:"355"`
	//ClOrdID is a non-required field for SettlementInstructions.
	ClOrdID *string `fix:"11"`
	//TransactTime is a required field for SettlementInstructions.
	TransactTime time.Time `fix:"60"`
	//SettlInstGrp Component
	SettlInstGrp settlinstgrp.Component
	Trailer      fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX50, "T", r
}
