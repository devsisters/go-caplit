package cat

// AUTO GENERATED - DO NOT EDIT

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/devsisters/go-caplit/example/capnp/dog"
	C "github.com/glycerine/go-capnproto"
	"io"
)

type Cat C.Struct

func NewCat(s *C.Segment) Cat      { return Cat(s.NewStruct(8, 2)) }
func NewRootCat(s *C.Segment) Cat  { return Cat(s.NewRootStruct(8, 2)) }
func AutoNewCat(s *C.Segment) Cat  { return Cat(s.NewStructAR(8, 2)) }
func ReadRootCat(s *C.Segment) Cat { return Cat(s.Root(0).ToStruct()) }
func (s Cat) Id() int32            { return int32(C.Struct(s).Get32(0)) }
func (s Cat) SetId(v int32)        { C.Struct(s).Set32(0, uint32(v)) }
func (s Cat) Name() string         { return C.Struct(s).GetObject(0).ToText() }
func (s Cat) SetName(v string)     { C.Struct(s).SetObject(0, s.Segment.NewText(v)) }
func (s Cat) Partner() dog.Dog     { return dog.Dog(C.Struct(s).GetObject(1).ToStruct()) }
func (s Cat) SetPartner(v dog.Dog) { C.Struct(s).SetObject(1, C.Object(v)) }
func (s Cat) WriteJSON(w io.Writer) error {
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
	_, err = b.WriteString("\"partner\":")
	if err != nil {
		return err
	}
	{
		s := s.Partner()
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
func (s Cat) MarshalJSON() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteJSON(&b)
	return b.Bytes(), err
}
func (s Cat) WriteCapLit(w io.Writer) error {
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
	_, err = b.WriteString("partner = ")
	if err != nil {
		return err
	}
	{
		s := s.Partner()
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
func (s Cat) MarshalCapLit() ([]byte, error) {
	b := bytes.Buffer{}
	err := s.WriteCapLit(&b)
	return b.Bytes(), err
}

type Cat_List C.PointerList

func NewCatList(s *C.Segment, sz int) Cat_List { return Cat_List(s.NewCompositeList(8, 2, sz)) }
func (s Cat_List) Len() int                    { return C.PointerList(s).Len() }
func (s Cat_List) At(i int) Cat                { return Cat(C.PointerList(s).At(i).ToStruct()) }
func (s Cat_List) ToArray() []Cat {
	n := s.Len()
	a := make([]Cat, n)
	for i := 0; i < n; i++ {
		a[i] = s.At(i)
	}
	return a
}
func (s Cat_List) Set(i int, item Cat) { C.PointerList(s).Set(i, C.Object(item)) }
