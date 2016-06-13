package gofc

import (
	"fmt"
	"net"
	"bytes"
	"../ofprotocol/ofp13"
)

// datapath
type Datapath struct{
	buffer chan *bytes.Buffer
	conn *net.TCPConn
	datapathId string
	// recvBuffer chan *ofp13.OFMessage
	sendBuffer chan *ofp13.OFMessage
	ofpversion string;
	ports int;
}


/**
 * ctor
 */
func NewDatapath(conn *net.TCPConn) *Datapath{
	dp := new(Datapath)
	dp.sendBuffer = make(chan *ofp13.OFMessage, 10)
	dp.conn = conn
	go dp.recvLoop()
	go dp.sendLoop()
	return dp
}



func (dp *Datapath)sendLoop() {
	for{
		select{
			// wait channel
			case msg := <- (dp.sendBuffer):
				// serialize data
				byteData := (*msg).Serialize()
				_, err := dp.conn.Write(byteData)
				if err != nil {
					fmt.Println("err to write conn")
					fmt.Println(err)
				}
			default:
		}
	}
}

func (dp *Datapath)recvLoop(){
	buf := make([]byte, 2048)
	for{
		// read
		dp.conn.Read(buf)

		// parse data
		// msg := ofp13.Parse(buf)

		// TODO:dispatch handler
	}
}


/**
 *
 */
func (dp *Datapath)send(message ofp13.OFMessage) bool{
	// push data
	(dp.sendBuffer) <- &message
	return true;
}

