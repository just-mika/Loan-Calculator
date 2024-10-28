/*
********************
Last names: Amon, Jaramillo, Pua
Language: Go
Paradigm(s): Procedural, Imperative
********************
*/

package main

import "fmt"

func calculateLoan(loanAmt float64, intRate float64, yrTerm int) {
	fmt.Println("================================")

	intRate /= 100
	monthInt := intRate / 12
	monthTerm := float64(yrTerm) * 12

	totalInt := loanAmt * monthTerm * monthInt
	monthRepay := (loanAmt + totalInt) / monthTerm

	fmt.Printf("Loan Amount: PHP %.2f\n", loanAmt)
	fmt.Printf("Annual Interest Rate: %.1f%%\n", intRate*100)
	fmt.Printf("Loan Term: %d months\n", int(monthTerm))
	fmt.Printf("Monthly Repayment: PHP %.2f\n", monthRepay)
	fmt.Printf("Total Interest: PHP %.2f\n", totalInt)
}

func main() {
	var loanAmt, intRate float64
	var yrTerm int

	fmt.Print("Loan Amount: PHP ")
	fmt.Scanln(&loanAmt)
	fmt.Print("Annual Interest Rate (in %): ")
	fmt.Scanln(&intRate)
	fmt.Print("Loan Term (in years): ")
	fmt.Scanln(&yrTerm)

	calculateLoan(loanAmt, intRate, yrTerm)
}
