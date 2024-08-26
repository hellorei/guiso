dev: guiso_initializing build_frontend guiso_banner ## Start development server with hot reload.
	@echo "$(TEAL)$(HEADER_DEV_ICO)$(BOLD)$(ITALIC)$(LAVENDER)$(HEADER_DEV)$(RESET)$(SUBTEXT0) Starting development server with hot reload...$(RESET)"
	@echo
	@trap 'make cleanup; exit 0' EXIT; \
	( \
		templ generate --watch --proxy=http://localhost:3000 \
			--cmd="go run ./app" \
			--proxybind=127.0.0.1 & \
		$(GOBIN)/tailwindcss -i ./global.css -o ./public/css/output.css --watch & \
		esbuild app/ts/*.ts --bundle --outdir=public/out --watch & \
		fswatch -o public/out | xargs -n1 -I{} templ generate --notify-proxy & \
		fswatch -o ./locales/*.toml | xargs -n1 -I{} sh -c 'templ generate --notify-proxy && echo "Regenerated templ files due to locale changes"' & \
		wait \
	)
	@echo

dev-clean: clean_generated dev ## Start development server with hot reload and clean generated files before starting.
