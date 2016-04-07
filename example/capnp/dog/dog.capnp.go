package dog

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type DogToy C.Struct

func NewDogToy(s *C.Segment) DogToy      { return DogToy(s.NewStruct(16, 0)) }
func NewRootDogToy(s *C.Segment) DogToy  { return DogToy(s.NewRootStruct(16, 0)) }
func AutoNewDogToy(s *C.Segment) DogToy  { return DogToy(s.NewStructAR(16, 0)) }
func ReadRootDogToy(s *C.Segment) DogToy { return DogToy(s.Root(0).ToStruct()) }
func (s DogToy) Id() int32               { return int32(C.Struct(s).Get32(0)) }
func (s DogToy) SetId(v int32)           { C.Struct(s).Set32(0, uint32(v)) }
func (s DogToy) Price() int64            { return int64(C.Struct(s).Get64(8)) }
func (s DogToy) SetPrice(v int64)        { C.Struct(s).Set64(8, uint64(v)) }
func (s DogToy) ToyType() ToyType        { return ToyType(C.Struct(s).Get16(4)) }
func (s DogToy) SetToyType(v ToyType)    { C.Struct(s).Set16(4, uint16(v)) }
func (s DogToy) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"id\":")
	if err != nil {
		return err
	}
	{
		s := s.Id()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"price\":")
	if err != nil {
		return err
	}
	{
		s := s.Price()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"toyType\":")
	if err != nil {
		return err
	}
	{
		s := s.ToyType()
		err = s.WriteJSON(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DogToy) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s DogToy) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("id = ")
	if err != nil {
		return err
	}
	{
		s := s.Id()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("price = ")
	if err != nil {
		return err
	}
	{
		s := s.Price()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("toyType = ")
	if err != nil {
		return err
	}
	{
		s := s.ToyType()
		err = s.WriteCapLit(b)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s DogToy) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type DogToy_List C.PointerList

func NewDogToyList(s *C.Segment, sz int) DogToy_List {
	return DogToy_List(s.NewCompositeList(16, 0, sz))
}
func (s DogToy_List) Len() int        { return C.PointerList(s).Len() }
func (s DogToy_List) At(i int) DogToy { return DogToy(C.PointerList(s).At(i).ToStruct()) }
func (s DogToy_List) ToArray() []DogToy {
	n := s.Len()
	a := make([]DogToy, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s DogToy_List) Set(i int, item DogToy) { C.PointerList(s).Set(i, C.Object(item)) }

type Dog C.Struct

func NewDog(s *C.Segment) Dog       { return Dog(s.NewStruct(8, 2)) }
func NewRootDog(s *C.Segment) Dog   { return Dog(s.NewRootStruct(8, 2)) }
func AutoNewDog(s *C.Segment) Dog   { return Dog(s.NewStructAR(8, 2)) }
func ReadRootDog(s *C.Segment) Dog  { return Dog(s.Root(0).ToStruct()) }
func (s Dog) Id() int32             { return int32(C.Struct(s).Get32(0)) }
func (s Dog) SetId(v int32)         { C.Struct(s).Set32(0, uint32(v)) }
func (s Dog) Name() string          { return C.Struct(s).GetObject(0).ToText() }
func (s Dog) SetName(v string)      { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Dog) Age() int8             { return int8(C.Struct(s).Get8(4)) }
func (s Dog) SetAge(v int8)         { C.Struct(s).Set8(4, uint8(v)) }
func (s Dog) Toys() DogToy_List     { return DogToy_List(C.Struct(s).GetObject(1)) }
func (s Dog) SetToys(v DogToy_List) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s Dog) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('{')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"id\":")
	if err != nil {
		return err
	}
	{
		s := s.Id()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"name\":")
	if err != nil {
		return err
	}
	{
		s := s.Name()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"age\":")
	if err != nil {
		return err
	}
	{
		s := s.Age()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(',')
	if err != nil {
		return err
	}
	_, err = b.WriteString("\"toys\":")
	if err != nil {
		return err
	}
	{
		s := s.Toys()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteJSON(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte('}')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Dog) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Dog) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	err = b.WriteByte('(')
	if err != nil {
		return err
	}
	_, err = b.WriteString("id = ")
	if err != nil {
		return err
	}
	{
		s := s.Id()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("name = ")
	if err != nil {
		return err
	}
	{
		s := s.Name()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("age = ")
	if err != nil {
		return err
	}
	{
		s := s.Age()
		buf, err = json.Marshal(s)
		if err != nil {
			return err
		}
		_, err = b.Write(buf)
		if err != nil {
			return err
		}
	}
	_, err = b.WriteString(", ")
	if err != nil {
		return err
	}
	_, err = b.WriteString("toys = ")
	if err != nil {
		return err
	}
	{
		s := s.Toys()
		{
			err = b.WriteByte('[')
			if err != nil {
				return err
			}
			for i, s := range s.ToArray() {
				if i != 0 {
					_, err = b.WriteString(", ")
				}
				if err != nil {
					return err
				}
				err = s.WriteCapLit(b)
				if err != nil {
					return err
				}
			}
			err = b.WriteByte(']')
		}
		if err != nil {
			return err
		}
	}
	err = b.WriteByte(')')
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s Dog) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Dog_List C.PointerList

func NewDogList(s *C.Segment, sz int) Dog_List { return Dog_List(s.NewCompositeList(8, 2, sz)) }
func (s Dog_List) Len() int                    { return C.PointerList(s).Len() }
func (s Dog_List) At(i int) Dog                { return Dog(C.PointerList(s).At(i).ToStruct()) }
func (s Dog_List) ToArray() []Dog {
	n := s.Len()
	a := make([]Dog, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Dog_List) Set(i int, item Dog) { C.PointerList(s).Set(i, C.Object(item)) }

type ToyType uint16

const (
	TOYTYPE_NONE ToyType = 0
	TOYTYPE_DOLL ToyType = 1
	TOYTYPE_ASDF ToyType = 2
)

func (c ToyType) String() string {
	switch c {
	case TOYTYPE_NONE:
		return "none"
	case TOYTYPE_DOLL:
		return "doll"
	case TOYTYPE_ASDF:
		return "asdf"
	default:
		return ""
	}
}

func ToyTypeFromString(c string) ToyType {
	switch c {
	case "none":
		return TOYTYPE_NONE
	case "doll":
		return TOYTYPE_DOLL
	case "asdf":
		return TOYTYPE_ASDF
	default:
		return 0
	}
}

type ToyType_List C.PointerList

func NewToyTypeList(s *C.Segment, sz int) ToyType_List { return ToyType_List(s.NewUInt16List(sz)) }
func (s ToyType_List) Len() int                        { return C.UInt16List(s).Len() }
func (s ToyType_List) At(i int) ToyType                { return ToyType(C.UInt16List(s).At(i)) }
func (s ToyType_List) ToArray() []ToyType {
	n := s.Len()
	a := make([]ToyType, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s ToyType_List) Set(i int, item ToyType) { C.UInt16List(s).Set(i, uint16(item)) }
func (s ToyType) WriteJSON(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	buf, err = json.Marshal(s.String())
	if err != nil {
		return err
	}
	_, err = b.Write(buf)
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ToyType) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s ToyType) WriteCapLit(w io.Writer) error {
	b := bufio.NewWriter(w)
	var err error
	var buf []byte
	_ = buf
	_, err = b.WriteString(s.String())
	if err != nil {
		return err
	}
	err = b.Flush()
	return err
}
func (s ToyType) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}
