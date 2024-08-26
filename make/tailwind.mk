# Tailwind
TAILWIND_VERSION=v3.4.10
TAILWIND_BINARY=tailwindcss-linux-x64
TAILWIND_INPUT := ./global.css
TAILWIND_OUTPUT := ./public/css/global.css
TAILWIND_TIMEOUT := 60

ifeq ($(UNAME_S),Darwin)
  ifeq ($(UNAME_M),arm64)
    TAILWIND_BINARY=tailwindcss-macos-arm64
  else
    TAILWIND_BINARY=tailwindcss-macos-x64
  endif
else ifeq ($(UNAME_S),Linux)
  ifeq ($(UNAME_M),x86_64)
    TAILWIND_BINARY=tailwindcss-linux-x64
  else ifeq ($(UNAME_M),armv7l)
    TAILWIND_BINARY=tailwindcss-linux-armv7
  else ifeq ($(UNAME_M),aarch64)
    TAILWIND_BINARY=tailwindcss-linux-arm64
  else
    $(error Unsupported Linux architecture: $(UNAME_M))
  endif
else ifeq ($(OS),Windows_NT)
  ifeq ($(PROCESSOR_ARCHITECTURE),AMD64)
    TAILWIND_BINARY=tailwindcss-windows-x64.exe
  else ifeq ($(PROCESSOR_ARCHITECTURE),ARM64)
    TAILWIND_BINARY=tailwindcss-windows-arm64.exe
  else
    $(error Unsupported Windows architecture: $(PROCESSOR_ARCHITECTURE))
  endif
else
  $(error Unsupported operating system: $(UNAME_S))
endif

install_tailwind: ## Install Tailwind CSS
	@echo "$(TEAL)$(HEADER_TAILWIND_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TAILWIND)$(RESET)$(SURFACE2) Installing Tailwind CSS...$(RESET)"
	@echo "Operating System: $(UNAME_S)"
	@echo "Architecture: $(UNAME_M)"
	@echo "Selected Tailwind Binary: $(TAILWIND_BINARY)"
	@mkdir -p $(GOBIN)
	@rm -f $(GOBIN)/tailwindcss
	@if command -v wget >/dev/null 2>&1; then \
		wget -q https://github.com/tailwindlabs/tailwindcss/releases/download/$(TAILWIND_VERSION)/$(TAILWIND_BINARY) -O $(GOBIN)/tailwindcss && chmod +x $(GOBIN)/tailwindcss; \
	elif command -v curl >/dev/null 2>&1; then \
		curl -sL https://github.com/tailwindlabs/tailwindcss/releases/download/$(TAILWIND_VERSION)/$(TAILWIND_BINARY) -o $(GOBIN)/tailwindcss && chmod +x $(GOBIN)/tailwindcss; \
	else \
		echo "Error: Neither wget nor curl is available. Please install one of them."; \
		exit 1; \
	fi
	@chmod +x $(GOBIN)/tailwindcss
	@echo "$(TEAL)$(HEADER_TAILWIND_ICO)$(GREEN)(✓)$(SURFACE2) Tailwind CSS downloaded!$(RESET)"
	@if [ -x "$(GOBIN)/tailwindcss" ]; then \
		echo "Tailwind version:"; \
		$(GOBIN)/tailwindcss --help | head -n 1 || echo "Failed to get Tailwind version"; \
	else \
		echo "$(RED)Error: Tailwind binary is not executable$(RESET)"; \
		exit 1; \
	fi
	@echo "$(TEAL)$(HEADER_TAILWIND_ICO)$(GREEN)(✓)$(SURFACE2) Tailwind CSS installation verified!$(RESET)"

tailwind: ## Build Tailwind CSS
	@echo "$(TEAL)$(HEADER_TAILWIND_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TAILWIND)$(RESET)$(SURFACE2) Building Tailwind CSS...$(RESET)"
	@echo "$(SURFACE2)$(HEADER_TAILWIND_ICO)Tailwind version: $(TAILWIND_VERSION)$(RESET)"
	@echo "$(SURFACE2)$(HEADER_TAILWIND_ICO)Tailwind input file: $(TAILWIND_INPUT)$(RESET)"
	@echo "$(SURFACE2)$(HEADER_TAILWIND_ICO)Tailwind output file: $(TAILWIND_OUTPUT)$(RESET)"
	@$(GOBIN)/tailwindcss --help | head -n 1 || echo "Failed to get Tailwind version"
	@echo "$(SURFACE2)$(HEADER_TAILWIND_ICO)Running Tailwind build...$(RESET)"
	@timeout $(TAILWIND_TIMEOUT)s $(GOBIN)/tailwindcss -i $(TAILWIND_INPUT) -o $(TAILWIND_OUTPUT) --minify || (echo "Tailwind build timed out or failed" && exit 1)
	@echo "$(TEAL)$(HEADER_TAILWIND_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_TAILWIND)$(RESET)$(GREEN) Tailwind build completed!$(RESET)"
	@echo
