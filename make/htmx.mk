# HTMX
HTMX_VERSION=1.9.10
HTMX_URL=https://unpkg.com/htmx.org@$(HTMX_VERSION)/dist/htmx.min.js

install_htmx: ## Download HTMX
	@echo "$(TEAL)$(HEADER_HTMX_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_HTMX)$(RESET)$(SURFACE2) Downloading HTMX...$(RESET)"
	@mkdir -p public/js
	@if command -v wget >/dev/null 2>&1; then \
		wget -q $(HTMX_URL) -O public/js/htmx.min.js; \
	elif command -v curl >/dev/null 2>&1; then \
		curl -sL $(HTMX_URL) -o public/js/htmx.min.js; \
	else \
		echo "Error: Neither wget nor curl is available. Please install one of them."; \
		exit 1; \
	fi
	@echo "$(TEAL)$(HEADER_HTMX_ICO)$(GREEN)(✓)$(SURFACE2) HTMX downloaded!$(RESET)"
	@echo
