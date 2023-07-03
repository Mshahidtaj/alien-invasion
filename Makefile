# Variables
APP_CONTAINER_NAME = mshahidtaj/alien-invasion
TAG?=latest

.PHONY: parse

help:
	@echo "************************** Alien Invasion Game ******************************* "
	@echo "1. build                  - Build the game into single binary"
	@echo "2. test                   - Run the test coverage for Alien Invasion"
	@echo "3. clean                  - Remove the alien inavsion binary file"
	@echo "4. run                    - Run the alien inavsion"
	@echo "5. build-alien-invasion   - Build the docker image for alien inavsion"
	@echo "6. push-alien-invasion    - Push the alien inavsion image to docker"

build:
	go build -o alien-invasion

test:
	go test -v ./...

clean:
	rm -f alien-invasion

run:
	./alien-invasion $(if $(worldmap),-worldmap=$(worldmap)) $(if $(n),-n=$(n))
	

build-alien-invasion:
	docker buildx build --tag $(APP_CONTAINER_NAME):$(TAG) .

push-alien-invasion:
	docker push $(APP_CONTAINER_NAME):$(TAG)