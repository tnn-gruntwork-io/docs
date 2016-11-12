# VPC-Mgmt Terraform Module

This Terraform Module launches a single VPC meant to house DevOps and other management services. By contrast, the apps
that power your business should run in an "app" VPC. (See the [vpc-app](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-app) module.)

## How do you use this module?

Check out the [examples folder](/examples).

## What's a VPC?

A [VPC](https:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt/aws.amazon.com/vpc) or Virtual Private Cloud is a logically isolated section of your AWS cloud. Each
VPC defines a virtual network within which you run your AWS resources, as well as rules for what can go in and out of
that network. This includes subnets, route tables that tell those subnets how to route inbound and outbound traffic,
security groups, firewalls for the subnet (known as "Network ACLs"), and any other network components such as VPN connections.

## Two Subnet Tiers

This VPC defines two "tiers" of subnets:

- **Public Subnets**: Resources in these subnets are directly addressable from the Internet. Only public-facing
  resources (typically just load balancers and the bastion host) should be put here.
- **https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt/Private/App Subnets**: Resources in these subnets are NOT directly addressable from the Internet but they can make
  outbound connections to the Internet through a NAT Gateway. You can connect to the resources in this subnet only from
  resources within the VPC, so you should put your app servers here and allow the load balancers in the Public Subnet
  to route traffic to them.

## VPC Architecture

The use of a Management VPC is inspired by the VPC Architecture described by Ben Whaley in his blog post [A Reference
VPC Architecture](https:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt/www.whaletech.co/2014/10/02/reference-vpc-architecture.html). That blog post proposed the
following VPC structure:

![VPC Diagram](http:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt/i.imgur.com/KC0OKZL.png)

To summarize:

- The only way operators can access our private network is through a mgmt VPC.
- The mgmt VPC uses [VPC Peering](#vpc-peering) so that, once in the mgmt VPC, you can access any other environment, but
  once in any other environment, you can only access the mgmt VPC (e.g. you cannot access prod from stage).
- We put "environment-agnostic" or management-level resources in the mgmt VPC such as Jenkins, a metrics store, an LDAP
  server, etc.

## VPC Peering

Learn more about VPC Peering in the [vpc-peering](https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-peering) module.

## SSH Access via the Bastion Host

To SSH into any of your EC2 Instances in a private subnet, we recommend launching a single "Bastion Host" to use as
an SSH jump host. For more info, see the [Bastion Host
examples](https:https://github.com/gruntwork-io/module-vpc/tree/master/modules/vpc-mgmt/github.com/gruntwork-io/module-server/tree/master/examples/bastion-host).

## Other VPC Core Concepts

Learn about [Other VPC Core Concepts](https://github.com/gruntwork-io/module-vpc/tree/master/modules/_docs/vpc-core-concepts.md) like subnets, NAT Gateways, and VPC Endpoints.
