package tabby

import (
	"bytes"
	"fmt"
	"os"
	"text/tabwriter"
)

// Tabby is returned when New() is called.
type Tabby struct {
	writer *tabwriter.Writer
}

// New returns a new *tabwriter.Writer with default config
func New() *Tabby {
	return &Tabby{
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0),
	}
}

// NewCustom returns a new *tabwriter.Writer with custom config
func NewCustom(minWidth, tabWidth, padding int, padchar byte, flags uint) *Tabby {
	return &Tabby{
		writer: tabwriter.NewWriter(os.Stdout, minWidth, tabWidth, padding, padchar, flags),
	}
}

// AddLine will write a new table line
func (t *Tabby) AddLine(args ...interface{}) {
	var b bytes.Buffer
	// Build the format string
	for idx := range args {
		b.WriteString("%v")
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	fmt.Fprintln(t.writer, fmt.Sprintf(b.String(), args...))
}

// AddHeader will write a new table header
func (t *Tabby) AddHeader(args ...interface{}) {
	var b bytes.Buffer
	for idx := range args {
		b.WriteString("%v")
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	fmt.Fprintln(t.writer, fmt.Sprintf(b.String(), args...))
	t.AddSeperator(args)
}

// AddSeperator will write a new dash seperator line based on the args length
func (t *Tabby) AddSeperator(args []interface{}) {
	var b bytes.Buffer
	for idx, arg := range args {
		length := len(fmt.Sprintf("%v", arg))
		b.WriteString(dashes(length))
		if idx+1 != len(args) {
			// Add a tab as long as its not the last column
			b.WriteString("\t")
		}
	}
	fmt.Fprintln(t.writer, b.String())
}

// Print will write the table to the terminal
func (t *Tabby) Print() {
	t.writer.Flush()
}

// dashes generates dash strings for heading borders
func dashes(l int) string {
	var b bytes.Buffer
	for i := 0; i < l; i++ {
		b.WriteString("-")
	}
	return b.String()
}
