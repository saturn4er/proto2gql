package generator

const headTemplate = `{{ if $.pkg }}
	package {{$.pkg}}
{{else}}
	package not_defined
{{end}}

import (
    {{ range $import := .imports -}}
        {{$import.Alias}} "{{$import.Path}}"
    {{ end }}
)

var (
	_ = {{$.errorspkg}}.New
	_ = {{$.gqlpkg}}.NewObject
	_ = {{$.ctxpkg}}.Background
	_ = {{$.strconvpkg}}.FormatBool
	_ = {{$.fmtpkg}}.Print
	_ = {{$.opentracingpkg}}.GlobalTracer
	_ = {{$.debugpkg}}.FreeOSMemory
)

type(
	_ = interceptors.CallInterceptor
)`
const bodyTemplate = `
// Enums
{{range $enum := .File.Enums -}}
	var {{call $.GQLInputTypeName .Type}} = {{$.gqlpkg}}.NewEnum({{$.gqlpkg}}.EnumConfig{
		Name:        "{{call $.GQLInputTypeName .Type}}",
		Description: {{.QuotedComment}},
		Values: {{$.gqlpkg}}.EnumValueConfigMap{
            {{range .Values -}}
				"{{.Name}}": &{{$.gqlpkg}}.EnumValueConfig{
					Value: {{.Value}},
					Description: {{.QuotedComment}},
				},
			{{end}}
		},
	})
{{end}}

// Messages
{{ range $msg := .File.Messages -}}
	{{if and $msg.OutputMessage (call $.MessageHaveFieldsExceptError $msg)}}
		var {{call $.GQLOutputTypeName .Type}} = {{$.gqlpkg}}.NewObject({{$.gqlpkg}}.ObjectConfig{
			Name: "{{call $.GQLOutputTypeName  .Type}}",
            Fields: {{$.gqlpkg}}.Fields{
			},
		})
	{{end -}}
    {{if and $msg.InputMessage (call $.MessageHaveFieldsExceptError $msg)}}
        var {{call $.GQLInputTypeName .Type}} = {{$.gqlpkg}}.NewInputObject({{$.gqlpkg}}.InputObjectConfig{
            Name: "{{call $.GQLInputTypeName .Type}}",
            Fields: {{$.gqlpkg}}.InputObjectConfigFieldMapThunk(func() {{$.gqlpkg}}.InputObjectConfigFieldMap {
                return {{$.gqlpkg}}.InputObjectConfigFieldMap{
                    // {{$msg.Name}} fields
                    {{range $msg.Fields}}
						
                        {{if not (call $.FieldContextKey $msg .Name)}}
                            "{{.Name}}": &{{$.gqlpkg}}.InputObjectFieldConfig{
                                {{ if .Type.Array }}
									Type: {{$.gqlpkg}}.NewList({{$.gqlpkg}}.NewNonNull({{call $.GQLInputTypeName .Type}})),
								{{ else }}
									Type: {{call $.GQLInputTypeName .Type}},
								{{ end }}
                            },
                        {{end}}
                    {{end}}
					{{range $msg.MapFields}}
                        {{if not (call $.FieldContextKey $msg .Name)}}
                            "{{.Name}}": &{{$.gqlpkg}}.InputObjectFieldConfig{
                                {{ if .Type.Array }}
									Type: {{$.gqlpkg}}.NewList({{$.gqlpkg}}.NewNonNull({{call $.GQLInputTypeName .Type}})),
								{{ else }}
									Type: {{call $.GQLInputTypeName .Type}},
								{{ end }}
                            },
                        {{end}}
                    {{end}}
                    {{range .OneOffs -}}
                        {{range .Fields -}}
                            "{{.Name}}": &{{$.gqlpkg}}.InputObjectFieldConfig{
                            	{{ if .Type.Array }}
									Type: {{$.gqlpkg}}.NewList({{$.gqlpkg}}.NewNonNull({{call $.GQLInputTypeName .Type}})),
								{{ else }}
									Type: {{call $.GQLInputTypeName .Type}},
								{{ end }}
                        	},
                        {{end}}
                    {{end}}
                }
            }),
        })
		// Output msg resolver
		func {{call $.GQLOutputTypeResolver .Type}}({{ if $.tracerEnabled }} tr {{$.tracerpkg}}.Tracer, {{end}}ctx {{$.ctxpkg}}.Context, i interface{}) (_ *{{call $.GoType $msg.Type}}, rerr error){
			{{ if $.tracerEnabled}}
				span := tr.CreateChildSpanFromContext(ctx, "{{call $.GQLOutputTypeResolver .Type}}")
				defer span.Finish()
				defer func(){
					if perr := recover(); perr != nil {
						span.SetTag("error", "true")
						span.SetTag("error_message", perr)
						span.SetTag("error_stack", string({{$.debugpkg}}.Stack()))
					}
					if rerr != nil {
						span.SetTag("error", "true")
						span.SetTag("error_message", rerr.Error())
					}
				}()
			{{end}}
			if i == nil {
					return nil, nil
			}
			args := i.(map[string]interface{})
			_ = args
			var result = new({{call $.GoType .Type}})
			{{range $field := $msg.Fields -}}
				{{ $ctxkey := (call $.FieldContextKey $msg .Name) }}
				{{- if $ctxkey }}
					{{.Name}}_ctx := ctx.Value("{{$ctxkey}}")
					if {{.Name}}_ctx != nil{
						if val, ok := {{.Name}}_ctx.({{call $.GoType .Type}}); ok {
							result.{{call $.ccase .Name}} = val
						}else{
							panic("bad value type for field {{$msg.Name}}.{{.Name}}. Should be {{call $.GoType .Type}}, found: "+ {{$.fmtpkg}}.Sprintf("%T", {{.Name}}_ctx))
						}
					}
				{{- else -}}
					{{- if .Repeated -}}
						{{- if .Type.IsMessage -}}
							// Repeated Message
							if args["{{$field.Name}}"] != nil {
								var {{$field.Name}}_list = args["{{$field.Name}}"].([]interface{})
								var {{$field.Name}}_  = make([]*{{ call $.GoType $field.Type}}, len({{$field.Name}}_list))
								for i, {{$field.Name}}_item := range {{$field.Name}}_list {
									{{ if $.tracerEnabled }}
										{{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(tr,  tr.ContextWithSpan(ctx, span), {{$field.Name}}_item)
									{{ else }}
										{{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(ctx, {{$field.Name}}_item)
									{{ end }}
									if err != nil {
										return nil, {{$.errorspkg}}.New("failed to parse {{$field.Name}}["+{{$.strconvpkg}}.Itoa(i)+"]: " + err.Error())
									}
									{{$field.Name}}_[i] = {{$field.Name}}_r
								}
								result.{{call $.ccase .Name}} = {{$field.Name}}_
							}
						{{- else if .Type.IsMap -}}
							// Map
							{{ if $.tracerEnabled }}
								{{$field.Name}}_, err := {{call $.GQLOutputTypeResolver $field.Type}}(tr,  tr.ContextWithSpan(ctx, span), args["{{.Name}}"])
							{{ else }}
								{{$field.Name}}_, err := {{call $.GQLOutputTypeResolver $field.Type}}(ctx, args["{{.Name}}"])
							{{ end }}
							if err != nil {
								return nil, {{$.errorspkg}}.New("failed to parse {{$field.Name}}: " + err.Error())
							}
							if {{$field.Name}}_ != nil {
								result.{{call $.ccase .Name}} = {{$field.Name}}_
							}
						{{- else if .Type.IsEnum -}}
							// Repeated Enum
							if args["{{$field.Name}}"] != nil {
								var {{$field.Name}}_list = args["{{$field.Name}}"].([]interface{})
								var {{$field.Name}}_  = make([]{{call $.GoType $field.Type}}, len({{$field.Name}}_list))
								for i, {{$field.Name}}_item := range {{$field.Name}}_list {
									{{$field.Name}}_r, ok := {{$field.Name}}_item.(int)
									if !ok {
										return nil, {{$.errorspkg}}.New("failed to parse {{$field.Name}}["+{{$.strconvpkg}}.Itoa(i)+"]")
									}
									{{$field.Name}}_[i] = {{ call $.GoType $field.Type}}({{$field.Name}}_r)
								}
								result.{{call $.ccase .Name}} = {{$field.Name}}_
							}
						{{- else -}}
							// Repeated Scalar type
							if args["{{$field.Name}}"] != nil {
								var {{$field.Name}}_list = args["{{$field.Name}}"].([]interface{})
								var {{$field.Name}}_  = make([]{{call $.GoType $field.Type}}, len({{$field.Name}}_list))
								for i, {{$field.Name}}_item := range {{$field.Name}}_list {
									{{$field.Name}}_r, ok := {{$field.Name}}_item.({{ call $.GoType $field.Type}})
									if !ok {
										return nil, {{$.errorspkg}}.New("failed to parse {{$field.Name}}["+{{$.strconvpkg}}.Itoa(i)+"]")
									}
									{{$field.Name}}_[i] = {{$field.Name}}_r
								}
								result.{{call $.ccase .Name}} = {{$field.Name}}_
							}
						{{- end -}}
					{{- else -}}
						{{- if $field.Type.IsScalar -}}
							// Non-repeated scalar
							if args["{{.Name}}"] != nil {
								result.{{call $.ccase  .Name}} = args["{{.Name}}"].({{call $.GoType .Type}})
							}
						{{- else if $field.Type.IsMessage -}}
							// Non-repeated message

							// {{$field.Type.Message.Name}}
                            {{ if $field.Type.Message.HaveFields}}
                                {{ if $.tracerEnabled }}
                                    {{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(tr,  tr.ContextWithSpan(ctx, span), args["{{.Name}}"])
                                {{ else }}
                                    {{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(ctx, args["{{.Name}}"])
                                {{ end }}
                                if err != nil {
                                    return nil, {{$.errorspkg}}.New("failed to parse {{$field.Name}}: " + err.Error())
                                }
                                result.{{call $.ccase  .Name}} = {{$field.Name}}_r
                            {{ else -}}
                                result.{{call $.ccase  .Name}} = new({{call $.GoType $field.Type}})
                            {{ end -}}

						{{- else if $field.Type.IsEnum -}}
							// Non-repeated enum
							if args["{{.Name}}"] != nil {
								result.{{call $.ccase  .Name}} = {{call $.GoType .Type}}(args["{{.Name}}"].(int))
							}
						{{end -}}
					{{ end}}
				{{end -}}
			{{- end}}
			{{range $oneoff := $msg.OneOffs -}}
					//Generated oneoff
					{{range $index, $field := .Fields}}
						{{- if eq $index 0 -}}
						if {{.Name}}_, ok := args["{{.Name}}"]; ok && {{.Name}}_ != nil {
							{{- if $field.Type.IsScalar -}}
								// Non-repeated scalar
								result.{{call $.ccase  $oneoff.Name}} = &{{call $.GoType $msg.Type}}_{{call $.ccase  .Name}}{{"{"}}{{.Name}}_.({{call $.GoType .Type}})}
							{{- else if $field.Type.IsMessage -}}
								// Non-repeated message
								{{ if $.tracerEnabled }}
									{{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(tr,  tr.ContextWithSpan(ctx, span), {{.Name}}_)
								{{ else }}
									{{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(ctx, {{.Name}}_)
								{{ end }}
								if err != nil {
									return nil, {{$.errorspkg}}.New("failed to parse {{$field.Name}}: " + err.Error())
								}
								result.{{call $.ccase  $oneoff.Name}} = &{{call $.GoType $msg.Type}}_{{call $.ccase  .Name}}{{"{"}}{{$field.Name}}_r}
							{{- else if $field.Type.IsEnum -}}
								// Non-repeated enum
								result.{{call $.ccase  $oneoff.Name}} = &{{call $.GoType $msg.Type}}_{{call $.ccase  .Name}}{{"{"}}{{call $.GoType .Type}}({{.Name}}_.(int))}
							{{end -}}
						}{{- else -}} else if {{.Name}}_, ok := args["{{.Name}}"]; ok && {{.Name}}_ != nil {
							{{- if $field.Type.IsScalar -}}
								// Non-repeated scalar
								result.{{call $.ccase  $oneoff.Name}} = &{{call $.GoType $msg.Type}}_{{call $.ccase  .Name}}{{"{"}}{{.Name}}_.({{call $.GoType .Type}})}
							{{- else if $field.Type.IsMessage -}}
								// Non-repeated message
								{{ if $.tracerEnabled }}
									{{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(tr,  tr.ContextWithSpan(ctx, span), {{.Name}}_)
								{{ else }}
									{{$field.Name}}_r, err := {{call $.GQLOutputTypeResolver $field.Type}}(ctx, {{.Name}}_)
								{{ end }}

								if err != nil {
									return nil, {{$.errorspkg}}.New("failed to parse {{$field.Name}}: " + err.Error())
								}
								result.{{call $.ccase  $oneoff.Name}} = &{{call $.GoType $msg.Type}}_{{call $.ccase  .Name}}{{"{"}}{{$field.Name}}_r}
							{{- else if $field.Type.IsEnum -}}
								// Non-repeated enum
								result.{{call $.ccase  $oneoff.Name}} = &{{call $.GoType $msg.Type}}_{{call $.ccase  .Name}}{{"{"}}{{call $.GoType .Type}}({{.Name}}_.(int))}
							{{end -}}
						}{{- end -}}

					{{end}}
			{{end}}
			return result, nil
		}
    {{end -}}
{{end -}}


// Maps
{{ range $map := .File.Maps -}}
	{{ if $map.Message.OutputMessage }}
		{{ if not (call $.IsErrorField .Message .Field.Name) }}
    		var {{call $.GQLOutputTypeName .Type}} = {{$.gqlpkg}}.NewList({{$.gqlpkg}}.NewNonNull({{$.gqlpkg}}.NewObject({{$.gqlpkg}}.ObjectConfig{
    			Name:   "{{call $.GQLOutputTypeName .Type}}",
    			Fields: {{$.gqlpkg}}.Fields{
					"key": &{{$.gqlpkg}}.Field{
						Name: "key",
						Type: {{call $.GQLOutputTypeName $map.KeyType}},
						Resolve: func(p {{$.gqlpkg}}.ResolveParams) (interface{}, error) {
							src := p.Source.(map[string]interface{})
							if src == nil {
								return nil, nil
							}
							return src["key"].({{call $.GoType $map.KeyType}}), nil
						},
					},
					"value": &{{$.gqlpkg}}.Field{
						Name: "value",
						Type: {{call $.GQLOutputTypeName $map.ValueType}},
						Resolve: func(p {{$.gqlpkg}}.ResolveParams) (interface{}, error) {
							src := p.Source.(map[string]interface{})
							if src == nil {
								return nil, nil
							}
							return src["value"].({{call $.GoType $map.ValueType}}), nil
						},
					},
				},
    		})))
		{{end}}
	{{end}}
    {{ if $map.Message.InputMessage }}
		// {{$map.Message.Name}}
        var {{call $.GQLInputTypeName .Type}} = {{$.gqlpkg}}.NewList({{$.gqlpkg}}.NewNonNull({{$.gqlpkg}}.NewInputObject({{$.gqlpkg}}.InputObjectConfig{
            Name: "{{call $.GQLInputTypeName .Type}}",
            Fields: {{$.gqlpkg}}.InputObjectConfigFieldMap{
				"key": &{{$.gqlpkg}}.InputObjectFieldConfig{
					Type: {{call $.GQLInputTypeName .Type.Map.KeyType}},
				},
                "value": &{{$.gqlpkg}}.InputObjectFieldConfig{
                    Type: {{call $.GQLInputTypeName .Type.Map.ValueType}},
                },
            },
        })))
        func {{call $.GQLOutputTypeResolver .Type}}({{ if $.tracerEnabled }} tr {{$.tracerpkg}}.Tracer, {{end}}ctx {{$.ctxpkg}}.Context, i interface{}) (_ {{call $.GoType $map.Type}}, rerr error){
        	{{ if $.tracerEnabled }}
					span := tr.CreateChildSpanFromContext(ctx, "{{call $.GQLOutputTypeResolver .Type}}")
					defer span.Finish()
					defer func(){
						if perr := recover(); perr != nil {
							span.SetTag("error", "true")
							span.SetTag("error_message", perr)
							span.SetTag("error_stack", string({{$.debugpkg}}.Stack()))
						}
						if rerr != nil {
							span.SetTag("error", "true")
							span.SetTag("error_message", rerr.Error())
						}
					}()
				{{end}}
				if i == nil {
					return nil, nil
				}
				result := make({{call $.GoType $map.Type}})
				vals := i.([]interface{})
				{{ if $map.ValueType.IsMessage }}
					for iv, v := range vals {
						args := v.(map[string]interface{})
						{{ if $.tracerEnabled }}
							vv, err := {{call $.GQLOutputTypeResolver $map.ValueType}}(tr,  tr.ContextWithSpan(ctx, span), args["value"])
						{{ else }}
							vv, err := {{call $.GQLOutputTypeResolver $map.ValueType}}(ctx, args["value"])
						{{ end }}
						if err != nil {
							return nil, {{$.errorspkg}}.New("failed to parse {{call $.GQLOutputTypeName .Type}}["+{{$.strconvpkg}}.Itoa(iv)+"]: " + err.Error())
						}
						result[args["key"].({{call $.GoType $map.KeyType}})] = vv
					}
				{{ else if $map.ValueType.IsEnum }}
					for _, v := range vals {
						args := v.(map[string]interface{})
						result[args["key"].({{call $.GoType $map.KeyType}})] = {{call $.GoType $map.ValueType}}(args["value"].(int))
					}
				{{else}}
					for _, v := range vals {
						args := v.(map[string]interface{})
						result[args["key"].({{call $.GoType $map.KeyType}})] = args["value"].({{call $.GoType $map.ValueType}})
					}
				{{end}}


				return result, nil
		}
    {{ end }}
{{end}}

func init() {
	 // Adding fields to output messages
	{{ range $msg := .File.Messages -}}
		{{if $msg.OutputMessage}}
			// {{$msg.Name}} message fields
			{{range .Fields -}}
				{{ if not (call $.IsErrorField $msg .Name) -}}
					{{call $.GQLOutputTypeName $msg.Type}}.AddFieldConfig("{{.Name}}", &{{$.gqlpkg}}.Field{
						Name: "{{.Name}}",
						{{ if .Type.Array }}
							Type: {{$.gqlpkg}}.NewList({{$.gqlpkg}}.NewNonNull({{call $.GQLOutputTypeName .Type}})),
						{{ else }}
							Type: {{call $.GQLOutputTypeName .Type}},
						{{ end }}
						Resolve: func(p {{$.gqlpkg}}.ResolveParams) (interface{}, error) {
                            src := p.Source.(*{{call $.GoType $msg.Type}})
							if src == nil {
								return nil, nil
							}
							{{ if and  .Type.Array .Type.IsEnum }}
								source := src.{{call $.ccase  .Name}}
								var result = make([]int, len(source))
								for i, val := range source {
									result[i] = int(val)
								}
								return result, nil
							{{ else if .Type.IsMap }}
								var res []map[string]interface{}
								for k, v := range src.{{call $.ccase  .Name}} {
									res = append(res, map[string]interface{}{
										"key": k,
										"value": v,
									})
								}
								return res, nil
							{{ else if .Type.IsEnum }}
								return int(src.{{call $.ccase  .Name}}), nil
							{{ else }}
								return src.{{call $.ccase  .Name}}, nil
							{{- end -}}
						},
					})
				{{end -}}
			{{end}}
            {{ range .MapFields}}
				{{ if not (call $.IsErrorField $msg .Name) -}}
                	// Map field
                	{{ call $.GQLOutputTypeName $msg.Type }}.AddFieldConfig("{{.Name}}", &{{$.gqlpkg}}.Field{
                	    Name: "{{.Name}}",
                	    Type: {{call $.GQLOutputTypeName .Type}},
                	    Resolve: func(p {{$.gqlpkg}}.ResolveParams) (interface{}, error) {
                	        src := p.Source.(*{{ call $.GoType $msg.Type}})
                	        if src == nil {
                	            return nil, nil
                	        }
                	        var res []map[string]interface{}
                	        for k, v := range src.{{call $.ccase  .Name}} {
                	            res = append(res, map[string]interface{}{
                	                "key": k,
                	                "value": v,
                	            })
                	        }
                	        return res, nil
                	    },
                	})
				{{end}}
            {{end}}
			{{ range .OneOffs -}}
				{{ range .Fields -}}
                	// One OFF output
					{{ call $.GQLOutputTypeName $msg.Type }}.AddFieldConfig("{{.Name}}", &{{$.gqlpkg}}.Field{
						Name: "{{.Name}}",
						Type: {{call $.GQLOutputTypeName .Type}},
						Resolve: func(p {{$.gqlpkg}}.ResolveParams) (interface{}, error) {
							src := p.Source.(*{{ call $.GoType $msg.Type}})
							if src == nil {
								return nil, nil
							}
							{{if .Type.IsMap}}
								var res []map[string]interface{}
								for k, v := range src.{{call $.ccase  .Name}} {
									res = append(res, map[string]interface{}{
										"key": k,
										"value": v,
									})
								}
								return res, nil
							{{else if .Type.IsEnum}}
                            	return int(src.Get{{call $.ccase  .Name}}()), nil
							{{else}}
								return src.Get{{call $.ccase  .Name}}(), nil
							{{end}}
						},
					})
				{{end}}
			{{ end }}
		{{end -}}
	{{end -}}
}
{{ range $service := .File.Services }}
	func Get{{.Name}}GraphQLQueriesFields(c {{$.protoPkg}}.{{.Name}}Client, ih *{{$.interceptorspkg}}.InterceptorHandler {{ if $.tracerEnabled }} ,tr {{$.tracerpkg}}.Tracer {{end}}) {{$.gqlpkg}}.Fields {
		{{ if (call $.ServiceHaveQueries $service) }}
			return {{$.gqlpkg}}.Fields{
				{{ $serviceName := (call $.ServiceName $service) }}
				{{range $method := .Methods -}}
					{{ if (call $.MethodIsQuery $method) }}
						{{ $methodName := (call $.MethodName $method) }}
					` + methodTemplate + `
					{{end}}
				{{ end }}
			}
		{{else}}
			return nil
		{{end}}
	}

	func Get{{.Name}}GraphQLMutationsFields(c {{$.protoPkg}}.{{.Name}}Client, ih *{{$.interceptorspkg}}.InterceptorHandler {{ if $.tracerEnabled }} ,tr {{$.tracerpkg}}.Tracer {{end}}) {{$.gqlpkg}}.Fields {
		{{ if (call $.ServiceHaveMutations $service) }}
			return {{$.gqlpkg}}.Fields{
				{{ $serviceName := (call $.ServiceName $service) }}
				{{range $method := .Methods -}}
					{{ if not (call $.MethodIsQuery $method) }}
						{{ $methodName := (call $.MethodName $method) }}
						` + methodTemplate + `
					{{end}}
				{{ end }}
			}
		{{else}}
			return nil
		{{end}}
	}
{{ end }}
`

