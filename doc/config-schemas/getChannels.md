# Get channels config

## Descripcion

Devolver una lista de los canales que estan registrados en el dispositivo.

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getChannels",
        "value": "{}"
    }
}
```

El Request Message de `getChannels` no necesita value. El driver solo con el configKey identificara la operacion que debe realizar.

## Response Message

```json
{
    ...,
    "data": [{
        "name": "string",
        "channelNumber": "111-aaa",
        "rtspSource": "rtsp://"
    }]
}
```

| Campo         | Tipo   | Descripcion              |
| ------------- | ------ | ------------------------ |
| name          | string | El nombre del canal      |
| channelNumber | string | El Id del canal          |
| rtspSource    | string | La fuente rtsp del canal |
