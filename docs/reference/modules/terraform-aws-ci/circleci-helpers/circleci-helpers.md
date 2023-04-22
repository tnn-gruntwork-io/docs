---
title: "CircleCI Helpers"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="CI Modules" version="0.51.6" lastModifiedVersion="0.51.4"/>

# CircleCI Helpers

<<<<<<< Updated upstream
<a href="https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/releases/tag/v0.51.4" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/releases/tag/v0.51.4" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This module contains helper scripts specially for CircleCI jobs, including:

*   `place-repo-in-gopath`: Places the git repo that triggered the CircleCI build in the GOPATH. This is useful because
    some tools (e.g. vendoring tools) require that they be executed in a repo located within the GOPATH.
*   `install-go-version`: Installs the given version of Golang.

## Installing the helpers

<<<<<<< Updated upstream
You can install the helpers using the [Gruntwork Installer](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/gruntwork-installer):

```bash
gruntwork-install --module-name "circleci-helpers" --repo "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci" --tag "0.0.1"
=======
You can install the helpers using the [Gruntwork Installer](https://github.com/tnn-gruntwork-io/gruntwork-installer):

```bash
gruntwork-install --module-name "circleci-helpers" --repo "https://github.com/tnn-gruntwork-io/terraform-aws-ci" --tag "0.0.1"
>>>>>>> Stashed changes
```

We recommend running this command in the `dependencies` section of `circle.yml`:

```yaml
dependencies:
  override:
    # Install the Gruntwork Installer
<<<<<<< Updated upstream
    - curl -Ls https://raw.githubusercontent.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/gruntwork-installer/main/bootstrap-gruntwork-installer.sh | bash /dev/stdin --version 0.0.9

    # Use the Gruntwork Installer to install the gruntwork-module-circleci-helpers module
    - gruntwork-install --module-name "circleci-helpers" --repo "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci" --tag "0.0.5"
=======
    - curl -Ls https://raw.githubusercontent.com/tnn-gruntwork-io/gruntwork-installer/main/bootstrap-gruntwork-installer.sh | bash /dev/stdin --version 0.0.9

    # Use the Gruntwork Installer to install the gruntwork-module-circleci-helpers module
    - gruntwork-install --module-name "circleci-helpers" --repo "https://github.com/tnn-gruntwork-io/terraform-aws-ci" --tag "0.0.5"
>>>>>>> Stashed changes
```

## Using the place-repo-in-gopath script

The `place-repo-in-gopath` script will create symlinks so GOPATH works correctly, as explained in [this
post](https://robots.thoughtbot.com/configure-circleci-for-go).

We recommend running this helper in the `dependencies` section of `circle.yml`:

```yaml
dependencies:
  override:
    # Install the Gruntwork Installer
<<<<<<< Updated upstream
    - curl -Ls https://raw.githubusercontent.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/gruntwork-installer/main/bootstrap-gruntwork-installer.sh | bash /dev/stdin --version 0.0.9

    # Use the Gruntwork Installer to install the gruntwork-module-circleci-helpers module
    - gruntwork-install --module-name "circleci-helpers" --repo "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci" --tag "0.0.1"
=======
    - curl -Ls https://raw.githubusercontent.com/tnn-gruntwork-io/gruntwork-installer/main/bootstrap-gruntwork-installer.sh | bash /dev/stdin --version 0.0.9

    # Use the Gruntwork Installer to install the gruntwork-module-circleci-helpers module
    - gruntwork-install --module-name "circleci-helpers" --repo "https://github.com/tnn-gruntwork-io/terraform-aws-ci" --tag "0.0.1"
>>>>>>> Stashed changes

    # Place the repo in the GOPATH
    - place-repo-in-gopath
```


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers/readme.md",
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers/variables.tf",
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/circleci-helpers/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "4b61bc7bd29a4af452aa88eeffdd71b3"
}
##DOCS-SOURCER-END -->
