---
title: "Scheduled Lambda Job Module"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="AWS Lambda" version="0.21.8" lastModifiedVersion="0.20.0"/>

# Scheduled Lambda Job Module

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/releases/tag/v0.20.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This module makes it easy to run an [AWS Lambda](https://aws.amazon.com/lambda/) function (such as one created with the
[lambda module](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/lambda)) on a scheduled basis. This is useful for periodic background jobs, such as taking a
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/releases/tag/v0.20.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This module makes it easy to run an [AWS Lambda](https://aws.amazon.com/lambda/) function (such as one created with the
[lambda module](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/lambda)) on a scheduled basis. This is useful for periodic background jobs, such as taking a
>>>>>>> Stashed changes
daily snapshot of your servers.

## Background info

For more information on AWS Lambda, how it works, and how to configure your functions, check out the [lambda module
<<<<<<< Updated upstream
documentation](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/lambda).
=======
documentation](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/lambda).
>>>>>>> Stashed changes

## Sample Usage

<Tabs>
<TabItem value="terraform" label="Terraform" default>

```hcl title="main.tf"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S SCHEDULED-LAMBDA-JOB MODULE
# ------------------------------------------------------------------------------------------------------

module "scheduled_lambda_job" {

<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/scheduled-lambda-job?ref=v0.21.8"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/scheduled-lambda-job?ref=v0.21.8"
>>>>>>> Stashed changes

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The ARN of the lambda function.
  lambda_function_arn = <string>

  # The name of the lambda function.
  lambda_function_name = <string>

  # An expression that defines the schedule for this lambda job. For example, cron(0
  # 20 * * ? *) or rate(5 minutes).
  schedule_expression = <string>

  # ----------------------------------------------------------------------------------------------------
  # OPTIONAL VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # Set to false to have this module skip creating resources. This weird parameter
  # exists solely because Terraform does not support conditional modules. Therefore,
  # this is a hack to allow you to conditionally decide if this module should create
  # anything or not.
  create_resources = true

  # JSON text that will be passed to the lambda function on each invoke.
  lambda_function_input = null

  # The namespace to use for all resources created by this module. If not set,
  # var.lambda_function_name, with '-scheduled' as a suffix, is used.
  namespace = null

  # A map of tags to apply to the event rule.
  tags = {}

}


```

</TabItem>
<TabItem value="terragrunt" label="Terragrunt" default>

```hcl title="terragrunt.hcl"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S SCHEDULED-LAMBDA-JOB MODULE
# ------------------------------------------------------------------------------------------------------

terraform {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/scheduled-lambda-job?ref=v0.21.8"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/scheduled-lambda-job?ref=v0.21.8"
>>>>>>> Stashed changes
}

inputs = {

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The ARN of the lambda function.
  lambda_function_arn = <string>

  # The name of the lambda function.
  lambda_function_name = <string>

  # An expression that defines the schedule for this lambda job. For example, cron(0
  # 20 * * ? *) or rate(5 minutes).
  schedule_expression = <string>

  # ----------------------------------------------------------------------------------------------------
  # OPTIONAL VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # Set to false to have this module skip creating resources. This weird parameter
  # exists solely because Terraform does not support conditional modules. Therefore,
  # this is a hack to allow you to conditionally decide if this module should create
  # anything or not.
  create_resources = true

  # JSON text that will be passed to the lambda function on each invoke.
  lambda_function_input = null

  # The namespace to use for all resources created by this module. If not set,
  # var.lambda_function_name, with '-scheduled' as a suffix, is used.
  namespace = null

  # A map of tags to apply to the event rule.
  tags = {}

}


```

</TabItem>
</Tabs>




## Reference

<Tabs>
<TabItem value="inputs" label="Inputs" default>

### Required

<HclListItem name="lambda_function_arn" requirement="required" type="string">
<HclListItemDescription>

The ARN of the lambda function.

</HclListItemDescription>
</HclListItem>

<HclListItem name="lambda_function_name" requirement="required" type="string">
<HclListItemDescription>

The name of the lambda function.

</HclListItemDescription>
</HclListItem>

<HclListItem name="schedule_expression" requirement="required" type="string">
<HclListItemDescription>

An expression that defines the schedule for this lambda job. For example, cron(0 20 * * ? *) or rate(5 minutes).

</HclListItemDescription>
</HclListItem>

### Optional

<HclListItem name="create_resources" requirement="optional" type="bool">
<HclListItemDescription>

Set to false to have this module skip creating resources. This weird parameter exists solely because Terraform does not support conditional modules. Therefore, this is a hack to allow you to conditionally decide if this module should create anything or not.

</HclListItemDescription>
<HclListItemDefaultValue defaultValue="true"/>
</HclListItem>

<HclListItem name="lambda_function_input" requirement="optional" type="string">
<HclListItemDescription>

JSON text that will be passed to the lambda function on each invoke.

</HclListItemDescription>
<HclListItemDefaultValue defaultValue="null"/>
</HclListItem>

<HclListItem name="namespace" requirement="optional" type="string">
<HclListItemDescription>

The namespace to use for all resources created by this module. If not set, <a href="#lambda_function_name"><code>lambda_function_name</code></a>, with '-scheduled' as a suffix, is used.

</HclListItemDescription>
<HclListItemDefaultValue defaultValue="null"/>
</HclListItem>

<HclListItem name="tags" requirement="optional" type="map(string)">
<HclListItemDescription>

A map of tags to apply to the event rule.

</HclListItemDescription>
<HclListItemDefaultValue defaultValue="{}"/>
</HclListItem>

</TabItem>
<TabItem value="outputs" label="Outputs">

<HclListItem name="event_rule_arn">
<HclListItemDescription>

Cloudwatch Event Rule Arn

</HclListItemDescription>
</HclListItem>

<HclListItem name="event_rule_schedule">
<HclListItemDescription>

Cloudwatch Event Rule schedule expression

</HclListItemDescription>
</HclListItem>

</TabItem>
</Tabs>


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/scheduled-lambda-job/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "fcf3dd6cd225e21b7f442a59b1bdb9c1"
}
##DOCS-SOURCER-END -->
