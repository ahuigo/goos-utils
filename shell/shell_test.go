// BEGIN: abpxx6d04wxr
package shell

import (
	"testing"
)

func TestExecCommand(t *testing.T) {
	stdout, stderr, errno := ExecCommand("sh", "-c", `echo abc;ls | wc -l `)
	t.Log("stdout:", stdout)
	t.Log("stderr:", stderr)
	t.Log("errno:", errno)
}
