all: ci/build-branches.yml

ci/build-branches.yml: branches-list.json build-branches.jsonnet
	jsonnet \
		--ext-code 'branches=$(shell cat ./branches-list.json)' \
		./build-branches.jsonnet > $@

branches-list.json:
	git ls-remote --heads origin | \
		awk '{ print $$2 }' | \
		awk -F '/' '{ print $$3 }' | \
		grep -v 'maintenance' | \
		jq -RscM '. / "\n" - [""]' > $@

.PHONY: branches-list.json