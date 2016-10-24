package ofp13

import (
	"encoding/hex"
	"math"
	"testing"
)

/*****************************************************/
/* OfpHello                                          */
/*****************************************************/
func TestSerializeHello(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x00,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	hello := NewOfpHello()
	actual := hello.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpHello is not equal to expected value.")
	}
}

func TestParseHello(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x00,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}

	hello := NewOfpHello()
	hello.Parse(packet)
	if hello.Header.Version != 4 || hello.Header.Type != 0 ||
		hello.Header.Length != 8 || hello.Header.Xid != 0 {
		t.Log("Version        : ", hello.Header.Version)
		t.Log("Type           : ", hello.Header.Type)
		t.Log("Length         : ", hello.Header.Length)
		t.Log("Transaction ID : ", hello.Header.Xid)
		t.Error("Parsed value of OfpHello is invalid.")
	}
}

/*****************************************************/
/* Echo Message                                      */
/*****************************************************/
func TestSerializeEchoRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x02,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	echo := NewOfpEchoRequest()
	actual := echo.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpEchoRequest is not equal to expected value.")
	}
}

func TestParseEchoRequest(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x02,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}

	echo := NewOfpEchoRequest()
	echo.Parse(packet)
	if echo.Version != 4 || echo.Type != 2 ||
		echo.Length != 8 || echo.Xid != 0 {
		t.Log("Version        : ", echo.Version)
		t.Log("Type           : ", echo.Type)
		t.Log("Length         : ", echo.Length)
		t.Log("Transaction ID : ", echo.Xid)
		t.Error("Parsed value of OfpEchoRequest is invalid.")
	}
}

func TestSerializeEchoReply(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x03,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	echo := NewOfpEchoReply()
	actual := echo.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpEchoReply is not equal to expected value.")
	}
}

func TestParseEchoReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x03,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}

	echo := NewOfpEchoReply()
	echo.Parse(packet)
	if echo.Version != 4 || echo.Type != 3 ||
		echo.Length != 8 || echo.Xid != 0 {
		t.Log("Version        : ", echo.Version)
		t.Log("Type           : ", echo.Type)
		t.Log("Length         : ", echo.Length)
		t.Log("Transaction ID : ", echo.Xid)
		t.Error("Parsed value of OfpEchoReply is invalid.")
	}
}

/*****************************************************/
/* Barrier Message                                   */
/*****************************************************/
func TestSerializeBarrierRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x14,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	m := NewOfpBarrierRequest()
	actual := m.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpBarrierRequest is not equal to expected value.")
	}
}

func TestParseBarrierReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x15,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}

	m := NewOfpBarrierReply()
	m.Parse(packet)
	if m.Version != 4 || m.Type != 21 ||
		m.Length != 8 || m.Xid != 0 {
		t.Log("Version        : ", m.Version)
		t.Log("Type           : ", m.Type)
		t.Log("Length         : ", m.Length)
		t.Log("Transaction ID : ", m.Xid)
		t.Error("Parsed value of OfpBarrierReply is invalid.")
	}
}

/*****************************************************/
/* OfpSwitchConfig                                   */
/*****************************************************/
func TestSerializeGetConfig(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x07,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	req := NewOfpGetConfig()
	actual := req.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpGetConfig is not equal to expected value.")
	}
}

func TestSerializeSetConfig(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x09,       // Type
		0x00, 0x0c, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x01, // Flags
		0x00, 0x01, // MissSendLen
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	req := NewOfpSetConfig(1, 1)
	actual := req.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpSetConfig is not equal to expected value.")
	}
}

func TestParseSwitchConfig(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x09,       // Type
		0x00, 0x0c, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x01, // Flags
		0x00, 0x01, // MissSendLen
	}

	rep := newOfpSwitchConfig(OFPT_SET_CONFIG, 1, 1)
	rep.Parse(packet)
	if rep.Header.Version != 4 || rep.Header.Type != 9 ||
		rep.Header.Length != 12 || rep.Header.Xid != 0 ||
		rep.Flags != 1 || rep.MissSendLen != 1 {
		t.Log("Version        : ", rep.Header.Version)
		t.Log("Type           : ", rep.Header.Type)
		t.Log("Length         : ", rep.Header.Length)
		t.Log("Transaction ID : ", rep.Header.Xid)
		t.Log("Flags          : ", rep.Flags)
		t.Log("MissSendLen    : ", rep.MissSendLen)
		t.Error("Parsed value of OfpSwitchConfig is invalid.")
	}
}

// TODO : test all other ofp13 messages

/*****************************************************/
/* OfpTableMod                                       */
/*****************************************************/
func TestSerializeTableMod(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x11,       // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x01,             // PortNo
		0x00, 0x00, 0x00, //Padding
		0x00, 0x00, 0x00, 0x03, // Config
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	req := NewOfpTableMod(1, OFPTC_DEPRECATED_MASK)
	actual := req.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableMod is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpPortStatus                                     */
/*****************************************************/
func TestParsePortStatus(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x0c,       // Type
		0x00, 0x50, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00,                                     // Reason
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x000, 0x01, // PortNo
		0x00, 0x00, 0x00, 0x00, // Padding
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // HwAddr
		0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Name
		0x00, 0x00, 0x00, 0x20, // Config
		0x00, 0x00, 0x00, 0x04, // State
		0x00, 0x00, 0x00, 0x01, // Curr
		0x00, 0x00, 0x00, 0x01, // Advertised
		0x00, 0x00, 0x00, 0x01, // Supported
		0x00, 0x00, 0x00, 0x01, // Peer
		0x00, 0x00, 0x01, 0x00, // CurrSpeed
		0x00, 0x00, 0x01, 0x00, // MaxSpeed
	}

	ps := NewOfpPortStatus()
	ps.Parse(packet)
	if ps.Header.Version != 4 || ps.Header.Type != OFPT_PORT_STATUS ||
		ps.Header.Length != 80 || ps.Header.Xid != 0 ||
		ps.Reason != OFPPR_ADD ||
		ps.Desc.PortNo != 1 ||
		ps.Desc.HwAddr.String() != "11:22:33:44:55:66" ||
		ps.Desc.Config != OFPPC_NO_FWD ||
		ps.Desc.State != OFPPS_LIVE ||
		ps.Desc.Curr != OFPPF_10MB_HD ||
		ps.Desc.Advertised != OFPPF_10MB_HD ||
		ps.Desc.Supported != OFPPF_10MB_HD ||
		ps.Desc.Peer != OFPPF_10MB_HD ||
		ps.Desc.CurrSpeed != 256 ||
		ps.Desc.MaxSpeed != 256 {
		t.Log("Version        : ", ps.Header.Version)
		t.Log("Type           : ", ps.Header.Type)
		t.Log("Length         : ", ps.Header.Length)
		t.Log("Transaction ID : ", ps.Header.Xid)
		t.Log("Reason         : ", ps.Reason)
		t.Log("PortNo         : ", ps.Desc.PortNo)
		t.Log("HwAddr         : ", ps.Desc.HwAddr.String())
		t.Log("Config         : ", ps.Desc.Config)
		t.Log("State          : ", ps.Desc.State)
		t.Log("Curr           : ", ps.Desc.Curr)
		t.Log("Advertised     : ", ps.Desc.Advertised)
		t.Log("Supported      : ", ps.Desc.Supported)
		t.Log("Peer           : ", ps.Desc.Peer)
		t.Log("CurrSpeed      : ", ps.Desc.CurrSpeed)
		t.Log("MaxSpeed       : ", ps.Desc.MaxSpeed)
		t.Error("Parsed value of OfpPortStatus is invalid.")
	}
}

/*****************************************************/
/* OfpPortMod                                        */
/*****************************************************/
func TestSerializePortMod(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x10,       // Type
		0x00, 0x28, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x01, // PortNo
		0x00, 0x00, 0x00, 0x00, //Padding
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // HwAddr
		0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, // Config
		0x00, 0x00, 0x00, 0x60, // Mask
		0x00, 0x00, 0x00, 0x03, // Advertise
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	req, _ := NewOfpPortMod(
		1,
		"11:22:33:44:55:66",
		0,
		(OFPPC_NO_FWD | OFPPC_NO_PACKET_IN),
		(OFPPF_10MB_HD | OFPPF_10MB_FD))
	actual := req.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpPortMod is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpFeaturesRequest                                */
/*****************************************************/
func TestSerializeFeaturesRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x05,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	req := NewOfpFeaturesRequest()
	actual := req.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpFeaturesRequest is not equal to expected value.")
	}
}

func TestParseFeaturesReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x06,       // Type
		0x00, 0x20, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Datapath ID
		0x00, 0x00, 0x01, 0x00, // nbuffers
		0xfe,       // ntables
		0x00,       // AuxiliaryId
		0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x47, // Capabilities
		0x00, 0x00, 0x00, 0x00, // Reserved
	}

	rep := NewOfpFeaturesReply()
	rep.Parse(packet)
	if rep.Header.Version != 4 || rep.Header.Type != 6 ||
		rep.Header.Length != 32 || rep.Header.Xid != 0 ||
		rep.NBuffers != 256 || rep.NTables != 254 ||
		rep.AuxiliaryId != 0 {
		t.Log("Version        : ", rep.Header.Version)
		t.Log("Type           : ", rep.Header.Type)
		t.Log("Length         : ", rep.Header.Length)
		t.Log("Transaction ID : ", rep.Header.Xid)
		t.Log("NBuffers       : ", rep.NBuffers)
		t.Log("NTables        : ", rep.NTables)
		t.Log("AuxiliaryId    : ", rep.AuxiliaryId)
		t.Error("Parsed value of OfpFeaturesReply is invalid.")
	}
}

/*****************************************************/
/* OfpFlowMod                                        */
/*****************************************************/
func TestSerializeFlowMod(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x0e,       // Type
		0x00, 0x58, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie Mask
		0x00,       // Table ID
		0x00,       // Command(OFPFC_ADD)
		0x00, 0x00, // Idle Timeout
		0x00, 0x00, // Hard Timeout
		0x00, 0x00, // Priority
		0xff, 0xff, 0xff, 0xff, // Buffer ID(OFP_NO_BUFFER)
		0xff, 0xff, 0xff, 0xff, // Out port(OFPP_ANY)
		0xff, 0xff, 0xff, 0xff, // Out group(OFPG_ANY)
		0x00, 0x01, // Flags
		0x00, 0x00, // Padding
		0x00, 0x01, // Match Type(OFPMT_OXM)
		0x00, 0x0e, // Length
		0x80, 0x00, // OFPXMC_OPENFLOW_BASIC
		0x06,                               // OFPXMT_OFB_ETH_DST, has mask is false.
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, //Value
		0x00, 0x00, // Padding
		0x00, 0x04, // OFPIT_APPLY_ACTIONS
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Length
		0x00, 0x00, // OFPAT_OUTPUT
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x00, // Port
		0x00, 0x00, // Max Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	fmod := NewOfpFlowMod(
		0,
		0,
		0,
		OFPFC_ADD,
		0,
		0,
		0,
		OFP_NO_BUFFER,
		OFPP_ANY,
		OFPG_ANY,
		OFPFF_SEND_FLOW_REM)
	ethdst, _ := NewOxmEthDst("11:22:33:44:55:66")
	if ethdst == nil {
		t.Error("Failed to create OxmEthDst.")
		return
	}
	fmod.AppendMatchField(ethdst)
	instruction := NewOfpInstructionActions(OFPIT_APPLY_ACTIONS)
	instruction.Append(NewOfpActionOutput(0, 0))
	fmod.AppendInstruction(instruction)

	actual := fmod.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpFlowMod is not equal to expected value.")
	}
}

// TODO: Test all match fields, instructions and actions.

// MatchField types
//  in_port			OFPXMT_OFB_IN_PORT
func TestSerializeOxmMatchInPort(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x00,                   // OFPXMT_OFB_IN_PORT, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmInPort(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmInPort is not equal to expected value.")
	}
}

