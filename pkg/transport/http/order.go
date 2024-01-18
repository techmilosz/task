package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (a *API) MountOrder(r chi.Router) {
	r.Route("/order", func(r chi.Router) {
		r.Get(fmt.Sprintf("/{%s}", ParamOrder), a.handleCalculate())
	})
}

type payloadRow struct {
	Packet int `json:"packet,omitempty"`
	Amount int `json:"amount,omitempty"`
}

func (a *API) handleCalculate() http.HandlerFunc {
	type calculatePayload struct {
		Data []payloadRow `json:"data,omitempty"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		val := chi.URLParam(r, ParamOrder)
		intVal, err := strconv.Atoi(val)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write(mustMarshalError(ErrInvalidParameter.Error()))
			return
		}

		res := a.calc.Calculate(intVal)

		payload := &calculatePayload{Data: transformCalculateResponse(res)}
		payloadJson, err := json.Marshal(payload)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(mustMarshalError(ErrInternalError.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(payloadJson)
	}
}

func transformCalculateResponse(vals []int) []payloadRow {
	rows := make(map[int]int)
	for _, val := range vals {
		rows[val]++
	}

	r := []payloadRow{}
	for packet, amount := range rows {
		r = append(r, payloadRow{Packet: packet, Amount: amount})
	}

	return r
}
