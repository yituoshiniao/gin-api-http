# InternalApiDtoAppJwtTokenSwgResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | code:  0 成功; 非0失败; | 
**Data** | [**InternalApiDtoAppJwtTokenResponse**](InternalApiDtoAppJwtTokenResponse.md) |  | 
**Msg** | **string** | 错误消息 | 
**TraceId** | **string** | traceId | 

## Methods

### NewInternalApiDtoAppJwtTokenSwgResponse

`func NewInternalApiDtoAppJwtTokenSwgResponse(code int32, data InternalApiDtoAppJwtTokenResponse, msg string, traceId string, ) *InternalApiDtoAppJwtTokenSwgResponse`

NewInternalApiDtoAppJwtTokenSwgResponse instantiates a new InternalApiDtoAppJwtTokenSwgResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApiDtoAppJwtTokenSwgResponseWithDefaults

`func NewInternalApiDtoAppJwtTokenSwgResponseWithDefaults() *InternalApiDtoAppJwtTokenSwgResponse`

NewInternalApiDtoAppJwtTokenSwgResponseWithDefaults instantiates a new InternalApiDtoAppJwtTokenSwgResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InternalApiDtoAppJwtTokenSwgResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetData

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetData() InternalApiDtoAppJwtTokenResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetDataOk() (*InternalApiDtoAppJwtTokenResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternalApiDtoAppJwtTokenSwgResponse) SetData(v InternalApiDtoAppJwtTokenResponse)`

SetData sets Data field to given value.


### GetMsg

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *InternalApiDtoAppJwtTokenSwgResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetTraceId

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *InternalApiDtoAppJwtTokenSwgResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *InternalApiDtoAppJwtTokenSwgResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


