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
	./bin/main/main -dddd  -go -qwerty fff dd help8 copy -r r234 -w ./e -qwerty2 -vvv "привет" -vvv2 -0033 -qwerty2 -BBB -lolololol uuu -k 678/876 MMM -ku 888 -c kkkk TEST -c uuu -c "tt  hjh h dfh dfhdfh dfh dfh kjh kjt" -g hhh -j gjkjgkjhgkj -i gkjhgkjhgkjh kkk -b --cccc -bbbb eee  fff  -rrrr AAA BBB -lolololol CCC -lolololol kkkk -g ooooo -g yyy TEST888 -k -c 55 -c 99

gittag:
	git tag v2.0.0
	git push origin --tags

help:
	go doc -all ./internal
	go doc -all ./cmd/app/main
	go doc -all ./pkg/structs
	