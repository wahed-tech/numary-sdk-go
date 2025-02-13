/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// Mapping struct for Mapping
type Mapping struct {
	Contracts []Contract `json:"contracts"`
}

// NewMapping instantiates a new Mapping object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapping(contracts []Contract) *Mapping {
	this := Mapping{}
	this.Contracts = contracts
	return &this
}

// NewMappingWithDefaults instantiates a new Mapping object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMappingWithDefaults() *Mapping {
	this := Mapping{}
	return &this
}

// GetContracts returns the Contracts field value
func (o *Mapping) GetContracts() []Contract {
	if o == nil {
		var ret []Contract
		return ret
	}

	return o.Contracts
}

// GetContractsOk returns a tuple with the Contracts field value
// and a boolean to check if the value has been set.
func (o *Mapping) GetContractsOk() (*[]Contract, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Contracts, true
}

// SetContracts sets field value
func (o *Mapping) SetContracts(v []Contract) {
	o.Contracts = v
}

func (o Mapping) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contracts"] = o.Contracts
	}
	return json.Marshal(toSerialize)
}

type NullableMapping struct {
	value *Mapping
	isSet bool
}

func (v NullableMapping) Get() *Mapping {
	return v.value
}

func (v *NullableMapping) Set(val *Mapping) {
	v.value = val
	v.isSet = true
}

func (v NullableMapping) IsSet() bool {
	return v.isSet
}

func (v *NullableMapping) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapping(val *Mapping) *NullableMapping {
	return &NullableMapping{value: val, isSet: true}
}

func (v NullableMapping) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapping) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


