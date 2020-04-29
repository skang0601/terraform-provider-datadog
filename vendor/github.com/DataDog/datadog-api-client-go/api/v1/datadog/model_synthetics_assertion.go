/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// SyntheticsAssertion TODO.
type SyntheticsAssertion struct {
	Operator SyntheticsAssertionOperator `json:"operator"`
	// TODO.
	Property *string `json:"property,omitempty"`
	// TODO.
	Target *interface{}            `json:"target,omitempty"`
	Type   SyntheticsAssertionType `json:"type"`
}

// NewSyntheticsAssertion instantiates a new SyntheticsAssertion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSyntheticsAssertion(operator SyntheticsAssertionOperator, type_ SyntheticsAssertionType) *SyntheticsAssertion {
	this := SyntheticsAssertion{}
	this.Operator = operator
	this.Type = type_
	return &this
}

// NewSyntheticsAssertionWithDefaults instantiates a new SyntheticsAssertion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSyntheticsAssertionWithDefaults() *SyntheticsAssertion {
	this := SyntheticsAssertion{}
	return &this
}

// GetOperator returns the Operator field value
func (o *SyntheticsAssertion) GetOperator() SyntheticsAssertionOperator {
	if o == nil {
		var ret SyntheticsAssertionOperator
		return ret
	}

	return o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertion) GetOperatorOk() (*SyntheticsAssertionOperator, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Operator, true
}

// SetOperator sets field value
func (o *SyntheticsAssertion) SetOperator(v SyntheticsAssertionOperator) {
	o.Operator = v
}

// GetProperty returns the Property field value if set, zero value otherwise.
func (o *SyntheticsAssertion) GetProperty() string {
	if o == nil || o.Property == nil {
		var ret string
		return ret
	}
	return *o.Property
}

// GetPropertyOk returns a tuple with the Property field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertion) GetPropertyOk() (*string, bool) {
	if o == nil || o.Property == nil {
		return nil, false
	}
	return o.Property, true
}

// HasProperty returns a boolean if a field has been set.
func (o *SyntheticsAssertion) HasProperty() bool {
	if o != nil && o.Property != nil {
		return true
	}

	return false
}

// SetProperty gets a reference to the given string and assigns it to the Property field.
func (o *SyntheticsAssertion) SetProperty(v string) {
	o.Property = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *SyntheticsAssertion) GetTarget() interface{} {
	if o == nil || o.Target == nil {
		var ret interface{}
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertion) GetTargetOk() (*interface{}, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *SyntheticsAssertion) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given interface{} and assigns it to the Target field.
func (o *SyntheticsAssertion) SetTarget(v interface{}) {
	o.Target = &v
}

// GetType returns the Type field value
func (o *SyntheticsAssertion) GetType() SyntheticsAssertionType {
	if o == nil {
		var ret SyntheticsAssertionType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SyntheticsAssertion) GetTypeOk() (*SyntheticsAssertionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SyntheticsAssertion) SetType(v SyntheticsAssertionType) {
	o.Type = v
}

func (o SyntheticsAssertion) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["operator"] = o.Operator
	}
	if o.Property != nil {
		toSerialize["property"] = o.Property
	}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableSyntheticsAssertion struct {
	value *SyntheticsAssertion
	isSet bool
}

func (v NullableSyntheticsAssertion) Get() *SyntheticsAssertion {
	return v.value
}

func (v *NullableSyntheticsAssertion) Set(val *SyntheticsAssertion) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsAssertion) IsSet() bool {
	return v.isSet
}

func (v *NullableSyntheticsAssertion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsAssertion(val *SyntheticsAssertion) *NullableSyntheticsAssertion {
	return &NullableSyntheticsAssertion{value: val, isSet: true}
}

func (v NullableSyntheticsAssertion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsAssertion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
