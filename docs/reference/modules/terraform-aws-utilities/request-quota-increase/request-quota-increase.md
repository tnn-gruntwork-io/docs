---
title: "Request AWS Quota Increase"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="Terraform Utility Modules" version="0.9.1" lastModifiedVersion="0.9.1"/>

# Request AWS Quota Increase

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.9.1" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases/tag/v0.9.1" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This module can be used to request a quota increase for an AWS Resource.

## Features

*   Request a quota increase for Network ACL Rules and NAT Gateway.

## Learn

### Core Concepts

*   [AWS Service Quotas Documentation](https://docs.aws.amazon.com/servicequotas/?id=docs_gateway)

### Example code

<<<<<<< Updated upstream
See the [request-quota-increase example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/request-quota-increase) for working sample code.
=======
See the [request-quota-increase example](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/examples/request-quota-increase) for working sample code.
>>>>>>> Stashed changes

## Usage

Use the module in your Terraform code, replacing `<VERSION>` with the latest version from the [releases
<<<<<<< Updated upstream
page](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases):

```hcl
module "path" {
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/quota-increase?ref=<VERSION>"
=======
page](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/releases):

```hcl
module "path" {
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/quota-increase?ref=<VERSION>"
>>>>>>> Stashed changes

    request_quota_increase = {
      nat_gateway = 40,
      nacl_rules = 25
    }
}
```

The argument to pass is:

<<<<<<< Updated upstream
*   `request_quota_increase`: A map with the desired resource and the new quota. The current supported resources are `nat_gateway` and `nacl_rules`. Feel free to contribute to this module to add support for more `quota_code` and `service_code` options in [main.tf](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/main.tf)!
=======
*   `request_quota_increase`: A map with the desired resource and the new quota. The current supported resources are `nat_gateway` and `nacl_rules`. Feel free to contribute to this module to add support for more `quota_code` and `service_code` options in [main.tf](https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/main.tf)!
>>>>>>> Stashed changes

When you run `apply`, the `new_quotas` output variable will confirm to you that a quota request has been made!

```hcl
new_quotas = {
  "nat_gateway" = {
    "adjustable" = true
    "arn" = "arn:aws:servicequotas:us-east-1:<account-id>:vpc/L-FE5A380F"
    "default_value" = 5
    "id" = "vpc/L-FE5A380F"
    "quota_code" = "L-FE5A380F"
    "quota_name" = "NAT gateways per Availability Zone"
    "request_id" = "<request_id>"
    "request_status" = "PENDING"
    "service_code" = "vpc"
    "service_name" = "Amazon Virtual Private Cloud (Amazon VPC)"
    "value" = 30
  }
}
```

## Manage

You can see a full history of quota request changes using the [AWS
Console](https://console.aws.amazon.com/servicequotas/home#!/requests) or the AWS CLI:

```
aws service-quotas list-requested-service-quota-change-history --region <REGION>
```

### Finding out the Service Code and Quota Code

When you need to add a new resource, you can check the available services with

```
aws service-quotas list-services --region <REGION> --output table
```

And use the `ServiceCode` from the output to get the code for the resources

```
aws service-quotas list-service-quotas --service-code <SERVICE_CODE>
```

### Request a new quota smaller than the current one

If the new value that you request is smaller than the current one, *nothing* will happen. The
`terraform apply` output will contain the current quota. For example, if the NAT Gateway current
quota is 30 and you ask for a new quota of 25, this is the output:

```hcl
new_quotas = {
  "nat_gateway" = {
    "adjustable" = true
      "arn" = "arn:aws:servicequotas:us-east-1:<account-id>:vpc/L-FE5A380F"
      "default_value" = 5
      "id" = "vpc/L-FE5A380F"
      "quota_code" = "L-FE5A380F"
      "quota_name" = "NAT gateways per Availability Zone"
      "service_code" = "vpc"
      "service_name" = "Amazon Virtual Private Cloud (Amazon VPC)"
      "value" = 30   <------ Returned the current quota, not the requested one.
  }
}
```

### What happens when you run `destroy`

When you run `terraform destroy` on this module, it does not affect your current quotas or your
existing quota requests. In other words, you don't have to worry about quotas being reset to old
values; once they have been increased, they stay that way!

## Sample Usage

<Tabs>
<TabItem value="terraform" label="Terraform" default>

```hcl title="main.tf"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S REQUEST-QUOTA-INCREASE MODULE
# ------------------------------------------------------------------------------------------------------

module "request_quota_increase" {

<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/request-quota-increase?ref=v0.9.1"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/request-quota-increase?ref=v0.9.1"
>>>>>>> Stashed changes

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # A map where the key is the resource and the value is the desired quota. The only
  # resources supported at the moment are 'nacl_rules' and 'nat_gateway'. You can
  # also use the `aws_servicequotas_service_quota` resource directly, there are
  # instructions on how to find the Service Code and Quota Code on the README!
  resources_to_increase = <map(number)>

}


```

</TabItem>
<TabItem value="terragrunt" label="Terragrunt" default>

```hcl title="terragrunt.hcl"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S REQUEST-QUOTA-INCREASE MODULE
# ------------------------------------------------------------------------------------------------------

terraform {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/request-quota-increase?ref=v0.9.1"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-utilities.git//modules/request-quota-increase?ref=v0.9.1"
>>>>>>> Stashed changes
}

inputs = {

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # A map where the key is the resource and the value is the desired quota. The only
  # resources supported at the moment are 'nacl_rules' and 'nat_gateway'. You can
  # also use the `aws_servicequotas_service_quota` resource directly, there are
  # instructions on how to find the Service Code and Quota Code on the README!
  resources_to_increase = <map(number)>

}


```

</TabItem>
</Tabs>




## Reference

<Tabs>
<TabItem value="inputs" label="Inputs" default>

### Required

<HclListItem name="resources_to_increase" requirement="required" type="map(number)">
<HclListItemDescription>

A map where the key is the resource and the value is the desired quota. The only resources supported at the moment are 'nacl_rules' and 'nat_gateway'. You can also use the `aws_servicequotas_service_quota` resource directly, there are instructions on how to find the Service Code and Quota Code on the README!

</HclListItemDescription>
<HclGeneralListItem title="Examples">
<details>
  <summary>Example</summary>


```hcl
   {
     nacl_rules  = 39,
     nat_gateway = 20,
   }

```
</details>

</HclGeneralListItem>
</HclListItem>

</TabItem>
<TabItem value="outputs" label="Outputs">

<HclListItem name="new_quotas">
</HclListItem>

</TabItem>
</Tabs>


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-utilities/tree/v0.9.1/modules/request-quota-increase/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "7704e4ed57a0999f05e322066671ccb9"
}
##DOCS-SOURCER-END -->
