/*
 * Giant Swarm legacy API
 *
 * Caution: This is an incomplete description of some legacy API functions.
 *
 * OpenAPI spec version: legacy
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package gsclientgen

type V4AddKeyPairBody struct {
	Description string `json:"description"`

	TtlHours int32 `json:"ttl_hours,omitempty"`

	CnPrefix string `json:"cn_prefix,omitempty"`

	CertificateOrganizations string `json:"certificate_organizations,omitempty"`
}
