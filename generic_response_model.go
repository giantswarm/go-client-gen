/* 
 * Giant Swarm legacy API
 *
 * Caution: This is an incomplete description of some legacy API functions.
 *
 * OpenAPI spec version: legacy
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package goclient

type GenericResponseModel struct {

	StatusCode int32 `json:"status_code,omitempty"`

	StatusText string `json:"status_text,omitempty"`
}