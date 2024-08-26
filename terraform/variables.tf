variable "project_id" {
  description = "The GCP project ID"
  type        = string
}

variable "service_name" {
  description = "The name of the Cloud Run service"
  type        = string
}

variable "region" {
  description = "The GCP region"
  type        = string
}

variable "image_url" {
  description = "The URL of the Docker image to deploy"
  type        = string
}

variable "authorized_users" {
  type        = string
  description = "A comma-separated list of authorized users, groups, and service accounts"
  default     = ""
}

variable "cpu_limit" {
  type        = string
  description = "The CPU limit for the container"
  default     = "1"
}

variable "memory_limit" {
  type        = string
  description = "The memory limit for the container"
  default     = "1Gi"
}

variable "min_instance_count" {
  type        = number
  description = "The minimum number of instances to run"
  default     = 0
}

variable "max_instance_count" {
  type        = number
  description = "The maximum number of instances to run"
  default     = 1
}

variable "STRIPE_URL" {
  type        = string
  description = "The Stripe donation link"
}

variable "ga4_measurement_id" {
  type        = string
  description = "The Google Analytics 4 Measurement ID"
}

variable "ga4_api_secret" {
  type        = string
  description = "The Google Analytics 4 API Secret"
}

variable "commit_sha" {
  type        = string
  description = "The SHA of the commit being deployed"
}