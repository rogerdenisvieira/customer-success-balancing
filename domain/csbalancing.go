package csbalancing

import (
	"errors"
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

	var notAttendedCustomers []Entity
	var customersByCustomerSuccess = make(map[int]int) // [CustomerSuccessID]CustomerQuantity
	cssByScore := SortEntitiesByScoreDesc(customerSuccess)
	customersByScore := SortEntitiesByScoreDesc(customers)

	if len(cssByScore) == 0 { // early return: no customers
		return 0
	}

	for _, customer := range customersByScore {

		suitableCS, error := FindSuitableCS(cssByScore, customer, customerSuccessAway)

		if error == nil {
			customersByCustomerSuccess[suitableCS.ID]++
		} else {
			notAttendedCustomers = append(notAttendedCustomers, customer)
		}
	}

	return FindBusiestCustomerSuccess(customersByCustomerSuccess)
}

// FindSuitableCS given a list of CSs and a customer, retrieve the most suitable Cs to attend the customer
func FindSuitableCS(customerSuccess []Entity, customer Entity, customerSuccessAway []int) (Entity, error) {

	var suitableCS Entity
	var errorsFound error

	availableCSSByScore := FindAvailableCustomerSuccess(customerSuccess, customerSuccessAway)

	if len(availableCSSByScore) == 0 { // early return: no CS available
		return suitableCS, errors.New("No CustomerSuccess available was found")
	}

	for availableCSSIndex, availableCS := range availableCSSByScore {

		// formattedText := fmt.Sprintf("CustomerScore: %d, CustomerSuccessScore: %d,  CustomerSuccessIndex: %d", customer.Score, availableCS.Score, availableCSSIndex)
		// fmt.Println(formattedText)

		if customer.Score > availableCS.Score && availableCSSIndex == 0 { // customer score greater than greatest CS score
			return suitableCS, errors.New("No suitable CustomerSuccess was found")
		} else if customer.Score <= availableCS.Score && availableCSSIndex == len(availableCSSByScore)-1 { // customer score less than or equals to last CS score
			return availableCS, nil
		} else if customer.Score <= availableCS.Score && customer.Score > availableCSSByScore[availableCSSIndex+1].Score { // customer score between current and next CS score
			return availableCS, nil
		}
	}

	return suitableCS, errorsFound

}

// O(n)
// FindBusiestCustomerSuccess finds the ID of the CustomerSuccess with most customers
func FindBusiestCustomerSuccess(customersByCustomerSuccess map[int]int) int {
	greatestCustomersQuantity := 0
	busiestCustomerID := 0

	for customerSuccessID, customersQuantity := range customersByCustomerSuccess {
		if len(customersByCustomerSuccess) == 1 { // there is only one customer success
			return customerSuccessID
		} else if customersQuantity != 0 && customersQuantity == greatestCustomersQuantity {
			busiestCustomerID = 0 // draw
		} else if customersQuantity > greatestCustomersQuantity {
			greatestCustomersQuantity = customersQuantity
			busiestCustomerID = customerSuccessID
		}

	}

	return busiestCustomerID
}

// FindAvailableCustomerSuccess retrieves all available customerSuccess given a list of customerSuccess away
func FindAvailableCustomerSuccess(customerSuccess []Entity, customerSuccessAwayIDs []int) []Entity {

	if len(customerSuccessAwayIDs) == 0 { // no customer success is away
		return customerSuccess
	}

	var availableCustomerSuccess []Entity

	for _, customerSuccessToBeChecked := range customerSuccess {

		if !ContainsInt(customerSuccessToBeChecked.ID, customerSuccessAwayIDs) {
			availableCustomerSuccess = append(availableCustomerSuccess, customerSuccessToBeChecked)
		}

	}

	return availableCustomerSuccess
}

//ContainsInt checks whether a array of integers contains a given int
func ContainsInt(number int, numbers []int) bool {
	for _, currentNumber := range numbers {

		if number == currentNumber {
			return true
		}
	}
	return false
}

// SortEntitiesByScoreDesc sorts Entities descending by their IDs
func SortEntitiesByScoreDesc(customerSuccess []Entity) []Entity {
	sort.Slice(customerSuccess[:], func(i, j int) bool {
		return customerSuccess[j].Score < customerSuccess[i].Score
	})

	return customerSuccess
}
