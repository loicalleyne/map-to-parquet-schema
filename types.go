package main

// avro 1.11.1 specification https://avro.apache.org/docs/1.11.1/specification/_print/

type AvroType int32

const (
	avroType_Null                 AvroType = 0
	avroType_Boolean              AvroType = 1
	avroType_Int                  AvroType = 2
	avroType_Long                 AvroType = 4
	avroType_Float                AvroType = 8
	avroType_Double               AvroType = 16
	avroType_Bytes                AvroType = 32
	avroType_String               AvroType = 64
	avroType_Record               AvroType = 128
	avroType_Enum                 AvroType = 256
	avroType_Array                AvroType = 512
	avroType_Map                  AvroType = 1024
	avroType_Union                AvroType = 2048
	avroType_Fixed                AvroType = 4096
	avroType_Decimal              AvroType = 8192
	avroType_UUID                 AvroType = 16384
	avroType_Date                 AvroType = 32768
	avroType_TimeMillis           AvroType = 65536
	avroType_TimeMicros           AvroType = 131072
	avroType_TimestampMillis      AvroType = 262144
	avroType_TimestampMicros      AvroType = 524288
	avroType_LocalTimestampMillis AvroType = 1048576
	avroType_LocalTimestampMicros AvroType = 2097152
	avroType_Duration             AvroType = 4194304
)

var AvroTypeName = map[AvroType]string{
	0:       "null",
	1:       "boolean",
	2:       "int",
	4:       "long",
	8:       "float",
	16:      "double",
	32:      "bytes",
	64:      "string",
	128:     "record",
	256:     "enum",
	512:     "array",
	1024:    "map",
	2048:    "union",
	4096:    "fixed",
	8192:    "decimal",
	16384:   "uuid",
	32768:   "date",
	65536:   "time-millis",
	131072:  "time-micros",
	262144:  "timestamp-millis",
	524288:  "timestamp-micros",
	1048576: "local-timestamp-millis",
	2097152: "local-timestamp-micros",
	4194304: "duration",
}

var AvroTypeValue = map[string]AvroType{
	"null":                   0,
	"boolean":                1,
	"int":                    2,
	"long":                   4,
	"float":                  8,
	"double":                 16,
	"bytes":                  32,
	"string":                 64,
	"record":                 128,
	"enum":                   256,
	"array":                  512,
	"map":                    1024,
	"union":                  2048,
	"fixed":                  4096,
	"decimal":                8192,
	"uuid":                   16384,
	"date":                   32768,
	"time-millis":            65536,
	"time-micros":            131072,
	"timestamp-millis":       262144,
	"timestamp-micros":       524288,
	"local-timestamp-millis": 1048576,
	"local-timestamp-micros": 2097152,
	"duration":               4194304,
}

type ParquetType int32

const (
	parquetType_BOOLEAN              ParquetType = 0
	parquetType_INT32                ParquetType = 1
	parquetType_INT64                ParquetType = 4
	parquetType_INT96                ParquetType = 8 // deprecated, only used by legacy implementations.
	parquetType_FLOAT                ParquetType = 16
	parquetType_DOUBLE               ParquetType = 32
	parquetType_BYTE_ARRAY           ParquetType = 64
	parquetType_FIXED_LEN_BYTE_ARRAY ParquetType = 128
)

var ParquetTypeName = map[ParquetType]string{
	0:   "BOOLEAN",
	1:   "INT32",
	4:   "INT64",
	8:   "INT96", // deprecated
	16:  "FLOAT",
	32:  "DOUBLE",
	64:  "BYTE_ARRAY",
	128: "FIXED_LEN_BYTE_ARRAY",
}

var ParquetTypeValue = map[string]ParquetType{
	"BOOLEAN":              0,
	"INT32":                1,
	"INT64":                4,
	"INT96":                8, // deprecated
	"FLOAT":                16,
	"DOUBLE":               32,
	"BYTE_ARRAY":           64,
	"FIXED_LEN_BYTE_ARRAY": 128,
}

type ParquetLogicalType int32

