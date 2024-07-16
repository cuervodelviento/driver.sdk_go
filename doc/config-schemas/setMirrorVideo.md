# Set Mirror Video

## Descripcion

Aplicar un efecto de espejo a la imagen del video

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setMirrorVideo",
        "value":{
            "enabled": true,
            "channelId": "111-aaa"

        }
    }
}
```

| Campo     | Tipo    | Descripcion                                  |
| --------- | ------- | -------------------------------------------- |
| enabled   | boolean | Habilitar o deshabilitar el efecto de espejo |
| channelId | string  | Id de canal de video                         |

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
| error | boolean | Error al hacer backup          |
| msg   | string  | Mensaje de error o log interno |
