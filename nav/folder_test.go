package nav

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestFolder_StandardizePath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		path     string
		expected string
	}{
		{path: "/a/b/c", expected: "a/b/c" },
		{path: "./a/b/c", expected: "a/b/c" },
		{path: "a/b/c", expected: "a/b/c" },
		{path: "b/c", expected: "b/c" },
		{path: "./b/c", expected: "b/c" },
		{path: "/b/c", expected: "b/c" },
	}

	for _, testCase := range testCases {
		actual := getStandardizedPath(testCase.path)
		assert.Equal(t, testCase.expected, actual, "path = %s\n", testCase.path)
	}
}

func TestFolder_GetTopFolderNameInPath(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		folderPath                  string
		expectedFolderName          string
		expectedNumRemainingFolders int
	}{
		{folderPath: "/a/b/c", expectedFolderName: "a", expectedNumRemainingFolders: 2 },
		{folderPath: "./a/b/c", expectedFolderName: "a", expectedNumRemainingFolders: 2 },
		{folderPath: "a/b/c", expectedFolderName: "a", expectedNumRemainingFolders: 2 },
		{folderPath: "b/c", expectedFolderName: "b", expectedNumRemainingFolders: 1 },
		{folderPath: "./b/c", expectedFolderName: "b", expectedNumRemainingFolders: 1 },
		{folderPath: "/b/c", expectedFolderName: "b", expectedNumRemainingFolders: 1 },
		{folderPath: "c", expectedFolderName: "c", expectedNumRemainingFolders: 0 },
	}

	for _, testCase := range testCases {
		actualName, actualNumRemainingFolders := getTopFolderNameInPath(testCase.folderPath)
		assert.Equal(t, testCase.expectedFolderName, actualName, "folderPath = %s\n", testCase.folderPath)
		assert.Equal(t, testCase.expectedNumRemainingFolders, actualNumRemainingFolders, "folderPath = %s\n", testCase.folderPath)
	}
}

func TestFolder_GetFolderPathTail(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		folderPath string
		expected   string
	}{
		{folderPath: "/a/b/c", expected: "b/c"},
		{folderPath: "a/b/c", expected: "b/c"},
		{folderPath: "./a/b/c", expected: "b/c"},
		{folderPath: "/1/2/3/4/5", expected: "2/3/4/5"},
	}

	for _, testCase := range testCases {
		actual := getFolderPathTail(testCase.folderPath)
		assert.Equal(t, testCase.expected, actual, "folderPath = %s\n", testCase.folderPath)
	}
}

func TestFolder_HasChildFolder(t *testing.T) {
	t.Parallel()

	// Setup test folder structure
	navTreeAsPath := "/a/b/c"

	rootFolder := NewRootFolder()
	a := NewFolder("/a", "a")
	b := NewFolder("/b", "b")
	c := NewFolder("a/b/c", "c")

	rootFolder.AddFolder(a)
	a.AddFolder(b)
	b.AddFolder(c)

	// Validate function
	aHasB := a.HasChildFolder("b")
	aHasC := a.HasChildFolder("c")
	bHasC := b.HasChildFolder("c")

	assert.True(t, aHasB, "navTree = %s\n", navTreeAsPath)
	assert.False(t, aHasC, "navTree = %s\n", navTreeAsPath)
	assert.True(t, bHasC, "navTree = %s\n", navTreeAsPath)
}

func TestFolder_CreateFolderIfNotExist1(t *testing.T) {
	t.Parallel()

	// Setup test folder structure
	rootFolder := NewRootFolder()
	a := NewFolder("a", "a")
	b := NewFolder("a/b", "b")
	c := NewFolder("a/b/c", "c")

	rootFolder.AddFolder(a)
	a.AddFolder(b)
	b.AddFolder(c)

	// Validate function
	d := rootFolder.CreateFolderIfNotExist("/a/b/c/d")

	expectedFolderName := "d"
	expectedFolderPath := "a/b/c/d"
	expectedParentFolder := c

	assert.Equal(t, expectedFolderName, d.Name, "%v\n", d)
	assert.Equal(t, expectedFolderPath, d.OutputPath, "%v\n", d)
	assert.Equal(t, expectedParentFolder, d.ParentFolder, "%v\n", d)
	assert.Equal(t, 1, len(c.ChildFolders), "%v\n", d)
	assert.Equal(t, d, c.ChildFolders[0], "%v\n", d)
}

func TestFolder_CreateFolderIfNotExist2(t *testing.T) {
	t.Parallel()

	// Setup test folder structure
	rootFolder := NewRootFolder()

	// Validate function
	d := rootFolder.CreateFolderIfNotExist("/a/b/c/d")

	expectedFolderName := "d"
	expectedFolderPath := "a/b/c/d"
	expectedParentFolderName := "c"

	assert.Equal(t, expectedFolderName, d.Name, "%v\n", d)
	assert.Equal(t, expectedFolderPath, d.OutputPath, "%v\n", d)
	assert.Equal(t, expectedParentFolderName, d.ParentFolder.Name, "%v\n", d)
	assert.Equal(t, 1, len(rootFolder.ChildFolders), "%v\n", d)
}

func TestFolder_CreateFolderIfNotExist3(t *testing.T) {
	t.Parallel()

	// Setup test folder structure
	rootFolder := NewRootFolder()
	hello := NewFolder("hello", "hello")
	world := NewFolder("world", "world")
	nadaa := NewFolder("hello/world/nadaa", "nadaa")

	rootFolder.AddFolder(hello)
	hello.AddFolder(world)
	world.AddFolder(nadaa)

	// Validate function
	jim := rootFolder.CreateFolderIfNotExist("hello/josh/jim")

	expectedFolderName := "jim"
	expectedFolderPath := "hello/josh/jim"
	expectedParentFolderName := "josh"

	assert.Equal(t, expectedFolderName, jim.Name, "%v\n", jim)
	assert.Equal(t, expectedFolderPath, jim.OutputPath, "%v\n", jim)
	assert.Equal(t, expectedParentFolderName, jim.ParentFolder.Name, "%v\n", jim)
	assert.Equal(t, 2, len(hello.ChildFolders), "%v\n", hello)
}