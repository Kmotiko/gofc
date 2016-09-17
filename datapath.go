package gofc

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/Kmotiko/gofc/ofprotocol/ofp13"
	"net"
)

// datapath
type Datapath struct {
	buffer     chan *bytes.Buffer
	conn       *net.TCPConn
	datapathId uint64
	sendBuffer chan *ofp13.OFMessage
	ofpversion string
	ports      int
}

/**
 * ctor
 */
func NewDatapath(conn *net.TCPConn) *Datapath {
	dp := new(Datapath)
	dp.sendBuffer = make(chan *ofp13.OFMessage, 10)
	dp.conn = conn
	go dp.recvLoop()
	go dp.sendLoop()
	return dp
}

func (dp *Datapath) sendLoop() {
	for {
		select {
		// wait channel
		case msg := <-(dp.sendBuffer):
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

func (dp *Datapath) recvLoop() {
	buf := make([]byte, 1024*64)
	for {
		// read
		size, err := dp.conn.Read(buf)
		if err != nil {
			fmt.Println("failed to read conn")
			fmt.Println(err)
			return
		}

		// tmp := make([]byte, 2048)
		for i := 0; i < size; {
			msgLen := binary.BigEndian.Uint16(buf[i+2:])
			dp.handlePacket(buf[i : i+(int)(msgLen)])
			i += (int)(msgLen)
		}
	}
}

func (dp *Datapath) handlePacket(buf []byte) {
	// parse data
	msg := ofp13.Parse(buf[0:])
	apps := GetAppManager().GetApplications()

	// handle hello
	if _, ok := msg.(*ofp13.OfpHello); ok {
		featureReq := ofp13.NewOfpFeaturesRequest()
		dp.Send(featureReq)
	}

	// dispatch handler
	for _, app := range apps {
		switch msgi := msg.(type) {
		// if message is OfpHeader
		case *ofp13.OfpHeader:
			switch msgi.Type {
			// handle echo request
			case ofp13.OFPT_ECHO_REQUEST:
				if obj, ok := app.(Of13EchoRequestHandler); ok {
					obj.HandleEchoRequest(msgi, dp)
				}

			// handle echo reply
			case ofp13.OFPT_ECHO_REPLY:
				if obj, ok := app.(Of13EchoReplyHandler); ok {
					obj.HandleEchoReply(msgi, dp)
				}
			default:
			}

		// case SwitchFeatures
		case *ofp13.OfpSwitchFeatures:
			if obj, ok := app.(Of13SwitchFeaturesHandler); ok {
				obj.HandleSwitchFeatures(msgi, dp)
			}

		// case PacketIn
		case *ofp13.OfpPacketIn:
			if obj, ok := app.(Of13PacketInHandler); ok {
				obj.HandlePacketIn(msgi, dp)
			}

		// case MultipartReply
		case *ofp13.OfpMultipartReply:
			switch msgi.Type {
			case ofp13.OFPMP_DESC:
				if obj, ok := app.(Of13DescStatsReplyHandler); ok {
					obj.HandleDescStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_FLOW:
				if obj, ok := app.(Of13FlowStatsReplyHandler); ok {
					obj.HandleFlowStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_AGGREGATE:
				if obj, ok := app.(Of13AggregateStatsReplyHandler); ok {
					obj.HandleAggregateStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_TABLE:
				// TODO: implement
			case ofp13.OFPMP_PORT_STATS:
				// TODO: implement
			case ofp13.OFPMP_QUEUE:
				// TODO: implement
			case ofp13.OFPMP_GROUP:
				// TODO: implement
			case ofp13.OFPMP_GROUP_DESC:
				// TODO: implement
			case ofp13.OFPMP_GROUP_FEATURES:
				// TODO: implement
			case ofp13.OFPMP_METER:
				// TODO: implement
			case ofp13.OFPMP_METER_CONFIG:
				// TODO: implement
			case ofp13.OFPMP_METER_FEATURES:
				// TODO: implement
			case ofp13.OFPMP_TABLE_FEATURES:
				// TODO: implement
			case ofp13.OFPMP_PORT_DESC:
				// TODO: implement
			case ofp13.OFPMP_EXPERIMENTER:
				// TODO: implement
			default:
			}

		// Recv Error
		case *ofp13.OfpErrorMsg:
			fmt.Println("ErrMsg")

		default:
			fmt.Println("UnSupport Message")
		}
	}
}

/**
 *
 */
func (dp *Datapath) Send(message ofp13.OFMessage) bool {
	// push data
	(dp.sendBuffer) <- &message
	return true
}
