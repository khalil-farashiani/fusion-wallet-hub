package handler

import "github.com/gorilla/mux"

func (h *Handler) SetBalanceRoute(router *mux.Router) {
	router.HandleFunc("/balances/users/{userID}", h.GetUserBalance).Methods("GET")
}
