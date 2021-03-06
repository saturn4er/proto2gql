package parser

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func testFileInfo(file *File) *File {
	var Int32Type = &Type{Scalar: "int32"}
	var RootMessage = file.Messages[0]
	var RootMessage2 = file.Messages[4]
	var RootMessage2Type = &Type{File: file, Message: RootMessage2}
	var RootEnumType = &Type{File: file, Enum: file.Enums[0]}
	var EmptyMessage = file.Messages[2]
	var EmptyMessageType = &Type{File: file, Message: EmptyMessage}
	var NestedMessage = file.Messages[1]
	var NestedMessageType = &Type{File: file, Message: file.Messages[1]}
	var NestedEnumType = &Type{File: file, Enum: file.Enums[1]}
	var NestedNestedEnumType = &Type{File: file, Enum: file.Enums[2]}
	var MessageWithEmpty = file.Messages[3]

	var CommonCommonEnumType = &Type{File: file.Imports[0], Enum: file.Imports[0].Enums[0]}
	var CommonCommonMessageType = &Type{File: file.Imports[0], Message: file.Imports[0].Messages[0]}

	return &File{
		Services: []*Service{
			{
				Name:          "ServiceExample",
				QuotedComment: `"Service, which do smth"`,
				Methods: []*Method{
					{Name: "getQueryMethod", InputMessage: RootMessage, OutputMessage: RootMessage, QuotedComment: `""`},
					{Name: "mutationMethod", InputMessage: RootMessage2, OutputMessage: NestedMessage, QuotedComment: `"rpc comment"`},
					{Name: "EmptyMsgs", InputMessage: EmptyMessage, OutputMessage: EmptyMessage, QuotedComment: `""`},
					{Name: "MsgsWithEpmty", InputMessage: MessageWithEmpty, OutputMessage: MessageWithEmpty, QuotedComment: `""`},
				},
			},
		},
		Messages: []*Message{
			{
				file:          file,
				Name:          "RootMessage",
				QuotedComment: `""`,
				TypeName:      TypeName{"RootMessage"},
				Fields: []*Field{
					{Name: "r_msg", Type: NestedMessageType, Repeated: true, QuotedComment: `"repeated Message"`},
					{Name: "r_scalar", Type: Int32Type, Repeated: true, QuotedComment: `"repeated Scalar"`},
					{Name: "r_enum", Type: RootEnumType, Repeated: true, QuotedComment: `"repeated Enum"`},
					{Name: "r_empty_msg", Type: EmptyMessageType, Repeated: true, QuotedComment: `"repeated empty message"`},
					{Name: "n_r_enum", Type: CommonCommonEnumType, QuotedComment: `"non-repeated Enum"`},
					{Name: "n_r_scalar", Type: Int32Type, QuotedComment: `"non-repeated Scalar"`},
					{Name: "n_r_msg", Type: CommonCommonMessageType, QuotedComment: `"non-repeated Message"`},
					{Name: "scalar_from_context", Type: Int32Type, QuotedComment: `"field from context"`},
					{Name: "n_r_empty_msg", Type: EmptyMessageType, QuotedComment: `"non-repeated empty message field"`},
				},
				OneOffs: []*OneOf{
					{Name: "enum_first_oneoff", Fields: []*Field{
						{Name: "e_f_o_e", Type: CommonCommonEnumType, QuotedComment: `""`},
						{Name: "e_f_o_s", Type: Int32Type, QuotedComment: `""`},
						{Name: "e_f_o_m", Type: CommonCommonMessageType, QuotedComment: `""`},
						{Name: "e_f_o_em", Type: EmptyMessageType, QuotedComment: `""`},
					}},
					{Name: "scalar_first_oneoff", Fields: []*Field{
						{Name: "s_f_o_s", Type: Int32Type, QuotedComment: `""`},
						{Name: "s_f_o_e", Type: RootEnumType, QuotedComment: `""`},
						{Name: "s_f_o_mes", Type: RootMessage2Type, QuotedComment: `""`},
						{Name: "s_f_o_m", Type: EmptyMessageType, QuotedComment: `""`},
					}},
					{Name: "message_first_oneoff", Fields: []*Field{
						{Name: "m_f_o_m", Type: RootMessage2Type, QuotedComment: `""`},
						{Name: "m_f_o_s", Type: Int32Type, QuotedComment: `""`},
						{Name: "m_f_o_e", Type: RootEnumType, QuotedComment: `""`},
						{Name: "m_f_o_em", Type: EmptyMessageType, QuotedComment: `""`},
					}},
					{Name: "empty_first_oneoff", Fields: []*Field{
						{Name: "em_f_o_em", Type: EmptyMessageType, QuotedComment: `""`},
						{Name: "em_f_o_s", Type: Int32Type, QuotedComment: `""`},
						{Name: "em_f_o_en", Type: RootEnumType, QuotedComment: `""`},
						{Name: "em_f_o_m", Type: RootMessage2Type, QuotedComment: `""`},
					}},
				},
				MapFields: []*MapField{
					{
						Name:          "map_enum",
						QuotedComment: `"enum_map"`,
						Type: &Type{
							File: file,
							Map: &Map{
								Message:   RootMessage,
								KeyType:   Int32Type,
								ValueType: NestedEnumType,
							},
						},
					},
					{
						Name:          "map_scalar",
						QuotedComment: `"scalar map"`,
						Type: &Type{
							File: file,
							Map: &Map{
								Message:   RootMessage,
								KeyType:   Int32Type,
								ValueType: Int32Type,
							},
						},
					},
					{
						Name:          "map_msg",
						QuotedComment: `""`,
						Type: &Type{
							File: file,
							Map: &Map{
								Message:   RootMessage,
								KeyType:   Int32Type,
								ValueType: NestedMessageType,
							},
						},
					},
				},
			},
			{
				file:          file,
				Name:          "NestedMessage",
				QuotedComment: `""`,
				TypeName:      TypeName{"RootMessage", "NestedMessage"},
				Fields: []*Field{
					{Name: "sub_r_enum", Type: NestedEnumType, Repeated: true, QuotedComment: `""`},
					{Name: "sub_sub_r_enum", Type: NestedNestedEnumType, Repeated: true, QuotedComment: `""`},
				},
			},
			{
				file:          file,
				Name:          "Empty",
				QuotedComment: `""`,
				TypeName:      TypeName{"Empty"},
			},
			{
				file:          file,
				Name:          "MessageWithEmpty",
				QuotedComment: `""`,
				TypeName:      TypeName{"MessageWithEmpty"},
				Fields: []*Field{
					{Name: "empt", Type: EmptyMessageType, QuotedComment: `""`},
				},
			},
			{
				file:          file,
				Name:          "RootMessage2",
				QuotedComment: `""`,
				TypeName:      TypeName{"RootMessage2"},
				Fields: []*Field{
					{Name: "some_field", Type: Int32Type, QuotedComment: `""`},
				},
			},
		},
		Enums: []*Enum{
			{
				file:          file,
				Name:          "RootEnum",
				QuotedComment: `""`,
				TypeName:      TypeName{"RootEnum"},
				Values: []*EnumValue{
					{Name: "RootEnumVal0", Value: 0, QuotedComment: `""`},
					{Name: "RootEnumVal1", Value: 1, QuotedComment: `""`},
					{Name: "RootEnumVal2", Value: 2, QuotedComment: `"It's a RootEnumVal2"`},
				},
			},
			{
				file:          file,
				Name:          "NestedEnum",
				QuotedComment: `""`,
				TypeName:      TypeName{"RootMessage", "NestedEnum"},
				Values: []*EnumValue{
					{Name: "NestedEnumVal0", Value: 0, QuotedComment: `""`},
					{Name: "NestedEnumVal1", Value: 1, QuotedComment: `""`},
				},
			},
			{
				file:          file,
				Name:          "NestedNestedEnum",
				QuotedComment: `""`,
				TypeName:      TypeName{"RootMessage", "NestedMessage", "NestedNestedEnum"},
				Values: []*EnumValue{
					{Name: "NestedNestedEnumVal0", Value: 0, QuotedComment: `""`},
					{Name: "NestedNestedEnumVal1", Value: 1, QuotedComment: `""`},
					{Name: "NestedNestedEnumVal2", Value: 2, QuotedComment: `""`},
					{Name: "NestedNestedEnumVal3", Value: 3, QuotedComment: `""`},
				},
			},
		},
	}
}

