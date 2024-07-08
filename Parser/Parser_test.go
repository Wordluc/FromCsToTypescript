package Parser

import (
	"GoFromCsToTypescript/Lexer"
	"testing"
)

func Test_parseParam1(t *testing.T) {
	l, _ := Lexer.New("public string name;")

	node, err := parseParam(l)
	if err != nil {
		panic(err)
	}
	if node.Name!="name" {
		t.Error("test:param name not found")
	}
	typ,ok:=node.Type.(SimpleTypeNode)
	if !ok {
		t.Error("test:param type not found")
	}
	if typ.Type!=String {
		t.Error("test:param type not string")
	}
}

func Test_parseParam2(t *testing.T) {
	l, _ := Lexer.New("int name;")

	node, err := parseParam(l)
	if err != nil {
		panic(err)
	}
	if node.Name!="name" {
		t.Error("test:param name not found")
	}
	typ,ok:=node.Type.(SimpleTypeNode)
	if !ok {
		t.Error("test:param simple type not found")
	}
	if typ.Type!=Number {
		t.Error("test:param type not number")
	}
}

func Test_parseParamCustom(t *testing.T) {
	l, _ := Lexer.New("public persona p;")

	node, err := parseParam(l)
	if err != nil {
		panic(err)
	}
	if node.Name!="p" {
		t.Error("test:param name not found")
	}
	typ,ok:=node.Type.(CustomTypeNode)
	if !ok {
		t.Error("test:param type not custom found")
	}
	if typ.Type!="persona" {
		t.Error("test:param type not persona")
	}
}
func Test_parseParamGenericType(t *testing.T) {
	l, _ := Lexer.New("public prova<int> name;")

	node, err := parseParam(l)
	if err != nil {
		panic(err)
	}
	if node.Name!="name" {
		t.Error("test:param name not found")
	}
	typeG,ok:=node.Type.(GenericTypeNode)
	if !ok {
		t.Error("test:param generic type not found")
	}
  if typeG.ParentName!="prova" {
		t.Error("test:param generic type not found")
	}

	typs:=typeG.ChildType
	for _,typ:=range typs {
		if typ.(SimpleTypeNode).Type!=Number {
			t.Error("test:param type not number")
		}
	}
}

func Test_parseParamComplexGenericType(t *testing.T) {
	l, _ := Lexer.New("public prova<Int64,string> name;")

	node, err := parseParam(l)
	if err != nil {
		panic(err)
	}
	if node.Name!="name" {
		t.Error("test:param name not found")
	}
	gType,ok:=node.Type.(GenericTypeNode)
	if !ok {
		t.Error("test:param generic type not found")
	}
	if gType.ChildType[0].(SimpleTypeNode).Type!=Number {
		t.Error("test:param type not number")
	}
	if gType.ChildType[1].(SimpleTypeNode).Type!=String {
		t.Error("test:param type not string")
	}
}
func Test_parseParamComposedGenericType(t *testing.T) {
	l, _ := Lexer.New("public prova<tipo<int,string>,string> name;")
	node, err := parseParam(l)
	if err != nil {
		panic(err)
	}
	if node.Name!="name" {
		t.Error("test:param name not found")
	}
	gType,ok:=node.Type.(GenericTypeNode)
	if !ok {
		t.Error("test:param generic type not found")
	}
	composedType:=gType.ChildType[0].(GenericTypeNode)
	if composedType.ChildType[0].(SimpleTypeNode).Type!=Number {
		t.Error("test:param type not number")
	}
	if composedType.ChildType[1].(SimpleTypeNode).Type!=String {
		t.Error("test:param type not string")
	}
	if composedType.ParentName!="tipo" {
		t.Error("test:parent name not tipo")
	}
	if gType.ChildType[1].(SimpleTypeNode).Type!=String {
		t.Error("test:param type not string")
	}
}
func Test_parseClass(t *testing.T) {
	l, _ := Lexer.New(`
   public class prova{
		 public string name;//ciao commento
   }`)
	class, err := parseClass(l)
	if err != nil {
		panic(err)
	}
	if class.Name!="prova" {
		t.Error("test:class name not found")
	}
	if len(class.Fields)!=1 {
		t.Error("test:class fields not found")
	}
	if class.Fields[0].Name!="name" {
		t.Error("test:class field name not found")
	}
	if class.Fields[0].Nullable==true {
		t.Error("test:class field nullable")
	}
	if class.Fields[0].Type.(SimpleTypeNode).Type!=String {
		t.Error("test:class field type not string")
	}
}
func Test_parseClassWithNullable(t *testing.T) {
	l, _ := Lexer.New(`
   public class prova{
		 public string? name;
   }`)
	class, err := parseClass(l)
	if err != nil {
		panic(err)
	}
	if class.Name!="prova" {
		t.Error("test:class name not found")
	}
	if len(class.Fields)!=1 {
		t.Error("test:class fields not found")
	}
	if class.Fields[0].Name!="name" {
		t.Error("test:class field name not found")
	}
	if class.Fields[0].Type.(SimpleTypeNode).Type!=String {
		t.Error("test:class field type not string")
	}
	if class.Fields[0].Nullable==false {
		t.Error("test:class field not nullable")
	}
}
func Test_parseClassWithNullableGeneric(t *testing.T) {
	parsed, err := ParseStr(`
   public class prova{
		 public pers<Byte>? name;
   }`)
	if err != nil {
		panic(err)
	}
	class:= parsed	
	if class.Name!="prova" {
		t.Error("test:class name not found")
	}
	if len(class.Fields)!=1 {
		t.Error("test:class fields not found")
	}
	if class.Fields[0].Name!="name" {
		t.Error("test:class field name not found")
	}
	gType:=class.Fields[0]
	if gType.Type.(GenericTypeNode).ChildType[0].(SimpleTypeNode).Type!=Number {
		t.Error("test:class field type not number")
	}
	if gType.Nullable==false {
		t.Error("test:class field not nullable")
	}
}

