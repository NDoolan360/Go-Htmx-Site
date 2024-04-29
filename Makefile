# Export all variables
export

# Set up environment variables
ENVFILE=$(PWD)/.env
ENVFILE_EXAMPLE=$(PWD)/.example.env
$(shell cp -n $(ENVFILE_EXAMPLE) $(ENVFILE))
include $(ENVFILE)

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

BIN_DIR=$(PWD)/bin
FUNC_DIR=$(PWD)/functions
GO_BINARY=$(FUNC_DIR)/go-htmx-site
TAILWIND_BINARY=$(BIN_DIR)/tailwindcss-$(OS)-$(ARCH)
TAILWIND_URL=https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(OS)-$(ARCH)
TAILWIND_STYLES=$(PWD)/static/styles/tailwind.css
STYLES=$(PWD)/static/styles/styles.css

$(TAILWIND_BINARY):
	mkdir -p $(BIN_DIR)
	curl -sLo $(TAILWIND_BINARY) $(TAILWIND_URL)
	chmod +x $(TAILWIND_BINARY)

$(STYLES): $(TAILWIND_STYLES) $(TAILWIND_BINARY)
	$(TAILWIND_BINARY) build -i $(TAILWIND_STYLES) -o $(STYLES) --minify


#################################################################################
# Commands                                                                      #
#################################################################################

.PHONY: templates
## generates the templates using `templ generate`
templates:
	templ generate

.PHONY: install
## installs the latest version of the `tailwindcss` CLI and the go packages named
## by the import paths
install: templates $(TAILWIND_BINARY)
	go get ./...
	go install ./...

.PHONY: build
## build the go binaries into the function folder
build: install $(STYLES)
	mkdir -p $(FUNC_DIR)
	go build -o $(FUNC_DIR) ./...

.PHONY: build-linux
## build the go binaries for linux into the function folder
build-linux: install $(STYLES)
	mkdir -p $(FUNC_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(FUNC_DIR) ./...

.PHONY: dev
## start the dev server
dev: templates $(STYLES)
	netlify dev

.PHONY: test
## test the api endpoints
test: templates
	@go test ./... -v

.PHONY: clean
## removes binaries and artifacts
clean:
	@go clean
	@rm -rvf $(BIN_DIR) $(FUNC_DIR) $(STYLES)
	@find . -name "*_templ.go" -type f -delete


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
