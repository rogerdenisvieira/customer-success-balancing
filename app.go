package main

// this package is actually pointing to the local one

// this package is actually pointing to the local one
import (
	"fmt"
	csbalancing "rogerdenisvieira/customer-success-balancing/domain" // this package is actually pointing to the local one
)

func main() {

	customers := []csbalancing.Entity{
		{ID: 1, Score: 90},
		{ID: 2, Score: 20},
		{ID: 3, Score: 70},
		{ID: 4, Score: 40},
		{ID: 5, Score: 60},
		{ID: 6, Score: 10},
	}

	fmt.Println(customers)
	sortedCustomer := csbalancing.SortEntitiesByScoreDesc(customers)
	fmt.Println(sortedCustomer)

	//fmt.Println(csbalancing.SumEntitiesScore(customers))

}
