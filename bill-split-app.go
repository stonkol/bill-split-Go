package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

//"io"
//"bufio"
//  "os"
//"strconv"

type Person struct {
	name string
	lent float32 // is recommended to don't stored in decimals
}

type Item struct {
	name        string
	category    string
	price       float32
	currency    string
	full_amount bool
}

type CurrencyFormat struct {
	full_name    string
	three_digits string
	one_digit    string
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
	// currency's acronyms
	const (
		e string = "eur"
		j string = "jpy"
		p string = "gbp"
	)

	// currency exchange
	const (
		pte float32 = 1.18
		etp float32 = 0.85 // pound to euro, euro to pound
		jte float32 = 0.01
	)

	// add all the same currencies
	var total, total_eur, total_gbp, total_jpy float32 = 0.0, 0.0, 0.0, 0.0

	//////////////////////////////
	/////// RETRIEVE LIST ////////
	//////////////////////////////
	const default_currency = "eur"

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

	//////////////////////////////
	/////////////////////////////
	//////// OPEN MD ////////////
	// Open the markdown file
	file, err := os.Open("the-bills/test-2p-2i.md")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Define slices to store persons and lent amounts
	person := make([]Person, 0)

	// Read each line of the file
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "# ") {
			// Extract the name from the line and create a new Person
			name := strings.TrimSpace(strings.TrimPrefix(line, "# "))
			person = append(person, Person{name: name})
			fmt.Println("detected:", person)
		}
	}

	// Print the names of the people
	for p := 0; p < len(person); p++ {
		fmt.Printf("Person %d: %s\n", p+1, person[p].name)
	}

	//////////////////////////////
	////////// TEST INPUT ////////
	//////////////////////////////
	// person[0].lent = 30
	//  person[1].lent = 20
	// Create a slice of Item structs
	//item := []Item {}

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

	currency := default_currency

	//  if currency == "eur" {
	//    total_eur += item[i].price // add to eur
	//  } else if currency == "gbp" {
	//    total_gbp += item[i].price
	//  } else {
	//    fmt.Println("Error: No currency detected.")
	//  }

	total += total_eur
	total += total_gbp * pte
	total += total_jpy * jte

	////////////////////////////////
	////////////////////////////////

	finalPrint(person[0].lent, person[1].lent, person[0].name, person[1].name, currency)
}
