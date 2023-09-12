/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsMonster struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Attributes *DungeonsandtrollsAttributes `json:"attributes,omitempty"`
	Effects []DungeonsandtrollsEffect `json:"effects,omitempty"`
	EquippedItems []DungeonsandtrollsItem `json:"equippedItems,omitempty"`
	Score float32 `json:"score,omitempty"`
	Icon string `json:"icon,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
	Faction string `json:"faction,omitempty"`
	OnDeath []DungeonsandtrollsDroppable `json:"onDeath,omitempty"`
}