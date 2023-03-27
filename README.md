# ZZap Parser

Парсер сервиса [ZZap](https://www.zzap.ru/https://www.zzap.ru/) с помощью [API](https://wiki.zzap.ru/api-zzap/https://wiki.zzap.ru/api-zzap/).

### Структура проекта:

* *[main](main/)* - Старт программы;
* *[internal](internal/)* - Ядро, которое парсит и в котором можно управлять работой программы. Бизнес-логика.
* *[pkg](zzap/)* - Пакеты необходимые для работы программы. Например, тут есть [пакет zzap](pkg/zzap/), который является прослойкой между [API ZZap](https://wiki.zzap.ru/api-zzap/https://wiki.zzap.ru/api-zzap/) и Вашим приложением.
