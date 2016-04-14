package attributes

import (
	"errors"
	"fmt"
	"reflect"
)

// Pool attributes
type Health float32
type Energy float32
type Spirit float32

// Scalar attributes
type Summoning float32
type Alteration float32
type Willpower float32
type Divinity float32
type Lifebringing float32

type Attribute interface{}

func AttributeString(a Attribute) string {
	r := reflect.ValueOf(a)
	if r.Kind() != reflect.Float32 {
		panic(errors.New(fmt.Sprintf("%v cannot be considered a float32.\n", a)))
	}
	v := r.Float()
	if v >= 0 {
		return "+" + fmt.Sprintf("%.2f", v)
	}
	return fmt.Sprintf("%.2f", v)
}
