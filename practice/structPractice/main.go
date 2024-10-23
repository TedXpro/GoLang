package main

import (
	"bufio"
	"fmt"
	"os"
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

	// fmt.Print("Create a new bill name: ")

	/*  read everyting they have entetred after clicking
	on the enter button which is the new line */
	// name, _ := reader.ReadString('\n')

	// name = strings.TrimSpace(name)

	name, _ := getInputs("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("created the bill -", b.name)

	return b
}

func main() {
	myBill := createBill()

	fmt.Println(myBill)
}
