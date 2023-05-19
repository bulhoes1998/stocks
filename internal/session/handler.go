package session

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/bulhoes1998/stock/utils"
	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

type SessionHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Signup(w http.ResponseWriter, r *http.Request)
}

type sessionHandler struct {
	sessionService SessionService
}

func NewSessionHandler(sessionService SessionService) SessionHandler {
	return &sessionHandler{
		sessionService: sessionService,
	}
}

func (h *sessionHandler) Login(w http.ResponseWriter, r *http.Request) {
	var login UserLogin
	json.NewDecoder(r.Body).Decode(&login)
	if login.Username == "" && login.Email == "" {
		utils.Encoder(w, "Username or Email is required", http.StatusBadRequest)
	}
	if login.Password == "" {
		utils.Encoder(w, "Password is required", http.StatusBadRequest)
		return
	}

	user, err := h.sessionService.Login(login)
	if err != nil {
		utils.Encoder(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidateUser([]byte(user.Password), []byte(login.Password)); err != nil {
		utils.Encoder(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, _ := CreateJWT(user.Id)
	fmt.Println(token)
	utils.Encoder(w, user, http.StatusOK)
}

func (h *sessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	h.sessionService.Logout()
}

func (h *sessionHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var signup UserSignup
	json.NewDecoder(r.Body).Decode(&signup)
	h.sessionService.Signup(signup)
}

func ValidateUser(hash, pass []byte) error {
	if err := bcrypt.CompareHashAndPassword(hash, pass); err != nil {
		return fmt.Errorf("username or password is incorrect")
	}

	return nil
}

func CreateJWT(id uint) (string, error) {
	level := jwt.MapClaims{}
	level["authorized"] = true
	level["exp"] = time.Now().Add(time.Hour * 5).Unix()
	level["user_id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, level)
	return token.SignedString([]byte("secret"))
}
