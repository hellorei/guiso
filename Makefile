# Include all makefiles
include make/*.mk

# Default target
.DEFAULT_GOAL := build

# Phony targets
# .PHONY: $(MAKECMDGOALS) $(shell sed -n -e '/^$$/ { n ; /^[^ .\#][^ ]*:/ { s/:.*$$// ; p ; } ; }' $(MAKEFILE_LIST))
.PHONY: $(shell sed -n -e '/^$$/ { n ; /^[^ .\#][^ ]*:/ { s/:.*$$// ; p ; } ; }' $(MAKEFILE_LIST))

# $(MAKECMDGOALS): guiso_banner

# Environment and OS-specific variables
include make/_env.mk