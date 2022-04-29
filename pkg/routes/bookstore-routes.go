package routes

import (
	"github.com/gorilla/mux"
	"go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBookIByd).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBookById).Methods("DELETE")


}