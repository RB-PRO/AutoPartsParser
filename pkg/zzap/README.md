# ZZap

[Документация по API](https://wiki.zzap.ru/api-zzap/https://wiki.zzap.ru/api-zzap/)

## Как это работает

Давайте рассмотрим на примере метода [GetRegionsV2](https://wiki.zzap.ru/api2-информация/#_Метод_регионы_поиска_GetRegionsVC__1https://wiki.zzap.ru/api2-информация/#_Метод_регионы_поиска_GetRegionsVC__1).

Реализация данного метода находится в [GetRegionsV2.go](GetRegionsV2.go "GetRegionsV2"), а соответствующий тест находится в файле [GetRegionsV2_test.go](GetRegionsV2_test.go).


## С чего начать

Пакет имеет максимально простую структуру со строгими структурами запроса-ответа.

### Настройка json

Создайте json-файл с названием "lap.json", который будет расположен в директории проекта.

###### Структура lap.json

```json
{
    "login": "login",
    "password":"password",
    "api_key":"api_key"
}
```

### Инициализируйте пакет

```go
lap, lapError := zzap.New()
if lapError != nil {
	t.Error(lapError)
}
fmt.Printf("New: Получил данные: %+v", lap)
```

Экземпляр lab необходим для работы с пакетом ZZap.
