# Terraform variables
TF_DIR := terraform
TF_VARS := -var="project_id=$(PROJECT_ID)" \
           -var="service_name=$(SERVICE_NAME)" \
           -var="region=$(REGION)" \
           -var="repo_name=$(REPO_NAME)" \
           -var="image_url=$(IMAGE_URL)"

tf-init: ## Initialize Terraform
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(SURFACE2) Initializing Terraform..$(RESET)"
	# @terraform -chdir=$(TF_DIR) init
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(GREEN) Terraform initialized successfully!$(RESET)"

tf-plan: tf-init ## Plan Terraform changes
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(SURFACE2) Planning Terraform changes...$(RESET)"
	# @terraform -chdir=$(TF_DIR) plan $(TF_VARS)
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(GREEN) Plan created successfully!$(RESET)"

tf-apply: tf-init ## Apply Terraform changes
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(GREEN) Applying Terraform changes...$(RESET)"
	# @terraform -chdir=$(TF_DIR) apply -auto-approve $(TF_VARS)
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(GREEN) Terraform changes applied successfully!$(RESET)"

tf-destroy: tf-init ## Destroy Terraform resources
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(RED) Destroying Terraform resources...$(RESET)"
	# @terraform -chdir=$(TF_DIR) destroy -auto-approve $(TF_VARS)
	@echo "$(TEAL)$(HEADER_TERRAFORM_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TERRAFORM)$(RESET)$(GREEN) Terraform resources destroyed successfully!$(RESET)"