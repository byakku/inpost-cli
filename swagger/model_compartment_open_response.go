/*
 * Inpost Mobile
 *
 * Extracted from retrofit
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type CompartmentOpenResponse struct {
	ActionTime int32 `json:"actionTime,omitempty"`
	Compartment *Compartment `json:"compartment,omitempty"`
	ConfirmActionTime int32 `json:"confirmActionTime,omitempty"`
	OpenCompartmentWaitingTime int32 `json:"openCompartmentWaitingTime,omitempty"`
}