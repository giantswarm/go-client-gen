/*
 * Giant Swarm legacy API
 *
 * Caution: This is an incomplete description of some legacy API functions.
 *
 * API version: legacy
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package gsclientgen

type V4InfoResponse struct {
	General *V4InfoResponseGeneral `json:"general,omitempty"`

	Workers *V4InfoResponseWorkers `json:"workers,omitempty"`
}
