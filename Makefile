.PHONY: build
build:
	docker build -f dockerfiles/Dockerfile -t article .
	docker-compose -f dockerfiles/docker-compose.yml up --build

.PHONY: clean
clean:
	docker-compose -f dockerfiles/docker-compose.yml down
