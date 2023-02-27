#SOFTWARE ENGINEERING COURSE
Стиль написания кода - camelCase
Подход к разработке  - Agile 
Таск менеджер  - github actions 



# BACKEND 

gorilla - logrus - postgreSQL/pq - testify - gomail

# Installation

## Supported Versions

This library supports the following Go implementations:

* Go 1.19

## Install Package

```bash
go get github.com/t67y110v/software-engineering
```

## Configuration setup

#### configs/apiserver.toml

```toml
[server]
bind_addr=":8080"
log_level="debug"

[database]
database_url="user=postgres password=p02tgre2 dbname=restapi sslmode=disable"


```

## Endpoints

| Name | Description |
|------|-------------|
| **/userCreate** | Creating a new user |
| **/userUpdate** | Updates user data |
| **/userDelete** | Deleting a user |
| **/getPositions** | User authorization |
| **/changePassword** | Changes user password |
| **/sendEmail** | Sends email |

## Request | Responds
### Json models in jsons/requests | jsons/responds
