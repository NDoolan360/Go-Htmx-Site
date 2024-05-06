CURRENT_DIR=$(dir $(abspath $(firstword $(MAKEFILE_LIST))))

# Export all variables
export

# Set up environment variables
ENVFILE=$(CURRENT_DIR).env
ifneq ("$(wildcard $(ENVFILE))","")
	include $(ENVFILE)
endif

# Get OS & ARCH info
SYSTEM := $(shell uname -s)
ifeq ($(SYSTEM),Linux)
    OS=linux
else ifeq ($(SYSTEM),Darwin)
    OS=macos
endif

PLATFORM := $(shell uname -m)
ifeq ($(PLATFORM),x86_64)
    ARCH=x64
else ifeq ($(PLATFORM),arm64)
    ARCH=arm64
endif


#################################################################################
# File Targets                                                                  #
#################################################################################

TEMP_DIR=$(CURRENT_DIR)tmp
TEMPL_TAR=$(TEMP_DIR)/templ_$(SYSTEM)_$(PLATFORM).tar.gz
TEMPL_BINARY=$(TEMP_DIR)/templ
TEMPL_URL=https://github.com/a-h/templ/releases/latest/download/templ_$(SYSTEM)_$(PLATFORM).tar.gz
TAILWIND_BINARY=$(TEMP_DIR)/tailwindcss-$(OS)-$(ARCH)
TAILWIND_URL=https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(OS)-$(ARCH)
TAILWIND_STYLES=$(CURRENT_DIR)website/static/styles/tailwind.css
STYLES=$(CURRENT_DIR)website/static/styles/styles.css
ALL_PROJECTS=./api/experience/... ./api/health/... ./api/projects/... ./website/...

$(TEMP_DIR):
	@mkdir -p $(TEMP_DIR)

$(TEMPL_TAR): $(TEMP_DIR)
	curl -sLo $(TEMPL_TAR) $(TEMPL_URL)

$(TEMPL_BINARY): $(TEMP_DIR) $(TEMPL_TAR)
	tar -xvzf $(TEMPL_TAR) -C $(TEMP_DIR) templ
	@chmod +x $(TEMPL_BINARY)
	@touch $(TEMPL_BINARY)

$(TAILWIND_BINARY): $(TEMP_DIR)
	curl -sLo $(TAILWIND_BINARY) $(TAILWIND_URL)
	@chmod +x $(TAILWIND_BINARY)

$(STYLES): $(TAILWIND_STYLES) $(TAILWIND_BINARY)
	$(TAILWIND_BINARY) build -i $(TAILWIND_STYLES) -o $(STYLES) --minify

#################################################################################
# Commands                                                                      #
#################################################################################

.PHONY: templates
## generates the templates using `templ generate`
templates: $(TEMPL_BINARY)
	@$(TEMPL_BINARY) generate

.PHONY: install
## `go generate` and `go install`
install: $(TEMPL_BINARY)
	@go generate $(ALL_PROJECTS)
	@go install $(ALL_PROJECTS)

.PHONY: build
## builds required deps for the site to run on netlify
build: install $(STYLES)
	@go run website/main.go

.PHONY: deploy-preview
## deploys the local changes as a preview
deploy-preview: build
	@netlify build
	@netlify deploy

.PHONY: dev
## start the dev server
dev: build
	@netlify build
	@netlify dev

.PHONY: test
## test the api endpoints and website page generator
test: install
	@go test $(ALL_PROJECTS)

.PHONY: lint
## lint all the modules
lint: install
	@golangci-lint run $(ALL_PROJECTS)

.PHONY: coverage
## test coverage across the code base
coverage: install
	@go test $(ALL_PROJECTS) -coverprofile=c.out
	@go tool cover -html="c.out"

.PHONY: clean
## removes binaries and artifacts
clean:
	@go clean
	@rm -rf $(TEMP_DIR)
	@rm -f $(STYLES)
	@find . -path "*/static/*.*ml" -type f -delete
	@find . -name "*_templ.go" -type f -delete

.PHONY: cleaner
## removes the same as clean and other ignored files
cleaner: clean
	@rm -rf .netlify
	@rm -f .env
	@rm -f c.out


#################################################################################
# Self Documenting Commands                                                     #
#################################################################################

.DEFAULT_GOAL := help

# Inspired by <http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html>
# sed script explained:
# /^##/:
# 	* save line in hold space
# 	* purge line
# 	* Loop:
# 		* append newline + line to hold space
# 		* go to next line
# 		* if line starts with doc comment, strip comment character off and loop
# 	* remove target prerequisites
# 	* append hold space (+ newline) to line
# 	* replace newline plus comments by `---`
# 	* print line
# Separate expressions are necessary because labels cannot be delimited by
# semicolon; see <http://stackoverflow.com/a/11799865/1968>
.PHONY: help
help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)"
	@echo
	@sed -n -e "/^## / { \
		h; \
		s/.*//; \
		:doc" \
		-e "H; \
		n; \
		s/^## //; \
		t doc" \
		-e "s/:.*//; \
		G; \
		s/\\n## /---/; \
		s/\\n/ /g; \
		p; \
	}" ${MAKEFILE_LIST} \
	| LC_ALL='C' sort --ignore-case \
	| awk -F '---' \
		-v ncol=$$(tput cols) \
		-v indent=19 \
		-v col_on="$$(tput setaf 6)" \
		-v col_off="$$(tput sgr0)" \
	'{ \
		printf "%s%*s%s ", col_on, -indent, $$1, col_off; \
		n = split($$2, words, " "); \
		line_length = ncol - indent; \
		for (i = 1; i <= n; i++) { \
			line_length -= length(words[i]) + 1; \
			if (line_length <= 0) { \
				line_length = ncol - indent - length(words[i]) - 1; \
				printf "\n%*s ", -indent, " "; \
			} \
			printf "%s ", words[i]; \
		} \
		printf "\n"; \
	}' \
	| more $(shell test $(shell uname) = Darwin && echo '--no-init --raw-control-chars')
