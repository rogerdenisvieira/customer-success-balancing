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

// type ScorePrinter interface {
// 	PrintScores(entities []Entity)
// }

// Todos os CSs têm níveis diferentes
// Não há limite de clientes por CS
// Clientes podem ficar sem serem atendidos
// Clientes podem ter o mesmo tamanho
// 0 < n < 1.000
// 0 < m < 1.000.000
// 0 < id do cs < 1.000
// 0 < id do cliente < 1.000.000
// 0 < nível do cs < 10.000
// 0 < tamanho do cliente < 100.000
// Valor máximo de t = n/2 arredondado para baixo

func CustomerSuccessBalancing(customerSuccess []Entity, customers []Entity, customerSuccessAway []int) int {
	// Write your solution here
	customersUnattended := 0

	return 0
}

func SortCustomerSuccessByScore(customerSuccess []Entity) {
	sort.Slice(customerSuccess[:], func(i, j int) bool {
		return customerSuccess[i].Score < customerSuccess[j].Score
	})
}

func PrintScores(entities []Entity) {

	for _, v := range entities {
		fmt.Println(v.Score)
	}

}

func GetCustomerTotalSize(customers []Entity) int {
	size := 0

	for _, v := range customers {
		size += v.Score
	}

	return size
}
