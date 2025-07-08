.PHONY: reset_tags rgo rgt

reset_tags: ##Reset GH tags
	git tag -l | xargs git tag -d

rgo: ##Run go
	go run *.go

rgt: ##Run go test
	cd logic && go test