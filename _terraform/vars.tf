# ---------------------------------------------------------------------------------------------------------------------
# ENVIRONMENT VARIABLES
# Define these secrets as environment variables
# ---------------------------------------------------------------------------------------------------------------------

# AWS_ACCESS_KEY_ID
# AWS_SECRET_ACCESS_KEY

# ---------------------------------------------------------------------------------------------------------------------
# MODULE PARAMETERS
# These variables are expected to be passed in by the operator
# ---------------------------------------------------------------------------------------------------------------------

variable "aws_region" {
  description = "The AWS region in which all resources will be created"
  default = "us-east-2"
}

variable "aws_account_ids" {
  description = "A list of AWS Account IDs. Only these IDs may be operated on by this template. The first account ID is considered the primary, where all the resources should be created."
  type = "list"
  default = ["466851281301"]
}

variable "website_domain_name" {
  description = "The fully-qualified domain name of the Gruntwork docs website. E.g. docs.gruntwork.io"
  default = "docs.gruntwork.io"
}