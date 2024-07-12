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
           "channelNumber": "111-aaa",
           "enabled": true,
        }]
    }
}
```

| Campo         | Tipo    | Descripcion                           |
| ------------- | ------- | ------------------------------------- |
| channelNumber | string  | Id del canal                          |
| enabled       | boolean | establecer el video en blanco y negro |

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
