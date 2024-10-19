package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cleanInput(input string) string {
	var cleanedinput string
	if strings.Contains(strings.ToLower(input), "rp") {
		cleanedinput = strings.Split(strings.TrimSpace(input), " ")[1]
		cleanedinput = strings.Replace(cleanedinput, ".", "", -1)
	} else {
		cleanedinput = strings.Replace(strings.TrimSpace(input), ".", "", -1)
	}
	return cleanedinput
}

func formatCurrency(c int) string{
	return fmt.Sprintf("%.3f", float64(c)/1000)
}

func Kasir(totalOrderStr, paymentStr string){
	totalOrder, _ := strconv.Atoi(cleanInput(totalOrderStr))
	payment, _ := strconv.Atoi(cleanInput(paymentStr))

	if payment < totalOrder {
		fmt.Println("False, kurang bayar")
		return 
	}
	
	result := ((payment - totalOrder)/100) * 100
	fmt.Printf("Kembalian yang harus diberikan kasir: %s\n", formatCurrency(result))
	fmt.Println("Pecahan uang:")
	
	pecahan := []int{100000, 50000, 20000, 10000, 5000, 2000, 200, 100}
	for _, v := range pecahan{
		hasilBagi := result / v
		if hasilBagi <= 0{
			continue
		}
		satuan := "lembar"
		formattedPecahan := formatCurrency(v)


		if v == 100 || v == 200 {
			satuan = "koin"
			formattedPecahan = strconv.Itoa(v)
		}
		result = result - (v * hasilBagi)
		fmt.Printf("%d %s %s\n", hasilBagi, satuan, formattedPecahan)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Total Belanja seorang customer:")
	totalOrderInput, _ := reader.ReadString('\n')

	fmt.Println("Pembeli membayar:")
	paymentInput, _ := reader.ReadString('\n')

	Kasir(totalOrderInput, paymentInput)

}
