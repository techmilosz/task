package api

import (
	"net/http"

	"repartners/docs"

	"github.com/go-chi/chi/v5"
)

func (a *API) MountDocs(r chi.Router) {
	r.Route("/docs", func(r chi.Router) {
		r.Get("/", a.handleSwagger())
		r.Get("/oapi", a.handleOAPI())
	})
}

func (a *API) handleSwagger() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentType, contentTypeHTML)
		w.Write(docs.Swagger)
	}
}

func (a *API) handleOAPI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(headerContentType, contentTypeYAML)
		w.Write(docs.OAPI)
	}
}
