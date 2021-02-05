.PHONY: dir
init:
	git checkout main; \
	git pull; \
	git checkout -b $(NAME)
	SNAKE=`echo $(NAME) | sed s/-/_/g)`; \
	mkdir -p $$SNAKE; \
	echo "package main" > $$SNAKE/main.go; \
	echo "package main" > $$SNAKE/main_test.go; \
	echo [](./$$SNAKE/main.go) >> README.md; \
	goland ./$$SNAKE/main.go

.PHONY: test
test:
	go test -v -coverprofile= ./...

