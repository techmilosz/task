package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *API) MountPacks(r chi.Router) {
	r.Route("/packs", func(r chi.Router) {
		r.Get("/", a.handleGetAllPacks())
		r.Post("/", a.handleAddPack())
		r.Delete(fmt.Sprintf("/{%s}", ParamValue), a.handleDeletePack())
	})
}

func (a *API) handleGetAllPacks() http.HandlerFunc {
	type getAllPacksResponse struct {
		Packs []int `json:"packs"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		packs := a.packs.GetAll()
		res := getAllPacksResponse{
			Packs: packs,
		}

		payload, err := json.Marshal(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(mustMarshalError(ErrInternalError.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(payload)
	}
}

func (a *API) handleAddPack() http.HandlerFunc {
	type addPackPayload struct {
		Value int `json:"value"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		payload, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(mustMarshalError(ErrInternalError.Error()))
			return
		}

		data := &addPackPayload{}
		if err := json.Unmarshal(payload, data); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustMarshalError(ErrInvalidBody.Error()))
			return
		}

		if data.Value < 0 {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustMarshalError(ErrInvalidBody.Error()))
			return
		}

		a.packs.Add(data.Value)
		w.WriteHeader(http.StatusNoContent)
	}
}

func (a *API) handleDeletePack() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, ParamValue)
		intVal, err := strconv.Atoi(val)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustMarshalError(ErrInvalidParameter.Error()))
			return
		}

		a.packs.Remove(intVal)
		w.WriteHeader(http.StatusNoContent)
	}
}
