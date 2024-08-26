# ANSI color codes
ROSEWATER := \033[38;2;245;224;220m
FLAMINGO := \033[38;2;242;205;205m
PINK := \033[38;2;245;194;231m
MAUVE := \033[38;2;203;166;247m
RED := \033[38;2;243;139;168m
MAROON := \033[38;2;235;160;172m
PEACH := \033[38;2;250;179;135m
YELLOW := \033[38;2;249;226;175m
GREEN := \033[38;2;166;227;161m
TEAL := \033[38;2;148;226;213m
SKY := \033[38;2;137;220;235m
SAPPHIRE := \033[38;2;116;199;236m
BLUE := \033[38;2;137;180;250m
LAVENDER := \033[38;2;180;190;254m
LAVENDER_LIGHT := \033[38;2;180;190;254m
LAVENDER_DARK := \033[38;2;127;132;156m
TEXT := \033[38;2;205;214;244m
SUBTEXT1 := \033[38;2;186;194;222m
SUBTEXT0 := \033[38;2;166;173;200m
OVERLAY2 := \033[38;2;147;153;178m
OVERLAY1 := \033[38;2;127;132;156m
OVERLAY0 := \033[38;2;108;112;134m
SURFACE2 := \033[38;2;88;91;112m
SURFACE1 := \033[38;2;69;71;90m
SURFACE0 := \033[38;2;69;71;90m
BASE := \033[38;2;30;30;46m
MANTLE := \033[38;2;24;24;37m
CRUST := \033[38;2;17;17;27m

BOLD := \033[1m
ITALIC := \033[3m

# Reset color
RESET := \033[0m

SEPARATOR := \033[38;2;69;71;90m═══════════════════════════════════════════════════════\033[0m

HEADER_TERRAFORM := "Terraform:"
HEADER_TERRAFORM_ICO := "󱁢 "
HEADER_TAILWIND := "Tailwind:"
HEADER_TAILWIND_ICO := "󱏿 "
HEADER_HTMX := "HTMX:"
HEADER_HTMX_ICO := "󱁡 "
HEADER_GO_ICO := "󱁢 "
HEADER_GO := "Go: :"
HEADER_GO_ICO := "󱁢 "
HEADER_ESBUILD := "Esbuild:"
HEADER_ESBUILD_ICO := " "
HEADER_ENV := "ENV:"
HEADER_ENV_ICO := "󱄑 "
HEADER_DOCKER := "Docker:"
HEADER_DOCKER_ICO := " "
HEADER_RUN := "Run:"
HEADER_RUN_ICO := " "
HEDER_TEMPL := "Templ:"
HEDER_TEMPL_ICO := " "
HEADER_CLEAN := "Cleaning"
HEADER_CLEAN_ICO := "󰃢 "
HEADER_DEV := "Dev:"
HEADER_DEV_ICO := " "
HEADER_ALPINE := "Alpine.js:"
HEADER_ALPINE_ICO := " "

ICO_BROOM := "󰃢 "
ICO_CONTAINER := " "
ICO_DEPLOY := " "
ICO_TEST := "󰙨 "
ICO_AB_TEST := "󰇉 "
ICO_PLAN := " "
ICO_HELP := "󰋖 "
ICO_TORI := " "
ICO_PACKAGE := "󰏓 "
ICO_PACKAGE_DOWN := "󰏔 "

version: ## Show version
	@echo "$(TEAL) $(LAVENDER)GUISO VERSION:$(RESET) $(VERSION)"

env: ## Show environment
	@echo "$(TEAL)$(HEADER_ENV_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_ENV)$(RESET)$(SUBTEXT0) Environment information$(RESET)"
	@echo "$(SEPARATOR)"
	@echo "$(TEAL) $(LAVENDER)GUISO VERSION:$(RESET) $(VERSION)"
	@echo "$(TEAL)󰫧 $(LAVENDER)ENV:$(RESET) $(ENV)"
	@echo "$(TEAL)󰻾 $(LAVENDER)PROJECT_ID:$(RESET) $(PROJECT_ID)"
	@echo "$(TEAL)󰢍 $(LAVENDER)SERVICE_NAME:$(RESET) $(SERVICE_NAME)"
	@echo "$(TEAL) $(LAVENDER)REGION:$(RESET) $(REGION)"
	@echo "$(TEAL) $(LAVENDER)REPO_NAME:$(RESET) $(REPO_NAME)"
	@echo "$(TEAL) $(LAVENDER)IMAGE_URL:$(RESET) $(IMAGE_URL)"
	@echo "$(TEAL)󰟓 $(LAVENDER)GOPATH:$(RESET) $(GOPATH)"
	@echo "$(TEAL)󰟓 $(LAVENDER)GOBIN:$(RESET) $(GOBIN)"

define GUISO_BANNER
	@echo "\033[38;2;148;226;213m   \033[38;2;141;221;218m   \033[38;2;134;216;223m   \033[38;2;127;211;228m   \033[38;2;120;206;233m   \033[38;2;116;199;236m _           "
	@echo "\033[38;2;148;226;213m   \033[38;2;141;221;218m____ \033[38;2;134;216;223m___ \033[38;2;127;211;228m __(\033[38;2;120;206;233m_)__\033[38;2;116;199;236m______  "
	@echo "\033[38;2;148;226;213m  / \033[38;2;141;221;218m__ \`\033[38;2;134;216;223m/ / \033[38;2;127;211;228m/ / \033[38;2;120;206;233m/ __\033[38;2;116;199;236m_/ __ \ "
	@echo "\033[38;2;148;226;213m / /\033[38;2;141;221;218m_/ /\033[38;2;134;216;223m /_\033[38;2;127;211;228m/ / \033[38;2;120;206;233m(__  \033[38;2;116;199;236m) /_/ / "
	@echo "\033[38;2;148;226;213m \_\033[38;2;141;221;218m_, /\033[38;2;134;216;223m\__,\033[38;2;127;211;228m_/_/\033[38;2;120;206;233m____/\033[38;2;116;199;236m\____/  "
	@echo "\033[38;2;148;226;213m/__\033[38;2;141;221;218m__/\033[38;2;88;91;112m/\033[38;2;96;101;125m/\033[38;2;104;111;138m/\033[38;2;112;121;151m/\033[38;2;120;131;164m/\033[38;2;128;141;177m/\033[38;2;136;151;190m/\033[38;2;144;161;203m/\033[38;2;152;171;216m/\033[38;2;160;181;229m/\033[38;2;168;181;242m/\033[38;2;176;190;254m/\033[38;2;180;190;254mv.$(VERSION)"
	@echo "\033[38;2;69;71;90m═══════════════════════════════════════════════════════\033[0m"
	@echo "\033[38;2;166;173;200mA customizable end-to-end framework for"
	@echo "building high-performance, containerized applications,"
	@echo "using Go, HTMX, Templ, Tailwind, and Terraform.\033[0m"
	@echo "\033[38;2;69;71;90m═══════════════════════════════════════════════════════\033[0m"
	@echo "\033[3m\033[38;2;116;199;236m Docs: \033[1mhttps://www.guiso.io\033[0m"
	@echo "\033[38;2;69;71;90m═══════════════════════════════════════════════════════\033[0m"
endef

guiso_banner:
	$(GUISO_BANNER)
	@echo

guiso_initializing:
	@echo
	@echo "$(TEAL)$(HEADER_GUISO)$(BOLD)$(ITALIC)$(LAVENDER)Initializing \033[38;2;148;226;213mG\033[38;2;141;221;218mU\033[38;2;134;216;223mI\033[38;2;127;211;228mS\033[38;2;120;206;233mO\033[38;2;116;199;236m⫻⫻⫻$(RESET)"
	@echo
