# VPC-App Network ACLs Terraform Module

This Terraform Module adds a default set of [Network
ACLs](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html) to a VPC created using the
[vpc-app](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app) module. The ACLs enforce the following security settings (based on [A Reference VPC
Architecture](https:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/www.whaletech.co/2014/10/02/reference-vpc-architecture.html)):

- **Public subnet**: Allow all requests.
- **Private app subnet**: Allow all requests https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/to/from the public subnets, private persistence subnets, and the Mgmt VPC.
  Allow all outbound TCP requests plus return traffic from any IP for those TCP requests on [ephemeral
  ports](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html#VPC_ACLs_Ephemeral_Ports).
- **Private persistence subnet**: Allow all requests https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/to/from the private app subnets and the Mgmt VPC.

## How do you use this module?

Check out the [vpc-network-acls example](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/examples/vpc-network-acls).

Check out [vars.tf](vars.tf) for all the configuration options available.

## What's a VPC?

A [VPC](https:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/aws.amazon.com/vpc) or Virtual Private Cloud is a logically isolated section of your AWS cloud. Each
VPC defines a virtual network within which you run your AWS resources, as well as rules for what can go in and out of
that network. This includes subnets, route tables that tell those subnets how to route inbound and outbound traffic,
security groups, firewalls for the subnet (known as "Network ACLs"), and any other network components such as VPN connections.

## What's a Network ACL?

[Network ACLs](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_ACLs.html) provide an extra layer of network
security, similar to a [security group](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app-network-acls/docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html).
Whereas a security group controls what inbound and outbound traffic is allowed for a specific resource (e.g. a single
EC2 instance), a network ACL controls what inbound and outbound traffic is allowed for an entire subnet.

