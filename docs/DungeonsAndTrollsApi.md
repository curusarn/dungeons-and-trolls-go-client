# {{classname}}

All URIs are relative to *https://dt.garage-trip.cz/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DungeonsAndTrollsAssignSkillPoints**](DungeonsAndTrollsApi.md#DungeonsAndTrollsAssignSkillPoints) | **Post** /v1/assign-skill-points | Send multiple commands to the Character bound to the logged user. The order of execution is defined in the message.
[**DungeonsAndTrollsBuy**](DungeonsAndTrollsApi.md#DungeonsAndTrollsBuy) | **Post** /v1/buy | Buy Items identified by the provided ID for the Character bound to the logged user.
[**DungeonsAndTrollsCommands**](DungeonsAndTrollsApi.md#DungeonsAndTrollsCommands) | **Post** /v1/commands | Send multiple commands to the Character bound to the logged user. The order of execution is defined in the message.
[**DungeonsAndTrollsGame**](DungeonsAndTrollsApi.md#DungeonsAndTrollsGame) | **Get** /v1/game | Sends all info about the game.
[**DungeonsAndTrollsGameLevel**](DungeonsAndTrollsApi.md#DungeonsAndTrollsGameLevel) | **Get** /v1/game/{level} | Sends all info about the game level.
[**DungeonsAndTrollsLevels**](DungeonsAndTrollsApi.md#DungeonsAndTrollsLevels) | **Get** /v1/levels | Sends info about
[**DungeonsAndTrollsMonstersCommands**](DungeonsAndTrollsApi.md#DungeonsAndTrollsMonstersCommands) | **Post** /v1/monsters-commands | Control monsters. Admin only.
[**DungeonsAndTrollsMove**](DungeonsAndTrollsApi.md#DungeonsAndTrollsMove) | **Post** /v1/move | Assign skill point to the attribute for the Character bound to the logged user.
[**DungeonsAndTrollsPickUp**](DungeonsAndTrollsApi.md#DungeonsAndTrollsPickUp) | **Post** /v1/pick-up | Equip the Item from the ground identified by the provided ID for the Character bound to the logged user (unused).
[**DungeonsAndTrollsPlayers**](DungeonsAndTrollsApi.md#DungeonsAndTrollsPlayers) | **Get** /v1/players | Sends all info about all players.
[**DungeonsAndTrollsRegister**](DungeonsAndTrollsApi.md#DungeonsAndTrollsRegister) | **Post** /v1/register | Register provided User to the Game and create a character.
[**DungeonsAndTrollsRespawn**](DungeonsAndTrollsApi.md#DungeonsAndTrollsRespawn) | **Post** /v1/respawn | Respawn the Character bound to the logged user.
[**DungeonsAndTrollsSkill**](DungeonsAndTrollsApi.md#DungeonsAndTrollsSkill) | **Post** /v1/skill | Use a skill (provided by an item) by the Character bound to the logged user.
[**DungeonsAndTrollsYell**](DungeonsAndTrollsApi.md#DungeonsAndTrollsYell) | **Post** /v1/yell | The Character bound to the logged user yells a messages (visible for everyone).

# **DungeonsAndTrollsAssignSkillPoints**
> interface{} DungeonsAndTrollsAssignSkillPoints(ctx, body, optional)
Send multiple commands to the Character bound to the logged user. The order of execution is defined in the message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsAttributes**](DungeonsandtrollsAttributes.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsAssignSkillPointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsAssignSkillPointsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsBuy**
> interface{} DungeonsAndTrollsBuy(ctx, body, optional)
Buy Items identified by the provided ID for the Character bound to the logged user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsIdentifiers**](DungeonsandtrollsIdentifiers.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsBuyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsBuyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsCommands**
> interface{} DungeonsAndTrollsCommands(ctx, body, optional)
Send multiple commands to the Character bound to the logged user. The order of execution is defined in the message.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsCommandsBatch**](DungeonsandtrollsCommandsBatch.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsCommandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsCommandsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsGame**
> DungeonsandtrollsGameState DungeonsAndTrollsGame(ctx, optional)
Sends all info about the game.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsGameOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsGameOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocking** | **optional.Bool**| default false | 
 **items** | **optional.Bool**| default true | 
 **fogOfWar** | **optional.Bool**| default false | 

### Return type

[**DungeonsandtrollsGameState**](dungeonsandtrollsGameState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsGameLevel**
> DungeonsandtrollsGameState DungeonsAndTrollsGameLevel(ctx, level, optional)
Sends all info about the game level.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **level** | **int32**|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsGameLevelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsGameLevelOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.Bool**| default false | 
 **items** | **optional.Bool**| default true | 
 **fogOfWar** | **optional.Bool**| default false | 

### Return type

[**DungeonsandtrollsGameState**](dungeonsandtrollsGameState.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsLevels**
> DungeonsandtrollsAvailableLevels DungeonsAndTrollsLevels(ctx, optional)
Sends info about

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsLevelsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsLevelsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocking** | **optional.Bool**| default false | 

### Return type

[**DungeonsandtrollsAvailableLevels**](dungeonsandtrollsAvailableLevels.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsMonstersCommands**
> interface{} DungeonsAndTrollsMonstersCommands(ctx, body, optional)
Control monsters. Admin only.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsCommandsForMonsters**](DungeonsandtrollsCommandsForMonsters.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsMonstersCommandsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsMonstersCommandsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsMove**
> interface{} DungeonsAndTrollsMove(ctx, body, optional)
Assign skill point to the attribute for the Character bound to the logged user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsPosition**](DungeonsandtrollsPosition.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsMoveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsMoveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsPickUp**
> interface{} DungeonsAndTrollsPickUp(ctx, body, optional)
Equip the Item from the ground identified by the provided ID for the Character bound to the logged user (unused).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsIdentifier**](DungeonsandtrollsIdentifier.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsPickUpOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsPickUpOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsPlayers**
> DungeonsandtrollsPlayersInfo DungeonsAndTrollsPlayers(ctx, optional)
Sends all info about all players.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsPlayersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsPlayersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **blocking** | **optional.Bool**| default false | 

### Return type

[**DungeonsandtrollsPlayersInfo**](dungeonsandtrollsPlayersInfo.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsRegister**
> DungeonsandtrollsRegistration DungeonsAndTrollsRegister(ctx, body)
Register provided User to the Game and create a character.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsUser**](DungeonsandtrollsUser.md)|  | 

### Return type

[**DungeonsandtrollsRegistration**](dungeonsandtrollsRegistration.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsRespawn**
> interface{} DungeonsAndTrollsRespawn(ctx, body, optional)
Respawn the Character bound to the logged user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**interface{}**](interface{}.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsRespawnOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsRespawnOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsSkill**
> interface{} DungeonsAndTrollsSkill(ctx, body, optional)
Use a skill (provided by an item) by the Character bound to the logged user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsSkillUse**](DungeonsandtrollsSkillUse.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsSkillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsSkillOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DungeonsAndTrollsYell**
> interface{} DungeonsAndTrollsYell(ctx, body, optional)
The Character bound to the logged user yells a messages (visible for everyone).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**DungeonsandtrollsMessage**](DungeonsandtrollsMessage.md)|  | 
 **optional** | ***DungeonsAndTrollsApiDungeonsAndTrollsYellOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DungeonsAndTrollsApiDungeonsAndTrollsYellOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **blocking** | **optional.**| default true | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

