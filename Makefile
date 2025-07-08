.PHONY: reset_tags rgo

reset_tags: ##Reset GH tags
	git tag -l | xargs git tag -d

rgo: ##Run go
	go run *.go