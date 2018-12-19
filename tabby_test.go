package tabby

import (
	"fmt"
	"testing"
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
