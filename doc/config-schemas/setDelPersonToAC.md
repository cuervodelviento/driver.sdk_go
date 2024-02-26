# Set Del Person To AC 


## Descripcion

Elimina una persona en un control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setDelPersonToAC",
        "value":{
            "personId": "string",
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
