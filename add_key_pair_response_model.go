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

type AddKeyPairResponseModel struct {
	Id string `json:"id,omitempty"`

	Description string `json:"description,omitempty"`

	TtlHours int32 `json:"ttl_hours,omitempty"`

	CreateDate string `json:"create_date,omitempty"`

	CertificateAuthorityData string `json:"certificate_authority_data,omitempty"`

	ClientKeyData string `json:"client_key_data,omitempty"`

	ClientCertificateData string `json:"client_certificate_data,omitempty"`
}
