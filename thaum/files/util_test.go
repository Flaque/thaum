package files

import (
	filet "github.com/Flaque/filet"
	testUtil "github.com/Flaque/thaum/thaum/testingutil"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

// Tests the exists() function
func TestExists(t *testing.T) {
	defer filet.CleanUp(t)

	assert := assert.New(t)

	tmp := filet.TmpFile(t, "", "")
	assert.True(exists(tmp.Name()), "Returned false but the file should exist")

	assert.False(exists("IDontExist.jpg"), "Returned true by the file should not exist")
}

// Tests the existsAbove() function
func TestExistsAbove(t *testing.T) {
	defer filet.CleanUp(t)

	// Create our test environment
	outerDir := filet.TmpDir(t, "")       // Outermost directory
	queryDir := filet.TmpDir(t, outerDir) // What we're looking for
	subDir := filet.TmpDir(t, outerDir)   // Some random directory
	subsubDir := filet.TmpDir(t, subDir)  // Another random dir inside the subDir

	// Run tests

	assert := assert.New(t)

	query := filepath.Base(queryDir)
	result, err := existsAbove(subsubDir, query)
	if err != nil {
		assert.FailNow("Error returned when calling existsAbove")
	}

	assert.Equal(result, queryDir)

	query = "IDontExistTrustMe"
	_, err = existsAbove(subsubDir, query)

	assert.EqualError(err, `Thaum can't find your thaum_files anywhere in the parent directories!`)
}

// More or less tests cwd()
func TestCwd(t *testing.T) {
	defer filet.CleanUp(t)

	assert.True(t, len(cwd()) > 0)
}

func TestThaumTemplates(t *testing.T) {
	defer filet.CleanUp(t)

	assert := assert.New(t)

	mySrc, _, template := testUtil.TmpThaumEnvironment(t, "")

	os.Chdir(mySrc)
	ts, err := ThaumTemplates()

	assert.NoError(err)
	assert.Equal(ts[0], filepath.Base(template), "Incorrect template returned")

}

func TestIsDsStore(t *testing.T) {
	const testPath = "path/to/.DS_Store"
	assert.True(t, IsDsStore(testPath))
}

func TestRemoveThaumExtension(t *testing.T) {
	assert.Equal(t, RemoveThaumExtension("this.thaum"),
		"this", true)
}