const (
	parquetLogicalType_STRING  ParquetLogicalType = 0  // use ConvertedType UTF8
	parquetLogicalType_MAP     ParquetLogicalType = 1  // use ConvertedType MAP
	parquetLogicalType_LIST    ParquetLogicalType = 2  // use ConvertedType LIST
	parquetLogicalType_ENUM    ParquetLogicalType = 4  // use ConvertedType ENUM
	parquetLogicalType_DECIMAL ParquetLogicalType = 8  // use ConvertedType DECIMAL + SchemaElement.{scale, precision}
	parquetLogicalType_DATE    ParquetLogicalType = 16 // use ConvertedType DATE
	// use ConvertedType TIME_MICROS for TIME(isAdjustedToUTC = *, unit = MICROS)
	// use ConvertedType TIME_MILLIS for TIME(isAdjustedToUTC = *, unit = MILLIS)
	parquetLogicalType_TIME ParquetLogicalType = 32
	// use ConvertedType TIMESTAMP_MICROS for TIMESTAMP(isAdjustedToUTC = *, unit = MICROS)
	// use ConvertedType TIMESTAMP_MILLIS for TIMESTAMP(isAdjustedToUTC = *, unit = MILLIS)
	parquetLogicalType_TIMESTAMP ParquetLogicalType = 64
	parquetLogicalType_INTEGER   ParquetLogicalType = 128  // use ConvertedType INT_* or UINT_*
	parquetLogicalType_UNKNOWN   ParquetLogicalType = 256  // no compatible ConvertedType
	parquetLogicalType_JSON      ParquetLogicalType = 512  // use ConvertedType JSON
	parquetLogicalType_BSON      ParquetLogicalType = 1024 // use ConvertedType BSON
	parquetLogicalType_UUID      ParquetLogicalType = 2048 // no compatible ConvertedType
)

var ParquetLogicalTypeName = map[ParquetLogicalType]string{
	0:    "STRING",
	1:    "MAP",
	2:    "LIST",
	4:    "ENUM",
	8:    "DECIMAL",
	16:   "DATE",
	32:   "TIME",
	64:   "TIMESTAMP",
	128:  "INTEGER",
	256:  "UNKNOWN",
	512:  "JSON",
	1024: "BSON",
	2048: "UUID",
}

var ParquetLogicalTypeValue = map[string]ParquetLogicalType{
	"STRING":    0,
	"MAP":       1,
	"LIST":      2,
	"ENUM":      4,
	"DECIMAL":   8,
	"DATE":      16,
	"TIME":      32,
	"TIMESTAMP": 64,
	"INTEGER":   128,
	"UNKNOWN":   256,
	"JSON":      512,
	"BSON":      1024,
	"UUID":      2048,
}

type ParquetConvertedType int32

const (
	parquetConvertedType_UTF8             ParquetConvertedType = 0
	parquetConvertedType_MAP              ParquetConvertedType = 1
	parquetConvertedType_MAP_KEY_VALUE    ParquetConvertedType = 2
	parquetConvertedType_LIST             ParquetConvertedType = 4
	parquetConvertedType_ENUM             ParquetConvertedType = 8
	parquetConvertedType_DECIMAL          ParquetConvertedType = 16
	parquetConvertedType_DATE             ParquetConvertedType = 32
	parquetConvertedType_TIME_MILLIS      ParquetConvertedType = 64
	parquetConvertedType_TIME_MICROS      ParquetConvertedType = 128
	parquetConvertedType_TIMESTAMP_MILLIS ParquetConvertedType = 256
	parquetConvertedType_TIMESTAMP_MICROS ParquetConvertedType = 512
	parquetConvertedType_UINT_8           ParquetConvertedType = 1024
	parquetConvertedType_UINT_16          ParquetConvertedType = 2048
	parquetConvertedType_UINT_32          ParquetConvertedType = 4096
	parquetConvertedType_UINT_64          ParquetConvertedType = 8192
	parquetConvertedType_INT_8            ParquetConvertedType = 16384
	parquetConvertedType_INT_16           ParquetConvertedType = 32768
	parquetConvertedType_INT_32           ParquetConvertedType = 65536
	parquetConvertedType_INT_64           ParquetConvertedType = 131072
	parquetConvertedType_JSON             ParquetConvertedType = 262144
	parquetConvertedType_BSON             ParquetConvertedType = 524288
	parquetConvertedType_INTERVAL         ParquetConvertedType = 1048576
)

var ParquetConvertedTypeName = map[ParquetConvertedType]string{
	0:       "UTF8",
	1:       "MAP",
	2:       "MAP_KEY_VALUE",
	4:       "LIST",
	8:       "ENUM",
	16:      "DECIMAL",
	32:      "DATE",
	64:      "TIME_MILLIS",
	128:     "TIME_MICROS",
	256:     "TIMESTAMP_MILLIS",
	512:     "TIMESTAMP_MICROS",
	1024:    "UINT_8",
	2048:    "UINT_16",
	4096:    "UINT_32",
	8192:    "UINT_64",
	16384:   "INT_8",
	32768:   "INT_16",
	65536:   "INT_32",
	131072:  "INT_64",
	262144:  "JSON",
	524288:  "BSON",
	1048576: "INTERVAL",
}

