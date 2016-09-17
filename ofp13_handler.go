package gofc

import (
	"github.com/Kmotiko/gofc/ofprotocol/ofp13"
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
/* OfpDescStatsReply                                 */
/*****************************************************/
type Of13DescStatsReplyHandler interface {
	HandleDescStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpFlowStatsReply                                 */
/*****************************************************/
type Of13FlowStatsReplyHandler interface {
	HandleFlowStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpAggregateStatsReply                            */
/*****************************************************/
type Of13AggregateStatsReplyHandler interface {
	HandleAggregateStatsReply(*ofp13.OfpMultipartReply, *Datapath)
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
