#!/bin/bash

set -euo pipefail

if [[ -n $(git status --porcelain) ]]; then
	echo "Staging area is dirty, please add all files created by the build to .gitignore"
	git diff --color-words
	git status -s
	exit 1
fi
