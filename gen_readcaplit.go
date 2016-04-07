package caplit

import (
	"fmt"
	"strings"
	"go/ast"
)

func GenCapnpReadCapLit(inputPath, outputPath string, packageName string, enumList []string) {
	capnpStructs := CapnpStructs(inputPath)

	int8Template := `
					v, err := strconv.ParseInt(value, 10, 8)
					if err != nil {
						return err
					}
					s.Set%v(int8(v))`

	int16Template := `
					v, err := strconv.ParseInt(value, 10, 16)
					if err != nil {
						return err
					}
					s.Set%v(int16(v))`

	int32Template := `
					v, err := strconv.ParseInt(value, 10, 32)
					if err != nil {
						return err
					}
					s.Set%v(int32(v))`

	intTemplate := `
					v, err := strconv.ParseInt(value, 10, 32)
					if err != nil {
						return err
					}
					s.Set%v(int(v))`

	int64Template := `
					v, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return err
					}
					s.Set%v(v)`

	uint8Template := `
					v, err := strconv.ParseUint(value, 10, 8)
					if err != nil {
						return err
					}
					s.Set%v(uint8(v))`

	uint16Template := `
					v, err := strconv.ParseUint(value, 10, 16)
					if err != nil {
						return err
					}
					s.Set%v(uint16(v))`

	uint32Template := `
					v, err := strconv.ParseUint(value, 10, 32)
					if err != nil {
						return err
					}
					s.Set%v(uint32(v))`

	uintTemplate := `
					v, err := strconv.ParseUint(value, 10, 32)
					if err != nil {
						return err
					}
					s.Set%v(uint(v))`

	uint64Template := `
					v, err := strconv.ParseUint(value, 10, 64)
					if err != nil {
						return err
					}
					s.Set%v(v)`

	float32Template := `
					v, err := strconv.ParseFloat(value, 32)
					if err != nil {
						return err
					}
					s.Set%v(float32(v))`

	float64Template := `
					v, err := strconv.ParseFloat(value, 64)
					if err != nil {
						return err
					}
					s.Set%v(v)`

	stringTemplate := `
					s.Set%v(value[1:len(value)-1])`

	boolTemplate := `
					v, err := strconv.ParseBool(value)
					if err != nil {
						return err
					}
					s.Set%v(v)`

	// value example : "[1,22,3,5]"
	textListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewTextList(len(valueList))
					for i, vs := range valueList {
						v.Set(i, vs)
					}
					s.Set%v(v)`

	int8ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewInt8List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseInt(vs, 10, 8)
						if err != nil {
							return err
						}
						v.Set(i, int8(elem))
					}
					s.Set%v(v)`

	int16ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewInt16List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseInt(vs, 10, 16)
						if err != nil {
							return err
						}
						v.Set(i, int16(elem))
					}
					s.Set%v(v)`

	int32ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewInt32List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseInt(vs, 10, 32)
						if err != nil {
							return err
						}
						v.Set(i, int32(elem))
					}
					s.Set%v(v)`

	intListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewIntList(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseInt(vs, 10, 32)
						if err != nil {
							return err
						}
						v.Set(i, int(elem))
					}
					s.Set%v(v)`

	int64ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewInt64List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseInt(vs, 10, 64)
						if err != nil {
							return err
						}
						v.Set(i, elem)
					}
					s.Set%v(v)`

	uint8ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewUInt8List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseUint(vs, 10, 8)
						if err != nil {
							return err
						}
						v.Set(i, uint8(elem))
					}
					s.Set%v(v)`

	uint16ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewUInt16List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseUint(vs, 10, 16)
						if err != nil {
							return err
						}
						v.Set(i, uint16(elem))
					}
					s.Set%v(v)`

	uint32ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewUInt32List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseUint(vs, 10, 32)
						if err != nil {
							return err
						}
						v.Set(i, uint32(elem))
					}
					s.Set%v(v)`

	uintListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewIntList(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseUint(vs, 10, 32)
						if err != nil {
							return err
						}
						v.Set(i, uint(elem))
					}
					s.Set%v(v)`

	uint64ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewUInt64List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseUint(vs, 10, 64)
						if err != nil {
							return err
						}
						v.Set(i, elem)
					}
					s.Set%v(v)`

	float32ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewFloat32List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseFloat(vs, 32)
						if err != nil {
							return err
						}
						v.Set(i, elem)
					}
					s.Set%v(float32(v))`

	float64ListTemplate := `
					valueList := arrayInStringParser(value)
					v := s.Segment.NewFloat64List(len(valueList))
					for i, vs := range valueList {
						elem, err := strconv.ParseFloat(vs, 64)
						if err != nil {
							return err
						}
						v.Set(i, elem)
					}
					s.Set%v(v)`

	voidTemplate := ``

	byteArrayTemplate := `
					s.Set%v([]byte(value))`

	enumTemplate := `
					v := %sFromString(value)
					s.Set%s(v)`

	enumListTemplate := `
					valueList := arrayInStringParser(value)
					v := New%s(s.Segment, len(valueList))
					for i, vs := range valueList {
						elem := %sFromString(vs)
						v.Set(i, elem)
					}
					s.Set%s(v)`

	structListTemplate := `
					valueList := arrayInStringParser(value)
					v := New%s(s.Segment, len(valueList))
					for i, vs := range valueList {
						elem := New%s(s.Segment)
						elem.ReadCapLit(bytes.NewReader([]byte(vs)))
						v.Set(i, elem)
					}
					s.Set%s(v)`

	structTemplate := `
					v := New%s(s.Segment)
					v.ReadCapLit(bytes.NewReader([]byte(value)))
					s.Set%s(v)`

	newCapnpStructs := make([]CapnpStruct,0)
	for _, capnpStruct := range capnpStructs {
		filter := func(t *ast.FuncDecl) bool {
			return receiverType(t) == capnpStruct.Name && strings.HasPrefix(t.Name.Name, "Set")
		}
		funcDecls := parseCapnpFuncDecl(inputPath + "/*.capnp.go", filter)
		capnpStruct.Keys = make([]CapnpStructParams,0)

		for _, funcDecl := range funcDecls {
			typeName := funcDecl.FuncDecl.Name.Name[3:]
			var capnpStructParams CapnpStructParams
			pt := paramType(funcDecl.FuncDecl)
			switch pt {
			case "":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: voidTemplate,
				}
			case "int8":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int8Template, typeName),
				}
			case "int16":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int16Template, typeName),
				}
			case "int32":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int32Template, typeName),
				}
			case "int":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(intTemplate, typeName),
				}
			case "int64":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int64Template, typeName),
				}
			case "uint8":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint8Template, typeName),
				}
			case "uint16":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint16Template, typeName),
				}
			case "uint32":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint32Template, typeName),
				}
			case "uint":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uintTemplate, typeName),
				}
			case "uint64":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint64Template, typeName),
				}
			case "float32":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(float32Template, typeName),
				}
			case "float64":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(float64Template, typeName),
				}
			case "bool":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(boolTemplate, typeName),
				}
			case "string":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(stringTemplate, typeName),
				}
			case "[]byte":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(byteArrayTemplate, typeName),
				}
			// ---- 아래는 Default Type의 List에 관한 처리
			case "TextList":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(textListTemplate, typeName),
				}
			case "Int8List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int8ListTemplate, typeName),
				}
			case "Int16List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int16ListTemplate, typeName),
				}
			case "Int32List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int32ListTemplate, typeName),
				}
			case "IntList":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(intListTemplate, typeName),
				}
			case "Int64List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(int64ListTemplate, typeName),
				}
			case "UInt8List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint8ListTemplate, typeName),
				}
			case "UInt16List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint16ListTemplate, typeName),
				}
			case "UInt32List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint32ListTemplate, typeName),
				}
			case "UIntList":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uintListTemplate, typeName),
				}
			case "UInt64List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(uint64ListTemplate, typeName),
				}
			case "Float32List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(float32ListTemplate, typeName),
				}
			case "Float64List":
				capnpStructParams = CapnpStructParams{
					Name:     firstToLower(typeName),
					Template: fmt.Sprintf(float64ListTemplate, typeName),
				}
			default:
				// Handle enums
				if elemInList(pt, enumList) {
					capnpStructParams = CapnpStructParams{
						Name:     firstToLower(typeName),
						Template: fmt.Sprintf(enumTemplate, pt, typeName),
					}
				} else {
					// Handle custom lists
					if strings.HasSuffix(pt, "_List") {
						rootPt := pt[:len(pt) - 5]
						pt = rootPt + "List"
						// Handle enums
						if elemInList(rootPt, enumList) {
							capnpStructParams = CapnpStructParams{
								Name:     firstToLower(typeName),
								Template: fmt.Sprintf(enumListTemplate, pt, rootPt, typeName),
							}
							// Handle list of custom structs
						} else {
							capnpStructParams = CapnpStructParams{
								Name:     firstToLower(typeName),
								Template: fmt.Sprintf(structListTemplate, pt, rootPt, typeName),
							}
						}
						// Handle custom structs
					} else {
						capnpStructParams = CapnpStructParams{
							Name:     firstToLower(typeName),
							Template: fmt.Sprintf(structTemplate, pt, typeName),
						}
					}
				}
			}
			capnpStruct.Keys = append(capnpStruct.Keys, capnpStructParams)
		}
		newCapnpStructs = append(newCapnpStructs, capnpStruct)
	}

	params := map[string]interface{}{
		"package": packageName,
		"structs": newCapnpStructs,
	}

	template := `package {{.package}}

// AUTO GENERATED - DO NOT EDIT
// GENERATED BY gen_readcaplit.go

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"strconv"
	"fmt"

	"github.com/glycerine/go-capnproto"
)

func arrayInStringParser(s string) []string {
	s = s[1:len(s)-1]
	l := len(s)
	innerCount := 0
	buff := ""
	ret_val := make([]string, 0)
	for i:=0; i<l; i++ {
		target := string(s[i])

		switch target {
		case "[":
			fallthrough
		case "(":
			innerCount++
			buff += target
		case "]":
			fallthrough
		case ")":
			innerCount--
			buff += target
		case ",":
			if innerCount == 0 {
				ret_val = append(ret_val, buff)
				buff = ""
			} else {
				buff += target
			}
		default:
			buff += target
		}
	}
	if buff != "" {
		ret_val = append(ret_val, buff)
	}
	return ret_val
}


func capLitParser() func(string) (string, string, string, bool) {
	const None = 0
	const In = 1
	const Key = 2
	const Value = 3

	status := None
	substatus := ""
	key := ""
	value := ""
    innerCount := 0
    inQuote := false

	return func(c string) (string, string, string, bool) {
        if substatus == "flush" {
            substatus = ""
            key = ""
            value = ""
        }

		if status == None {
			switch c {
			case "(":
				status = Key
			case ")":
				status = None
			default:
				panic(fmt.Sprintf("parse error : None status, %s",c))
			}
		} else if status == In {
			switch c {
			case ",":
				status = Key
			}
		} else if status == Key {
			switch c {
			case "=":
				status = Value
			case " ":
			default:
				key += c
			}
		} else if status == Value {
			// array( [] ), group(), struct(), plain text
			if inQuote {
				value += c
			} else if substatus == "plain" {
				switch c {
				case ",":
					status = Key
					fallthrough
				case ")":
					substatus = "flush"
				default:
					value += c
				}
			} else if substatus == "(" {
				value += c
				switch c {
				case "(":
					innerCount++
				case ")":
					innerCount--
					if innerCount < 0 {
						panic("parse error : ( and ) does not match")
					} else if innerCount == 0 {
						substatus = "flush"
						status = In
					}
				}
			} else if substatus == "[" {
				value += c
				switch c {
				case "[":
					innerCount++
				case "]":
					innerCount--
					if innerCount < 0 {
						panic("parse error : [ and ] does not match")
					} else if innerCount == 0 {
						substatus = "flush"
						status = In
					}
				}
			} else {
				value += c
				switch c {
				case "[":
					fallthrough
				case "(":
					innerCount++
					substatus = c
				default:
					substatus = "plain"
				}
			}
		}

        if c == "\"" {
            inQuote = !inQuote
        }

		return substatus, key, value, inQuote
	}
}

{{range .structs}}
func (s {{.Name}}) ReadCapLit(r io.Reader) error {
	b := bufio.NewReader(r)
	parseError := errors.New("Parse error")
	var substatus, key, value string
	var inQuote bool
	parser := capLitParser()
	for {
		b, err := b.ReadByte()
		if err != nil {
			break
		}

		c := string(b)
		if c == "\n" || c == "\t" {
			continue
		}

		if c == " " && !inQuote {
			continue
		}

		substatus, key, value, inQuote = parser(c)

		if substatus == "flush" {
			switch key {
			{{range .Keys}}
			case "{{.Name}}": {{.Template}}
			{{end}}
			default:
				return parseError
			}

			substatus = ""
		}
	}

	if substatus != "" {
		return parseError
	}

	return nil
}

func (s {{.Name}}) GetSegment() *capn.Segment { return s.Segment }
{{end}}
`

	generate(template, outputPath, params)
}
