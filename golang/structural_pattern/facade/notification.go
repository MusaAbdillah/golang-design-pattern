package main

import "fmt"

type Notification struct {
}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("Sending credit wallet notification")
	return
}

func (n *Notification) sendWalletDebiteNotification() {
	fmt.Println("Sending debite wallet notification")
	return
}


