// Code generated by github.com/Morkow/gqlgen, DO NOT EDIT.

package todo

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
)

// Passed to createTodo to create a new todo
type TodoInput struct {
	Text string `json:"text"`
	Done *bool  `json:"done"`
}

type Role string

const (
	RoleAdmin Role = "ADMIN"
	RoleOwner Role = "OWNER"
)

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleOwner:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
