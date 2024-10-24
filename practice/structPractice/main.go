package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputs(prompt string, reader *bufio.Reader) (string, error){
	fmt.Print(prompt);

	/*  read everyting they have entetred after clicking
	on the enter button which is the new line */
	input, err := reader.ReadString('\n')

	return strings.TrimSpace(input), err;
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInputs("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("created the bill -", b.name)

	return b
}

func promptOptions(b bill){
	reader := bufio.NewReader(os.Stdin);

	prompt := "Choose option (a - add item, s - save bill, t - add tip): ";
	opt, _ := getInputs(prompt, reader);

	switch opt {
	case "a":
		name, _ := getInputs("Enter Item Name: ", reader);
		price, _:= getInputs("Enter the Price for " + name + ": ", reader)

		p, err := strconv.ParseFloat(price, 64);
		if err != nil {
			fmt.Println("The Price must be a number!!")
			promptOptions(b)
			return
		}

		b.addItem(name, p);
		fmt.Println("Item added - ", name, price)
		
		promptOptions(b)
	case "t":
		tip, _ := getInputs("Enter tip amount($): ", reader)
		t, err := strconv.ParseFloat(tip, 64);
		if err != nil {
			fmt.Println("The Price must be a number!!")
			promptOptions(b)
			return
		}

		b.updateTip(t);
		fmt.Println("Tip added - ",tip)
		
		promptOptions(b)
	case "s" :
		b.saveBill()
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()

	promptOptions(myBill)
	fmt.Println(myBill)
}
