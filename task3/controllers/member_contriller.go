package controllers

import (
	"goPractice/task3/models"
)

func NewUser() *models.Member {
	return &models.Member{

		BorrowedBooks: make(map[int]models.Book),
	}
}
