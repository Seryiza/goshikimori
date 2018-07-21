# API Шикимори для Golang
## Описание
Пакет предназначен для взаимодействия с [API Шикимори](https://shikimori.org/api/doc).

Сейчас поддерживается **только [API 1.0](https://shikimori.org/api/doc/1.0)**.

## Зависимости
* [github.com/golang/oauth2](https://github.com/golang/oauth2)
* [github.com/seryiza/loadOAuth](https://github.com/seryiza/loadOAuth)

## Установка
```bash
go get github.com/seryiza/go-shikimori
```

## Использование
### OAuth2
Прежде всего, для OAuth2 потребуется создать приложение [на самом Шикимори](https://shikimori.org/oauth/applications). Для авторизации потребуется *название приложения*, *client id* и *client secret*.

### Работа с API
Объект `api.Shikimori` предназначен для взаимодействия с API Шикимори. Можно использовать через:

* HTTP-запросы `Shikimori.Client` (работа с GET/POST/PUT/DELETE + JSON ложится на Вас):
```go
  // shiki := api.Shikimori{...}

  // => https://shikimori.org/api/users/whoami
  whoamiURL := shiki.ApiURL("/api/users/whoami")
  resp, _ := shiki.Client.Get(whoamiURL)
  whoamiJSON, _ := ioutil.ReadAll(resp.Body)
  fmt.Println(string(whoamiJSON))
  // => {"id":206253,"nickname":"Seryiza",...,"locale":"ru"}
```

* Методы структуры `Shikimori` (высокоуровневая работа с API и ее сущностями):
```go
// shiki := api.Shikimori{...}

user, err := shiki.Whoami()
fmt.Println(user, err)
// => &api.User{
//      ID: 206253,
//      Nickname: "Seryiza",
//      ...
//    }
```

## Пример
```go
conf := &oauth2.Config{
  ClientID:     "your client id",
  ClientSecret: "your client secret",
  // RedirectURL задан для использования как CLI
  // (!) Примечание: RedirectURL должен совпадать с тем,
  // что Вы задали на странице приложения Шикимори
  RedirectURL:  auth.StandaloneRedirectURL,
  Endpoint:     auth.ShikimoriEndpoint,
}

url := auth.GetAuthCodeURL(conf)
fmt.Println("Введите код из этой ссылки: ", url)

var code string
if _, err := fmt.Scanln(&code); err != nil {
  panic(err)
}

ctx := context.Background()
ctx = auth.AddShikimoriTransport(ctx, "your application name")

tok, err := conf.Exchange(ctx, code)
if err != nil {
  panic(err)
}

client := conf.Client(ctx, tok)
shiki := &api.Shikimori{
  Client: client,
}

user, err := shiki.Whoami()
fmt.Println(user, err)
```

Также есть вспомогательные функции `api.DefaultClientByCode`, `api.DefaultClientByToken`, `helpers.GetShikimori` (используя env-переменные) для меньшего написания кода.