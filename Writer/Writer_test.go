package Writer

import (
	"strings"
	"testing"
)

func IsEqual(a, b string) bool {
	var stringsSliceA = strings.Fields(a)
	var stringsSliceB = strings.Fields(b)
	for i := 0; i < len(stringsSliceA); i++ {
		if stringsSliceA[i] != stringsSliceB[i] {
			return false
		}
	}
	return true
}

func Test_WriteSimpleClass(t *testing.T) {

	str, err := Convert(`
	  public class Person{
			public string name;
			public int age;
	  }
		`)
	exp := `
	  export interface Person {
		  name : string;
		  age : number;
    }`
	if err != nil {
		panic(err)
	}
	if !IsEqual(str, exp) {
		t.Error("test:write simple class not found")
	}
}
func Test_WriteSimpleClassWithNullable(t *testing.T) {

	str, err := Convert(`
	  public class Person{
			public string name{get;set;}
			public int age{get;set;}
			public string? cognome{get;set;}
	  }
		`)
	exp := `
	  export interface Person {
			name : string;			
			age : number;
			cognome : string?;
		}`
	if err != nil {
		panic(err)
	}

	if !IsEqual(str, exp) {
		t.Error("test:write simple with nullable class not found")
	}
}

func Test_WriteClassWithGeneric(t *testing.T) {
	str, err := Convert(`
	  public class Person{
			public IEnumerable<List<prova>,float>? name{get;set;}
	  }
		`)
	exp := `
	  export interface Person {
			name : Array<Array<prova>,number>?;
		}`
	if err != nil {
		panic(err)
	}
	if !IsEqual(str, exp) {
		t.Error("test:write simple with nullable class not found")
	}
}
func Test_WriteClassWithCustomType(t *testing.T) {
	str, err := Convert(`
	  public class Person{
			public List<Prova.persona> p{get;set;}
	  }
		`)
	exp := `
	  export interface Person {
			p : Array<Prova.persona>;
		}`
	if err != nil {
		panic(err)
	}

	if !IsEqual(str, exp) {
		t.Error("test:write simple with nullable class not found")
	}
}
func Test_WriteClassWithExtends(t *testing.T) {
	str, err := Convert(`
	public class Lavoratore:Person{
			public List<Lavoro.tipo> lavori{get;set;}
	  }
		`)
	exp := `
		export interface Lavoratore extends Person {
				lavori : Array<Lavoro.tipo>;
		}`
	if err != nil {
		panic(err)
	}
   println(str)
	if !IsEqual(str, exp) {
		t.Error("test:write simple with nullable class not found")
	}
}