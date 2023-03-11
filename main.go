package main

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)


type book struct{
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}


var books=[]book{

	{ID:"1" ,Title: "DLD",Author: "Rehan Saleem" ,Quantity: 50 },
	{ID:"2" ,Title: "Programming Fundamentals" ,Author:"Sheraz" ,Quantity:  60},
	{ID:"3" ,Title:"Machine Learning" ,Author:"Abdullah" ,Quantity: 100 },
}

func getBooks(c *gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}

func main(){
	router :=gin.Default()
	router.GET("/books",getBooks)
	router.Run("localhost:8080")
}