func TestParser_Parse(t *testing.T) {
	Convey("Test Parser.Parse", t, func() {
		parser := New(map[string]string{"common/commo.proto": "common/common.proto"}, []string{"../testdata"})
		test, err := parser.Parse("../testdata/test.proto")
		So(err, ShouldBeNil)
		So(test, ShouldNotBeNil)
		test2, err := parser.Parse("../testdata/test2.proto")
		So(err, ShouldBeNil)
		So(test2, ShouldNotBeNil)
		So(test, ShouldNotEqual, test2)

		Convey("Imports should be the same", func() {
			So(len(test.Imports), ShouldEqual, 1)
			So(len(test2.Imports), ShouldEqual, 1)
			So(test.Imports[0], ShouldEqual, test2.Imports[0])
		})
		Convey("If we trying to parse same file, it should return pointer to parsed one", func() {
			test22, err := parser.Parse("../testdata/test2.proto")
			So(err, ShouldBeNil)
			So(test22, ShouldEqual, test2)
		})
		f := testFileInfo(test)
		Convey("test.proto Should contains valid enums", func() {
			So(test.Enums, ShouldHaveLength, len(f.Enums))
			for i, enum := range test.Enums {
				validEnum := f.Enums[i]
				Convey("Should contain "+validEnum.Name, func() {
					So(enum.file, ShouldEqual, validEnum.file)
					So(enum.Name, ShouldEqual, validEnum.Name)
					So(enum.Type.Enum, ShouldEqual, enum)
					So(enum.Type.File, ShouldEqual, test)
					So(enum.TypeName, ShouldResemble, validEnum.TypeName)
					So(enum.QuotedComment, ShouldEqual, validEnum.QuotedComment)
					Convey(validEnum.Name+" enum should contains valid values", func() {
						So(enum.Values, ShouldHaveLength, len(validEnum.Values))
						for i, value := range enum.Values {
							validValue := validEnum.Values[i]
							Convey(validEnum.Name+" enum should contains valid "+validValue.Name+" value", func() {
								So(value.Name, ShouldEqual, validValue.Name)
								So(value.Value, ShouldEqual, validValue.Value)
								So(value.QuotedComment, ShouldEqual, validValue.QuotedComment)
							})
						}
					})
				})
			}
		})

		Convey("test.proto Should contains valid messages", func() {
			So(test.Messages, ShouldHaveLength, len(f.Messages))
			for i, msg := range test.Messages {
				validMsg := f.Messages[i]
				Convey("Should have valid parsed "+strings.Join(validMsg.TypeName, "_")+" message ", func() {
					So(msg.file, ShouldEqual, validMsg.file)
					So(msg.Name, ShouldEqual, validMsg.Name)
					So(msg.Type.Message, ShouldEqual, msg)
					So(msg.Type.File, ShouldEqual, test)
					So(msg.TypeName, ShouldResemble, validMsg.TypeName)
					So(msg.QuotedComment, ShouldEqual, validMsg.QuotedComment)
					So(msg.Fields, ShouldHaveLength, len(validMsg.Fields))
					for i, fld := range msg.Fields {
						validFld := validMsg.Fields[i]
						Convey("Should have valid parsed "+strings.Join(validMsg.TypeName, "_")+"."+validFld.Name+" field", func() {
							So(fld.Name, ShouldEqual, validFld.Name)
							So(fld.Repeated, ShouldEqual, validFld.Repeated)
							So(fld.QuotedComment, ShouldEqual, validFld.QuotedComment)
							CompareTypes(fld.Type, validFld.Type)
						})
					}
					So(msg.MapFields, ShouldHaveLength, len(validMsg.MapFields))
					for i, fld := range msg.MapFields {
						validFld := validMsg.MapFields[i]
						Convey("Should have valid parsed "+strings.Join(validMsg.TypeName, "_")+"."+validFld.Name+" field", func() {
							So(fld.Name, ShouldEqual, validFld.Name)
							So(fld.QuotedComment, ShouldEqual, validFld.QuotedComment)
							CompareTypes(fld.Type, validFld.Type)
						})
					}
					So(msg.OneOffs, ShouldHaveLength, len(validMsg.OneOffs))
					for i, oneOf := range msg.OneOffs {
						validOneOf := validMsg.OneOffs[i]
						Convey("Should have valid parsed "+strings.Join(validMsg.TypeName, "_")+"."+validOneOf.Name+" one of", func() {
							So(oneOf.Name, ShouldEqual, validOneOf.Name)
							So(oneOf.Fields, ShouldHaveLength, len(validOneOf.Fields))
							for i, fld := range oneOf.Fields {
								validFld := validOneOf.Fields[i]
								Convey("Should have valid parsed "+strings.Join(validMsg.TypeName, "_")+"."+validOneOf.Name+"."+validFld.Name+" one of field", func() {
									So(fld.Name, ShouldEqual, validFld.Name)
									So(fld.QuotedComment, ShouldEqual, validFld.QuotedComment)
									CompareTypes(fld.Type, validFld.Type)
								})
							}

						})
					}
				})

			}
		})
		Convey("test.proto Should contain valid services", func() {
			So(test.Services, ShouldHaveLength, len(f.Services))
			for i, srv := range test.Services {
				validSrv := f.Services[i]
				Convey("Should have valid parsed "+validSrv.Name+" service ", func() {
					So(srv.Name, ShouldEqual, validSrv.Name)
					So(srv.QuotedComment, ShouldEqual, validSrv.QuotedComment)
					Convey(validSrv.Name+" should contains valid methods", func() {
						So(srv.Methods, ShouldHaveLength, len(validSrv.Methods))
						for i, method := range srv.Methods {
							validMethod := validSrv.Methods[i]
							Convey(validSrv.Name+" should contains valid "+validMethod.Name+" method", func() {
								So(method.Name, ShouldEqual, validMethod.Name)
								So(method.QuotedComment, ShouldEqual, validMethod.QuotedComment)
								Convey(validSrv.Name+"."+validMethod.Name+" should have valid input message type", func() {
									CompareTypes(method.InputMessage.Type, validMethod.InputMessage.Type)
								})
								Convey(validSrv.Name+"."+validMethod.Name+" should have valid output message type", func() {
									CompareTypes(method.OutputMessage.Type, validMethod.OutputMessage.Type)
								})
							})
						}
					})
				})
			}
		})
	})
}

func CompareTypes(t1, t2 *Type) {
	So(t1, ShouldNotBeNil)
	So(t2, ShouldNotBeNil)
	if t1.IsScalar() {
		So(t1.Scalar, ShouldEqual, t2.Scalar)
	}
	if t1.IsMessage() {
		So(t1.Message, ShouldEqual, t2.Message)
		So(t1.File, ShouldEqual, t2.File)
	}
	if t1.IsEnum() {
		So(t1.Enum, ShouldEqual, t2.Enum)
		So(t1.File, ShouldEqual, t2.File)
	}
	if t1.IsMap() {
		So(t1.Map.Message, ShouldEqual, t2.Map.Message)
		CompareTypes(t1.Map.KeyType, t2.Map.KeyType)
		CompareTypes(t1.Map.ValueType, t2.Map.ValueType)
		So(t1.File, ShouldEqual, t2.File)
	}
}
