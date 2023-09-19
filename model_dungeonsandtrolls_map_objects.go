/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsMapObjects struct {
	Position *DungeonsandtrollsCoordinates `json:"position,omitempty"`
	Monsters []DungeonsandtrollsMonster `json:"monsters,omitempty"`
	Players []DungeonsandtrollsCharacter `json:"players,omitempty"`
	IsStairs bool `json:"isStairs,omitempty"`
	Portal *DungeonsandtrollsWaypoint `json:"portal,omitempty"`
	Decorations []DungeonsandtrollsDecoration `json:"decorations,omitempty"`
	Effects []DungeonsandtrollsEffect `json:"effects,omitempty"`
	Items []DungeonsandtrollsItem `json:"items,omitempty"`
	IsFree bool `json:"isFree,omitempty"`
	IsWall bool `json:"isWall,omitempty"`
	IsDoor bool `json:"isDoor,omitempty"`
	IsSpawn bool `json:"isSpawn,omitempty"`
}
