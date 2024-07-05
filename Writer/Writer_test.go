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
			cognome : string | null;
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
			name : Array<Array<prova>,number> | null;
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
			p : Array<persona>;
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
				lavori : Array<tipo>;
		}`
	if err != nil {
		panic(err)
	}
	if !IsEqual(str, exp) {
		t.Error("test:write simple with nullable class not found")
	}
}
func Test_WriteClass1(t *testing.T) {
	_, err := Convert(`
	public class TelematicInfo
{
   public short? AnnoModello       { get; init; }
   public bool Accolta             { get; init; }
   public bool? Scartata           { get; init; }
   public bool AccoltaParzialmente { get; init; }
   public string AdEProtocol       { get; init; }
   public bool FlagISA             { get; init; }
   public bool FlagConferma        { get; init; }
   public int CounterIVP           { get; init; }
   public Byte DocumentType        { get; init; }
   public bool ManuallySent        { get; init; }
   public DateTime? SentDate       { get; init; }
   public int? IdCustomer          { get; init; }
   public int? IdWfInstanceKey     { get; init; }
   public string State             { get; init; }
   public short YearRelease        { get; init; }
   public string SupplyCode        { get; init; }
   public string ModelCode         { get; init; }
   public int TelematicSource      { get; init; }
   public string? TakerUser        { get; init; }
   public DateTime? TakerDate      { get; init; }
	 public string? QuittanceUser    { get; init; }
   public DateTime? QuittanceDate  { get; init; }
}		`)
	if err != nil {
		panic(err)
	}
}
