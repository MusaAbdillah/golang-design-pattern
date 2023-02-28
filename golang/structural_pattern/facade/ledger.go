package main

import "fmt"

type Ledger struct {
	
}

func (l *Ledger) makeEntry(accountID, txnType string, amount int){
	fmt.Printf("Make ledger entry for account id %s with txnType %s for amount %d\n", accountID, txnType, amount)
	return 
}
