# Очистка СУБД Postgres от временных тестовых СУБД

## Назначение

При выполнении тестов на Postgres каждый тест
порождает СУБД с именем `test_<testname>` и естественно их формируется
достаточно много.

Данный инструмент позволяет массово устранить все эти временные СУБД с целевого сервера.

Требуется SU доступ к СУБД

## Общий синтаксис запуска

1. Если не собирать в exe `go run cmd/cleantestdb/cleantestdb.go <args>`
2. Если собрать `go build cmd/cleantestdb/cleantestdb.go`, то появится
бинарник `cleantestdb.exe` на Windows или `cleantestdb` на Linux и
можно будет просто `./cleantestdb <args>`
3. Также можно загрузить директорию с утилитой как проект GoLand и запускать оттуда 
   
## Аргументы

Если выполнить `./cleantestdb -h` или `go run cmd/cleantestdb/cleantestdb.go -h` будет
показана следующая справочная информация:

```
Usage of ...\cleantestdb :
  -host string
        PG host (default "127.0.0.1")
  -listonly
        Only list no drop
  -pass string
        PG password (default "postgres")
  -port int
        PG port (default 5436)
  -prefix string
        Prefix for databases to clean (default "test_")
  -sysdb string
        PG sys table nasme (default "postgres")
  -user string
        PG username (default "postgres")

```

По умолчанию утилита нацелена на локальный тестовый PG на порту 5436.

Особо важные параметры:

1. `-listonly` - только выведет список баз, подлежащих удалению, но удалять не будет
2. `-prefix=PREFIX` - позволяет указать или полное имя СУБД или собственный префикс
имени, отличный от `test_`
   
> Внимание! Для защиты системных СУБД утилита не позволяет выставлять пустой
> префикс или префикс без символа подчеркивания!
