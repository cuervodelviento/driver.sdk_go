# 


## Descripcion

Reproducir un clip de audio

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionPlayAudioClip",
        "value":{
            "channelNumber": 1,
            "start": "2019-01-01T00:00:00Z",
            "end": "2019-01-01T00:00:00Z"
        }
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| channelNumber | int | Numero de canal de video |
| start | string | Fecha de inicio |
| end | string | Fecha de fin |




## Response Message
```json
{
    ...,
    "data": {
        "filename": "heatmap.jpg"
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| filename | string | Nombre del archivo de la imagen del heatmap |
