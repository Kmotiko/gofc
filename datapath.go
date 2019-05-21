package gofc

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"

	"github.com/nutanix/gofc/ofprotocol/ofp13"
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
	return dp
}

func (dp *Datapath) sendLoop() {
	for {
		// wait channel
		msg := <-(dp.sendBuffer)
		// serialize data
		byteData := (*msg).Serialize()
		_, err := dp.conn.Write(byteData)
		if err != nil {
			fmt.Println("failed to write conn")
			fmt.Println(err)
			return
		}
	}
}

func (dp *Datapath) recvLoop() {
	// for more information see https://www.opennetworking.org/wp-content/uploads/2014/10/openflow-spec-v1.3.0.pdf
	const kHeaderSize = 8
	reader := bufio.NewReader(dp.conn)
	buf := make([]byte, 1024*64)
	for {
		for i := 0; i < kHeaderSize; i++ {
			b, err := reader.ReadByte() // read len prefix and len
			if err != nil {
				fmt.Println("failed to read conn")
				fmt.Println(err)
				return
			}
			buf[i] = b
		}
		msgLen := (int)(binary.BigEndian.Uint16(buf[2:]))
		bufPos := kHeaderSize
		for bufPos < msgLen {
			read, err := reader.Read(buf[bufPos:msgLen])
			if err != nil {
				fmt.Println("failed to read conn")
				fmt.Println(err)
				return
			}
			bufPos += read
		}
		if bufPos != msgLen {
			fmt.Printf("Strange ofp packet, len in packet %d, received len %d\n", msgLen, bufPos)
		}
		dp.handlePacket(buf[0:bufPos])
	}
}

func (dp *Datapath) handlePacket(buf []byte) {
	// parse data
	msg := ofp13.Parse(buf[0:])

	if _, ok := msg.(*ofp13.OfpHello); ok {
		// handle hello
		featureReq := ofp13.NewOfpFeaturesRequest()
		dp.Send(featureReq)
	} else {
		// dispatch handler
		dp.dispatchHandler(msg)
	}
}

func (dp *Datapath) dispatchHandler(msg ofp13.OFMessage) {
	apps := GetAppManager().GetApplications()
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

			// handle Barrier reply
			case ofp13.OFPT_BARRIER_REPLY:
				if obj, ok := app.(Of13BarrierReplyHandler); ok {
					obj.HandleBarrierReply(msgi, dp)
				}
			default:
			}

		// Recv Error
		case *ofp13.OfpErrorMsg:
			if obj, ok := app.(Of13ErrorMsgHandler); ok {
				obj.HandleErrorMsg(msgi, dp)
			}

		// Recv PortStatus
		case *ofp13.OfpPortStatus:
			if obj, ok := app.(Of13PortStatusHandler); ok {
				obj.HandlePortStatus(msgi, dp)
			}

		// Recv RoleReply
		case *ofp13.OfpRole:
			if obj, ok := app.(Of13RoleReplyHandler); ok {
				obj.HandleRoleReply(msgi, dp)
			}

		// Recv GetAsyncReply
		case *ofp13.OfpAsyncConfig:
			if obj, ok := app.(Of13AsyncConfigHandler); ok {
				obj.HandleAsyncConfig(msgi, dp)
			}

		// case SwitchFeatures
		case *ofp13.OfpSwitchFeatures:
			if obj, ok := app.(Of13SwitchFeaturesHandler); ok {
				obj.HandleSwitchFeatures(msgi, dp)
			}

		// case GetConfigReply
		case *ofp13.OfpSwitchConfig:
			if obj, ok := app.(Of13SwitchConfigHandler); ok {
				obj.HandleSwitchConfig(msgi, dp)
			}
		// case PacketIn
		case *ofp13.OfpPacketIn:
			if obj, ok := app.(Of13PacketInHandler); ok {
				obj.HandlePacketIn(msgi, dp)
			}

		// case FlowRemoved
		case *ofp13.OfpFlowRemoved:
			if obj, ok := app.(Of13FlowRemovedHandler); ok {
				obj.HandleFlowRemoved(msgi, dp)
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
				if obj, ok := app.(Of13TableStatsReplyHandler); ok {
					obj.HandleTableStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_PORT_STATS:
				if obj, ok := app.(Of13PortStatsReplyHandler); ok {
					obj.HandlePortStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_QUEUE:
				if obj, ok := app.(Of13QueueStatsReplyHandler); ok {
					obj.HandleQueueStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_GROUP:
				if obj, ok := app.(Of13GroupStatsReplyHandler); ok {
					obj.HandleGroupStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_GROUP_DESC:
				if obj, ok := app.(Of13GroupDescStatsReplyHandler); ok {
					obj.HandleGroupDescStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_GROUP_FEATURES:
				if obj, ok := app.(Of13GroupFeaturesStatsReplyHandler); ok {
					obj.HandleGroupFeaturesStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_METER:
				if obj, ok := app.(Of13MeterStatsReplyHandler); ok {
					obj.HandleMeterStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_METER_CONFIG:
				if obj, ok := app.(Of13MeterConfigStatsReplyHandler); ok {
					obj.HandleMeterConfigStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_METER_FEATURES:
				if obj, ok := app.(Of13MeterFeaturesStatsReplyHandler); ok {
					obj.HandleMeterFeaturesStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_TABLE_FEATURES:
				if obj, ok := app.(Of13TableFeaturesStatsReplyHandler); ok {
					obj.HandleTableFeaturesStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_PORT_DESC:
				if obj, ok := app.(Of13PortDescStatsReplyHandler); ok {
					obj.HandlePortDescStatsReply(msgi, dp)
				}
			case ofp13.OFPMP_EXPERIMENTER:
				// TODO: implement
			default:
			}

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
