package challenge

import (
	"testing"
)

func TestHexToBase64(t *testing.T) {
	inputs := []string{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "DEAD", "DEADBABE"}
	outputs := []string{"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t", "3q0=", "3q26vg=="}

	for idx, input := range inputs {
		result := HexToBase64(input)
		if result != outputs[idx] {
			t.Error("For", input, ":", outputs[idx], "!=", result)
		}
	}
}