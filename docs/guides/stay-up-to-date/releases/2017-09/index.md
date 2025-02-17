
# Gruntwork release 2017-09

<p style={{marginTop: "-25px"}}><small><a href="/guides">Guides</a> / <a href="/guides/stay-up-to-date">Update Guides</a> / <a href="/guides/stay-up-to-date/releases">Releases</a> / 2017-09</small></p>

This page is lists all the updates to the [Gruntwork Infrastructure as Code 
Library](https://gruntwork.io/infrastructure-as-code-library/) that were released in 2017-09. For instructions 
on how to use these updates in your code, check out the [updating 
documentation](/guides/working-with-code/using-modules#updating).

Here are the repos that were updated:

- [terraform-aws-asg](#terraform-aws-asg)
- [terraform-aws-kafka](#terraform-aws-kafka)
- [terraform-aws-monitoring](#terraform-aws-monitoring)
- [terraform-aws-security](#terraform-aws-security)
- [terraform-aws-server](#terraform-aws-server)
- [terraform-aws-static-assets](#terraform-aws-static-assets)
- [terraform-aws-zookeeper](#terraform-aws-zookeeper)


## terraform-aws-asg


<<<<<<< Updated upstream
### [v0.6.1](https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.1)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/30/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.1">Release notes</a></small>
=======
### [v0.6.1](https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.1)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/30/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.1">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/module-asg/pull/11: Fix the Python script used by the asg-rolling-deploy module so it properly checks the `tmp` folder to decide whether to extract the boto3 library.
=======
  https://github.com/tnn-gruntwork-io/module-asg/pull/11: Fix the Python script used by the asg-rolling-deploy module so it properly checks the `tmp` folder to decide whether to extract the boto3 library.
>>>>>>> Stashed changes

</div>


<<<<<<< Updated upstream
### [v0.6.0](https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/29/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.0">Release notes</a></small>
=======
### [v0.6.0](https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/29/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-asg/releases/tag/v0.6.0">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/module-asg/pull/10: The `server-group` module now assigns EBS permissions based on the `ServerGroupName` tag instead of the `Name` tag, as the latter is too brittle. This change is backwards incompatible, so we&apos;re bumping the patch version number, but unless you are doing something weird and overriding `ServerGroupName` (very unlikely!), you shouldn&apos;t have to do anything to make this work with your code.
=======
  https://github.com/tnn-gruntwork-io/module-asg/pull/10: The `server-group` module now assigns EBS permissions based on the `ServerGroupName` tag instead of the `Name` tag, as the latter is too brittle. This change is backwards incompatible, so we&apos;re bumping the patch version number, but unless you are doing something weird and overriding `ServerGroupName` (very unlikely!), you shouldn&apos;t have to do anything to make this work with your code.
>>>>>>> Stashed changes

</div>



## terraform-aws-kafka


<<<<<<< Updated upstream
### [v0.0.7](https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.7)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/29/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.7">Release notes</a></small>
=======
### [v0.0.7](https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.7)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/29/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.7">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/package-kafka/pull/7: Update to the new module-asg and package-zookeeper, with the server-group module fix that uses the `ServerGroupName` tag instead of the `Name` tag to assign EBS volume permissions, which should make things less brittle.
=======
  https://github.com/tnn-gruntwork-io/package-kafka/pull/7: Update to the new module-asg and package-zookeeper, with the server-group module fix that uses the `ServerGroupName` tag instead of the `Name` tag to assign EBS volume permissions, which should make things less brittle.
>>>>>>> Stashed changes

</div>


<<<<<<< Updated upstream
### [v0.0.6](https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.6)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/26/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.6">Release notes</a></small>
=======
### [v0.0.6](https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.6)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/26/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-kafka/releases/tag/v0.0.6">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/package-kafka/pull/6: You can now run Kafka on Ubuntu! The major change is actually not in package-kafka itself, but in the updated `attach-eni` script that now supports Ubuntu as of [v0.2.0](https://github.com/tnn-gruntwork-io/module-server/releases/tag/v0.2.0). See the [kafka-ami](https://github.com/tnn-gruntwork-io/package-kafka/tree/master/examples/kafka-ami) example for working sample code.
=======
  https://github.com/tnn-gruntwork-io/package-kafka/pull/6: You can now run Kafka on Ubuntu! The major change is actually not in package-kafka itself, but in the updated `attach-eni` script that now supports Ubuntu as of [v0.2.0](https://github.com/tnn-gruntwork-io/module-server/releases/tag/v0.2.0). See the [kafka-ami](https://github.com/tnn-gruntwork-io/package-kafka/tree/master/examples/kafka-ami) example for working sample code.
>>>>>>> Stashed changes

</div>



## terraform-aws-monitoring


<<<<<<< Updated upstream
### [v0.5.2](https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.2)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/21/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.2">Release notes</a></small>
=======
### [v0.5.2](https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.2)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/21/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.2">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

   #35: Fix a bug where we were using `high_read_latency_threshold` instead of `high_write_latency_threshold` on the `rds_high_write_latency` aws_cloudwatch_metric_alarm.

</div>


<<<<<<< Updated upstream
### [v0.5.1](https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.1)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/15/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.1">Release notes</a></small>
=======
### [v0.5.1](https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.1)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/15/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/releases/tag/v0.5.1">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

  - BUG FIX/#34: The Target Group alarm `tg_high_http_code_target_4xx_count` was previously using variables meant for `tg_high_http_code_target_5xx_count`. This has now been fixed.

</div>



## terraform-aws-security


<<<<<<< Updated upstream
### [v0.6.2](https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.2)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/24/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.2">Release notes</a></small>
=======
### [v0.6.2](https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.2)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/24/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.2">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/module-security/pull/47: ssh-iam should no longer error out when syncing users that changed IAM groups.
=======
  https://github.com/tnn-gruntwork-io/module-security/pull/47: ssh-iam should no longer error out when syncing users that changed IAM groups.
>>>>>>> Stashed changes

</div>


<<<<<<< Updated upstream
### [v0.6.1](https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.1)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/12/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.1">Release notes</a></small>
=======
### [v0.6.1](https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.1)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/12/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.1">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/module-security/pull/46: Add module to install and configure NTP.
=======
  https://github.com/tnn-gruntwork-io/module-security/pull/46: Add module to install and configure NTP.
>>>>>>> Stashed changes

</div>


<<<<<<< Updated upstream
### [v0.6.0](https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/7/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.0">Release notes</a></small>
=======
### [v0.6.0](https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/7/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-security/releases/tag/v0.6.0">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/module-security/pull/45
=======
  https://github.com/tnn-gruntwork-io/module-security/pull/45
>>>>>>> Stashed changes

BACKWARDS INCOMPATIBLE CHANGE

The `iam_groups_for_cross_account_access` input parameter of the `iam-groups` module is now a list of maps rather than a map. This keeps the order of groups more constant when you add groups, rather than trying to delete and recreate all the old groups (note that if you remove a group, the order will still change, which is an unfortunate Terraform limitation: https://github.com/hashicorp/terraform/issues/14275). 

To use the new version of the `iam-groups` module, instead of specifying a map:

```hcl
iam_groups_for_cross_account_access  = &#x7B;
  &quot;stage-full-access&quot;: &quot;arn:aws:iam::123445678910:role/mgmt-full-access&quot;,
  &quot;prod-read-only-access&quot;: &quot;arn:aws:iam::9876543210:role/prod-read-only-access&quot;
&#x7D;
```

You need to specify a list of maps:

```hcl
iam_groups_for_cross_account_access = [
  &#x7B;
    group_name   = &quot;stage-full-access&quot;
    iam_role_arn = &quot;arn:aws:iam::123445678910:role/mgmt-full-access&quot;
  &#x7D;,
  &#x7B;
    group_name   = &quot;prod-read-only-access&quot;
    iam_role_arn = &quot;arn:aws:iam::9876543210:role/prod-read-only-access&quot;
  &#x7D;
]
```

</div>



## terraform-aws-server


<<<<<<< Updated upstream
### [v0.2.0](https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.2.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/25/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.2.0">Release notes</a></small>
=======
### [v0.2.0](https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.2.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/25/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.2.0">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/module-server/pull/18: The `attach-eni` script will now automatically configure route tables on Debian servers. This should allow ENIs to work &quot;automagically&quot; just like they do on Amazon Linux. This release is backwards compatible from an API perspective, but we&apos;ve bumped the minor version number to indicate that it&apos;s a fairly large change in terms of behavior.
=======
  https://github.com/tnn-gruntwork-io/module-server/pull/18: The `attach-eni` script will now automatically configure route tables on Debian servers. This should allow ENIs to work &quot;automagically&quot; just like they do on Amazon Linux. This release is backwards compatible from an API perspective, but we&apos;ve bumped the minor version number to indicate that it&apos;s a fairly large change in terms of behavior.
>>>>>>> Stashed changes

</div>


<<<<<<< Updated upstream
### [v0.1.13](https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.1.13)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/11/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.1.13">Release notes</a></small>
=======
### [v0.1.13](https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.1.13)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/11/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-server/releases/tag/v0.1.13">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/module-server/pull/17: The `single-server` module now exposes `security_group_name` and `iam_group_name` parameters that let you customize the security group and IAM group names, respectively. The default uses the `name` input as before, so this is a backwards compatible change.
=======
  https://github.com/tnn-gruntwork-io/module-server/pull/17: The `single-server` module now exposes `security_group_name` and `iam_group_name` parameters that let you customize the security group and IAM group names, respectively. The default uses the `name` input as before, so this is a backwards compatible change.
>>>>>>> Stashed changes

</div>



## terraform-aws-static-assets


<<<<<<< Updated upstream
### [v0.1.0](https://github.com/tnn-gruntwork-io/terraform-aws-static-assets/releases/tag/v0.1.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/20/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-static-assets/releases/tag/v0.1.0">Release notes</a></small>
=======
### [v0.1.0](https://github.com/tnn-gruntwork-io/terraform-aws-static-assets/releases/tag/v0.1.0)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/20/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-static-assets/releases/tag/v0.1.0">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

  BACKWARDS INCOMPATIBLE CHANGE

<<<<<<< Updated upstream
https://github.com/tnn-gruntwork-io/package-static-assets/pull/3: The `s3-cloudfront` module now allows you to use multiple domain names with your CloudFront distribution. To support this, the following parameters have been renamed:
=======
https://github.com/tnn-gruntwork-io/package-static-assets/pull/3: The `s3-cloudfront` module now allows you to use multiple domain names with your CloudFront distribution. To support this, the following parameters have been renamed:
>>>>>>> Stashed changes

* Input: `create_route53_entry` -&gt; `create_route53_entries` 
* Input: `domain_name` -&gt; `domain_names` and is now a list
* Output: `cloudfront_domain_name` -&gt; `cloudfront_domain_names` and is now a list



</div>



## terraform-aws-zookeeper


<<<<<<< Updated upstream
### [v0.0.7](https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.7)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/29/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.7">Release notes</a></small>
=======
### [v0.0.7](https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.7)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/29/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.7">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/package-zookeeper/pull/7: Update to the new module-asg, with the server-group module fix that uses the `ServerGroupName` tag instead of the `Name` tag to assign EBS volume permissions, which should make things less brittle.
=======
  https://github.com/tnn-gruntwork-io/package-zookeeper/pull/7: Update to the new module-asg, with the server-group module fix that uses the `ServerGroupName` tag instead of the `Name` tag to assign EBS volume permissions, which should make things less brittle.
>>>>>>> Stashed changes

</div>


<<<<<<< Updated upstream
### [v0.0.6](https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.6)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/26/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.6">Release notes</a></small>
=======
### [v0.0.6](https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.6)

<p style={{marginTop: "-20px", marginBottom: "10px"}}>
  <small>Published: 9/26/2017 | <a href="https://github.com/tnn-gruntwork-io/terraform-aws-zookeeper/releases/tag/v0.0.6">Release notes</a></small>
>>>>>>> Stashed changes
</p>

<div style={{"overflow":"hidden","textOverflow":"ellipsis","display":"-webkit-box","WebkitLineClamp":10,"lineClamp":10,"WebkitBoxOrient":"vertical"}}>

<<<<<<< Updated upstream
  https://github.com/tnn-gruntwork-io/package-zookeeper/pull/6: You can now run ZooKeeper on Ubuntu! The major change is actually not in package-zookeeper itself, but in the updated `attach-eni` script that now supports Ubuntu as of [v0.2.0](https://github.com/tnn-gruntwork-io/module-server/releases/tag/v0.2.0). See the [zookeeper-ami example](https://github.com/tnn-gruntwork-io/package-zookeeper/tree/master/examples/zookeeper-ami) for working sample code.
=======
  https://github.com/tnn-gruntwork-io/package-zookeeper/pull/6: You can now run ZooKeeper on Ubuntu! The major change is actually not in package-zookeeper itself, but in the updated `attach-eni` script that now supports Ubuntu as of [v0.2.0](https://github.com/tnn-gruntwork-io/module-server/releases/tag/v0.2.0). See the [zookeeper-ami example](https://github.com/tnn-gruntwork-io/package-zookeeper/tree/master/examples/zookeeper-ami) for working sample code.
>>>>>>> Stashed changes

</div>




<!-- ##DOCS-SOURCER-START
{
  "sourcePlugin": "releases",
  "hash": "8a6374d2defa5976b21ea46fa6fdc485"
}
##DOCS-SOURCER-END -->
