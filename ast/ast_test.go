package ast

import (
	"github.com/google/go-cmp/cmp/cmpopts"
)

var IgnoreOpt = cmpopts.IgnoreUnexported(Reference{}, FunctionType{})
