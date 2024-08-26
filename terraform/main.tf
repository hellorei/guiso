terraform {
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.42"
    }
    null = {
      source  = "hashicorp/null"
      version = "~> 3.2"
    }
  }
}

provider "google" {
  project = var.project_id
  region  = var.region
}

resource "google_project_service" "serviceusage" {
  service = "serviceusage.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "artifact_registry_api" {
  service = "artifactregistry.googleapis.com"
  disable_on_destroy = false
  depends_on = [google_project_service.serviceusage]
}

resource "google_project_service" "run_api" {
  service = "run.googleapis.com"
  disable_on_destroy = false
}

resource "google_project_service" "container_registry_api" {
  service = "containerregistry.googleapis.com"
  disable_on_destroy = false
}

resource "google_cloud_run_v2_service" "app" {
  name     = var.service_name
  location = var.region
  ingress  = "INGRESS_TRAFFIC_ALL"

  template {
    containers {
      image = var.image_url

      ports {
        container_port = 8080
      }

      env {
        name  = "STRIPE_URL"
        value = var.STRIPE_URL
      }
      env {
        name  = "GA4_MEASUREMENT_ID"
        value = var.ga4_measurement_id
      }
      env {
        name  = "GA4_API_SECRET"
        value = var.ga4_api_secret
      }

      resources {
        limits = {
          cpu    = var.cpu_limit
          memory = var.memory_limit
        }
      }
    }
    scaling {
      min_instance_count = var.min_instance_count
      max_instance_count = var.max_instance_count
    }
  }

  traffic {
    type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
    percent = 100
  }

  lifecycle {
    create_before_destroy = true
  }
}

# # Force new deployment by adding a revision suffix with commit SHA
# resource "google_cloud_run_v2_service" "revision" {
#   name     = "${var.service_name}-${substr(var.commit_sha, 0, 8)}"
#   location = var.region

#   template {
#     containers {
#       image = var.image_url
#     }
#   }

#   traffic {
#     type    = "TRAFFIC_TARGET_ALLOCATION_TYPE_LATEST"
#     percent = 100
#   }

#   lifecycle {
#     create_before_destroy = true
#   }

#   depends_on = [google_cloud_run_v2_service.app]
# }

resource "null_resource" "image_url_check" {
  count = var.image_url == "" ? "Please provide a valid image_url" : 0
}

locals {
  authorized_users_list = split(",", var.authorized_users)
}

resource "google_cloud_run_v2_service_iam_member" "invoker" {
  for_each = toset(local.authorized_users_list)
  project  = google_cloud_run_v2_service.app.project
  location = google_cloud_run_v2_service.app.location
  name     = google_cloud_run_v2_service.app.name
  role     = "roles/run.invoker"
  member   = each.value
}

resource "google_project_iam_member" "github_actions_deployer_roles" {
  for_each = toset([
    "roles/artifactregistry.writer",
    "roles/run.admin",
    "roles/iam.serviceAccountUser",
    "roles/storage.objectCreator",
    "roles/iam.workloadIdentityUser"
  ])
  
  project = var.project_id
  role    = each.key
  member  = "serviceAccount:github-actions-deployer@${var.project_id}.iam.gserviceaccount.com"
}
