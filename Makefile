NAME = noise-pollution
build:
	@go build -o ./$(NAME)

install:
	@go build -o /usr/bin/$(NAME)

run: build
	@./$(NAME)
