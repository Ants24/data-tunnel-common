package column

// type ColumnStringType interface {
// 	ColumnType
// 	Len() int
// }

type CharType struct {
	Length int
}

func (t *CharType) ID() Type       { return CHAR }
func (t *CharType) Name() string   { return "char" }
func (t *CharType) String() string { return "char" }
func (t *CharType) Len() int       { return t.Length }

type VarCharType struct {
	Length int
}

func (t *VarCharType) ID() Type       { return CHAR }
func (t *VarCharType) Name() string   { return "varchar" }
func (t *VarCharType) String() string { return "varchar" }
func (t *VarCharType) Len() int       { return t.Length }

type BitType struct {
	Length int
}

func (t *BitType) ID() Type       { return BIT }
func (t *BitType) Name() string   { return "bit" }
func (t *BitType) String() string { return "bit" }
func (t *BitType) Len() int       { return t.Length }

type BitVaryingType struct {
	Length int
}

func (t *BitVaryingType) ID() Type       { return BIT }
func (t *BitVaryingType) Name() string   { return "bit" }
func (t *BitVaryingType) String() string { return "bit" }
func (t *BitVaryingType) Len() int       { return t.Length }
