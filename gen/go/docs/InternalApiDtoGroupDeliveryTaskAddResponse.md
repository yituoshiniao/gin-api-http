# InternalApiDtoGroupDeliveryTaskAddResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | code:  0 成功; 非0失败; | 
**Data** | **map[string]interface{}** | 数据data | 
**Msg** | **string** | 错误消息 | 
**TraceId** | **string** | traceId | 

## Methods

### NewInternalApiDtoGroupDeliveryTaskAddResponse

`func NewInternalApiDtoGroupDeliveryTaskAddResponse(code int32, data map[string]interface{}, msg string, traceId string, ) *InternalApiDtoGroupDeliveryTaskAddResponse`

NewInternalApiDtoGroupDeliveryTaskAddResponse instantiates a new InternalApiDtoGroupDeliveryTaskAddResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalApiDtoGroupDeliveryTaskAddResponseWithDefaults

`func NewInternalApiDtoGroupDeliveryTaskAddResponseWithDefaults() *InternalApiDtoGroupDeliveryTaskAddResponse`

NewInternalApiDtoGroupDeliveryTaskAddResponseWithDefaults instantiates a new InternalApiDtoGroupDeliveryTaskAddResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetData

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetMsg

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetTraceId

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *InternalApiDtoGroupDeliveryTaskAddResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


