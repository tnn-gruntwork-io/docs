# docs-preprocessor

The purpose of docs-preprocessor is to prepare a single folder to be used to convert markdown to HTML.

In particular, docs-preprocessor starts with a folder structure that looks like this:

```
- introduction
  - overview.md
  - setup.md
- packages
  - module-vpc
    - docs
      - core-concepts.md
    - examples
      - ...
      - vpc-app
        ─ ...
        ─ README.md
    ─ modules
      - ...
      - vpc-app
        - ...
        - README.md
    - README.md
```

and transforms it to this:

```
- introduction
  - overview.md
  - setup.md
- packages
  - module-vpc
    - modules
      - vpc-app
        - overview.md
        - examples.md
    - overview.md
    - core-concepts.md
```

The tool is written in Go (as opposed to bash) to support more exotic mapping logic in the future if needed.

...

## docs_preprocessor

- Convert all relative links to absolute links
  - e.g. /modules/vpc-app --> https://github.com/gruntwork-io/module-vpc/blob/master/modules/vpc-app
- Leave all http:// and https:// links unchanged
- Global links should be treated differently, but I'm not sure how yet?

## docs_generator

This function assumes that it starts with a "sanitized" file structure optimized for a public website layout. Its 
objective is to produce HTML of each markdown page, taking care to transform links where necessary, and also to copy 
any other files that belong in the documentation. In addition, it maintains a navigation tree in memory and outputs
 this in HTML on each page.
 
The docs-generator has two passes through the file system. The first creates the nav tree.

The second pass generates the markdown file along with an HTML representation of the nav-tree, and also updates any links
that need to be fixed.

### Nav Tree Pass

- Is it a folder that doesn't begin with `_`? Then it's a section label.
- Is it a package folder? Consider decorating it differently.
- Is it a .md file? Create a node
- Is it an image? Do nothing.

### Markdown Generation Pass

- Is it a folder that doesn't begin with `_`? Do nothing.
- Is it a .md file? Create the HTML in the same folder
- Is it an image? Copy it.


