package main

import (
	"github.com/Kmotiko/gofc"
	"github.com/Kmotiko/gofc/ofprotocol/ofp13"
	"log"
	"net"
	"sync"
)

// Structure to track hosts that we discover.
type Host struct {
	mac  net.HardwareAddr
	port uint16
}

// A thread safe map to store our hosts.
type HostMap struct {
	hosts map[uint64]map[string]Host
	sync.RWMutex
}

func NewHostMap() *HostMap {
	h := new(HostMap)
	h.hosts = make(map[uint64]map[string]Host)
	return h
}

func (m *HostMap) Host(dpid uint64, mac net.HardwareAddr) (h Host, ok bool) {
	m.RLock()
	defer m.RUnlock()
	h, ok = m.hosts[dpid][mac.String()]
	return
}

func (m *HostMap) SetHost(dpid uint64, mac net.HardwareAddr, port uint16) {
	m.Lock()
	defer m.Unlock()
	var hostmap map[string]Host
	if m.hosts[dpid] == nil {
		hostmap = make(map[string]Host)
	}else {
		hostmap = m.hosts[dpid]
	}
	hostmap[mac.String()] = Host{mac, port}
	m.hosts[dpid] =hostmap
}

var hostMap = *NewHostMap()

type SimpleSwitch struct {
	*HostMap
}

func NewSimpleSwitch() *SimpleSwitch {
	return &SimpleSwitch{&hostMap}
}

func (l2sw *SimpleSwitch) add_flow(dp *gofc.Datapath, priority uint16, match *ofp13.OfpMatch, instructions []ofp13.OfpInstruction) {
	fmod := ofp13.NewOfpFlowModAdd(
		0,
		0,
		0,
		priority,
		ofp13.OFPFF_SEND_FLOW_REM,
		match,
		instructions,
	)
	dp.Send(fmod)
}

func (l2sw *SimpleSwitch) HandlePacketIn(msg *ofp13.OfpPacketIn, dp *gofc.Datapath) {

	eth := msg.DataEth
	// Ignore link discovery packet types. or ipv6
	if eth.Ethertype == 0x88cc || eth.Ethertype == 0x86dd {
		return
	}

	oxm := msg.Match.OxmFields[0]
	if oxm.OxmField() == ofp13.OFPXMT_OFB_IN_PORT {
		if OxmInPort, ok := oxm.(*ofp13.OxmInPort); ok {
			dpid := dp.GetDatapathId()
			log.Printf("packet in dpid:%d src:%v dst:%v in_port:%d", dp.GetDatapathId(), eth.HWSrc, eth.HWDst, OxmInPort.Value)
			l2sw.SetHost(dpid, eth.HWSrc, uint16(OxmInPort.Value))
			if host, ok := l2sw.Host(dpid, eth.HWDst); ok {
				log.Printf("dpid:%d  found dest:%v in port:%d", dpid, host.mac, host.port)

				eth_src, _ := ofp13.NewOxmEthDst(eth.HWSrc.String())
				match_src := ofp13.NewOfpMatch()
				match_src.Append(eth_src)
				instruction := ofp13.NewOfpInstructionActions(ofp13.OFPIT_APPLY_ACTIONS)
				instruction.Append(ofp13.NewOfpActionOutput(OxmInPort.Value, 0))
				instructions := make([]ofp13.OfpInstruction, 0)
				instructions = append(instructions, instruction)
				l2sw.add_flow(dp, 1, match_src, instructions)

				eth_dst, _ := ofp13.NewOxmEthDst(eth.HWDst.String())
				match_dst := ofp13.NewOfpMatch()
				match_dst.Append(eth_dst)
				instruction1 := ofp13.NewOfpInstructionActions(ofp13.OFPIT_APPLY_ACTIONS)
				instruction1.Append(ofp13.NewOfpActionOutput(uint32(host.port), 0))
				instructions1 := make([]ofp13.OfpInstruction, 0)
				instructions1 = append(instructions1, instruction1)
				l2sw.add_flow(dp, 1, match_dst, instructions1)

			} else {
				actions := make([]ofp13.OfpAction, 0)
				action := ofp13.NewOfpActionOutput(ofp13.OFPP_FLOOD, 0)
				actions = append(actions, action)
				bufferid := uint32(ofp13.OFP_NO_BUFFER)
				data := msg.Data
				if msg.BufferId != ofp13.OFP_NO_BUFFER {
					bufferid = msg.BufferId
					data = nil
				}
				m := ofp13.NewOfpPacketOut(bufferid, OxmInPort.Value, actions, data)
				dp.Send(m)
			}
		}

	}
}

func main() {
	// regist app
	ofc := NewSimpleSwitch()
	gofc.GetAppManager().RegistApplication(ofc)

	// start server
	gofc.ServerLoop(gofc.DEFAULT_PORT)
}
