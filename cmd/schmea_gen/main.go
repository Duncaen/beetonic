package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/dave/jennifer/jen"
)

var typeFields map[string][]jen.Code

type Occurs struct {
	Unbounded bool
	Value     uint32
}

func (o *Occurs) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value == "unbounded" {
		o.Unbounded = true
	} else {
		u, err := strconv.ParseUint(attr.Value, 10, 32)
		if err != nil {
			return err
		}
		o.Unbounded = false
		o.Value = uint32(u)
	}
	return nil
}

type Schema struct {
	XMLName      xml.Name      `xml:"schema"`
	Elements     []Element     `xml:"element"`
	ComplexTypes []ComplexType `xml:"complexType"`
	SimpleTypes  []SimpleType  `xml:"simpleType"`
}

type Element struct {
	XMLName   xml.Name `xml:"element"`
	Name      string   `xml:"name,attr"`
	Type      string   `xml:"type,attr"`
	MinOccurs Occurs   `xml:"minOccurs,attr"`
	MaxOccurs Occurs   `xml:"maxOccurs,attr"`
}

func (e *Element) StructField() *jen.Statement {
	if e.MinOccurs.Value == 1 && e.MaxOccurs.Value == 1 {
		return jen.Id(strings.Title(e.Name)).Add(Type(e.Type)).Tag(
			map[string]string{
				"xml":  e.Name,
				"json": e.Name,
			},
		)
	}
	if e.MinOccurs.Value == 0 && e.MaxOccurs.Value == 1 {
		return e.PtrStructField()
	}
	return jen.Id(strings.Title(e.Name)).Index().Add(Type(e.Type)).Tag(
		map[string]string{
			"xml":  e.Name,
			"json": e.Name,
		},
	)
}

func (e *Element) PtrStructField() *jen.Statement {
	return jen.Id(strings.Title(e.Name)).Op("*").Add(Type(e.Type)).Tag(
		map[string]string{
			"xml":  e.Name,
			"json": fmt.Sprintf("%s,omitempty", e.Name),
		},
	)
}

type ComplexType struct {
	XMLName        xml.Name        `xml:"complexType"`
	Name           string          `xml:"name,attr"`
	Mixed          bool            `xml:"mixed,attr"`
	Choice         Choice          `xml:"choice"`
	Attributes     []Attribute     `xml:"attribute"`
	Sequence       *Sequence       `xml:"sequence"`
	ComplexContent *ComplexContent `xml:"complexContent"`
}

func (ct *ComplexType) Fields() []jen.Code {
	stmts := []jen.Code{}
	for _, a := range ct.Attributes {
		stmts = append(stmts, a.Go())
	}
	if ct.Mixed {
		stmts = append(stmts, jen.Id("Value").String().Tag(
			map[string]string{
				"xml":  ",chardata",
				"json": "value",
			},
		))
	}
	for _, e := range ct.Choice.Elements {
		stmts = append(stmts, e.PtrStructField())
	}
	if ct.Sequence != nil {
		stmts = append(stmts, ct.Sequence.Fields()...)
	}
	if ct.ComplexContent != nil {
		log.Println(ct.Name)
		stmts = append(stmts, ct.ComplexContent.Extension.Fields()...)
	}
	return stmts
}

func (ct *ComplexType) Go() *jen.Statement {
	return jen.Type().Id(ct.Name).Struct(ct.Fields()...)
}

type ComplexContent struct {
	XMLName   xml.Name  `xml:"complexContent"`
	Extension Extension `xml:"extension"`
}

type Extension struct {
	XMLName    xml.Name    `xml:"extension"`
	Base       string      `xml:"base,attr"`
	Sequence   *Sequence   `xml:"sequence"`
	Attributes []Attribute `xml:"attribute"`
}

func (e *Extension) Fields() []jen.Code {
	base := e.Base
	switch {
	case strings.HasPrefix(base, "sub:"):
		base = base[4:]
	}
	stmts, ok := typeFields[base]
	if !ok {
		log.Fatalf("type %q not found\n", base)
	}
	for _, a := range e.Attributes {
		stmts = append(stmts, a.Go())
	}
	if e.Sequence != nil {
		stmts = append(stmts, e.Sequence.Fields()...)
	}
	return stmts
}

type Choice struct {
	XMLName   xml.Name  `xml:"choice"`
	MinOccurs Occurs    `xml:"minOccurs,attr"`
	MaxOccurs Occurs    `xml:"maxOccurs,attr"`
	Elements  []Element `xml:"element"`
}

type Attribute struct {
	XMLName xml.Name `xml:"attribute"`
	Name    string   `xml:"name,attr"`
	Type    string   `xml:"type,attr"`
	Use     string   `xml:"use,attr"`
}

func (a *Attribute) Go() *jen.Statement {
	xmlAttr := fmt.Sprintf("%s,attr", a.Name)
	jsonAttr := a.Name
	field := jen.Id(strings.Title(a.Name))
	if a.Use == "optional" {
		xmlAttr += ",omitempty"
		jsonAttr += ",omitempty"
	}

	return field.Add(Type(a.Type)).Tag(
		map[string]string{
			"xml":  xmlAttr,
			"json": jsonAttr,
		},
	)
}

