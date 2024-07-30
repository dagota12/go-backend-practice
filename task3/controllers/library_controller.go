package controllers

import (
	"fmt"
	"goPractice/task3/models"
	"math/rand"
	"time"
)

type Library struct {
	Members map[int]models.Member `json:"members"`
	Books   map[int]models.Book   `json:"books"`
}

type LibraryManager interface {
	AddBook(book models.Book)
	AddMember(book models.Member)
	ListMembers() map[int]models.Member
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]models.Book),
		Members: make(map[int]models.Member),
	}
}

func (l *Library) AddMember(member models.Member) error {
	if _, ok := l.Members[member.ID]; ok {
		return fmt.Errorf("member with id %d already exists", member.ID)
	}
	l.Members[member.ID] = member
	return nil
}
func (l *Library) ListMembers() map[int]models.Member {
	return l.Members
}

func (l *Library) AddBook(book models.Book) error {
	if _, ok := l.Books[book.ID]; ok {
		return fmt.Errorf("book with ID %d already exists", book.ID)
	}
	l.Books[book.ID] = book
	return nil
}
func (l *Library) RemoveBook(bookID int) error {
	if _, ok := l.Books[bookID]; !ok {
		return fmt.Errorf("book with ID %d not found", bookID)
	}
	delete(l.Books, bookID)
	return nil
}

// BorrowBook borrows a book with the given bookID and memberID.
//
// Parameters:
// - bookID: the ID of the book to be borrowed.
// - memberID: the ID of the member borrowing the book.
//
// Returns:
// - error: an error if the book or member is not found, or if the book is already borrowed.

func (lib *Library) BorrowBook(bookID int, memberID int) error {
	book, ok := lib.Books[bookID]
	if !ok {
		return fmt.Errorf("book not found")
	}
	if !book.Status {
		return fmt.Errorf("book with ID %d is already borrowed", bookID)
	}
	member, ok := lib.Members[memberID]
	if !ok {
		return fmt.Errorf("membember not found")
	}
	if _, ok := lib.Members[memberID].BorrowedBooks[bookID]; ok {
		return fmt.Errorf("book already borrowed")
	}
	book.Status = false
	member.BorrowedBooks[bookID] = book
	lib.Books[bookID] = book
	lib.Members[memberID] = member

	return nil
}

// ReturnBook returns the borrowed book with the given bookID and memberID to the library.
//
// Parameters:
// - bookID: the ID of the book to be returned.
// - memberID: the ID of the member returning the book.
//
// Returns:
// - error: an error if the book or member is not found, or if the book was not borrowed by the member.
func (lib *Library) ReturnBook(bookID int, memberID int) error {
	if !lib.findBook(bookID) || !lib.findMember(memberID) {
		return fmt.Errorf("book not found or member not found")
	}
	book, ok := lib.Members[memberID].BorrowedBooks[bookID]
	if !ok {
		return fmt.Errorf("you did not borrow this Book")
	}

	member := lib.Members[memberID]
	delete(member.BorrowedBooks, bookID)
	book.Status = true
	lib.Books[bookID] = book
	lib.Members[memberID] = member
	return nil
}
func (lib *Library) findBook(bookID int) bool {
	_, ok := lib.Books[bookID]
	return ok
}
func (lib *Library) findMember(memberID int) bool {
	_, ok := lib.Members[memberID]
	return ok
}
func (l *Library) ListAvailableBooks() []models.Book {
	var availableBooks []models.Book
	for _, book := range l.Books {
		//if book is available
		if book.Status {
			availableBooks = append(availableBooks, book)
		}
	}
	return availableBooks
}
func (lib *Library) ListBorrowedBooks(memberID int) map[int]models.Book {
	member, ok := lib.Members[memberID]
	if !ok {
		fmt.Println("Member not found return nil")
		return nil
	}

	return member.BorrowedBooks
}
func Gen_id() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(1000000)
}
