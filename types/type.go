package types

var (
	IntType   = &Tuple{}
	TypeType  = &Tuple{}
	ErrorType = &Tuple{}
)

type Type interface{}

type Tuple struct {
	Fields []TupleField
}

type TupleField struct {
	Name string
	Type Type
}

var _ Type = &Tuple{}

type FunctionType struct {
	Inputs  []Type
	Outputs []Type
}

var _ Type = &FunctionType{}
