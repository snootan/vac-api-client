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
	"fmt"
)

// Cluster - Cluster where workload are deployed
type Cluster struct {
	KubernetesCluster *KubernetesCluster
}

// KubernetesClusterAsCluster is a convenience function that returns KubernetesCluster wrapped in Cluster
func KubernetesClusterAsCluster(v *KubernetesCluster) Cluster {
	return Cluster{
		KubernetesCluster: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Cluster) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into KubernetesCluster
	err = newStrictDecoder(data).Decode(&dst.KubernetesCluster)
	if err == nil {
		jsonKubernetesCluster, _ := json.Marshal(dst.KubernetesCluster)
		if string(jsonKubernetesCluster) == "{}" { // empty struct
			dst.KubernetesCluster = nil
		} else {
			match++
		}
	} else {
		dst.KubernetesCluster = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.KubernetesCluster = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(Cluster)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(Cluster)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Cluster) MarshalJSON() ([]byte, error) {
	if src.KubernetesCluster != nil {
		return json.Marshal(&src.KubernetesCluster)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Cluster) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.KubernetesCluster != nil {
		return obj.KubernetesCluster
	}

	// all schemas are nil
	return nil
}

type NullableCluster struct {
	value *Cluster
	isSet bool
}

func (v NullableCluster) Get() *Cluster {
	return v.value
}

func (v *NullableCluster) Set(val *Cluster) {
	v.value = val
	v.isSet = true
}

func (v NullableCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCluster(val *Cluster) *NullableCluster {
	return &NullableCluster{value: val, isSet: true}
}

func (v NullableCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


