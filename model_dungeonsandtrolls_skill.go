/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsSkill struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Target *SkillTarget `json:"target,omitempty"`
	Effects []DungeonsandtrollsEffect `json:"effects,omitempty"`
	Cost *DungeonsandtrollsAttributes `json:"cost,omitempty"`
	Range_ *DungeonsandtrollsAttributes `json:"range,omitempty"`
	Radius *DungeonsandtrollsAttributes `json:"radius,omitempty"`
	Duration *DungeonsandtrollsAttributes `json:"duration,omitempty"`
	DamageAmount *DungeonsandtrollsAttributes `json:"damageAmount,omitempty"`
	DamageType *DungeonsandtrollsDamageType `json:"damageType,omitempty"`
	CasterEffects *DungeonsandtrollsSkillEffect `json:"casterEffects,omitempty"`
	TargetEffects *DungeonsandtrollsSkillEffect `json:"targetEffects,omitempty"`
}