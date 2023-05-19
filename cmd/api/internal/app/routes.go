package app

import (
	"github.com/go-chi/chi"
)

func routes(r *chi.Mux, h *Handlers) {
	r.Post("/login", h.Session.Login)
	r.Post("/logout", h.Session.Logout)
	r.Post("/signup", h.Session.Signup)
}
