# Logs

All of your services have been configured using the [`cloudwatch-agent`
<<<<<<< Updated upstream
module](https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/tree/master/modules/agents/cloudwatch-agent)
=======
module](https://github.com/tnn-gruntwork-io/terraform-aws-monitoring/tree/master/modules/agents/cloudwatch-agent)
>>>>>>> Stashed changes
to send their logs to [CloudWatch Logs](https://console.aws.amazon.com/cloudwatch/home?#logs:). Instead of SSHing to
each server to see a log file, and worrying about losing those log files if the server fails, you can just go to the
[CloudWatch Logs Page](https://console.aws.amazon.com/cloudwatch/home?#logs:) and browse and search log events for all
your servers in near-real-time.
