ESBUILD_OUTPUT := ./public/js/dist

install_esbuild: ## Install esbuild
	@echo "$(TEAL)$(HEADER_ESBUILD_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_ESBUILD)$(RESET)$(SURFACE2) Installing esbuild...$(RESET)"
	@mkdir -p $(GOBIN)
	@if command -v wget >/dev/null 2>&1; then \
		wget -qO- https://github.com/evanw/esbuild/releases/download/v$(ESBUILD_VERSION)/esbuild-$(shell uname -s | tr '[:upper:]' '[:lower:]')-$(shell uname -m).tar.gz | tar -xzC $(GOBIN); \
	elif command -v curl >/dev/null 2>&1; then \
		curl -fsSL https://github.com/evanw/esbuild/releases/download/v$(ESBUILD_VERSION)/esbuild-$(shell uname -s | tr '[:upper:]' '[:lower:]')-$(shell uname -m).tar.gz | tar -xzC $(GOBIN); \
	else \
		echo "Error: Neither wget nor curl is available. Please install one of them."; \
		exit 1; \
	fi
	@echo "$(TEAL)$(HEADER_ESBUILD_ICO)$(GREEN)(✓)$(SURFACE2) esbuild installed!$(RESET)"

esbuild: $(ESBUILD_OUTPUT) ## Build TypeScript with esbuild
$(ESBUILD_OUTPUT): $(FRONTEND_SOURCES)
	@echo "$(TEAL)$(HEADER_ESBUILD_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_ESBUILD)$(RESET)$(SURFACE2) Building TypeScript with esbuild...$(RESET)"
	@esbuild --minify app/ts/*.ts --bundle --outdir=$(ESBUILD_OUTPUT) --external:htmx.org
	@echo
