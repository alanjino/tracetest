/*
TraceTest

OpenAPI definition for TraceTest endpoint and resources

API version: 0.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AssertionResult struct for AssertionResult
type AssertionResult struct {
	Assertion   *Assertion            `json:"assertion,omitempty"`
	AllPassed   *bool                 `json:"allPassed,omitempty"`
	SpanResults []AssertionSpanResult `json:"spanResults,omitempty"`
}

// NewAssertionResult instantiates a new AssertionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssertionResult() *AssertionResult {
	this := AssertionResult{}
	return &this
}

// NewAssertionResultWithDefaults instantiates a new AssertionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssertionResultWithDefaults() *AssertionResult {
	this := AssertionResult{}
	return &this
}

// GetAssertion returns the Assertion field value if set, zero value otherwise.
func (o *AssertionResult) GetAssertion() Assertion {
	if o == nil || o.Assertion == nil {
		var ret Assertion
		return ret
	}
	return *o.Assertion
}

// GetAssertionOk returns a tuple with the Assertion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssertionResult) GetAssertionOk() (*Assertion, bool) {
	if o == nil || o.Assertion == nil {
		return nil, false
	}
	return o.Assertion, true
}

// HasAssertion returns a boolean if a field has been set.
func (o *AssertionResult) HasAssertion() bool {
	if o != nil && o.Assertion != nil {
		return true
	}

	return false
}

// SetAssertion gets a reference to the given Assertion and assigns it to the Assertion field.
func (o *AssertionResult) SetAssertion(v Assertion) {
	o.Assertion = &v
}

// GetAllPassed returns the AllPassed field value if set, zero value otherwise.
func (o *AssertionResult) GetAllPassed() bool {
	if o == nil || o.AllPassed == nil {
		var ret bool
		return ret
	}
	return *o.AllPassed
}

// GetAllPassedOk returns a tuple with the AllPassed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssertionResult) GetAllPassedOk() (*bool, bool) {
	if o == nil || o.AllPassed == nil {
		return nil, false
	}
	return o.AllPassed, true
}

// HasAllPassed returns a boolean if a field has been set.
func (o *AssertionResult) HasAllPassed() bool {
	if o != nil && o.AllPassed != nil {
		return true
	}

	return false
}

// SetAllPassed gets a reference to the given bool and assigns it to the AllPassed field.
func (o *AssertionResult) SetAllPassed(v bool) {
	o.AllPassed = &v
}

// GetSpanResults returns the SpanResults field value if set, zero value otherwise.
func (o *AssertionResult) GetSpanResults() []AssertionSpanResult {
	if o == nil || o.SpanResults == nil {
		var ret []AssertionSpanResult
		return ret
	}
	return o.SpanResults
}

// GetSpanResultsOk returns a tuple with the SpanResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssertionResult) GetSpanResultsOk() ([]AssertionSpanResult, bool) {
	if o == nil || o.SpanResults == nil {
		return nil, false
	}
	return o.SpanResults, true
}

// HasSpanResults returns a boolean if a field has been set.
func (o *AssertionResult) HasSpanResults() bool {
	if o != nil && o.SpanResults != nil {
		return true
	}

	return false
}

// SetSpanResults gets a reference to the given []AssertionSpanResult and assigns it to the SpanResults field.
func (o *AssertionResult) SetSpanResults(v []AssertionSpanResult) {
	o.SpanResults = v
}

func (o AssertionResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Assertion != nil {
		toSerialize["assertion"] = o.Assertion
	}
	if o.AllPassed != nil {
		toSerialize["allPassed"] = o.AllPassed
	}
	if o.SpanResults != nil {
		toSerialize["spanResults"] = o.SpanResults
	}
	return json.Marshal(toSerialize)
}

type NullableAssertionResult struct {
	value *AssertionResult
	isSet bool
}

func (v NullableAssertionResult) Get() *AssertionResult {
	return v.value
}

func (v *NullableAssertionResult) Set(val *AssertionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAssertionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAssertionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssertionResult(val *AssertionResult) *NullableAssertionResult {
	return &NullableAssertionResult{value: val, isSet: true}
}

func (v NullableAssertionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssertionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}