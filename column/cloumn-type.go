package column

import "fmt"

type Type int

const (
	// NULL type having no physical storage
	NULL Type = iota

	// BOOL is a 1 bit, LSB bit-packed ordering
	BOOL

	// UINT8 is an Unsigned 8-bit little-endian integer
	UINT8
	UINT4
	INT4
	// INT8 is a Signed 8-bit little-endian integer
	INT8

	// UINT16 is an Unsigned 16-bit little-endian integer
	UINT16

	// INT16 is a Signed 16-bit little-endian integer
	INT16

	// UINT32 is an Unsigned 32-bit little-endian integer
	UINT32

	// INT32 is a Signed 32-bit little-endian integer
	INT32

	// UINT64 is an Unsigned 64-bit little-endian integer
	UINT64

	// INT64 is a Signed 64-bit little-endian integer
	INT64

	// FLOAT16 is a 2-byte floating point value
	FLOAT16

	// FLOAT32 is a 4-byte floating point value
	FLOAT32

	// FLOAT64 is an 8-byte floating point value
	FLOAT64

	// DATE32 is int32 days since the UNIX epoch
	DATE

	TIME

	DATETIME

	YEAR
	// TIMESTAMP is an exact timestamp encoded with int64 since UNIX epoch
	// Default unit millisecond
	TIMESTAMP

	// INTERVAL_MONTHS is YEAR_MONTH interval in SQL style
	INTERVAL_YEAR_MONTHS

	// INTERVAL_DAY_TIME is DAY_TIME in SQL Style
	INTERVAL_DAY_TIME

	CHAR

	VARCHAR

	JSON

	JSONB

	TEXT

	BLOB

	CLOB

	NUMERIC

	DECIMAL

	BIT

	SET

	ENUM

	UNKNOWN

	REAL

	XML
)

type ColumnType interface {
	fmt.Stringer
	ID() Type
	Name() string
}

type Json struct{}

func (j *Json) ID() Type       { return JSON }
func (j *Json) Name() string   { return "json" }
func (t *Json) String() string { return "json" }

type Jsonb struct{}

func (j *Jsonb) ID() Type       { return JSONB }
func (j *Jsonb) Name() string   { return "jsonb" }
func (t *Jsonb) String() string { return "jsonb" }

type TextType struct{}

func (j *TextType) ID() Type       { return TEXT }
func (j *TextType) Name() string   { return "text" }
func (t *TextType) String() string { return "text" }

type ByteaType struct{}

func (j *ByteaType) ID() Type       { return TEXT }
func (j *ByteaType) Name() string   { return "bytea" }
func (t *ByteaType) String() string { return "bytea" }

type BooleanType struct{}

func (j *BooleanType) ID() Type       { return BOOL }
func (j *BooleanType) Name() string   { return "boolean" }
func (t *BooleanType) String() string { return "boolean" }

type Blob struct{}

func (j *Blob) ID() Type       { return BLOB }
func (j *Blob) Name() string   { return "blob" }
func (t *Blob) String() string { return "blob" }

type Clob struct{}

func (j *Clob) ID() Type       { return CLOB }
func (j *Clob) Name() string   { return "clob" }
func (t *Clob) String() string { return "clob" }

type SetType struct {
	Values []string
}

func (j *SetType) ID() Type       { return SET }
func (j *SetType) Name() string   { return "set" }
func (t *SetType) String() string { return "set" }

type RealType struct{}

func (j *RealType) ID() Type       { return REAL }
func (j *RealType) Name() string   { return "real" }
func (t *RealType) String() string { return "real" }

type XmlType struct{}

func (j *XmlType) ID() Type       { return XML }
func (j *XmlType) Name() string   { return "real" }
func (t *XmlType) String() string { return "real" }

type EnumType struct {
	Values []string
}

func (j *EnumType) ID() Type       { return SET }
func (j *EnumType) Name() string   { return "set" }
func (t *EnumType) String() string { return "set" }

var (
	PrimitiveTypes = struct {
		Json    ColumnType
		Jsonb   ColumnType
		Bytea   ColumnType
		Text    ColumnType
		Boolean ColumnType
		Blob    ColumnType
		Clob    ColumnType
		Real    ColumnType
		Xml     ColumnType
	}{
		Json:    &Json{},
		Jsonb:   &Jsonb{},
		Bytea:   &ByteaType{},
		Text:    &TextType{},
		Boolean: &BooleanType{},
		Blob:    &Blob{},
		Clob:    &Clob{},
		Real:    &RealType{},
		Xml:     &XmlType{},
	}
)
