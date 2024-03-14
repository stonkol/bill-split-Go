package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

type Person struct {
	name string
	lent float32 // is recommended to don't stored in decimals
}

type Item struct {
	//	name        string
	category string
	price    float32
	//	currency    string
	no_Split bool
}

type Currency struct {
	// full_name    string
	three_digits string
	one_digit    string
	total        float32
}

// currency's acronyms
//	const (
//		e string = "eur"
//	j string = "jpy"
//	p string = "gbp"
//	)

const default_currency string = "eur"

// Exchange rates matrix
var exchangeRates = map[string]map[string]float32{
	"USD": {"USD": 1.0, "EUR": 0.84, "GBP": 0.73, "JPY": 109.56},
	"EUR": {"USD": 1.19, "EUR": 1.0, "GBP": 0.87, "JPY": 130.50},
	"GBP": {"USD": 1.38, "EUR": 1.15, "GBP": 1.0, "JPY": 150.34},
	"JPY": {"USD": 0.0091, "EUR": 0.0077, "GBP": 0.0067, "JPY": 1.0},
}

// Exchange function
func currencyConverter(totalPrice float32, currencyIn, currencyOut string) (float32, error) {
	rate, ok := exchangeRates[currencyIn][currencyOut]
	if !ok {
		return 0, fmt.Errorf("exchange rate not found for %s to %s", currencyIn, currencyOut)
	}
	return totalPrice * rate, nil
}

func scanCalcItems() []Person {
	// Define slices to store persons and lent amounts
	person := make([]Person, 0)

	//////////////////////////////
	/////// RETRIEVE LIST ////////
	//////////////////////////////
	//////// OPEN MD
	file, err := os.Open("the-bills/test-2p-2c-2i.md")
	if err != nil {
		fmt.Println("Error:", err)
		return person
	}
	defer file.Close()

	p := -1

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		lineStartNumber := regexp.MustCompile(`^\s*\d+(\.\d+)?`) // ^\s*\d+(\.\d+)? matches lines that start with an optional whitespace, followed by one or more digits, which can be followed by a decimal point and one or more digits.

		///////////////////////////////////////////////////
		/////////////////// PERSON ///////////////////////
		/////////////////////////////////////////////////////
		if strings.HasPrefix(line, "# ") {
			p++

			// Extract the name from the line and create a new Person
			name := strings.TrimSpace(strings.TrimPrefix(line, "# "))
			person = append(person, Person{name: name})
			fmt.Println("\n p, name, line:", p, name, line)
			// person[p].name = name
			// fmt.Println("detected:", line)

			///////////////////////////////////////////////////
			//////////////////  CURRENCY /////////////////////
			//////////////////////////////////////////////////
		} else if strings.HasPrefix(line, "## ") {
			currency := make([]Currency, 0)

			three_digits := strings.TrimSpace(strings.TrimPrefix(line, "## "))
			// Convert currency1 to uppercase if it's not already
			if three_digits != strings.ToUpper(three_digits) {
				three_digits = strings.ToUpper(three_digits)
			}
			currency = append(currency, Currency{three_digits: three_digits})

			fmt.Println("currency:", three_digits)

			// currencyScanner := bufio.NewScanner(strings.NewReader(line))

			///////////////////////////////////////////////////
			//////////////////// ITEM /////////////////////////
			///////////////////////////////////////////////////
			itemScanner := bufio.NewScanner(strings.NewReader(line))
			for itemScanner.Scan() {
				fmt.Println("itemScanner loop ->")
				// currencyLine := currencyScanner.Text()
				itemLine := itemScanner.Text()

				///////////// NEW PERSON DETECTED: Break
				if strings.HasPrefix(itemLine, "# ") {

					// add that currency total
					// currencyConverter(currencyTotal, whichcurrency)
					person[p].lent = currency[p].total
					/////////////////////////
					break // TODO: will this break go out of if and for?

					///////////// NUMBER detected /////////////
				} else if lineStartNumber.MatchString(itemLine) { // TODO: not sure if the statement is corrrectly written

					// FindStringSubmatch returns a slice of strings containing the text of the leftmost match and the matches found by the capturing groups.
					match := lineStartNumber.FindStringSubmatch(itemLine)
					// match is the slice returned by FindStringSubmatch, and match[1] refers to the first captured group
					itemPrice := match[1]

					fmt.Println("item detected:", itemPrice, ". In the line:", itemLine)
					// total += itemPrice

					///////////// F detected //////////////////
				} else if strings.HasPrefix(line, "f") {
					// TODO:
					// f_prefix_trimmed := strings.TrimSpace(strings.TrimPrefix(line, "f"))
					// item.amount = trim the part after the number
					// item.amount = append(item, Item{})
					// item.amount *= 2
					// total += item amounts

					fmt.Println("Non splitted item found:", itemLine)
				}
			}
		}
	}

	return person
}

