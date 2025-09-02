package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"stores-lite/internal/service"
)

type rest struct{ svc *service.Service }

func RegisterREST(r *chi.Mux, svc *service.Service) {
	h := &rest{svc: svc}
	r.Route("/api", func(api chi.Router) {
		api.Post("/products", h.createProduct)
		api.Get("/healthz", func(w http.ResponseWriter, r *http.Request) { w.Write([]byte("ok")) })
		api.Post("/checkout", h.checkout)
	})
}

func (h *rest) createProduct(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name       string `json:"name"`
		PriceCents int64  `json:"priceCents"`
		Stock      int64  `json:"stock"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest); return
	}
	p, err := h.svc.CreateProduct(r.Context(), req.Name, req.PriceCents, req.Stock)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest); return
	}
	json.NewEncoder(w).Encode(p)
}

func (h *rest) checkout(w http.ResponseWriter, r *http.Request) {
	var req struct {
		ProductID int64 `json:"productId"`
		Qty       int64 `json:"qty"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest); return
	}
	o, err := h.svc.Checkout(context.Background(), req.ProductID, req.Qty)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest); return
	}
	json.NewEncoder(w).Encode(o)
}

func atoi64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
