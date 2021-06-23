# Старт в режиме HTTP

Пример запуска

`go run main.go --first --last --http 8080 --pprofhttp same`

также уже и добавлена команда для make на такой вызов:

`go ta_http`

Также добавил [openapi.json](openapi.json) файл и разрешил CORS,
можно в принципе и с [https://editor.swagger.io/](https://editor.swagger.io/) тестировать,

также приложил файл для [insomnia](insomnia.json) с набором вызовов