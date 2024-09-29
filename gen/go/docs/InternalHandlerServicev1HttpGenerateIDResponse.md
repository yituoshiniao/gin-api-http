# InternalHandlerServicev1HttpGenerateIDResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **int32** | code:  0 成功; 非0失败; | 
**Data** | **[]string** | 生成id数组 | 
**Msg** | **string** | 错误消息 | 
**TraceId** | **string** | traceId | 

## Methods

### NewInternalHandlerServicev1HttpGenerateIDResponse

`func NewInternalHandlerServicev1HttpGenerateIDResponse(code int32, data []string, msg string, traceId string, ) *InternalHandlerServicev1HttpGenerateIDResponse`

NewInternalHandlerServicev1HttpGenerateIDResponse instantiates a new InternalHandlerServicev1HttpGenerateIDResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalHandlerServicev1HttpGenerateIDResponseWithDefaults

`func NewInternalHandlerServicev1HttpGenerateIDResponseWithDefaults() *InternalHandlerServicev1HttpGenerateIDResponse`

NewInternalHandlerServicev1HttpGenerateIDResponseWithDefaults instantiates a new InternalHandlerServicev1HttpGenerateIDResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) SetCode(v int32)`

SetCode sets Code field to given value.


### GetData

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) SetData(v []string)`

SetData sets Data field to given value.


### GetMsg

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetTraceId

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *InternalHandlerServicev1HttpGenerateIDResponse) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


