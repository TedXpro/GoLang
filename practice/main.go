package main

import (
	"fmt"
	"sort"
	"strings"
)

var name = "Hello"

func main() {
	// fmt.Println("Hello World!");

	// dataTypes();

	// formatStrings();

	// arrays();

	// libraries();

	// loops();

	// names := []string{"mario", "Yoshi", "luigi", "wario"};
	// cycleNames(names, helper);

	firstName, secondName := getInitials("Yohannes Woldeyes");
	fmt.Println(firstName, secondName);

	firstName, secondName = getInitials("Henock");
	fmt.Println(firstName, secondName);

}

func getInitials(name string) (string, string) {
	s := strings.ToUpper(name);
	fmt.Println(s);
	names := strings.Split(s, " ");
	fmt.Println(names);

	var initials [] string;
	for _, val := range(names) {
		initials = append(initials, val[: 1]);
		fmt.Println(val);
	}
	fmt.Println(initials);

	if len(initials) > 1 {
		return	initials[0], initials[1];
	}

	return initials[0], "_";
}

func helper(name string){
	fmt.Println("Good Morning ", name);
}

func cycleNames(names []string, f func(string)){
	for _, val := range(names){
		f(val);
	}
}

func loops(){
	// x := 0;
	// for x < 5 {
	// 	fmt.Println("value of x is", x);
	// 	x++;
	// }

	// for i := 0; i < 5; i++ {
	// 	fmt.Println("value of i is", i);
	// }

	names:= []string{"mario", "luigi", "peach", "wario"};
	for i:= 0; i < len(names); i++{
		fmt.Println(names[i]);
	}

	for index, val := range(names) {
		fmt.Println(index + 1, ": ", val);
	}

	for _, val := range(names) {
		fmt.Println(val);
		val = "new String";
		fmt.Println(val);
	}
	println(names);
}

func libraries() {
	// greetings := "Hello there! Hello World";
	// fmt.Println(strings.Contains(greetings, "Hello"));
	// fmt.Println(strings.ReplaceAll(greetings, "Hello", "Hi"));

	// fmt.Println(strings.Count(greetings, "Hello"));
	// fmt.Println(strings.Index(greetings, "ll"));

	// fmt.Println(strings.Split(greetings, " "));
	ages := []int{45, 60, 50, 70, 80, 90};
	fmt.Println("before Sorting", ages);
	sort.Ints(ages);
	fmt.Println("after Sorting", ages);

	index:= sort.SearchInts(ages, 75);
	fmt.Println(index);
	ages = append(ages[0: index], 75, ages[index]);
	fmt.Println(ages);
}

func arrays() {
	// var ages [3]int = [3]int{20, 25, 02}
	var ages = [3]int{10, 20, 30}
	fmt.Println(ages, len(ages))

	names := [4]string{"mario", "luigi", "yoshi", "peach"}
	names[3] = "wario"
	fmt.Println(names)

	// sclices;
	var scores = []int{100, 50, 60}
	fmt.Println(scores, len(scores))
	scores[2] = 1000
	scores2 := append(scores, 90) // it creates a new array.
	fmt.Println(scores, len(scores))
	fmt.Println(scores2, len(scores2))

	// slice ranges
	rangeOne := scores2[1:4] // from position one to position 4 not including 4
	fmt.Println(rangeOne)
	// rangeOne = scores2[1: len(scores2)]; // till the last element. Not recommended
	fmt.Println(rangeOne)
	rangeOne = scores2[1:] // from position one to the last element
	fmt.Println(rangeOne)
	rangeOne = scores2[:3] // from position one to the last element
	fmt.Println(rangeOne)

	rangeOne = append(rangeOne, 1234)
	fmt.Println(rangeOne, len(rangeOne))
}

func dataTypes() {
	/*  strings */
	var nameOne string = "mario234"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Abebe"
	nameTwo = "Peach"
	nameThree = "Bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Yoshi"
	fmt.Println(nameFour)

	nameFour = "Wario"
	fmt.Println(nameFour, name)

	/*  integers */
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	/* 	bits and memory */
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint = 12310231123123123123

	fmt.Println(numOne, numThree, numTwo)

	/*  Float */
	var scoreOne float64 = 90.23
	var scoreTwo = 1230123.123
	scoreThree := 8909.123
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}

func formatStrings() {
	/* format Strings  */
	fmt.Print("hello, ")
	fmt.Print("world! \n")
	fmt.Print("new line \n")

	fmt.Println("hello again")
	fmt.Println("goodbye!")

	age := 35
	name := "yohannes"
	fmt.Println("my age is", age, "and my name is", name)

	fmt.Printf("my age is %v and my name is %q\n", age, name)

	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %f points\n", 22.5)
	fmt.Printf("you scored %0.1f points\n", 2254.5923)
	str := fmt.Sprintf("my age is %v and my name is %v\n", age, name)
	fmt.Println(str)
}
