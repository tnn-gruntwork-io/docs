package nav

import "fmt"

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

func (f *Folder) AddFolder(folder *Folder) error {
	fmt.Printf("Adding %s folder to %s\n", folder.name, f.name)
	return nil
}

func (f *Folder) AddPage(page *Page) error {
	fmt.Printf("Adding %s page to %s\n", page.title, f.name)
	return nil
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

// Sample function only. Actual pages and folders will be added automatically by the docs-processor.
func SetupNav() {
	root := NewRootFolder()
	intro := NewFolder("", "Introduction")
	overview := NewPage("", "Overview", "", "")
	tools := NewPage("", "Tools", "", "")

	packages := NewFolder("", "Packages")
	networkTopology := NewFolder("", "Network Topology")
	networkTopologyOverview := NewPage("", "Overview", "", "")
	vpcApp := NewFolder("", "vpc-app")
	vpcAppOverview := NewPage("", "Overview", "", "")


	root.AddFolder(intro)
	intro.AddPage(overview)
	intro.AddPage(tools)

	root.AddFolder(packages)
	packages.AddFolder(networkTopology)
	networkTopology.AddPage(networkTopologyOverview)
	networkTopology.AddFolder(vpcApp)
	vpcApp.AddPage(vpcAppOverview)
}