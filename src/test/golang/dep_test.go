package golang

import (
	"testing"

	"bitwormhole.com/starter/afs/files"
	"bitwormhole.com/starter/base/util"
	"bitwormhole.com/starter/vlog"
)

func TestDep(t *testing.T) {
	util.Close(nil)
	files.FS()
	vlog.GetLogger()
}
