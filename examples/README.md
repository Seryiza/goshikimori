# Examples
Здесь хранятся различные примеры кода.

Вынес в отдельный пакет (а не как `example_test.go`) из-за периодичного запуска для получения конфигурации/токенов (используя [gorram](https://github.com/natefinch/gorram)).

Например:
```
$ gorram github.com/seryiza/go-shikimori/examples ExampleGetToken
$ gorram github.com/seryiza/goshikimori/examples PrintShikiGet "users/386084/info"
```