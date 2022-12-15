package main

var tokenTypeHash = map[int]string{
	0:  "Control",
	1:  "Var Declaration",
	2:  "Function Declaration",
	3:  "Function Call",
	4:  "If Statement",
	5:  "Else Statement",
	6:  "Else If Statement",
	7:  "Return Statement",
	8:  "Var Set",
	9:  "Main Declaration",
	10: "",
}

var typeTable = map[string]string{
	"i64":  "int64",
	"i32":  "int32",
	"i16":  "int16",
	"i8":   "int8",
	"i":    "int",
	"u":    "uint",
	"u8":   "uint8",
	"u16":  "uint16",
	"u32":  "uint32",
	"u64":  "uint64",
	"bool": "bool",
	"ch":   "byte",
	"lch":  "rune",
	"str":  "string",
	"f32":  "float32",
	"f64":  "float64",
}

var funcTransTable = map[string]string{}

var importFuncTransTable = map[string]string{
	"print": "fmt.Println",
}
var importTable = map[string]string{
	"print": "fmt",
}
