---
title: "Operating System Module"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="Terraform Utility Modules" version="0.9.1" lastModifiedVersion="0.8.0"/>

# Operating System Module

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.8.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This is a module that can be used to figure out what operating system is being used to run Terraform. This may be used
to modify Terraform's behavior depending on the OS, such as modifying the way you format file paths on Linux vs
Windows (see also the [join-path module](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path)).
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.8.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This is a module that can be used to figure out what operating system is being used to run Terraform. This may be used
to modify Terraform's behavior depending on the OS, such as modifying the way you format file paths on Linux vs
Windows (see also the [join-path module](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/join-path)).
>>>>>>> Stashed changes

This module uses Python under the hood so, the Python must be installed on the OS.

## Example code

<<<<<<< Updated upstream
See the [operating-system example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/operating-system) for working sample code.
=======
See the [operating-system example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/operating-system) for working sample code.
>>>>>>> Stashed changes

## Usage

Simply use the module in your Terraform code, replacing `<VERSION>` with the latest version from the [releases
<<<<<<< Updated upstream
page](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases):

```hcl
module "os" {
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/operating-system?ref=<VERSION>"
=======
page](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases):

```hcl
module "os" {
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/operating-system?ref=<VERSION>"
>>>>>>> Stashed changes
}
```

*   You can now get the name of the operating system from the `name` output, which will be set to either `Linux`,
    `Darwin`, or `Windows`

*   You can also get the path separator for the current OS—backslash for Windows, forward slash everywhere else—from the
    `path_separator` output.

```hcl
operating_system_name = "${module.os.name}"
path_separator        = "${module.os.path_separator}"
```


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/operating-system/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "20974b3df2df0de3966b9f2272bb7064"
}
##DOCS-SOURCER-END -->
