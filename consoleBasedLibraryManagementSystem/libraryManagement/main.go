package main

import (
    "fmt"
    "libraryManagement/controllers"
)

func main() {
    for {
        controllers.ShowMenu()
        choice := controllers.GetUserInputs("Enter your choice: ")

        switch choice {
        case "1":
            controllers.AddBook()
        case "2":
            controllers.RemoveBook()
        case "3":
            controllers.BorrowBook()
        case "4":
            controllers.ReturnBook()
        case "5":
            controllers.ListAvailableBooks()
        case "6":
            controllers.ListBorrowedBooks()
        case "7":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }
}
