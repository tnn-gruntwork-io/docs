---
title: "Security Modules"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="Security Modules" version="0.67.8" lastModifiedVersion="0.65.9"/>

# Security Modules

<<<<<<< Updated upstream
<a href="https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.65.9" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.65.9" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This module can configure a Linux server to automatically install critical security updates.

## Features

*   Automatically install critical security updates once per night so your servers don’t go unpatched for long periods of time.

*   Built-in random delay (between 0 and 30 minutes, by default) so all of your servers don’t update at the exact same time.

*   Supports Ubuntu 18.04 and 20.04 via [unattended-upgrades](https://help.ubuntu.com/lts/serverguide/automatic-updates.html).

*   Supports Amazon Linux, Amazon Linux 2, and CentOS via [yum-cron](http://man7.org/linux/man-pages/man8/yum-cron.8.html).

## Learn

Note

This repo is a part of [the Gruntwork Infrastructure as Code Library](https://gruntwork.io/infrastructure-as-code-library/), a collection of reusable, battle-tested, production ready infrastructure code. If you’ve never used the Infrastructure as Code Library before, make sure to read [How to use the Gruntwork Infrastructure as Code Library](https://gruntwork.io/guides/foundations/how-to-use-gruntwork-infrastructure-as-code-library/)!

### Core concepts

<<<<<<< Updated upstream
*   [How to install Auto Update](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#installation)

*   [How Auto Update works on Ubuntu](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#ubuntu-support)

*   [How Auto Update works on Amazon Linux and CentOS](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#amazon-linux-and-centos-support)

*   [Auto Update Limitations](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#limitations)

*   [Core Security Concepts](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/README.adoc#core-concepts)

### Repo organization

*   [modules](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules): the main implementation code for this repo, broken down into multiple standalone, orthogonal submodules.

*   [examples](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples): This folder contains working examples of how to use the submodules.

*   [test](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/test): Automated tests for the modules and examples.
=======
*   [How to install Auto Update](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#installation)

*   [How Auto Update works on Ubuntu](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#ubuntu-support)

*   [How Auto Update works on Amazon Linux and CentOS](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#amazon-linux-and-centos-support)

*   [Auto Update Limitations](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/core-concepts.md#limitations)

*   [Core Security Concepts](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/README.adoc#core-concepts)

### Repo organization

*   [modules](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules): the main implementation code for this repo, broken down into multiple standalone, orthogonal submodules.

*   [examples](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples): This folder contains working examples of how to use the submodules.

*   [test](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/test): Automated tests for the modules and examples.
>>>>>>> Stashed changes

## Deploy

### Non-production deployment (quick start for learning)

If you just want to try this repo out for experimenting and learning, check out the following resources:

<<<<<<< Updated upstream
*   [auto-update example](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples/auto-update): The `examples/auto-update` folder contains sample code optimized for learning, experimenting, and testing (but not production usage).
=======
*   [auto-update example](https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/examples/auto-update): The `examples/auto-update` folder contains sample code optimized for learning, experimenting, and testing (but not production usage).
>>>>>>> Stashed changes

### Production deployment

If you want to deploy this repo in production, check out the following resources:

<<<<<<< Updated upstream
*   [Packer template in the Acme example Reference Architecture](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/infrastructure-modules-multi-account-acme/blob/main/services/eks-cluster/packer/eks-node.json): Production-ready sample code from the Acme Reference Architecture examples.
=======
*   [Packer template in the Acme example Reference Architecture](https://github.com/tnn-gruntwork-io/infrastructure-modules-multi-account-acme/blob/main/services/eks-cluster/packer/eks-node.json): Production-ready sample code from the Acme Reference Architecture examples.
>>>>>>> Stashed changes


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/readme.adoc",
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/variables.tf",
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/readme.adoc",
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-security/tree/v0.67.8/modules/auto-update/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "3e07024fb5b82c7f04764f6890000bdb"
}
##DOCS-SOURCER-END -->
