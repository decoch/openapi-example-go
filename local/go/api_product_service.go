/*
 * EC
 *
 * EC API
 *
 * API version: 2023.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"net/http"
	"errors"
)

// ProductApiService is a service that implements the logic for the ProductApiServicer
// This service should implement the business logic for every endpoint for the ProductApi API.
// Include any external packages or services that will be required by this service.
type ProductApiService struct {
}

// NewProductApiService creates a default api service
func NewProductApiService() ProductApiServicer {
	return &ProductApiService{}
}

// DeleteProductById - Cancel a specified product
func (s *ProductApiService) DeleteProductById(ctx context.Context, productId int64) (ImplResponse, error) {
	// TODO - update DeleteProductById with the required logic for this service method.
	// Add api_product_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	//return Response(201, nil),nil

	//TODO: Uncomment the next line to return response Response(0, ModelError{}) or use other options such as http.Ok ...
	//return Response(0, ModelError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteProductById method not implemented")
}

// GetProductById - Get a specified product
func (s *ProductApiService) GetProductById(ctx context.Context, productId int64) (ImplResponse, error) {
	// TODO - update GetProductById with the required logic for this service method.
	// Add api_product_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Product{}) or use other options such as http.Ok ...
	//return Response(200, Product{}), nil

	//TODO: Uncomment the next line to return response Response(0, ModelError{}) or use other options such as http.Ok ...
	//return Response(0, ModelError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetProductById method not implemented")
}

// GetProducts - Get all products
func (s *ProductApiService) GetProducts(ctx context.Context, pageSize int32, pageToken string) (ImplResponse, error) {
	// TODO - update GetProducts with the required logic for this service method.
	// Add api_product_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	//TODO: Uncomment the next line to return response Response(200, Products{}) or use other options such as http.Ok ...
	//return Response(200, Products{}), nil

	//TODO: Uncomment the next line to return response Response(0, ModelError{}) or use other options such as http.Ok ...
	//return Response(0, ModelError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetProducts method not implemented")
}