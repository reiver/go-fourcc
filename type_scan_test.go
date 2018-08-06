package fourcc

import (
	"testing"
)

func TestTypeScan(t *testing.T) {

	tests := []struct{
		Value    interface{}
		Expected Type
	}{
		{
			Value: "PAL8",
			Expected: FOURCC('P','A','L','8'),
		},



		{
			Value: []byte("PAL8"),
			Expected: FOURCC('P','A','L','8'),
		},



		{
			Value: uint32(0),
			Expected: FOURCC(0x00,0x00,0x00,0x00),
		},
		{
			Value: uint32(0xffffffff),
			Expected: FOURCC('ÿ','ÿ','ÿ','ÿ'),
		},



		{
			Value: uint64(0),
			Expected: FOURCC(0x00,0x00,0x00,0x00),
		},
		{
			Value: uint64(0xffffffff),
			Expected: FOURCC('ÿ','ÿ','ÿ','ÿ'),
		},
	}


	for testNumber, test := range tests {

		var actual Type

		if err := actual.Scan(test.Value); nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}

func TestTypeScanError(t *testing.T) {

	tests := []struct{
		Value interface{}
	}{
		{
			Value: false,
		},
		{
			Value: true,
		},



		{
			Value: int8(5),
		},
		{
			Value: int16(5),
		},
		{
			Value: int32(5),
		},



		{
			Value: uint8(5),
		},
		{
			Value: uint16(5),
		},



		{
			Value: "apple banana cherry",
		},



		{
			Value: []byte("apple banana cherry"),
		},
	}


	for testNumber, test := range tests {

		var actual Type

		if err := actual.Scan(test.Value); nil == err {
			t.Errorf("For test #%d, expected an error, but did not actually got one: (%T) %q", testNumber, err, err)
			continue
		}
	}
}
