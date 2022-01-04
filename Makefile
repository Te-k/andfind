BUILD_FOLDER  = $(shell pwd)/build

.PHONY: build
default: build ;


fmt:
	gofmt -w -s ./


lint:
	@echo "[lint] Running linter on codebase"
	@golint ./...


build:
	env GOOS=linux GOARCH=arm GOARM=7 go build -o $(BUILD_FOLDER)/andfind_arm7
	env GOOS=linux GOARCH=arm GOARM=6 go build -o $(BUILD_FOLDER)/andfind_arm6
	env GOOS=linux GOARCH=amd64 go build -o $(BUILD_FOLDER)/andfind_x64

clean:
	rm -rf $BUILD_FOLDER

test: build
	adb shell "rm -rf /data/local/tmp/andfind_arm7"
	adb push $(BUILD_FOLDER)/andfind_arm7 /data/local/tmp
	adb shell "chmod +x /data/local/tmp/andfind_arm7"
	adb shell "/data/local/tmp/andfind_arm7"
