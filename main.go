package main

import (
	"fmt"
	"math/big"

	"github.com/leekchan/accounting"
)

/**
mar = 6188.08
feb = 5822.46
jan = 6030.54
dec = 5842.54
nov = 5933.12
oct = 5665.84
morland = 1500 * 9
ghost_writer = 3938.88
may = 3384.61
june = 6666.66
july = 6666.66 - 2787.60
august = 6666.66


total = 16,839.38
tax = 950.4 + 566.2 + 1123.6 + 257.6
total - tax

*/

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}

	salary := big.NewFloat(70633)
	allowanceTax := getTaxOnAllowance(salary)
	basicTax := getBasicTaxAmount(salary)
	higherTax := getHigherTaxAmount(salary)
	additionalTax := getAdditionalTaxAmount(salary)

	totalTax := new(big.Float).Add(allowanceTax, basicTax)

	totalTax.Add(totalTax, higherTax)
	totalTax.Add(totalTax, additionalTax)

	fmt.Println(allowanceTax)
	fmt.Println(basicTax)
	fmt.Println(higherTax)
	fmt.Println(additionalTax)

	fmt.Println("Tax is: ", ac.FormatMoneyBigFloat(totalTax))
	fmt.Println("Percentage is: ", new(big.Float).Quo(totalTax, salary))
	fmt.Println("Percentage is: ", totalTax.Quo(totalTax, salary).Mul(totalTax, big.NewFloat(100)))
}

func getTaxOnAllowance(salary *big.Float) *big.Float {
	lowerLimit := big.NewFloat(100_000)

	result := salary.Cmp(lowerLimit)

	if result < 1 {
		return big.NewFloat(0)
	}

	upperLimit := big.NewFloat(125_140)
	allowance := big.NewFloat(12_570)
	taxRate := big.NewFloat(0.2)

	result = salary.Cmp(upperLimit)

	if result > 0 {
		return new(big.Float).Mul(allowance, taxRate)
	}

	taxable := new(big.Float).Sub(upperLimit, salary)
	taxable.Quo(taxable, big.NewFloat(2))
	taxable.Sub(allowance, taxable)
	taxable.Mul(taxable, taxRate)

	return taxable
}

func getBasicTaxAmount(salary *big.Float) *big.Float {
	lowerLimit := big.NewFloat(12_570)
	upperLimit := big.NewFloat(50_270)
	taxRate := big.NewFloat(0.2)

	result := salary.Cmp(upperLimit)
	tax := big.NewFloat(0)

	if result > -1 {
		tax.Sub(upperLimit, lowerLimit)
		tax.Mul(tax, taxRate)
		return tax
	}

	result = salary.Cmp(lowerLimit)

	if result < 1 {
		return tax
	}

	tax.Sub(salary, lowerLimit)
	tax.Mul(tax, taxRate)
	return tax
}

func getHigherTaxAmount(salary *big.Float) *big.Float {
	lowerLimit := big.NewFloat(50_270)
	upperLimit := big.NewFloat(125_140)
	taxRate := big.NewFloat(0.4)

	result := salary.Cmp(upperLimit)
	tax := big.NewFloat(0)

	if result > -1 {
		tax.Sub(upperLimit, lowerLimit)
		tax.Mul(tax, taxRate)
		return tax
	}

	result = salary.Cmp(lowerLimit)

	if result < 1 {
		return tax
	}

	tax.Sub(salary, lowerLimit)
	tax.Mul(tax, taxRate)
	return tax
}

func getAdditionalTaxAmount(salary *big.Float) *big.Float {
	lowerLimit := big.NewFloat(125_140)
	taxRate := big.NewFloat(0.45)

	result := salary.Cmp(lowerLimit)
	tax := big.NewFloat(0)

	if result < 1 {
		return tax
	}

	tax.Sub(salary, lowerLimit)
	tax.Mul(tax, taxRate)
	return tax
}
