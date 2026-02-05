.PHONY: help tag push-tag push-tags release

VERSION ?= 

help: ## Display this help message
	@echo "Available targets:"
	@echo "  tag        - Create a git tag (use: make tag VERSION=v1.0.0)"
	@echo "  push-tag   - Push a specific tag to remote (use: make push-tag VERSION=v1.0.0)"
	@echo "  push-tags  - Push all tags to remote"
	@echo "  release    - Create and push a tag (use: make release VERSION=v1.0.0)"

tag: ## Create a git tag (use: make tag VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make tag VERSION=v1.0.0"; \
		exit 1; \
	fi
	git tag -a $(VERSION) -m "Release $(VERSION)"
	@echo "Tag $(VERSION) created!"

push-tag: ## Push a specific tag to remote (use: make push-tag VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make push-tag VERSION=v1.0.0"; \
		exit 1; \
	fi
	git push origin $(VERSION)
	@echo "Tag $(VERSION) pushed!"

push-tags: ## Push all tags to remote
	git push --tags
	@echo "All tags pushed!"

release: ## Create and push a tag (use: make release VERSION=v1.0.0)
	@if [ -z "$(VERSION)" ]; then \
		echo "Error: VERSION is required. Usage: make release VERSION=v1.0.0"; \
		exit 1; \
	fi
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)
	@echo "Release $(VERSION) complete!"