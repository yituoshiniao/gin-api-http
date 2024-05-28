/*
gin-http API

gin-http服务文档

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the InternalModuleMockv2ApplicationEntityUserScore type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InternalModuleMockv2ApplicationEntityUserScore{}

// InternalModuleMockv2ApplicationEntityUserScore struct for InternalModuleMockv2ApplicationEntityUserScore
type InternalModuleMockv2ApplicationEntityUserScore struct {
	CreateTime *int32 `json:"create_time,omitempty"`
	Id *int32 `json:"id,omitempty"`
	Score *int32 `json:"score,omitempty"`
	ScoreResult *int32 `json:"score_result,omitempty"`
	UpdateTime *int32 `json:"update_time,omitempty"`
	UserId *string `json:"user_id,omitempty"`
	ZeroTimestamp *int32 `json:"zero_timestamp,omitempty"`
}

// NewInternalModuleMockv2ApplicationEntityUserScore instantiates a new InternalModuleMockv2ApplicationEntityUserScore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInternalModuleMockv2ApplicationEntityUserScore() *InternalModuleMockv2ApplicationEntityUserScore {
	this := InternalModuleMockv2ApplicationEntityUserScore{}
	return &this
}

// NewInternalModuleMockv2ApplicationEntityUserScoreWithDefaults instantiates a new InternalModuleMockv2ApplicationEntityUserScore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInternalModuleMockv2ApplicationEntityUserScoreWithDefaults() *InternalModuleMockv2ApplicationEntityUserScore {
	this := InternalModuleMockv2ApplicationEntityUserScore{}
	return &this
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetCreateTime() int32 {
	if o == nil || IsNil(o.CreateTime) {
		var ret int32
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetCreateTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.CreateTime) {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) HasCreateTime() bool {
	if o != nil && !IsNil(o.CreateTime) {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int32 and assigns it to the CreateTime field.
func (o *InternalModuleMockv2ApplicationEntityUserScore) SetCreateTime(v int32) {
	o.CreateTime = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *InternalModuleMockv2ApplicationEntityUserScore) SetId(v int32) {
	o.Id = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *InternalModuleMockv2ApplicationEntityUserScore) SetScore(v int32) {
	o.Score = &v
}

// GetScoreResult returns the ScoreResult field value if set, zero value otherwise.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetScoreResult() int32 {
	if o == nil || IsNil(o.ScoreResult) {
		var ret int32
		return ret
	}
	return *o.ScoreResult
}

// GetScoreResultOk returns a tuple with the ScoreResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetScoreResultOk() (*int32, bool) {
	if o == nil || IsNil(o.ScoreResult) {
		return nil, false
	}
	return o.ScoreResult, true
}

// HasScoreResult returns a boolean if a field has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) HasScoreResult() bool {
	if o != nil && !IsNil(o.ScoreResult) {
		return true
	}

	return false
}

// SetScoreResult gets a reference to the given int32 and assigns it to the ScoreResult field.
func (o *InternalModuleMockv2ApplicationEntityUserScore) SetScoreResult(v int32) {
	o.ScoreResult = &v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetUpdateTime() int32 {
	if o == nil || IsNil(o.UpdateTime) {
		var ret int32
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetUpdateTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.UpdateTime) {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) HasUpdateTime() bool {
	if o != nil && !IsNil(o.UpdateTime) {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given int32 and assigns it to the UpdateTime field.
func (o *InternalModuleMockv2ApplicationEntityUserScore) SetUpdateTime(v int32) {
	o.UpdateTime = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetUserId() string {
	if o == nil || IsNil(o.UserId) {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *InternalModuleMockv2ApplicationEntityUserScore) SetUserId(v string) {
	o.UserId = &v
}

// GetZeroTimestamp returns the ZeroTimestamp field value if set, zero value otherwise.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetZeroTimestamp() int32 {
	if o == nil || IsNil(o.ZeroTimestamp) {
		var ret int32
		return ret
	}
	return *o.ZeroTimestamp
}

// GetZeroTimestampOk returns a tuple with the ZeroTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) GetZeroTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.ZeroTimestamp) {
		return nil, false
	}
	return o.ZeroTimestamp, true
}

// HasZeroTimestamp returns a boolean if a field has been set.
func (o *InternalModuleMockv2ApplicationEntityUserScore) HasZeroTimestamp() bool {
	if o != nil && !IsNil(o.ZeroTimestamp) {
		return true
	}

	return false
}

// SetZeroTimestamp gets a reference to the given int32 and assigns it to the ZeroTimestamp field.
func (o *InternalModuleMockv2ApplicationEntityUserScore) SetZeroTimestamp(v int32) {
	o.ZeroTimestamp = &v
}

func (o InternalModuleMockv2ApplicationEntityUserScore) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InternalModuleMockv2ApplicationEntityUserScore) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreateTime) {
		toSerialize["create_time"] = o.CreateTime
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.ScoreResult) {
		toSerialize["score_result"] = o.ScoreResult
	}
	if !IsNil(o.UpdateTime) {
		toSerialize["update_time"] = o.UpdateTime
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.ZeroTimestamp) {
		toSerialize["zero_timestamp"] = o.ZeroTimestamp
	}
	return toSerialize, nil
}

type NullableInternalModuleMockv2ApplicationEntityUserScore struct {
	value *InternalModuleMockv2ApplicationEntityUserScore
	isSet bool
}

func (v NullableInternalModuleMockv2ApplicationEntityUserScore) Get() *InternalModuleMockv2ApplicationEntityUserScore {
	return v.value
}

func (v *NullableInternalModuleMockv2ApplicationEntityUserScore) Set(val *InternalModuleMockv2ApplicationEntityUserScore) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalModuleMockv2ApplicationEntityUserScore) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalModuleMockv2ApplicationEntityUserScore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalModuleMockv2ApplicationEntityUserScore(val *InternalModuleMockv2ApplicationEntityUserScore) *NullableInternalModuleMockv2ApplicationEntityUserScore {
	return &NullableInternalModuleMockv2ApplicationEntityUserScore{value: val, isSet: true}
}

func (v NullableInternalModuleMockv2ApplicationEntityUserScore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalModuleMockv2ApplicationEntityUserScore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


