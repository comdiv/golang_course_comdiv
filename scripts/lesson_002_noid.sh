#/bin/bash

# проверим, что настроен GO
if [[ "$GOPATH" == "" ]]; then
	echo "GOPATH not set"
	exit 1
fi;

# упростим использование GO
PATH=$GOPATH:$PATH


# точный путь до текущего скрипта для гарантии корректного размещения результатов
CWD="$(cd -P -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)"

# целевая директория для файла проекта
PROJECTDIR=$CWD/../tmp/lesson_002_noid  

# гарантировано удаляем директорию проекта, чтобы было убедительнее что мы все делаем чисто
rm -rf $PROJECTDIR

# создаем пустую директорию проекта
mkdir -p $PROJECTDIR

# переходим в директорию проекта
cd  $PROJECTDIR

# пишем в проект наш исходный файл на GO
echo '    
package main

import "fmt"

func main() {
  fmt.Println("Hello, Go!")
}

' > hello.go

# выполняем и ловим результат в переменную
RESULT=$(go run hello.go)

# выводим полученный результат в консоль
echo "Result is: $RESULT"

# проверяем что программа действительно выполнилась корректно
if [[ "$RESULT" == "Hello, Go!" ]]; then
	echo "All is Ok!"
else 
	echo "Something going wrong should be 'Hello, Go!'"
	exit 2
fi;
