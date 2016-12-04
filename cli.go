package main

import (
	"fmt"
	"strings"

	"github.com/gobwas/glob"
	"github.com/urfave/cli"
	"github.com/gruntwork-io/docs/errors"
	"github.com/gruntwork-io/docs/globs"
	"reflect"
	"os"
)

// Customize the --help text for the app so we don't show extraneous info
const CUSTOM_HELP_TEXT = `NAME:
   {{.Name}} - {{.Usage}}
   
USAGE:
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
   {{if .VisibleFlags}}
OPTIONS:
   {{range .VisibleFlags}}
   {{.}}
   {{end}}{{end}}{{if len .Authors}}
   {{if .Version}}{{if not .HideVersion}}
VERSION:
   {{.Version}}
   {{end}}{{end}}
AUTHOR(S):
   {{range .Authors}}{{.}}{{end}}
   {{end}}{{if .Copyright}}
COPYRIGHT:
   {{.Copyright}}
   {{end}}
`

const OPT_HTML_FILES_PATH = "html-files-path"
const OPT_CONFIG_FILE_PATH = "config-file-path"
const OPT_GLOBAL_DOCS_PATH = "global-docs-path"
const OPT_INPUT_PATH = "input-path"
const OPT_OUTPUT_PATH = "output-path"
const OPT_DOC_PATTERN = "doc-pattern"
const OPT_EXCLUDE_PATTERN = "exclude-pattern"

var DEFAULT_DOC_PATTERNS = []string{"*.md", "*.txt", "*.jpg", "*.png", "*.gif"}
var DEFAULT_EXCLUDES = []string{".git*", "vendor", "vendor/*", "test/vendor", "test/vendor/*"}

type Opts struct {
	HtmlFilesPath  string
	ConfigFilePath string
	GlobalDocsPath string
	InputPath      string
	OutputPath     string
	DocPatterns    []glob.Glob
	Excludes       []glob.Glob
}

type EnvVars struct {
	GithubOauthToken string `name:"GITHUB_OAUTH_TOKEN"`
}

const ENV_VARS_STRUCT_TAG_NAME = "name"

func CreateCli(version string) *cli.App {
	// Override the exiter to do nothing, since we want our own code to handle errors
	cli.OsExiter = func(code int) {}
	cli.AppHelpTemplate = CUSTOM_HELP_TEXT

	app := cli.NewApp()

	app.Name = "docs-preprocessor"
	app.Author = "Gruntwork <www.gruntwork.io>"
	app.Usage = `Transforms the existing folder structure of a collection of repos into one more suitable for a public documentation website.`
	app.UsageText = "docs-preprocessor [OPTIONS]"
	app.Version = version
	app.Action = runApp

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  OPT_HTML_FILES_PATH,
			Usage: "Use the HTML, CSS, and JS files in `PATH` to generate the final website.",
		},
		cli.StringFlag{
			Name:  OPT_CONFIG_FILE_PATH,
			Usage: "Use the JSON file at the given `PATH` to identify the Gruntwork Packages for which docs should be generated, plus other config.",
		},
		cli.StringFlag{
			Name:  OPT_GLOBAL_DOCS_PATH,
			Usage: "Copy the files at `PATH` to generate documentation files global to all Gruntwork Packages.",
		},
		cli.StringFlag{
			Name:  OPT_INPUT_PATH,
			Usage: "Generate documentation from the files and folders in `PATH`.",
		},
		cli.StringFlag{
			Name:  OPT_OUTPUT_PATH,
			Usage: "Write the output to `PATH`.",
		},
		cli.StringSliceFlag{
			Name: OPT_DOC_PATTERN,
			Usage: fmt.Sprintf(`Copy files that match the PATTERN to the output path, unchanged. 
	Supports standard file patterns (e.g. *.txt, foo/**/bar). Make sure 
	to quote PATTERN so bash doesn't expand it. May be specified more 
	than once. Default: %s`, strings.Join(DEFAULT_DOC_PATTERNS, " ")),
		},
		cli.StringSliceFlag{
			Name:  OPT_EXCLUDE_PATTERN,
			Usage: fmt.Sprintf(`A PATTERN to exclude while copying to the output path. Supports 
	standard file patterns (e.g. *.tf, foo/**/bar). Make sure to quote 
	PATTERN so bash doesn't expand it. May be specified more than once. 
	Default: %s`, strings.Join(DEFAULT_EXCLUDES, " ")),
		},
	}

	return app
}

