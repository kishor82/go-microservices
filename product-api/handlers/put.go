package handlers

import (
	"net/http"

	"github.com/kishor82/go-microservices/product-api/data"
)

// swagger:route PUT /products products updateProduct
// Update a products details
//
// responses:
//  201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update products
func (p *Products) Update(rw http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Debug("updating record", "id", prod.ID)

	err := p.productDB.UpdateProduct(&prod)
	if err == data.ErrProductNotFound {
		p.l.Error("product not found", "error", err)

		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{
			Message: "Product not found in database",
		}, rw)
	}
	// write the no content success header
	rw.WriteHeader(http.StatusNoContent)
}
