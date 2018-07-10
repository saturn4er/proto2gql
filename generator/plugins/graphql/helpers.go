package graphql

import (
	"reflect"
)

func typeIsScalar(p GoType) bool {
	switch p.Kind {
	case reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128,
		reflect.String:
		return true
	}
	return false
}

func ResolverCall(resolverPkg, resolverFuncName string) ValueResolver {
	return func(arg string, ctx BodyContext) string {
		if ctx.TracerEnabled {
			return ctx.Importer.Prefix(resolverPkg) + resolverFuncName + "(tr, tr.ContextWithSpan(ctx, span), " + arg + ")"
		}
		return ctx.Importer.Prefix(resolverPkg) + resolverFuncName + "(ctx, " + arg + ")"
	}
}
