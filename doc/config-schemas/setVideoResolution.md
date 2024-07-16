# Set video resolution

## Descripcion

Establecer la resolucion de video para un canal

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setVideoResolution",
        "value": {
            "channelId": "111-aaa",
            "resolution": "1920x1080"
        }
    }
}
```

| Campo      | Tipo   | Descripcion         |
| ---------- | ------ | ------------------- |
| channelId  | string | Id del canal        |
| resolution | string | Resolucion de video |

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
