package converters

import (
	"api/converters"
	"testing"
)

func TestAsciiToBinary(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"dilara", "01100100 01101001 01101100 01100001 01110010 01100001 "},
	}
	for _, tc := range testcases {
		res := converters.AsciiToBinary(tc.in)
		if res != tc.want {
			t.Errorf("sonuç: %s, olması gereken: %s", res, tc.want)
		}
	}
}
