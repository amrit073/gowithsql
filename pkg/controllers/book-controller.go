package controllers

import (
	"encoding/json"
	//"fmt"
	//"fmt"
	//"io/ioutil"

	// 	"fmt"
	"github.com/amrit073/gosql/pkg/models"
	"github.com/amrit073/gosql/pkg/utils"
	"github.com/gorilla/mux"

	// 	"github.com/amrit073/gosql/pkg/utils"
	// 	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// 	nb := &models.Book{}

	// 	var m models.Book
	// 	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// 	if m.Author == "" || m.Name == "" || m.Publication == "" {
	//
	// 		w.Header().Set("Content-type", "application/json")
	// 		w.WriteHeader(http.StatusBadRequest)
	// 		fmt.Fprint(w, "{success:false, msg:\"need to provide all fields\"}")
	// 		return
	// 	}
	// 	fmt.Println(m.Author)
	// 	utils.ParseBody(r, nb)
	// 	result := nb.CreateBook()
	// 	_, err := json.Marshal(result)
	// 	w.Header().Set("Content-type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	if err != nil {
	// 		fmt.Fprint(w, "{success:true}")
	// 	}
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)
	Id, _ := strconv.ParseUint(parms["bookId"], 10, 64)
	res, _ := json.Marshal(models.GetBookBYId(Id))
	w.WriteHeader(http.StatusOK)
	w.Write(res)

	//models.GetBookBYId(parms["bookId"])

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	parms := mux.Vars(r)
	Id, _ := strconv.ParseUint(parms["bookId"], 10, 64)
	res, _ := json.Marshal(models.DeleteBook(Id))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	updatebook := &models.Book{}
	utils.ParseBody(r, updatebook)
	parms := mux.Vars(r)
	Id, _ := strconv.ParseUint(parms["bookId"], 10, 64)
	models.DeleteBook(Id)
	b := updatebook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
