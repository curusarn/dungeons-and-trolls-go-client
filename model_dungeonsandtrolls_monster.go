/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsMonster struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`
	Items []DungeonsandtrollsSimpleItem `json:"items,omitempty"`
	Effects []DungeonsandtrollsEffect `json:"effects,omitempty"`
	LifePercentage float32 `json:"lifePercentage,omitempty"`
	Faction string `json:"faction,omitempty"`
	Attributes *DungeonsandtrollsAttributes `json:"attributes,omitempty"`
	EquippedItems []DungeonsandtrollsItem `json:"equippedItems,omitempty"`
	Score float32 `json:"score,omitempty"`
	Algorithm string `json:"algorithm,omitempty"`
	OnDeath []DungeonsandtrollsDroppable `json:"onDeath,omitempty"`
	MaxAttributes *DungeonsandtrollsAttributes `json:"maxAttributes,omitempty"`
	LastDamageTaken int32 `json:"lastDamageTaken,omitempty"`
	Stun *DungeonsandtrollsStun `json:"stun,omitempty"`
}
