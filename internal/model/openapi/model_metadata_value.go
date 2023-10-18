/*
Model Registry REST API

REST API for Model Registry to create and manage ML model metadata

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MetadataValue - A value in properties.
type MetadataValue struct {
	MetadataValueOneOf  *MetadataValueOneOf
	MetadataValueOneOf1 *MetadataValueOneOf1
	MetadataValueOneOf2 *MetadataValueOneOf2
	MetadataValueOneOf3 *MetadataValueOneOf3
	MetadataValueOneOf4 *MetadataValueOneOf4
	MetadataValueOneOf5 *MetadataValueOneOf5
}

// MetadataValueOneOfAsMetadataValue is a convenience function that returns MetadataValueOneOf wrapped in MetadataValue
func MetadataValueOneOfAsMetadataValue(v *MetadataValueOneOf) MetadataValue {
	return MetadataValue{
		MetadataValueOneOf: v,
	}
}

// MetadataValueOneOf1AsMetadataValue is a convenience function that returns MetadataValueOneOf1 wrapped in MetadataValue
func MetadataValueOneOf1AsMetadataValue(v *MetadataValueOneOf1) MetadataValue {
	return MetadataValue{
		MetadataValueOneOf1: v,
	}
}

// MetadataValueOneOf2AsMetadataValue is a convenience function that returns MetadataValueOneOf2 wrapped in MetadataValue
func MetadataValueOneOf2AsMetadataValue(v *MetadataValueOneOf2) MetadataValue {
	return MetadataValue{
		MetadataValueOneOf2: v,
	}
}

// MetadataValueOneOf3AsMetadataValue is a convenience function that returns MetadataValueOneOf3 wrapped in MetadataValue
func MetadataValueOneOf3AsMetadataValue(v *MetadataValueOneOf3) MetadataValue {
	return MetadataValue{
		MetadataValueOneOf3: v,
	}
}

// MetadataValueOneOf4AsMetadataValue is a convenience function that returns MetadataValueOneOf4 wrapped in MetadataValue
func MetadataValueOneOf4AsMetadataValue(v *MetadataValueOneOf4) MetadataValue {
	return MetadataValue{
		MetadataValueOneOf4: v,
	}
}

// MetadataValueOneOf5AsMetadataValue is a convenience function that returns MetadataValueOneOf5 wrapped in MetadataValue
func MetadataValueOneOf5AsMetadataValue(v *MetadataValueOneOf5) MetadataValue {
	return MetadataValue{
		MetadataValueOneOf5: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *MetadataValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetadataValueOneOf
	err = json.Unmarshal(data, &dst.MetadataValueOneOf)
	if err == nil {
		jsonMetadataValueOneOf, _ := json.Marshal(dst.MetadataValueOneOf)
		if string(jsonMetadataValueOneOf) == "{}" { // empty struct
			dst.MetadataValueOneOf = nil
		} else {
			match++
		}
	} else {
		dst.MetadataValueOneOf = nil
	}

	// try to unmarshal data into MetadataValueOneOf1
	err = json.Unmarshal(data, &dst.MetadataValueOneOf1)
	if err == nil {
		jsonMetadataValueOneOf1, _ := json.Marshal(dst.MetadataValueOneOf1)
		if string(jsonMetadataValueOneOf1) == "{}" { // empty struct
			dst.MetadataValueOneOf1 = nil
		} else {
			match++
		}
	} else {
		dst.MetadataValueOneOf1 = nil
	}

	// try to unmarshal data into MetadataValueOneOf2
	err = json.Unmarshal(data, &dst.MetadataValueOneOf2)
	if err == nil {
		jsonMetadataValueOneOf2, _ := json.Marshal(dst.MetadataValueOneOf2)
		if string(jsonMetadataValueOneOf2) == "{}" { // empty struct
			dst.MetadataValueOneOf2 = nil
		} else {
			match++
		}
	} else {
		dst.MetadataValueOneOf2 = nil
	}

	// try to unmarshal data into MetadataValueOneOf3
	err = json.Unmarshal(data, &dst.MetadataValueOneOf3)
	if err == nil {
		jsonMetadataValueOneOf3, _ := json.Marshal(dst.MetadataValueOneOf3)
		if string(jsonMetadataValueOneOf3) == "{}" { // empty struct
			dst.MetadataValueOneOf3 = nil
		} else {
			match++
		}
	} else {
		dst.MetadataValueOneOf3 = nil
	}

	// try to unmarshal data into MetadataValueOneOf4
	err = json.Unmarshal(data, &dst.MetadataValueOneOf4)
	if err == nil {
		jsonMetadataValueOneOf4, _ := json.Marshal(dst.MetadataValueOneOf4)
		if string(jsonMetadataValueOneOf4) == "{}" { // empty struct
			dst.MetadataValueOneOf4 = nil
		} else {
			match++
		}
	} else {
		dst.MetadataValueOneOf4 = nil
	}

	// try to unmarshal data into MetadataValueOneOf5
	err = json.Unmarshal(data, &dst.MetadataValueOneOf5)
	if err == nil {
		jsonMetadataValueOneOf5, _ := json.Marshal(dst.MetadataValueOneOf5)
		if string(jsonMetadataValueOneOf5) == "{}" { // empty struct
			dst.MetadataValueOneOf5 = nil
		} else {
			match++
		}
	} else {
		dst.MetadataValueOneOf5 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MetadataValueOneOf = nil
		dst.MetadataValueOneOf1 = nil
		dst.MetadataValueOneOf2 = nil
		dst.MetadataValueOneOf3 = nil
		dst.MetadataValueOneOf4 = nil
		dst.MetadataValueOneOf5 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MetadataValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MetadataValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MetadataValue) MarshalJSON() ([]byte, error) {
	if src.MetadataValueOneOf != nil {
		return json.Marshal(&src.MetadataValueOneOf)
	}

	if src.MetadataValueOneOf1 != nil {
		return json.Marshal(&src.MetadataValueOneOf1)
	}

	if src.MetadataValueOneOf2 != nil {
		return json.Marshal(&src.MetadataValueOneOf2)
	}

	if src.MetadataValueOneOf3 != nil {
		return json.Marshal(&src.MetadataValueOneOf3)
	}

	if src.MetadataValueOneOf4 != nil {
		return json.Marshal(&src.MetadataValueOneOf4)
	}

	if src.MetadataValueOneOf5 != nil {
		return json.Marshal(&src.MetadataValueOneOf5)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MetadataValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MetadataValueOneOf != nil {
		return obj.MetadataValueOneOf
	}

	if obj.MetadataValueOneOf1 != nil {
		return obj.MetadataValueOneOf1
	}

	if obj.MetadataValueOneOf2 != nil {
		return obj.MetadataValueOneOf2
	}

	if obj.MetadataValueOneOf3 != nil {
		return obj.MetadataValueOneOf3
	}

	if obj.MetadataValueOneOf4 != nil {
		return obj.MetadataValueOneOf4
	}

	if obj.MetadataValueOneOf5 != nil {
		return obj.MetadataValueOneOf5
	}

	// all schemas are nil
	return nil
}

type NullableMetadataValue struct {
	value *MetadataValue
	isSet bool
}

func (v NullableMetadataValue) Get() *MetadataValue {
	return v.value
}

func (v *NullableMetadataValue) Set(val *MetadataValue) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataValue) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataValue(val *MetadataValue) *NullableMetadataValue {
	return &NullableMetadataValue{value: val, isSet: true}
}

func (v NullableMetadataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}