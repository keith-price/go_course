//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type item struct {
	price int
	name  string
}

func printListInfo(list [4]item) {
	var cost, numItems int
	for i := 0; i < len(list); i++ {
		item := list[i]
		cost += item.price

		if item.name != "" {
			numItems += 1
		}
	}

	fmt.Println(list[numItems-1])
	fmt.Println(len(list))
	fmt.Println(cost)

}

func main() {

	shoppingList := [4]item{
		{29, "Eels"},
		{21, "Small Eels"},
		{2, "Eel Surprise"},
	}

	printListInfo(shoppingList)

	shoppingList[3] = item{333, "Boil in the bag arses"}

	printListInfo(shoppingList)

}
