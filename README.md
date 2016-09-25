gofc
==================

[![Build Status](https://travis-ci.org/Kmotiko/gofc.svg?branch=master)](https://travis-ci.org/Kmotiko/gofc)


## What is this?

OpenFlow Controller written in golang.
Now, only support OpenFlow 1.3.

## How to use

To run OpenFlow Controller with gofc, you have to follow the steps outlined bellow.

 * Implement Handler
 * Regist and Run Controller


### Implement Handler
First, Implement handler interface.
For example, if you want to handle PacketIn message, you should implement HandlePacketIn like bellow.

```
import (
	"github.com/Kmotiko/gofc"
	"github.com/Kmotiko/gofc/ofprotocol/ofp13"
}

type SampleController struct {
	// add any paramter used in controller.
}

func (c *SampleController) HandlePacketIn(msg *ofp13.OfpPacketIn, dp *gofc.Datapath) {
	// create flow mod
	fm := ofp13.NewOfpFlowMod(
		0,
		0,
		0,
		ofp13.OFPFC_ADD,
		0,
		0,
		0,
		ofp13.OFP_NO_BUFFER,
		ofp13.OFPP_ANY,
		ofp13.OFPG_ANY,
		ofp13.OFPFF_SEND_FLOW_REM)

	// TODO: set match field and instructions.

	// send FlowMod message
	dp.Send(fm)
}
```

Defenition of gofc interfaces is written in ofp13_handler.go.



### Regist and Run Controller

If you complete to implement of handlers, you can regist that to gofc and start it.

```
func main() {
	// regist app
	ofc := NewSampleController()
	gofc.GetAppManager().RegistApplication(ofc)

	// start Controller
	gofc.ServerLoop()
}
```

## OpenFlow Messages Support Status

### Messages

 - [x] Hello
 - [ ] SwitchConfig
 - [x] EchoRequest
 - [x] EchoReply
 - [x] SwitchFeatures
 - [x] PacketIn
 - [ ] TableMod
 - [ ] PortMod
 - [x] FlowMod(Partial Support)
 - [ ] GroupMod
 - [x] MeterMod
 - [ ] PacketOut
 - [ ] FlowRemoved
 - [ ] OfpErroMsg
 - [ ] OfpExperimenterMsg
 - [x] OfpMultipartRequest
 - [x] OfpMultipartReply


### Match Field

 - [ ] In Port
 - [ ] In Phy Port
 - [ ] Metadata
 - [x] Eth Dst
 - [ ] Eth Src
 - [ ] Eth Type
 - [ ] Vlan ID
 - [ ] Vlan PCP
 - [ ] IP Dscp
 - [ ] IP ECN
 - [ ] IP Proto
 - [ ] IPv4 Src
 - [ ] IPv4 Dst
 - [ ] TCP Src
 - [ ] TCP Dst
 - [ ] UDP Src
 - [ ] UDP Dst
 - [ ] SCTP Src
 - [ ] SCTP Dst
 - [ ] ICMPv4 Type
 - [ ] ICMPv4 Code
 - [ ] Arp Op
 - [ ] Arp SPA
 - [ ] Arp TPA
 - [ ] Arp SHA
 - [ ] Arp THA
 - [ ] IPv6 Src
 - [ ] IPv6 Dst
 - [ ] IPv6 FLabel
 - [ ] IPv6 Type
 - [ ] IPv6 Code
 - [ ] IPv6 ND Target
 - [ ] IPv6 ND SLL
 - [ ] IPv6 ND TLL
 - [ ] MPLS Label
 - [ ] MPLS TC
 - [ ] MPLS BOS
 - [ ] PBB ISID
 - [ ] TUNNEL ID
 - [ ] IPv6 ExtraHeader

### Instructions

 - [ ] Go to table
 - [ ] Write Metadata
 - [x] Action
 - [ ] Meter

### Actions

 - [x] Output
 - [ ] Copy TTL OUT
 - [ ] Copy TTL IN
 - [ ] Set MPLS TTL
 - [ ] Dec  MPLS TTL
 - [ ] PUSH VLAN
 - [ ] PUSH MPLS
 - [ ] PUSH PBB
 - [ ] POP VLAN
 - [ ] POP MPLS
 - [ ] POP PBB
 - [ ] Group
 - [ ] Set NW TTL
 - [ ] Dec NW TTL
 - [ ] Set Field
 - [ ] Experimenter Header

### MeterBand

 - [x] MeterBandDrop
 - [ ] MeterBandDscpRemark
 - [ ] MeterBandExperimenter

### Multipart Message

 - [x] DescStats
 - [x] FlowStats
 - [x] AggregateStats
 - [ ] TableStats
 - [ ] PortStats
 - [ ] QueueStats
 - [ ] GroupStats
 - [ ] GroupDescStats
 - [ ] GroupFeaturesStats
 - [ ] MeterStats
 - [ ] MeterConfigStats
 - [ ] MeterFeaturesStats
 - [ ] TableFeaturesStats
 - [ ] PortDescStats
 - [ ] ExperimenterStats

## License

This software is distributed with MIT license.

```
Copyright (c) 2016 Kmotiko
```
