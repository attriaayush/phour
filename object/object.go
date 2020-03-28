package object

import "fmt"

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
)

// NewEnvironment ..
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// ObjectType . .
type ObjectType string

// Object . .
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer for when we encounter any integers during evaluation
// and then convert that into an object by saving the value
// in the struct
type Integer struct {
	Value int64
}

// Boolean ..
type Boolean struct {
	Value bool
}

// Null ..
type Null struct{}

// ReturnValue ..
type ReturnValue struct {
	Value Object
}

// Error handler
type Error struct {
	Message string
}

// Environment struct
type Environment struct {
	store map[string]Object
}

// Type for Integer
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Inspect for Integer
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type for Boolean
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Inspect for Boolean
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type for Null
func (b *Null) Type() ObjectType { return NULL_OBJ }

// Inspect for Null
func (b *Null) Inspect() string { return "null" }

// Type for Return
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect for Return
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Type for Error
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect for Error
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Get for Environment Getter
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set for Environment Setter
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
