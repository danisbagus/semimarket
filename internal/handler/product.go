package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/danisbagus/semimarket-product/internal/core/port"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Service port.IProducService
}

func (rc ProductHandler) GetProductList(w http.ResponseWriter, r *http.Request) {
	dataList, err := rc.Service.GetAll()
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, dataList)
}

func (rc ProductHandler) GetProductDetail(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	ProductID, _ := strconv.Atoi(vars["product_id"])

	data, err := rc.Service.GetDetail(int64(ProductID))
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}
	writeResponse(w, http.StatusOK, data)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}