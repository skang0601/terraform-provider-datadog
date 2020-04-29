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

// DashboardDeleteResponse Response from the DeleteDashboard call.
type DashboardDeleteResponse struct {
	// ID of the deleted dashboard.
	DeletedDashboardId *string `json:"deleted_dashboard_id,omitempty"`
}

// NewDashboardDeleteResponse instantiates a new DashboardDeleteResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboardDeleteResponse() *DashboardDeleteResponse {
	this := DashboardDeleteResponse{}
	return &this
}

// NewDashboardDeleteResponseWithDefaults instantiates a new DashboardDeleteResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardDeleteResponseWithDefaults() *DashboardDeleteResponse {
	this := DashboardDeleteResponse{}
	return &this
}

// GetDeletedDashboardId returns the DeletedDashboardId field value if set, zero value otherwise.
func (o *DashboardDeleteResponse) GetDeletedDashboardId() string {
	if o == nil || o.DeletedDashboardId == nil {
		var ret string
		return ret
	}
	return *o.DeletedDashboardId
}

// GetDeletedDashboardIdOk returns a tuple with the DeletedDashboardId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DashboardDeleteResponse) GetDeletedDashboardIdOk() (*string, bool) {
	if o == nil || o.DeletedDashboardId == nil {
		return nil, false
	}
	return o.DeletedDashboardId, true
}

// HasDeletedDashboardId returns a boolean if a field has been set.
func (o *DashboardDeleteResponse) HasDeletedDashboardId() bool {
	if o != nil && o.DeletedDashboardId != nil {
		return true
	}

	return false
}

// SetDeletedDashboardId gets a reference to the given string and assigns it to the DeletedDashboardId field.
func (o *DashboardDeleteResponse) SetDeletedDashboardId(v string) {
	o.DeletedDashboardId = &v
}

func (o DashboardDeleteResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeletedDashboardId != nil {
		toSerialize["deleted_dashboard_id"] = o.DeletedDashboardId
	}
	return json.Marshal(toSerialize)
}

type NullableDashboardDeleteResponse struct {
	value *DashboardDeleteResponse
	isSet bool
}

func (v NullableDashboardDeleteResponse) Get() *DashboardDeleteResponse {
	return v.value
}

func (v *NullableDashboardDeleteResponse) Set(val *DashboardDeleteResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboardDeleteResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboardDeleteResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboardDeleteResponse(val *DashboardDeleteResponse) *NullableDashboardDeleteResponse {
	return &NullableDashboardDeleteResponse{value: val, isSet: true}
}

func (v NullableDashboardDeleteResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboardDeleteResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
