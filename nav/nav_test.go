package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const EMPTY_STRING = ""

func TestContainsFolder1(t *testing.T) {
	t.Parallel()

	root := NewRootFolder()
	a := NewFolder(EMPTY_STRING, "a")
	b := NewFolder(EMPTY_STRING, "b")
	c := NewFolder(EMPTY_STRING, "c")
	d := NewFolder(EMPTY_STRING, "d")
	e := NewFolder(EMPTY_STRING, "e")
	f := NewFolder(EMPTY_STRING, "f")
	g := NewFolder(EMPTY_STRING, "g")
	h := NewFolder(EMPTY_STRING, "h")

	allFolders := []*Folder{ root, a, b, c, d, e, f, g, h }

	// NavTree #1
	root.AddFolder(a)
		a.AddFolder(b)
			b.AddFolder(c)
			b.AddFolder(d)
		a.AddFolder(e)
	root.AddFolder(f)
		f.AddFolder(g)

	actual := root.ContainsFolder("d")
	assert.True(t, actual, "The given folder should have been found in NavTree %s but was not.", "#1")

	actual = a.ContainsFolder("g")
	assert.False(t, actual, "The given folder should NOT have been found in NavTree %s but was.", "#1")

	// NavTree #2
	resetFolders(allFolders)

	root.AddFolder(a)
	root.AddFolder(b)
	root.AddFolder(c)
	root.AddFolder(d)

	actual = root.ContainsFolder("e")
	assert.False(t, actual, "The given folder should NOT have been found in NavTree %s but was.", "#2")

	// NavTree #3
	resetFolders(allFolders)

	root.AddFolder(a)
	root.AddFolder(b)
	root.AddFolder(c)
		c.AddFolder(d)
		c.AddFolder(e)
		c.AddFolder(f)
			f.AddFolder(g)
			f.AddFolder(h)

	actual = root.ContainsFolder("h")
	assert.True(t, actual, "The given folder should have been found in NavTree %s but was not.", "#3")

	actual = f.ContainsFolder("h")
	assert.True(t, actual, "The given folder should have been found in NavTree %s but was not.", "#3")

	actual = b.ContainsFolder("e")
	assert.False(t, actual, "The given folder should NOT have been found in NavTree %s but was.", "#3")
}

func resetFolders(folders []*Folder) {
	for _, folder := range folders {
		folder.parentFolder = nil
		folder.childFolders = nil
		folder.childPages = nil
	}
}