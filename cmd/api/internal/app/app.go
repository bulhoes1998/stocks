package app

import (
	"log"
	"net/http"

	"github.com/bulhoes1998/stock/cmd/api/internal/database"
	"github.com/bulhoes1998/stock/internal/session"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handlers struct {
	Session session.SessionHandler
}

func BuildApplication() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	sessionRepository := session.NewSessionRepository(db)
	sessionService := session.NewSessionService(sessionRepository)

	ctrl := &Handlers{
		Session: session.NewSessionHandler(sessionService),
	}

	routes(r, ctrl)

	http.ListenAndServe(":8080", r)
}
