/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsCharacter struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Attributes *DungeonsandtrollsAttributes `json:"attributes,omitempty"`
	Money int32 `json:"money,omitempty"`
	Equip []DungeonsandtrollsItem `json:"equip,omitempty"`
	Score float32 `json:"score,omitempty"`
	SkillPoints float32 `json:"skillPoints,omitempty"`
	Effects []DungeonsandtrollsEffect `json:"effects,omitempty"`
}
