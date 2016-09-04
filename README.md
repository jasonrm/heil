# Heil - RabbitMQ auth via OAuth2 Proxy

## Requirements

* [https://github.com/bitly/oauth2_proxy](oauth2_proxy)
* [rabbitmq-auth-backend-http](https://github.com/rabbitmq/rabbitmq-auth-backend-http)

| Request Path      | Response Status Code | Body                                      |
|-------------------|----------------------|-------------------------------------------|
| /auth/user        | 200                  | `allow administrator` if password matches |
| /auth/vhost       | 200                  | `allow`                                   |
| /auth/resource    | 200                  | `allow`                                   |
| _everything else_ | 403                  | _empty_                                   |

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
