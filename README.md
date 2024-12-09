# Meower relation service

Cервис, отслеживающий связи между пользователями.

Предоставляет доступ к связям пользователей (follow, mute) посредством grpc c интерфейсом и сообщениями описанным в [api](https://github.com/Karzoug/meower-api/tree/main/proto/relation).

### Стек
- Основной язык: go
- База данных: neo4j
- Брокер: kafka
- Наблюдаемость: opentelemetry, jaeger, prometheus
- Контейнеры: docker, docker compose