/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MetadataValueOneOf2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataValueOneOf2{}

// MetadataValueOneOf2 struct for MetadataValueOneOf2
type MetadataValueOneOf2 struct {
	StringValue *string `json:"string_value,omitempty"`
}

// NewMetadataValueOneOf2 instantiates a new MetadataValueOneOf2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataValueOneOf2() *MetadataValueOneOf2 {
	this := MetadataValueOneOf2{}
	return &this
}

// NewMetadataValueOneOf2WithDefaults instantiates a new MetadataValueOneOf2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataValueOneOf2WithDefaults() *MetadataValueOneOf2 {
	this := MetadataValueOneOf2{}
	return &this
}

// GetStringValue returns the StringValue field value if set, zero value otherwise.
func (o *MetadataValueOneOf2) GetStringValue() string {
	if o == nil || IsNil(o.StringValue) {
		var ret string
		return ret
	}
	return *o.StringValue
}

// GetStringValueOk returns a tuple with the StringValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetadataValueOneOf2) GetStringValueOk() (*string, bool) {
	if o == nil || IsNil(o.StringValue) {
		return nil, false
	}
	return o.StringValue, true
}

// HasStringValue returns a boolean if a field has been set.
func (o *MetadataValueOneOf2) HasStringValue() bool {
	if o != nil && !IsNil(o.StringValue) {
		return true
	}

	return false
}

// SetStringValue gets a reference to the given string and assigns it to the StringValue field.
func (o *MetadataValueOneOf2) SetStringValue(v string) {
	o.StringValue = &v
}

func (o MetadataValueOneOf2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataValueOneOf2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StringValue) {
		toSerialize["string_value"] = o.StringValue
	}
	return toSerialize, nil
}

type NullableMetadataValueOneOf2 struct {
	value *MetadataValueOneOf2
	isSet bool
}

func (v NullableMetadataValueOneOf2) Get() *MetadataValueOneOf2 {
	return v.value
}

func (v *NullableMetadataValueOneOf2) Set(val *MetadataValueOneOf2) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataValueOneOf2) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataValueOneOf2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataValueOneOf2(val *MetadataValueOneOf2) *NullableMetadataValueOneOf2 {
	return &NullableMetadataValueOneOf2{value: val, isSet: true}
}

func (v NullableMetadataValueOneOf2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataValueOneOf2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}