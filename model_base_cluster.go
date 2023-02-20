/*
Catalog Governor Service REST API

This is the service to track assets deployed in customer clusters

API version: 0.1.0
Contact: content-building-ecosystem@vmware.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// BaseCluster Base properties for a cluster
type BaseCluster struct {
	// name of the cluster where the workload is deployed
	Name string `json:"name"`
	// Kind of the cluster
	Kind string `json:"kind"`
	AdditionalProperties map[string]interface{}
}

type _BaseCluster BaseCluster

// NewBaseCluster instantiates a new BaseCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseCluster(name string, kind string) *BaseCluster {
	this := BaseCluster{}
	this.Name = name
	this.Kind = kind
	return &this
}

// NewBaseClusterWithDefaults instantiates a new BaseCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseClusterWithDefaults() *BaseCluster {
	this := BaseCluster{}
	return &this
}

// GetName returns the Name field value
func (o *BaseCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseCluster) SetName(v string) {
	o.Name = v
}

// GetKind returns the Kind field value
func (o *BaseCluster) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *BaseCluster) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *BaseCluster) SetKind(v string) {
	o.Kind = v
}

func (o BaseCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["kind"] = o.Kind
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BaseCluster) UnmarshalJSON(bytes []byte) (err error) {
	varBaseCluster := _BaseCluster{}

	if err = json.Unmarshal(bytes, &varBaseCluster); err == nil {
		*o = BaseCluster(varBaseCluster)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "kind")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBaseCluster struct {
	value *BaseCluster
	isSet bool
}

func (v NullableBaseCluster) Get() *BaseCluster {
	return v.value
}

func (v *NullableBaseCluster) Set(val *BaseCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseCluster(val *BaseCluster) *NullableBaseCluster {
	return &NullableBaseCluster{value: val, isSet: true}
}

func (v NullableBaseCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

