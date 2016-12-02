package nav

import (
	"path/filepath"
	"regexp"
	"strings"
	"github.com/gruntwork-io/docs/errors"
	"github.com/gruntwork-io/docs/file"
	"bytes"
	"fmt"
	"github.com/shurcooL/github_flavored_markdown"
	"html/template"
	"github.com/gruntwork-io/docs/config"
)

const MARKDOWN_LINKS_REGEX = `\]\(([\w-\./]+)\)`
const GRUNTWORK_GITHUB_URL_REGEX = `http[s]?://github.com/gruntwork-io/([A-Za-z0-9_.-]+)(/[A-Za-z0-9_.-]+)*`
const PACKAGE_GITHUB_REPO_URL_PREFIX = "https://github.com/gruntwork-io/<package-name>/tree/master"
const PACKAGE_FILE_REGEX = `^packages/([\w -]+)(/.*)$`
const PACKAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 2
const MARKDOWN_FILE_PATH_REGEX = `^.*/(.*)\.md$`
const MARKDOWN_FILE_PATH_REGEX_NUM_CAPTURE_GROUPS = 1

const CSS_CLASS_NAME_FOR_PRIVATE_GITHUB_URLS = "private-link-modal-link"

// TODO: Figure out better way to reference this file
const HTML_TEMPLATE_MAIN_FILENAME = "main.template.html"
const HTML_TEMPLATE_404_FILENAME = "404.template.html"
const HTML_PAGE_404_FILENAME = "404.html"

// A Page represents a page of documentation, usually formatted as a markdown file.
type Page struct {
	File
	Title        string  // the title of the page
	BodyMarkdown string  // the body of the page as Markdown
	BodyHtml     string  // the body of the page as HTML (does not include surrounding HTML)
	GithubUrl    string  // the Gruntwork Repo GitHub URL to which this page corresponds
	ParentFolder *Folder // the nav folder in which this page resides
	RootFolder   *Folder // the root folder in which this page exists
}

