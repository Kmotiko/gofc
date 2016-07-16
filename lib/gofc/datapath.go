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
	datapathId uint64
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
					fmt.Println("failed to write conn")
					fmt.Println(err)
					return
				}
			default:
				continue
		}
	}
}

func (dp *Datapath)recvLoop(){
	buf := make([]byte, 2048)
	for{
		// read
		_,err := dp.conn.Read(buf)
		if err != nil {
			fmt.Println("failed to read conn")
			fmt.Println(err)
			return
		}

		// parse data
		msg := ofp13.Parse(buf)
		apps := GetAppManager().GetApplications()

		// dispatch handler
		for _,app := range apps {
			switch msgi := msg.(type) {
				// handle hello
			case *ofp13.OfpHello:
				if obj, ok := app.(ofp13.HelloHandler); ok{
					obj.HandleHello(msgi, dp)
				}

				// if message is OfpHeader
			case *ofp13.OfpHeader:
				switch msgi.Type {
					// handle echo request
				case ofp13.OFPT_ECHO_REQUEST:
					if obj, ok := app.(ofp13.EchoRequestHandler); ok{
						obj.HandleEchoRequest(msgi, dp)
					}

					// handle echo reply
				case ofp13.OFPT_ECHO_REPLY:
					if obj, ok := app.(ofp13.EchoReplyHandler); ok{
						obj.HandleEchoReply(msgi, dp)
					}
				default:
				}

			// case SwitchFeatures
			case *ofp13.OfpSwitchFeatures:
				if obj, ok := app.(ofp13.SwitchFeaturesHandler); ok{
					obj.HandleSwitchFeatures(msgi, dp)
				}

			// case PacketIn
			case *ofp13.OfpPacketIn:
				if obj, ok := app.(ofp13.PacketInHandler); ok{
					obj.HandlePacketIn(msgi, dp)
				}

			// Recv Error
			case *ofp13.OfpErrorMsg:
				fmt.Println("ErrMsg")

			default:
				fmt.Println("UnSupport Message")
			}
		}
	}
}


/**
 *
 */
func (dp *Datapath)Send(message ofp13.OFMessage) bool{
	// push data
	(dp.sendBuffer) <- &message
	return true;
}
