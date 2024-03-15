package main

// regexp: for detecting if the line started as a number
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Person struct {
	name string
	lent float32 // is recommended not to store in decimals
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

//       TODO: currency's acronyms
//const (
//	e string = "eur"
//	j string = "jpy"
//  p string = "gbp"
//)

const default_currency string = "EUR"

var prefCurrency string = default_currency

var exchangedPrice float32 = 0.0

// var err string

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
	exchangedPrice = totalPrice * rate
	// fmt.Printf("exchanged price is: %.2f \n", exchangedPrice) //check

	return exchangedPrice, nil
}

func finalPrint(p1_lent, p2_lent float32, p1_name, p2_name, currency string) {
	// calculate who owes more
	difference := (p1_lent - p2_lent) / 2

	// print who owes who and how much
	fmt.Print("\n  ")
	if difference > 0 {
		fmt.Printf("%s owes %s %.2f %s  \n", p2_name, p1_name, difference, currency)
	} else if difference < 0 {
		fmt.Printf("%s owes %s %.2f %s  \n ", p1_name, p2_name, math.Abs(float64(difference)), currency)
	} else if difference == 0 {
		fmt.Println("You have lent the same amount.")
	} else {
		fmt.Println("Error: there is no difference data")
	}
	fmt.Println("")
}

func scanCalcItems(filePath string) []Person {
	// Define slices to store persons and lent amounts
	person := make([]Person, 0)

	//////////////////////////////
	/////// RETRIEVE LIST ////////
	//////////////////////////////
	//////// OPEN MD

	// file, err := os.Open("the-bills/test-2p-2c-2i.md")
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error:", err)
		return person
	}
	defer file.Close()

	// ^\s*\d+(\.\d+)? matches lines that start with an optional whitespace,
	// followed by one or more digits, which can be followed by a decimal point
	// and one or more digits.
	lineStartNumber := regexp.MustCompile(`^\s*\d+(\.\d+)?`)

	// INIT person and currency
	p, c := -1, -1
	currency := make([]Currency, 0)
	var three_digits string
	var totalLend, currencyTotal, currencyTotalExchanged float32 = 0.0, 0.0, 0.0
	personChangedTrue := false

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()

		///////////////////////////////////////////////////
		/////////////////// PERSON ///////////////////////
		/////////////////////////////////////////////////////
		if strings.HasPrefix(line, "# ") {

			p++

			name := strings.TrimSpace(strings.TrimPrefix(line, "# "))
			person = append(person, Person{name: name})

			/////// PERSON 1 ENDS, PERSON 2 GOING TO START ////////
			if p > 0 {

				// process last currency of the previous person
				personChangedTrue = true

				if three_digits != prefCurrency {
					currencyTotalExchanged, err = currencyConverter(currencyTotal, three_digits, prefCurrency)
					fmt.Printf("   %.2f %s -> %.2f %s.\n", currencyTotal, three_digits, currencyTotalExchanged, prefCurrency)
				} else {
					currencyTotalExchanged = currencyTotal
					// fmt.Println("Is already prefCurrency:", prefCurrency)
				}
				currency = append(currency, Currency{total: currencyTotalExchanged})
				totalLend += currencyTotalExchanged
				currencyTotal, currencyTotalExchanged = 0.0, 0.0

				///////////

				person[p-1].lent = totalLend
				totalLend = 0.0
				// fmt.Println("\n ->", person[p-1].name, "lent:", person[p-1].lent)
				fmt.Printf("\n -> %s lent: %.2f \n", person[p-1].name, person[p-1].lent)

				//////// PERSON 1 start//////////
			} else if p == 0 {
				//////// ERROR /////////
			} else {
				fmt.Println("Error: else of Person.")
			}
			//			fmt.Println("\n P", p, "name:", name, "person:", person[p].name)
			fmt.Printf("\n---------- %s ----------\n", person[p].name)

			///////////////////////////////////////////////////
			//////////////////  CURRENCY /////////////////////
			//////////////////////////////////////////////////
		} else if strings.HasPrefix(line, "## ") {

			c++

			// convert the previous currency with his items total price
			// check if three_digits have already stored something
			if len(three_digits) > 0 && !personChangedTrue {
				if three_digits != prefCurrency {
					currencyTotalExchanged, err = currencyConverter(currencyTotal, three_digits, prefCurrency)
					fmt.Printf("    %.2f %s -> %.2f %s.\n", currencyTotal, three_digits, currencyTotalExchanged, prefCurrency)
				} else {
					currencyTotalExchanged = currencyTotal
					// fmt.Println("Is already prefCurrency:", prefCurrency)
				}
				currency = append(currency, Currency{total: currencyTotalExchanged})
				totalLend += currencyTotalExchanged
				fmt.Printf("   ––––––––––––––\n   = %.2f %s \n", currencyTotalExchanged, three_digits)
				currencyTotal, currencyTotalExchanged = 0.0, 0.0
			}

			if personChangedTrue {
				personChangedTrue = false
			}

			////////////// RETRIEVE
			three_digits = strings.TrimSpace(strings.TrimPrefix(line, "## "))

			// Convert currency to uppercase, if it's not already
			if three_digits != strings.ToUpper(three_digits) {
				three_digits = strings.ToUpper(three_digits)
			}
			// currency = append(currency, Currency{three_digits: three_digits})
			// currency[c].three_digits = three_digits

			fmt.Println("\n +", three_digits)

			// currencyScanner := bufio.NewScanner(strings.NewReader(line))

			///////////////////////////////////////////////////
			//////////////////// ITEM /////////////////////////
			///////////////////////////////////////////////////
			//////////////// NUMBER detected //////////////////
		} else if lineStartNumber.MatchString(line) {

			// FindStringSubmatch returns a slice of strings containing the text
			// of the leftmost match and the matches found by the capturing groups.
			match := lineStartNumber.FindStringSubmatch(line)
			// match is the slice returned by FindStringSubmatch,
			// and match[1] refers to the first captured group
			itemPrice := match[0]
			itemDescription := strings.TrimSpace(strings.TrimPrefix(line, itemPrice))

			// Convert the string to a float32
			itemPriceFloat, err := strconv.ParseFloat(itemPrice, 32)
			if err != nil {
				// Handle error if the string is not a valid number1
				panic(err)
			}

			// fmt.Println("match:", match, "match[0], match[1]:", match[0], match[1]) // check
			fmt.Printf("   %.2f - %s\n", itemPriceFloat, itemDescription)

			// totalLend += float32(itemPriceFloat)
			currencyTotal += float32(itemPriceFloat)

			//////////////////////////////////////////
			///////////// F detected //////////////////
			/////////////////////////////////////////
		} else if strings.HasPrefix(line, "f") {
			// TODO:
			// f_prefix_trimmed := strings.TrimSpace(strings.TrimPrefix(line, "f"))
			// item.amount = trim the part after the number
			// item.amount = append(item, Item{})
			// item.amount *= 2
			// total += item amounts

			fmt.Println("Non splitted item found:", line)

			///////////////////////////////////////////
			//////////////// BLANK LINE ///////////////
			///////////////////////////////////////////
		} else if len(strings.TrimSpace(line)) == 0 {
			continue

			/////////////////////////////////////////
			/////////////// ERROR ///////////////////
			/////////////////////////////////////////
		} else {
			fmt.Println("ERROR: else of itemScanner")
		}
	}

	/////////////////////////////////////////
	////////////// LAST LINE ////////////////
	/////////////////////////////////////////
	if !scanner.Scan() {

		// process last currency of the last person
		if three_digits != prefCurrency {
			currencyTotalExchanged, err = currencyConverter(currencyTotal, three_digits, prefCurrency)
			fmt.Printf("   %.2f %s -> %.2f %s.\n", currencyTotal, three_digits, currencyTotalExchanged, prefCurrency)
		} else {
			currencyTotalExchanged = currencyTotal
			fmt.Println("Is already prefCurrency:", prefCurrency)
		}
		// currency = append(currency, Currency{total: currencyTotalExchanged})
		totalLend += currencyTotalExchanged

		// calculate last person total lent
		person[p].lent = totalLend
		fmt.Printf("\n -> %s lent: %.2f \n", person[p].name, person[p].lent)
	}

	// finalPrint(person[0].lent, person[1].lent, person[0].name, person[1].name, prefCurrency)

	return person
}

