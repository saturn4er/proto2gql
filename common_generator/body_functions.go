package common_generator

type BodyFunctions struct {
	importer *Importer
}

func (b BodyFunctions) importFunc(path string) func() string {
	return func() string {
		return b.importer.New(path)
	}
}
func (b *BodyFunctions) gqlPkg() func() string {
	return b.importFunc(GraphqlPkgPath)
}
func (b *BodyFunctions) scalarsPkg() func() string {
	return b.importFunc(ScalarsPkgPath)
}
func (b *BodyFunctions) interceptorsPkg() func() string {
	return b.importFunc(InterceptorsPkgPath)
}
func (b *BodyFunctions) opentracingPkg() func() string {
	return b.importFunc(OpentracingPkgPath)
}
func (b *BodyFunctions) tracerPkg() func() string {
	return b.importFunc(TracerPkgPath)
}
