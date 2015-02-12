package log

import (
	"fmt"
)

// return type name
func ParseToType(v Any) string {
	switch t := v.(type) {
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("%T", t)
	}
}

// return value
func ParseToString(v Any) string {
	switch t := v.(type) {
	case string:
		return t
	case int, int32, int64, uint, uint32, uint64:
		return fmt.Sprintf("%d", t)
	case float32, float64:
		return fmt.Sprintf("%f", t)
	case bool:
		return fmt.Sprintf("%t", t)
	case nil:
		return "<nil>"
	default:
		return fmt.Sprintf("%+v", t)
	}
}

// check if the typename has pointer
func HasPointer(typeName string) bool {
	switch typeName {
	case "string", "int", "int32", "int64", "uint", "uint32", "uint64":
		return false
	case "float32", "float64", "bool":
		return false
	default:
		return true
	}
}

// the type for anything
type Any interface{}
