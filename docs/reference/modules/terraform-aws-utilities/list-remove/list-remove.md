---
title: "List Remove Module"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="Terraform Utility Modules" version="0.9.1" lastModifiedVersion="0.7.0"/>

# List Remove Module

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.7.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.7.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This is a module that can be used to remove items in a given list from another list. This functionality is not yet
available as an interpolation function.

For example, suppose you have a list of availability zones (`["us-east-1a", "us-east-1b", "us-east-1c", "us-east-1d",
"us-east-1e"]`) and you want to remove specific zones that don't support the features you need (`["us-east-1b",
"us-east-1c"]`). You can use this module:

```hcl
module "list_remove" {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/list-remove?ref=v0.0.8"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/list-remove?ref=v0.0.8"
>>>>>>> Stashed changes

  original_list = ["us-east-1a", "us-east-1b", "us-east-1c", "us-east-1d", "us-east-1e"]
  items_to_remove = ["us-east-1b", "us-east-1c"]
}

output "output_list" {
  value = "${module.list_remove.output_list}"
}
```

The output `new_list` should be the list `["us-east-1a", "us-east-1d", "us-east-1e"]`.

**NOTE**: This will dedup the input list due to the way it is implemented. This module will not work if you are expecting duplicate items to remain.

## Example code

<<<<<<< Updated upstream
See the [list-remove example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/list-remove) for working sample code.
=======
See the [list-remove example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/list-remove) for working sample code.
>>>>>>> Stashed changes

## Sample Usage

<Tabs>
<TabItem value="terraform" label="Terraform" default>

```hcl title="main.tf"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S LIST-REMOVE MODULE
# ------------------------------------------------------------------------------------------------------

module "list_remove" {

<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/list-remove?ref=v0.9.1"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/list-remove?ref=v0.9.1"
>>>>>>> Stashed changes

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The list of items that you want to remove from the original list.
  items_to_remove = <list(any)>

  # The list of items where you want to remove items from.
  original_list = <list(any)>

}


```

</TabItem>
<TabItem value="terragrunt" label="Terragrunt" default>

```hcl title="terragrunt.hcl"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S LIST-REMOVE MODULE
# ------------------------------------------------------------------------------------------------------

terraform {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/list-remove?ref=v0.9.1"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/list-remove?ref=v0.9.1"
>>>>>>> Stashed changes
}

inputs = {

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The list of items that you want to remove from the original list.
  items_to_remove = <list(any)>

  # The list of items where you want to remove items from.
  original_list = <list(any)>

}


```

</TabItem>
</Tabs>




## Reference

<Tabs>
<TabItem value="inputs" label="Inputs" default>

### Required

<HclListItem name="items_to_remove" requirement="required" type="list(any)">
<HclListItemDescription>

The list of items that you want to remove from the original list.

</HclListItemDescription>
<HclListItemTypeDetails>

```hcl
Any types represent complex values of variable type. For details, please consult `variables.tf` in the source repo.
```

</HclListItemTypeDetails>
</HclListItem>

<HclListItem name="original_list" requirement="required" type="list(any)">
<HclListItemDescription>

The list of items where you want to remove items from.

</HclListItemDescription>
<HclListItemTypeDetails>

```hcl
Any types represent complex values of variable type. For details, please consult `variables.tf` in the source repo.
```

</HclListItemTypeDetails>
</HclListItem>

</TabItem>
<TabItem value="outputs" label="Outputs">

<HclListItem name="output_list">
</HclListItem>

</TabItem>
</Tabs>


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/list-remove/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "93f7e48ad3cc474cd8842133677b71e3"
}
##DOCS-SOURCER-END -->