func TestParseOxmMatchInPort(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x00,                   // OFPXMT_OFB_IN_PORT, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x01, // Value
	}

	oxm := NewOxmInPort(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IN_PORT ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmInPort is invalid.")
	}
}

//  in_phy_port		OFPXMT_OFB_IN_PHY_PORT
func TestSerializeOxmMatchInPhyPort(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x02,                   // OFPXMT_OFB_IN_PHY_PORT, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmInPhyPort(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmInPhyPort is not equal to expected value.")
	}
}

func TestParseOxmMatchInPhyPort(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x02,                   // OFPXMT_OFB_IN_PHY_PORT, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x01, // Value
	}

	oxm := NewOxmInPhyPort(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IN_PHY_PORT ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmInPhyPort is invalid.")
	}
}

//  metadata		OFPXMT_OFB_METADATA
func TestSerializeOxmMatchMetadata(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x04,                                           // OFPXMT_OFB_METADATA, Has mask is false
		0x08,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmMetadata(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmMetadata is not equal to expected value.")
	}
}

func TestParseOxmMatchMetadata(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x04,                                           // OFPXMT_OFB_METADATA, Has mask is false
		0x08,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // Value
	}

	oxm := NewOxmMetadata(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_METADATA ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 8 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmMetadata is invalid.")
	}
}

func TestSerializeOxmMatchMetadataW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x05,                                           // OFPXMT_OFB_METADATA, Has mask is true
		0x10,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmMetadataW(1, math.MaxUint64)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmMetadata is not equal to expected value.")
	}
}

func TestParseOxmMatchMetadataW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x05,                                           // OFPXMT_OFB_METADATA, Has mask is true
		0x10,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Value
	}

	oxm := NewOxmMetadataW(0, 0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_METADATA ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 16 ||
		oxm.Value != 1 ||
		oxm.Mask != math.MaxUint64 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmMetadata is invalid.")
	}
}

//  eth_dst			OFPXMT_OFB_ETH_DST
func TestSerializeOxmMatchEthDst(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x06,                               // OFPXMT_OFB_ETH_DST, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmEthDst("11:22:33:44:55:66")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmEthDst is not equal to expected value.")
	}
}

func TestParseOxmMatchEthDst(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x06,                               // OFPXMT_OFB_ETH_DST, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}

	oxm, _ := NewOxmEthDst("00:00:00:00:00:00")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ETH_DST ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 6 ||
		oxm.Value.String() != "11:22:33:44:55:66" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value.String())
		t.Error("Parsed value of OxmEthDst is invalid.")
	}
}

func TestSerializeOxmMatchEthDstW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x07,                               // OFPXMT_OFB_ETH_DST, Has mask is true
		0x0c,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmEthDstW("11:22:33:44:55:66", "ff:ff:ff:ff:ff:ff")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmEthDst is not equal to expected value.")
	}
}

func TestParseOxmMatchEthDstW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x07,                               // OFPXMT_OFB_ETH_DST, Has mask is true
		0x0c,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Value
	}

	oxm, _ := NewOxmEthDstW("11:22:33:44:55:66", "ff:ff:ff:ff:ff:ff")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ETH_DST ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 12 ||
		oxm.Value.String() != "11:22:33:44:55:66" ||
		oxm.Mask.String() != "ff:ff:ff:ff:ff:ff" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value.String())
		t.Log("Mask           : ", oxm.Mask.String())
		t.Error("Parsed value of OxmEthDst is invalid.")
	}
}

//  eth_src			OFPXMT_OFB_ETH_SRC
func TestSerializeOxmMatchEthSrc(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x08,                               // OFPXMT_OFB_ETH_SRC, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmEthSrc("11:22:33:44:55:66")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmEthSrc is not equal to expected value.")
	}
}

func TestParseOxmMatchEthSrc(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x08,                               // OFPXMT_OFB_ETH_SRC, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}

	oxm, _ := NewOxmEthSrc("00:00:00:00:00:00")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ETH_SRC ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 6 ||
		oxm.Value.String() != "11:22:33:44:55:66" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value.String())
		t.Error("Parsed value of OxmEthSrc is invalid.")
	}
}

func TestSerializeOxmMatchEthSrcW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x09,                               // OFPXMT_OFB_ETH_SRC, Has mask is true
		0x0c,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmEthSrcW("11:22:33:44:55:66", "ff:ff:ff:ff:ff:ff")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmEthSrc is not equal to expected value.")
	}
}

func TestParseOxmMatchEthSrcW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x09,                               // OFPXMT_OFB_ETH_SRC, Has mask is true
		0x0c,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Value
	}

	oxm, _ := NewOxmEthSrcW("11:22:33:44:55:66", "ff:ff:ff:ff:ff:ff")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ETH_SRC ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 12 ||
		oxm.Value.String() != "11:22:33:44:55:66" ||
		oxm.Mask.String() != "ff:ff:ff:ff:ff:ff" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value.String())
		t.Log("Mask           : ", oxm.Mask.String())
		t.Error("Parsed value of OxmEthSrc is invalid.")
	}
}

//  eth_type		OFPXMT_OFB_ETH_TYPE
func TestSerializeOxmMatchEthType(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0a,       // OFPXMT_OFB_ETH_TYPE, Has mask is false
		0x02,       // Length
		0x80, 0x00, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmEthType(0x8000)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmEthType is not equal to expected value.")
	}
}

func TestParseOxmMatchEthType(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0a,       // OFPXMT_OFB_ETH_TYPE, Has mask is false
		0x02,       // Length
		0x80, 0x00, // Value
	}

	oxm := NewOxmEthType(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ETH_TYPE ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 0x8000 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmEthTYpe is invalid.")
	}
}

//  vlan_vid		OFPXMT_OFB_VLAN_VID
func TestSerializeOxmMatchVlanVid(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0c,       // OFPXMT_OFB_VLAN_VID, Has mask is false
		0x02,       // Length
		0x00, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmVlanVid(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmVlanVid is not equal to expected value.")
	}
}

func TestParseOxmMatchEthVlanVid(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0c,       // OFPXMT_OFB_VLAN_VID, Has mask is false
		0x02,       // Length
		0x00, 0x01, // Value
	}

	oxm := NewOxmVlanVid(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_VLAN_VID ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmVlanVid is invalid.")
	}
}

func TestSerializeOxmMatchVlanVidW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0d,       // OFPXMT_OFB_VLAN_VID, Has mask is true
		0x04,       // Length
		0x00, 0x01, // Value
		0xff, 0xff, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmVlanVidW(1, 0xffff)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmVlanVid is not equal to expected value.")
	}
}

func TestParseOxmMatchEthVlanVidW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0d,       // OFPXMT_OFB_VLAN_VID, Has mask is false
		0x04,       // Length
		0x00, 0x01, // Value
		0xff, 0xff, // Mask
	}

	oxm := NewOxmVlanVidW(0, 0xffff)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_VLAN_VID ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 4 ||
		oxm.Value != 1 ||
		oxm.Mask != 0xffff {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmVlanVid is invalid.")
	}
}

//  vlan_pcp		OFPXMT_OFB_VLAN_PCP
func TestSerializeOxmMatchVlanPcp(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0e, // OFPXMT_OFB_VLAN_PCP, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmVlanPcp(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmVlanPcp is not equal to expected value.")
	}
}

func TestParseOxmMatchVlanPcp(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x0e, // OFPXMT_OFB_VLAN_PCP, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmVlanPcp(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_VLAN_PCP ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmVlanPcp is invalid.")
	}
}

//  ip_dscp			OFPXMT_OFB_IP_DSCP
func TestSerializeOxmMatchIpDscp(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x10, // OFPXMT_OFB_IP_DSCP, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIpDscp(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpDscp is not equal to expected value.")
	}
}

func TestParseOxmMatchIpDscp(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x10, // OFPXMT_OFB_IP_DSCP, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmIpDscp(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IP_DSCP ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpDscp is invalid.")
	}
}

//  ip_ecn			OFPXMT_OFB_IP_ECN
func TestSerializeOxmMatchIpEcn(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x12, // OFPXMT_OFB_IP_ECN, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIpEcn(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpEcn is not equal to expected value.")
	}
}

func TestParseOxmMatchIpEcn(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x12, // OFPXMT_OFB_IP_ECN, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmIpEcn(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IP_ECN ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpEcn is invalid.")
	}
}

//  ip_proto		OFPXMT_OFB_IP_PROTO
func TestSerializeOxmMatchIpProto(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x14, // OFPXMT_OFB_IP_PROTO, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIpProto(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpProto is not equal to expected value.")
	}
}

func TestParseOxmMatchIpProto(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x14, // OFPXMT_OFB_IP_PROTO, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmIpProto(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IP_PROTO ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpProto is invalid.")
	}
}

//  ipv4_src		OFPXMT_OFB_IPV4_SRC
func TestSerializeOxmMatchIpv4Src(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x16,                   // OFPXMT_OFB_IPV4_SRC, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv4Src("192.168.10.1")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv4Src is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv4Src(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x16,                   // OFPXMT_OFB_IPV4_SRC, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}

	oxm, _ := NewOxmIpv4Src("0.0.0.0")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV4_SRC ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value.String() != "192.168.10.1" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpv4Src is invalid.")
	}
}

func TestSerializeOxmMatchIpv4SrcW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x17,                   // OFPXMT_OFB_IPV4_SRC, Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // MASK
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv4SrcW("192.168.10.1", 24)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv4Src is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv4SrcW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x17,                   // OFPXMT_OFB_IPV4_SRC, Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // MASK
	}

	oxm, _ := NewOxmIpv4SrcW("0.0.0.0", 24)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV4_SRC ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 8 ||
		oxm.Value.String() != "192.168.10.1" ||
		oxm.Mask.String() != "ffffff00" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmIpv4Src is invalid.")
	}
}

//  ipv4_dst		OFPXMT_OFB_IPV4_DST
func TestSerializeOxmMatchIpv4Dst(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x18,                   // OFPXMT_OFB_IPV4_DST, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv4Dst("192.168.10.1")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv4Dst is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv4Dst(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x18,                   // OFPXMT_OFB_IPV4_DST, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}

	oxm, _ := NewOxmIpv4Dst("0.0.0.0")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV4_DST ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value.String() != "192.168.10.1" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpv4Dst is invalid.")
	}
}

func TestSerializeOxmMatchIpv4DstW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x19,                   // OFPXMT_OFB_IPV4_DST, Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // MASK
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv4DstW("192.168.10.1", 24)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv4Dst is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv4DstW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x19,                   // OFPXMT_OFB_IPV4_DST Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // MASK
	}

	oxm, _ := NewOxmIpv4DstW("0.0.0.0", 24)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV4_DST ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 8 ||
		oxm.Value.String() != "192.168.10.1" ||
		oxm.Mask.String() != "ffffff00" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmIpv4Dst is invalid.")
	}
}

//  tcp_src			OFPXMT_OFB_TCP_SRC
func TestSerializeOxmMatchTcpSrc(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x1a,       // OFPXMT_OFB_TCP_SRC, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmTcpSrc(8000)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmTcpSrc is not equal to expected value.")
	}
}

func TestParseOxmMatchTcpSrc(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x1a,       // OFPXMT_OFB_TCP_SRC, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}

	oxm := NewOxmTcpSrc(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_TCP_SRC ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 8000 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmTcpSrc is invalid.")
	}
}

//  tcp_dst			OFPXMT_OFB_TCP_DST
func TestSerializeOxmMatchTcpDst(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x1c,       // OFPXMT_OFB_TCP_DST, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmTcpDst(8000)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmTcpDst is not equal to expected value.")
	}
}

func TestParseOxmMatchTcpDst(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x1c,       // OFPXMT_OFB_TCP_DST, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}

	oxm := NewOxmTcpDst(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_TCP_DST ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 8000 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmTcpDst is invalid.")
	}
}

//  udp_src			OFPXMT_OFB_UDP_SRC
func TestSerializeOxmMatchUdpSrc(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x1e,       // OFPXMT_OFB_UDP_SRC, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmUdpSrc(8000)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmUdpSrc is not equal to expected value.")
	}
}

