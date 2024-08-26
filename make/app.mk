build_frontend: tailwind esbuild install_htmx ## Build frontend assets

build: deps tailwind esbuild generate ## Build the application
	@echo "$(TEAL)$(ICO_BUILD)$(BOLD)$(ITALIC)$(LAVENDER)Build:$(RESET)$(SURFACE2) Building..."
	@go build $(LDFLAGS) -o $(GOBIN)/$(BINARY_NAME) $(MAIN_PACKAGE)
	@echo "$(TEAL)$(ICO_BUILD)$(GREEN)(✓)$(SURFACE2) Build done!$(RESET)"

run: build ## Run the application
	@echo "$(TEAL)$(ICO_RUN)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_RUN)$(RESET)$(SURFACE2) Running...$(RESET)"
	@$(GOBIN)/$(BINARY_NAME)

deps: ## Ensure dependencies are up to date
	@echo "$(TEAL)$(HEADER_GO_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_GO)$(RESET)$(SURFACE2) Ensuring dependencies are up to date...$(RESET)"
	@go mod tidy
	@go mod verify
	@echo "$(TEAL)$(HEADER_GO_ICO)$(GREEN)(✓)$(SURFACE2) Dependencies up to date!$(RESET)"

generate: ## Generate templ files
	@echo "$(TEAL)$(HEADER_TEMPL_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TEMPL)$(RESET)$(SURFACE2) Generating templ files...$(RESET)"
	@templ generate
	@echo "$(TEAL)$(HEADER_TEMPL_ICO)$(GREEN)(✓)$(SURFACE2) Templ files generated!$(RESET)"
