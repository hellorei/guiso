# Alpine.js installation and build configuration
ALPINE_VERSION := 3.13.3
ALPINE_FILE := alpine.min.js
ALPINE_OUTPUT_DIRECTORY := public/js
ALPINE_URL := https://cdn.jsdelivr.net/npm/alpinejs@$(ALPINE_VERSION)/dist/cdn.min.js

install_alpinejs: ## Download Alpine.js
	@echo "$(TEAL)$(HEADER_ALPINE_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_ALPINE)$(RESET)$(SURFACE2) Downloading Alpine.js...$(RESET)"
	@mkdir -p $(ALPINE_OUTPUT_DIRECTORY)
	@if command -v wget >/dev/null 2>&1; then \
		wget -q $(ALPINE_URL) -O $(ALPINE_OUTPUT_DIRECTORY)/$(ALPINE_FILE); \
	elif command -v curl >/dev/null 2>&1; then \
		curl -sL $(ALPINE_URL) -o $(ALPINE_OUTPUT_DIRECTORY)/$(ALPINE_FILE); \
	else \
		echo "Error: Neither wget nor curl is available. Please install one of them."; \
		exit 1; \
	fi
	@echo "$(TEAL)$(HEADER_ALPINE_ICO)$(GREEN)(✓)$(SURFACE2) Alpine.js downloaded!$(RESET)"
	@echo