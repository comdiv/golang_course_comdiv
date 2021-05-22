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

Потом добавил более крайние варанты агрессивные по данным

1. 10000 уникальных по возрастанию
2. 10000 уникальных по убыванию
3. 10000 с большим количеством частых дубляжей (по % 20)

Потом сделал оптимизацию слайсевого с использованием дополнительной карты `map[int]int` для
дубликатов (заполняется только для дубляжей). Это увеличело потребление памяти но при этом резко ускорило в
целом решение на примитиах golang  (я так понял карты не являлись нарушением условия задачи).


Теперь такие бенчмарки:
```
BenchmarkSortedLinkedList_InsertRandom-12                     27          41291074 ns/op          137024 B/op       4282 allocs/op
BenchmarkSortedLinkedList_InsertAscNoDups-12                2726            430632 ns/op          320034 B/op      10001 allocs/op
BenchmarkSortedLinkedList_InsertDescNoDups-12               2854            419276 ns/op          320033 B/op      10001 allocs/op
BenchmarkSortedLinkedList_InsertManyDups-12                 5078            227701 ns/op             672 B/op         21 allocs/op
BenchmarkSortedLinkedList_Delete-12                           64          18923483 ns/op               0 B/op          0 allocs/op
BenchmarkSortedLinkedList_GetAll-12                        25339             47893 ns/op           81920 B/op          1 allocs/op
BenchmarkSortedLinkedList_GetUnique-12                     25089             47189 ns/op           81920 B/op          1 allocs/op
BenchmarkLinkedFind-12                                     88675             13585 ns/op               0 B/op          0 allocs/op
BenchmarkSortedSliced_InsertRandom-12                        127          15893527 ns/op        80170642 B/op       4257 allocs/op
BenchmarkSortedSliced_InsertAscNoDups-12                    8110            143861 ns/op          386394 B/op         22 allocs/op
BenchmarkSortedSliced_InsertDescNoDups-12                     16          69593181 ns/op        428602503 B/op     20023 allocs/op
BenchmarkSortedSliced_InsertManyDups-12                     3602            337573 ns/op            1689 B/op         12 allocs/op
BenchmarkSortedSliced_Delete-12                              795           1502937 ns/op               3 B/op          0 allocs/op
BenchmarkSortedSliced_GetAll-12                             5365            221075 ns/op           81920 B/op          1 allocs/op
BenchmarkSortedSliced_GetUnique-12                          5444            221063 ns/op           81920 B/op          1 allocs/op
BenchmarkLastIndexOf_5_sorted-12                        220747221                5.434 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_5_non_sorted-12                    278496711                4.311 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_10_sorted-12                       162669978                7.331 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_10_non_sorted-12                   189784501                6.326 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_20_sorted-12                       133960431                8.962 ns/op           0 B/op          0 allocs/op
BenchmarkLastIndexOf_20_non_sorted-12                   100000000               10.05 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_100_sorted-12                      97160484                12.53 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_100_non_sorted-12                  24504898                48.49 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_1000_sorted-12                     69303270                17.56 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_1000_non_sorted-12                  4117498               290.3 ns/op             0 B/op          0 allocs/op
BenchmarkLastIndexOf_10000_sorted-12                    38682224                30.90 ns/op            0 B/op          0 allocs/op
BenchmarkLastIndexOf_10000_non_sorted-12                  429424              2798 ns/op               0 B/op          0 allocs/op
```

Вот более ясная сравнительная таблица по ключевым пробам

### Ключевые пробы в разрезе производительности (ns/op)

|Проба|Связанный список|Только срез|Срез+(Карта)|Лучший|Худший|
|-----|----------------:|-----------:|----------:|-----------|------|
|Вставка рандомных значений|41291074|65148495|15893527|Срез+|Срез|
|Вставка по возрастанию|430632|141894|143861|Срез,Срез+|Список|
|Вставка по убыванию|419276|70529293|69593181|Список|Срез,Срез+|
|Вставка многих дубляжей|230748|5578888|337573|Список|Срез|
|Удаление|18906929|2336147|1502937|Срез+|Список|
|Возврат всех|47891|15392|221075|Срез|Срез+|
|Возврат уникальных|20535|18338|7617|Срез+|Список| 
|Итого баллов|6|6|9|Срез+|Список==Срез|

### Ключевые пробы в разрезе памяти

|Проба|Связанный список|Только срез|Срез+(Карта)|Лучший|Худший|
|-----|----------------:|-----------:|----------:|-----------|------|
|Вставка рандомных значений|137024|373221826|80171115|Список|Срез|
|Вставка по возрастанию|320033|386330|386394|-|-|
|Вставка по убыванию|320034|428602295|428602581|Список|Срез,Срез+|
|Вставка многих дубляжей|672|3698144|1694|Список|Срез|
|Итого баллов|6|0|2|Список|Срез|

Итого расставим по местам:

1. Список - 12 баллов по сумме производительности и памяти - 1 место (за счет памяти)
2. Срез+карта - 11 баллов - на втором месте из-за памяти, но при этом производительнее    
2. Просто срез - 6 баллов - намного хуже чем первые 2 решения

Вывод. Если данных мало и они меряются десятками на массив в целом достаточно просто среза,
больших отличий не будет.
Но если нужно прокачать действительно большой набор значений, то потребуется или связанный список
или срез+карта.

Прогнал еще раз на большом объеме 100000 вместо обычных для этих моих бенчей 10000.
Проивзолительность списка просела на запись (деградировала), а у срез+карты нет - и 
понятно - у меня в списке линейный поиск позиции, а в срезе - бинарный.
Производительность срез+карта на порядок лучше.
Но при этом срез+карта сожрал очень много памяти - потребление памяти на порядок хуже.

В общем если память не критична - лучше срез+карта.
В любом случае просто срез - показал себя так себе.


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