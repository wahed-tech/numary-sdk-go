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

// GetBalances200Response struct for GetBalances200Response
type GetBalances200Response struct {
	Cursor GetBalances200ResponseCursor `json:"cursor"`
}

// NewGetBalances200Response instantiates a new GetBalances200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBalances200Response(cursor GetBalances200ResponseCursor) *GetBalances200Response {
	this := GetBalances200Response{}
	this.Cursor = cursor
	return &this
}

// NewGetBalances200ResponseWithDefaults instantiates a new GetBalances200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBalances200ResponseWithDefaults() *GetBalances200Response {
	this := GetBalances200Response{}
	return &this
}

// GetCursor returns the Cursor field value
func (o *GetBalances200Response) GetCursor() GetBalances200ResponseCursor {
	if o == nil {
		var ret GetBalances200ResponseCursor
		return ret
	}

	return o.Cursor
}

// GetCursorOk returns a tuple with the Cursor field value
// and a boolean to check if the value has been set.
func (o *GetBalances200Response) GetCursorOk() (*GetBalances200ResponseCursor, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cursor, true
}

// SetCursor sets field value
func (o *GetBalances200Response) SetCursor(v GetBalances200ResponseCursor) {
	o.Cursor = v
}

func (o GetBalances200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cursor"] = o.Cursor
	}
	return json.Marshal(toSerialize)
}

type NullableGetBalances200Response struct {
	value *GetBalances200Response
	isSet bool
}

func (v NullableGetBalances200Response) Get() *GetBalances200Response {
	return v.value
}

func (v *NullableGetBalances200Response) Set(val *GetBalances200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetBalances200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetBalances200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetBalances200Response(val *GetBalances200Response) *NullableGetBalances200Response {
	return &NullableGetBalances200Response{value: val, isSet: true}
}

func (v NullableGetBalances200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetBalances200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


