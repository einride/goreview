semantic_release_dir := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))
semantic_release_version := 17.1.1
semantic_release := $(semantic_release_dir)/node_modules/.bin/semantic-release

$(semantic_release):
	npm install --no-save --no-audit --prefix $(semantic_release_dir) semantic-release@$(semantic_release_version)
	npm install --no-save --no-audit --prefix $(semantic_release_dir) @semantic-release/changelog
	npm install --no-save --no-audit --prefix $(semantic_release_dir) conventional-changelog-conventionalcommits@4.4.0

.PHONY: semantic-release
semantic-release: $(semantic_release)
	$(semantic_release)
