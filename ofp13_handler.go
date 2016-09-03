package gofc

import (
	"./ofprotocol/ofp13"
)

/*****************************************************/
/* OfpSwitchFeatures                                 */
/*****************************************************/
type Of13SwitchFeaturesHandler interface {
	HandleSwitchFeatures(*ofp13.OfpSwitchFeatures, *Datapath)
}

/*****************************************************/
/* OfpPacketIn                                       */
/*****************************************************/
type Of13PacketInHandler interface {
	HandlePacketIn(*ofp13.OfpPacketIn, *Datapath)
}

/*****************************************************/
/* OfpErrorMsg                                       */
/*****************************************************/
type Of13ErroMsgHandler interface {
	HandleErroMsg(*ofp13.OfpErrorMsg, *Datapath)
}

/*****************************************************/
/* Echo Message                                      */
/*****************************************************/
type Of13EchoRequestHandler interface {
	HandleEchoRequest(*ofp13.OfpHeader, *Datapath)
}

type Of13EchoReplyHandler interface {
	HandleEchoReply(*ofp13.OfpHeader, *Datapath)
}
