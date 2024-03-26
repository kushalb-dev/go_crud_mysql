package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/kushalb-dev/bookstore_management/pkg/models"
	"github.com/kushalb-dev/bookstore_management/pkg/utils"
)

var Book models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBooks()
	resp, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "pkgapplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	bookById, _ := models.GetBookById(ID)

	resp, _ := json.Marshal(bookById)
	w.Header().Set("content-type", "pkgapplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	b := Book.CreateBook()
	resp, _ := json.Marshal(b)
	w.Header().Set("content-type", "pkgapplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing at delete book")
	}
	deletedBook, _ := models.DeleteBook(ID)
	resp, _ := json.Marshal(deletedBook)
	w.Header().Set("content-type", "pkgapplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatedBook := &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookId := vars["id"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing at updatebook")
	}

	bookDetails, db := models.GetBookById(ID)

	if updatedBook.Author != "" {
		bookDetails.Author = updatedBook.Author
	}

	if updatedBook.Name != "" {
		bookDetails.Name = updatedBook.Name
	}

	if updatedBook.Publication != "" {
		bookDetails.Publication = updatedBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("content-type", "pkgapplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
