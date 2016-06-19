package gofc

import (
	"net"
	"fmt"
	"../ofprotocol/ofp13"
)

/**
 * basic controller
 */
type OFController struct{
	echoInterval int32 // echo interval
}

func NewOFController() *OFController{
	ofc := new(OFController)
	ofc.echoInterval = 60
	return ofc
}

func (c *OFController)HandleHello(msg *ofp13.OfpHello, dp *Datapath){
	fmt.Println("recv Hello")
	// send feature request
	featureReq := ofp13.NewOfpFeaturesRequest()
	dp.Send(featureReq)
}


func (c *OFController)HandleSwitchFeatures(msg *ofp13.OfpSwitchFeatures, dp *Datapath){
	fmt.Println("recv SwitchFeatures")
	// handle FeatureReply
	dp.datapathId = msg.DatapathId;
}

func (c *OFController)HandleEchoRequest(msg *ofp13.OfpHeader, dp *Datapath){
	fmt.Println("recv EchoReq")
	// send EchoReply
	echo := ofp13.NewOfpEchoReply()
	dp.Send(echo)
}

func (c *OFController)ConnectionUp(){
	// handle connection up
}

func (c *OFController)ConnectionDown(){
	// handle connection down
}

func (c *OFController)sendEchoLoop(){
	// send echo request forever
}



func ServerLoop(){
	serverStr := ":6633"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverStr)
	listener, err := net.ListenTCP("tcp", tcpAddr)

	ofc := NewOFController()
	GetAppManager().RegistApplication(ofc)

	if err != nil {
		return
	}

	// wait for connect from switch
	for{
		conn,err := listener.AcceptTCP()
		if err != nil {
			return
		}
		go handleConnection(conn)
	}
}


/**
 *
 */
func handleConnection(conn *net.TCPConn){
	// send hello
	hello := ofp13.NewOfpHello()
	_, err := conn.Write(hello.Serialize())
	if err != nil{
		fmt.Println(err)
	}

	// create datapath
	dp := NewDatapath(conn)
	go dp.recvLoop()
	go dp.sendLoop()
}
