package parser

type Type struct {
	File    *File
	Message *Message
	Enum    *Enum
	Scalar  string
	Map     *Map
}

func (p *Type) IsScalar() bool {
	return p.Scalar != ""
}
func (p *Type) IsMessage() bool {
	return p.Message != nil
}

func (p *Type) IsEnum() bool {
	return p.Enum != nil
}

func (p *Type) IsMap() bool {
	return p.Map != nil
}

func (p *Type) String() string {
	switch {
	case p.IsMessage():
		return p.Message.Name + " message"
	case p.IsMap():
		return p.Map.Message.Name + "." + p.Map.Field.Name + " map"
	case p.IsEnum():
		return p.Enum.Name + " enum"
	case p.IsScalar():
		return p.Scalar
	default:
		return "unknown type"
	}
}
