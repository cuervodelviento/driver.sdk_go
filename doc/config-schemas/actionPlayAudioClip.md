# Action Play Audio Clip


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
            "audioClipUrl": "http://www.example.com/audio.mp3",
            "speakerId": "1"
        }
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| audioClipUrl | string | URL del clip de audio |
| speakerId | string | ID del altavoz |




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

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| error | boolean | Error |
| msg | string | Mensaje de error o log interno|