func finalPrint(p1_lent, p2_lent float32, p1_name, p2_name, currency string) {
	// calculate who owes more
	difference := (p1_lent - p2_lent) / 2

	// print who owes who and hpw much
	if difference > 0 {
		fmt.Println(p2_name, "owes", difference, currency)
	} else if difference < 0 {
		fmt.Println(p1_name, "owes", math.Abs(float64(difference)), currency)
	} else if difference == 0 {
		fmt.Println("You have lent the same amount.")
	} else {
		fmt.Println("Error: there is no difference data")
	}
}

func main() {
	// maybe make map to store multiple currencies in one variable?
	// var total, total_eur, total_gbp, total_jpy float32 = 0.0, 0.0, 0.0, 0.0

	person := scanCalcItems()

	// Print the names of the people
	fmt.Println("\n______________")
	for p := 0; p < len(person); p++ {
		fmt.Printf("Person %d: %s\n", p+1, person[p].name)
	}

	//////////////////////////////
	////////// TEST INPUT ////////
	//////////////////////////////
	// Declare an array of currencies
	currencies := [4]string{"EUR", "jpy", "usd", "GBP"}

	// Convert each currency to uppercase if it's not already
	for i, currency := range currencies {
		if currency != strings.ToUpper(currency) {
			currencies[i] = strings.ToUpper(currency)
		}
	}
	converted, err := currencyConverter(1, currencies[0], currencies[1])
	fmt.Print("Currency Check: 1 EUR = ", converted, "JPY, ")

	converted, err = currencyConverter(1, currencies[0], currencies[2])
	fmt.Print(converted, "USD, ")

	converted, err = currencyConverter(1, currencies[0], currencies[3])
	fmt.Println(converted, "GBP")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// person[0].lent = 30
	//  person[1].lent = 20
	// Create a slice of Item structs
	// item := []Item {}

	// for p := 0; p < 2; p++ {
	// 	if p == 0 {
	// 		item := []Item{
	// 			{"Bread", "f", 0, "eur", false},
	// 			{"Carrot", "f", 20, "eur", false},
	// 			{"Apple", "f", 10, "eur", false},
	// 		}
	// 		for i := 0; i < len(item); i++ {
	// 			total_eur += item[i].price
	// 		}
	// 		total = total_eur
	// 		person[p].lent = total_eur
	// 		fmt.Println("p1_lent: ", person[p].lent)
	// 	} else if p == 1 {
	// 		item := []Item{
	// 			{"Bread", "f", 10, "eur", false},
	// 			{"Carrot", "f", 20, "eur", false},
	// 			{"Apple", "f", 10, "eur", false},
	// 		}
	// 		for i := 0; i < len(item); i++ {
	// 			total_eur += item[i].price
	// 		}
	// 		total = total_eur
	// 		person[p].lent = total_eur
	// 		fmt.Println("p2_lent: ", person[p].lent)
	// 	}
	// }

	// item[0].name = "carrot"

	//	person[0].name = "Elephant"
	//	person[1].name = "Mamut"

	//	currency := default_currency

	//  if currency == "eur" {
	//    total_eur += item[i].price // add to eur
	//  } else if currency == "gbp" {
	//    total_gbp += item[i].price
	//  } else {
	//    fmt.Println("Error: No currency detected.")
	//  }

	// total += total_eur
	// total += total_gbp * pte
	// total += total_jpy * jte

	////////////////////////////////
	////////////////////////////////

	finalPrint(person[0].lent, person[1].lent, person[0].name, person[1].name, default_currency)
}

//var item []Item
//  item := Item{
//	  name:       "Default Item",
//	  price:      0.0,
//	  currency:   "EUR",
//	  category:   "other",
//	  full_amount: false,
//}

// // map imput list
// var p, i int16 := 0, 0
// if first letter is "f" {
//    item
// } else if first letter is num {
//    item
// }
// } else if "#" first letter {
//   person[p].name = "(string after "# ")"
//   i++
// } else if "##" first 2 letters{
//   item[i].currency = (the part after "## ")
//
// // calculate it
// for p := 0; p <=10; p++ {
//   if detect new "#" {
//
//     for item[i] {}
//         if  {}
//         total += item[i].price
//     p++
//   } else if "page end" {
//       p = 11
//   }
// }

// detect the first '#', it will be the first user_1

// add money on user_1_lent

// if detect another "#", then +1 to person[i] will start

// is adding money to each currency
//
//
// 			///////////////////////////////////////////////////
//////////////////  CURRENCY /////////////////////
//////////////////////////////////////////////////
// for {
// 	if strings.HasPrefix(line, "## ") {
// 		currency := make([]Currency, 0)
// 		three_digits := strings.TrimSpace(strings.TrimPrefix(line, "## "))
// 		currency = append(currency, Currency{three_digits: three_digits})
// 		fmt.Println("new currency:", line)
//
// 		// detect a blank line
// 	} else if len(strings.TrimSpace(line)) == 0 {
// 		// calculate total
//
// 		break
// 	}
// }
