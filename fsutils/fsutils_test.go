package fsutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cwd string = "."
var nonExistantPath string = "non-existant-path"
var knownGoodFile string = "fsutils.go"

func TestPathExists(t *testing.T) {
	assert.True(t, PathExists(cwd))
	assert.False(t, PathExists(nonExistantPath))
}

func TestIsFile(t *testing.T) {
	assert.True(t, IsFile(knownGoodFile))
	assert.False(t, IsFile(cwd))
	assert.False(t, IsFile(nonExistantPath))
}

func TestIsDir(t *testing.T) {
	assert.False(t, IsDir(knownGoodFile))
	assert.True(t, IsDir(cwd))
	assert.False(t, IsDir(nonExistantPath))
}
