export

MODULES = ./web ./api/health ./api/projects ./api/experience
TEMPLATES = ./web/templates ./api/projects ./api/experience
MAKE_HTML_MODULE = ./web
MAKE_HTML_BIN = make_html
STATIC_DIR = web/static
ENVFILE = ./.env
DEV_PORTS = 3999 8888

include $(ENVFILE)

.PHONY: generate install deploy test dev clean

.env:
	cp -n .example.env .env
	echo "Make sure the .env file has the correct values"

go.work:
	go work init
	go work use $(MODULES) $(TEMPLATES)

generate: go.work
ifneq (, $(shell which templ))
	for template in $(TEMPLATES); do templ generate -path $$template; done
else
	for template in $(TEMPLATES); do go run github.com/a-h/templ/cmd/templ@latest generate -path $$template; done
endif

install: go.work generate
	go install $(MODULES) $(TEMPLATES)

dist: install
	rm -rf $@
	cp -R $(STATIC_DIR)/. $@
ifneq (, $(shell which minify))
	minify -o $@/styles.css -b web/styles/*.css
else
	go run github.com/tdewolff/minify/v2/cmd/minify@latest -o $@/styles.css -b web/styles/*.css
endif
	go build -o $(MAKE_HTML_BIN) $(MAKE_HTML_MODULE)
	cd $@; ../$(MAKE_HTML_BIN);

deploy: dist
	netlify deploy

test: install
	go test $(MODULES)

dev: dist
	parallel -tmux ::: 'watchexec -e go,templ,md,css,js,avif,ico "make dist"' 'netlify dev --target-port 8888'

clean:
	$(foreach port, $(DEV_PORTS), lsof -i:$(port) -t | xargs kill;)
	mv .env .env.bak
	git clean -fdX
	mv .env.bak .env
