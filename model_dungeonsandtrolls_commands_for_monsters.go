/*
 * Dungeons and Trolls
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.10.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type DungeonsandtrollsCommandsForMonsters struct {
	// Monster ID with corresponding batch of commands.
	Commands map[string]DungeonsandtrollsCommandsBatch `json:"commands,omitempty"`
}
