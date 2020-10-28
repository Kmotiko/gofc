package app

import (
	"fmt"
	"github.com/Kmotiko/gofc"
	"github.com/Kmotiko/gofc/ofprotocol/ofp13"
	"log"
	"strings"
	"sync"
	"time"
)


type SimpleMonitor struct {
	*sync.RWMutex
	dpmap map[uint64]*gofc.Datapath
}

func (sm *SimpleMonitor) HandleSwitchFeatures(msg *ofp13.OfpSwitchFeatures, dp *gofc.Datapath) {
	log.Printf("add dpmap with key:%v",msg.DatapathId)
	sm.Lock()
	defer sm.Unlock()
	sm.dpmap[msg.DatapathId] = dp
}

func NewSimpleMonitor() *SimpleMonitor {
	sm := &SimpleMonitor{&sync.RWMutex{},make(map[uint64]*gofc.Datapath)}
	go sm.monitor()
	return sm
}

func (sm *SimpleMonitor)  monitor() {
	log.Printf("start monitor")
	mf := ofp13.NewOfpMatch()
	fs_req := ofp13.NewOfpFlowStatsRequest(0, 0, ofp13.OFPP_ANY, ofp13.OFPG_ANY, 0, 0, mf)
	ps_req := ofp13.NewOfpPortStatsRequest(ofp13.OFPP_ANY, 0)
	for {
		for dpid,dp := range sm.dpmap {
			log.Printf("sent flow stat request, dpid:%d",dpid)
			dp.Send(fs_req)
			dp.Send(ps_req)
		}
		time.Sleep(time.Second*10)
	}
}

func (sm *SimpleMonitor)  HandleFlowStatsReply(msg *ofp13.OfpMultipartReply, dp *gofc.Datapath) {
	//log.Printf("dpid:%d , flow stat reply:%+v",dp.GetDatapathId(),msg)
	for _,mr := range msg.Body {
		if fs, ok := mr.(*ofp13.OfpFlowStats);ok{
			matchfield := make([]string,0)
			matchfield = append(matchfield,fmt.Sprintf("table:%d",fs.TableId))
			matchfield = append(matchfield,fmt.Sprintf("duration:%d",fs.DurationSec))
			matchfield = append(matchfield,fmt.Sprintf("n_packets:%d",fs.PacketCount))
			matchfield = append(matchfield,fmt.Sprintf("n_bytes:%d",fs.ByteCount))
			matchfield = append(matchfield,fmt.Sprintf("priority:%d",fs.Priority))

			if len(fs.Match.OxmFields) != 0 {
				//only care about the rules which has match field we set

				for _,oxmf := range fs.Match.OxmFields {
					switch ioxmf := oxmf.(type) {
					case *ofp13.OxmEth:
						switch ioxmf.TlvHeader {
						case ofp13.OXM_OF_ETH_DST:
							matchfield = append(matchfield,fmt.Sprintf("dst_mac:%s",ioxmf.Value.String()))
						case ofp13.OXM_OF_ETH_SRC:
							matchfield = append(matchfield,fmt.Sprintf("src_mac:%s",ioxmf.Value.String()))
						}
					case *ofp13.OxmInPort:
						if ioxmf.TlvHeader == ofp13.OXM_OF_IN_PORT {
							matchfield = append(matchfield,fmt.Sprintf("in_port:%d",ioxmf.Value))
						}
					}
				}

				for _,instru := range fs.Instructions {
					switch iinstru := instru.(type) {
					case *ofp13.OfpInstructionActions:
						if iinstru.Header.Type == ofp13.OFPIT_APPLY_ACTIONS {
							for _,action := range iinstru.Actions {
								switch iaction :=action.(type) {
								case *ofp13.OfpActionOutput:
									if iaction.ActionHeader.Type == ofp13.OFPAT_OUTPUT{
										matchfield = append(matchfield,fmt.Sprintf("actions=output:%d",iaction.Port))
									}
								}
							}
						}
					}
				}
			}
			log.Printf("dpid[%d]>>> %s",dp.GetDatapathId(),strings.Join(matchfield," "))
		}
	}


}

func (sm *SimpleMonitor)  HandlePortStatsReply(msg *ofp13.OfpMultipartReply, dp *gofc.Datapath) {
	//log.Printf("dpid:%d ,port stat reply:%+v",dp.GetDatapathId(),msg)
	for _,mr := range msg.Body {
		if ps, ok := mr.(*ofp13.OfpPortStats); ok {
			log.Printf("dpid[%d]>>> %+v",dp.GetDatapathId(),ps)
		}
	}
}