func TestParseOxmMatchUdpSrc(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x1e,       // OFPXMT_OFB_UDP_SRC, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}

	oxm := NewOxmUdpSrc(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_UDP_SRC ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 8000 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmUdpSrc is invalid.")
	}
}

//  udp_dst			OFPXMT_OFB_UDP_DST
func TestSerializeOxmMatchUdpDst(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x20,       // OFPXMT_OFB_UDP_DST, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmUdpDst(8000)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmUdpDst is not equal to expected value.")
	}
}

func TestParseOxmMatchUdpDst(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x20,       // OFPXMT_OFB_UDP_DST, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}

	oxm := NewOxmUdpDst(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_UDP_DST ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 8000 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmUdpDst is invalid.")
	}
}

//  sctp_src		OFPXMT_OFB_SCTP_SRC
func TestSerializeOxmMatchSctpSrc(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x22,       // OFPXMT_OFB_SCTP_SRC, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmSctpSrc(8000)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmSctpSrc is not equal to expected value.")
	}
}

func TestParseOxmMatchSctpSrc(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x22,       // OFPXMT_OFB_SCTP_SRC, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}

	oxm := NewOxmSctpSrc(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_SCTP_SRC ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 8000 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmSctpSrc is invalid.")
	}
}

//  sctp_dst		OFPXMT_OFB_SCTP_DST
func TestSerializeOxmMatchSctpDst(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x24,       // OFPXMT_OFB_SCTP_DST, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmSctpDst(8000)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmSctpDst is not equal to expected value.")
	}
}

func TestParseOxmMatchSctpDst(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x24,       // OFPXMT_OFB_SCTP_DST, Has mask is false
		0x02,       // Length
		0x1f, 0x40, // Value
	}

	oxm := NewOxmSctpDst(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_SCTP_DST ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 8000 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmSctpDst is invalid.")
	}
}

//  icmpv4_type		OFPXMT_OFB_ICMPV4_TYPE
func TestSerializeOxmMatchIcmpv4Type(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x26, // OFPXMT_OFB_ICMPV4_TYPE, Has mask is false
		0x01, // Length
		0x08, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIcmpType(8)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIcmpType is not equal to expected value.")
	}
}

func TestParseOxmMatchIcmpv4Type(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x26, // OFPXMT_OFB_ICMPV4_TYPE, Has mask is false
		0x01, // Length
		0x08, // Value
	}

	oxm := NewOxmIcmpType(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ICMPV4_TYPE ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 8 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIcmpType is invalid.")
	}
}

//  icmpv4_code		OFPXMT_OFB_ICMPV4_CODE
func TestSerializeOxmMatchIcmpv4Code(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x28, // OFPXMT_OFB_ICMPV4_CODE, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIcmpCode(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIcmpCode is not equal to expected value.")
	}
}

func TestParseOxmMatchIcmpv4Code(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x28, // OFPXMT_OFB_ICMPV4_CODE, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmIcmpType(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ICMPV4_CODE ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIcmpCode is invalid.")
	}
}

//  arp_op			OFPXMT_OFB_ARP_OP
func TestSerializeOxmMatchArpOp(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2a,       // OFPXMT_OFB_ARP_OP, Has mask is false
		0x02,       // Length
		0x00, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmArpOp(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmArpOp is not equal to expected value.")
	}
}

func TestParseOxmMatchArpOp(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2a,       // OFPXMT_OFB_ARP_OP, Has mask is false
		0x02,       // Length
		0x00, 0x01, // Value
	}

	oxm := NewOxmArpOp(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ARP_OP ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmArpOp is invalid.")
	}
}

//  arp_spa			OFPXMT_OFB_ARP_SPA
func TestSerializeOxmMatchArpSpa(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2c,                   // OFPXMT_OFB_ARP_SPA, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmArpSpa("192.168.10.1")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmArpSpa is not equal to expected value.")
	}
}

func TestParseOxmMatchArpSpa(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2c,                   // OFPXMT_OFB_ARP_SPA, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}

	oxm, _ := NewOxmArpSpa("192.168.10.1")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ARP_SPA ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value.String() != "192.168.10.1" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmArpSpa is invalid.")
	}
}

func TestSerializeOxmMatchArpSpaW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2d,                   // OFPXMT_OFB_ARP_SPA, Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmArpSpaW("192.168.10.1", 24)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmArpSpa is not equal to expected value.")
	}
}

func TestParseOxmMatchArpSpaW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2d,                   // OFPXMT_OFB_ARP_SPA, Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // Mask
	}

	oxm, _ := NewOxmArpSpaW("192.168.10.1", 24)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ARP_SPA ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 8 ||
		oxm.Value.String() != "192.168.10.1" ||
		oxm.Mask.String() != "ffffff00" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmArpSpa is invalid.")
	}
}

//  arp_tpa			OFPXMT_OFB_ARP_TPA
func TestSerializeOxmMatchArpTpa(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2e,                   // OFPXMT_OFB_ARP_TPA, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmArpTpa("192.168.10.1")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmArpTpa is not equal to expected value.")
	}
}

func TestParseOxmMatchArpTpa(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2e,                   // OFPXMT_OFB_ARP_TPA, Has mask is false
		0x04,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
	}

	oxm, _ := NewOxmArpTpa("192.168.10.1")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ARP_TPA ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value.String() != "192.168.10.1" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmArpTpa is invalid.")
	}
}

func TestSerializeOxmMatchArpTpaW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2f,                   // OFPXMT_OFB_ARP_TPA, Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmArpTpaW("192.168.10.1", 24)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmArpTpa is not equal to expected value.")
	}
}

func TestParseOxmMatchArpTpaW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x2f,                   // OFPXMT_OFB_ARP_TPA, Has mask is true
		0x08,                   // Length
		0xc0, 0xa8, 0x0a, 0x01, // Value
		0xff, 0xff, 0xff, 0x00, // Mask
	}

	oxm, _ := NewOxmArpTpaW("192.168.10.1", 24)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ARP_TPA ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 8 ||
		oxm.Value.String() != "192.168.10.1" ||
		oxm.Mask.String() != "ffffff00" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmArpTpa is invalid.")
	}
}

//  arp_sha			OFPXMT_OFB_ARP_SHA
func TestSerializeOxmMatchArpSha(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x30,                               // OFPXMT_OFB_ARP_SHA, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmArpSha("11:22:33:44:55:66")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmArpSha is not equal to expected value.")
	}
}

func TestParseOxmMatchArpSha(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x30,                               // OFPXMT_OFB_ARP_SHA, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}

	oxm, _ := NewOxmArpSha("11:22:33:44:55:66")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ARP_SHA ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 6 ||
		oxm.Value.String() != "11:22:33:44:55:66" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmArpSha is invalid.")
	}
}

//  arp_tha			OFPXMT_OFB_ARP_THA
func TestSerializeOxmMatchArpTha(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x32,                               // OFPXMT_OFB_ARP_THA, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmArpTha("11:22:33:44:55:66")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmArpTha is not equal to expected value.")
	}
}

func TestParseOxmMatchArpTha(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x32,                               // OFPXMT_OFB_ARP_THA, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}

	oxm, _ := NewOxmArpTha("11:22:33:44:55:66")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ARP_THA ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 6 ||
		oxm.Value.String() != "11:22:33:44:55:66" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmArpTha is invalid.")
	}
}

//  ipv6_src		OFPXMT_OFB_IPV6_SRC
func TestSerializeOxmMatchIpv6Src(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x34,                   // OFPXMT_OFB_IPV6_SRC, Has mask is false
		0x10,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv6Src("fe80::e4d:eeff:fe00:90cc")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6Src is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6Src(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x34,                   // OFPXMT_OFB_IPV6_SRC, Has mask is false
		0x10,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
	}

	oxm, _ := NewOxmIpv6Src("::1")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_SRC ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 16 ||
		oxm.Value.String() != "fe80::e4d:eeff:fe00:90cc" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpv6Src is invalid.")
	}
}

func TestSerializeOxmMatchIpv6SrcW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x35,                   // OFPXMT_OFB_IPV6_SRC, Has mask is true
		0x20,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
		0xff, 0xc0, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv6SrcW("fe80::e4d:eeff:fe00:90cc", 10)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6Src is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6SrcW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x35,                   // OFPXMT_OFB_IPV6_SRC, Has mask is true
		0x20,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
		0xff, 0xc0, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Mask
	}

	oxm, _ := NewOxmIpv6SrcW("::1", 10)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_SRC ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 32 ||
		oxm.Value.String() != "fe80::e4d:eeff:fe00:90cc" ||
		oxm.Mask.String() != "ffc00000000000000000000000000000" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmIpv6Src is invalid.")
	}
}

//  ipv6_dst		OFPXMT_OFB_IPV6_DST
func TestSerializeOxmMatchIpv6Dst(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x36,                   // OFPXMT_OFB_IPV6_DST, Has mask is false
		0x10,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv6Dst("fe80::e4d:eeff:fe00:90cc")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6Dst is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6Dst(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x36,                   // OFPXMT_OFB_IPV6_DST, Has mask is false
		0x10,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
	}

	oxm, _ := NewOxmIpv6Dst("::1")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_DST ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 16 ||
		oxm.Value.String() != "fe80::e4d:eeff:fe00:90cc" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpv6Dst is invalid.")
	}
}

func TestSerializeOxmMatchIpv6DstW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x37,                   // OFPXMT_OFB_IPV6_DST, Has mask is true
		0x20,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
		0xff, 0xc0, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv6DstW("fe80::e4d:eeff:fe00:90cc", 10)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6Dst is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6DstW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x37,                   // OFPXMT_OFB_IPV6_DST, Has mask is true
		0x20,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
		0xff, 0xc0, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Mask
	}

	oxm, _ := NewOxmIpv6DstW("::1", 10)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_DST ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 32 ||
		oxm.Value.String() != "fe80::e4d:eeff:fe00:90cc" ||
		oxm.Mask.String() != "ffc00000000000000000000000000000" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmIpv6Dst is invalid.")
	}
}

//  ipv6_flabel		OFPXMT_OFB_IPV6_FLABEL
func TestSerializeOxmMatchIpv6FLabel(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x38,                   // OFPXMT_OFB_IPV6_FLABEL, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x0a, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIpv6FLabel(10)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6FLabel is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6FLabel(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x38,                   // OFPXMT_OFB_IPV6_FLABEL, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x0a, // Value
	}

	oxm := NewOxmIpv6FLabel(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_FLABEL ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value != 10 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpv6FLabel is invalid.")
	}
}

func TestSerializeOxmMatchIpv6FLabelW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x39,                   // OFPXMT_OFB_IPV6_FLABEL, Has mask is true
		0x08,                   // Length
		0x00, 0x00, 0x00, 0x0a, // Value
		0x00, 0x00, 0xff, 0xff, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIpv6FLabelW(10, 0xffff)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6FLabel is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6FLabelW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x39,                   // OFPXMT_OFB_IPV6_FLABEL, Has mask is true
		0x08,                   // Length
		0x00, 0x00, 0x00, 0x0a, // Value
		0x00, 0x00, 0xff, 0xff, // Mask
	}

	oxm := NewOxmIpv6FLabelW(0, 0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_FLABEL ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 8 ||
		oxm.Value != 10 ||
		oxm.Mask != 0xffff {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmIpv6FLabel is invalid.")
	}
}

//  icmpv6_type		OFPXMT_OFB_ICMPV6_TYPE
func TestSerializeOxmMatchIpv6Type(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x3a, // OFPXMT_OFB_ICMPV6_TYPE, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIcmpv6Type(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIcmpv6Type is not equal to expected value.")
	}
}

func TestParseOxmMatchIcmpv6Type(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x3a, // OFPXMT_OFB_ICMPV6_TYPE, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmIcmpv6Type(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ICMPV6_TYPE ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIcmpv6Type is invalid.")
	}
}

