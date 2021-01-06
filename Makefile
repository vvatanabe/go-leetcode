.PHONY: dir
dir:
	SNAKE=`echo $(NAME) | sed s/-/_/g)`; \
	mkdir -p $$SNAKE; \
	echo "package main" > $$SNAKE/main.go; \
	echo "package main" > $$SNAKE/main_test.go;

.PHONY: test
test:
	go test -v -coverprofile= ./...

