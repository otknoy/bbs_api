/*
 * BBS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Comment struct {
	Meta CommentMeta `json:"meta"`

	Text string `json:"text,omitempty"`
}

// AssertCommentRequired checks if the required fields are not zero-ed
func AssertCommentRequired(obj Comment) error {
	elements := map[string]interface{}{
		"meta": obj.Meta,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCommentMetaRequired(obj.Meta); err != nil {
		return err
	}
	return nil
}

// AssertRecurseCommentRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Comment (e.g. [][]Comment), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseCommentRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aComment, ok := obj.(Comment)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertCommentRequired(aComment)
	})
}
