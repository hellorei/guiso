docker-build: ## Build Docker image
	@echo "$(BLUE)🐳 Building Docker image...$(RESET)"
	@docker build -t $(BINARY_NAME) .

docker-deploy-build: ## Build Docker image for deployment
	@echo "$(BLUE)🐳 Building Docker image for deployment...$(RESET)"
	@docker build \
		--build-arg STRIPE_URL=$(STRIPE_URL) BINARY_NAME=$(BINARY_NAME) \
		-t $(BINARY_NAME) \
		.

docker-run: ## Run Docker container
	@echo "$(BLUE)🐳 Running Docker container...$(RESET)"
	@docker run -p 8080:8080 $(BINARY_NAME)

docker: docker-build docker-run ## Build and run Docker container

docker-push: ## Push Docker image
	@echo "$(BLUE)🐳 Pushing Docker image...$(RESET)"
	@docker push $(BINARY_NAME)

docker-tag: ## Tag Docker image
	@echo "$(BLUE)🐳 Tagging Docker image...$(RESET)"
	@docker tag $(BINARY_NAME) $(BINARY_NAME):$(GIT_TAG)

docker-tag-latest: ## Tag Docker image as latest
	@echo "$(BLUE)🐳 Tagging Docker image as latest...$(RESET)"
	@docker tag $(BINARY_NAME) $(BINARY_NAME):latest

docker-push-tag: ## Push Docker image with tag
	@echo "$(BLUE)🐳 Pushing Docker image with tag...$(RESET)"
	@docker push $(BINARY_NAME):$(GIT_TAG)

docker-push-latest: ## Push Docker image as latest
	@echo "$(BLUE)🐳 Pushing Docker image as latest...$(RESET)"
	@docker push $(BINARY_NAME):latest

docker-push-tag-latest: docker-push-tag docker-push-latest ## Push Docker image with tag and latest