func Test_parseRecord(t *testing.T) {
	parsed, e := ParseStr("record persona(string nome, string cognome,eta<int>? e);")

	if e != nil {
		panic(e)
	}
	record:=parsed

	if record.Name!="persona" {
		t.Error("test:class name not found")
	}
	if len(record.Fields)!=3 {
		t.Error("test:class fields not found")
	}
	if record.Fields[0].Name!="nome" {
		t.Error("test:class field name not found")
	}
	if record.Fields[0].Type.(SimpleTypeNode).Type!=String {
		t.Error("test:class field type not string")
	}
	if record.Fields[1].Name!="cognome" {
		t.Error("test:class field name not found")
	}
	if record.Fields[1].Type.(SimpleTypeNode).Type!=String {
		t.Error("test:class field type not string")
	}
	if record.Fields[2].Name!="e" {
		t.Error("test:class field name not found")
	}
	if record.Fields[2].Type.(GenericTypeNode).ChildType[0].(SimpleTypeNode).Type!=Number {
		t.Error("test:class field type not number")
	}
	if record.Fields[2].Nullable==false {
		t.Error("test:class field not nullable")
	}
	if record.Fields[2].Type.(GenericTypeNode).ParentName!="eta" {
		t.Error("test:class field name not correct")
	}
}
func Test_parseRecordMixed(t *testing.T) {
	record, e := ParseStr("public record persona(string nome){string cognome};")

	if e != nil {
		panic(e)
	}


	if record.Name!="persona" {
		t.Error("test:class name not found")
	}
	if len(record.Fields)!=2{
		t.Error("test:class fields not found")
	}
	if record.Fields[0].Name!="nome" {
		t.Error("test:class field name not found")
	}
	if record.Fields[0].Type.(SimpleTypeNode).Type!=String {
		t.Error("test:class field type not string")
	}
	if record.Fields[1].Name!="cognome" {
		t.Error("test:class field name not found")
	}
	if record.Fields[1].Type.(SimpleTypeNode).Type!=String {
		t.Error("test:class field type not string")
	}
}
func Test_parseClassWithGetSet(t *testing.T) {
	class, err := ParseStr(`
   public class prova{
		 public pers<int>? name{get;set;}
		 public pers<int>? cognome{get;set;}="dddd";
		 public string? eta;
   }`)
	if err != nil {
		panic(err)
	}
	if class.Name!="prova" {
		t.Error("test:class name not found")
	}
	if len(class.Fields)!=3 {
		t.Error("test:class fields not found")
	}
	if class.Fields[0].Name!="name" {
		t.Error("test:class field name not found")
	}
	gType:=class.Fields[0]
	if gType.Type.(GenericTypeNode).ChildType[0].(SimpleTypeNode).Type!=Number {
		t.Error("test:class field type not number")
	}
	if gType.Nullable==false {
		t.Error("test:class field not nullable")
	}
	if class.Fields[1].Name!="cognome" {
		t.Error("test:class field name not found")
	}
	if class.Fields[2].Name!="eta" {
		t.Error("test:class field name not found")
	}
}

func Test_parseClassWithCustomType(t *testing.T) {
	class, err := ParseStr(`
	  public class Person{
			public Prova.persona p{get;set;}
	  }
`)
	if err != nil {
		panic(err)
	}
	if class.Name!="Person" {
		t.Error("test:class name not found")
	}
	if len(class.Fields)!=1 {
		t.Error("test:class fields not found")
	}
	if class.Fields[0].Name!="p" {
		t.Error("test:class field name not found")
	}
	gType:=class.Fields[0]
	if gType.Type.(CustomTypeNode).Type!="persona" {
		t.Error("test:class field type ")
	}
}
func Test_parseClassWithExtends(t *testing.T) {
	class, err := ParseStr(`
	class Lavoratore:Persona{
			public Prova.persona p{get;set;}
	  }
`)
	if err != nil {
		panic(err)
	}
	if class.Name!="Lavoratore" {
		t.Error("test:class name not found")
	}
	if len(class.Fields)!=1 {
		t.Error("test:class fields not found")
	}
	if class.Fields[0].Name!="p" {
		t.Error("test:class field name not found")
	}
	gType:=class.Fields[0]
	println(gType.Type.(CustomTypeNode).Type)
	if gType.Type.(CustomTypeNode).Type!="persona" {
		t.Error("test:class field type ")
	}

	if class.ExtendType[0].(CustomTypeNode).Type!="Persona" {
		t.Error("test:class extends not found")
	}
}
