# InternalApiDtoUserScoreFindResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | code:  0 成功; 非0失败; | 
**Data** | [**InternalModuleMockv2ApplicationEntityUserScore**](InternalModuleMockv2ApplicationEntityUserScore.md) |  | 
**Msg** | **string** | 错误消息 | 
**TraceId** | **string** | traceId | 

## Methods

### NewInternalApiDtoUserScoreFindResponse

`func NewInternalApiDtoUserScoreFindResponse(code int32, data InternalModuleMockv2ApplicationEntityUserScore, msg string, traceId string, ) *InternalApiDtoUserScoreFindResponse`

NewInternalApiDtoUserScoreFindResponse instantiates a new InternalApiDtoUserScoreFindResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApiDtoUserScoreFindResponseWithDefaults

`func NewInternalApiDtoUserScoreFindResponseWithDefaults() *InternalApiDtoUserScoreFindResponse`

NewInternalApiDtoUserScoreFindResponseWithDefaults instantiates a new InternalApiDtoUserScoreFindResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InternalApiDtoUserScoreFindResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InternalApiDtoUserScoreFindResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InternalApiDtoUserScoreFindResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetData

`func (o *InternalApiDtoUserScoreFindResponse) GetData() InternalModuleMockv2ApplicationEntityUserScore`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternalApiDtoUserScoreFindResponse) GetDataOk() (*InternalModuleMockv2ApplicationEntityUserScore, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternalApiDtoUserScoreFindResponse) SetData(v InternalModuleMockv2ApplicationEntityUserScore)`

SetData sets Data field to given value.


### GetMsg

`func (o *InternalApiDtoUserScoreFindResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *InternalApiDtoUserScoreFindResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *InternalApiDtoUserScoreFindResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetTraceId

`func (o *InternalApiDtoUserScoreFindResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *InternalApiDtoUserScoreFindResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *InternalApiDtoUserScoreFindResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


