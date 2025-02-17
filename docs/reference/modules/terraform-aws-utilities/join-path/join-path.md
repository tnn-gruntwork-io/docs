---
title: "Join Path Module"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="Terraform Utility Modules" version="0.9.1" lastModifiedVersion="0.7.0"/>

# Join Path Module

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.7.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.7.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This is a module that can be used to join a list of given path parts (that is, file and folder names) into a single
path with the appropriate path separator (backslash or forward slash) for the current operating system. This is useful
for ensuring the paths you build will work properly on Windows, Linux, and OS X.

This module uses Python under the hood so, the Python must be installed on the OS.

## Example code

<<<<<<< Updated upstream
See the [join-path example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/join-path) for working sample code.
=======
See the [join-path example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/join-path) for working sample code.
>>>>>>> Stashed changes

## Usage

Simply use the module in your Terraform code, replacing `<VERSION>` with the latest version from the [releases
<<<<<<< Updated upstream
page](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases), and specifying the path parts using the
=======
page](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases), and specifying the path parts using the
>>>>>>> Stashed changes
`path_parts` input:

```hcl
module "path" {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/join-path?ref=<VERSION>"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/join-path?ref=<VERSION>"
>>>>>>> Stashed changes
  
  path_parts = ["foo", "bar", "baz.txt"]
}
```

You can now get the joined path using the `path` output:

```hcl
# Will be set to "foo/bar/baz.txt" on Linux and OS X, "foo\bar\baz.txt" on Windows
joined_path = "${module.path.path}" 
```

## Sample Usage

<Tabs>
<TabItem value="terraform" label="Terraform" default>

```hcl title="main.tf"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S JOIN-PATH MODULE
# ------------------------------------------------------------------------------------------------------

module "join_path" {

<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/join-path?ref=v0.9.1"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/join-path?ref=v0.9.1"
>>>>>>> Stashed changes

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # A list of folder and file names to combine into a path, using the proper path
  # separator for the current OS.
  path_parts = <list(string)>

}


```

</TabItem>
<TabItem value="terragrunt" label="Terragrunt" default>

```hcl title="terragrunt.hcl"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S JOIN-PATH MODULE
# ------------------------------------------------------------------------------------------------------

terraform {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/join-path?ref=v0.9.1"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/join-path?ref=v0.9.1"
>>>>>>> Stashed changes
}

inputs = {

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # A list of folder and file names to combine into a path, using the proper path
  # separator for the current OS.
  path_parts = <list(string)>

}


```

</TabItem>
</Tabs>




## Reference

<Tabs>
<TabItem value="inputs" label="Inputs" default>

### Required

<HclListItem name="path_parts" requirement="required" type="list(string)">
<HclListItemDescription>

A list of folder and file names to combine into a path, using the proper path separator for the current OS.

</HclListItemDescription>
<HclGeneralListItem title="Examples">
<details>
  <summary>Example</summary>


```hcl
   path_parts = ["foo", "bar", "baz.txt"] => outputs "foo/bar/baz.txt" on Linux

```
</details>

</HclGeneralListItem>
</HclListItem>

</TabItem>
<TabItem value="outputs" label="Outputs">

<HclListItem name="path">
</HclListItem>

</TabItem>
</Tabs>


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "e846c907e790eec949efd70a96e15d59"
}
##DOCS-SOURCER-END -->
