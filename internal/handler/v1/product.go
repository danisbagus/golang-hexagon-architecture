package v1

import (
	"encoding/json"
	"net/http"

	"github.com/danisbagus/golang-hexagon-architecture/internal/core/port"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Service port.IProductService
}

func (rc ProductHandler) GetProductListV1(w http.ResponseWriter, r *http.Request) {
	data, _ := rc.Service.GetList()
	writeResponse(w, http.StatusOK, data)
}

func (rc ProductHandler) GetProductDetailV1(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ProductID := vars["product_id"]
	data, _ := rc.Service.GetDetail(ProductID)
	writeResponse(w, http.StatusOK, data)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
