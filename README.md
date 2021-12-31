
# Setup
* Add `gopherjs` binary to path: `export GOPHERJS_GOROOT="$(go1.17.1 env GOROOT)"  # Also add this line to your .profile or equivalent.`

# Usage
* build: `GOARCH=wasm GOOS=js go build -o dist/main.wasm main.go`
* start: `go run serve.go`
* format: `go fmt`
