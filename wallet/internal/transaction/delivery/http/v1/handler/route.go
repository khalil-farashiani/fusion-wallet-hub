package handler

import "github.com/gorilla/mux"

func (h *Handler) SetTransactionRoute(r *mux.Router) {
	r.HandleFunc("/transactions", h.CreateTransaction).Methods("POST")
	r.HandleFunc("/transactions", h.GetAllTransaction).Methods("GET")
	r.HandleFunc("/transactions/users/{userID}", h.GetUserTransaction).Methods("GET")
}
