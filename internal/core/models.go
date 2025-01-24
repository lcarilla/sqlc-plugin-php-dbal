package core

import "github.com/sqlc-dev/plugin-sdk-go/plugin"

type Query struct {
	ClassName    string
	Cmd          string
	Comments     []string
	MethodName   string
	FieldName    string
	ConstantName string
	SQL          string
	SourceName   string
	Ret          QueryValue
	Arg          Params
}

type Field struct {
	ID      int
	Name    string
	Type    phpType
	Comment string
}

type Struct struct {
	Table   plugin.Identifier
	Name    string
	Fields  []Field
	Comment string
}

type QueryValue struct {
	Emit   bool
	Name   string
	Struct *Struct
	Typ    phpType
}

func (v QueryValue) EmitStruct() bool {
	return v.Emit
}

func (v QueryValue) IsStruct() bool {
	return v.Struct != nil
}

func (v QueryValue) isEmpty() bool {
	return v.Typ == (phpType{}) && v.Name == "" && v.Struct == nil
}

func (v QueryValue) Type() string {
	if v.Typ != (phpType{}) {
		return v.Typ.String()
	}
	if v.Struct != nil {
		return v.Struct.Name
	}
	panic("no type for QueryValue: " + v.Name)
}

type PhpTmplCtx struct {
	Package             string
	DataClasses         []Struct
	Queries             []Query
	Settings            *plugin.Settings
	SqlcVersion         string
	SourceName          string
	EmitJSONTags        bool
	EmitPreparedQueries bool
	EmitInterface       bool
}

type phpType struct {
	Name     string
	IsArray  bool
	IsNull   bool
	DataType string
	Engine   string
}

func (t phpType) String() string {
	v := t.Name
	if t.IsArray {
		v = "array"
	} else if t.IsNull {
		v = "?" + v
	}
	return v
}

func (t phpType) IsDateTimeImmutable() bool {
	return t.Name == "\\DateTimeImmutable"
}

func (t phpType) IsInt() bool {
	return t.Name == "int"
}

func (t phpType) IsFloat() bool {
	return t.Name == "float"
}

func (t phpType) IsString() bool {
	return t.Name == "string"
}