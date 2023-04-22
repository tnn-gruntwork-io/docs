---
title: "AWS Helpers"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="CI Modules" version="0.51.6" lastModifiedVersion="0.50.11"/>

# AWS Helpers

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/releases/tag/v0.50.11" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/releases/tag/v0.50.11" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This module contains helper scripts that automate common AWS tasks:

*   `publish-ami`: This script copies the given AMI to the specified AWS regions and makes it public.

## Installing the helpers

<<<<<<< Updated upstream
You can install the helpers using the [Gruntwork Installer](https://github.com/tnn-gruntwork-io/gruntwork-installer):

```bash
gruntwork-install --module-name "aws-helpers" --repo "https://github.com/tnn-gruntwork-io/terraform-aws-ci" --tag "v0.0.1"
=======
You can install the helpers using the [Gruntwork Installer](https://github.com/tnn-gruntwork-io/gruntwork-installer):

```bash
gruntwork-install --module-name "aws-helpers" --repo "https://github.com/tnn-gruntwork-io/terraform-aws-ci" --tag "v0.0.1"
>>>>>>> Stashed changes
```

We recommend running this command in the `dependencies` section of `circle.yml`:

```yaml
dependencies:
  override:
    # Install the Gruntwork Installer
<<<<<<< Updated upstream
    - curl -Ls https://raw.githubusercontent.com/tnn-gruntwork-io/gruntwork-installer/main/bootstrap-gruntwork-installer.sh | bash /dev/stdin --version v0.0.16

    # Use the Gruntwork Installer to install the gruntwork-module-circleci-helpers module
    - gruntwork-install --module-name "aws-helpers" --repo "https://github.com/tnn-gruntwork-io/terraform-aws-ci" --tag "v0.0.1"
=======
    - curl -Ls https://raw.githubusercontent.com/tnn-gruntwork-io/gruntwork-installer/main/bootstrap-gruntwork-installer.sh | bash /dev/stdin --version v0.0.16

    # Use the Gruntwork Installer to install the gruntwork-module-circleci-helpers module
    - gruntwork-install --module-name "aws-helpers" --repo "https://github.com/tnn-gruntwork-io/terraform-aws-ci" --tag "v0.0.1"
>>>>>>> Stashed changes
```


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/aws-helpers/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "d1dc6cc1862d2fc43d260b51f240791c"
}
##DOCS-SOURCER-END -->
