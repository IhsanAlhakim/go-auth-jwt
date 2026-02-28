package handlers

import (
	"database/sql"
	"net/http"

	"github.com/IhsanAlhakim/go-auth-api/internal/auth"
	"github.com/IhsanAlhakim/go-auth-api/internal/validation"
)

type User struct {
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")

	var user User

	row := h.db.QueryRow("SELECT username, email FROM users WHERE id = ?", userId)
	if err := row.Scan(&user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	Response(w, P{Data: user}, http.StatusOK)
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {

	var user User

	if err := BindJSON(r, &user); err != nil {
		if err == ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := validation.CheckStructEmptyProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validation.CheckStructWhitespaceProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.GenerateHashPassword(user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = h.db.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", user.Username, user.Email, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Response(w, P{Message: "User Created!"}, http.StatusCreated)
}

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")

	var user = struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	}{}

	if err := BindJSON(r, &user); err != nil {
		if err == ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := validation.CheckStructEmptyProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validation.CheckStructWhitespaceProperty(user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := h.db.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", user.Username, user.Email, userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Response(w, P{Message: "User Updated!"}, http.StatusOK)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	userId := r.PathValue("id")

	var user User

	row := h.db.QueryRow("SELECT username FROM users WHERE id = ?", userId)
	if err := row.Scan(&user.Username); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	_, err := h.db.Exec("DELETE FROM users WHERE id = ?", userId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	Response(w, P{Message: "User Deleted!"}, http.StatusOK)
}
