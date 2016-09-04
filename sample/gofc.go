package main

import (
	"fmt"
	"github.com/Kmotiko/gofc"
	"github.com/Kmotiko/gofc/ofprotocol/ofp13"
)

type SampleController struct {
	// add any paramter used in controller.
}

func NewSampleController() *SampleController {
	ofc := new(SampleController)
	return ofc
}

func (c *SampleController) HandleSwitchFeatures(msg *ofp13.OfpSwitchFeatures, dp *gofc.Datapath) {
	// create flow mod
	fm := ofp13.NewOfpFlowMod()

	// create match
	ethdst, _ := ofp13.NewOxmEthDst("00:00:00:00:00:00")
	if ethdst == nil {
		fmt.Println(ethdst)
		return
	}
	fm.AppendMatchField(ethdst)

	// create Instruction
	instruction := ofp13.NewOfpInstructionActions(ofp13.OFPIT_APPLY_ACTIONS)

	// create actions
	instruction.Append(ofp13.NewOfpActionOutput(0, 0))

	// append Instruction
	fm.AppendInstruction(instruction)

	dp.Send(fm)
}

func main() {
	// regist app
	ofc := NewSampleController()
	gofc.GetAppManager().RegistApplication(ofc)

	// start server
	gofc.ServerLoop()
}
