package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test04(t *testing.T) {
	scenarios := []struct {
		name  string
		input struct {
			jumlahCutiBersama int
			tglJoin           string
			tglCuti           string
			durasiCuti        int
		}
		output1 bool
		output2 string
	}{
		{
			name: "1",
			input: struct {
				jumlahCutiBersama int
				tglJoin           string
				tglCuti           string
				durasiCuti        int
			}{
				jumlahCutiBersama: 7,
				tglJoin:           "2021-05-01",
				tglCuti:           "2021-07-05",
				durasiCuti:        1,
			},
			output1: false,
			output2: "Alasan: Karena belum 180 hari sejak tanggal join karyawan",
		},

		{
			name: "2",
			input: struct {
				jumlahCutiBersama int
				tglJoin           string
				tglCuti           string
				durasiCuti        int
			}{
				jumlahCutiBersama: 7,
				tglJoin:           "2021-05-01",
				tglCuti:           "2021-11-05",
				durasiCuti:        3,
			},
			output1: false,
			output2: "Alasan: Karena hanya boleh mengambil 1 hari cuti",
		},
		{
			name: "3",
			input: struct {
				jumlahCutiBersama int
				tglJoin           string
				tglCuti           string
				durasiCuti        int
			}{
				jumlahCutiBersama: 7,
				tglJoin:           "2021-01-05",
				tglCuti:           "2021-12-18",
				durasiCuti:        1,
			},
			output1: true,
			output2: "",
		},
		{
			name: "4",
			input: struct {
				jumlahCutiBersama int
				tglJoin           string
				tglCuti           string
				durasiCuti        int
			}{
				jumlahCutiBersama: 7,
				tglJoin:           "2021-01-05",
				tglCuti:           "2021-12-18",
				durasiCuti:        3,
			},
			output1: true,
			output2: "",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			result, message := CheckLeaveAvaibility(scenario.input.jumlahCutiBersama, scenario.input.tglJoin, scenario.input.tglCuti, scenario.input.durasiCuti)
			assert.Equal(t, scenario.output1, result)
			assert.Equal(t, scenario.output2, message)
		})
	}

}
