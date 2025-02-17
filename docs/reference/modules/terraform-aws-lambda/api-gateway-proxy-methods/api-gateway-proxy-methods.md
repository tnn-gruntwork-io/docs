---
title: "API Gateway Lambda Function Proxy Methods Module"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="AWS Lambda" version="0.21.8" lastModifiedVersion="0.20.0"/>

# API Gateway Lambda Function Proxy Methods Module

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/releases/tag/v0.20.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This module must be used in conjunction with [the api-gateway-proxy module](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy) to configure an API
Gateway REST API to route all requests from a root path to a lambda function.

Refer to [the module docs](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy/README.md) for the `api-gateway-proxy` module for more details on how to
use this module. Specifically, see the section [Can I expose additional lambda functions in a decentralized
manner?](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy/core-concepts.md#can-i-expose-additional-lambda-functions-in-a-decentralized-manner)
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-lambda/releases/tag/v0.20.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This module must be used in conjunction with [the api-gateway-proxy module](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy) to configure an API
Gateway REST API to route all requests from a root path to a lambda function.

Refer to [the module docs](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy/README.md) for the `api-gateway-proxy` module for more details on how to
use this module. Specifically, see the section [Can I expose additional lambda functions in a decentralized
manner?](https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy/core-concepts.md#can-i-expose-additional-lambda-functions-in-a-decentralized-manner)
>>>>>>> Stashed changes

## Sample Usage

<Tabs>
<TabItem value="terraform" label="Terraform" default>

```hcl title="main.tf"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S API-GATEWAY-PROXY-METHODS MODULE
# ------------------------------------------------------------------------------------------------------

module "api_gateway_proxy_methods" {

<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/api-gateway-proxy-methods?ref=v0.21.8"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/api-gateway-proxy-methods?ref=v0.21.8"
>>>>>>> Stashed changes

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The API Gateway REST API resource as returned by the terraform resource or data
  # source. This can also be able arbitrary object that has the keys id,
  # root_resource_id, and execution_arn of the API Gateway REST API.
  api_gateway_rest_api = <any>

  # Name of the AWS Lambda function that is being invoked for the API requests.
  lambda_function_name = <string>

  # ----------------------------------------------------------------------------------------------------
  # OPTIONAL VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The URL path prefix to proxy. Requests to any path under this path prefix will
  # be routed to the lambda function. Note that if the path prefix is empty string
  # (default), all requests (including to the root path) will be proxied. Note that
  # this only supports single levels for now (e.g., you can configure to route `foo`
  # and everything below that path like `foo/api/v1`, but you cannot configure to
  # route something like `api/foo/*`). Example: api will route all requests under
  # api/, such as /api, /api/v1, /api/v2/myresource/action, etc.
  path_prefix = ""

  # Configures only the root path to route to the lambda function, and not the other
  # subpaths. When true, the path_prefix must be empty string or no resources will
  # be created.
  root_only = false

}


```

</TabItem>
<TabItem value="terragrunt" label="Terragrunt" default>

```hcl title="terragrunt.hcl"

# ------------------------------------------------------------------------------------------------------
# DEPLOY GRUNTWORK'S API-GATEWAY-PROXY-METHODS MODULE
# ------------------------------------------------------------------------------------------------------

terraform {
<<<<<<< Updated upstream
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/api-gateway-proxy-methods?ref=v0.21.8"
=======
  source = "git::git@github.com:tnn-gruntwork-io/terraform-aws-lambda.git//modules/api-gateway-proxy-methods?ref=v0.21.8"
>>>>>>> Stashed changes
}

inputs = {

  # ----------------------------------------------------------------------------------------------------
  # REQUIRED VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The API Gateway REST API resource as returned by the terraform resource or data
  # source. This can also be able arbitrary object that has the keys id,
  # root_resource_id, and execution_arn of the API Gateway REST API.
  api_gateway_rest_api = <any>

  # Name of the AWS Lambda function that is being invoked for the API requests.
  lambda_function_name = <string>

  # ----------------------------------------------------------------------------------------------------
  # OPTIONAL VARIABLES
  # ----------------------------------------------------------------------------------------------------

  # The URL path prefix to proxy. Requests to any path under this path prefix will
  # be routed to the lambda function. Note that if the path prefix is empty string
  # (default), all requests (including to the root path) will be proxied. Note that
  # this only supports single levels for now (e.g., you can configure to route `foo`
  # and everything below that path like `foo/api/v1`, but you cannot configure to
  # route something like `api/foo/*`). Example: api will route all requests under
  # api/, such as /api, /api/v1, /api/v2/myresource/action, etc.
  path_prefix = ""

  # Configures only the root path to route to the lambda function, and not the other
  # subpaths. When true, the path_prefix must be empty string or no resources will
  # be created.
  root_only = false

}


```

</TabItem>
</Tabs>




## Reference

<Tabs>
<TabItem value="inputs" label="Inputs" default>

### Required

<HclListItem name="api_gateway_rest_api" requirement="required" type="any">
<HclListItemDescription>

The API Gateway REST API resource as returned by the terraform resource or data source. This can also be able arbitrary object that has the keys id, root_resource_id, and execution_arn of the API Gateway REST API.

</HclListItemDescription>
<HclListItemTypeDetails>

```hcl
Any types represent complex values of variable type. For details, please consult `variables.tf` in the source repo.
```

</HclListItemTypeDetails>
<HclGeneralListItem title="More Details">
<details>


```hcl

   Ideally we can define an object type, but that would require defining every attribute of the API Gateway REST API
   resource, which can be painful if you can't pass through the entire resource (e.g., as in terragrunt dependencies).

```
</details>

</HclGeneralListItem>
</HclListItem>

<HclListItem name="lambda_function_name" requirement="required" type="string">
<HclListItemDescription>

Name of the AWS Lambda function that is being invoked for the API requests.

</HclListItemDescription>
</HclListItem>

### Optional

<HclListItem name="path_prefix" requirement="optional" type="string">
<HclListItemDescription>

The URL path prefix to proxy. Requests to any path under this path prefix will be routed to the lambda function. Note that if the path prefix is empty string (default), all requests (including to the root path) will be proxied. Note that this only supports single levels for now (e.g., you can configure to route `foo` and everything below that path like `foo/api/v1`, but you cannot configure to route something like `api/foo/*`). Example: api will route all requests under api/, such as /api, /api/v1, /api/v2/myresource/action, etc.

</HclListItemDescription>
<HclListItemDefaultValue defaultValue="&quot;&quot;"/>
</HclListItem>

<HclListItem name="root_only" requirement="optional" type="bool">
<HclListItemDescription>

Configures only the root path to route to the lambda function, and not the other subpaths. When true, the path_prefix must be empty string or no resources will be created.

</HclListItemDescription>
<HclListItemDefaultValue defaultValue="false"/>
</HclListItem>

</TabItem>
<TabItem value="outputs" label="Outputs">

<HclListItem name="path_prefix_root_integration_id">
<HclListItemDescription>

ID of the API Gateway integration for the root method of the path_prefix resource.

</HclListItemDescription>
</HclListItem>

<HclListItem name="path_prefix_root_method_http_method">
<HclListItemDescription>

HTTP method of the API Gateway method for the root of the path_prefix resource.

</HclListItemDescription>
</HclListItem>

<HclListItem name="path_prefix_root_method_id">
<HclListItemDescription>

ID of the API Gateway method for the root of the path_prefix resource.

</HclListItemDescription>
</HclListItem>

<HclListItem name="path_prefix_root_resource_id">
<HclListItemDescription>

ID of the API Gateway method for the root of the path_prefix resource.

</HclListItemDescription>
</HclListItem>

<HclListItem name="proxy_integration_id">
<HclListItemDescription>

ID of the API Gateway integration for the proxy method.

</HclListItemDescription>
</HclListItem>

<HclListItem name="proxy_method_http_method">
<HclListItemDescription>

HTTP method of the API Gateway method for the proxy.

</HclListItemDescription>
</HclListItem>

<HclListItem name="proxy_method_id">
<HclListItemDescription>

ID of the API Gateway method for the proxy.

</HclListItemDescription>
</HclListItem>

<HclListItem name="proxy_resource_id">
<HclListItemDescription>

ID of the API Gateway method for the proxy resource.

</HclListItemDescription>
</HclListItem>

<HclListItem name="root_intergration_id">
<HclListItemDescription>

ID of the API Gateway integration for the root method (only created if path_prefix is empty string).

</HclListItemDescription>
</HclListItem>

<HclListItem name="root_method_http_method">
<HclListItemDescription>

HTTP method of the API Gateway method for the root proxy (only created if path_prefix is empty string).

</HclListItemDescription>
</HclListItem>

<HclListItem name="root_method_id">
<HclListItemDescription>

ID of the API Gateway method for the root proxy (only created if path_prefix is empty string).

</HclListItemDescription>
</HclListItem>

</TabItem>
</Tabs>


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-lambda/tree/v0.21.8/modules/api-gateway-proxy-methods/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "d9996a83db132e4f3a94dc42fdb3a0ee"
}
##DOCS-SOURCER-END -->
