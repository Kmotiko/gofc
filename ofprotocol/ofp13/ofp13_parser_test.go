package ofp13

import (
	"encoding/hex"
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
//  in_phy_port		OFPXMT_OFB_IN_PHY_PORT
//  metadata		OFPXMT_OFB_METADATA
//  eth_dst			OFPXMT_OFB_ETH_DST
//  eth_src			OFPXMT_OFB_ETH_SRC
//  eth_type		OFPXMT_OFB_ETH_TYPE
//  vlan_vid		OFPXMT_OFB_VLAN_VID
//  vlan_pcp		OFPXMT_OFB_VLAN_PCP
//  ip_dscp			OFPXMT_OFB_IP_DSCP
//  ip_ecn			OFPXMT_OFB_IP_ECN
//  ip_proto		OFPXMT_OFB_IP_PROTO
//  ipv4_src		OFPXMT_OFB_IPV4_SRC
//  ipv4_dst		OFPXMT_OFB_IPV4_DST
//  tcp_src			OFPXMT_OFB_TCP_SRC
//  tcp_dst			OFPXMT_OFB_TCP_DST
//  udp_src			OFPXMT_OFB_UDP_SRC
//  udp_dst			OFPXMT_OFB_UDP_DST
//  sctp_src		OFPXMT_OFB_SCTP_SRC
//  sctp_dst		OFPXMT_OFB_SCTP_DST
//  icmpv4_typ		OFPXMT_OFB_ICMPV4_TYPE
//  icmpv4_code		OFPXMT_OFB_ICMPV4_CODE
//  arp_op			OFPXMT_OFB_ARP_OP
//  arp_spa			OFPXMT_OFB_ARP_SPA
//  arp_tpa			OFPXMT_OFB_ARP_TPA
//  arp_sha			OFPXMT_OFB_ARP_SHA
//  arp_tha			OFPXMT_OFB_ARP_THA
//  ipv6_src		OFPXMT_OFB_IPV6_SRC
//  ipv6_dst		OFPXMT_OFB_IPV6_DST
//  ipv6_flabel		OFPXMT_OFB_IPV6_FLABEL
//  icmpv6_type		OFPXMT_OFB_ICMPV6_TYPE
//  icmpv6_code		OFPXMT_OFB_ICMPV6_CODE
//  ipv6_nd_target	OFPXMT_OFB_IPV6_ND_TARGET
//  ipv6_nd_sll		OFPXMT_OFB_IPV6_ND_SLL
//  ipv6_nd_tll		OFPXMT_OFB_IPV6_ND_TLL
//  mpls_label		OFPXMT_OFB_MPLS_LABEL
//  mpls_tc			OFPXMT_OFB_MPLS_TC
//  mpls_bos		OFPXMT_OFB_MPLS_BOS
//  pbb_isid		OFPXMT_OFB_PBB_ISID
//  tunnel_id		OFPXMT_OFB_TUNNEL_ID
//  ipv6_exthdr		OFPXMT_OFB_IPV6_EXTHDR

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
// TODO: implements and test

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
	// TODO: check more details
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
