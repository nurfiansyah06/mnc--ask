package main

import (
	"fmt"
	"math"
)

func hitungKembalian(total, spendingMoney int) interface{} {
	if spendingMoney < total {
		return false
	}

	changeMoney := spendingMoney - total
	roundChange := int(math.Floor(float64(changeMoney)/100) * 100)

	fragment := []int{100000, 50000, 20000, 10000, 5000, 2000, 1000, 500, 200, 100}

	numbOfFragment := make(map[int]int)

	remainingChange := roundChange
	for _, p := range fragment {
		if remainingChange >= p {
			numbOfFragment[p] = remainingChange / p
			remainingChange = remainingChange % p
		}
	}

	return struct {
		Changes         int
		NumbOfFragments map[int]int
	}{
		Changes:         roundChange,
		NumbOfFragments: numbOfFragment,
	}
}

func main() {
	var total, paidMoney int

	fmt.Print("Masukkan total belanja: Rp ")
	fmt.Scan(&total)
	fmt.Print("Masukkan uang yang dibayarkan: Rp ")
	fmt.Scan(&paidMoney)

	result := hitungKembalian(total, paidMoney)

	switch res := result.(type) {
	case bool:
		fmt.Println("False, kurang bayar")
	case struct {
		Changes         int
		NumbOfFragments map[int]int
	}:
		fmt.Printf("Kembalian yang harus diberikan kasir: %d\n", res.Changes)
		fmt.Println("Pecahan uang:")
		for fragment, amount := range res.NumbOfFragments {
			if fragment >= 1000 {
				fmt.Printf("%d lembar %d\n", amount, fragment)
			} else {
				fmt.Printf("%d koin %d\n", amount, fragment)
			}
		}
	}
}
