# Network ACL Outbound Terraform Module

This Terraform Module launches is a simple helper for adding outbound rules to a [Network
ACL](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/network-acl-outbound/docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html). Network ACLs can be a bit tricky to work with
because they are stateless, which means that opening an outbound port is often not enough; you also need to open
[ephemeral inbound ports](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/network-acl-outbound/docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html#VPC_ACLs_Ephemeral_Ports)
which the remote services can use to respond. This can be very easy to forget, so this module adds not only the
outbound to an ACL, but also the ephemeral inbound ports for return traffic.

See the [network-acl-inbound](https://github.com/gruntwork-io/module-vpc/tree/master/modules/network-acl-inbound) module for the analogous version of this module, but for opening
inbound ports.

## How do you use this module?

Check out the [vpc-app-network-acls](https://github.com/gruntwork-io/module-vpc/tree/master/modules/network-acl-outbound/modules/vpc-app-network-acls) and
[vpc-mgmt-network-acls](https://github.com/gruntwork-io/module-vpc/tree/master/modules/network-acl-outbound/modules/vpc-mgmt-network-acls) modules for examples.

Check out [vars.tf](vars.tf) for all the configuration options available.

## What's a Network ACL?

[Network ACLs](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/network-acl-outbound/docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html) provide an extra layer of network
security, similar to a [security group](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/network-acl-outbound/docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html).
Whereas a security group controls what inbound and outbound traffic is allowed for a specific resource (e.g. a single
EC2 instance), a network ACL controls what inbound and outbound traffic is allowed for an entire subnet.

