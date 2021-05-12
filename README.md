[![Go Report Card](https://goreportcard.com/badge/github.com/comdiv/golang_course_comdiv)](https://goreportcard.com/report/github.com/comdiv/golang_course_comdiv)

# golang_course_comdiv
Репозиторий для изучения golang Садыков Фагим

## Работы

### Занятие 002 2021-05-12

#### 002.02 Исполнить `hello, go!` без использования IDE

Для облегчения сдачи и контроля выполнения исполнил в виде [Bash скрипта](./scripts/lesson_002_noid.sh)

Внутри происходит инициализация пустой директории, создание файла .go и его выполнение 
с контролем корректности итогов выполнения.

Файлы GO гарантировано порождаются во временной директории `./tmp/lesson_002_noid` не зависимо от 
точки файловой системы откуда запущен скриптю

Для запуска

1. Если в системе настроен GOPATH, то просто `./scripts/lesson_002_noid.sh`
2. Если нет, то соответственно `GOPATH=<путь до go/bin> ./scripts/lesson_002_noid.sh`

В результате должно быть напечатано следующее:

```
Out from go run
Hello, Go!
Using just created lesson_002_noid(.exe)
Result is: Hello, Go!
All is Ok!
```

### Занятие 001 2021-05-11

[сслка на задание](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzM5NDA2NTc2OTY0/details)

#### 001.01 Структура проекта

Создан пустой репозиторий.
В репозиторий сразу добавлен

1. [README.md](README.md)
2. [.gitignore](.gitignore)
3. [Makefile](Makefile) - заготовка
4. [go.mod](go.mod) - заготовка


##### 001.02 Первая программа
В соответствии с [заданием](https://stepik.org/lesson/228260/step/1?unit=200793), 
добавлено начальное [`Hello, go` приложение](cmd/lesson_001_hellogo.go).

Запуск программы добавлен в [Makefile](Makefile)

1. `make run_lesson_001_hellogo` - запустит только этот пример
2. `make runall` - для запуска в составе всех примеров


#### 001.03  "локальное окружение"

1. Windows 10
2. GoLang 1.16.4
3. GoLand

Файлы для GoLand проверены на отсутствие каких-то локальных ссылок и прочих непереносимых сведений
и включены в проект.
