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
	println(typ.Type)
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
	l, _ := Lexer.New("public prova<int,string> name;")

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
		 public string name;
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
	l, _ := Lexer.New(`
   public class prova{
		 public pers<int>? name;
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
	gType:=class.Fields[0]
	if gType.Type.(GenericTypeNode).ChildType[0].(SimpleTypeNode).Type!=Number {
		t.Error("test:class field type not number")
	}
	if gType.Nullable==false {
		t.Error("test:class field not nullable")
	}
}