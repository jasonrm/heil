# Heil - RabbitMQ Authentication via OAuth2 Proxy

## Requirements

* [https://github.com/bitly/oauth2_proxy](oauth2_proxy)
* [rabbitmq-auth-backend-http](https://github.com/rabbitmq/rabbitmq-auth-backend-http)

## Behavior

| Request Path      | Response Status Code | Body                  | Condition            |
|-------------------|----------------------|-----------------------|----------------------|
| /auth/user        | 200                  | `allow administrator` | if password matches  |
| /auth/vhost       | 200                  | `allow`               | always               |
| /auth/resource    | 200                  | `allow`               | always               |
| _everything else_ | 403                  | _empty_               | always               |

## Usage

`docker run -i -t -p 8080:8080 jasonrm/heil -basic-auth-password complexPassword`

## Options

```
Usage of ./heli:
  -basic-auth-password string
        Must be the same as set on oauth2_proxy
  -port int
        port number (default 8080)
```
