build:
	@echo Building...
	@go build -v .
test:
	@echo Testing...
	@go test -v ./...
coverage:
	@echo "" > coverage.txt;
	@for d in $(shell go list -v ./...); do \
		go test -race -coverprofile=profile.out -covermode=atomic $$d; \
		if [ -f profile.out ]; then \
			cat profile.out >> coverage.txt; \
			rm profile.out; \
		fi; \
	done
clean:
	rm coverage.txt