type Sequence struct {
	XMLName  xml.Name  `xml:"sequence"`
	Elements []Element `xml:"element"`
}

func (s *Sequence) Fields() []jen.Code {
	stmts := []jen.Code{}
	for _, e := range s.Elements {
		stmts = append(stmts,
			e.StructField(),
			// jen.Id(strings.Title(e.Name)).Index().Add(Type(e.Type)).Tag(
			// 	map[string]string{"xml": e.Name},
			// ),
		)
	}
	return stmts
}

func (s *Sequence) Struct() jen.Code {
	return jen.Struct(s.Fields()...)
}

type SimpleType struct {
	XMLName     xml.Name    `xml:"simpleType"`
	Name        string      `xml:"name,attr"`
	Restriction Restriction `xml:"restriction"`
}

func validateRangeFunc(name string) jen.Code {
	min := fmt.Sprintf("%sMin", name)
	max := fmt.Sprintf("%sMax", name)
	return jen.Func().Params(jen.Id("x").Id(name)).Id("Validate").Params().Error().Block(
		jen.If(jen.Id("x").Op("<").Id(min).Op("||").
			Id("x").Op(">").Id(max)).Block(
			jen.Return().Qual("fmt", "Errorf").Call(jen.Lit("Out of range")),
		),
		jen.Return().Nil(),
	)
}

func marshalIntFunc(name string) jen.Code {
	format := jen.Qual("strconv", "FormatFloat").Call(jen.Float64().Parens(jen.Id("x")), jen.LitRune('f'), jen.Lit(1), jen.Lit(32))
	return jen.Func().Params(jen.Id("x").Id(name)).Id("MarshalText").
		Params().Params(jen.Op("[]").Id("byte"), jen.Id("error")).Block(
		jen.If(jen.Id("err").Op(":=").Id("x").Dot("Validate").Call().Op(";").Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Id("err")),
		),
		jen.Return(jen.Index().Byte().Parens(format), jen.Nil()),
	)
}

func unmarshalIntFunc(name string) jen.Code {
	parse := jen.List(jen.Id("v"), jen.Id("err")).Op(":=").Qual("strconv", "ParseInt").Call(
		jen.String().Parens(jen.Id("text")),
		jen.Lit(10),
		jen.Lit(32),
	)
	return jen.Func().Params(jen.Id("x").Op("*").Id(name)).Id("UnmarshalText").
		Params(jen.Id("text").Op("[]").Id("byte")).Params(jen.Id("error")).Block(
		parse,
		jen.If(jen.Id("err").Op("!=").Nil()).Block(jen.Return().Id("err")),
		jen.Op("*").Id("x").Op("=").Id(name).Parens(jen.Id("v")),
		jen.Return().Id("x").Dot("Validate").Call(),
	)
}

func marshalFloatFunc(name string) jen.Code {
	format := jen.Qual("strconv", "FormatInt").Call(jen.Int64().Parens(jen.Id("x")), jen.Lit(10))

	return jen.Func().Params(jen.Id("x").Id(name)).Id("MarshalText").
		Params().Params(jen.Op("[]").Id("byte"), jen.Id("error")).Block(
		jen.If(jen.Id("err").Op(":=").Id("x").Dot("Validate").Call().Op(";").Id("err").Op("!=").Nil()).Block(
			jen.Return(jen.Nil(), jen.Id("err")),
		),
		jen.Return(jen.Index().Byte().Parens(format), jen.Nil()),
	)
}

func unmarshalFloatFunc(name string) jen.Code {
	parse := jen.List(jen.Id("v"), jen.Id("err")).Op(":=").Qual("strconv", "ParseFloat").Call(
		jen.String().Parens(jen.Id("text")),
		jen.Lit(32),
	)
	return jen.Func().Params(jen.Id("x").Op("*").Id(name)).Id("UnmarshalText").
		Params(jen.Id("text").Op("[]").Id("byte")).Params(jen.Id("error")).Block(
		parse,
		jen.If(jen.Id("err").Op("!=").Nil()).Block(jen.Return().Id("err")),
		jen.Op("*").Id("x").Op("=").Id(name).Parens(jen.Id("v")),
		jen.Return().Id("x").Dot("Validate").Call(),
	)
}

// func (st *SimpleType) Fields() []jen.Code {
// 	var stmts []jen.Code
// }

