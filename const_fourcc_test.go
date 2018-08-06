package fourcc

import (
	"testing"
)

func TestConstFOURCC(t *testing.T) {

	tests := []struct{
		Value    Const
		Expected Type
	}{
		{
			Value:    Const("PAL8"),
			Expected: FOURCC('P','A','L','8'),

		},



		{
			Value:    Const("GREY"),
			Expected: FOURCC('G','R','E','Y'),

		},
		{
			Value:    Const("Y06 "),
			Expected: FOURCC('Y','0','6',' '),

		},
		{
			Value:    Const("Y10 "),
			Expected: FOURCC('Y','1','0',' '),

		},
		{
			Value:    Const("Y12 "),
			Expected: FOURCC('Y','1','2',' '),

		},
		{
			Value:    Const("Y16 "),
			Expected: FOURCC('Y','1','6',' '),

		},

	}


	for testNumber, test := range tests {

		actual, err := test.Value.FOURCC()
		if nil != err {
			t.Errorf("For test #%d, did not expect an error, but actually got one: (%T) %q", testNumber, err, err)
			continue
		}

		if expected := test.Expected; expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}

func TestConstFOURCCError(t *testing.T) {

	tests := []struct{
		Value Const
	}{
		{
			Value: Const("APPLE"),
		},
		{
			Value: Const("BANANA"),
		},
		{
			Value: Const("CHERRY"),
		},



		{
			Value: Const(""),
		},
		{
			Value: Const("1"),
		},
		{
			Value: Const("12"),
		},
		{
			Value: Const("123"),
		},

		{
			Value: Const("12345"),
		},
		{
			Value: Const("123456"),
		},
		{
			Value: Const("1234567"),
		},
	}


	for testNumber, test := range tests {

		if _, err := test.Value.FOURCC(); nil == err {
			t.Errorf("For test #%d, did not expected an error, but did not actually got one: (%T) %q", testNumber, err, err)
			continue
		}
	}
}
