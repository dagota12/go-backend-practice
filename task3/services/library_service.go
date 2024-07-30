package services

import (
	"bufio"
	"fmt"
	"goPractice/task3/controllers"
	"goPractice/task3/models"
	"os"
	"strconv"
	"strings"
	"time"
)

type LibraryService struct {
	library *controllers.Library
}

func NewLibraryService() *LibraryService {
	return &LibraryService{
		library: controllers.NewLibrary(),
	}
}

func (libService *LibraryService) Run() {
	reader := bufio.NewReader(os.Stdin)
	library := libService.library

	for {
		fmt.Println("\nLibrary Management System")
		fmt.Println("0.Add User")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books")
		fmt.Println("7. List users")
		fmt.Println("8. Exit")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		switch choice {
		case "0":
			handleAddUser(reader, library)
		case "1":
			handleAddBook(reader, library)
		case "2":
			handleRemoveBook(reader, library)
		case "3":
			handleBorrowBook(reader, library)
		case "4":
			handleReturnBook(reader, library)
		case "5":
			handleListAvailableBooks(library)
		case "6":
			handleListBorrowedBooks(reader, library)
		case "7":
			handleListUsers(library)
		case "8":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func handleListUsers(library *controllers.Library) {
	users := library.ListMembers()
	if len(users) == 0 {
		fmt.Println("No users found")
		return
	}
	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s\n", user.ID, user.Name)
	}
}

func handleAddUser(reader *bufio.Reader, library *controllers.Library) {
	id := time.Now().UnixNano()
	fmt.Print("Enter user name: ")
	// var name string
	name, _ := reader.ReadString('\n')

	user := models.Member{
		ID:            int(id),
		Name:          strings.TrimSpace(name),
		BorrowedBooks: make(map[int]models.Book), // Initialize borrowed books map
	}

	err := library.AddMember(user)
	if err != nil {
		fmt.Println("Error adding user:", err)
		return
	}
	fmt.Println("User added successfully")
}

func handleAddBook(reader *bufio.Reader, library *controllers.Library) {
	id := time.Now().UnixNano()

	fmt.Print("Enter book title: ")
	// var title string
	// fmt.Scanln(&title)
	title, _ := reader.ReadString('\n')

	fmt.Print("Enter book author: ")
	author, _ := reader.ReadString('\n')
	// fmt.Scanln(&author)

	book := models.Book{
		ID:     int(id),
		Name:   strings.TrimSpace(title),
		Author: strings.TrimSpace(author),
		Status: true,
	}

	err := library.AddBook(book)
	if err != nil {
		fmt.Println("Error adding book:", err)
		return
	}
	fmt.Println("Book added successfully")
}

func handleRemoveBook(reader *bufio.Reader, library *controllers.Library) {
	fmt.Print("Enter book ID to remove: ")
	idStr, _ := reader.ReadString('\n')
	id, _ := strconv.Atoi(strings.TrimSpace(idStr))

	err := library.RemoveBook(id)
	if err != nil {
		fmt.Println("Error removing book:", err)
		return
	}
	fmt.Println("Book removed successfully")
}

func handleBorrowBook(reader *bufio.Reader, library *controllers.Library) {
	fmt.Print("Enter member ID: ")
	memberIdStr, _ := reader.ReadString('\n')

	memberId, _ := strconv.Atoi(strings.TrimSpace(memberIdStr))

	fmt.Print("Enter book ID: ")
	bookIdStr, _ := reader.ReadString('\n')
	bookId, _ := strconv.Atoi(strings.TrimSpace(bookIdStr))
	if bookIdStr == "" || memberIdStr == "" {
		fmt.Println("Invalid input")
		return
	}

	err := library.BorrowBook(bookId, memberId)
	if err != nil {
		fmt.Println("Error borrowing book:", err)
		return
	}
	fmt.Println("Book borrowed successfully")
}

func handleReturnBook(reader *bufio.Reader, library *controllers.Library) {
	fmt.Print("Enter member ID: ")
	memberIdStr, _ := reader.ReadString('\n')
	memberId, _ := strconv.Atoi(strings.TrimSpace(memberIdStr))

	fmt.Print("Enter book ID: ")
	bookIdStr, _ := reader.ReadString('\n')
	bookId, _ := strconv.Atoi(strings.TrimSpace(bookIdStr))

	err := library.ReturnBook(bookId, memberId)
	if err != nil {
		fmt.Println("Error returning book:", err)
		return
	}
	fmt.Println("Book returned successfully")
}

func handleListAvailableBooks(library *controllers.Library) {
	availableBooks := library.ListAvailableBooks()
	if len(availableBooks) == 0 {
		fmt.Println("No available books")
		return
	}
	fmt.Println("Available Books:")
	for _, book := range availableBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Name, book.Author)
	}
}

func handleListBorrowedBooks(reader *bufio.Reader, library *controllers.Library) {
	fmt.Print("Enter member ID: ")
	memberIdStr, _ := reader.ReadString('\n')
	memberId, _ := strconv.Atoi(strings.TrimSpace(memberIdStr))

	borrowedBooks := library.ListBorrowedBooks(memberId)
	if len(borrowedBooks) == 0 {
		fmt.Println("Member has no borrowed books")
		return
	}
	fmt.Println("Borrowed Books:")
	for _, book := range borrowedBooks {
		fmt.Printf("ID: %d, Title: %s, Author: %s\n", book.ID, book.Name, book.Author)
	}
}
