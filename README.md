## API для пользователей
Запуск

> go run cmd/app/main.go

Настройки приложения хранятся в файле config.json

Перед стартом приложения необходимо запустить БД mysql и создать в ней таблицу **user**:

```
create table user (
id      int auto_increment
primary key,
name    varchar(128)            not null,
age     int                     not null,
friends varchar(128) default '' not null
);
```

## Прокси
Запуск

> go run cmd/proxy/main.go 6060 6061

Если не указать порты, то по умолчанию будут использованы 8080 и 8081.