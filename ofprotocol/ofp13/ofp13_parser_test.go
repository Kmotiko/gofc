package ofp13

import (
	"encoding/hex"
	"testing"
)

func TestSerializeHello(t *testing.T){
	expect := []byte{
		0x04, // Version
		0x00, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}
	e_str := hex.EncodeToString(expect)

	hello := NewOfpHello()
	actual := hello.Serialize()
	a_str := hex.EncodeToString(actual)
	if len(expect) != len(actual) || e_str != a_str{
		t.Log("Expected Value is : ", e_str)
		t.Log("Actual Value is   : ", a_str)
		t.Error("Serialized binary of OfpHello is not equal to expected value.")
	}
}

func TestParseHello(t *testing.T){
	packet := []byte{
		0x04, // Version
		0x00, // Type
		0x00, 0x08, // Length
		0x00, 0x00, 0x00, 0x00, // Transaction ID
	}

	hello := NewOfpHello()
	hello.Parse(packet)
	if hello.Header.Version != 4 || hello.Header.Type != 0 ||
	hello.Header.Length != 8 || hello.Header.Xid != 0{
		t.Log("Version        : ", hello.Header.Version)
		t.Log("Type           : ", hello.Header.Type)
		t.Log("Length         : ", hello.Header.Length)
		t.Log("Transaction ID : ", hello.Header.Xid)
		t.Error("Parsed value of OfpHello is invalid.")
	}
}


// TODO : test all of ofp13 messages
