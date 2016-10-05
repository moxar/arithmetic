package arithmetic

// import (
// 	"fmt"
// 	"strings"
// )
//
// type Function func(...interface{}) (interface{}, error)
//
// type function struct {
// 	label string
// 	solve Function
// }
//
// func (f function) String() string {
// 	return f.label
// }
//
// var functions = map[string]function{}
//
// func RegisterFunction(label string, f Function) {
//
// 	if f == nil {
// 		panic(fmt.Sprintf("function %s is nil", label))
// 	}
//
// 	label = strings.ToLower(label)
//
// 	if _, ok := functions[label]; ok {
// 		panic(fmt.Sprintf("%s already defined as function", label))
// 	}
//
// 	if _, ok := variables[label]; ok {
// 		panic(fmt.Sprintf("%s already defined as variable", label))
// 	}
//
// 	var o function
// 	o.label = label
// 	o.solve = f
//
// 	functions[label] = o
// }
