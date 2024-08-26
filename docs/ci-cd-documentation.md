# CI/CD Pipeline Documentation

## Overview

This project uses a GitHub Actions-based CI/CD pipeline to automate testing, building, and deploying our application to Google Cloud Run using Terraform for infrastructure management.

## Motivation

The pipeline aims to:

1. Ensure code quality through automated testing
2. Provide consistent and repeatable deployments
3. Separate staging and production environments for safer releases
4. Automate infrastructure provisioning and updates

## Implementation

The pipeline consists of three main stages:

1. **Test**: Runs automated tests for the application.
2. **Build**: Builds a Docker image and pushes it to Google Artifact Registry.
3. **Deploy**: Uses Terraform to deploy the application to Google Cloud Run.

Two deployment jobs are defined:

- `deploy-staging`: Deploys to staging when changes are pushed to the `main` branch.
- `deploy-production`: Deploys to production when a new tag is pushed.

## Usage

### Continuous Deployment to Staging

1. Develop features in feature branches.
2. Create a pull request to merge into `main`.
3. Once merged, the pipeline automatically deploys to the staging environment.

### Production Releases

1. Ensure all desired changes are in the `main` branch and deployed to staging.
2. Create and push a new tag:
   ```
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. The pipeline will automatically deploy the tagged version to production.

## Modifying the Pipeline

To modify the pipeline:

1. Edit the `.github/workflows/ci-cd.yml` file.
2. Commit and push changes to the `main` branch.
3. GitHub Actions will use the updated workflow for subsequent runs.

## Provisioning Resources

Resources are managed using Terraform:

1. Terraform files are located in the `terraform/` directory.
2. To modify infrastructure, update the relevant `.tf` files.
3. The CI/CD pipeline automatically applies Terraform changes during deployment.

## Environment Variables and Secrets

The following secrets need to be set in GitHub repository settings:

- `GCP_PROJECT_ID`: Your Google Cloud project ID
- `SERVICE_NAME`: The base name for your Cloud Run service
- `GCP_REGION`: The GCP region for deployments
- `REPO_NAME`: The name of your Artifact Registry repository
- `GCP_SA_KEY`: The service account key for GCP authentication

## Manual Deployment

While the CI/CD pipeline automates deployments, you can also deploy manually using Make commands:

1. Ensure you have the necessary permissions and tools installed (gcloud, terraform).
2. Authenticate with Google Cloud: `gcloud auth application-default login`.
3. Set the required environment variables:
   ```
   export PROJECT_ID=your-project-id
   export SERVICE_NAME=your-service-name
   export REGION=your-gcp-region
   export REPO_NAME=your-repo-name
   export IMAGE_URL=your-image-url
   ```
4. Use the following Make commands:
   - To initialize Terraform: `make tf-init`
   - To plan Terraform changes: `make tf-plan`
   - To apply Terraform changes: `make tf-apply`
   - To destroy Terraform resources: `make tf-destroy`

Remember to use different variable values for staging and production environments when deploying manually.

These Make commands simplify the Terraform operations and ensure consistent use of variables across different environments and deployment methods.

## Makefile Structure

The project uses a Makefile to simplify common development and deployment tasks. The Makefile includes targets for both application-specific tasks (like building and testing) and infrastructure management (via Terraform).

All targets in the Makefile are declared as phony targets to ensure they always run, regardless of whether a file with the target's name exists. This includes both application-specific targets and Terraform-related targets.

To see all available make commands, you can run:

```
make help
```

This will display a list of all defined targets and their descriptions, helping you quickly find the command you need for various tasks.

```
.
├── Makefile
└── make/
    ├── app.mk
    ├── ci.mk
    ├── deploy.mk
    ├── dev.mk
    ├── help.mk
    └── terraform.mk
```
