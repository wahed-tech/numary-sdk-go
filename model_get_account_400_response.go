/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.8
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// GetAccount400Response struct for GetAccount400Response
type GetAccount400Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewGetAccount400Response instantiates a new GetAccount400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAccount400Response(errorCode string) *GetAccount400Response {
	this := GetAccount400Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewGetAccount400ResponseWithDefaults instantiates a new GetAccount400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAccount400ResponseWithDefaults() *GetAccount400Response {
	this := GetAccount400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *GetAccount400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *GetAccount400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *GetAccount400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *GetAccount400Response) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAccount400Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *GetAccount400Response) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *GetAccount400Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o GetAccount400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableGetAccount400Response struct {
	value *GetAccount400Response
	isSet bool
}

func (v NullableGetAccount400Response) Get() *GetAccount400Response {
	return v.value
}

func (v *NullableGetAccount400Response) Set(val *GetAccount400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAccount400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAccount400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAccount400Response(val *GetAccount400Response) *NullableGetAccount400Response {
	return &NullableGetAccount400Response{value: val, isSet: true}
}

func (v NullableGetAccount400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAccount400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


