# API Шикимори для Golang
[![GoDoc](https://godoc.org/github.com/Seryiza/goshikimori?status.svg)](https://godoc.org/github.com/Seryiza/goshikimori) [![Go Report Card](https://goreportcard.com/badge/github.com/seryiza/goshikimori)](https://goreportcard.com/report/github.com/seryiza/goshikimori)

## Описание
Пакет предназначен для взаимодействия с [API Шикимори](https://shikimori.org/api/doc).

## Зависимости
* [github.com/golang/oauth2](https://github.com/golang/oauth2)
* [github.com/seryiza/loadOAuth](https://github.com/seryiza/loadOAuth) (пакет `helpers`)
* [github.com/headzoo/surf](https://github.com/headzoo/surf)

## Установка
```bash
go get github.com/seryiza/goshikimori
```

## Использование
### OAuth2
Прежде всего, для OAuth2 потребуется создать приложение [на самом Шикимори](https://shikimori.org/oauth/applications). Для авторизации потребуется *название приложения*, *client id* и *client secret*.

### Работа с API
Объект `goshikimori.Shikimori` предназначен для взаимодействия с API Шикимори. Можно использовать через:

* HTTP-запросы `Shikimori.Get/Post/...`:
```go
  // var shiki *goshikimori.Shikimori

  resp, _ := shiki.Get("users/whoami")    // для GET https://shikimori.org/api/users/whoami
  userJSON, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(userJSON))
  // => {"id":206253,"nickname":"Seryiza",...,"locale":"ru"}
```

* HTTP-запросы с приведением к структурам `Shikimori.JSONGet/JSONPost/...`:
```go
// var shiki *goshikimori.Shikimori

var im structs.User
err := shiki.JSONGet("users/whoami", &im)
if err != nil {
  panic(err)
}

fmt.Println(im, err)
// => models.User{
//      ID: 206253,
//      Nickname: "Seryiza",
//      ...
//    }
```

## Пример
Также есть примеры в папке `examples`.

```go
conf := &oauth2.Config{
  ClientID:     "your shikimori client id",
  ClientSecret: "your shikimori client secret",
  RedirectURL:  auth.StandaloneRedirectURL,
  Endpoint:     auth.ShikimoriEndpoint,
}

url := auth.GetAuthCodeURL(conf)
fmt.Println("Enter code from here: ", url)

var code string
if _, err := fmt.Scanln(&code); err != nil {
  panic(err)
}

ctx := context.Background()
ctx = goshikimori.AddTransport(ctx, "your shikimori oauth app name")

tok, err := conf.Exchange(ctx, code)
if err != nil {
  panic(err)
}

client := conf.Client(ctx, tok)
shiki := goshikimori.NewShikimori(client, "1.0")  // 1.0 -- version of Shikimori API

resp, err := shiki.Get("users/whoami")
if err != nil {
  panic(err)
}

user := &structs.User{}
jd := json.NewDecoder(resp.Body)
if err = jd.Decode(user); err != nil {
  panic(err)
}

fmt.Printf("I'm %s", user.Nickname)
```

Также есть вспомогательные функции `helpers` (использующие файлы и env-переменные) для написания меньшего кода:

```go
// GetShikimori получает oauth2 конфиг и токен из файлов или env-переменных
// и создает Shikimori из стандартного http клиента
shiki, err := helpers.GetShikimori("1.0")
if err != nil {
  panic(err)
}
// Сохраняем токен в файл, если он изменился
defer helpers.SaveToken(shiki)

user := &structs.User{}
_, err = shiki.JSONGet("users/whoami", user)
if err != nil {
  panic(err)
}

fmt.Printf("I'm %s", user.Nickname)
```

## Тестирование
Тесты также проверяют запрос-ответ от Шикимори. Для корректной работы всех тестов необходимо задать следующие envirement-переменные:
* `SHIKI_APP_NAME` -- название OAuth-приложения на Шикимори
* `SHIKI_CLIENTID` -- публичный Client ID приложения на Шикимори
* `SHIKI_CLIENTSECRET` -- секретный ключ приложения на Шикимори
* `SHIKI_REDIRECT_URL` -- url для перенаправления приложения (на Шикимори)
* `SHIKI_TOKEN_FILE` -- путь к файлу с json-токеном
* `SHIKI_CONF_FILE` -- путь к файлу с json oauth-конфигурацией
* `SHIKI_LOGIN` -- псевдоним пользователя
* `SHIKI_PASS` -- пароль пользователя (используется при получении токена, если предыдущий был истекшим)
