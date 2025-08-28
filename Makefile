.PHONY: update_modules
update_modules:
	@go get -u all
	@go mod tidy
	@go mod vendor


.PHONY: update_internal
update_internal:
	@go get -u github.com/becash/apis
	@go mod tidy
	@go mod vendor