# InternalModuleAuthApplicationServiceAuthPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aud** | Pointer to **string** |  | [optional] 
**Audience** | Pointer to **int32** | Audience 当前时间 | [optional] 
**Exp** | Pointer to **int32** |  | [optional] 
**ExpiresAt** | Pointer to **int32** | ExpiresAt  token 过期时间 | [optional] 
**Iat** | Pointer to **int32** |  | [optional] 
**Iss** | Pointer to **string** |  | [optional] 
**Jti** | Pointer to **string** |  | [optional] 
**Nbf** | Pointer to **int32** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 
**Uid** | Pointer to **string** | UID 用户id | [optional] 
**Username** | Pointer to **string** | Username 用户名 | [optional] 

## Methods

### NewInternalModuleAuthApplicationServiceAuthPayload

`func NewInternalModuleAuthApplicationServiceAuthPayload() *InternalModuleAuthApplicationServiceAuthPayload`

NewInternalModuleAuthApplicationServiceAuthPayload instantiates a new InternalModuleAuthApplicationServiceAuthPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalModuleAuthApplicationServiceAuthPayloadWithDefaults

`func NewInternalModuleAuthApplicationServiceAuthPayloadWithDefaults() *InternalModuleAuthApplicationServiceAuthPayload`

NewInternalModuleAuthApplicationServiceAuthPayloadWithDefaults instantiates a new InternalModuleAuthApplicationServiceAuthPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAud

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetAud() string`

GetAud returns the Aud field if non-nil, zero value otherwise.

### GetAudOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetAudOk() (*string, bool)`

GetAudOk returns a tuple with the Aud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAud

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetAud(v string)`

SetAud sets Aud field to given value.

### HasAud

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasAud() bool`

HasAud returns a boolean if a field has been set.

### GetAudience

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetAudience() int32`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetAudienceOk() (*int32, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetAudience(v int32)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetExp

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetExp() int32`

GetExp returns the Exp field if non-nil, zero value otherwise.

### GetExpOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetExpOk() (*int32, bool)`

GetExpOk returns a tuple with the Exp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExp

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetExp(v int32)`

SetExp sets Exp field to given value.

### HasExp

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasExp() bool`

HasExp returns a boolean if a field has been set.

### GetExpiresAt

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetIat

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetIat() int32`

GetIat returns the Iat field if non-nil, zero value otherwise.

### GetIatOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetIatOk() (*int32, bool)`

GetIatOk returns a tuple with the Iat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIat

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetIat(v int32)`

SetIat sets Iat field to given value.

### HasIat

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasIat() bool`

HasIat returns a boolean if a field has been set.

### GetIss

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasIss() bool`

HasIss returns a boolean if a field has been set.

### GetJti

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetJti() string`

GetJti returns the Jti field if non-nil, zero value otherwise.

### GetJtiOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetJtiOk() (*string, bool)`

GetJtiOk returns a tuple with the Jti field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJti

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetJti(v string)`

SetJti sets Jti field to given value.

### HasJti

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasJti() bool`

HasJti returns a boolean if a field has been set.

### GetNbf

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetNbf() int32`

GetNbf returns the Nbf field if non-nil, zero value otherwise.

### GetNbfOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetNbfOk() (*int32, bool)`

GetNbfOk returns a tuple with the Nbf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbf

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetNbf(v int32)`

SetNbf sets Nbf field to given value.

### HasNbf

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasNbf() bool`

HasNbf returns a boolean if a field has been set.

### GetSub

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasSub() bool`

HasSub returns a boolean if a field has been set.

### GetUid

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetUid() string`

GetUid returns the Uid field if non-nil, zero value otherwise.

### GetUidOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetUidOk() (*string, bool)`

GetUidOk returns a tuple with the Uid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUid

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetUid(v string)`

SetUid sets Uid field to given value.

### HasUid

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasUid() bool`

HasUid returns a boolean if a field has been set.

### GetUsername

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InternalModuleAuthApplicationServiceAuthPayload) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InternalModuleAuthApplicationServiceAuthPayload) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *InternalModuleAuthApplicationServiceAuthPayload) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


