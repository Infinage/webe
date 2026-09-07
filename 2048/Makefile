.PHONY: dev watch-css watch-go

dev:
	@$(MAKE) -j2 watch-css watch-go

watch-css:
	tailwindcss -i static/input.css -o static/index.css --watch

watch-go:
	 air
