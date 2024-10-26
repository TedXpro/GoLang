package services

import (
	"errors"
	"libraryManagement/models"
)

type libraryManager interface{
    AddBook(book models.Book)
    RemoveBook(bookId int)
    BorrowBook(bookId int, memberId int) error
    ReturnBook(bookId int, memberId int) error
    ListAvailableBooks() []models.Book
    ListBorrowedBooks(memberId int) []models.Book

}

type Library struct {
    Books map[int]models.Book
    Members map[int]models.Member
}

func NewLibrary() *Library{
    return &Library{
        Books: make(map[int]models.Book),
        Members: make(map[int]models.Member),
    }
}

func (l *Library) AddBook(book models.Book){
    l.Books[book.ID] = book
}

func (l *Library) RemoveBook(bookId int){
    delete(l.Books, bookId)
}

func (l *Library) BorrowBook(bookId int, memberId int) error{
    book, bookExists := l.Books[bookId]
    member, memberExists := l.Members[memberId]

    if !bookExists {
        return errors.New("Book Not Found!")
    }

    if !memberExists{
        return errors.New("Member Not Found")
    }

    book.Status = "Borrowed"
    l.Books[bookId] = book

    member.BorrowedBooks = append(member.BorrowedBooks, book)
    l.Members[memberId] = member

    return nil
}

func (l *Library) ReturnBook(bookId int, memberId int) error{
    book, bookExists := l.Books[bookId]
    member, memberExists := l.Members[memberId]

    if !bookExists {
        return errors.New("Book Not Found!")
    }

    if !memberExists{
        return errors.New("Member Not Found")
    }

    if book.Status == "Available"{
        return errors.New("Book Is Not Borrowed!")
    }

    book.Status = "Available"
    l.Books[bookId] = book

    for i, b := range member.BorrowedBooks{
        if b.ID == bookId{
            member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i + 1:]...)
            break
        }
    }

    l.Members[memberId] = member

    return nil
}

func (l *Library) ListAvailableBooks() []models.Book{
    var availableBooks []models.Book
    for _, book := range l.Books{
        if book.Status == "Available"{
            availableBooks = append(availableBooks, book)
        }
    }

    return availableBooks
}

func (l * Library) ListBorrowedBooks(memberId int) []models.Book{
    member, exists := l.Members[memberId]
    if !exists{
        return nil
    }

    return member.BorrowedBooks
}