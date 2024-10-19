package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func cleanInput(input string) string {
	return strings.TrimSpace(input)
}

func CheckLeaveAvaibility(totalCutiBersama int, tglJoinStr string, tglCutiStr string, durasiCuti int) (bool, string) {
	const totalCutiKantor = 14
	totalCutiPribadi := totalCutiKantor - totalCutiBersama

	tglJoin, err := time.Parse("2006-01-02", cleanInput(tglJoinStr))
	if err != nil {
		return false, "Invalid date format"
	}

	tglCuti, err := time.Parse("2006-01-02", cleanInput(tglCutiStr))
	if err != nil {
		return false, "Invalid date format"
	}

	diffDate := tglCuti.Sub(tglJoin)

	days := int(diffDate.Hours() / 24)
	if days < 180 {
		return false, "Alasan: Karena belum 180 hari sejak tanggal join karyawan"
	}

	avaibilityDays := tglJoin.AddDate(0, 0, 180)
	lastMonth := time.Date(tglJoin.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	leaveDays := int(lastMonth.Sub(avaibilityDays).Hours()/24)
	quotaLeave := int(math.Floor((float64(leaveDays) / float64(365)) * float64(totalCutiPribadi)))

	if durasiCuti > quotaLeave{
		return false, fmt.Sprintf("Alasan: Karena hanya boleh mengambil %d hari cuti", quotaLeave)
	}

	return true, ""
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Jumlah Cuti Bersama:")
	jumlahCutiBersama, _ := reader.ReadString('\n')

	fmt.Println("Tanggal Join Karyawan: ")
	tglJoin, _ := reader.ReadString('\n')

	fmt.Println("Tanggal Rencana Cuti: ")
	tglCuti, _ := reader.ReadString('\n')

	fmt.Println("Durasi cuti (1): ")
	durasiCutiInput, _ := reader.ReadString('\n')

	totalCutiBersama, _ := strconv.Atoi(cleanInput(jumlahCutiBersama))
	durasiCuti, _ := strconv.Atoi(cleanInput(durasiCutiInput))

	result, message := CheckLeaveAvaibility(totalCutiBersama, tglJoin, tglCuti, durasiCuti)
	fmt.Println(result)
	fmt.Println(message)

}
