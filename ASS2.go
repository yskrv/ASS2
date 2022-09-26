package main

import "fmt"

type makeup interface {
	cost() int
	description() string
}

type eveningMakeup struct {
}

func (e eveningMakeup) cost() int {
	return 10000
}

func (e eveningMakeup) description() string {
	return "Evening makeup : 10000tg"
}

type everydayMakeup struct {
}

func (e everydayMakeup) cost() int {
	return 7000
}

func (e everydayMakeup) description() string {
	return "Everyday makeup : 7000tg"
}

type manicure struct {
	mup makeup
}

func (m manicure) description() string {
	return m.mup.description() + "\nManicure : 6000tg "
}

func (m manicure) cost() int {
	return m.mup.cost() + 6000
}

type hairstyle struct {
	mup makeup
}

func (h hairstyle) description() string {
	return h.mup.description() + "\nHairstyle : 7000tg"
}

func (h hairstyle) cost() int {
	return h.mup.cost() + 7000
}

type pedicure struct {
	mup makeup
}

func (p pedicure) description() string {
	return p.mup.description() + "\nPedicure : 8000tg"
}
func (p pedicure) cost() int {
	return p.mup.cost() + 8000
}

func main() {
	order1 := pedicure{hairstyle{eveningMakeup{}}}
	fmt.Println(order1.description())
	fmt.Printf("Total cost : %vtg ", order1.cost())
}