// When you run the CLI, this is the action function that gets called
func runApp(cliContext *cli.Context) error {
	if !cliContext.Args().Present() && cliContext.NumFlags() == 0 {
		cli.ShowAppHelp(cliContext)
		return nil
	}

	opts, err := parseOpts(cliContext)
	if err != nil {
		return err
	}

	envVars, err := getEnvVars()
	if err != nil {
		return err
	}

	if err := GenerateDocs(opts, envVars); err != nil {
		return err
	}

	return nil
}

func parseOpts(cliContext *cli.Context) (*Opts, error) {
	htmlPath := cliContext.String(OPT_HTML_FILES_PATH)
	if htmlPath == "" {
		return nil, errors.WithStackTrace(MissingParam(OPT_HTML_FILES_PATH))
	}

	repoManifestPath := cliContext.String(OPT_CONFIG_FILE_PATH)
	if repoManifestPath == "" {
		return nil, errors.WithStackTrace(MissingParam(OPT_CONFIG_FILE_PATH))
	}

	globalDocsPath := cliContext.String(OPT_GLOBAL_DOCS_PATH)
	if globalDocsPath == "" {
		return nil, errors.WithStackTrace(MissingParam(OPT_GLOBAL_DOCS_PATH))
	}

	inputPath := cliContext.String(OPT_INPUT_PATH)
	if inputPath == "" {
		return nil, errors.WithStackTrace(MissingParam(OPT_INPUT_PATH))
	}

	outputPath := cliContext.String(OPT_OUTPUT_PATH)
	if outputPath == "" {
		return nil, errors.WithStackTrace(MissingParam(OPT_OUTPUT_PATH))
	}

	docPatterns := cliContext.StringSlice(OPT_DOC_PATTERN)
	if len(docPatterns) == 0 {
		docPatterns = DEFAULT_DOC_PATTERNS
	}
	docGlobs, err := globs.ToGlobs(docPatterns)
	if err != nil {
		return nil, err
	}

	excludePatterns := cliContext.StringSlice(OPT_EXCLUDE_PATTERN)
	if len(excludePatterns) == 0 {
		excludePatterns = DEFAULT_EXCLUDES
	}
	excludeGlobs, err := globs.ToGlobs(excludePatterns)
	if err != nil {
		return nil, err
	}

	return &Opts{
		HtmlFilesPath: htmlPath,
		ConfigFilePath: repoManifestPath,
		GlobalDocsPath: globalDocsPath,
		InputPath: inputPath,
		OutputPath: outputPath,
		DocPatterns: docGlobs,
		Excludes: excludeGlobs,
	}, nil
}

// obtain the env vars from the local environment
func getEnvVars() (*EnvVars, error) {
	githubOauthTokenEnvVarName, err := getOsEnvVarName("GithubOauthToken")
	if err != nil {
		return nil, errors.WithStackTrace(err)
	}

	githubOauthTokenEnvVarValue := os.Getenv(githubOauthTokenEnvVarName)
	if githubOauthTokenEnvVarValue == "" {
		return nil, errors.WithStackTrace(MissingEnvVar(githubOauthTokenEnvVarName))
	}

	return &EnvVars{
		GithubOauthToken: githubOauthTokenEnvVarValue,
	}, nil
}

// Given a field name of the EnvVars struct, return the corresponding OS Environemnt Variable name
func getOsEnvVarName(fieldName string) (string, error) {
	var osEnvVarName string

	envVars := &EnvVars{}
	field, ok := reflect.TypeOf(envVars).Elem().FieldByName(fieldName)
	if !ok {
		return osEnvVarName, fmt.Errorf("The struct field '%s' was not found in the EnvVars struct.\n", fieldName)
	}

	tag := field.Tag
	osEnvVarName = tag.Get(ENV_VARS_STRUCT_TAG_NAME)

	return osEnvVarName, nil
}

// Custom error types
type MissingParam string
func (paramName MissingParam) Error() string {
	return fmt.Sprintf("Required parameter --%s cannot be empty", string(paramName))
}

type MissingEnvVar string
func (envVarName MissingEnvVar) Error() string {
	return fmt.Sprintf("Required environment variable %s cannot be empty", string(envVarName))
}
