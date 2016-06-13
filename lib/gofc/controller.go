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


func ServerLoop(){
	serverStr := ":6633"
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverStr)
	listener, err := net.ListenTCP("tcp", tcpAddr)

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
	defer conn.Close()

	// handle initialize sequence
	// send hello
	hello := ofp13.NewOfpHello()
	_, err := conn.Write(hello.Serialize())
	if err != nil{
		fmt.Println(err)
	}

	// make buffer and read
	buf := make([]byte, 1024)
	// switch message type
	for{
		fmt.Println("read message")
		_, err = conn.Read(buf)
		if err != nil {
			return
		}
		msg := ofp13.Parse(buf)
		switch msg.(type) {
		// case Hello
		case *ofp13.OfpHello:
			fmt.Println("recv Hello")
			// send feature request
			featureReq := ofp13.NewOfpFeaturesRequest()
			_, err = conn.Write(featureReq.Serialize())
			if err != nil{
				fmt.Println(err)
			}
		// case SwitchFeatures
		case *ofp13.OfpSwitchFeatures:
			fmt.Println("recv SwitchFeatures")
			// create datapath
			dp := NewDatapath(conn)
			// create and send echo
			echo := ofp13.NewOfpEchoRequest()
			dp.send(echo)
			// TODO:regist datapath
			break
		// Recv Error
		case *ofp13.OfpErrorMsg:
			fmt.Println("ErrMsg")
			break
		default:
			fmt.Println("UnSupport Message")
			break
		}
	}
}
