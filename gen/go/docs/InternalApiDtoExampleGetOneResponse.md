# InternalApiDtoExampleGetOneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | code:  0 成功; 非0失败; | 
**Data** | [**InternalApiDtoUserPortraitData**](InternalApiDtoUserPortraitData.md) |  | 
**Msg** | **string** | 错误消息 | 
**TraceId** | **string** | traceId | 

## Methods

### NewInternalApiDtoExampleGetOneResponse

`func NewInternalApiDtoExampleGetOneResponse(code int32, data InternalApiDtoUserPortraitData, msg string, traceId string, ) *InternalApiDtoExampleGetOneResponse`

NewInternalApiDtoExampleGetOneResponse instantiates a new InternalApiDtoExampleGetOneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApiDtoExampleGetOneResponseWithDefaults

`func NewInternalApiDtoExampleGetOneResponseWithDefaults() *InternalApiDtoExampleGetOneResponse`

NewInternalApiDtoExampleGetOneResponseWithDefaults instantiates a new InternalApiDtoExampleGetOneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InternalApiDtoExampleGetOneResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InternalApiDtoExampleGetOneResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InternalApiDtoExampleGetOneResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetData

`func (o *InternalApiDtoExampleGetOneResponse) GetData() InternalApiDtoUserPortraitData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternalApiDtoExampleGetOneResponse) GetDataOk() (*InternalApiDtoUserPortraitData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternalApiDtoExampleGetOneResponse) SetData(v InternalApiDtoUserPortraitData)`

SetData sets Data field to given value.


### GetMsg

`func (o *InternalApiDtoExampleGetOneResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *InternalApiDtoExampleGetOneResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *InternalApiDtoExampleGetOneResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetTraceId

`func (o *InternalApiDtoExampleGetOneResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *InternalApiDtoExampleGetOneResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *InternalApiDtoExampleGetOneResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


