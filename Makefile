build:
	env GOOS=linux GOARCH=arm GOARM=7 go build main.go

clean:
	rm -rf main

test:
	env GOOS=linux GOARCH=arm GOARM=7 go build main.go
	adb shell "rm -rf /data/local/tmp/main"
	adb push main /data/local/tmp
	adb shell "chmod +x /data/local/tmp/main"
	adb shell "/data/local/tmp/main"
