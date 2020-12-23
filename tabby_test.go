package tabby

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
	"text/tabwriter"
)

func Test_buildFormatString(t *testing.T) {
	items := make([]interface{}, 3)
	items[0] = "s1"
	items[1] = "s2"
	items[2] = "s3"
	tabby := &Tabby{}
	fmtString := tabby.buildFormatString(items)
	if fmt.Sprintf("%q", fmtString) != fmt.Sprintf("%q", "%v\t%v\t%v\n") {
		t.Errorf("fmtString incorrect, got: %v, want: %v.", fmt.Sprintf("%q", fmtString), fmt.Sprintf("%q", "%v\t%v\t%v\n"))
	}
}

func Test_New(t *testing.T) {
	tabby := New()
	if tabby.writer == nil {
		t.Errorf("New returning uninitialized writer")
	}
}

func Test_NewCustom(t *testing.T) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	tabby := NewCustom(w)
	if reflect.TypeOf(tabby) != reflect.TypeOf(&Tabby{}) {
		fmt.Println(reflect.TypeOf(tabby))
		t.Errorf("NewCustom incorect type returned")
	}
}

func Test_AddLine(t *testing.T) {
	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 1, '.', 0)
	tabby := NewCustom(w)
	tabby.AddLine("test")
	if b.String() != "test\n" {
		t.Errorf("AddLine not writing to io.Writer")
	}
}

func Test_AddHeader(t *testing.T) {
	var b bytes.Buffer
	w := tabwriter.NewWriter(&b, 0, 0, 1, '.', 0)
	tabby := NewCustom(w)
	tabby.AddHeader("test")
	if b.String() != "test\n----\n" {
		t.Errorf("AddHeader not writing to io.Writer")
	}
}

func Test_FileWriter(t *testing.T) {
	fd, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	w := tabwriter.NewWriter(fd, 0, 0, 4, ' ', 0)
	tabby := NewCustom(w)
	tabby.AddHeader("NAME", "TITLE", "DEPARTMENT")
	tabby.AddLine("John Smith", "Developer", "Engineering")
	tabby.Print()
}

func BenchmarkBuffer(b *testing.B) {
	fd, _ := os.OpenFile("temp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer fd.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buff bytes.Buffer
		buff.WriteString("TestString")
		buff.WriteString("\n")
		buff.WriteTo(fd)
	}
}

func BenchmarkFmt(b *testing.B) {
	fd, _ := os.OpenFile("temp.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	defer fd.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buff bytes.Buffer
		buff.WriteString("TestString")
		fmt.Fprintln(fd, buff.String())
	}
}
