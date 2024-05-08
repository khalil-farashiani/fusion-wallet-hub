package handler

import "github.com/gorilla/mux"

func (h *Handler) SetDiscountRoute(router *mux.Router) {
	router.HandleFunc("/redeem", h.CreateRedeem).Methods("POST")
	router.HandleFunc("/register-redeem", h.RegisterRedeem).Methods("GET")
	router.HandleFunc("/redeem/reports", h.GetRedeemReport).Methods("GET")
}