// Populate the remaining properties of this Page instance.
// Note that we cannot populate the body properties here because those depend on the p.RootFolder being fully built out.
func (p *Page) PopulateProperties() error {
	var err error

	p.Title = p.getTitle()

	p.GithubUrl, err = convertPackageLinkPathToGithubUrl(p.InputPath, "./")
	if err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

// Populate the body-related properties.
// Note that this must be done AFTER the NavTree at p.RootFolder is fully built out.
func (p *Page) PopulateBodyProperties(rootOutputPath string, packages []config.GruntworkPackage) error {
	var err error

	p.BodyMarkdown, err = p.getSanitizedMarkdownBody(rootOutputPath)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	p.BodyHtml = getHtmlFromMarkdown(p.BodyMarkdown)

	p.BodyHtml, err = addCssClassToPrivateGitHubUrls(p.BodyHtml, CSS_CLASS_NAME_FOR_PRIVATE_GITHUB_URLS, packages)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

// Add this page to the NavTree that starts at the rootFolder, creating any necessary folders along the way.
func (p *Page) AddToNavTree() error {
	containingFolderPath := getContainingFolder(p.OutputPath)
	containingFolder := p.RootFolder.CreateFolderIfNotExist(containingFolderPath)

	containingFolder.AddPage(p)

	return nil
}

// Return the relative path between this page and the given folder
func (p *Page) GetRelPathToFolder(folder *Folder) string {
	relPath, err := filepath.Rel(p.OutputPath, folder.OutputPath)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while computing the relative path between %s and %s", p.OutputPath, folder.OutputPath))
	}

	// Golang's filepath.Rel() function treats all parts of the path as a directory, even the ending file, so we remove an extraneous ../
	relPath = strings.Replace(relPath, "../", "", 1)

	return relPath
}

// Return the relative path between this page and the given page
func (p *Page) GetRelPathToPage(page *Page) string {
	relPath, err := filepath.Rel(p.OutputPath, page.OutputPath)
	if err != nil {
		panic(fmt.Sprintf("Fatal error while computing the relative path between %s and %s", p.OutputPath, page.OutputPath))
	}

	relPath, err = replaceMdFileExtensionWithHtmlFileExtension(relPath)
	if err != nil {
		panic(fmt.Sprintf("Failed while calling page.replaceMdFileExtensionWithHtmlFileExtension(%s)", relPath))
	}

	// Golang's filepath.Rel() function treats all parts of the path as a directory, even the ending file, so we remove an extraneous ../
	relPath = strings.Replace(relPath, "../", "", 1)

	return relPath
}

// Get the folder that contains the file specified in the given path
func getContainingFolder(path string) string {
	return filepath.Dir(path)
}

// Output the full HTML body of this page
func (p *Page) WriteFullPageHtmlToOutputPath(htmlFilesPath string, rootFolder *Folder, rootOutputPath string) error {
	bodyHtml := p.getBodyHtml()
	navTreeHtml := getNavTreeHtml(rootFolder, p)

	mainHtmlTemplatePath := filepath.Join(htmlFilesPath, "/", HTML_TEMPLATE_MAIN_FILENAME)
	fullHtml, err := getFullHtml(mainHtmlTemplatePath, bodyHtml, navTreeHtml, p.Title, p.GithubUrl, p.InputPath)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	absOutputPath := filepath.Join(rootOutputPath, p.OutputPath)
	absOutputPathDotHtml, err := replaceMdFileExtensionWithHtmlFileExtension(absOutputPath)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	fmt.Printf("Outputting %s to %s\n", p.InputPath, absOutputPathDotHtml)

	err = file.WriteFile(fullHtml, absOutputPathDotHtml)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

func Write404PageToOutputPath(htmlFilesPath string, rootFolder *Folder, rootOutputPath string) error {
	bodyHtml := template.HTML("")
	navTreeHtml := getNavTreeHtml(rootFolder, nil)

	notFoundHtmlTemplatePath := filepath.Join(htmlFilesPath, "/", HTML_TEMPLATE_404_FILENAME)
	fullHtml, err := getFullHtml(notFoundHtmlTemplatePath, bodyHtml, navTreeHtml, "Gruntwork Doc Not Found", "", "")
	if err != nil {
		return errors.WithStackTrace(err)
	}

	absOutputPath := filepath.Join(rootOutputPath, HTML_PAGE_404_FILENAME)

	fmt.Printf("Outputting %s to %s\n", notFoundHtmlTemplatePath, absOutputPath)

	err = file.WriteFile(fullHtml, absOutputPath)
	if err != nil {
		return errors.WithStackTrace(err)
	}

	return nil
}

// Get the NavTree of the given Root Folder with the current page as the "active" page as HTML
func getNavTreeHtml(rootFolder *Folder, activePage *Page) template.HTML {
	return rootFolder.GetAsNavTreeHtml(activePage)
}

// Get the NavTree of the givn Root Folder with the current page as the "active" page as HTML
func (p *Page) getBodyHtml() template.HTML {
	return template.HTML(p.BodyHtml)
}

// Return a NewPage
func NewPage(file *File, rootNavFolder *Folder) *Page {
	return &Page{
		File: *file,
		RootFolder: rootNavFolder,
	}
}

// Get the Page's title from the page's output filename.
func (p *Page) getTitle() string {
	fileNameFull := filepath.Base(p.OutputPath)
	fileNameComponents := strings.Split(fileNameFull, ".")
	title := fileNameComponents[0]
	title = convertDashesToSpacesAndCapitalize(title)
	return strings.Title(title)
}

// Get the Page's markdown body, sanitized for public HTML output (i.e. convert inline links to fully qualified URLs)
func (p *Page) getSanitizedMarkdownBody(rootOutputPath string) (string, error) {
	var body string

	body, err := file.ReadFile(p.FullInputPath)
	if err != nil {
		return body, errors.WithStackTrace(err)
	}

	body, err = convertMarkdownLinksToUrls(p.InputPath, body)
	if err != nil {
		return body, errors.WithStackTrace(err)
	}

	body, err = convertExternalGruntworkGithubUrlsToInternalLinks(body, p.RootFolder, rootOutputPath)
	if err != nil {
		return body, errors.WithStackTrace(err)
	}

	return body, nil
}

// Given a body and its inputPath, convert all paths in the body (e.g. "/foo" or "../bar") to fully qualified URLs.
func convertMarkdownLinksToUrls(inputPath, body string) (string, error) {
	var newBody string

	newBody = body
	linkPaths := getAllLinkPaths(body)

	for _, linkPath := range linkPaths {
		url, err := convertPackageLinkPathToUrl(inputPath, linkPath)
		if err != nil {
			return newBody, errors.WithStackTrace(err)
		}

		// If we blindly replace all instances of our linkPath, some of them will be found in existing URLs!
		// So we wrap them in parentheses (to support markdown-formatted links)...
		linkPathWithParens := fmt.Sprintf("(%s)", linkPath)
		urlWithParens := fmt.Sprintf("(%s)", url)

		newBody = strings.Replace(newBody, linkPathWithParens, urlWithParens, -1)

		// ... and spaces (to support standalone links).
		linkPathWithSpaces := fmt.Sprintf(" %s ", linkPath)
		urlWithSpaces := fmt.Sprintf(" %s ", url)

		newBody = strings.Replace(newBody, linkPathWithSpaces, urlWithSpaces, -1)
	}

	return newBody, nil
}

// Given a body and its inputPath, convert all Gruntwork GitHub URL links to internal links where possible
func convertExternalGruntworkGithubUrlsToInternalLinks(body string, rootNavFolder *Folder, rootOutputPath string) (string, error) {
	var newBody string

	newBody = body
	urls := getGruntworkGithubUrls(body)

	for _, url := range urls {
		internalLink, err := convertGruntworkGithubUrlToInternalLink(url, rootNavFolder, rootOutputPath)
		if err != nil {
			return newBody, errors.WithStackTrace(err)
		}

		newBody = strings.Replace(newBody, url, internalLink, 1)
	}

	return newBody, nil
}

// For each of the given GruntworkPackages, search the HTML body for URLs that point to that repo and add the given cssClass.
// This is important because we want to show a popup window on private links, and we identify such links by their CSS class.
func addCssClassToPrivateGitHubUrls(body string, cssClass string, packages []config.GruntworkPackage) (string, error) {
	var newBody string

	newBody = body

	for _, gPackage := range packages {
		regexStr := fmt.Sprintf(`<a.*href="%s.*".*>`, gPackage.GithubUrl)
		regex, err := regexp.Compile(regexStr)
		if err != nil {
			return newBody, errors.WithStackTrace(err)
		}

		matches := regex.FindAllStringSubmatch(newBody, -1)

		for _, match := range matches {
			oldATag := match[0]
			newATag := strings.Replace(oldATag, "<a ", fmt.Sprintf(`<a class="%s"`, cssClass), -1)

			newBody = strings.Replace(newBody, oldATag, newATag, -1)
		}
	}

	return newBody, nil
}

// Given a body of text find all instances of link paths (e.g. /foo or ../bar)
func getAllLinkPaths(body string) []string {
	var relPaths []string

	regex := regexp.MustCompile(MARKDOWN_LINKS_REGEX)
	submatches := regex.FindAllStringSubmatch(body, -1)

	if len(submatches) == 0 {
		return relPaths
	}

	for _, submatch := range submatches {
		relPath := submatch[1]
		relPaths = append(relPaths, relPath)
	}

	return relPaths
}

// Given a body of text find all instances of Github URLs to other Gruntwork repos
func getGruntworkGithubUrls(body string) []string {
	var urls []string

	regex := regexp.MustCompile(GRUNTWORK_GITHUB_URL_REGEX)
	submatches := regex.FindAllStringSubmatch(body, -1)

	if len(submatches) == 0 {
		return urls
	}

	for _, submatch := range submatches {
		url := submatch[0]
		urls = append(urls, url)
	}

	return urls
}

// Given a github URL like https://github.com/gruntwork-io/..., convert it to the corresponding internal link (e.g. /foo/bar)
func convertGruntworkGithubUrlToInternalLink(githubUrl string, rootNavFolder *Folder, rootOutputPath string) (string, error) {
	linkPath := convertGruntworkGithubUrlToInternalLinkAux("", githubUrl, rootNavFolder)

	if linkPath == "" {
		return githubUrl, nil
	}

	linkPath, err := replaceMdFileExtensionWithHtmlFileExtension(linkPath)
	if err != nil {
		return linkPath, errors.WithStackTrace(err)
	}

	linkPath = "/" + linkPath

	return linkPath, nil
}

// Helper function for convertGruntworkGithubUrlToInternalLink() that recursively scans all pages for a match
func convertGruntworkGithubUrlToInternalLinkAux(linkPath string, githubUrl string, folder *Folder) string {
	if linkPath != "" {
		return linkPath
	}

	for _, page := range folder.ChildPages {
		if page.GithubUrl == githubUrl || page.GithubUrl == githubUrl + "README.md" || page.GithubUrl == githubUrl + "/README.md" {
			linkPath = page.OutputPath
			return linkPath
		}
	}

	for _, childFolder := range folder.ChildFolders {
		linkPath = convertGruntworkGithubUrlToInternalLinkAux(linkPath, githubUrl, childFolder)
	}

	return linkPath
}


// Convert a linkPath (e.g. /foo/bar) that directs to a Package page to a GitHub URL. If no GitHub URL can be created, return the empty string.
func convertPackageLinkPathToGithubUrl(inputPath, linkPath string) (string, error) {
	var url string

	url, err := convertPackageLinkPathToUrl(inputPath, linkPath)
	if err != nil {
		return url, errors.WithStackTrace(err)
	}

	if linkPath == url {
		url = ""
		return url, nil
	}

	return url, nil
}

// Convert a linkPath (e.g. /foo/bar) that directs to a Package page to a fully qualified URL. For non-Package links, just return the original link
func convertPackageLinkPathToUrl(inputPath, linkPath string) (string, error) {
	var url string

	if isPackageInputPath(inputPath) {
		packageName, err := getPackageName(inputPath)
		if err != nil {
			return url, errors.WithStackTrace(err)
		}

		urlPrefix := strings.Replace(PACKAGE_GITHUB_REPO_URL_PREFIX, "<package-name>", packageName, 1)

		if strings.HasPrefix(linkPath, "/") {
			url = urlPrefix + linkPath
		} else if strings.HasPrefix(linkPath, "./") {
			relPath, err := getPathRelativeToPackageRoot(inputPath)
			if err != nil {
				return url, errors.WithStackTrace(err)
			}

			url = urlPrefix + relPath
		} else if strings.HasPrefix(linkPath, "../") {
			relPath, err := getPathRelativeToPackageRoot(inputPath)
			if err != nil {
				return url, errors.WithStackTrace(err)
			}

			relPath = filepath.Join(relPath, "../", linkPath)
			url = urlPrefix + relPath
		} else {
			// linkPath has no / in its first few characters; assume it's a link like "vars.tf"
			relPath, err := getPathRelativeToPackageRoot(inputPath)
			if err != nil {
				return url, errors.WithStackTrace(err)
			}

			relPath = filepath.Join(relPath, "../", linkPath)
			url = urlPrefix + relPath
		}

		return url, nil
	}

	// If this isn't a Package Page, just output the linkPath unmodified
	return linkPath, nil
}

// Return true if the given inputPath is part of a Gruntwork Package
func isPackageInputPath(inputPath string) bool {
	regex := regexp.MustCompile(PACKAGE_FILE_REGEX)
	return regex.MatchString(inputPath)
}

// Extract the Package name from the given inputPath
func getPackageName(inputPath string) (string, error) {
	var subpath string

	regex := regexp.MustCompile(PACKAGE_FILE_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != PACKAGE_FILE_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return subpath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{string: inputPath, regExName: "PACKAGE_FILE_REGEX", regEx: PACKAGE_FILE_REGEX })
	}

	subpath = submatches[0][1]

	return subpath, nil
}

// Given an inputPath like packages/package-vpc/modules/vpc-app, extract the path relative to the package root (i.e. /modules/vpc-app)
func getPathRelativeToPackageRoot(inputPath string) (string, error) {
	var subpath string

	regex := regexp.MustCompile(PACKAGE_FILE_REGEX)
	submatches := regex.FindAllStringSubmatch(inputPath, -1)

	if len(submatches) == 0 || len(submatches[0]) != PACKAGE_FILE_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return subpath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{string: inputPath, regExName: "PACKAGE_FILE_REGEX", regEx: PACKAGE_FILE_REGEX })
	}

	subpath = submatches[0][2]

	return subpath, nil
}

