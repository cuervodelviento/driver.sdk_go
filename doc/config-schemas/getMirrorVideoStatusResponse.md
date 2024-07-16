# Get Mirror Video Status

## Descripcion

Obtener el estado del efecto de espejo de la imagen del video

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setMirrorVideoStatus",
        "value":{
            "channelId": "111-aaa"
        }
    }
}
```

| Campo     | Tipo   | Descripcion           |
| --------- | ------ | --------------------- |
| channelId | string | Id del canal de video |

## Response Message

```json
{
    ...,
    "data": {
       "enabled": true,
    }
}
```

| Campo   | Tipo    | Descripcion                 |
| ------- | ------- | --------------------------- |
| enabled | boolean | Estado del efecto de espejo |
