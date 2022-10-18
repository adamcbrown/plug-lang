package ast

import (
	"github.com/acbrown/plug-lang/lexer/token"
)

type Context struct {
	scopes      []map[string]Expr
	resolutions map[*Reference]Expr

	errors []NodeErr
}

func NewContext() *Context {
	return &Context{
		scopes:      nil,
		resolutions: make(map[*Reference]Expr),
		errors:      nil,
	}
}

func (ctx *Context) Resolve(name string) (Expr, bool) {
	for i := range ctx.scopes {
		scope := ctx.scopes[len(ctx.scopes)-1-i]
		if node, ok := scope[name]; ok {
			return node, true
		}
	}
	return nil, false
}

func (ctx *Context) Resolved(ref *Reference) (Expr, bool) {
	res, ok := ctx.resolutions[ref]
	return res, ok
}

func (ctx *Context) PushScope(scope map[string]Expr) {
	ctx.scopes = append(ctx.scopes, scope)
}

func (ctx *Context) PopScope() {
	ctx.scopes = ctx.scopes[:len(ctx.scopes)-1]
}

func (ctx *Context) AddError(err NodeErr) {
	ctx.errors = append(ctx.errors, err)
}

func (ctx *Context) AddReference(ref *Reference, expr Expr) {
	ctx.resolutions[ref] = expr
}

type Node interface {
	Start() int
	End() int
	AddReferences(ctx *Context)
}

type ParseErr struct {
	Msg string
	Tok token.Token
}

type NodeErr struct {
	Msg string
	N   Node
}