const methodTemplate = `
 				"{{$methodName}}": &{{$.gqlpkg}}.Field{
					Name: "{{call $.MethodName $method}}",
					Type: {{call $.GQLOutputTypeName .OutputMessage.Type}},
					Args: {{$.gqlpkg}}.FieldConfigArgument{
						{{range $field := $method.InputMessage.Fields -}}
							{{if not (call $.FieldContextKey $method.InputMessage .Name) -}}
								"{{.Name}}": &{{$.gqlpkg}}.ArgumentConfig{
									{{if ne .QuotedComment "\"\"" -}}
										Description: {{.QuotedComment}},
									{{end -}}
									{{ if .Type.Array -}}
										Type: {{$.gqlpkg}}.NewList({{$.gqlpkg}}.NewNonNull({{call $.GQLInputTypeName .Type}})),
									{{ else -}}
										Type: {{call $.GQLInputTypeName .Type}},
									{{ end -}}
								},
							{{end -}}
						{{end -}}
		                {{range $field := $method.InputMessage.MapFields -}}
							{{if not (call $.FieldContextKey $method.InputMessage .Name) -}}
								"{{.Name}}": &{{$.gqlpkg}}.ArgumentConfig{
									{{if ne .QuotedComment "\"\"" -}}
										Description: {{.QuotedComment}},
									{{end -}}
									Type: {{call $.GQLInputTypeName .Type}},
								},
							{{end -}}
						{{end -}}
						{{range $oneoff := $method.InputMessage.OneOffs -}}
							{{range $field := $oneoff.Fields -}}
								"{{.Name}}": &{{$.gqlpkg}}.ArgumentConfig{
									{{if ne .QuotedComment "\"\"" -}}
									Description: {{.QuotedComment}},
									{{- end}}
									Type: {{call $.GQLInputTypeName .Type}},
								},
							{{end -}}
						{{end -}}
					},
					Resolve: func(p {{$.gqlpkg}}.ResolveParams) (_ interface{}, rerr error) {       
						{{ if $.tracerEnabled }}	
							span := tr.CreateChildSpanFromContext(p.Context, "{{$serviceName}}.{{$methodName}} Resolver")
							defer span.Finish()
							defer func(){
								if rerr != nil {
									span.SetTag("error", "true")
									span.SetTag("error_message", rerr.Error())
								}
							}()
						{{end}}
						if ih == nil {
							{{ if .InputMessage.HaveFields }}
								{{ if $.tracerEnabled -}}	
									req, err := {{call $.GQLOutputTypeResolver .InputMessage.Type}}(tr, tr.ContextWithSpan(p.Context, span), p.Args)
								{{ else -}}
									req, err := {{call $.GQLOutputTypeResolver .InputMessage.Type}}(p.Context, p.Args)
								{{ end -}}
								if err != nil {
									return nil, err
								}
								return c.{{call $.ccase .Name}}(p.Context, req)
							{{ else }}
 								return c.{{call $.ccase .Name}}(p.Context, new({{call $.GoType .InputMessage.Type}}))
							{{ end }}
						}
						ctx := &{{$.interceptorspkg}}.Context{
							Service: "{{$serviceName}}",
							Method: "{{$methodName}}",
							Params: p,
						}
						req, err := ih.ResolveArgs(ctx, func(ctx *{{$.interceptorspkg}}.Context, next {{$.interceptorspkg}}.ResolveArgsInvoker) (result interface{}, err error) {
							{{ if .InputMessage.HaveFields }}
								{{ if $.tracerEnabled -}}	
									return {{call $.GQLOutputTypeResolver .InputMessage.Type}}(tr, tr.ContextWithSpan(p.Context, span), p.Args)
								{{ else -}}
									return {{call $.GQLOutputTypeResolver .InputMessage.Type}}(p.Context, p.Args)
								{{ end -}}
							{{ else }}
								return new({{call $.GoType .InputMessage.Type}}), nil
							{{ end }}
						})
						if err != nil {
							return nil, err
						}
						res, err := ih.Call(ctx, req, func(ctx *{{$.interceptorspkg}}.Context, req interface{}, next {{$.interceptorspkg}}.CallMethodInvoker, opts ...grpc.CallOption) (result interface{}, err error) {
							r, ok := req.(*{{call $.GoType .InputMessage.Type}})
							if !ok {
								return nil, {{$.errorspkg}}.New({{$.fmtpkg}}.Sprintf("resolve args interceptor returns bad request type(%T). Should be: *{{call $.GoType .InputMessage.Type}}", req))
							}
							res, err := c.{{call $.ccase .Name}}(ctx.Params.Context, r, opts...)
							{{$errFld := (call $.MessageErrorField $method.OutputMessage)}}
							{{if $errFld}}
								if res != nil {
									 ctx.PayloadError =  res.{{call $.ccase $errFld}}
								}
							{{end}}
							return res, err				
						})
						rc, ok :=res.(*{{call $.GoType .OutputMessage.Type}})
						if !ok {
							return nil, {{$.errorspkg}}.New({{$.fmtpkg}}.Sprintf("Resolve Interceptor returns bad value type(%T). Should return *{{call $.GoType .OutputMessage.Type}}", res))
						}
						return rc, err
					},
				},
`
