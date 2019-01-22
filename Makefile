.PHONY: all
all: \
	dep-ensure

include build/rules.mk
.PHONY: build/rules.mk
build/rules.mk:
	git submodule update --init --recursive build

.PHONY: dep-ensure
dep-ensure: $(DEP)
	$(DEP) ensure -v
