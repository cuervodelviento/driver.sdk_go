# Get Current Video Resolution By Channel

## Descripcion

Obtener la resolucion actual del video por canal

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getCurrentVideoResolutionByChannel",
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
       "resolution": "1920x1080",
    }
}
```

| Campo      | Tipo   | Descripcion          |
| ---------- | ------ | -------------------- |
| resolution | string | Resolucion del canal |
