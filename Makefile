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
# Commands                                                                      #
#################################################################################

BUILD_FILE=$(PWD)/public/tailwind.css
CSS_FILE=$(PWD)/public/styles.css
BINARY=$(PWD)/tailwindcss-$(OS)-$(ARCH)
BINARY_URL=https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(OS)-$(ARCH)

.PHONY: build
## generates `tailwind.css` from the `styles.css` using tailwindcss binary
build: $(BUILD_FILE)
$(BUILD_FILE): $(CSS_FILE) $(BINARY)
	$(BINARY) build -i $(CSS_FILE) -o $(BUILD_FILE) --minify

.PHONY: install
## installs the latest version of the `tailwindcss` CLI
install: $(BINARY)
$(BINARY):
	curl -sLo $(BINARY) $(BINARY_URL)
	chmod +x $(BINARY)

.PHONY: dev
## start the dev server
dev: $(BUILD_FILE)
	go run main.go

.PHONY: test
## test the api endpoints
test:
	go test ./...

.PHONY: clean
## removes binaries and artifacts
clean:
	rm -vf $(BINARY) $(BUILD_FILE)


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
