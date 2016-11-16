package nav

// A Page represents a page of documentation, usually formatted as a markdown file.
type Page struct {
	File
	Title        string  // the title of the page
	MarkdownBody string  // the body of the page as HTML (does not include surrounding HTML)
	GithubUrl    string  // the Gruntwork Repo GitHub URL to which this page corresponds
	ParentFolder *Folder // the nav folder in which this page resides
}

// TODO
func (p *Page) PopulateAllProperties() error {
	return nil
}

// TODO
func (p *Page) AddToNavTree(rootFolder *Folder) error {
	return nil
}

// TODO
func (p *Page) WriteToOutputPathAsHtml() error {
	return nil
}

func NewPage(file *File) *Page {
	return &Page{
		File: *file,
	}
}