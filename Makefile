.PHONY: dir
init:
	git checkout main; \
	git pull; \
	git checkout -b $(NAME)
	SNAKE=`echo $(NAME) | sed s/-/_/g)`; \
	mkdir -p problemset/$$SNAKE; \
	echo "package main" > problemset/$$SNAKE/main.go; \
	echo "package main" > problemset/$$SNAKE/main_test.go; \
	echo "- [Title](./problemset/$$SNAKE/main.go)" >> README.md;

.PHONY: test
test:
	go test -v -coverprofile= ./...

