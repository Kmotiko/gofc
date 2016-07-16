package ofp13

/*****************************************************/
/* OfpHeader                                         */
/*****************************************************/
type HelloHandler interface{
	HandleHello(*ofp13.OfpHello, *Datapath)
}


/*****************************************************/
/* OfpSwitchFeatures                                 */
/*****************************************************/
type SwitchFeaturesHandler interface{
	HandleSwitchFeatures(*ofp13.OfpSwitchFeatures, *Datapath)
}


/*****************************************************/
/* OfpPacketIn                                       */
/*****************************************************/
type PacketInHandler interface{
	HandlePacketIn(*ofp13.OfpPacketIn, *Datapath)
}


/*****************************************************/
/* OfpErrorMsg                                       */
/*****************************************************/
type ErroMsgHandler interface{
	HandleErroMsg(*ofp13.OfpErrorMsg, *Datapath)
}


/*****************************************************/
/* Echo Message                                      */
/*****************************************************/
type EchoRequestHandler interface{
	HandleEchoRequest(*ofp13.OfpHeader, *Datapath)
}

type EchoReplyHandler interface{
	HandleEchoReply(*ofp13.OfpHeader, *Datapath)
}
