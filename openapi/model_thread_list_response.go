/*
 * BBS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ThreadListResponse struct {
	ThreadList []ThreadListResponseThreadList `json:"thread_list"`
}

// AssertThreadListResponseRequired checks if the required fields are not zero-ed
func AssertThreadListResponseRequired(obj ThreadListResponse) error {
	elements := map[string]interface{}{
		"thread_list": obj.ThreadList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.ThreadList {
		if err := AssertThreadListResponseThreadListRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseThreadListResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ThreadListResponse (e.g. [][]ThreadListResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseThreadListResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aThreadListResponse, ok := obj.(ThreadListResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertThreadListResponseRequired(aThreadListResponse)
	})
}