func (st *SimpleType) Go() []jen.Code {
	var stmts []jen.Code
	if len(st.Restriction.Enumerations) > 0 {
		switch st.Restriction.Base {
		case "xs:string":
			stmts = append(stmts, jen.Type().Id(st.Name).Int())
			var values []jen.Code
			for _, e := range st.Restriction.Enumerations {
				values = append(values, jen.Lit(e.Value))
			}
			stmts = append(stmts, jen.Var().Id(fmt.Sprintf("%sValues", st.Name)).Op("=").Index().String().Values(values...))
		default:
			panic(fmt.Sprintf("unhandled restriction %q", st.Restriction.Base))
		}
		defs := []jen.Code{}
		var values []jen.Code
		for i, e := range st.Restriction.Enumerations {
			name := fmt.Sprintf("%s%s", st.Name, strings.Title(e.Value))
			defs = append(defs,
				jen.Id(name).Id(st.Name).Op("=").Lit(i),
			)
			values = append(values, jen.Id(name))
		}
		stmts = append(stmts, jen.Const().Defs(defs...))
	} else {
		stmts = append(stmts, jen.Type().Id(st.Name).Add(Type(st.Restriction.Base)))
		if st.Restriction.MinInclusive != nil && st.Restriction.MinInclusive != nil {
			switch st.Restriction.Base {
			case "xs:int":
				var min, max int64
				var err error
				if min, err = strconv.ParseInt(st.Restriction.MinInclusive.Value, 10, 32); err != nil {
					log.Fatal(err)
				}
				if max, err = strconv.ParseInt(st.Restriction.MaxInclusive.Value, 10, 32); err != nil {
					log.Fatal(err)
				}
				stmts = append(stmts, jen.Const().Defs(
					jen.Id(fmt.Sprintf("%sMin", st.Name)).Id(st.Name).Op("=").Lit(int(min)),
					jen.Id(fmt.Sprintf("%sMax", st.Name)).Id(st.Name).Op("=").Lit(int(max)),
				))
			case "xs:double":
				var min, max float64
				var err error
				if min, err = strconv.ParseFloat(st.Restriction.MinInclusive.Value, 32); err != nil {
					log.Fatal(err)
				}
				if max, err = strconv.ParseFloat(st.Restriction.MaxInclusive.Value, 32); err != nil {
					log.Fatal(err)
				}
				stmts = append(stmts, jen.Const().Defs(
					jen.Id(fmt.Sprintf("%sMin", st.Name)).Id(st.Name).Op("=").Lit(min),
					jen.Id(fmt.Sprintf("%sMax", st.Name)).Id(st.Name).Op("=").Lit(max),
				))
			default:
				panic(fmt.Sprintf("unhandled restriction %q", st.Restriction.Base))
			}
		}
	}
	stmts = append(stmts, jen.Commentf("make sure %s implements Validate", st.Name))
	stmts = append(stmts, jen.Var().Id("_").Id("Validate").Op("=").Parens(jen.Op("*").Id(st.Name)).Parens(jen.Nil()))
	return stmts
}

type Restriction struct {
	XMLName      xml.Name      `xml:"restriction"`
	Base         string        `xml:"base,attr"`
	Enumerations []Enumeration `xml:"enumeration"`
	Pattern      *Pattern      `xml:"pattern"`
	MinInclusive *struct {
		Value string `xml:"value,attr"`
	} `xml:"minInclusive"`
	MaxInclusive *struct {
		Value string `xml:"value,attr"`
	} `xml:"maxInclusive"`
}

type Enumeration struct {
	XMLName xml.Name `xml:"enumeration"`
	Value   string   `xml:"value,attr"`
}

type Pattern struct {
	XMLName xml.Name `xml:"pattern"`
	Value   string   `xml:"value,attr"`
}

func Type(s string) *jen.Statement {
	if strings.HasPrefix(s, "sub:") {
		if s[4:] == "type" {
			panic("type")
		}
		return jen.Id(s[4:])
	} else {
		switch s {
		case "xs:boolean":
			return jen.Bool()
		case "xs:int":
			return jen.Int()
		case "xs:long":
			return jen.Int64()
		case "xs:double":
			return jen.Float32()
		case "xs:float":
			return jen.Float64()
		case "xs:string":
			return jen.String()
		case "xs:dateTime":
			return jen.Id("DateTime")
		default:
			panic(fmt.Sprintf("unhandled type %q", s))
		}
	}
}

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	var schema Schema
	err = xml.Unmarshal(bytes, &schema)
	if err != nil {
		log.Fatal("Unmarschal: ", err)
	}

	typeFields = make(map[string][]jen.Code)
	// for _, st := range schema.SimpleTypes {
	// 	typeFields[st.Name] = st.Fields()
	// }
	for _, ct := range schema.ComplexTypes {
		log.Println(ct.Name)
		typeFields[ct.Name] = ct.Fields()
	}

	file := jen.NewFile("spec")
	file.Commentf("go:generate goyacc -o gopher.go -p parser gopher.y")
	for _, e := range schema.Elements {
		file.Add(jen.Type().Id("SubsonicResponse").Struct(
			jen.Id("XMLName").Qual("encoding/xml", "Name").Tag(
				map[string]string{
					"xml":  e.Name,
					"json": "-",
				},
			),
			jen.Id("XMLNS").Id("string").Tag(
				map[string]string{
					"xml":  "xmlns,attr",
					"json": "-",
				},
			),
			Type(e.Type).Tag(map[string]string{"json": "subsonic-response"}),
		))
	}
	for _, st := range schema.SimpleTypes {
		for _, x := range st.Go() {
			file.Add(x)
		}
	}
	for _, ct := range schema.ComplexTypes {
		file.Add(ct.Go())
	}
	fmt.Printf("%#v", file)
}
