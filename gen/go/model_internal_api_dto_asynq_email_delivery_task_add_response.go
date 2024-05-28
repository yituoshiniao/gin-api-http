/*
gin-http API

gin-http服务文档

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the InternalApiDtoAsynqEmailDeliveryTaskAddResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalApiDtoAsynqEmailDeliveryTaskAddResponse{}

// InternalApiDtoAsynqEmailDeliveryTaskAddResponse struct for InternalApiDtoAsynqEmailDeliveryTaskAddResponse
type InternalApiDtoAsynqEmailDeliveryTaskAddResponse struct {
	// code:  0 成功; 非0失败;
	Code int32 `json:"code"`
	// 数据data
	Data map[string]interface{} `json:"data"`
	// 错误消息
	Msg string `json:"msg"`
	// traceId
	TraceId string `json:"traceId"`
}

type _InternalApiDtoAsynqEmailDeliveryTaskAddResponse InternalApiDtoAsynqEmailDeliveryTaskAddResponse

// NewInternalApiDtoAsynqEmailDeliveryTaskAddResponse instantiates a new InternalApiDtoAsynqEmailDeliveryTaskAddResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalApiDtoAsynqEmailDeliveryTaskAddResponse(code int32, data map[string]interface{}, msg string, traceId string) *InternalApiDtoAsynqEmailDeliveryTaskAddResponse {
	this := InternalApiDtoAsynqEmailDeliveryTaskAddResponse{}
	this.Code = code
	this.Data = data
	this.Msg = msg
	this.TraceId = traceId
	return &this
}

// NewInternalApiDtoAsynqEmailDeliveryTaskAddResponseWithDefaults instantiates a new InternalApiDtoAsynqEmailDeliveryTaskAddResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalApiDtoAsynqEmailDeliveryTaskAddResponseWithDefaults() *InternalApiDtoAsynqEmailDeliveryTaskAddResponse {
	this := InternalApiDtoAsynqEmailDeliveryTaskAddResponse{}
	return &this
}

// GetCode returns the Code field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) SetCode(v int32) {
	o.Code = v
}

// GetData returns the Data field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) SetData(v map[string]interface{}) {
	o.Data = v
}

// GetMsg returns the Msg field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetMsg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Msg
}

// GetMsgOk returns a tuple with the Msg field value
// and a boolean to check if the value has been set.
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetMsgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Msg, true
}

// SetMsg sets field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) SetMsg(v string) {
	o.Msg = v
}

// GetTraceId returns the TraceId field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) SetTraceId(v string) {
	o.TraceId = v
}

func (o InternalApiDtoAsynqEmailDeliveryTaskAddResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalApiDtoAsynqEmailDeliveryTaskAddResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["data"] = o.Data
	toSerialize["msg"] = o.Msg
	toSerialize["traceId"] = o.TraceId
	return toSerialize, nil
}

func (o *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"data",
		"msg",
		"traceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varInternalApiDtoAsynqEmailDeliveryTaskAddResponse := _InternalApiDtoAsynqEmailDeliveryTaskAddResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varInternalApiDtoAsynqEmailDeliveryTaskAddResponse)

	if err != nil {
		return err
	}

	*o = InternalApiDtoAsynqEmailDeliveryTaskAddResponse(varInternalApiDtoAsynqEmailDeliveryTaskAddResponse)

	return err
}

type NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse struct {
	value *InternalApiDtoAsynqEmailDeliveryTaskAddResponse
	isSet bool
}

func (v NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse) Get() *InternalApiDtoAsynqEmailDeliveryTaskAddResponse {
	return v.value
}

func (v *NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse) Set(val *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse(val *InternalApiDtoAsynqEmailDeliveryTaskAddResponse) *NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse {
	return &NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse{value: val, isSet: true}
}

func (v NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalApiDtoAsynqEmailDeliveryTaskAddResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

