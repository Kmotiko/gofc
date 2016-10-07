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

// TODO : test all other ofp13 messages

/*****************************************************/
/* OfpSwitchConfig                                   */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableMod                                       */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpPortStatus                                     */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpPortMod                                        */
/*****************************************************/
// TODO: implements and test

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
// OFPIT_WRITE_METADATA
// OFPIT_WRITE_ACTIONS
// OFPIT_APPLY_ACTIONS
// OFPIT_CLEAR_ACTIONS
// OFPIT_METER
// OFPIT_EXPERIMENTER

// Actions
// OFPAT_OUTPUT
// OFPAT_COPY_TTL_OUT
// OFPAT_COPY_TTL_IN
// OFPAT_SET_MPLS_TTL
// OFPAT_DEC_MPLS_TTL
// OFPAT_PUSH_VLAN
// OFPAT_POP_VLAN
// OFPAT_PUSH_MPLS
// OFPAT_POP_MPLS
// OFPAT_SET_QUEUE
// OFPAT_GROUP
// OFPAT_SET_NW_TTL
// OFPAT_DEC_NW_TTL
// OFPAT_SET_FIELD
// OFPAT_PUSH_PBB
// OFPAT_POP_PBB
// OFPAT_EXPERIMENTER

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
// TODO: implements and test

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
// TODO: implements and test

/*****************************************************/
/* OfpErrorMsg                                       */
/*****************************************************/
// TODO: implements and test

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
// TODO: implements and test

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
// TODO: implements and test

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
// TODO: implements and test

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
// TODO: implements and test

/*****************************************************/
/* OfpTableFeaturePropInstructions                   */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableFeaturePropNextTables                     */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableFeaturePropActions                        */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableFeaturePropOxm                            */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableFeaturePropExperimenter                   */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpTableFeatures                                  */
/*****************************************************/
// TODO: implements and test

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
/* OfpRoleRequest                                    */
/*****************************************************/
// TODO: implements and test

/*****************************************************/
/* OfpAsyncConfig                                    */
/*****************************************************/
// TODO: implements and test