//  icmpv6_code		OFPXMT_OFB_ICMPV6_CODE
func TestSerializeOxmMatchIpv6Code(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x3c, // OFPXMT_OFB_ICMPV6_CODE, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIcmpv6Code(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIcmpv6Code is not equal to expected value.")
	}
}

func TestParseOxmMatchIcmpv6Code(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x3c, // OFPXMT_OFB_ICMPV6_CODE, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmIcmpv6Code(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_ICMPV6_CODE ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIcmpv6Code is invalid.")
	}
}

//  ipv6_nd_target	OFPXMT_OFB_IPV6_ND_TARGET
func TestSerializeOxmMatchIpv6NdTarget(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x3e,                   // OFPXMT_OFB_IPV6_ND_TARGET, Has mask is false
		0x10,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv6NdTarget("fe80::0e4d:eeff:fe00:90cc")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6NdTarget is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6NdTarget(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x3e,                   // OFPXMT_OFB_IPV6_ND_TARGET, Has mask is false
		0x10,                   // Length
		0xfe, 0x80, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x0e, 0x4d, 0xee, 0xff, //
		0xfe, 0x00, 0x90, 0xcc, // Value
	}

	oxm, _ := NewOxmIpv6NdTarget("::1")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_ND_TARGET ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 16 ||
		oxm.Value.String() != "fe80::e4d:eeff:fe00:90cc" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpv6NdTarget is invalid.")
	}
}

//  ipv6_nd_sll		OFPXMT_OFB_IPV6_ND_SLL
func TestSerializeOxmMatchIpv6NdSll(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x40,                               // OFPXMT_OFB_IPV6_ND_SLL, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv6NdSll("11:22:33:44:55:66")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6NdSll is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6NdSll(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x40,                               // OFPXMT_OFB_IPV6_ND_SLL, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}

	oxm, _ := NewOxmIpv6NdSll("00:00:00:00:00:00")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_ND_SLL ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 6 ||
		oxm.Value.String() != "11:22:33:44:55:66" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpc6NdSll is invalid.")
	}
}

//  ipv6_nd_tll		OFPXMT_OFB_IPV6_ND_TLL
func TestSerializeOxmMatchIpv6NdTll(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x42,                               // OFPXMT_OFB_IPV6_ND_TLL, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm, _ := NewOxmIpv6NdTll("11:22:33:44:55:66")
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6NdTll is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6NdTll(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x42,                               // OFPXMT_OFB_IPV6_ND_TLL, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
	}

	oxm, _ := NewOxmIpv6NdTll("00:00:00:00:00:00")
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_ND_TLL ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 6 ||
		oxm.Value.String() != "11:22:33:44:55:66" {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpc6NdTll is invalid.")
	}
}

//  mpls_label		OFPXMT_OFB_MPLS_LABEL
func TestSerializeOxmMatchMplsLabel(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x44,                   // OFPXMT_OFB_MPLS_LABEL, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmMplsLabel(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmMplsLabel is not equal to expected value.")
	}
}

func TestParseOxmMatchMplsLabel(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x44,                   // OFPXMT_OFB_MPLS_LABEL, Has mask is false
		0x04,                   // Length
		0x00, 0x00, 0x00, 0x01, // Value
	}

	oxm := NewOxmMplsLabel(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_MPLS_LABEL ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 4 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmMplsLabel is invalid.")
	}
}

//  mpls_tc			OFPXMT_OFB_MPLS_TC
func TestSerializeOxmMatchMplsTc(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x46, // OFPXMT_OFB_MPLS_TC, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmMplsTc(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmMplsTc is not equal to expected value.")
	}
}

func TestParseOxmMatchMplsTc(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x46, // OFPXMT_OFB_MPLS_TC, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmMplsTc(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_MPLS_TC ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmMplsTc is invalid.")
	}
}

//  mpls_bos		OFPXMT_OFB_MPLS_BOS
func TestSerializeOxmMatchMplsBos(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x48, // OFPXMT_OFB_MPLS_BOS, Has mask is false
		0x01, // Length
		0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmMplsBos(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmMplsBos is not equal to expected value.")
	}
}

func TestParseOxmMatchMplsBos(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x48, // OFPXMT_OFB_MPLS_BOS, Has mask is false
		0x01, // Length
		0x01, // Value
	}

	oxm := NewOxmMplsBos(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_MPLS_BOS ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 1 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmMplsBos is invalid.")
	}
}

//  pbb_isid		OFPXMT_OFB_PBB_ISID
func TestSerializeOxmMatchPbbIsid(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4a,             // OFPXMT_OFB_PBB_ISID, Has mask is false
		0x03,             // Length
		0x00, 0x00, 0x03, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmPbbIsid([3]uint8{0, 0, 3})
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmPbbIsid is not equal to expected value.")
	}
}

func TestParseOxmMatchPbbIsid(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4a,             // OFPXMT_OFB_PBB_ISID, Has mask is false
		0x03,             // Length
		0x00, 0x00, 0x03, // Value
	}

	oxm := NewOxmPbbIsid([3]uint8{0, 0, 0})
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_PBB_ISID ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 3 ||
		oxm.Value[0] != 0 ||
		oxm.Value[1] != 0 ||
		oxm.Value[2] != 3 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmPbbIsid is invalid.")
	}
}

func TestSerializeOxmMatchPbbIsidW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4b,             // OFPXMT_OFB_PBB_ISID, Has mask is true
		0x06,             // Length
		0x00, 0x00, 0x03, // Value
		0xff, 0xff, 0xff, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmPbbIsidW([3]uint8{0, 0, 3}, [3]uint8{255, 255, 255})
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmPbbIsid is not equal to expected value.")
	}
}

func TestParseOxmMatchPbbIsidW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4b,             // OFPXMT_OFB_PBB_ISID, Has mask is true
		0x06,             // Length
		0x00, 0x00, 0x03, // Value
		0xff, 0xff, 0xff, // Mask
	}

	oxm := NewOxmPbbIsidW([3]uint8{0, 0, 0}, [3]uint8{0, 0, 0})
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_PBB_ISID ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 6 ||
		oxm.Value[0] != 0 ||
		oxm.Value[1] != 0 ||
		oxm.Value[2] != 3 ||
		oxm.Mask[0] != 0xff ||
		oxm.Mask[1] != 0xff ||
		oxm.Mask[2] != 0xff {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmPbbIsid is invalid.")
	}
}

//  tunnel_id		OFPXMT_OFB_TUNNEL_ID
func TestSerializeOxmMatchTunnelId(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4c,                                           // OFPXMT_OFB_TUNNEL_ID, Has mask is false
		0x08,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmTunnelId(10)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmTunnelId is not equal to expected value.")
	}
}

func TestParseOxmMatchTunnelId(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4c,                                           // OFPXMT_OFB_TUNNEL_ID, Has mask is false
		0x08,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, // Value
	}

	oxm := NewOxmTunnelId(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_TUNNEL_ID ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 8 ||
		oxm.Value != 10 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmTunnelId is invalid.")
	}
}

func TestSerializeOxmMatchTunnelIdW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4d,                                           // OFPXMT_OFB_TUNNEL_ID, Has mask is false
		0x10,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmTunnelIdW(10, 0xffffffffffffffff)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmTunnelId is not equal to expected value.")
	}
}

func TestParseOxmMatchTunnelIdW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4d,                                           // OFPXMT_OFB_TUNNEL_ID, Has mask is false
		0x10,                                           // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0a, // Value
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // Value
	}

	oxm := NewOxmTunnelIdW(0, 0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_TUNNEL_ID ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 16 ||
		oxm.Value != 10 ||
		oxm.Mask != 0xffffffffffffffff {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmTunnelId is invalid.")
	}
}

//  ipv6_exthdr		OFPXMT_OFB_IPV6_EXTHDR
func TestSerializeOxmMatchIpv6ExtHdr(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4e,       // OFPXMT_OFB_IPV6_EXTHDR, Has mask is false
		0x02,       // Length
		0x00, 0x01, // Value
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIpv6ExtHeader(1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6ExtHeader is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6ExtHdr(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4e,       // OFPXMT_OFB_IPV6_EXTHDR, Has mask is false
		0x02,       // Length
		0x00, 0x01, // Value
	}

	oxm := NewOxmIpv6ExtHeader(0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_EXTHDR ||
		oxm.OxmHasMask() != 0 ||
		oxm.Length() != 2 ||
		oxm.Value != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Error("Parsed value of OxmIpv6ExtHeader is invalid.")
	}
}

func TestSerializeOxmMatchIpv6ExtHdrW(t *testing.T) {
	expect := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4f,       // OFPXMT_OFB_IPV6_EXTHDR, Has mask is true
		0x04,       // Length
		0x00, 0x01, // Value
		0x00, 0x01, // Mask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	oxm := NewOxmIpv6ExtHeaderW(1, 1)
	actual := oxm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OxmIpv6ExtHeader is not equal to expected value.")
	}
}

func TestParseOxmMatchIpv6ExtHdrW(t *testing.T) {
	packet := []byte{
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x4f,       // OFPXMT_OFB_IPV6_EXTHDR, Has mask is true
		0x04,       // Length
		0x00, 0x01, // Value
		0x00, 0x01, // Mask
	}

	oxm := NewOxmIpv6ExtHeaderW(0, 0)
	oxm.Parse(packet)
	if oxm.OxmClass() != OFPXMC_OPENFLOW_BASIC ||
		oxm.OxmField() != OFPXMT_OFB_IPV6_EXTHDR ||
		oxm.OxmHasMask() != 1 ||
		oxm.Length() != 4 ||
		oxm.Value != 1 ||
		oxm.Mask != 1 {
		t.Log("Class          : ", oxm.OxmClass())
		t.Log("Field          : ", oxm.OxmField())
		t.Log("HasMask        : ", oxm.OxmHasMask())
		t.Log("Length         : ", oxm.Length())
		t.Log("Value          : ", oxm.Value)
		t.Log("Mask           : ", oxm.Mask)
		t.Error("Parsed value of OxmIpv6ExtHeader is invalid.")
	}
}

// Instructions
// OFPIT_GOTO_TABLE
func TestSerializeInstructionGotoTable(t *testing.T) {
	expect := []byte{
		0x00, 0x01, // Type
		0x00, 0x08, // Length
		0x01,             // TableId
		0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	inst := NewOfpInstructionGotoTable(1)
	actual := inst.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpInstructionGotoTable is not equal to expected value.")
	}
}

func TestParseInstructionGotoTable(t *testing.T) {
	packet := []byte{
		0x00, 0x01, // Type
		0x00, 0x08, // Length
		0x01,             // TableId
		0x00, 0x00, 0x00, // Padding
	}

	inst := NewOfpInstructionGotoTable(0)
	inst.Parse(packet)
	if inst.InstructionType() != OFPIT_GOTO_TABLE ||
		inst.Header.Length != 8 ||
		inst.TableId != 1 {
		t.Log("Type           : ", inst.InstructionType())
		t.Log("Length         : ", inst.Header.Length)
		t.Log("TableId        : ", inst.TableId)
		t.Error("Parsed value of OfpInstructionGotoTable is invalid.")
	}
}

// OFPIT_WRITE_METADATA
func TestSerializeInstructionWriteMetadata(t *testing.T) {
	expect := []byte{
		0x00, 0x02, // Type
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // Metadata
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // MetadataMask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	inst := NewOfpInstructionWriteMetadata(1, 0xffffffffffffffff)
	actual := inst.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpInstructionWriteMetadata is not equal to expected value.")
	}
}

func TestParseInstructionWriteMetadata(t *testing.T) {
	packet := []byte{
		0x00, 0x02, // Type
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // Metadata
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, // MetadataMask
	}

	inst := NewOfpInstructionWriteMetadata(0, 0)
	inst.Parse(packet)
	if inst.InstructionType() != OFPIT_WRITE_METADATA ||
		inst.Header.Length != 24 ||
		inst.Metadata != 1 ||
		inst.MetadataMask != 0xffffffffffffffff {
		t.Log("Type           : ", inst.InstructionType())
		t.Log("Length         : ", inst.Header.Length)
		t.Log("Metadata       : ", inst.Metadata)
		t.Log("MetadataMask   : ", inst.MetadataMask)
		t.Error("Parsed value of OfpInstructionWriteMetadata is invalid.")
	}
}

