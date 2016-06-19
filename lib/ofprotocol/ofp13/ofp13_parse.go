package ofp13

import(
	"encoding/binary"
)

func Parse(packet []byte) (msg OFMessage){
	switch packet[1]{
	case OFPT_HELLO:
		msg = new(OfpHello)
		msg.Parse(packet)
	case OFPT_FEATURES_REPLY:
		msg = NewOfpFeaturesReply()
		msg.Parse(packet)
	case OFPT_ECHO_REQUEST:
		msg = NewOfpEchoRequest()
		msg.Parse(packet)
	case OFPT_ECHO_REPLY:
		msg = NewOfpEchoReply()
		msg.Parse(packet)
	case OFPT_PACKET_IN:
		msg = NewOfpPacketIn()
		msg.Parse(packet)
	default:
	}
	return msg
}



/*****************************************************/
/* OfpHeader                                         */
/*****************************************************/
func NewOfpHeader() OfpHeader{
	// 4 means ofp version 1.3
	h := OfpHeader{4, 0, 8, 0}
	return h
}

func (h *OfpHeader)Serialize() []byte{
	packet := make([]byte, 8)
	packet[0] = h.Version
	packet[1] = h.Type
	binary.BigEndian.PutUint16(packet[2:], h.Length)
	binary.BigEndian.PutUint32(packet[4:], h.Xid)
	return packet;
}

func (h *OfpHeader)Parse(packet []byte) {
	h.Version = packet[0]
	h.Type = packet[1]
	h.Length = binary.BigEndian.Uint16(packet[2:])
	h.Xid = binary.BigEndian.Uint32(packet[4:])
}

func (h *OfpHeader)Size() int{
	return 8
}



/*****************************************************/
/* OfpHelloElemHeader                                */
/*****************************************************/
func NewOfpHelloElemHeader() *OfpHelloElemHeader{
	e := new(OfpHelloElemHeader)
	e.Length = 8
	return e
}


func (h *OfpHelloElemHeader)Serialize() []byte{
	packet := make([]byte, 8)
	binary.BigEndian.PutUint16(packet[0:], h.Type)
	binary.BigEndian.PutUint16(packet[2:], h.Length)

	return packet
}


func (h *OfpHelloElemHeader)Parse(packet []byte){
	h.Type = binary.BigEndian.Uint16(packet[0:])
	h.Length = binary.BigEndian.Uint16(packet[2:])
}


func (h *OfpHelloElemHeader)Size() int{
	return 8
}



/*****************************************************/
/* OfpHello                                          */
/*****************************************************/
func NewOfpHello() *OfpHello{
	hello := new(OfpHello)
	hello.Header = NewOfpHeader()
	hello.Elements = make([]OfpHelloElemHeader, 0)
	return hello;
}

///
///
func (m *OfpHello)Serialize() []byte{
	packet := make([]byte, m.Size())
	// header
	h_packet := m.Header.Serialize()
	// append header
	copy(packet[0:], h_packet)

	// serialize hello body
	index := len(h_packet)
	e_packet := make([]byte, 0)
	for _, elem := range m.Elements{
		e_packet = elem.Serialize()
		copy(packet[index:], elem.Serialize())
		index += len(e_packet)
	}

	return packet;
}

func (m *OfpHello)Parse(packet []byte){
	m.Header.Parse(packet[0:])
	index := 8

	for index < len(packet){
		e := NewOfpHelloElemHeader()
		e.Parse(packet[index:])
		index += e.Size()
		// m.Elements = append(m.Elements, e)
	}
	return
}


func (m *OfpHello)Size() int{
	size := m.Header.Size()
	for _,e := range m.Elements{
		size += e.Size()
	}
	return size
}


/*****************************************************/
/* OfpFeaturesRequest                                */
/*****************************************************/
func NewOfpFeaturesRequest() *OfpHeader{
	m := NewOfpHeader()
	m.Type = OFPT_FEATURES_REQUEST
	return &m
}


/*****************************************************/
/* OfpSwitchFeatures                                 */
/*****************************************************/
func NewOfpFeaturesReply() *OfpSwitchFeatures{
	m := new(OfpSwitchFeatures)
	m.Header = NewOfpHeader()
	return m
}


func (m *OfpSwitchFeatures)Serialize() []byte{
	packet := make([]byte, m.Size())
	h_packet := m.Header.Serialize()
	copy(packet[0:], h_packet)
	index := m.Header.Size()
	binary.BigEndian.PutUint64(packet[index:8], m.DatapathId)
	index += 8
	binary.BigEndian.PutUint32(packet[index:4], m.NBuffers)
	index += 4
	packet[index] = m.NTables
	index += 1
	packet[index] = m.AuxiliaryId
	index += 1
	packet[index] = m.Pad[0]
	index += 1
	packet[index] = m.Pad[1]
	index += 1
	binary.BigEndian.PutUint32(packet[index:4], m.Capabilities)
	index += 4
	binary.BigEndian.PutUint32(packet[index:4], m.Reserved)

	return packet
}


func (m *OfpSwitchFeatures)Parse(packet []byte){
	m.Header.Parse(packet)
	index := m.Header.Size()
	m.DatapathId = binary.BigEndian.Uint64(packet[index:])
	index += 8
	m.NBuffers = binary.BigEndian.Uint32(packet[index:])
	index += 4
	m.NTables = packet[index]
	index += 1
	m.AuxiliaryId = packet[index]
	index += 1
	m.Pad[0] = packet[index]
	index += 1
	m.Pad[1] = packet[index]
	index += 1
	m.Capabilities = binary.BigEndian.Uint32(packet[index:])
	index += 4
	m.Reserved = binary.BigEndian.Uint32(packet[index:])
	index += 4
}

