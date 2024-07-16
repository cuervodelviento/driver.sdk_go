# Set Channel In Black And White


## Descripcion

Poner un canal de video en blanco y negro

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setChannelInBlackAndWhite",
        "value":{
            "channelId": "string",
        } // no needed
    }
}
```



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
