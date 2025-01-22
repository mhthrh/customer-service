set-app-path:
	sudo chmod +x ./shell/set.sh
	zsh ./shell/set.sh
	#sh ./shell/set.sh
update-lib:
	go get github.com/mhthrh/common-lib #get latest commit instead of "go get github.com/mhthrh/common-lib@latest-hashcode"
test:
	go test ./... -v -bench . -failfast -cover -count=1