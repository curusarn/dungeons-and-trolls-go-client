/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsEvent struct {
	Message string `json:"message,omitempty"`
	Type_ *DungeonsandtrollsEventType `json:"type,omitempty"`
	Coordinates *DungeonsandtrollsCoordinates `json:"coordinates,omitempty"`
	Damage float32 `json:"damage,omitempty"`
	PlayerId string `json:"playerId,omitempty"`
	SkillName string `json:"skillName,omitempty"`
	Radius float32 `json:"radius,omitempty"`
	Target *DungeonsandtrollsCoordinates `json:"target,omitempty"`
	Skill *DungeonsandtrollsSkill `json:"skill,omitempty"`
}
