package csbalancing

import (
	"fmt"
	"sort"
)

// Entity ...
type Entity struct {
	ID    int
	Score int
}

/*
Todos os CSs têm níveis diferentes
Não há limite de clientes por CS
Clientes podem ficar sem serem atendidos
Clientes podem ter o mesmo tamanho
0 < n < 1.000
0 < m < 1.000.000
0 < id do cs < 1.000
0 < id do cliente < 1.000.000
0 < nível do cs < 10.000
0 < tamanho do cliente < 100.000
Valor máximo de t = n/2 arredondado para baixo
*/

// CustomerSuccessBalancing retrieves the customer success' ID with most customers
func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	// Write your solution here

	// fmt.Println(customerSuccess)
	// fmt.Println(customers)

	var customersByCustomerSuccess = make(map[int]int) // [CustomerSuccessID]CustomerQuantity
	customerSuccessByScore := SortEntitiesByScoreDesc(customerSuccess)
	customersByScore := SortEntitiesByScoreDesc(customers)

	availableCustomerSuccessByScore := FindAvailableCustomerSuccess(customerSuccessByScore, customerSuccessAway)

	for _, customer := range customersByScore {
		for availabeCustomerSuccessIndex, availabeCustomerSuccess := range availableCustomerSuccessByScore {

			if customer.Score <= availabeCustomerSuccess.Score && availabeCustomerSuccessIndex == len(availableCustomerSuccessByScore)-1 {
				customersByCustomerSuccess[availabeCustomerSuccess.ID]++
			}

			if customer.Score <= availabeCustomerSuccess.Score && customer.Score > availableCustomerSuccessByScore[availabeCustomerSuccessIndex+1].Score {
				customersByCustomerSuccess[availabeCustomerSuccess.ID]++
			}
		}
	}

	fmt.Println(customersByCustomerSuccess)

	return FindBusiestCustomerSuccess(customersByCustomerSuccess)
}

// O(m*n)
// FindBusiestCustomerSuccess finds the ID of the CustomerSuccess with most customers
func FindBusiestCustomerSuccess(customersByCustomerSuccess map[int]int) int {
	greatestCustomersQuantity := 0
	busiestCustomerID := 0

	for customerSuccessID, customersQuantity := range customersByCustomerSuccess {
		if customersQuantity > greatestCustomersQuantity {
			greatestCustomersQuantity = customersQuantity
			busiestCustomerID = customerSuccessID
		}
	}

	return busiestCustomerID
}

// FindAvailableCustomerSuccess retrieves all available customerSuccess given a list of customerSuccess away
func FindAvailableCustomerSuccess(customerSuccess []Entity, customerSuccessAwayIDs []int) []Entity {

	var availableCustomerSuccess []Entity

	for _, customerSuccessAwayID := range customerSuccessAwayIDs {
		for _, customerSuccessToBeChecked := range customerSuccess {
			if customerSuccessToBeChecked.ID != customerSuccessAwayID {
				availableCustomerSuccess = append(availableCustomerSuccess, customerSuccessToBeChecked)
			}
		}
	}
	return availableCustomerSuccess
}

// SortEntitiesByScoreDesc sorts Entities descending by their IDs
func SortEntitiesByScoreDesc(customerSuccess []Entity) []Entity {
	sort.Slice(customerSuccess[:], func(i, j int) bool {
		return customerSuccess[j].Score < customerSuccess[i].Score
	})

	return customerSuccess
}
