/*
Ledger API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1.7.0-beta.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ledgerclient

import (
	"encoding/json"
)

// ListTransactions200ResponseCursor struct for ListTransactions200ResponseCursor
type ListTransactions200ResponseCursor struct {
	PageSize int32 `json:"page_size"`
	HasMore *bool `json:"has_more,omitempty"`
	Previous *string `json:"previous,omitempty"`
	Next *string `json:"next,omitempty"`
	Data []Transaction `json:"data"`
}

// NewListTransactions200ResponseCursor instantiates a new ListTransactions200ResponseCursor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListTransactions200ResponseCursor(pageSize int32, data []Transaction) *ListTransactions200ResponseCursor {
	this := ListTransactions200ResponseCursor{}
	this.PageSize = pageSize
	this.Data = data
	return &this
}

// NewListTransactions200ResponseCursorWithDefaults instantiates a new ListTransactions200ResponseCursor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListTransactions200ResponseCursorWithDefaults() *ListTransactions200ResponseCursor {
	this := ListTransactions200ResponseCursor{}
	return &this
}

// GetPageSize returns the PageSize field value
func (o *ListTransactions200ResponseCursor) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *ListTransactions200ResponseCursor) GetPageSizeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *ListTransactions200ResponseCursor) SetPageSize(v int32) {
	o.PageSize = v
}

// GetHasMore returns the HasMore field value if set, zero value otherwise.
func (o *ListTransactions200ResponseCursor) GetHasMore() bool {
	if o == nil || o.HasMore == nil {
		var ret bool
		return ret
	}
	return *o.HasMore
}

// GetHasMoreOk returns a tuple with the HasMore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactions200ResponseCursor) GetHasMoreOk() (*bool, bool) {
	if o == nil || o.HasMore == nil {
		return nil, false
	}
	return o.HasMore, true
}

// HasHasMore returns a boolean if a field has been set.
func (o *ListTransactions200ResponseCursor) HasHasMore() bool {
	if o != nil && o.HasMore != nil {
		return true
	}

	return false
}

// SetHasMore gets a reference to the given bool and assigns it to the HasMore field.
func (o *ListTransactions200ResponseCursor) SetHasMore(v bool) {
	o.HasMore = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *ListTransactions200ResponseCursor) GetPrevious() string {
	if o == nil || o.Previous == nil {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactions200ResponseCursor) GetPreviousOk() (*string, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *ListTransactions200ResponseCursor) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *ListTransactions200ResponseCursor) SetPrevious(v string) {
	o.Previous = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ListTransactions200ResponseCursor) GetNext() string {
	if o == nil || o.Next == nil {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListTransactions200ResponseCursor) GetNextOk() (*string, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ListTransactions200ResponseCursor) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *ListTransactions200ResponseCursor) SetNext(v string) {
	o.Next = &v
}

// GetData returns the Data field value
func (o *ListTransactions200ResponseCursor) GetData() []Transaction {
	if o == nil {
		var ret []Transaction
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListTransactions200ResponseCursor) GetDataOk() (*[]Transaction, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *ListTransactions200ResponseCursor) SetData(v []Transaction) {
	o.Data = v
}

func (o ListTransactions200ResponseCursor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["page_size"] = o.PageSize
	}
	if o.HasMore != nil {
		toSerialize["has_more"] = o.HasMore
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListTransactions200ResponseCursor struct {
	value *ListTransactions200ResponseCursor
	isSet bool
}

func (v NullableListTransactions200ResponseCursor) Get() *ListTransactions200ResponseCursor {
	return v.value
}

func (v *NullableListTransactions200ResponseCursor) Set(val *ListTransactions200ResponseCursor) {
	v.value = val
	v.isSet = true
}

func (v NullableListTransactions200ResponseCursor) IsSet() bool {
	return v.isSet
}

func (v *NullableListTransactions200ResponseCursor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListTransactions200ResponseCursor(val *ListTransactions200ResponseCursor) *NullableListTransactions200ResponseCursor {
	return &NullableListTransactions200ResponseCursor{value: val, isSet: true}
}

func (v NullableListTransactions200ResponseCursor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListTransactions200ResponseCursor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