func (m *OfpSwitchFeatures)Size() int{
	return m.Header.Size() + 24
}


/*****************************************************/
/* OfpPacketIn                                       */
/*****************************************************/
func NewOfpPacketIn() *OfpPacketIn{
	m := new(OfpPacketIn)
	m.Header = NewOfpHeader()
	m.Header.Type = OFPT_PACKET_IN
	// m.Match = NewOfpMatch()
	return m
}

func (m *OfpPacketIn)Serialize() []byte{
	packet := make([]byte, m.Size())
	h_packet := m.Header.Serialize()
	copy(packet[0:], h_packet)
	index := m.Header.Size()

	binary.BigEndian.PutUint32(packet[index:4], m.BufferId)
	index += 4
	binary.BigEndian.PutUint16(packet[index:2], m.TotalLen)
	index += 2
	packet[index] = m.Reason
	index++
	packet[index] = m.TableId
	index++

	m_packet := m.Match.Serialize()
	copy(packet[index:], m_packet)

	return packet
}


func (m *OfpPacketIn)Parse(packet []byte){
	m.Header.Parse(packet)
	index := m.Header.Size()

	m.BufferId = binary.BigEndian.Uint32(packet[index:])
	index += 4
	m.TotalLen = binary.BigEndian.Uint16(packet[index:])
	index += 2
	m.Reason = packet[index]
	index++
	m.TableId = packet[index]
	index++
	m.Cookie = binary.BigEndian.Uint64(packet[index:])
	index += 8

	// parse match field
	m.Match.Parse(packet[index:])
}

func (m *OfpPacketIn)Size() int{
	return m.Header.Size() + 16 + m.Match.Size() + 2 + len(m.Data)
}


/*****************************************************/
/* OfpMatch                                          */
/*****************************************************/
func NewOfpMatch() *OfpMatch{
	m := new(OfpMatch)
	m.OxmFields = make([]OxmField, 0)
	return m
}


func (m *OfpMatch)Serialize() []byte{
	packet := make([]byte, m.Size())
	index := 0
	binary.BigEndian.PutUint16(packet[index:], m.Type)
	index++
	binary.BigEndian.PutUint16(packet[index:], m.Type)
	index++
	for _, e := range m.OxmFields{
		binary.BigEndian.PutUint16(packet[index:], e.Class)
		index += 2
		tmp := e.Type << 1
		if e.HasMask == true {
			tmp += 1
		}
		packet[index] = tmp
		index++
		packet[index] = e.Length
		index++
		copy(packet[index:], e.Value)
		index += len(e.Value)
		copy(packet[index:], e.Pad)
		index += len(e.Pad)
	}
	return packet
}

func (m *OfpMatch)Parse(packet []byte) {
	index := 0
	m.Type = binary.BigEndian.Uint16(packet[index:])
	index += 2
	m.Length = binary.BigEndian.Uint16(packet[index:])
	index += 2

	for index < (int(m.Length) - 4){
		oxm := new(OxmField)
		oxm.Class = binary.BigEndian.Uint16(packet[index:])
		index += 2

		tmp := packet[index]
		oxm.Type = tmp >> 1
		oxm.HasMask = false
		if tmp & 0x01 > 0{
			oxm.HasMask = true
		}
		index += 1

		oxm.Length = packet[index]
		index += 1
		oxm.Value = make([]byte, oxm.Length)
		copy(oxm.Value[0:], packet[index:])
		index += int(oxm.Length)

		m.OxmFields = append(m.OxmFields, *oxm)
	}
}

func (m *OfpMatch)Size() int{
	size := 4
	for _,e := range m.OxmFields{
		size += 4
		size += len(e.Value)
	}
	size += (size % 8)
	return size
}


/*****************************************************/
/* OfpErrorMsg                                       */
/*****************************************************/
func NewOfpErrorMsg() *OfpErrorMsg{
	header := NewOfpHeader()
	header.Type = OFPT_ERROR
	return nil
}

func (m *OfpErrorMsg)Serialize() []byte{
	packet := make([]byte, m.Size())
	h_packet := m.Header.Serialize()
	copy(packet[0:], h_packet)
	index := m.Header.Size()
	binary.BigEndian.PutUint16(packet[index:], m.Type)
	index += 2
	binary.BigEndian.PutUint16(packet[index:], m.Code)
	index += 2
	for _,d := range m.Data{
		packet[index] = d
		index += 1
	}
	return packet
}

func (m *OfpErrorMsg)Parse(packet []byte){
	m.Header.Parse(packet)
	index := m.Header.Size()
	m.Type = binary.BigEndian.Uint16(packet[index:])
	index += 2
	m.Code = binary.BigEndian.Uint16(packet[index:])
	index += 2
	for int(index) < len(packet){
		m.Data = append(m.Data, packet[index])
		index += 1
	}
}

func (m *OfpErrorMsg)Size() int{
	return m.Header.Size() + 8 + len(m.Data)
}



/*****************************************************/
/* Echo Message                                   */
/*****************************************************/
func NewOfpEchoRequest() *OfpHeader{
	echo := NewOfpHeader()
	echo.Type = OFPT_ECHO_REQUEST
	return &echo
}

func NewOfpEchoReply() *OfpHeader{
	echo := NewOfpHeader()
	echo.Type = OFPT_ECHO_REPLY
	return &echo
}
