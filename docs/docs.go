// Documentation of Posts CRUD API Gateway.
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package docs

import "github.com/iman_task/api-gateway/api/models"

// Empty Response
// swagger:response emptyResponse
type swagEmptyResponse struct{}

// Error BadRequest
// swagger:response badRequest
type swagErrBadRequest struct {
	// in:body
	Body models.ValidationError
}

// Error Not Found
// swagger:response notFoundResponse
type swagErrNotFound struct {
	// in:body
	Body models.ValidationError
}

// Error Interval Server
// swagger:response internalService
type swagErrInternal struct {
	// in:body
	Body models.InternalServerError
}

// QueryParams of endpoints
//swagger:parameters postsList
type QueryParams struct {
	// limit
	//
	//required: false
	//in:query
	Limit int `json:"limit"`

	// page number
	//
	//required: false
	//in:query
	Page int `json:"page"`

	// search text
	//
	//required: false
	//in:query
	Search string `json:"search"`
}