// OFPIT_WRITE_ACTIONS
// OFPIT_APPLY_ACTIONS
// OFPIT_CLEAR_ACTIONS
func TestSerializeInstructionActions(t *testing.T) {
	expect := []byte{
		0x00, 0x04, // Type
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x01, // Port
		0xff, 0xe5, // MaxLen
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionOutput(1, 0xffe5)
	inst := NewOfpInstructionActions(OFPIT_APPLY_ACTIONS)
	inst.Append(action)
	actual := inst.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpInstructionActions is not equal to expected value.")
	}
}

func TestParseInstructionActions(t *testing.T) {
	packet := []byte{
		0x00, 0x04, // Type
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x01, // Port
		0xff, 0xe5, // MaxLen
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
	}

	inst := NewOfpInstructionActions(OFPIT_WRITE_ACTIONS)
	inst.Parse(packet)
	if inst.InstructionType() != OFPIT_APPLY_ACTIONS ||
		inst.Header.Length != 24 {
		t.Log("Type           : ", inst.InstructionType())
		t.Log("Length         : ", inst.Header.Length)
		t.Error("Parsed value of OfpInstructionActions is invalid.")
	}

	if action, ok := inst.Actions[0].(*OfpActionOutput); ok == true {
		if action.ActionHeader.Type != OFPAT_OUTPUT ||
			action.ActionHeader.Length != 16 ||
			action.Port != 1 ||
			action.MaxLen != 0xffe5 {
			t.Error("Parsed value of OfpInstructionActions is invalid.")
		}
	} else {
		t.Log("Type           : ", action.ActionHeader.Type)
		t.Log("Length         : ", action.ActionHeader.Length)
		t.Error("Parsed value of OfpInstructionActions is invalid.")
	}
}

// OFPIT_METER
func TestSerializeInstructionMeter(t *testing.T) {
	expect := []byte{
		0x00, 0x06, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x01, // MeterId
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	inst := NewOfpInstructionMeter(1)
	actual := inst.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpInstructionMeter is not equal to expected value.")
	}
}

func TestParseInstructionMeter(t *testing.T) {
	packet := []byte{
		0x00, 0x06, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x01, // MeterId
	}

	inst := NewOfpInstructionMeter(0)
	inst.Parse(packet)
	if inst.InstructionType() != OFPIT_METER ||
		inst.Header.Length != 8 ||
		inst.MeterId != 1 {
		t.Log("Type           : ", inst.InstructionType())
		t.Log("Length         : ", inst.Header.Length)
		t.Log("MeterId        : ", inst.MeterId)
		t.Error("Parsed value of OfpInstructionMeter is invalid.")
	}
}

// OFPIT_EXPERIMENTER
func TestSerializeInstructionExperimenter(t *testing.T) {
	expect := []byte{
		0xff, 0xff, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // MeterId
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	inst := NewOfpInstructionExperimenter(0)
	actual := inst.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpInstructionExperimenter is not equal to expected value.")
	}
}

func TestParseInstructionExperimenter(t *testing.T) {
	packet := []byte{
		0xff, 0xff, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Experimenter
	}

	inst := NewOfpInstructionExperimenter(0)
	inst.Parse(packet)
	if inst.InstructionType() != OFPIT_EXPERIMENTER ||
		inst.Header.Length != 8 ||
		inst.Experimenter != 0 {
		t.Log("Type           : ", inst.InstructionType())
		t.Log("Length         : ", inst.Header.Length)
		t.Log("Experimenter   : ", inst.Experimenter)
		t.Error("Parsed value of OfpInstructionExperimenter is invalid.")
	}
}

// Actions
// OFPAT_OUTPUT
func TestSerializeActionOutput(t *testing.T) {
	expect := []byte{
		0x00, 0x00, // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x01, // Port
		0xff, 0xe5, // MaxLen
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionOutput(1, OFPCML_MAX)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionOutput is not equal to expected value.")
	}
}

// OFPAT_COPY_TTL_OUT
func TestSerializeActionTtlOut(t *testing.T) {
	expect := []byte{
		0x00, 0x0b, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionCopyTtlOut()
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionTtlOut is not equal to expected value.")
	}
}

// OFPAT_COPY_TTL_IN
func TestSerializeActionTtlIn(t *testing.T) {
	expect := []byte{
		0x00, 0x0c, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionCopyTtlIn()
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionTtlIn is not equal to expected value.")
	}
}

// OFPAT_SET_MPLS_TTL
func TestSerializeActionSetMplsTtl(t *testing.T) {
	expect := []byte{
		0x00, 0x0f, // Type
		0x00, 0x08, // Length
		0x40,             // MplsTtl
		0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionSetMplsTtl(64)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionSetMplsTtl is not equal to expected value.")
	}
}

// OFPAT_DEC_MPLS_TTL
func TestSerializeActionDecMplsTtl(t *testing.T) {
	expect := []byte{
		0x00, 0x10, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionDecMplsTtl()
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionDecMplsTtl is not equal to expected value.")
	}
}

// OFPAT_PUSH_VLAN
func TestSerializeActionPushVlan(t *testing.T) {
	expect := []byte{
		0x00, 0x11, // Type
		0x00, 0x08, // Length
		0x81, 0x00, // EtherType
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionPushVlan()
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionPushVlan is not equal to expected value.")
	}
}

// OFPAT_POP_VLAN
func TestSerializeActionPopVlan(t *testing.T) {
	expect := []byte{
		0x00, 0x12, // Type
		0x00, 0x08, // Length
		0x08, 0x00, // EtherType
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionPopVlan(0x0800)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionPopVlan is not equal to expected value.")
	}
}

// OFPAT_PUSH_MPLS
func TestSerializeActionPushMpls(t *testing.T) {
	expect := []byte{
		0x00, 0x13, // Type
		0x00, 0x08, // Length
		0x88, 0x47, // EtherType
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionPushMpls()
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionPushMpls is not equal to expected value.")
	}
}

// OFPAT_POP_MPLS
func TestSerializeActionPopMpls(t *testing.T) {
	expect := []byte{
		0x00, 0x14, // Type
		0x00, 0x08, // Length
		0x08, 0x00, // EtherType
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionPopMpls(0x0800)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionPopMpls is not equal to expected value.")
	}
}

// OFPAT_SET_QUEUE
func TestSerializeActionSetQueue(t *testing.T) {
	expect := []byte{
		0x00, 0x15, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x01, // QueueId
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionSetQueue(1)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionSetQueue is not equal to expected value.")
	}
}

// OFPAT_GROUP
func TestSerializeActionGroup(t *testing.T) {
	expect := []byte{
		0x00, 0x16, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x01, // GroupId
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionGroup(1)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionGroup is not equal to expected value.")
	}
}

// OFPAT_SET_NW_TTL
func TestSerializeActionSetNwTtl(t *testing.T) {
	expect := []byte{
		0x00, 0x17, // Type
		0x00, 0x08, // Length
		0x40,             // TTL
		0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionSetNwTtl(64)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionSetNwTtl is not equal to expected value.")
	}
}

// OFPAT_DEC_NW_TTL
func TestSerializeActionDecNwTtl(t *testing.T) {
	expect := []byte{
		0x00, 0x18, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionDecNwTtl()
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionDecNwTtl is not equal to expected value.")
	}
}

// OFPAT_SET_FIELD
func TestSerializeActionSetField(t *testing.T) {
	expect := []byte{
		0x00, 0x19, // Type
		0x00, 0x10, // Length
		0x80, 0x00, // OFPXMC_OPENFLOW_BASIC
		0x06,                               // OFPXMT_OFB_ETH_DST, HasMask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0x00, 0x00, // Padding

	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	seteth, _ := NewOxmEthDst("11:22:33:44:55:66")
	action := NewOfpActionSetField(seteth)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionSetField is not equal to expected value.")
	}
}

// OFPAT_PUSH_PBB
func TestSerializeActionPushPbb(t *testing.T) {
	expect := []byte{
		0x00, 0x1a, // Type
		0x00, 0x08, // Length
		0x88, 0xe7, // EtherType
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionPushPbb()
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionPushPbb is not equal to expected value.")
	}
}

// OFPAT_POP_PBB
func TestSerializeActionPopPbb(t *testing.T) {
	expect := []byte{
		0x00, 0x1b, // Type
		0x00, 0x08, // Length
		0x08, 0x00, // EtherType
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionPopPbb(0x0800)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionPopPbb is not equal to expected value.")
	}
}

// OFPAT_EXPERIMENTER
func TestSerializeActionExperimenter(t *testing.T) {
	expect := []byte{
		0xff, 0xff, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // ExperimenterId
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	action := NewOfpActionExperimenter(0)
	actual := action.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpActionExperimenter is not equal to expected value.")
	}
}

// TODO : test all other ofp13 messages
/*****************************************************/
/* OfpGroupMod                                       */
/*****************************************************/
func TestSerializeGroupMod(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x0f,       // Type
		0x00, 0x30, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, // Command
		0x01,                   // Type
		0x00,                   // Padding
		0x00, 0x00, 0x00, 0x01, // GroupId
		0x00, 0x20, // Length
		0x00, 0x00, // Weight
		0x00, 0x00, 0x00, 0x01, // Watch Port
		0xff, 0xff, 0xff, 0xff, // Watch Group
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, // Output
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x01, // Port
		0xff, 0xe5, // Max Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	gm := NewOfpGroupMod(OFPGC_ADD, OFPGT_SELECT, 1)
	bucket := NewOfpBucket(0, 1, OFPG_ANY)
	bucket.Append(NewOfpActionOutput(1, OFPCML_MAX))
	gm.Append(bucket)
	actual := gm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpGroupMod is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpPacketOut                                      */
/*****************************************************/
func TestSerializePacketOut(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x0d,       // Type
		0x00, 0x28, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0xff, 0xff, 0xff, 0xff, // BufferId
		0x00, 0x00, 0x00, 0x01, // InPort
		0x00, 0x01, // ActionLen
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, // OFPAT_OUTPUT
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x01, // Port
		0xff, 0xe5, // Max Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	actions := make([]OfpAction, 0)
	action := NewOfpActionOutput(1, OFPCML_MAX)
	actions = append(actions, action)
	m := NewOfpPacketOut(0xffffffff, 1, actions, nil)
	actual := m.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpPacketOut is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpMeterMod                                       */
/*****************************************************/
func TestSerializeMeterModRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x1d,       // Type
		0x00, 0x20, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, // Command(OFPMC_ADD)
		0x00, 0x01, // Flags
		0x00, 0x00, 0x00, 0x01, // Meter ID
		0x00, 0x01, // OFPMBT_DROP
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x64, // Rate
		0x00, 0x00, 0x00, 0x00, // Burst Size
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	mm := NewOfpMeterMod(OFPMC_ADD, OFPMF_KBPS, 1)
	mdrop := NewOfpMeterBandDrop(100, 0)
	mm.AppendMeterBand(mdrop)
	actual := mm.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpMeterMod is not equal to expected value.")
	}
}

// TODO: test all other MeterBand Type

