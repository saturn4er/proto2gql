package parser

type ProtoType struct {
	File     *File
	Message  *Message
	Enum     *Enum
	Scalar   string
	Map      *Map
	Array    bool
	Optional bool
}

func (p *ProtoType) IsScalar() bool {
	return p.Scalar != ""
}
func (p *ProtoType) IsMessage() bool {
	return p.Message != nil
}

func (p *ProtoType) IsEnum() bool {
	return p.Enum != nil
}

func (p *ProtoType) IsMap() bool {
	return p.Map != nil
}

func (p *ProtoType) String() string {
	var res string
	if p.Array {
		res += "repeated "
	}
	if p.Optional {
		res += "optional "
	}
	switch {
	case p.IsMessage():
		res += p.Message.Name + " message"
	case p.IsMap():
		res += p.Map.Message.Name + "." + p.Map.Field.Name + " map"
	case p.IsEnum():
		res += p.Enum.Name + " enum"
	case p.IsScalar():
		res += p.Scalar
	default:
		return "unknown type"
	}
	return res
}
