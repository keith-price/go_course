//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

// these constants are better than just using true/false as they are less ambiguous
const (
	active   = true
	inactive = false
)

type item struct {
	name     string
	tagState bool
}

func changeTagState(item *item) {
	item.tagState = !item.tagState
}

func deactivateAll(items []item) {
	for i := range items {
		items[i].tagState = inactive
	}

}

func main() {

	items := make([]item, 4)

	items[0] = item{"Teeth", active}
	items[1] = item{"Eggs", active}
	items[2] = item{"Yeast", active}
	items[3] = item{"Ovals", active}

	fmt.Println(items)

	changeTagState(&items[0])

	fmt.Println(items)

	deactivateAll(items)

	fmt.Println(items)

}
