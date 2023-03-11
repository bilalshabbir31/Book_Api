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

func createBook(c *gin.Context){
	var newbook book
	if err:=c.BindJSON(&newbook); err!=nil{
		return
	}
	books=append(books,newbook)
	c.IndentedJSON(http.StatusCreated,newbook)
}

func bookbyid(c *gin.Context){
	id :=c.Param("id")
	book,err:=getBookById(id)

	if err!=nil{
		c.IndentedJSON(http.StatusNotFound,gin.H{"message":"book not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK,book)
}

func getBookById(id string)(*book,error ){
	for i,b := range books{
		if b.ID==id{
			return &books[i], nil
		}
	}
	return nil,errors.New("book not found")
}

func main(){
	router :=gin.Default()
	router.GET("/books",getBooks)
	router.GET("/books/:id",bookbyid)
	router.POST("/books",createBook)
	router.Run("localhost:8080")
}
