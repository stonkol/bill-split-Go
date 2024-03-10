package main

import (
	"fmt"
)

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
	var total_eur, total_gbp, total_jpy float32 = 0.0, 0.0, 0.0

	//////////////////////////////
	/////// RETRIEVE LIST ////////
	//////////////////////////////

	// default currency which will be converted
	const default_currency = "eur"

	var user_1_name, user_2_name string
	var user_1_lent, user_2_lent float32 = 0.0, 0.0

	// detect the first '#', it will be the first user_1

	// add money on user_1_lent

	// after detecting the next "#", will start

	// is adding money to

	////////// TEST INPUT //////////
	total_eur = total_eur + 3.0
	total_gbp = total_gbp + 5.5
	total_jpy = total_jpy + 0.0

	total_eur = total_gbp * pte
	total_eur = total_jpy * jte

	////////////////////////////////

	// calculate who owes more
	var final_lent float32
	// print who owes who and hpw much
	if user_1_lent > user_2_lent {
		final_lent = (user_1_lent - user_2_lent) / 2
		fmt.Println(user_2_name, " owes ", final_lent, default_currency)
	} else if user_2_lent > user_1_lent {
		final_lent = (user_2_lent - user_1_lent) / 2
		fmt.Println(user_1_name, " owes ", final_lent, default_currency)
	} else {
		fmt.Println("Error: Couldn't determined who owes")
	}
}