// TODO : test all other ofp13 messages
/*****************************************************/
/* OfpPacketIn                                       */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpFlowRemoved                                    */
/*****************************************************/
func TestParseFlowRemoved(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x0b,       // Type
		0x00, 0x40, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie
		0x00, 0x01, // Priority
		0x02,                   // Reason
		0x01,                   // TableId
		0x00, 0x00, 0x00, 0x00, // DurationSec
		0x00, 0x00, 0x00, 0x00, // DurationNSec
		0x00, 0x00, // IdleTimeout
		0x00, 0x00, // HardTimeout
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // PacketCount
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // ByteCount
		0x00, 0x01, // Type
		0x00, 0x0e, // Length
		0x80, 0x00, // OFPXMC_OPENFLOW_BASIC
		0x06,                               // OFPXMT_OFB_ETH_DST, HasMask is false
		0x06,                               // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Value
		0x00, 0x00, // Padding
	}

	m := NewOfpFlowRemoved()
	m.Parse(packet)
	if m.Header.Version != 4 || m.Header.Type != OFPT_FLOW_REMOVED ||
		m.Header.Length != 64 || m.Header.Xid != 0 ||
		m.Cookie != 0 || m.Priority != 1 ||
		m.Reason != OFPRR_DELETE || m.TableId != 1 ||
		m.DurationSec != 0 || m.DurationNSec != 0 ||
		m.IdleTimeout != 0 || m.HardTimeout != 0 ||
		m.PacketCount != 0 || m.ByteCount != 0 ||
		m.Match.Type != OFPMT_OXM || m.Match.Length != 14 {
		t.Log("Version        : ", m.Header.Version)
		t.Log("Type           : ", m.Header.Type)
		t.Log("Length         : ", m.Header.Length)
		t.Log("Transaction ID : ", m.Header.Xid)
		t.Log("Cookie         : ", m.Cookie)
		t.Log("Priority       : ", m.Priority)
		t.Log("Reason         : ", m.Reason)
		t.Log("TableId        : ", m.TableId)
		t.Log("DurationSec    : ", m.DurationSec)
		t.Log("DurationNSec   : ", m.DurationNSec)
		t.Log("IdleTimeout    : ", m.IdleTimeout)
		t.Log("HardTimeout    : ", m.HardTimeout)
		t.Log("PacketCount    : ", m.PacketCount)
		t.Log("ByteCount      : ", m.ByteCount)
		t.Log("MatchType      : ", m.Match.Type)
		t.Log("MatchLength    : ", m.Match.Length)
		t.Error("Parsed value of OfpFlowRemoved is invalid.")
	}
}

/*****************************************************/
/* OfpErrorMsg                                       */
/*****************************************************/
func TestParseErrorMsg(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x01,       // Type
		0x00, 0x20, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x01, // Type
		0x00, 0x01, // Code
		0x04,       //
		0x0f,       //
		0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Data
	}

	err := NewOfpErrorMsg()
	err.Parse(packet)
	if err.Header.Version != 4 || err.Header.Type != OFPT_ERROR ||
		err.Header.Length != 32 || err.Header.Xid != 0 ||
		err.Type != OFPET_BAD_REQUEST || err.Code != OFPBRC_BAD_TYPE {
		t.Log("Version        : ", err.Header.Version)
		t.Log("Type           : ", err.Header.Type)
		t.Log("Length         : ", err.Header.Length)
		t.Log("Transaction ID : ", err.Header.Xid)
		t.Error("Parsed value of OfpErrorMsg is invalid.")
	}
}

/*****************************************************/
/* OfpErrorExperimenterMsg                           */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpMultipartRequest                               */
/*****************************************************/
//
// OFPMP_DESC
// OFPMP_FLOW
// OFPMP_AGGREGATE
// OFPMP_TABLE
// OFPMP_PORT_STATS
// OFPMP_QUEUE
// OFPMP_GROUP
// OFPMP_GROUP_DESC
// OFPMP_GROUP_FEATURES
// OFPMP_METER
// OFPMP_METER_CONFIG
// OFPMP_METER_FEATURES
// OFPMP_TABLE_FEATURES
// OFPMP_PORT_DESC
// OFPMP_EXPERIMENTER
/*****************************************************/
/* OfpDesc                                           */
/*****************************************************/
func TestSerializeDescStatsRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x12,       // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, // Type(OFPMP_DESC)
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	mp := NewOfpDescStatsRequest(0)
	actual := mp.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpDescStatsRequest is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpFlowStatsRequest                               */
/*****************************************************/
func TestSerializeFlowStatsRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x12,       // Type
		0x00, 0x40, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x01, // Type(OFPMP_FLOW)
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00,             // Table ID
		0x00, 0x00, 0x00, // Padding
		0xff, 0xff, 0xff, 0xff, // OutPort
		0xff, 0xff, 0xff, 0xff, // OutGroup
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie Mask
		0x00, 0x01, // Match Type(OFPMT_OXM)
		0x00, 0x0e, // Length
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x06,                               // OFPXMT_OFB_ETH_DST, has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	mf := NewOfpMatch()
	ethdst, _ := NewOxmEthDst("11:22:33:44:55:66")
	mf.Append(ethdst)
	mp := NewOfpFlowStatsRequest(0, 0, OFPP_ANY, OFPG_ANY, 0, 0, mf)
	actual := mp.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpFlowStatsRequest is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpAggregateStatsRequest                          */
/*****************************************************/
func TestSerializeAggregateStatsRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x12,       // Type
		0x00, 0x40, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x02, // Type(OFPMP_AGGREGATE)
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00,             // Table ID
		0x00, 0x00, 0x00, // Padding
		0xff, 0xff, 0xff, 0xff, // OutPort
		0xff, 0xff, 0xff, 0xff, // OutGroup
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie Mask
		0x00, 0x01, // Match Type(OFPMT_OXM)
		0x00, 0x0e, // Length
		0x80, 0x00, // Class(OFPXMC_OPENFLOW_BASIC)
		0x06,                               // OFPXMT_OFB_ETH_DST, has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	mf := NewOfpMatch()
	ethdst, _ := NewOxmEthDst("11:22:33:44:55:66")
	mf.Append(ethdst)
	mp := NewOfpAggregateStatsRequest(0, 0, OFPP_ANY, OFPG_ANY, 0, 0, mf)
	actual := mp.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpFlowStatsRequest is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpTableStatsRequest                              */
/*****************************************************/
func TestSerializeTableStatsRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x12,       // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x03, // Type(OFPMP_TABLE)
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	mp := NewOfpTableStatsRequest(0)
	actual := mp.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableStatsRequest is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpPortStatsRequest                               */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueueStatsRequest                              */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpGroupStatsRequest                              */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpGroupDescStatsRequest                          */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpGroupFeaturesStatsRequest                      */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpMeterStatsRequest                              */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpMeterConfigStatsRequest                        */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpMeterFeatruresStatsRequest                     */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableFeatruresStatsRequest                     */
/*****************************************************/
func TestSerializeTableFeaturesStatsRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x12,       // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x0c, // Type(OFPMP_TABLE_FEATURES)
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	mp := NewOfpTableFeaturesStatsRequest(0, nil)
	actual := mp.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturesStatsRequest is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpPortDescStatsRequest                           */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpExperimenterStatsRequest                       */
/*****************************************************/
// TODO: implements and test

/* OfpMultipartReply                                 */
//
// OFPMP_DESC
// OFPMP_FLOW
// OFPMP_AGGREGATE
// OFPMP_TABLE
// OFPMP_PORT_STATS
// OFPMP_QUEUE
// OFPMP_GROUP
// OFPMP_GROUP_DESC
// OFPMP_GROUP_FEATURES
// OFPMP_METER
// OFPMP_METER_CONFIG
// OFPMP_METER_FEATURES
// OFPMP_TABLE_FEATURES
// OFPMP_PORT_DESC
// OFPMP_EXPERIMENTER

/*****************************************************/
/* OfpDesc                                           */
/*****************************************************/
func TestParseOfpDescStatsReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x13,       // Type
		0x04, 0x30, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, // Type(OFPMP_DESC)
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
		0x4e, 0x69, 0x63, 0x69,
		0x72, 0x61, 0x2c, 0x20,
		0x49, 0x6e, 0x63, 0x2e,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, // Manufacturer desc
		0x4f, 0x70, 0x65, 0x6e,
		0x20, 0x76, 0x53, 0x77,
		0x69, 0x74, 0x63, 0x68,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, // Hardware desc
		0x32, 0x2e, 0x30, 0x2e,
		0x32, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, // Software desc
		0x4e, 0x6f, 0x6e, 0x65,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, // SerialNum
		0x4e, 0x6f, 0x6e, 0x65,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, // Datapath desc
	}

	rep := NewOfpMultipartReply()
	rep.Parse(packet)

	mp := rep.Body[0]
	if obj, ok := mp.(*OfpDescStats); ok {
		mfrdesc := (string)(obj.MfrDesc[0:12])
		hwdesc := (string)(obj.HwDesc[0:12])
		swdesc := (string)(obj.SwDesc[0:5])
		serialnum := (string)(obj.SerialNum[0:4])
		dpdesc := (string)(obj.DpDesc[0:4])

		if rep.Header.Version != 4 || rep.Header.Type != 19 ||
			rep.Header.Length != 1072 || rep.Header.Xid != 0 ||
			mfrdesc != "Nicira, Inc." ||
			hwdesc != "Open vSwitch" ||
			swdesc != "2.0.2" ||
			serialnum != "None" ||
			dpdesc != "None" {
			t.Log("Version        : ", rep.Header.Version)
			t.Log("Type           : ", rep.Header.Type)
			t.Log("Length         : ", rep.Header.Length)
			t.Log("Transaction ID : ", rep.Header.Xid)
			t.Log("MfrDesc        : ", mfrdesc)
			t.Log("HwDesc         : ", hwdesc)
			t.Log("SwDesc         : ", swdesc)
			t.Log("SerialNum      : ", serialnum)
			t.Log("DpDesc         : ", dpdesc)
			t.Error("Parsed value of OfpDescStatsReply is invalid.")
		}
	}
}

/*****************************************************/
/* OfpFlowStatsReply                                 */
/*****************************************************/
func TestParseFlowStatsReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x13,       // Type
		0x00, 0x68, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x01, // OFPMP_FLOW
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x58, // Length
		0x00,                   // Table ID
		0x00,                   // Padding
		0x00, 0x00, 0x00, 0x00, // Duration Sec
		0x00, 0x00, 0x00, 0x00, // Duration nSec
		0x00, 0x00, // Priority
		0x00, 0x00, // Idle Timeout
		0x00, 0x00, // Hard Timeout
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Cookie
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Packet count
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Byte count
		0x00, 0x01, // Type(OFPMT_OXM)
		0x00, 0x0e, // Length
		0x80, 0x00, // OFPXMC_OPENFLOW_BASIC
		0x06,                               // OFPXMT_OFB_ETH_DST, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0x00, 0x00, // Padding
		0x00, 0x04, // OFPIT_APPLY_ACTIONS
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x19, // Type(OFPAT_SET_FIELD)
		0x00, 0x10, // Length
		0x80, 0x00, // OFPXMC_OPENFLOW_BASIC
		0x06,                               // OFPXMT_OFB_ETH_DST, Has mask is false
		0x06,                               // Length
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, // Value
		0x00, 0x00, // Padding
	}

	rep := NewOfpMultipartReply()
	rep.Parse(packet)
	if rep.Header.Version != 4 || rep.Header.Type != 19 ||
		rep.Header.Length != 104 || rep.Header.Xid != 0 {
		t.Log("Version        : ", rep.Header.Version)
		t.Log("Type           : ", rep.Header.Type)
		t.Log("Length         : ", rep.Header.Length)
		t.Log("Transaction ID : ", rep.Header.Xid)
		t.Error("Parsed value of OfpFlowStatsReply is invalid.")
	}
	// TODO: check more details
}

/*****************************************************/
/* OfpAggregateStatsReply                            */
/*****************************************************/
func TestParseAgregateStatsReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x13,       // Type
		0x00, 0x28, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x02, // OFPMP_AGGREGATE
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Packet count
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Byte count
		0x00, 0x00, 0x00, 0x01, // Flow Count
		0x00, 0x00, 0x00, 0x00, // Padding
	}

	rep := NewOfpMultipartReply()
	rep.Parse(packet)
	if rep.Header.Version != 4 || rep.Header.Type != 19 ||
		rep.Header.Length != 40 || rep.Header.Xid != 0 {
		t.Log("Version        : ", rep.Header.Version)
		t.Log("Type           : ", rep.Header.Type)
		t.Log("Length         : ", rep.Header.Length)
		t.Log("Transaction ID : ", rep.Header.Xid)
		t.Error("Parsed value of OfpFlowStatsReply is invalid.")
	}
	// TODO: check more details
}

