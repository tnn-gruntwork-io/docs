package nav

import (
	"strings"
	"github.com/gruntwork-io/docs/util"
	"path/filepath"
	"fmt"
)

type Folder struct {
	path         string    // the path where this folder will exist when finally output
	name         string    // the name of the folder
	childPages   []*Page   // the list of pages contained in this folder
	childFolders []*Folder // the list of folders containers in this folder
	parentFolder *Folder   // the folder in which this folder resides
	isRoot       bool      // true if this is the topmost folder
}

func (f *Folder) AddFolder(childFolder *Folder) {
	f.childFolders = append(f.childFolders, childFolder)
	childFolder.parentFolder = f
}

func (f *Folder) AddPage(childPage *Page) {
	f.childPages = append(f.childPages, childPage)
	childPage.ParentFolder = f
}

func (f *Folder) AddToFolder(parentFolder *Folder) {
	f.parentFolder = parentFolder
	parentFolder.childFolders = append(parentFolder.childFolders, f)
}

// Returns true if this folder or any of its recursive children have the given folderName
func (f *Folder) ContainsFolderRecursive(folderName string) bool {
	if f.name == folderName {
		return true
	}

	for _, folder := range f.childFolders {
		if folder.ContainsFolderRecursive(folderName) {
			return true
		}
	}

	return false
}

// Returns true if this folder or any of its direct children (but no further descendants) have the given folderName
func (f *Folder) HasChildFolder(folderName string) bool {
	for _, folder := range f.childFolders {
		if folder.name == folderName {
			return true
		}
	}

	return false
}

// Returns the direct child folder of the given name, or nil if not found
func (f *Folder) GetChildFolder(folderName string) *Folder {
	for _, folder := range f.childFolders {
		if folder.name == folderName {
			return folder
		}
	}

	return nil
}

// Returns the given folder if it exists in the current folder or any recursive child folder. Otherwise returns nil.
func (f *Folder) GetFolder(folderName string) *Folder {
	if f.name == folderName {
		return f
	}

	for _, folder := range f.childFolders {
		if folder.GetFolder(folderName) != nil {
			return folder
		}
	}

	return nil
}

// Given a folderPath x/y/z, create each such folder if it does not already exist
func (f *Folder) CreateFolderIfNotExist(folderPath string) *Folder {
	folderPath = getStandardizedPath(folderPath)
	folderNameToCreate, numRemainingFolders := getTopFolderNameInPath(folderPath)

	// Base case
	if numRemainingFolders == 0 {
		if f.HasChildFolder(folderNameToCreate) {
			return f.GetChildFolder(folderNameToCreate)
		} else {
			newChildFolderPath := filepath.Join(f.path, folderNameToCreate)
			childFolder := NewFolder(newChildFolderPath, folderNameToCreate)
			f.AddFolder(childFolder)

			return childFolder
		}
	}

	// Recursive Case
	var childFolder *Folder

	if f.HasChildFolder(folderNameToCreate) {
		childFolder = f.GetChildFolder(folderNameToCreate)
	} else {
		newChildFolderPath := filepath.Join(f.path, folderNameToCreate)
		childFolder = NewFolder(newChildFolderPath, folderNameToCreate)
		f.AddFolder(childFolder)
	}

	folderPathTail := getFolderPathTail(folderPath)

	return childFolder.CreateFolderIfNotExist(folderPathTail)
}

// Given a folderPath such as /x/y/z or ./x/y/z, return the top folder name
func getTopFolderNameInPath(folderPath string) (string, int) {
	folderPath = getStandardizedPath(folderPath)

	folderNames := strings.Split(folderPath, "/")
	numRemainingFolders := len(folderNames) - 1

	return folderNames[0], numRemainingFolders
}

// Given a folderPath such as /x/y/z or ./x/y/z, return the top folder name
func getFolderPathTail(folderPath string) string {
	folderPath = getStandardizedPath(folderPath)

	folderNames := strings.Split(folderPath, "/")
	folderNamesTail := util.GetStrSliceTail(folderNames)
	folderNamesTailStr := strings.Join(folderNamesTail, "/")

	return folderNamesTailStr
}

// Convert a path of the form ./x/y/z, /x/y/z, or x/y/z to the form x/y/z
func getStandardizedPath(path string) string {
	if strings.HasPrefix(path, "/") {
		path = strings.Replace(path, "/", "", 1)
	}

	if strings.HasPrefix(path, "./") {
		path = strings.Replace(path, "./", "", 1)
	}

	return path
}

// TODO
func (f *Folder) OutputAllFilesAsHtml() error {
	return nil
}

// Print the entire tree of a given folder
func (f *Folder) PrintFolderTree() {
	f.printFolderTreeAux(0)
}

// Helper function for printing a complete tree
func (f *Folder) printFolderTreeAux(folderDepth int) {
	fmt.Printf("%s", strings.Repeat("- ", folderDepth))
	fmt.Printf("FOLDER: %s\n", f.name)

	for _, folder := range f.childFolders {
		folder.printFolderTreeAux(folderDepth + 1)
	}

	for _, page := range f.childPages {
		fmt.Printf("%s", strings.Repeat("- ", folderDepth + 1))
		fmt.Printf("%s\n", page.Title)
	}
}

// Print a nicely formatted string of the folder
func (f *Folder) PrintFolder() {
	var parentFolderName string
	var childFolders string
	var childPages string

	if f.parentFolder != nil {
		parentFolderName = f.parentFolder.name
	}

	if f.childFolders != nil {
		childFolderNames := []string{}
		for _, childFolder := range f.childFolders {
			childFolderNames = append(childFolderNames, childFolder.name)
		}
		childFolders = fmt.Sprintf("%v", childFolderNames)
	}

	if f.childPages != nil {
		childPageNames := []string{}
		for _, childPage := range f.childPages {
			childPageNames = append(childPageNames, childPage.Title)
		}
		childPages = fmt.Sprintf("%v", childPageNames)
	}

	fmt.Printf("[ name=%s, path=%s, parentFolder=%s, childFolders=%s, childPages=%s ]\n",
		f.name,
		f.path,
		parentFolderName,
		childFolders,
		childPages,
	)
}

func NewRootFolder() *Folder {
	return &Folder{
		name: "ROOT-FOLDER",
		isRoot: true,
	}
}

func NewFolder(path, name string) *Folder {
	return &Folder{
		path: path,
		name: name,
	}
}

