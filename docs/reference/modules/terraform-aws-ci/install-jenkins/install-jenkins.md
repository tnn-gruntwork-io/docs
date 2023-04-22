---
title: "Install Jenkins Module"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="CI Modules" version="0.51.6" lastModifiedVersion="0.51.6"/>

# Install Jenkins Module

<<<<<<< Updated upstream
<a href="https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/releases/tag/v0.51.6" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-ci/releases/tag/v0.51.6" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>
>>>>>>> Stashed changes

This module contains two scripts for working with [Jenkins CI server](https://jenkins.io):

1.  `install.sh`: This script will install Jenkins on a Linux server. Currently, Ubuntu and CentOS are supported.
    This script also installs the `run-jenkins` script.
2.  `run-jenkins`: This script can be used to configure and run Jenkins. You will typically run this script while your
    server is booting.

## Example code

<<<<<<< Updated upstream
*   Check out the [jenkins example](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/examples/jenkins) for working sample code.
*   See [install.sh](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/install.sh) and [run-jenkins.sh](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/run-jenkins) for all options you can pass to these scripts.
=======
*   Check out the [jenkins example](https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/examples/jenkins) for working sample code.
*   See [install.sh](https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/install.sh) and [run-jenkins.sh](https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/run-jenkins) for all options you can pass to these scripts.
>>>>>>> Stashed changes

## Install Jenkins

The easiest way to install and run these scripts is to use the [Gruntwork
<<<<<<< Updated upstream
Installer](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/gruntwork-installer) (make sure to replace `VERSION` below with the latest
version from the [releases page](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/releases)):
=======
Installer](https://github.com/tnn-gruntwork-io/gruntwork-installer) (make sure to replace `VERSION` below with the latest
version from the [releases page](https://github.com/tnn-gruntwork-io/terraform-aws-ci/releases)):
>>>>>>> Stashed changes

```bash
gruntwork-install \
  --module-name 'install-jenkins' \
<<<<<<< Updated upstream
  --repo 'https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci' \
=======
  --repo 'https://github.com/tnn-gruntwork-io/terraform-aws-ci' \
>>>>>>> Stashed changes
  --tag '<VERSION>' \
  --version 2.164.3
```

The command above will copy `install.sh` to your server, run it, install Jenkins 2.164.3, and copy the `run-jenkins`
script into `/usr/local/bin`. We recommend running this command in a [Packer template](https://www.packer.io/) so you
<<<<<<< Updated upstream
can create an AMI with Jenkins installed. Check out the [jenkins example](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/examples/jenkins) for an example of such a
=======
can create an AMI with Jenkins installed. Check out the [jenkins example](https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/examples/jenkins) for an example of such a
>>>>>>> Stashed changes
Packer template.

## Run Jenkins

Once you have an AMI with Jenkins installed, you need to deploy it on an EC2 Instance in AWS. The easiest way to do
<<<<<<< Updated upstream
this is with the [jenkins-server module](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/jenkins-server). When the EC2 Instance is booting, you should
=======
this is with the [jenkins-server module](https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/jenkins-server). When the EC2 Instance is booting, you should
>>>>>>> Stashed changes
typically do two things in [User Data](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/user-data.html):

1.  Mount an EBS volume for the Jenkins home directory. You want to use an EBS volume so that your Jenkins data is
    persisted even if the EC2 Instance is replaced (e.g., after a crash or upgrade). The `mount-ebs-volume` script in the
<<<<<<< Updated upstream
    [persistent-ebs-volume module](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-server/tree/main/modules/persistent-ebs-volume)
=======
    [persistent-ebs-volume module](https://github.com/tnn-gruntwork-io/terraform-aws-server/tree/main/modules/persistent-ebs-volume)
>>>>>>> Stashed changes
    makes it easy to attach and mount a volume.

2.  Execute the `run-jenkins` script to start Jenkins, set its home directory to the mount point for the EBS volume,
    and configure it to use a certain amount of memory:

    ```bash
    run-jenkins \
      --memory "1g" \
      --jenkins-home "/jenkins"
    ```

<<<<<<< Updated upstream
Check out the [jenkins example](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/examples/jenkins) for an example of such a User Data script.
=======
Check out the [jenkins example](https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/examples/jenkins) for an example of such a User Data script.
>>>>>>> Stashed changes


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/readme.md",
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/variables.tf",
    "https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-ci/tree/v0.51.6/modules/install-jenkins/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "0ca69718b0260e4c89c3af38f1e6002e"
}
##DOCS-SOURCER-END -->
