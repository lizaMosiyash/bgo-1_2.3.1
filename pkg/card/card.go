package card

import (
	"fmt"
	"sort"
)

type Transaction struct {
	Id int64
	Sum int64
	Date int64
	Status string
	MccCode string
}

type Card struct {
	Id           int
	Balance      int
	Currency     string
	Number       string
	Transactions []*Transaction
}




func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, transaction)
}

func SortBySum (transactions []*Transaction) {
	sort.SliceStable(transactions, func(i, j int) bool {return transactions[i].Sum > transactions[j].Sum })
	for i, _ := range transactions {
		fmt.Println("for sum", transactions[i].Sum)
	}
}