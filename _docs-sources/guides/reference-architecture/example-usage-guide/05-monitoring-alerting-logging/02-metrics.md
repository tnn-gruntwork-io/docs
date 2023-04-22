# Metrics

You can find all the metrics for your AWS account on the [CloudWatch Metrics
Page](https://console.aws.amazon.com/cloudwatch/home?#metricsV2:).

- Most AWS services emit metrics by default, which you'll find under the "AWS Namespaces" (e.g. EC2, ECS, RDS).

- Custom metrics show up under "Custom Namespaces." In particular, the [`cloudwatch-memory-disk-metrics-scripts`
<<<<<<< Updated upstream
  module](https://github.com/tnn-tnn-tnn-tnn-tnn-gruntwork-io/terraform-aws-monitoring/tree/master/modules/metrics/) is installed on every
=======
  module](https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/tree/master/modules/metrics/) is installed on every
>>>>>>> Stashed changes
  server to emit metrics not available from AWS by default, including memory and disk usage. You'll find these under
  the "Linux System" Namespace.

You may want to create a [Dashboard](https://console.aws.amazon.com/cloudwatch/home?#dashboards:)
with the most useful metrics for your services and have that open on a big screen at all times.
