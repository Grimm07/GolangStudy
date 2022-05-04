package GroceryRegisterSimulator

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

var products = map[string]float64{
	"Pizza Rolls":       4.50,
	"Double Stuf Oreos": 4.00,
	"Classic Lays":      3.00,
	"Milk":              3.00,
	"Waffles":           2.75,
	"Grape Jelly":       2.75,
	"Peanut Butter":     2.50,
	"Bread":             2.00,
	"Baby Carrots":      2.00,
	"Mac & Cheese":      1.00,
}

type Order struct {
	SubTotal float64
	Total    float64
	Items    []string
}

// Shows user products available to purchase (each key and corresponding value in the products map).
func ShowProducts() {
	fmt.Println("\n--------- PRODUCTS ---------")
	for k, v := range products {
		//formatting spacing
		space := "      "
		if len(k) < 8 {
			space = "\t\t       "
		} else if len(k) < 14 {
			space = "\t       "
		}
		fmt.Print(k, space)
		fmt.Printf("$%.2f\n", v)
	}
	fmt.Println()
}

// Prints a receipt which shows the user's entered products and their quantities, and the subtotal and total for the order.
func (o *Order) PrintReceipt() {
	fmt.Println("\n--------- RECEIPT ---------")
	for i := 0; i < len(o.Items); i++ {
		fmt.Println(o.Items[i])
	}
	fmt.Printf("\nSubtotal: $%.2f\n", o.SubTotal)
	fmt.Printf("Total:    $%.2f\n\n", o.Total)
}

// Calculates total from subtotal using an arbitrary tax.
func (o *Order) CalculateTotal() {
	tax := 0.0625
	total := o.SubTotal + o.SubTotal*float64(tax)
	o.Total = math.Floor(total*100) / 100
}

// User is asked to enter the product they want and its quantity until they enter "STOP".
// Each entry is added to Items in Order (struct) - the SubTotal is updated for each entry.
func (o *Order) BuildReceipt() {
	EntryStr := ""
	StopStr := "STOP"
	InvalidStr := "Invalid Input!"
	scanner := bufio.NewScanner(os.Stdin)

	// entry should be in format of 'product [quantity]' (there can be multiple spaces between)
	entry, _ := regexp.Compile("^[A-Za-z&\\ ]+\\ \\[[1-9]\\]$")
	// the product name consists of at least one: A-Z, a-z, &, and/or space
	product, _ := regexp.Compile("\\b[A-Za-z&\\ ]+\\b")
	// a number 1-9
	quantity, _ := regexp.Compile("\\b[1-9]\\b")
	for EntryStr != StopStr {
		fmt.Print("Enter 'product [quantity]' to add an item or 'STOP' when you're done: ")
		if scanner.Scan() {
			EntryStr = scanner.Text()
		}
		EntryCheck := entry.MatchString(EntryStr)
		// if valid entry according to regular expression
		if EntryCheck {
			ProductString := product.FindString(EntryStr)
			// if ProductString is in products map, InProducts is set to true, and the value from
			// products[ProductString] is assigned to price
			if price, InProducts := products[ProductString]; InProducts {
				QuantityNum, _ := strconv.Atoi(quantity.FindString(EntryStr))
				// What the user enters is appended to Items in Order (struct)
				o.Items = append(o.Items, EntryStr)
				// Price of purchase is added to SubTotal in Order
				o.SubTotal += price * float64(QuantityNum)
			} else {
				fmt.Println(InvalidStr)
			} // Notify user that what they entered is invalid
		} else if EntryStr != StopStr {
			fmt.Println(InvalidStr)
		}
	}
}

func Register() {
	ShowProducts()
	fmt.Println("Product names entered must have same case as shown, and the quantity entered must be 1-9.")
	o := &Order{}
	o.BuildReceipt()
	o.CalculateTotal()
	o.PrintReceipt()
}
