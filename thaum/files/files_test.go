package files

import (
	filet "github.com/flaque/filet"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCompiledFile(t *testing.T) {
	defer filet.CleanUp(t)

	inPath := filet.TmpFile(t, filet.TmpDir(t, ""), "").Name()
	outPath := filet.TmpDir(t, "") + "/test"

	createCompiledFile(TemplateFile{"", inPath, make(map[string]string)},
		outPath)

	assert.True(t, exists(outPath), "outpath does not exist")
}

func TestCreateCompiledDir(t *testing.T) {
	defer filet.CleanUp(t)

	outPath := filet.TmpDir(t, "") + "/test"

	createCompiledDir(outPath)

	assert.True(t, exists(outPath), "outpath does not exist")
}
