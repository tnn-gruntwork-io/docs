---
title: "Backup PKI Assets Module"
hide_title: true
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';
import VersionBadge from '../../../../../src/components/VersionBadge.tsx';
import { HclListItem, HclListItemDescription, HclListItemTypeDetails, HclListItemDefaultValue, HclGeneralListItem } from '../../../../../src/components/HclListItem.tsx';
import { ModuleUsage } from "../../../../../src/components/ModuleUsage";

<VersionBadge repoTitle="Open VPN Package Infrastructure Package" version="0.25.0" lastModifiedVersion="0.19.0"/>

# Backup PKI Assets Module

<<<<<<< Updated upstream
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/releases/tag/v0.19.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This module is used to backup the OpenVPN Public Key Infrastructure (PKI) to S3 on a server that has been installed using
the [install-openvpn](https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/install-openvpn) module.
=======
<a href="https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki" className="link-button" title="View the source code for this module in GitHub.">View Source</a>

<a href="https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/releases/tag/v0.19.0" className="link-button" title="Release notes for only versions which impacted this module.">Release Notes</a>

This module is used to backup the OpenVPN Public Key Infrastructure (PKI) to S3 on a server that has been installed using
the [install-openvpn](https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/install-openvpn) module.
>>>>>>> Stashed changes

The PKI is the set of certificates used to verify the server and users' identities for VPN authentication purposes. This
normally lives on the OpenVPN server in the `/etc/openvpn-ca` and `/etc/openvpn` directories. If we didn't back these files
up, we would have to reissue client certificates if the OpenVPN server ever needed to be rebuilt.


<!-- ##DOCS-SOURCER-START
{
  "originalSources": [
<<<<<<< Updated upstream
    "https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki/outputs.tf"
=======
    "https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki/readme.md",
    "https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki/variables.tf",
    "https://github.com/tnn-gruntwork-io/terraform-aws-openvpn/tree/v0.25.0/modules/backup-openvpn-pki/outputs.tf"
>>>>>>> Stashed changes
  ],
  "sourcePlugin": "module-catalog-api",
  "hash": "d75c505937c2efc8de67a71016b2ed40"
}
##DOCS-SOURCER-END -->
