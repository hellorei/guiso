test: build ## Run tests
	@echo "$(GREEN)👨‍🍳 Running tests...$(RESET)"
	@go test ./...

lint: ## Run linter
	@echo "$(YELLOW)🕵️ Linting...$(RESET)"
	@golangci-lint run

fmt: ## Format code
	@echo "$(BLUE)👷 Formatting...$(RESET)"
	@go fmt ./...

vet: ## Vet code
	@echo "$(PURPLE)👮‍♀️ Vetting...$(RESET)"
	@go vet ./...

ci: fmt vet lint test ## Run all CI tasks
	@echo "$(CYAN)Running CI tasks: \n\tfmt\n\tvet\n\tlint\n\ttest$(RESET)"