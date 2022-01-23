.PHONY: default
default: test

clean:
	rm -rf ./dist
	rm -f coverage.out

test: clean
	go test -short -coverprofile=coverage.out `go list ./...`
	go tool cover -func=coverage.out

test-ci: clean
	go test `go list ./...` -coverprofile=coverage.out -json > test-results.json
	go tool cover -func=coverage.out

test-report: test
	go tool cover -html=coverage.out

test-package: test
	goreleaser release --skip-sign --skip-validate --skip-publish --rm-dist --snapshot

build: test
	goreleaser build --single-target --rm-dist --snapshot

tag: test
	@export VERSION_TAG=`cat VERSION` \
	&& git tag v$$VERSION_TAG \
	&& git push origin v$$VERSION_TAG

release: tag
	goreleaser release --rm-dist