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
            "channelNumber": "111-aaa"
        }
    }
}
```

| Campo         | Tipo   | Descripcion           |
| ------------- | ------ | --------------------- |
| channelNumber | string | Id del canal de video |

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
