---
title: "AWS IAM Access Analyzer"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="Security Modules" version="0.67.8" lastModifiedVersion="0.66.0"/>

# AWS IAM Access Analyzer

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.66.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.66.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This repo contains a Module for creating and enabling [IAM Access Analyzer](https://docs.aws.amazon.com/IAM/latest/UserGuide/access-analyzer-getting-started.html)

This module is not meant to be used directly. Instead, it’s used under the hood in the `account-baseline-root` & `account-baseline-security` modules. Please see those modules and their respective examples for specific configuration and usage.

## Features

*   Create an IAM Access Analyzer service for different regions in one module

*   Enable the IAM Access Analyzer service for a given AWS account

## Learn

Note

This repo is a part of [the Gruntwork Infrastructure as Code Library](https://gruntwork.io/infrastructure-as-code-library/), a collection of reusable, battle-tested, production ready infrastructure code. If you’ve never used the Infrastructure as Code Library before, make sure to read [How to use the Gruntwork Infrastructure as Code Library](https://gruntwork.io/guides/foundations/how-to-use-gruntwork-infrastructure-as-code-library/)!

### Core concepts

<<<<<<< Updated upstream
*   [What is the AWS IAM Access Analyzer?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#what-is-the-aws-iam-access-analyzer?)

*   [What resources does IAM Access Analyzer analyze?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#what-resources-does-iam-access-analyzer-analyze?)

*   [IAM Access Analyzer documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html)

*   [How to use a multi-region module](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/codegen/core-concepts.md#how-to-use-a-multi-region-module)

### Repo organization

*   [modules](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules): the main implementation code for this repo, broken down into multiple standalone, orthogonal submodules.

*   [examples](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples): This folder contains working examples of how to use the submodules.

*   [test](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/test): Automated tests for the modules and examples.
=======
*   [What is the AWS IAM Access Analyzer?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#what-is-the-aws-iam-access-analyzer?)

*   [What resources does IAM Access Analyzer analyze?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#what-resources-does-iam-access-analyzer-analyze?)

*   [IAM Access Analyzer documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/what-is-access-analyzer.html)

*   [How to use a multi-region module](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/codegen/core-concepts.md#how-to-use-a-multi-region-module)

### Repo organization

*   [modules](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules): the main implementation code for this repo, broken down into multiple standalone, orthogonal submodules.

*   [examples](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples): This folder contains working examples of how to use the submodules.

*   [test](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/test): Automated tests for the modules and examples.
>>>>>>> Stashed changes

## Deploy

### Non-production deployment (quick start for learning)

If you just want to try this out for experimenting and learning, check out the following resources:

<<<<<<< Updated upstream
*   [examples folder](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples): The `examples` folder contains sample code optimized for learning, experimenting, and testing (but not production usage).

## Manage

*   [Who can manage the analyzer?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#who-can-manage-the-analyzer?)

*   [What to do with the access analyzer findings?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#what-to-do-with-the-access-analyzer-findings?)
=======
*   [examples folder](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples): The `examples` folder contains sample code optimized for learning, experimenting, and testing (but not production usage).

## Manage

*   [Who can manage the analyzer?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#who-can-manage-the-analyzer?)

*   [What to do with the access analyzer findings?](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/core-concepts.md#what-to-do-with-the-access-analyzer-findings?)
>>>>>>> Stashed changes

## Sample Usage

<Tabs>
<TabItem value="terraform" label="Terraform" default>

```hcl title="main.tf"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S IAM-ACCESS-ANALYZER-MULTI-REGION MODULE
# ------------------------------------------------------------------------------------------------------

module "iam_access_analyzer_multi_region" {

<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-security.git//modules/iam-access-analyzer-multi-region?ref=v0.67.8"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-security.git//modules/iam-access-analyzer-multi-region?ref=v0.67.8"
>>>>>>> Stashed changes

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The AWS Account ID the template should be operated on. This avoids
  # misconfiguration errors caused by environment variables.
  aws_account_id = <string>

  # ----------------------------------------------------------------------------------------------------
  # OPTIONAL VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # A feature flag to enable or disable this module.
  create_resources = true

  # The name of the IAM Access Analyzer module
  iam_access_analyzer_name = "iam-access-analyzer"

  # If set to ACCOUNT, the analyzer will only be scanning the current AWS account
  # it's in. If set to ORGANIZATION - will scan the organization AWS account and the
  # child accounts.
  iam_access_analyzer_type = "ACCOUNT"

}


```

</TabItem>
<TabItem value="terragrunt" label="Terragrunt" default>

```hcl title="terragrunt.hcl"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S IAM-ACCESS-ANALYZER-MULTI-REGION MODULE
# ------------------------------------------------------------------------------------------------------

terraform {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-security.git//modules/iam-access-analyzer-multi-region?ref=v0.67.8"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-security.git//modules/iam-access-analyzer-multi-region?ref=v0.67.8"
>>>>>>> Stashed changes
}

inputs = {

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The AWS Account ID the template should be operated on. This avoids
  # misconfiguration errors caused by environment variables.
  aws_account_id = <string>

  # ----------------------------------------------------------------------------------------------------
  # OPTIONAL VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # A feature flag to enable or disable this module.
  create_resources = true

  # The name of the IAM Access Analyzer module
  iam_access_analyzer_name = "iam-access-analyzer"

  # If set to ACCOUNT, the analyzer will only be scanning the current AWS account
  # it's in. If set to ORGANIZATION - will scan the organization AWS account and the
  # child accounts.
  iam_access_analyzer_type = "ACCOUNT"

}


```

</TabItem>
</Tabs>


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/readme.adoc",
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/readme.adoc",
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/iam-access-analyzer-multi-region/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "3bd36029185ba7a3b323dd3bdb43e503"
}
##DOCS-SOURCER-END -->
