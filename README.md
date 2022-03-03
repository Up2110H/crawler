# crawler

### Зависимости
* Go v1.17
* github.com/gin-gonic/gin v1.7.7
* golang.org/x/net

### Установка/Запуск
```
go mod tidy
go run cmd/main.go
```

### Endpoint
```
post localhost:8000/api/titles
body:
{
  "urls": [
    "url1",
    "url2",
    ...
    "urlN"
  ]
}
```

### Тестирование
`go test ./..`
