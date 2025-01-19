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
	class Person{
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
		cognome? : string;
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
	record Person{
		public IEnumerable<List<prova>,float>? name{get;set;}
	}
	`)
	exp := `
	export interface Person {
		name? : Array<Array<prova>,number>;
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
func Test_WriteTwoSimpleRecord(t *testing.T) {

	str, err := Convert(`
	record Person1(
		string name1,
		int age1
	);
	record Person2{
		public string name2;
		public int age2;
	}`)
	exp := `
	export interface Person1 {
		name1 : string;
		age1 : number;
	}
	export interface Person2 {
		name2 : string;
		age2 : number;
	}`
	if err != nil {
		panic(err)
	}
	if !IsEqual(str, exp) {
		t.Error("test:write simple class not found")
	}
}
func Test_WriteTwoSimpleClass2(t *testing.T) {

	str, err := Convert(`
	record Person1{
		public persona.persona p;
		public int age1;
	}
	record Person2{
		public int age2;
	}`)
	exp := `
	export interface Person1 {
		p : persona;
		age1 : number;
	}
	export interface Person2 {
		age2 : number;
	}`
	if err != nil {
		panic(err)
	}
	if !IsEqual(str, exp) {
		t.Error("test:write simple class not found")
	}
}

func Test_WriteDto1(t *testing.T) {

	str, err := Convert(`
	public class dfdksfds{
		public IEnumerable<Ti> downlaod = Enumerable.Empty<Ti>();
	}
	`)
	exp := `
	export interface dfdksfds {
		downlaod : Array<Ti>;
	}
	`
	if err != nil {
		panic(err)
	}
	if !IsEqual(str, exp) {
		t.Error("test:no match")
	}
}

func Test_WriteDto2(t *testing.T) {

	str, err := Convert(`
	public record Tfd(String ModelKi, Telemat Source, Int32? Suppl, String? ccc, DateTime? CreationDate)
	{
		public IEnumerable<String> Errors                                  { get; init; };
		public IEnumerable<Pdf> CombinedDetails { get; init; };
	}
	`)
	exp := `
	export interface Tfd {
		ModelKi : String;
		Source : Telemat;
		Suppl? : number;
		ccc? : String;
		CreationDate? : Date;
		Errors : Array<String>;
		CombinedDetails : Array<Pdf>;
	}
	`
	if err != nil {
		panic(err)
	}
	if !IsEqual(str, exp) {
		t.Error("test:no match")
	}
}
func Test_WriteDto3(t *testing.T) {

	str, err := Convert(`
	public class pdfff {
		public IEnumerable<tii> repord = Enumerable.Empty();
	}

	//qui mettere tutto su record. lasciando Errors e SupplyNumber. Comunque possono essere messi sul costruttore.
	public record pred(String ModelKind, Tele Source, Int32? Sup = null, String? pro = null, DateTime? CreationDate = null)
	{
		public IEnumerable Errors { get; init; } = Enumerable.Empty();
		public IEnumerable CombinedDetails { get; init; } = Enumerable.Empty();
	}

	public record predd(Boolean ConfirmationFlag = false, Boolean IsFlag = false) {
		public IEnumerable Details { get; init; } = Enumerable.Empty();
	}`)
	exp := `
	export interface pdfff {
		repord : Array<tii>;
	}

	export interface pred {
		ModelKind : String;
		Source : Tele;
		Sup? : number;
		pro? : String;
		CreationDate? : Date;
		Errors : IEnumerable;
		CombinedDetails : IEnumerable;
	}

	export interface predd {
		ConfirmationFlag : Boolean;
		IsFlag : Boolean;
		Details : IEnumerable;
	}`
	if err != nil {
		panic(err)
	}
	println(str)
	if !IsEqual(str, exp) {
		t.Error("test:no match")
	}
}