/*****************************************************/
/* OfpTableStatsReply                                */
/*****************************************************/
func TestParseTableStatsReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x13,       // Type
		0x00, 0x28, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x03, // OFPMP_TABLE
		0x00, 0x00, // Flags
		0x00, 0x00, 0x00, 0x00, // Padding
		0x01,             // TableId
		0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x01, // ActiveCount
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // LookupCount
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, // MatchedCount
	}

	rep := NewOfpMultipartReply()
	rep.Parse(packet)
	if rep.Header.Version != 4 || rep.Header.Type != 19 ||
		rep.Header.Length != 40 || rep.Header.Xid != 0 ||
		rep.Body[0].(*OfpTableStats).TableId != 1 ||
		rep.Body[0].(*OfpTableStats).ActiveCount != 1 ||
		rep.Body[0].(*OfpTableStats).LookupCount != 1 ||
		rep.Body[0].(*OfpTableStats).MatchedCount != 1 {
		t.Log("Version        : ", rep.Header.Version)
		t.Log("Type           : ", rep.Header.Type)
		t.Log("Length         : ", rep.Header.Length)
		t.Log("Transaction ID : ", rep.Header.Xid)
		t.Log("TableId        : ", rep.Body[0].(*OfpTableStats).TableId)
		t.Log("ActiveCount    : ", rep.Body[0].(*OfpTableStats).ActiveCount)
		t.Log("LookupCount    : ", rep.Body[0].(*OfpTableStats).LookupCount)
		t.Log("MatchedCount   : ", rep.Body[0].(*OfpTableStats).MatchedCount)
		t.Error("Parsed value of OfpTableStatsReply is invalid.")
	}
}

/*****************************************************/
/* OfpPortStatsReply                                 */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueueStats                                     */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpGroupStatsReply                                */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpGroupDesc                                      */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpGroupFeatures                                  */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpMeterStats                                     */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpMeterConfig                                    */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpMeterFeatures                                  */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableFeaturePropHeader                         */
/*****************************************************/
func TestSerializeTableFeaturePropHeader(t *testing.T) {
	expect := []byte{
		0x00, 0x00, // Type
		0x00, 0x04, // Length
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	p := NewOfpTableFeaturePropHeader(OFPTFPT_INSTRUCTIONS, 4)
	actual := p.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturePropHeader is not equal to expected value.")
	}
}

func TestParseTableFeatureStatsPropHeader(t *testing.T) {
	packet := []byte{
		0x00, 0x00, // Type
		0x00, 0x04, // Length
	}

	p := NewOfpTableFeaturePropHeader(OFPTFPT_INSTRUCTIONS, 0)
	p.Parse(packet)
	if p.Type != OFPTFPT_INSTRUCTIONS || p.Length != 4 {
		t.Log("Type           : ", p.Type)
		t.Log("Length         : ", p.Length)
		t.Error("Parsed value of OfpTableFeaturePropHeader is invalid.")
	}
}

/*****************************************************/
/* OfpTableFeaturePropInstructions                   */
/*****************************************************/
func TestSerializeTableFeaturePropInstructions(t *testing.T) {
	expect := []byte{
		0x00, 0x00, // Type
		0x00, 0x0c, // Length
		0x00, 0x01, // Type(OFPIT_GOTO_TABLE)
		0x00, 0x04, // Length
		0x00, 0x02, // Type(OFPIT_WRITE_METADATA)
		0x00, 0x04, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	ids := make([]*OfpInstructionId, 2)
	ids[0] = NewOfpInstructionId(OFPIT_GOTO_TABLE, 4)
	ids[1] = NewOfpInstructionId(OFPIT_WRITE_METADATA, 4)
	p := NewOfpTableFeaturePropInstructions(
		OFPTFPT_INSTRUCTIONS, ids)
	actual := p.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturePropInstructions is not equal to expected value.")
	}
}

func TestParseTableFeatureStatsPropInstructions(t *testing.T) {
	packet := []byte{
		0x00, 0x00, // Type
		0x00, 0x0c, // Length
		0x00, 0x01, // Type(OFPIT_GOTO_TABLE)
		0x00, 0x04, // Length
		0x00, 0x02, // Type(OFPIT_WRITE_METADATA)
		0x00, 0x04, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
	}

	p := NewOfpTableFeaturePropInstructions(OFPTFPT_INSTRUCTIONS, nil)
	p.Parse(packet)
	if p.PropHeader.Type != OFPTFPT_INSTRUCTIONS ||
		p.PropHeader.Length != 12 ||
		p.InstructionIds[0].Type != OFPIT_GOTO_TABLE ||
		p.InstructionIds[0].Length != 4 ||
		p.InstructionIds[1].Type != OFPIT_WRITE_METADATA ||
		p.InstructionIds[1].Length != 4 {
		t.Log("Type           : ", p.PropHeader.Type)
		t.Log("Length         : ", p.PropHeader.Length)
		t.Log("Id[0] Type     : ", p.InstructionIds[0].Type)
		t.Log("Id[0] Length   : ", p.InstructionIds[0].Length)
		t.Log("Id[1] Type     : ", p.InstructionIds[1].Type)
		t.Log("Id[1] Length   : ", p.InstructionIds[1].Length)
		t.Error("Parsed value of OfpTableFeaturePropInstructions is invalid.")
	}
}

/*****************************************************/
/* OfpTableFeaturePropNextTables                     */
/*****************************************************/
func TestSerializeTableFeaturePropNextTables(t *testing.T) {
	expect := []byte{
		0x00, 0x02, // Type
		0x00, 0x07, // Length
		0x01,
		0x02,
		0x03,
		0x00, //Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	p := NewOfpTableFeaturePropNextTables(
		OFPTFPT_NEXT_TABLES, []uint8{1, 2, 3})
	actual := p.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturePropNextTables is not equal to expected value.")
	}
}

func TestParseTableFeatureStatsPropNextTables(t *testing.T) {
	packet := []byte{
		0x00, 0x02, // Type
		0x00, 0x07, // Length
		0x01,
		0x02,
		0x03,
		0x00, //Padding
	}

	p := NewOfpTableFeaturePropNextTables(OFPTFPT_NEXT_TABLES, nil)
	p.Parse(packet)
	if p.PropHeader.Type != OFPTFPT_NEXT_TABLES ||
		p.PropHeader.Length != 7 ||
		p.NextTableIds[0] != 1 ||
		p.NextTableIds[1] != 2 ||
		p.NextTableIds[2] != 3 {
		t.Log("Type           : ", p.PropHeader.Type)
		t.Log("Length         : ", p.PropHeader.Length)
		t.Log("Id[0]          : ", p.NextTableIds[0])
		t.Log("Id[1]          : ", p.NextTableIds[1])
		t.Log("Id[2]          : ", p.NextTableIds[2])
		t.Error("Parsed value of OfpTableFeaturePropNextTables is invalid.")
	}
}

/*****************************************************/
/* OfpTableFeaturePropActions                        */
/*****************************************************/
func TestSerializeTableFeaturePropActions(t *testing.T) {
	expect := []byte{
		0x00, 0x04, // Type
		0x00, 0x14, // Length
		0x00, 0x00, // Type(OFPAT_OUTPUT)
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x0b, // Type(OFPAT_COPY_TTL_OUT)
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	ids := make([]OfpActionHeader, 2)
	ids[0] = NewOfpActionHeader(OFPAT_OUTPUT, 8)
	ids[1] = NewOfpActionHeader(OFPAT_COPY_TTL_OUT, 8)
	p := NewOfpTableFeaturePropActions(OFPTFPT_WRITE_ACTIONS, ids)
	actual := p.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturePropActions is not equal to expected value.")
	}
}

func TestParseTableFeatureStatsPropActions(t *testing.T) {
	packet := []byte{
		0x00, 0x04, // Type
		0x00, 0x14, // Length
		0x00, 0x00, // Type(OFPAT_OUTPUT)
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x0b, // Type(OFPAT_COPY_TTL_OUT)
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, // Padding
	}

	p := NewOfpTableFeaturePropActions(OFPTFPT_WRITE_ACTIONS, nil)
	p.Parse(packet)
	if p.PropHeader.Type != OFPTFPT_WRITE_ACTIONS ||
		p.PropHeader.Length != 20 ||
		p.ActionIds[0].Type != OFPAT_OUTPUT ||
		p.ActionIds[0].Length != 8 ||
		p.ActionIds[1].Type != OFPAT_COPY_TTL_OUT ||
		p.ActionIds[1].Length != 8 {
		t.Log("Type               : ", p.PropHeader.Type)
		t.Log("Length             : ", p.PropHeader.Length)
		t.Log("ActionId[0] Type   : ", p.ActionIds[0].Type)
		t.Log("ActionId[0] Length : ", p.ActionIds[0].Length)
		t.Log("ActionId[1] Type   : ", p.ActionIds[1].Type)
		t.Log("ActionId[1] Length : ", p.ActionIds[1].Length)
		t.Error("Parsed value of OfpTableFeaturePropActions is invalid.")
	}
}

/*****************************************************/
/* OfpTableFeaturePropOxm                            */
/*****************************************************/
func TestSerializeTableFeaturePropOxm(t *testing.T) {
	expect := []byte{
		0x00, 0x08, // Type
		0x00, 0x0c, // Length
		0x80, 0x00, 0x00, 0x04, // OXM_OF_IN_PORT
		0x80, 0x00, 0x02, 0x04, // OXM_OF_IN_PHY_PORT
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	ids := make([]uint32, 2)
	ids[0] = OXM_OF_IN_PORT
	ids[1] = OXM_OF_IN_PHY_PORT
	p := NewOfpTableFeaturePropOxm(OFPTFPT_MATCH, ids)
	actual := p.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturePropOxm is not equal to expected value.")
	}
}

func TestParseTableFeatureStatsPropOxm(t *testing.T) {
	packet := []byte{
		0x00, 0x08, // Type
		0x00, 0x0c, // Length
		0x80, 0x00, 0x00, 0x04, // OXM_OF_IN_PORT
		0x80, 0x00, 0x02, 0x04, // OXM_OF_IN_PHY_PORT
		0x00, 0x00, 0x00, 0x00, // Padding
	}

	p := NewOfpTableFeaturePropOxm(OFPTFPT_MATCH, nil)
	p.Parse(packet)
	if p.PropHeader.Type != OFPTFPT_MATCH ||
		p.PropHeader.Length != 12 ||
		p.OxmIds[0] != OXM_OF_IN_PORT ||
		p.OxmIds[1] != OXM_OF_IN_PHY_PORT {
		t.Log("Type               : ", p.PropHeader.Type)
		t.Log("Length             : ", p.PropHeader.Length)
		t.Log("OxmId[0]           : ", p.OxmIds[0])
		t.Log("OxmId[1]           : ", p.OxmIds[1])
		t.Error("Parsed value of OfpTableFeaturePropOxm is invalid.")
	}
}

/*****************************************************/
/* OfpTableFeaturePropExperimenter                   */
/*****************************************************/
func TestSerializeTableFeaturePropExperimenter(t *testing.T) {
	expect := []byte{
		0xff, 0xfe, // Type
		0x00, 0x14, // Length
		0x00, 0x00, 0x00, 0x01, // Experimenter
		0x00, 0x00, 0x00, 0x01, // ExpType
		0x00, 0x00, 0x00, 0x01, // Data
		0x00, 0x00, 0x00, 0x02, // Data
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	data := make([]uint32, 2)
	data[0] = 1
	data[1] = 2
	p := NewOfpTableFeaturePropExperimenter(OFPTFPT_EXPERIMENTER, 1, 1, data)
	actual := p.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturePropExperimenter is not equal to expected value.")
	}
}

