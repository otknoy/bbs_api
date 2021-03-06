/*
 * BBS API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type Board struct {
	Name string `json:"name"`

	ServerId string `json:"server_id"`

	BoardId string `json:"board_id"`
}

// AssertBoardRequired checks if the required fields are not zero-ed
func AssertBoardRequired(obj Board) error {
	elements := map[string]interface{}{
		"name":      obj.Name,
		"server_id": obj.ServerId,
		"board_id":  obj.BoardId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseBoardRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Board (e.g. [][]Board), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseBoardRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aBoard, ok := obj.(Board)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertBoardRequired(aBoard)
	})
}
