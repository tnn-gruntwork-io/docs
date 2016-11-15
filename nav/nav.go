package nav

type Page struct {
	path         string  // the path where this page will exist when finally output
	title        string  // the title of the page
	htmlBody     string  // the body of the page as HTML (does not include surrounding HTML)
	githubUrl    string  // the Gruntwork Repo GitHub URL to which this page corresponds
	parentFolder *Folder // the folder in which this page resides
}

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
	childPage.parentFolder = f
}

func (f *Folder) AddToFolder(parentFolder *Folder) {
	f.parentFolder = parentFolder
	parentFolder.childFolders = append(parentFolder.childFolders, f)
}

// Returns true if this folder is the folderName, or any of its recursive children
func (f *Folder) ContainsFolder(folderName string) bool {
	if f.name == folderName {
		return true
	}

	for _, folder := range f.childFolders {
		if folder.ContainsFolder(folderName) {
			return true
		}
	}

	return false
}

func (p *Page) AddToFolder(parentFolder *Folder) {
	p.parentFolder = parentFolder
	parentFolder.childPages = append(parentFolder.childPages, p)
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

func NewPage(path, title, htmlBody, githubUrl string) *Page {
	return &Page{
		path: path,
		title: title,
		htmlBody: htmlBody,
		githubUrl: githubUrl,
	}
}