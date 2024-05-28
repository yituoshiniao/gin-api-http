# InternalApiDtoAppJwtTokenResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JwtPayload** | Pointer to [**InternalModuleAuthApplicationServiceAuthPayload**](InternalModuleAuthApplicationServiceAuthPayload.md) | token payload 信息 | [optional] 
**Token** | Pointer to **string** | Token jwt token | [optional] 

## Methods

### NewInternalApiDtoAppJwtTokenResponse

`func NewInternalApiDtoAppJwtTokenResponse() *InternalApiDtoAppJwtTokenResponse`

NewInternalApiDtoAppJwtTokenResponse instantiates a new InternalApiDtoAppJwtTokenResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApiDtoAppJwtTokenResponseWithDefaults

`func NewInternalApiDtoAppJwtTokenResponseWithDefaults() *InternalApiDtoAppJwtTokenResponse`

NewInternalApiDtoAppJwtTokenResponseWithDefaults instantiates a new InternalApiDtoAppJwtTokenResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJwtPayload

`func (o *InternalApiDtoAppJwtTokenResponse) GetJwtPayload() InternalModuleAuthApplicationServiceAuthPayload`

GetJwtPayload returns the JwtPayload field if non-nil, zero value otherwise.

### GetJwtPayloadOk

`func (o *InternalApiDtoAppJwtTokenResponse) GetJwtPayloadOk() (*InternalModuleAuthApplicationServiceAuthPayload, bool)`

GetJwtPayloadOk returns a tuple with the JwtPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtPayload

`func (o *InternalApiDtoAppJwtTokenResponse) SetJwtPayload(v InternalModuleAuthApplicationServiceAuthPayload)`

SetJwtPayload sets JwtPayload field to given value.

### HasJwtPayload

`func (o *InternalApiDtoAppJwtTokenResponse) HasJwtPayload() bool`

HasJwtPayload returns a boolean if a field has been set.

### GetToken

`func (o *InternalApiDtoAppJwtTokenResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *InternalApiDtoAppJwtTokenResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *InternalApiDtoAppJwtTokenResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *InternalApiDtoAppJwtTokenResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


