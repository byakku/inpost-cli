/*
 * Inpost Mobile
 *
 * Extracted from retrofit
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AuthenticateRequest struct {
	PhoneOS string `json:"phoneOS,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
}
