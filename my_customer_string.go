package gqlgen_todos

import (
	"fmt"
	"io"
)

// MyCustomStringScalar _
type MyCustomStringScalar string

// MarshalGQL _
func (m MyCustomStringScalar) MarshalGQL(w io.Writer) {
	w.Write([]byte(m))
}

// UnmarshalGQL _
func (m *MyCustomStringScalar) UnmarshalGQL(v interface{}) error {
	switch v := v.(type) {
	case string:
		*m = MyCustomStringScalar(v)
		return nil
	default:
		return fmt.Errorf("%T is not a string", v)
	}
}
