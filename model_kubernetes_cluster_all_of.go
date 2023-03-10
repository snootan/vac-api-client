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

// KubernetesClusterAllOf Kubernetes cluster
type KubernetesClusterAllOf struct {
	Workloads []KubernetesWorkload `json:"workloads"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesClusterAllOf KubernetesClusterAllOf

// NewKubernetesClusterAllOf instantiates a new KubernetesClusterAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesClusterAllOf(workloads []KubernetesWorkload) *KubernetesClusterAllOf {
	this := KubernetesClusterAllOf{}
	this.Workloads = workloads
	return &this
}

// NewKubernetesClusterAllOfWithDefaults instantiates a new KubernetesClusterAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesClusterAllOfWithDefaults() *KubernetesClusterAllOf {
	this := KubernetesClusterAllOf{}
	return &this
}

// GetWorkloads returns the Workloads field value
func (o *KubernetesClusterAllOf) GetWorkloads() []KubernetesWorkload {
	if o == nil {
		var ret []KubernetesWorkload
		return ret
	}

	return o.Workloads
}

// GetWorkloadsOk returns a tuple with the Workloads field value
// and a boolean to check if the value has been set.
func (o *KubernetesClusterAllOf) GetWorkloadsOk() ([]KubernetesWorkload, bool) {
	if o == nil {
		return nil, false
	}
	return o.Workloads, true
}

// SetWorkloads sets field value
func (o *KubernetesClusterAllOf) SetWorkloads(v []KubernetesWorkload) {
	o.Workloads = v
}

func (o KubernetesClusterAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["workloads"] = o.Workloads
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesClusterAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesClusterAllOf := _KubernetesClusterAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesClusterAllOf); err == nil {
		*o = KubernetesClusterAllOf(varKubernetesClusterAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "workloads")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesClusterAllOf struct {
	value *KubernetesClusterAllOf
	isSet bool
}

func (v NullableKubernetesClusterAllOf) Get() *KubernetesClusterAllOf {
	return v.value
}

func (v *NullableKubernetesClusterAllOf) Set(val *KubernetesClusterAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesClusterAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesClusterAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesClusterAllOf(val *KubernetesClusterAllOf) *NullableKubernetesClusterAllOf {
	return &NullableKubernetesClusterAllOf{value: val, isSet: true}
}

func (v NullableKubernetesClusterAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesClusterAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


