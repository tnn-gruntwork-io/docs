package docfile

import (
	"fmt"
	"strings"

	"github.com/gruntwork-io/docs/logger"
	"github.com/gruntwork-io/docs/file"
	"github.com/gruntwork-io/docs/errors"
	"github.com/gruntwork-io/docs/nav"
)

const IS_GLOBAL_DOC_REGEX = `^global/(.*\.md)$`

// Represents a non-overview document that's part of a specific module.
type GlobalDoc struct {
	relPath string
	absPath string
}

func NewGlobalDoc(absPath string, relPath string) *GlobalDoc {
	return &GlobalDoc { absPath: absPath, relPath: relPath }
}

func (d *GlobalDoc) IsMatch() bool {
	return checkRegex(d.relPath, IS_GLOBAL_DOC_REGEX)
}

func (d *GlobalDoc) AddToNavTree(rootFolder *nav.Folder) error {

	// Create the page corresponding to this document
	page := nav.NewPage(d.absPath, d.getTitle(), )

	// ...
	folderName := d.getFolderName()

	// Make into a function addFolderIfNotExist(folderName)
	if ! rootFolder.ContainsFolder(folderName) {
		folderPath := d.getFolderPath()
		globalSection := nav.NewFolder(folderPath, folderName)
		rootFolder.AddFolder(globalSection)
	}

	containingFolder := rootFolder.GetFolder(folderName)

	containingFolder.AddFolder()



	return nil
}

func (d *GlobalDoc) Copy(outputPathRoot string) error {
	outRelPath := d.getRelOutputPath()
	outAbsPath := fmt.Sprintf("%s/%s", outputPathRoot, outRelPath)

	logger.Logger.Printf("Copying GLOBAL-DOC file %s to %s\n", d.absPath, outAbsPath)
	err := file.CopyFile(d.absPath, outAbsPath)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

func (d *GlobalDoc) getRelOutputPath() string {
	return strings.Replace(d.relPath, "global/", "", -1)
}

func (d *GlobalDoc) getFolderName() string {
	return "intro"
}

func (d *GlobalDoc) getFolderPath() string {
	return "global/intro"
}

func (d *GlobalDoc) getTitle() string {
	return "something"
}

func (d *GlobalDoc) getHtmlBody() string {
	return ""
}