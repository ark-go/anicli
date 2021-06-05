.PHONY: 
space:
	$(info eeeee)
build:
	go build -o ./bin/main/main cmd/app/main/main.go
	$(info Компиляция)
buildwin:
	GOOS=windows GOARCH=amd64 go build -o ./bin/main/winmain.exe cmd/app/main/main.go
#.SILENT: 
run: build
	$(info Запуск)
	./bin/main/main   copy -r r234 -w ./e help -help -mm  reverse  addPath -p 44

gittag:
	git tag v2.0.1
	git push origin --tags

help:
	go doc -all ./internal
	go doc -all ./cmd/app/main
	go doc -all ./pkg/structs
	