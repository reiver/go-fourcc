package fourcc

import (
	"testing"
)

func TestTypeString(t *testing.T) {

	tests := []struct{
		Value    Type
		Expected string
	}{
		{
			Value: Type{(uint32('a') | (uint32('b') << 8) | (uint32('c') << 16) | (uint32('d') << 24))},
			Expected:           "abcd",
		},
		{
			Value: Type{(uint32('A') | (uint32('B') << 8) | (uint32('C') << 16) | (uint32('D') << 24))},
			Expected:           "ABCD",
		},
		{
			Value: Type{(uint32('1') | (uint32('2') << 8) | (uint32('3') << 16) | (uint32('4') << 24))},
			Expected:           "1234",
		},



		{
			Value: Type{(uint32('P') | (uint32('A') << 8) | (uint32('L') << 16) | (uint32('8') << 24))},
			Expected:           "PAL8",
		},



		{
			Value: Type{(uint32('M') | (uint32('J') << 8) | (uint32('P') << 16) | (uint32('G') << 24))},
			Expected:           "MJPG",
		},
	}


	for testNumber, test := range tests {

		if expected, actual := test.Expected, test.Value.String(); expected != actual {
			t.Errorf("For test #%d, expected %q, but actually got %q.", testNumber, expected, actual)
			continue
		}
	}
}
