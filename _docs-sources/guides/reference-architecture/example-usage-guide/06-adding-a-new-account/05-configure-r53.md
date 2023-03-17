# Configure Route53

The domain setup process used when setting up a new account is still currently a manual process

### Gruntwork CLI

The current recommended method for configuring Route53 in additional accounts is by using the Gruntwork CLI to bootstrap domain names:

* [Gruntwork CLI: Bootstrap Domain Names](https://github.com/gruntwork-io/gruntwork#bootstrap-the-domain-names)

### Manually in the AWS Console

Alternatively you can manually configure Route53 domains in the AWS Console
  * [AWS Documentation: Registering a new domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/domain-register.html)
  * [AWS Documentation: Configuring DNS routing for a new domain](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-configuring-new-domain.html) (when you register a new domain AWS will automatically configure Route53: creating the hosted zone, creating and configuring the name servers)
