# msb
mesh service bus - enterprise service bus on service mesh

## 1. Project Goal

* Cost reduction for development and operation
* Make it easier to integrate with any components
* Provide as many functionalities as possible

`If you have to learn a lot to achieve one thing, you should rethink if it is really suitable for you.`

## 2. Key mechanisms

### 2.1 Entrypoint definition

* HTTP Web: http/https
* HTTP API: http/https
* Job Scheduler
* Event/Message

### 2.2 General processing flow

* Caller -> Entrypoint -> Process

### 2.3 Process and kernel API definition

* Process
* Aggregation
* Binding
* Router
* Service
* Splitter
* Transformer

## Dependencies

```text
* Toml encoding/decoding:
github.com/toml-lang/toml

* Logging
github.com/sirupsen/logrus

* Http router
github.com/julienschmidt/httprouter

* Job scheduler
github.com/go-co-op/gocron

* Database driver
** Postgresql
github.com/jackc/pgx/v4
```

## Misc

1. Logo generator. Thanks to: https://patorjk.com/software/taag/#p=display&h=0&v=0&f=Doh&t=MSB

