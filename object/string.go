package object

import (
	"bytes"
	"hash/fnv"
	"strings"
)

type String struct {
	Value string
}

var stringProperties = map[string]BuiltinObjectFunction{
	"length":      length,
	"toUpperCase": toUpperCase,
	"charAt":      charAt,
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string {
	var out bytes.Buffer

	out.WriteByte('"')
	out.WriteString(s.Value)
	out.WriteByte('"')

	return out.String()
}
func (s *String) HashKey() HashKey {
	// TODO: Potential for collisions. Can we do better? Will Java's String hashCode algorithm work?
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
func (s *String) Properties() map[string]BuiltinObjectFunction {
	return stringProperties
}

//func wrapStringFunc() BuiltinObjectFunction {
//	return func() {}
//}

func length(obj Object, _ ...Object) Object {
	return func(s *String) Object {
		return &Integer{Value: int64(len(s.Value))}
	}(obj.(*String))
}

func toUpperCase(obj Object, _ ...Object) Object {
	return func(s *String) Object {
		return &String{Value: strings.ToUpper(s.Value)}
	}(obj.(*String))
}

func charAt(obj Object, args ...Object) Object {
	return func(s *String, args ...Object) Object {
		index := 0
		if len(args) > 0 && args[0].Type() == INTEGER_OBJ {
			index = int(args[0].(*Integer).Value)
		}

		value := ""
		if index < len(s.Value) {
			value = string(s.Value[index])
		}

		return &String{Value: value}
	}(obj.(*String), args...)
}

//func init() {
//	GlobalEnvironment[STRING_OBJ] = map[string]BuiltinObjectFunction{
//		"length": func(s Object, _ ...Object) Object {
//			return &Integer{Value: int64(len(s.(*String).Value))}
//		},
//		"toUpperCase": func(s Object, _ ...Object) Object {
//			return &String{Value: strings.ToUpper(s.(*String).Value)}
//		},
//		"charAt": func(s Object, args ...Object) Object {
//			index := 0
//			if len(args) > 0 && args[0].Type() == INTEGER_OBJ {
//				index = int(args[0].(*Integer).Value)
//			}
//
//			value := ""
//			if index < len(s.(*String).Value) {
//				value = string(s.(*String).Value[index])
//			}
//
//			return &String{Value: value}
//		},
//	}
//}
