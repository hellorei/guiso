HELP_FUN = \
    %help; \
    while(<>) { \
        if(/^([a-zA-Z_-]+):.*\#\#(.*)$$/) { \
            push(@{$$help{$$1}}, $$2); \
        } \
    }; \
    print "\033[38;2;69;71;90m═══════════════════════════════════════════════════════\033[0m\n\n"; \
    for (sort keys %help) { \
        print " $(OVERLAY2) $(BOLD)$(PEACH)$$_:$(RESET)"; \
        for (@{$$help{$$_}}) { \
            print "$(ITALIC)$(SUBTEXT0)$$_ $(RESET)"; \
        } \
        print "\n"; \
    }

help: guiso_banner ## Show this help
	@echo "$(TEAL)$(ICO_HELP)$(BOLD)$(ITALIC)$(LAVENDER)Help$(RESET)$(SUBTEXT0) Available commands ($(BOLD)command$(RESET)$(SUBTEXT0): $(ITALIC)usage)$(RESET)"
	@echo
	@perl -e '$(HELP_FUN)' $(MAKEFILE_LIST)
	@echo
