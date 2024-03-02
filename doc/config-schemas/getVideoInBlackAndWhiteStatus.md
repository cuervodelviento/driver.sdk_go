# Get video in black and white status config

## Descripcion

Obtener informacion sobre si el video en blanco y negro esta habilitado

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getVideoInBlackAndWhiteStatus",
        "value": {
            "channelNumber": 1
        }
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| channelNumber | number | Numero del canal |



## Response Message
```json
{
    ...,
    "data": {
        "enabled": false
    }
}
```

Solo el nombre de usuario para no revelear el password del admin

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| enabled | boolean | Si el video en blanco y negro esta habilitado sera true |
