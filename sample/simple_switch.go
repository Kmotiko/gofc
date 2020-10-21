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
	hosts map[string]Host
	sync.RWMutex
}

func NewHostMap() *HostMap {
	h := new(HostMap)
	h.hosts = make(map[string]Host)
	return h
}

func (m *HostMap) Host(mac net.HardwareAddr) (h Host, ok bool) {
	m.RLock()
	defer m.RUnlock()
	h, ok = m.hosts[mac.String()]
	return
}


func (m *HostMap) SetHost(mac net.HardwareAddr, port uint16) {
	m.Lock()
	defer m.Unlock()
	m.hosts[mac.String()] = Host{mac, port}
}

var hostMap = *NewHostMap()

type SimpleSwitch struct {
	*HostMap
}

func NewSimpleSwitch() *SimpleSwitch {
	return &SimpleSwitch{&hostMap}
}

func (l2sw *SimpleSwitch) HandlePacketIn(msg *ofp13.OfpPacketIn, dp *gofc.Datapath) {

	//log.Printf("msg DataEth:%+v", msg.DataEth)
	eth := msg.DataEth
	// Ignore link discovery packet types.
	if eth.Ethertype == 0xa0f1 || eth.Ethertype == 0x88cc {
		log.Printf("link discovery packet:%v",eth)
		return
	}

	oxm := msg.Match.OxmFields[0]
	if oxm.OxmField() == ofp13.OFPXMT_OFB_IN_PORT {
		if OxmInPort,ok := oxm.(*ofp13.OxmInPort);ok {
			log.Printf("SetHost mac:%v with port:%d",eth.HWSrc,OxmInPort.Value)
			l2sw.SetHost(eth.HWSrc, uint16(OxmInPort.Value))
			if host, ok := l2sw.Host(eth.HWDst); ok {
				log.Printf("found dest:%v in port:%d",host.mac,host.port)

				eth_src,_ := ofp13.NewOxmEthSrc(eth.HWSrc.String())
				eth_dst,_ := ofp13.NewOxmEthDst(eth.HWDst.String())
				match_src_dst := ofp13.NewOfpMatch()
				match_src_dst.Append(eth_src)
				match_src_dst.Append(eth_dst)
				instruction := ofp13.NewOfpInstructionActions(ofp13.OFPIT_APPLY_ACTIONS)
				instruction.Append(ofp13.NewOfpActionOutput(uint32(host.port), 0))
				instructions := make([]ofp13.OfpInstruction, 0)
				instructions = append(instructions, instruction)

				fmod := ofp13.NewOfpFlowModAdd(
					0,
					0,
					0,
					1,
					ofp13.OFPFF_SEND_FLOW_REM,
					match_src_dst,
					instructions,
				)
				dp.Send(fmod)
				log.Printf("add flow mod:%+v",fmod)

				eth_src1,_ := ofp13.NewOxmEthSrc(eth.HWDst.String())
				eth_dst1,_ := ofp13.NewOxmEthDst(eth.HWSrc.String())
				match_dst_src := ofp13.NewOfpMatch()
				match_dst_src.Append(eth_src1)
				match_dst_src.Append(eth_dst1)
				instruction1 := ofp13.NewOfpInstructionActions(ofp13.OFPIT_APPLY_ACTIONS)
				instruction1.Append(ofp13.NewOfpActionOutput(OxmInPort.Value, 0))
				instructions1 := make([]ofp13.OfpInstruction, 0)
				instructions1 = append(instructions1, instruction1)

				fmod1 := ofp13.NewOfpFlowModAdd(
					0,
					0,
					0,
					1,
					ofp13.OFPFF_SEND_FLOW_REM,
					match_dst_src,
					instructions1,
				)
				dp.Send(fmod1)
				log.Printf("add flow mod:%+v",fmod)
			} else {
				actions := make([]ofp13.OfpAction, 0)
				action := ofp13.NewOfpActionOutput(ofp13.OFPP_FLOOD, 0)
				actions = append(actions, action)
				m := ofp13.NewOfpPacketOut(0xffffffff, OxmInPort.Value, actions, msg.Data)
				dp.Send(m)
				log.Printf("inport:%d flood end",OxmInPort.Value)
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
