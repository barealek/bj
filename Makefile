build:
	go build -ldflags "-s -w" -o bin/bj.exe cmd/cli/main.go
	# upx bin/bj.exe

clean:
	rm -rf bin/
