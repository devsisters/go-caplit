package cat

// AUTO GENERATED - DO NOT EDIT
// GENERATED BY gen_readcaplit.go

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	. "github.com/devsisters/go-caplit/example/capnp/dog"
	"github.com/glycerine/go-capnproto"
)

func arrayInStringParser(s string) []string {
	s = s[1 : len(s)-1]
	l := len(s)
	innerCount := 0
	buff := ""
	ret_val := make([]string, 0)
	for i := 0; i < l; i++ {
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
			if innerCount < 0 {
				panic("( and ) are not matched.")
			}
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
				panic(fmt.Sprintf("parse error : None status, %s", c))
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

func (s Cat) ReadCapLit(r io.Reader) error {
	b := bufio.NewReader(r)
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

			case "id":
				v, err := strconv.ParseInt(value, 10, 32)
				if err != nil {
					return err
				}
				s.SetId(int32(v))

			case "name":
				runedValue := []rune(value)
				if string(runedValue[0]) != "\"" || string(runedValue[len(runedValue)-1]) != "\"" {
					return errors.New("First and last character of string must be \"")
				}
				s.SetName(value[1 : len(value)-1])

			case "partner":
				v := NewDog(s.Segment)
				err := v.ReadCapLit(bytes.NewReader([]byte(value)))
				if err != nil {
					return err
				}
				s.SetPartner(v)

			default:
				return errors.New(fmt.Sprintf("cannot find key in Cat : %v", key))
			}

			substatus = ""
		}
	}

	if substatus != "" {
		return errors.New("mismatched bracket in Cat")
	}

	return nil
}

func (s Cat) GetSegment() *capn.Segment { return s.Segment }