var ParquetConvertedTypeValue = map[string]ParquetConvertedType{
	"UTF8":             0,
	"MAP":              1,
	"MAP_KEY_VALUE":    2,
	"LIST":             4,
	"ENUM":             8,
	"DECIMAL":          16,
	"DATE":             32,
	"TIME_MILLIS":      64,
	"TIME_MICROS":      128,
	"TIMESTAMP_MILLIS": 256,
	"TIMESTAMP_MICROS": 512,
	"UINT_8":           1024,
	"UINT_16":          2048,
	"UINT_32":          4096,
	"UINT_64":          8192,
	"INT_8":            16384,
	"INT_16":           32768,
	"INT_32":           65536,
	"INT_64":           131072,
	"JSON":             262144,
	"BSON":             524288,
	"INTERVAL":         1048576,
}

type FieldRepetitionType int32

const (
	/** This field is required (can not be null) and each record has exactly 1 value. */
	REQUIRED FieldRepetitionType = 0

	/** The field is optional (can be null) and each record has 0 or 1 values. */
	OPTIONAL FieldRepetitionType = 1

	/** The field is repeated and can contain 0 or more values */
	REPEATED FieldRepetitionType = 2
)

var FieldRepetitionTypeName = map[FieldRepetitionType]string{
	0: "REQUIRED",
	1: "OPTIONAL",
	2: "REPEATED",
}

var FieldRepetitionTypeValue = map[string]FieldRepetitionType{
	"REQUIRED": 0,
	"OPTIONAL": 1,
	"REPEATED": 2,
}

type DecimalType struct {
	scale     int32
	precision int32
}

// Time units for logical types
type milliSeconds struct{}

func newMilliSeconds() *milliSeconds { return new(milliSeconds) }

type microSeconds struct{}

func newMicroSeconds() *microSeconds { return new(microSeconds) }

type nanoSeconds struct{}

func newNanoSeconds() *nanoSeconds { return new(nanoSeconds) }

type TimeUnit struct {
	Precision interface{}
}

// Timestamp logical type annotation
// Allowed for physical types: INT64
type TimestampType struct {
	isAdjustedToUTC bool
	unit            TimeUnit
}

// Time logical type annotation
// Allowed for physical types: INT32 (millis), INT64 (micros, nanos)
type TimeType struct {
	isAdjustedToUTC bool
	unit            TimeUnit
}

// Integer logical type annotation
// bitWidth must be 8, 16, 32, or 64.
// Allowed for physical types: INT32, INT64
type IntType struct {
	bitWidth int
	isSigned bool
}

func AvroToParquetType(typeString string) string {
	if avroTypeValue, ok := AvroTypeValue[typeString]; ok {
		return avroToParquetType(avroTypeValue)
	}
	return "UNKNOWN"
}

func avroToParquetType(a AvroType) string {
	switch a {
	case avroType_Null:
		return ""
	case avroType_Boolean:
		return "BOOLEAN"
	case avroType_Int:
		return "INT32"
	case avroType_Long:
		return "INT64"
	case avroType_Float:
		return "FLOAT"
	case avroType_Double:
		return "DOUBLE"
	case avroType_Bytes:
		return "BYTE_ARRAY"
	case avroType_String:
		return "BYTE_ARRAY"
	case avroType_Record:
		return ""
	case avroType_Enum:
		return "ENUM"
	case avroType_Array:
		return ""
	case avroType_Map:
		return ""
	case avroType_Union:
		return ""
	case avroType_Fixed:
		return "BYTE_ARRAY"
	case avroType_Decimal:
		return "FLOAT"
	case avroType_UUID:
		return "BYTE_ARRAY"
	case avroType_Date:
		return "DATE"
	case avroType_TimeMillis:
		return "TIME_MILLIS"
	case avroType_TimeMicros:
		return "TIME_MICROS"
	case avroType_TimestampMillis:
		return "TIMESTAMP_MILLIS" // TIMESTAMP(isAdjustedToUTC=true, unit=MILLIS)
	case avroType_TimestampMicros:
		return "TIMESTAMP_MICROS" // TIMESTAMP(isAdjustedToUTC=true, unit=MICROS)
	case avroType_LocalTimestampMillis:
		return "TIMESTAMP_MILLIS" // TIMESTAMP(isAdjustedToUTC=false, unit=MILLIS)
	case avroType_LocalTimestampMicros:
		return "TIMESTAMP_MICROS" // TIMESTAMP(isAdjustedToUTC=false, unit=MICROS)
	case avroType_Duration:
		return "INTERVAL"
	}
	return ""
}

func GoToParquetType(goType string) string {
	switch goType {
	case "bool":
		return "BOOLEAN"
	case "string":
		return "BYTE_ARRAY"
	case "int", "int8", "int16", "int32", "uint", "uint8", "uint16", "uint32":
		return "INT32"
	case "int64", "uint64":
		return "INT64"
	case "float32":
		return "FLOAT"
	case "float64":
		return "DOUBLE"
	case "nil":
		return "UNKNOWN"
	default:
		return "BYTE_ARRAY"
	}
}
