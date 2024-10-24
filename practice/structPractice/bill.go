package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// reciever functions for my struct
func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	for key, val := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key + ":", val);
		total += val;
	}

	fs += "----------------------------------------------\n";
	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.tip);

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "total:", total + b.tip);

	return fs;
}

func (b *bill) updateTip(tip float64){
	b.tip = tip;
	fmt.Println(b.tip);
}

func (b *bill) addItem(name string, price float64){
	b.items[name] = price;
}

func (b *bill) saveBill(){
	data := []byte(b.format())

	err := os.WriteFile("billsFolder/"+b.name+".txt", data,0644) // 0644 is permission for the file.
	if(err != nil) {
		panic(err)
	}

	fmt.Println("Bill was saved to file");
}