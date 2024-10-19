package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T){
	scenarios := []struct{
		name string
		input struct{
			n int
			strings string
		}
		output string
	}{
		{
			name: "1",
			input: struct{n int; strings string}{
				n: 4,
				strings: "abcd acbd aabd acbd",
			},
			output: "2 4 ",
		},
		{
			name: "2",
			input: struct{n int; strings string}{
				n: 5,
				strings: "pisang goreng enak sekali rasanya",
			},
			output: "false",
		},
		{
			name: "3",
			input: struct{n int; strings string}{
				n: 11,
				strings: "Satu Sate Tujuh Tusuk Tujuh Sate Bonus Tiga Puluh Tujuh Tusuk",
			},
			output: "2 6 3 5 10 4 11 ",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result := CountPairString(scenario.input.n, scenario.input.strings)
			assert.Equal(t, scenario.output, result)
		})
	}


}