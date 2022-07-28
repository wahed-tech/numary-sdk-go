/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.0-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// ListAccounts400Response struct for ListAccounts400Response
type ListAccounts400Response struct {
	ErrorCode string `json:"error_code"`
	ErrorMessage *string `json:"error_message,omitempty"`
}

// NewListAccounts400Response instantiates a new ListAccounts400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListAccounts400Response(errorCode string) *ListAccounts400Response {
	this := ListAccounts400Response{}
	this.ErrorCode = errorCode
	return &this
}

// NewListAccounts400ResponseWithDefaults instantiates a new ListAccounts400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListAccounts400ResponseWithDefaults() *ListAccounts400Response {
	this := ListAccounts400Response{}
	return &this
}

// GetErrorCode returns the ErrorCode field value
func (o *ListAccounts400Response) GetErrorCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value
// and a boolean to check if the value has been set.
func (o *ListAccounts400Response) GetErrorCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ErrorCode, true
}

// SetErrorCode sets field value
func (o *ListAccounts400Response) SetErrorCode(v string) {
	o.ErrorCode = v
}

// GetErrorMessage returns the ErrorMessage field value if set, zero value otherwise.
func (o *ListAccounts400Response) GetErrorMessage() string {
	if o == nil || o.ErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListAccounts400Response) GetErrorMessageOk() (*string, bool) {
	if o == nil || o.ErrorMessage == nil {
		return nil, false
	}
	return o.ErrorMessage, true
}

// HasErrorMessage returns a boolean if a field has been set.
func (o *ListAccounts400Response) HasErrorMessage() bool {
	if o != nil && o.ErrorMessage != nil {
		return true
	}

	return false
}

// SetErrorMessage gets a reference to the given string and assigns it to the ErrorMessage field.
func (o *ListAccounts400Response) SetErrorMessage(v string) {
	o.ErrorMessage = &v
}

func (o ListAccounts400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error_code"] = o.ErrorCode
	}
	if o.ErrorMessage != nil {
		toSerialize["error_message"] = o.ErrorMessage
	}
	return json.Marshal(toSerialize)
}

type NullableListAccounts400Response struct {
	value *ListAccounts400Response
	isSet bool
}

func (v NullableListAccounts400Response) Get() *ListAccounts400Response {
	return v.value
}

func (v *NullableListAccounts400Response) Set(val *ListAccounts400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListAccounts400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListAccounts400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListAccounts400Response(val *ListAccounts400Response) *NullableListAccounts400Response {
	return &NullableListAccounts400Response{value: val, isSet: true}
}

func (v NullableListAccounts400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListAccounts400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


