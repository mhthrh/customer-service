IMAGE_NAME=customer-service
set-app-path-local:
	sudo chmod +x ./script/set.sh
	zsh ./script/set.sh
	#sh ./script/set.sh
set-app-path-docker:
	chmod +x ./script/docker-set.sh
	 ./script/docker-set.sh
update-go-nest:
	sudo chmod +x ./script/update-lib.sh
	zsh ./script/update-lib.sh
	#sh ./script/update-lib.sh
test_app:
	go test ./test/... -v -bench . -failfast -cover -count=1
build:
	docker build --progress=plain -t $(IMAGE_NAME) .
run: build
	docker run --rm -p 6985:6985 $(IMAGE_NAME)

.PHONY: build run test_app update-go-nest set-app-path set-app-path-docker

