package datatunnelcommon

// 范围是 -128 到 127（有符号）或 0 到 255（无符号），占用 1 字节。
type TinyintType struct {
	Unsigned bool
}

func (t *TinyintType) ID() Type       { return INT4 }
func (t *TinyintType) Name() string   { return "int4" }
func (t *TinyintType) String() string { return "int4" }

func (t *TinyintType) GetUnsigned() bool { return t.Unsigned }

// 范围是 -32768 到 32767（有符号）或 0 到 65535（无符号），占用 2 字节。
type SmallIntType struct {
	Unsigned bool
}

func (t *SmallIntType) ID() Type       { return INT8 }
func (t *SmallIntType) Name() string   { return "int8" }
func (t *SmallIntType) String() string { return "int8" }

func (t *SmallIntType) GetUnsigned() bool { return t.Unsigned }

// 中等大小的整数值，范围是 -8388608 到 8388607（有符号）或 0 到 16777215（无符号），占用 3 字节。
type MediumIntType struct {
	Unsigned bool
}

func (t *MediumIntType) ID() Type       { return INT16 }
func (t *MediumIntType) Name() string   { return "int16" }
func (t *MediumIntType) String() string { return "int16" }

func (t *MediumIntType) GetUnsigned() bool { return t.Unsigned }

type IntType struct {
	Unsigned bool
}

func (t *IntType) ID() Type       { return INT32 }
func (t *IntType) Name() string   { return "int32" }
func (t *IntType) String() string { return "int32" }

func (t *IntType) GetUnsigned() bool { return t.Unsigned }

type BigIntType struct {
	Unsigned bool
}

func (t *BigIntType) ID() Type       { return INT64 }
func (t *BigIntType) Name() string   { return "int64" }
func (t *BigIntType) String() string { return "int64" }

func (t *BigIntType) GetUnsigned() bool { return t.Unsigned }

type DoubleType struct{}

func (t *DoubleType) ID() Type       { return FLOAT64 }
func (t *DoubleType) Name() string   { return "double" }
func (t *DoubleType) String() string { return "double" }

type FloatType struct{}

func (t *FloatType) ID() Type       { return FLOAT32 }
func (t *FloatType) Name() string   { return "float" }
func (t *FloatType) String() string { return "float" }

// var (
// 	DigitTypes = struct {
// 		Int4  ColumnType
// 		Int8  ColumnType
// 		Int16 ColumnType
// 		Int32 ColumnType
// 		Int64 ColumnType
// 	}{
// 		Int4:  &Int4Type{},
// 		Int8:  &Int8Type{},
// 		Int16: &Int16Type{},
// 		Int32: &Int32Type{},
// 		Int64: &Int64Type{},
// 	}
// )

type DecimalType struct {
	PrecisionNum int
	ScaleNum     int
}

func (t *DecimalType) ID() Type       { return DECIMAL }
func (t *DecimalType) Name() string   { return "decimal" }
func (t *DecimalType) String() string { return "decimal" }

func (t *DecimalType) Precision() int { return t.PrecisionNum }

func (t *DecimalType) Scale() int { return t.ScaleNum }
