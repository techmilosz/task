package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	ErrBodyLimitExceeded = errors.New("request body is too big")
	ErrInvalidBody       = errors.New("invalid body")
	ErrInvalidParameter  = errors.New("invalid parameter")
	ErrInternalError     = errors.New("internal server error")
)

const (
	ParamOrder = "order"
	ParamValue = "value"

	contentTypeJSON = "application/json"
	// 4MB
	maxRequestSize = 1024 * 1024 * 4

	headerContentType = "Content-Type"
)

type calculator interface {
	Calculate(order int) []int
}

type packsManager interface {
	Add(int)
	Remove(int)
	GetAll() []int
}

type API struct {
	packs packsManager
	calc  calculator
}

func New(packsManager packsManager, calculator calculator) *API {
	return &API{
		packs: packsManager,
		calc:  calculator,
	}
}

func (a *API) Run(ctx context.Context, port uint) error {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.SetHeader(headerContentType, contentTypeJSON))
	r.Use(middleware.AllowContentType(contentTypeJSON))
	r.Use(middleware.RequestSize(maxRequestSize))

	a.MountPacks(r)
	a.MountOrder(r)

	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
