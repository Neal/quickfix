//Package indicationofinterest msg type = 6.
package indicationofinterest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoIOIQualifiers is a repeating group in IndicationofInterest
type NoIOIQualifiers struct {
	//IOIQualifier is a non-required field for NoIOIQualifiers.
	IOIQualifier *string `fix:"104"`
}

//NoRoutingIDs is a repeating group in IndicationofInterest
type NoRoutingIDs struct {
	//RoutingType is a non-required field for NoRoutingIDs.
	RoutingType *int `fix:"216"`
	//RoutingID is a non-required field for NoRoutingIDs.
	RoutingID *string `fix:"217"`
}

//Message is a IndicationofInterest FIX Message
type Message struct {
	FIXMsgType string `fix:"6"`
	Header     fix42.Header
	//IOIid is a required field for IndicationofInterest.
	IOIid string `fix:"23"`
	//IOITransType is a required field for IndicationofInterest.
	IOITransType string `fix:"28"`
	//IOIRefID is a non-required field for IndicationofInterest.
	IOIRefID *string `fix:"26"`
	//Symbol is a required field for IndicationofInterest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for IndicationofInterest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for IndicationofInterest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for IndicationofInterest.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for IndicationofInterest.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for IndicationofInterest.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for IndicationofInterest.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for IndicationofInterest.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for IndicationofInterest.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for IndicationofInterest.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for IndicationofInterest.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for IndicationofInterest.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for IndicationofInterest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for IndicationofInterest.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for IndicationofInterest.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for IndicationofInterest.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for IndicationofInterest.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for IndicationofInterest.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for IndicationofInterest.
	EncodedSecurityDesc *string `fix:"351"`
	//Side is a required field for IndicationofInterest.
	Side string `fix:"54"`
	//IOIShares is a required field for IndicationofInterest.
	IOIShares string `fix:"27"`
	//Price is a non-required field for IndicationofInterest.
	Price *float64 `fix:"44"`
	//Currency is a non-required field for IndicationofInterest.
	Currency *string `fix:"15"`
	//ValidUntilTime is a non-required field for IndicationofInterest.
	ValidUntilTime *time.Time `fix:"62"`
	//IOIQltyInd is a non-required field for IndicationofInterest.
	IOIQltyInd *string `fix:"25"`
	//IOINaturalFlag is a non-required field for IndicationofInterest.
	IOINaturalFlag *bool `fix:"130"`
	//NoIOIQualifiers is a non-required field for IndicationofInterest.
	NoIOIQualifiers []NoIOIQualifiers `fix:"199,omitempty"`
	//Text is a non-required field for IndicationofInterest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for IndicationofInterest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for IndicationofInterest.
	EncodedText *string `fix:"355"`
	//TransactTime is a non-required field for IndicationofInterest.
	TransactTime *time.Time `fix:"60"`
	//URLLink is a non-required field for IndicationofInterest.
	URLLink *string `fix:"149"`
	//NoRoutingIDs is a non-required field for IndicationofInterest.
	NoRoutingIDs []NoRoutingIDs `fix:"215,omitempty"`
	//SpreadToBenchmark is a non-required field for IndicationofInterest.
	SpreadToBenchmark *float64 `fix:"218"`
	//Benchmark is a non-required field for IndicationofInterest.
	Benchmark *string `fix:"219"`
	Trailer   fix42.Trailer
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
	return enum.BeginStringFIX42, "6", r
}
