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

func TestAsciiToOctal(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"dilara", "144 151 154 141 162 141 "},
	}
	for _, tc := range testcases {
		res := converters.AsciiToOctal(tc.in)
		if res != tc.want {
			t.Errorf("sonuç: %s, olması gereken: %s", res, tc.want)
		}
	}
}

func TestBinaryToAscii(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"01100100 01101001 01101100 01100001 01110010 01100001", "d i l a r a "},
	}
	for _, tc := range testcases {
		res := converters.BinaryToAscii(tc.in)
		if res != tc.want {
			t.Errorf("sonuç: %s, olması gereken: %s", res, tc.want)
		}
	}
}
func TestBinaryToOctal(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"01100100 01101001 01101100 01100001 01110010 01100001", "144 151 154 141 162 141 "},
	}
	for _, tc := range testcases {
		res := converters.BinaryToOctal(tc.in)
		if res != tc.want {
			t.Errorf("sonuç: %s, olması gereken: %s", res, tc.want)
		}
	}
}

func TestOctaltoAscii(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"144 151 154 141 162 141", "d i l a r a "},
	}
	for _, tc := range testcases {
		res := converters.OctalToAscii(tc.in)
		if res != tc.want {
			t.Errorf("sonuç: %s, olması gereken: %s", res, tc.want)
		}
	}
}

func TestOctaltoBinary(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"144 151 154 141 162 141", "01100100 01101001 01101100 01100001 01110010 01100001 "},
	}
	for _, tc := range testcases {
		res := converters.OctalToBinary(tc.in)
		if res != tc.want {
			t.Errorf("sonuç: %s, olması gereken: %s", res, tc.want)
		}
	}
}