func TestParseTableFeatureStatsPropExperimenter(t *testing.T) {
	packet := []byte{
		0xff, 0xfe, // Type
		0x00, 0x14, // Length
		0x00, 0x00, 0x00, 0x01, // Experimenter
		0x00, 0x00, 0x00, 0x01, // ExpType
		0x00, 0x00, 0x00, 0x01, // Data
		0x00, 0x00, 0x00, 0x02, // Data
		0x00, 0x00, 0x00, 0x00, // Padding
	}

	p := NewOfpTableFeaturePropExperimenter(OFPTFPT_EXPERIMENTER, 0, 0, nil)
	p.Parse(packet)
	if p.PropHeader.Type != OFPTFPT_EXPERIMENTER ||
		p.PropHeader.Length != 20 ||
		p.ExperimenterData[0] != 1 ||
		p.ExperimenterData[1] != 2 {
		t.Log("Type               : ", p.PropHeader.Type)
		t.Log("Length             : ", p.PropHeader.Length)
		t.Log("Data[0]            : ", p.ExperimenterData[0])
		t.Log("Data[1]            : ", p.ExperimenterData[1])
		t.Error("Parsed value of OfpTableFeaturePropExperimenter is invalid.")
	}
}

/*****************************************************/
/* OfpTableFeatures                                  */
/*****************************************************/
func TestSerializeTableFeatures(t *testing.T) {
	expect := []byte{
		0x00, 0x58, // Length
		0x00,                         // TableId
		0x00, 0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Name
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // MetadataMatch
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // MetadataWrite
		0x00, 0x00, 0x00, 0x00, // Config
		0x00, 0x00, 0x00, 0x00, // MaxEntries
		0x00, 0x00, // Type(OFPTFPT_INSTRUCTIONS)
		0x00, 0x0c, // Length
		0x00, 0x01, // Type(OFPIT_GOTO_TABLE)
		0x00, 0x04, // Length
		0x00, 0x02, // Type(OFPIT_WRITE_METADATA)
		0x00, 0x04, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x02, // Type(OFPTFPT_NEXT_TABLES)
		0x00, 0x07, // Length
		0x01,
		0x02,
		0x03,
		0x00, //Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	ids := make([]*OfpInstructionId, 2)
	ids[0] = NewOfpInstructionId(OFPIT_GOTO_TABLE, 4)
	ids[1] = NewOfpInstructionId(OFPIT_WRITE_METADATA, 4)
	pi := NewOfpTableFeaturePropInstructions(
		OFPTFPT_INSTRUCTIONS, ids)

	pn := NewOfpTableFeaturePropNextTables(
		OFPTFPT_NEXT_TABLES, []uint8{1, 2, 3})

	props := make([]OfpTableFeatureProp, 2)
	props[0] = pi
	props[1] = pn

	p := NewOfpTableFeatures(
		0,
		nil,
		0,
		0,
		0,
		0,
		props)

	actual := p.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpTableFeaturesStats is not equal to expected value.")
	}
}

func TestParseTableFeature(t *testing.T) {
	packet := []byte{
		0x00, 0x58, // Length
		0x00,                         // TableId
		0x00, 0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, //
		0x00, 0x00, 0x00, 0x00, // Name
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // MetadataMatch
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // MetadataWrite
		0x00, 0x00, 0x00, 0x00, // Config
		0x00, 0x00, 0x00, 0x00, // MaxEntries
		0x00, 0x00, // Type(OFPTFPT_INSTRUCTIONS)
		0x00, 0x0c, // Length
		0x00, 0x01, // Type(OFPIT_GOTO_TABLE)
		0x00, 0x04, // Length
		0x00, 0x02, // Type(OFPIT_WRITE_METADATA)
		0x00, 0x04, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x02, // Type(OFPTFPT_NEXT_TABLES)
		0x00, 0x07, // Length
		0x01,
		0x02,
		0x03,
		0x00, //Padding
	}

	p := NewOfpTableFeatures(
		0,
		nil,
		0,
		0,
		0,
		0,
		nil)
	p.Parse(packet)
	if p.Length != 88 ||
		p.TableId != 0 ||
		p.MetadataMatch != 0 ||
		p.MetadataWrite != 0 ||
		p.Config != 0 ||
		p.MaxEntries != 0 ||
		p.Properties[0].(*OfpTableFeaturePropInstructions).
			PropHeader.Type !=
			OFPTFPT_INSTRUCTIONS ||
		p.Properties[0].(*OfpTableFeaturePropInstructions).
			PropHeader.Length != 12 ||
		p.Properties[0].(*OfpTableFeaturePropInstructions).
			InstructionIds[0].Type != OFPIT_GOTO_TABLE ||
		p.Properties[0].(*OfpTableFeaturePropInstructions).
			InstructionIds[0].Length != 4 ||
		p.Properties[0].(*OfpTableFeaturePropInstructions).
			InstructionIds[1].Type != OFPIT_WRITE_METADATA ||
		p.Properties[0].(*OfpTableFeaturePropInstructions).
			InstructionIds[1].Length != 4 ||
		p.Properties[1].(*OfpTableFeaturePropNextTables).
			PropHeader.Type != OFPTFPT_NEXT_TABLES ||
		p.Properties[1].(*OfpTableFeaturePropNextTables).
			PropHeader.Length != 7 ||
		p.Properties[1].(*OfpTableFeaturePropNextTables).
			NextTableIds[0] != 1 ||
		p.Properties[1].(*OfpTableFeaturePropNextTables).
			NextTableIds[1] != 2 ||
		p.Properties[1].(*OfpTableFeaturePropNextTables).
			NextTableIds[2] != 3 {
		t.Log("Length             : ", p.Length)
		t.Log("TableId            : ", p.TableId)
		t.Log("MetadataMatch      : ", p.MetadataMatch)
		t.Log("MetadataWrite      : ", p.MetadataWrite)
		t.Log("Config             : ", p.Config)
		t.Log("MaxEntries         : ", p.MaxEntries)
		t.Error("Parsed value of OfpTableFeaturesStats is invalid.")
	}
}

/*****************************************************/
/* OfpExperimenterMultipartHeader                    */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueuePropHeader                                */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueuePropMinRate                               */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueuePropMaxRate                               */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueuePropExperimenter                          */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpPacketQueue                                    */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueueGetConfigRequest                          */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpActionSetQueue                                 */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueueStatsRequest                              */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpQueueGetConfigRequest                          */
/*****************************************************/
func TestSerializeQueueGetConfigRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x16,       // Type
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x01, // Port
		0x00, 0x00, 0x00, 0x00, // Padding
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	m := NewOfpQueueGetConfigRequest(1)
	actual := m.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpQueueGetConfigRequest is not equal to expected value.")
	}
}

/*****************************************************/
/* OfpQueueGetConfigReply                            */
/*****************************************************/
func TestParseQueueGetConfigReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x17,       // Type
		0x00, 0x30, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x01, // Port
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x01, // QueueId
		0x00, 0x00, 0x00, 0x01, // Port
		0x00, 0x20, // Length
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x01, // Property
		0x00, 0x10, // Length
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x01, // Rate
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // Padding
	}

	m := NewOfpQueueGetConfigReply()
	m.Parse(packet)
	if m.Header.Version != 4 || m.Header.Type != 23 ||
		m.Header.Length != 48 || m.Header.Xid != 0 ||
		m.Port != 1 ||
		m.Queue[0].QueueId != 1 ||
		m.Queue[0].Port != 1 ||
		m.Queue[0].Length != 32 ||
		m.Queue[0].Properties[0].(*OfpQueuePropMinRate).Property() != OFPQT_MIN_RATE ||
		m.Queue[0].Properties[0].(*OfpQueuePropMinRate).Rate != 1 {
		t.Log("Version        : ", m.Header.Version)
		t.Log("Type           : ", m.Header.Type)
		t.Log("Length         : ", m.Header.Length)
		t.Log("Transaction ID : ", m.Header.Xid)
		t.Error("Parsed value of OfpQueueGetConfigReply is invalid.")
	}
}

/*****************************************************/
/* OfpRoleRequest                                    */
/*****************************************************/
func TestSerializeRoleRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x18,       // Type
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x02, // Role
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // GenerationId
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	m := NewOfpRoleRequest(OFPCT_ROLE_MASTER, 0)
	actual := m.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpRoleRequest is not equal to expected value.")
	}
}

func TestParseRoleReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x19,       // Type
		0x00, 0x18, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x02, // Role
		0x00, 0x00, 0x00, 0x00, // Padding
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, // GenerationId
	}

	m := NewOfpRoleReply()
	m.Parse(packet)
	if m.Header.Version != 4 || m.Header.Type != 25 ||
		m.Header.Length != 24 || m.Header.Xid != 0 ||
		m.Role != OFPCT_ROLE_MASTER ||
		m.GenerationId != 0 {
		t.Log("Version        : ", m.Header.Version)
		t.Log("Type           : ", m.Header.Type)
		t.Log("Length         : ", m.Header.Length)
		t.Log("Transaction ID : ", m.Header.Xid)
		t.Log("Role           : ", m.Role)
		t.Log("GenerationId   : ", m.GenerationId)
		t.Error("Parsed value of OfpRoleReply is invalid.")
	}
}

/*****************************************************/
/* OfpAsyncConfig                                    */
/*****************************************************/
func TestSerializeGetAsyncRequest(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x1a,       // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	m := NewOfpGetAsyncRequest()
	actual := m.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpGetAsyncRequest is not equal to expected value.")
	}
}

func TestParseGetAsyncReply(t *testing.T) {
	packet := []byte{
		0x04,       // Version
		0x1b,       // Type
		0x00, 0x20, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x04, // PacketInMask
		0x00, 0x00, 0x00, 0x00, // PacketInMask
		0x00, 0x00, 0x00, 0x01, // PortStatusMask
		0x00, 0x00, 0x00, 0x00, // PortStatusMask
		0x00, 0x00, 0x00, 0x01, // FlowRemovedMask
		0x00, 0x00, 0x00, 0x00, // FlowRemovedMask
	}

	m := NewOfpGetAsyncReply()
	m.Parse(packet)
	if m.Header.Version != 4 || m.Header.Type != 27 ||
		m.Header.Length != 32 || m.Header.Xid != 0 ||
		m.PacketInMask[0] != 4 ||
		m.PacketInMask[1] != 0 ||
		m.PortStatusMask[0] != 1 ||
		m.PortStatusMask[1] != 0 ||
		m.FlowRemovedMask[0] != 1 ||
		m.FlowRemovedMask[1] != 0 {
		t.Log("Version        : ", m.Header.Version)
		t.Log("Type           : ", m.Header.Type)
		t.Log("Length         : ", m.Header.Length)
		t.Log("Transaction ID : ", m.Header.Xid)
		t.Log("PacketInMask   : ", m.PacketInMask[0])
		t.Log("PacketInMask   : ", m.PacketInMask[1])
		t.Log("PortStatusMask : ", m.PortStatusMask[0])
		t.Log("PortStatusMask : ", m.PortStatusMask[1])
		t.Log("FlowRemovedMask: ", m.FlowRemovedMask[0])
		t.Log("FlowRemovedMask: ", m.FlowRemovedMask[1])
		t.Error("Parsed value of OfpGetAsyncReply is invalid.")
	}
}

func TestSerializeSetAsync(t *testing.T) {
	expect := []byte{
		0x04,       // Version
		0x1c,       // Type
		0x00, 0x20, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
		0x00, 0x00, 0x00, 0x04, // PacketInMask
		0x00, 0x00, 0x00, 0x00, // PacketInMask
		0x00, 0x00, 0x00, 0x01, // PortStatusMask
		0x00, 0x00, 0x00, 0x00, // PortStatusMask
		0x00, 0x00, 0x00, 0x01, // FlowRemovedMask
		0x00, 0x00, 0x00, 0x00, // FlowRemovedMask
	}
	e_str := hex.EncodeToString(expect)

	// reset xid for test
	xid = 0

	m := NewOfpSetAsync(
		[2]uint32{4, 0},
		[2]uint32{1, 0},
		[2]uint32{1, 0})
	actual := m.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str {
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpSetAsync is not equal to expected value.")
	}
}
