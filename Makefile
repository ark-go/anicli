.PHONY: 
space:
	$(info eeeee)
build:
	go build -o ./bin/main/main cmd/app/main/main.go
	$(info Компиляция)

.SILENT: 
run: build
	$(info Запуск)
	./bin/main/main  BBB -lolololol uuu -k 678/876 888 -c kkkk -c uuu -bbbb hhh -j gjkjgkjhgkj -i gkjhgkjhgkjh kkk -b --cccc -bbbb eee  fff  -rrrr AAA BBB -lolololol CCC kkkk -g ooooo -g yyy

gittag:
	git tag v1.0.4
	git push origin --tags
# git tag -d v1.0.2  - удаляем тег в локальном репозитарии	
# git push origin --delete v1.0.2  - удалить тег на сервере