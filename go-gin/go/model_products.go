/*
 * EC
 *
 * EC API
 *
 * API version: 2023.1.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Products struct {

	NextPageToken string `json:"nextPageToken,omitempty"`

	Items []Product `json:"items,omitempty"`
}
