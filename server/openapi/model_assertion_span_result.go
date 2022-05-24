/*
 * TraceTest
 *
 * OpenAPI definition for TraceTest endpoint and resources
 *
 * API version: 0.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type AssertionSpanResult struct {
	SpanId string `json:"spanId,omitempty"`

	ObservedValue string `json:"observedValue,omitempty"`

	Passed bool `json:"passed,omitempty"`

	Error string `json:"error,omitempty"`
}

// AssertAssertionSpanResultRequired checks if the required fields are not zero-ed
func AssertAssertionSpanResultRequired(obj AssertionSpanResult) error {
	return nil
}

// AssertRecurseAssertionSpanResultRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of AssertionSpanResult (e.g. [][]AssertionSpanResult), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseAssertionSpanResultRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aAssertionSpanResult, ok := obj.(AssertionSpanResult)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertAssertionSpanResultRequired(aAssertionSpanResult)
	})
}