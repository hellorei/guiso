install: ## Install dependencies
	# Check if Go is installed
	@echo "Checking Go installation..."
	@go version || { echo "$(TEAL)$(ICO_PACKAGE_DOWN)$(BOLD)$(ITALIC)$(LAVENDER)Install:$(RESET)$(SURFACE2) Go is not installed. Please install Go first.$(RESET)"; exit 1; }
	@echo "$(TEAL)$(ICO_PACKAGE_DOWN)$(BOLD)$(ITALIC)$(LAVENDER)Install:$(RESET)$(SURFACE2) Installing dependencies...$(RESET)"
	@go install github.com/a-h/templ/cmd/templ@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/evanw/esbuild/cmd/esbuild@latest
	@go install github.com/bokwoon95/wgo@$(WGO_VERSION)
	@$(MAKE) install_tailwind
	@$(MAKE) install_htmx
	@$(MAKE) install_alpinejs
	@echo "$(TEAL)$(ICO_PACKAGE)$(GREEN)(✓)$(SURFACE2) All Dependencies installed!$(RESET)"
