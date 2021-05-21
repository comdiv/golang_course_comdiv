# Занятие 004 2021-05-18

1. [Stepic: 2.3 Указатели](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQ0NDU3MzcyMTE4/details)
2. [Unit testing & benchmarking](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/a/MzQwODQ3MjkyMDky/details)
3. [Материал к структурированию пакетов](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/m/MzQwODQ3MjkyMDA2/details)
4. [Функции и пакеты. Тезисы к обсуждению](https://classroom.google.com/u/0/c/MzM5NDA2NTc2ODk5/m/MzQwODQ3MjkyMTA2/details)

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