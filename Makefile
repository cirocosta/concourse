all: ci/build-branches.yml

pipeline: ci/build-branches.yml
	fly -t hh set-pipeline -p build-branches -c ./ci/build-branches.yml

ci/build-branches.yml: branches-list.json build-branches.jsonnet
	jsonnet \
		--ext-code 'branches=$(shell cat ./branches-list.json)' \
		./build-branches.jsonnet > $@

branches-list.json:
	git ls-remote --heads origin | \
		awk '{ print $$2 }' | \
		awk -F "/" '{ for (i=3; i<NF; i++) printf $$i "/"; print $$NF }' | \
		grep -v 'maintenance' | \
		jq -RscM '. / "\n" - [""]' > $@

.PHONY: branches-list.json
