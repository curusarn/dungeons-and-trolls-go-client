/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsEffect struct {
	DamageAmount float32 `json:"damageAmount,omitempty"`
	DamageType *DungeonsandtrollsDamageType `json:"damageType,omitempty"`
	Effects *DungeonsandtrollsAttributes `json:"effects,omitempty"`
	Duration int32 `json:"duration,omitempty"`
	CasterId string `json:"CasterId,omitempty"`
}
