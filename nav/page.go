package nav

import (
	"path/filepath"
	"regexp"
	"strings"
	//"github.com/gruntwork-io/docs/file"
	"github.com/gruntwork-io/docs/errors"
	"github.com/gruntwork-io/docs/file"
)

const FILE_PATHS_REGEX = `(?:http:/|https:/)?(/[A-Za-z0-9_/.-]+)|([A-Za-z0-9_/.-]+/[A-Za-z0-9_.-]*)`
const PACKAGE_GITHUB_REPO_URL_PREFIX = "https://github.com/gruntwork-io/<package-name>/tree/master"
const PACKAGE_FILE_REGEX = `^packages/([\w -]+)(/.*)$`
const PACKAGE_FILE_REGEX_NUM_CAPTURE_GROUPS = 2

// A Page represents a page of documentation, usually formatted as a markdown file.
type Page struct {
	File
	Title        string  // the title of the page
	MarkdownBody string  // the body of the page as HTML (does not include surrounding HTML)
	GithubUrl    string  // the Gruntwork Repo GitHub URL to which this page corresponds
	ParentFolder *Folder // the nav folder in which this page resides
}

// Populate all the remaining properties of this Page instance
func (p *Page) PopulateAllProperties() error {
	var err error

	p.Title = p.getTitle()

	p.MarkdownBody, err = p.getSanitizedMarkdownBody()
	if err != nil {
		return errors.WithStackTrace(err)
	}

	p.GithubUrl, err = convertPackageLinkToUrl(p.InputPath, "./")
	if err != nil {
		return errors.WithStackTrace(err)
	}

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

// Get the Page's title from the page's output filename.
func (p *Page) getTitle() string {
	fileNameFull := filepath.Base(p.OutputPath)
	fileNameComponents := strings.Split(fileNameFull, ".")
	title := fileNameComponents[0]
	return strings.Title(title)
}

// Get the Page's markdown body, sanitized for public HTML output (i.e. convert inline links to fully qualified URLs)
func (p *Page) getSanitizedMarkdownBody() (string, error) {
	var body string

	body, err := file.ReadFile(p.FullInputPath)
	if err != nil {
		return body, errors.WithStackTrace(err)
	}

	body, err = convertMarkdownLinksToUrls(body, p.InputPath)
	if err != nil {
		return body, errors.WithStackTrace(err)
	}

	return body, nil
}

// Given a doc file with the given body at the given inputPath, convert all paths in the body (e.g. "/foo" or "../bar")
// to fully qualified URLs.
func convertMarkdownLinksToUrls(inputPath, body string) (string, error) {
	var newBody string

	newBody = body
	linkPaths := getAllLinkPaths(body)

	for _, linkPath := range linkPaths {
		url, err := convertPackageLinkToUrl(inputPath, linkPath)
		if err != nil {
			return newBody, errors.WithStackTrace(err)
		}

		newBody = strings.Replace(newBody, linkPath, url, -1)
	}

	return newBody, nil
}

// Given a body of text find all instances of link paths (e.g. /foo or ../bar)
func getAllLinkPaths(body string) []string {
	var relPaths []string

	regex := regexp.MustCompile(FILE_PATHS_REGEX)
	submatches := regex.FindAllStringSubmatch(body, -1)

	if len(submatches) == 0 {
		return relPaths
	}

	for _, submatch := range submatches {
		relPath := submatch[0]

		// Cowardly use string search because Golang regular expressions don't support positive lookahead.
		if ! strings.Contains(relPath, "http://") && ! strings.Contains(relPath, "https://") {
			relPaths = append(relPaths, relPath)
		}
	}

	return relPaths
}

// Convert a link to another Package page to a fully qualified URL. For non-Package links, just return the original link
func convertPackageLinkToUrl(inputPath, linkPath string) (string, error) {
	var url string

	if isPackageInputPath(inputPath) {
		packageName, err := getPackageName(inputPath)
		if err != nil {
			return url, errors.WithStackTrace(err)
		}

		urlPrefix := strings.Replace(PACKAGE_GITHUB_REPO_URL_PREFIX, "<package-name>", packageName, 1)

		if strings.HasPrefix(linkPath, "/") {
			url = urlPrefix + linkPath
		}

		if strings.HasPrefix(linkPath, "./") {
			relPath, err := getPathRelativeToPackageRoot(inputPath)
			if err != nil {
				return url, errors.WithStackTrace(err)
			}

			url = urlPrefix + relPath
		}

		if strings.HasPrefix(linkPath, "../") {
			relPath, err := getPathRelativeToPackageRoot(inputPath)
			if err != nil {
				return url, errors.WithStackTrace(err)
			}

			relPath = filepath.Join(relPath, "../", linkPath)
			url = urlPrefix + relPath
		}

		return url, nil
	}

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
		return subpath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "PACKAGE_FILE_REGEX", regEx: PACKAGE_FILE_REGEX })
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
		return subpath, errors.WithStackTrace(&WrongNumberOfCaptureGroupsReturnedFromPageRegEx{ inputPath: inputPath, regExName: "PACKAGE_FILE_REGEX", regEx: PACKAGE_FILE_REGEX })
	}

	subpath = submatches[0][2]

	return subpath, nil
}
