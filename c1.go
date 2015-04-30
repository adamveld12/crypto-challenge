package challenge

import (
	"fmt"
)

var um byte = 0x3F
var lm byte = 0x03

func decimalToBase64(dec byte) string {
	out := byte(0)
	if dec >= 0 && dec <= 25 {
		out = (dec + 'A')
	} else if dec >= 26 && dec <= 51 {
		out = (dec - 26 + 'a')
	} else if dec >= 52 && dec <= 61 {
		out = (dec - 52 + '0')
	} else if dec == 62 {
		out = '+'
	} else if dec == 63 {
		out = '/'
	} else {
		panic(fmt.Sprint("out of range ", dec))
	}
	return string(out)
}

func HexToBase64(hex string) string {
	data := ""

	// for every 6 runes, turn them into 3 decimal values
	for i := 0; i < len(hex); i += 6 {
		decVals := fill(hex, i)

		section := ""

		// then for every 6 bits, turn those into 4 base64 values
		// first 6 bits from [0]
		section += decimalToBase64(((um << 2) & decVals[0]) >> 2)

		// 2 bits from end of [0], 4 bits from start of [1]
		section += decimalToBase64(((lm & decVals[0]) << 4) | ((0xF0 & decVals[1]) >> 4))

		// 4 bits from end of [1], 2 bits from start of [2]
		section += decimalToBase64(((0x0F & decVals[1]) << 2) | ((0xC0 & decVals[2]) >> 6))

		// 6 bits from end of [3]
		section += decimalToBase64((0x3F & decVals[2]))

		data += section
	}

	return string(data)
}

func fill(hex string, index int) []byte {
	hexValue := make([]byte, 0)
	for idx := 0; idx < 6 && idx+index < len(hex); idx++ {
		hexValue = append(hexValue, hex[idx+index])
	}

	for idx := len(hexValue); idx < 6; idx++ {
		hexValue = append(hexValue, 255)
	}

	output := make([]byte, 0)
	for idx := 0; idx < 6; idx += 2 {
		decimal0, decimal1 := hexRuneToDecimal(hexValue[idx])<<4, hexRuneToDecimal(hexValue[idx+1])
		output = append(output, decimal0|decimal1)
	}

	return output
}

func hexRuneToDecimal(hex byte) byte {
	if hex >= '0' && hex <= '9' {
		return byte(hex - '0')
	} else if hex >= 'a' && hex <= 'f' {
		return byte(hex-'a') + 10
	} else if hex >= 'A' && hex <= 'F' {
		return byte(hex-'A') + 10
	}
	return 0
}
