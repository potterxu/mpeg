package ts_test

import (
	"reflect"
	"testing"

	"github.com/32bitkid/bitreader"
	"github.com/potterxu/mpeg/ts"
)

func TestAdaptationField(t *testing.T) {
	reader := bitreader.NewReader(adaptationFieldReader())
	packet, err := ts.NewPacket(reader)
	if err != nil {
		t.Fatal(err)
	}

	if expected, actual := packet.AdaptationFieldControl, ts.FieldThenPayload; actual != expected {
		t.Fatalf("unexpected AdaptationFieldControl. expected %v, got %v", actual, expected)
	}

	if packet.AdaptationField == nil {
		t.Fatal("exptected adaptation field to be set")
	}

	if expected, actual := 80, len(packet.Payload); actual != expected {
		t.Fatalf("payload was not the correct size. expected %v, got %v", expected, actual)
	}

}

func TestTransPrivateDataField(t *testing.T) {
	reader := bitreader.NewReader(detailedAdaptationFieldReader())
	packet, err := ts.NewPacket(reader)
	if err != nil {
		t.Fatal(err)
	}

	if expected, actual := packet.AdaptationFieldControl, ts.FieldThenPayload; actual != expected {
		t.Fatalf("unexpected AdaptationFieldControl. expected %v, got %v", actual, expected)
	}

	if packet.AdaptationField == nil {
		t.Fatal("exptected adaptation field to be set")
	}

	if expected, actual := false, packet.AdaptationField.DiscontinuityIndicator; actual != expected {
		t.Fatalf("unexpected DiscontinuityIndicator. expected %v, got %v", expected, actual)
	}

	if expected, actual := true, packet.AdaptationField.RandomAccessIndicator; actual != expected {
		t.Fatalf("unexpected RandomAccessIndicator. expected %v, got %v", expected, actual)
	}

	if expected, actual := true, packet.AdaptationField.ElementaryStreamPriorityIndicator; actual != expected {
		t.Fatalf("unexpected ElementaryStreamPriorityIndicator. expected %v, got %v", expected, actual)
	}

	if expected, actual := true, packet.AdaptationField.TransportPrivateDataFlag; actual != expected {
		t.Fatalf("unexpected TransportPrivateDataFlag. expected %v, got %v", expected, actual)
	}

	if expected, actual := []byte{0x02, 0x06, 0x28, 0xC8, 0x28, 0xEC, 0x10, 0xB3}, packet.AdaptationField.TransportPrivateData.Data; !reflect.DeepEqual(actual, expected) {
		t.Fatalf("unexpected TransportPrivateData. expected %v, got %v", expected, actual)
	}

	if expected, actual := 173, len(packet.Payload); actual != expected {
		t.Fatalf("payload was not the correct size. expected %v, got %v", expected, actual)
	}
}

func TestTransPcrField(t *testing.T) {
	reader := bitreader.NewReader(pcrDataPacketReader())
	packet, err := ts.NewPacket(reader)
	if err != nil {
		t.Fatal(err)
	}

	if expected, actual := packet.AdaptationFieldControl, ts.FieldThenPayload; actual != expected {
		t.Fatalf("unexpected AdaptationFieldControl. expected %v, got %v", actual, expected)
	}

	if packet.AdaptationField == nil {
		t.Fatal("exptected adaptation field to be set")
	}

	if expected, actual := true, packet.AdaptationField.PCRFlag; actual != expected {
		t.Fatalf("unexpected PCRFlag. expected %v, got %v", expected, actual)
	}

	if expected, actual := uint64(411841357261), packet.AdaptationField.PCR; actual != expected {
		t.Fatalf("unexpected PCR. expected %v, got %v", expected, actual)
	}

	if expected, actual := 176, len(packet.Payload); actual != expected {
		t.Fatalf("payload was not the correct size. expected %v, got %v", expected, actual)
	}

}
