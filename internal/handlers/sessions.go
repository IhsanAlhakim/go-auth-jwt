package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/IhsanAlhakim/go-auth-api/internal/auth"
	"github.com/IhsanAlhakim/go-auth-api/internal/validation"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) SignIn(w http.ResponseWriter, r *http.Request) {

	var credentials = struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	if err := BindJSON(r, &credentials); err != nil {
		if err == ErrEmptyBody {
			http.Error(w, err.Error(), http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	if err := validation.CheckStructEmptyProperty(credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := validation.CheckStructWhitespaceProperty(credentials); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var (
		user   User
		userId string
	)

	row := h.db.QueryRow("SELECT id, username, password FROM users WHERE email = ?", credentials.Email)
	if err := row.Scan(&userId, &user.Username, &user.Password); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := auth.VerifyPassword(user.Password, credentials.Password); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	token, err := auth.GenerateToken(userId, user.Username, *h.cfg)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{
		Name:    h.cfg.TokenCookieName,
		Value:   token,
		Expires: time.Now().Add(time.Duration(1) * time.Hour),
	}
	http.SetCookie(w, cookie)
	Response(w, P{Message: "Sign In Successfull"}, http.StatusOK)
}

func (h *Handler) SignOut(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:    h.cfg.TokenCookieName,
		Expires: time.Unix(0, 0),
		MaxAge:  -1,
	}
	http.SetCookie(w, cookie)
	Response(w, P{Message: "Sign Out Successful"}, http.StatusOK)
}
