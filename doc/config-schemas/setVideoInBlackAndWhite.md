# Set video in black and white

## Descripcion

Establecer el video en blanco y negro

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setVideoInBlackAndWhite"
        "value": [{
           "channelId": "111-aaa",
           "enabled": true,
        }]
    }
}
```

| Campo     | Tipo    | Descripcion                           |
| --------- | ------- | ------------------------------------- |
| channelId | string  | Id del canal                          |
| enabled   | boolean | establecer el video en blanco y negro |

## Response Message

```json
{
    ...,
    "data": {
        "error" : false,
        "msg" : ""
    }
}
```

| Campo | Tipo    | Descripcion                    |
| ----- | ------- | ------------------------------ |
| error | boolean | Error al actualizar            |
| msg   | string  | Mensaje de error o log interno |
