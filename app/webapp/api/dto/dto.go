package dto

import (
	"errors"
	"reflect"
	"timeCardSimple/app/domain/employee"
	"timeCardSimple/app/domain/id"
	"timeCardSimple/app/domain/timecard"
)

var ErrUnknownType = errors.New("dto: unknown type to marshal")

type Transformer interface {
	Transform(interface{}) interface{}
}

type TransformerFunc func(interface{}) interface{}

func (tf TransformerFunc) Transform(v interface{}) interface{} {
	return tf(v)
}

var transformers map[reflect.Type]Transformer

func init() {
	transformers = map[reflect.Type]Transformer{
		reflect.TypeOf([]id.ID{}):  TransformerFunc(Ids),
		reflect.TypeOf(id.Empty()): TransformerFunc(Id),

		reflect.TypeOf([]*employee.Employee{}): TransformerFunc(Employees),
		reflect.TypeOf(&employee.Employee{}):   TransformerFunc(Employee),

		reflect.TypeOf(&timecard.Timecard{}): TransformerFunc(Timecard),
	}

}

func Transform(v interface{}) (interface{}, error) {
	if v == nil {
		return map[string]interface{}{}, nil
	}

	transformer, ok := transformers[reflect.TypeOf(v)]
	if !ok {
		return nil, ErrUnknownType
	}
	return transformer.Transform(v), nil
}
