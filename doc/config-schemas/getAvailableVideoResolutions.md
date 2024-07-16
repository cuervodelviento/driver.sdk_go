# Get available video resolutions

## Descripcion

Obtener todas las resoluciones de video disponibles en el dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAvailableVideoResolutions",
        "value": {
            "channelId": "111-aaa"
        }
    }
}
```

| Campo     | Tipo   | Descripcion |
| --------- | ------ | ----------- |
| channelId | string | Id de canal |

## Response Message

```json
{
    ...,
    "data": {
        "resolutions": [
            "1920x1080",
            "1280x720",
            "640x480"
        ]
    }
}
```

| Campo       | Tipo     | Descripcion                       |
| ----------- | -------- | --------------------------------- |
| resolutions | []string | Lista de resoluciones disponibles |
