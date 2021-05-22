# Занятие 004 2021-05-18

1. [Stepic: 2.3 Указатели](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQ0NDU3MzcyMTE4/details)
2. [Unit testing & benchmarking](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQwODQ3MjkyMDky/details)
3. [Материал к структурированию пакетов](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/m/MzQwODQ3MjkyMDA2/details)
4. [Функции и пакеты. Тезисы к обсуждению](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/m/MzQwODQ3MjkyMTA2/details)

## Обновление

Переделал SortedList

1. Ну собственно все прорефакторил в части организации кода, теперь он кроме [main](../cmd/lesson_005/sortedintlist/main.go)
   весь собран [тут](../internal/sortedintlist)
2. Сделал обе реализации и на самодельном LinkedList и на Slice
3. Так как реализации одного и того же 2, то я сделал:
    1. интерфейс который они оба реализуют и REPL может с любым работать
    2. шаблонные модельные тесты для этого интерфейса и просто вызовы их как для одной так и другой реализации
    3. шаблонные бенчмарки также для обоих реализаций
4. Для слайсов мне еще понадобился эффективный для соортированных списко LastIndexOf - тоже написал +
   в нем есть оптимизация на бинарный поиск для больших сортированных слайсов, проверено соответствующими
   бенчмарками что прирост действительно существенный
   
Все тесты прошли, бенчмарки собраны.

Вот на моей машине:

```
goos: windows
goarch: amd64
pkg: github.com/comdiv/golang_course_comdiv/internal/sortedintlist/test
cpu: AMD Ryzen 5 2600 Six-Core Processor
BenchmarkSortedLinkedList_Insert-12                   27          41504422 ns/op          137024 B/op       4282 allocs/op
BenchmarkSortedLinkedList_Delete-12                   62          19290894 ns/op               0 B/op          0 allocs/op
BenchmarkSortedLinkedList_GetAll-12                23646             50648 ns/op           81920 B/op          1 allocs/op
BenchmarkSortedLinkedList_GetUnique-12             24657             48550 ns/op           81920 B/op          1 allocs/op
BenchmarkLinkedFind-12                             88176             13699 ns/op               0 B/op          0 allocs/op
BenchmarkSortedSliced_Insert-12                       19          69020447 ns/op        373221741 B/op      8676 allocs/op
BenchmarkSortedSliced_Delete-12                      525           2325148 ns/op               1 B/op          0 allocs/op
BenchmarkSortedSliced_GetAll-12                    83980             14687 ns/op           81920 B/op          1 allocs/op
BenchmarkSortedSliced_GetUnique-12                 82430             14157 ns/op           81920 B/op          1 allocs/op
BenchmarkLastIndexOf_5_sorted-12                227660871                5.327 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_5_non_sorted-12            279002871                4.319 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_10_sorted-12               160709370                7.390 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_10_non_sorted-12           188338281                6.359 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_20_sorted-12               131622720                8.925 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_20_non_sorted-12           100000000               10.17 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_100_sorted-12              94085916                12.75 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_100_non_sorted-12          25532023                45.40 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_1000_sorted-12             63152244                18.82 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_1000_non_sorted-12          4128564               287.0 ns/op             0 B/op          0 allocs/op
BenchmarkLastIndexOf_10000_sorted-12            37481844                32.37 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_10000_non_sorted-12          444460              2787 ns/op               0 B/op          0 allocs/op
PASS
ok      github.com/comdiv/golang_course_comdiv/internal/sortedintlist/test      77.350s

```

Выводы

1. SortedIntLinkedList в общем случае чуть быстрее на вставку и меньше потребляет памяти (так как там оптимизация для дубликатов)
2. Но в остальных пробах - удаление, чтение, поиск - лучше получилось на слайсах

## Тайминг

1. `(00:15)(00:15)` Реструктурировал документацию, подготовливал диретории и [README](../README.md) для задания
1. `(00:22)(00:37)` [Stepic: 2.3 Указатели](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQ0NDU3MzcyMTE4/details) 
и чтение [статьи](https://habr.com/en/post/339192/)
1. `(04:30)(05:00)`  [Unit testing & benchmarking](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQwODQ3MjkyMDky/details)
> ушло реально много времени, но решил прямо нормально запилить какие-то свои структуры данных, методы
> работу с файлами, тесты, бенчмарки и все такое

## [Stepic: 2.3 Указатели](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQ0NDU3MzcyMTE4/details)

Задания выполнил : [функция обмена по указателям](../cmd/lesson_005/lesson_005_2_3__6.go) и
[тест к ней](../cmd/lesson_005/lesson_005_2_3__6_test.go)

Никаких особых вопросов к указателям не возникло - обычные указатели, только без арифметики.

##  [Unit testing & benchmarking](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQwODQ3MjkyMDky/details)

1. Все же решил после чтения [этого](https://softwareengineering.stackexchange.com/questions/286406/use-of-this-in-golang#:~:text=Don't%20use%20generic%20names,and%20serves%20no%20documentary%20purpose.),
что я буду как и автор камента использовать `this` для методов в том же файле, что и структура к которой они относятся
   
В общем сделал консоль, которая динамически читает команды.

Запуск `go run github.com/comdiv/golang_course_comdiv/cmd/lesson_005/InsertDeleteSortedArray`

Сначала выведет такой help:

```
any positive int ( 10 )  - add it to list
any negative int ( -10)  - remove it counterpart from list
size  - prints list size (unique value count)
count - prints list size (all value count)
all - prints all values (with duplicates)
unique - prints only unique values
```

Код содержится в [этой директории](../cmd/lesson_005/sortedintlist)

Тестами более менее охвачена даже консоль (подмена Stdin, Stdout файлами)

Решил не делать на массивах и слайсах, слайс - только производная.

Сделал как классический 2-направленный LinkedList, который сразу формируется как сортированный.

Также сразу сделал так, чтобы он работал и как SortedList и как SortedSet - при этом объем 
памяти не зависит от количества дубликатов (собственно это особенность упорядоченных списков).

Можно тестировать и при помощи make `make test_sorted_linked_list`