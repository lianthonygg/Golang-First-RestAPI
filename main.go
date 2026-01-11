package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"restapi/go/internal/service"
	"restapi/go/internal/store"
	"restapi/go/internal/transport"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	q := `
		CREATE TABLE IF NOT EXISTS books (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			author TEXT NOT NULL
		)
	`

	if _, err := db.Exec(q); err != nil {
		log.Fatal(err.Error())
	}

	bookStore := store.New(db)
	bookService := service.New(bookStore)
	bookHandler := transport.New(bookService)

	http.HandleFunc("/books", bookHandler.HandleBooks)
	http.HandleFunc("/books/", bookHandler.HandleBookById)

	fmt.Println("Servidor Corriendo en el puerto 8080")
	fmt.Println("API Endpoints:")
	fmt.Println("	GET 		  /books			    - Obtener Todos los Libros")
	fmt.Println("	GET 		  /books/{id}			- Obtener un Libro Especifico")
	fmt.Println("	POST 		  /books			    - Crear un nuevo Libro")
	fmt.Println("	PUT 		  /books/{id}			- Actualizar un Libro")
	fmt.Println("	DELETE 		/books/{id}			- Eliminar un Libro")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
