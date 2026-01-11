package transport

import (
	"encoding/json"
	"net/http"
	"restapi/go/internal/model"
	"restapi/go/internal/service"
	"strconv"
	"strings"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler{
		service: s,
	}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		books, err := h.service.GetAllBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)

	case http.MethodPost:
		var book model.Book

		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		created, err := h.service.CreateBook(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application-json")
		json.NewEncoder(w).Encode(created)

	default:
		http.Error(w, "Method not Implemented", http.StatusMethodNotAllowed)
	}
}

func (h *BookHandler) HandleBookById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Id Required", http.StatusBadRequest)
	}

	switch r.Method {
	case http.MethodGet:
		book, err := h.service.GetByIdBook(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)

	case http.MethodPut:
		var book model.Book

		err := json.NewDecoder(r.Body).Decode(&book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		updated, err := h.service.UpdateBook(id, &book)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application-json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		err := h.service.RemoveBook(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application-json")
		json.NewEncoder(w).Encode("Book Deleted")

	default:
		http.Error(w, "Method not Implemented", http.StatusMethodNotAllowed)
	}
}
