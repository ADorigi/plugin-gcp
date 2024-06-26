# run plugin in debug mode
# (Assuming Kaytu CLI is running in debug mode AND Kaytu CLI has plugin-gcp installed)
debug:
	go run main.go --server localhost:30422

# removes the dist folder, which is generated by goreleaser build(make goreleaser)
clean:
	rm -rf dist

# cleans the dist directory and generates build using the goreleaser's configuration for current GOOS and GOARCH
goreleaser: clean
	REPOSITORY_NAME="kaytu" REPOSITORY_OWNER="kaytu-io" goreleaser build --snapshot --single-target

# test components of plugin/gcp
testgcp:
	go test -v ./plugin/gcp -count=1