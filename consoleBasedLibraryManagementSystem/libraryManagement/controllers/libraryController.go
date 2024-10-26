package controllers

import (
	"bufio"
	"fmt"
	"libraryManagement/models"
	"libraryManagement/services"
	"os"
	"strconv"
	"strings"
)

var library = services.NewLibrary()

func ShowMenu(){
	fmt.Println("Library Management System")
	fmt.Println("1. Add a new Book")
	fmt.Println("2. Remove an existing book")
	fmt.Println("3. Borrow a book")
	fmt.Println("4. Return a book")
	fmt.Println("5. List all available books")
	fmt.Println("6. List all borrowed books by a member")
	fmt.Println("7. Exit")
}

func GetUserInputs(prompt string) string{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func AddBook(){
	id, _:= strconv.Atoi(GetUserInputs("Enter Book ID: "))
	title := GetUserInputs("Enter Book Title: ")
	author := GetUserInputs("Enter Author Name: ")

	book := models.Book{
		ID : id, 
		Title: title,
		Author: author,
		Status: "Available",
	}

	library.AddBook(book)
	fmt.Println("Book added successfully!")
}

func RemoveBook() {
    id, _ := strconv.Atoi(GetUserInputs("Enter Book ID to remove: "))
    library.RemoveBook(id)
    fmt.Println("Book removed successfully!")
}

func BorrowBook() {
    bookID, _ := strconv.Atoi(GetUserInputs("Enter Book ID to borrow: "))
    memberID, _ := strconv.Atoi(GetUserInputs("Enter Member ID: "))

    err := library.BorrowBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book borrowed successfully!")
    }
}

func ReturnBook() {
    bookID, _ := strconv.Atoi(GetUserInputs("Enter Book ID to return: "))
    memberID, _ := strconv.Atoi(GetUserInputs("Enter Member ID: "))

    err := library.ReturnBook(bookID, memberID)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Book returned successfully!")
    }
}

func ListAvailableBooks() {
    books := library.ListAvailableBooks()
    fmt.Println("Available Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}

func ListBorrowedBooks() {
    memberID, _ := strconv.Atoi(GetUserInputs("Enter Member ID: "))
    books := library.ListBorrowedBooks(memberID)
    fmt.Println("Borrowed Books:")
    for _, book := range books {
        fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
    }
}