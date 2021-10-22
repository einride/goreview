PRETTIER_VERSION := 2.2.1
PRETTIER_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))/bin/$(PRETTIER_VERSION)
PRETTIER := $(PRETTIER_DIR)/node_modules/.bin/prettier

$(PRETTIER):
	npm install --no-save --no-audit --prefix $(PRETTIER_DIR) prettier@$(PRETTIER_VERSION)
	chmod +x $@
	touch $@

.PHONY: yaml-format
yaml-format: $(PRETTIER)
	$(PRETTIER) --parser yaml --write ./**/*.y*ml *.y*ml

.PHONY: markdown-format
markdown-format: $(PRETTIER)
	$(PRETTIER) --parser markdown --check ./**/*.md --check *.md -w
