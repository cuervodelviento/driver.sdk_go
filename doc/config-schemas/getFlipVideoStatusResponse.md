# Get Flip Video Status

## Descripcion

Obtener informacion sobre si el video esta volteado (verticalmente) o no

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getFlipVideoStatus",
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
