package server

import (
	"net/http"


	"github.com/muratmirgun/api/pkg/mid"

	"github.com/muratmirgun/api/internal/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"github.com/muratmirgun/api/internal/healthcheck"
	"github.com/muratmirgun/api/pkg/log"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// Routing setup api routing
func Routing(db *sqlx.DB, logger log.Logger) chi.Router {
	validate = validator.New()

	// setup server routing
	r := chi.NewRouter()

	// homepage welcome page
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.HTML(w, r, "<html><head><title>Go Starter Kit</title></head><body>Welcome to Go Starter Kit</head></body></html>")
	})

	// register health check route
	healthcheck.RegisterHandlers(r)

	// register v1 api path group
	r.Route("/v1", func(r chi.Router) {
		r.Use(mid.APIVersionCtx("v1"))
		user.RegisterHandlers(r, db, logger, validate)
	})

	return r
}