// Given a markdown body return an HTML body
func getHtmlFromMarkdown(markdown string) string {
	bytesInput := []byte(markdown)
	bytesOutput := github_flavored_markdown.Markdown(bytesInput)
	return string(bytesOutput)
}

// Given a path like /foo/bar.md, return /foo/bar.html
func replaceMdFileExtensionWithHtmlFileExtension(path string) (string, error) {
	var updatedPath string

	regex := regexp.MustCompile(MARKDOWN_FILE_PATH_REGEX)
	submatches := regex.FindAllStringSubmatch(path, -1)

	if len(submatches) == 0 || len(submatches[0]) != MARKDOWN_FILE_PATH_REGEX_NUM_CAPTURE_GROUPS + 1 {
		return updatedPath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{string: path, regExName: "MARKDOWN_FILE_PATH_REGEX", regEx: MARKDOWN_FILE_PATH_REGEX })
	}

	filename := submatches[0][1]
	filenameDotMd := fmt.Sprintf("%s.%s", filename, "md")
	filenameDotHtml := fmt.Sprintf("%s.%s", filename, "html")

	updatedPath = strings.Replace(path, filenameDotMd, filenameDotHtml, -1)

	return updatedPath, nil
}

// Return the full HTML rendering of this page within the given HTML template
func getFullHtml(htmlTemplatePath string, pageBodyHtml template.HTML, navTreeHtml template.HTML, pageTitle string, githubUrl string, pageInputPath string) (string, error) {
	var templateOutput string

	type htmlTemplateProperties struct {
		PageTitle     string
		PageBody      template.HTML
		NavTree       template.HTML
		GithubUrl     string
		IsPackagePage bool
	}

	htmlTemplate, err := getTemplate(htmlTemplatePath, pageTitle)
	if err != nil {
		return templateOutput, errors.WithStackTrace(err)
	}

	buf := new(bytes.Buffer)
	err = htmlTemplate.Execute(buf, &htmlTemplateProperties{
		PageTitle: pageTitle,
		PageBody: pageBodyHtml,
		NavTree: navTreeHtml,
		GithubUrl: githubUrl,
		IsPackagePage: isPackageInputPath(pageInputPath),
	})
	if err != nil {
		return templateOutput, errors.WithStackTrace(err)
	}

	templateOutput = buf.String()

	return templateOutput, nil
}

// Get the file at the given path as a *template.Template
func getTemplate(path string, templateTitle string) (*template.Template, error) {
	var templateRef *template.Template

	templateBody, err := file.ReadFile(path)
	if err != nil {
		return templateRef, errors.WithStackTrace(err)
	}

	templateRef, err = template.New(templateTitle).Parse(templateBody)
	if err != nil {
		return templateRef, errors.WithStackTrace(err)
	}

	return templateRef, nil
}

func replaceGithubLinksWithInternalLinks(htmlBody string, rootNavFolder *Folder) {

	// - get all links on page matching github regex
	// - for each link
	//   - if that link can be found in any rootNavFolder children, replace it with the corresponding outputPath of that page
}
