package gofc

import (
	"github.com/Kmotiko/gofc/ofprotocol/ofp13"
)

/*****************************************************/
/* OfpErrorMsg                                       */
/*****************************************************/
type Of13ErrorMsgHandler interface {
	HandleErrorMsg(*ofp13.OfpErrorMsg, *Datapath)
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

/*****************************************************/
/* BarrierReply Message                              */
/*****************************************************/
type Of13BarrierReplyHandler interface {
	HandleBarrierReply(*ofp13.OfpHeader, *Datapath)
}

/*****************************************************/
/* OfpSwitchFeatures                                 */
/*****************************************************/
type Of13SwitchFeaturesHandler interface {
	HandleSwitchFeatures(*ofp13.OfpSwitchFeatures, *Datapath)
}

/*****************************************************/
/* OfpSwitchConfig                                   */
/*****************************************************/
type Of13SwitchConfigHandler interface {
	HandleSwitchConfig(*ofp13.OfpSwitchConfig, *Datapath)
}

/*****************************************************/
/* OfpPacketIn                                       */
/*****************************************************/
type Of13PacketInHandler interface {
	HandlePacketIn(*ofp13.OfpPacketIn, *Datapath)
}

/*****************************************************/
/* OfpFlowRemoved                                    */
/*****************************************************/
type Of13FlowRemovedHandler interface {
	HandleFlowRemoved(*ofp13.OfpFlowRemoved, *Datapath)
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
/* OfpTableStatsReply                                */
/*****************************************************/
type Of13TableStatsReplyHandler interface {
	HandleTableStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpPortStatsReply                                 */
/*****************************************************/
type Of13PortStatsReplyHandler interface {
	HandlePortStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpQueueStatsReply                                */
/*****************************************************/
type Of13QueueStatsReplyHandler interface {
	HandleQueueStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpGroupStatsReply                                */
/*****************************************************/
type Of13GroupStatsReplyHandler interface {
	HandleGroupStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpGroupDescStatsReply                            */
/*****************************************************/
type Of13GroupDescStatsReplyHandler interface {
	HandleGroupDescStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpGroupFeaturesStatsReply                        */
/*****************************************************/
type Of13GroupFeaturesStatsReplyHandler interface {
	HandleGroupFeaturesStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpMeterStatsReply                                */
/*****************************************************/
type Of13MeterStatsReplyHandler interface {
	HandleMeterStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpMeterConfigStatsReply                          */
/*****************************************************/
type Of13MeterConfigStatsReplyHandler interface {
	HandleMeterConfigStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpMeterFeaturesStatsReply                        */
/*****************************************************/
type Of13MeterFeaturesStatsReplyHandler interface {
	HandleMeterFeaturesStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpTableFeaturesStatsReply                        */
/*****************************************************/
type Of13TableFeaturesStatsReplyHandler interface {
	HandleTableFeaturesStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* OfpPortDescStatsReply                             */
/*****************************************************/
type Of13PortDescStatsReplyHandler interface {
	HandlePortDescStatsReply(*ofp13.OfpMultipartReply, *Datapath)
}

/*****************************************************/
/* RoleReply Message                                 */
/*****************************************************/
type Of13RoleReplyHandler interface {
	HandleRoleReply(*ofp13.OfpRole, *Datapath)
}

/*****************************************************/
/* GetAsyncReply Message                             */
/*****************************************************/
type Of13AsyncConfigHandler interface {
	HandleAsyncConfig(*ofp13.OfpAsyncConfig, *Datapath)
}
