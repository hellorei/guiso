# Cleanup files
TEMPL_PID_FILE := $(GOBIN)/.templ_pid
TAILWIND_PID_FILE := $(GOBIN)/.tailwind_pid
ESBUILD_PID_FILE := $(GOBIN)/.esbuild_pid
FSWATCH_PID_FILE := $(GOBIN)/.fswatch_pid
GO_PID_FILE := $(GOBIN)/.go_pid

cleanup: ## Cleanup files
	@echo "$(TEAL)$(HEADER_CLEAN_ICO)$(BOLD)$(ITALIC)$(OVERLAY1)Shutting Down:$(RESET)$(SURFACE2) Cleaning up processes...$(RESET)"
	@pkill -f tailwindcss || true
	@pkill -f esbuild || true
	@pkill -f fswatch || true
	@pkill -f "go run ./app" || true
	@pkill -f "templ generate" || true
	@echo "$(TEAL)󱠡 $(BOLD)$(ITALIC)$(LAVENDER)Done!$(RESET)"

clean_templ: ## Clean templ files
	@echo "$(TEAL)$(HEADER_CLEAN_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_CLEAN) Templ:$(RESET)$(SURFACE2) Cleaning templ files...$(RESET)"
	@find . -name "*_templ.go" -type f -delete
	@rm -rf ~/.cache/templ
	@echo "$(TEAL)$(HEADER_CLEAN_ICO)$(GREEN)(✓)$(SURFACE2) Templ files cleaned!$(RESET)"

clean_generated: clean_templ ## Clean generated files
	@echo "$(TEAL)$(HEADER_CLEAN_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_CLEAN) Generated:$(RESET)$(SURFACE2) Cleaning generated files...$(RESET)"
	@rm -f ./public/out/*
	@rm -f ./public/js/dist/*
	@rm -f ./public/css/global.css
	@echo "$(TEAL)$(HEADER_CLEAN_ICO)$(GREEN)(✓)$(SURFACE2) Generated files cleaned!$(RESET)"

clean_all: clean_generated ## Perform full clean (including tools)
	@echo "$(TEAL)$(HEADER_CLEAN_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_CLEAN) Go:$(RESET)$(SURFACE2) Performing full clean (including tools)...$(RESET)"
	@go clean
	@rm -rf $(GOBIN)
	@echo "$(TEAL)$(HEADER_CLEAN_ICO)$(GREEN)(✓)$(SURFACE2) Cleaning done!$(RESET)"