func testings() {
	//////////////////////////////
	////////// TEST INPUT ////////
	//////////////////////////////
	fmt.Println("\n___ TESTINGS ___")

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
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run bill-split-app.go /path/to/file.md")
		os.Exit(1)
	}

	filepath := os.Args[1]
	person := scanCalcItems(filepath)

	//////
	//person := scanCalcItems()
	// testings()

	fmt.Println("\n\n=============== BILL ===============\n")

	for p := 0; p < len(person); p++ {
		fmt.Printf("  P%d: %s lend: %.2f %s\n", p+1, person[p].name, person[p].lent, prefCurrency)
	}
	finalPrint(person[0].lent, person[1].lent, person[0].name, person[1].name, prefCurrency)

	fmt.Println("====================================\n\n")
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
//
//
///////////////////////////////////////////////////
//////////////////// ITEM /////////////////////////
///////////////////////////////////////////////////
//itemScanner := bufio.NewScanner(strings.NewReader(line))
//scanner := bufio.NewScanner(file)

// for scanner.Scan() {
// fmt.Println("itemScanner loop ->") // check
// currencyLine := currencyScanner.Text()
//	itemLine := scanner.Text()

//	fmt.Println("itemLine:", line)

///////////// NEW PERSON DETECTED: Break
//if strings.HasPrefix(itemLine, "# ") {

//	fmt.Println("new Person detected.")
// add that currency total
// currencyConverter(currencyTotal, whichcurrency)
//person[p].lent = currency[p].total
/////////////////////////
