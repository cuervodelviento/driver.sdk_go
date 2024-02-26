# Set Add Person To AC 


## Descripcion

Agrega una persona en un control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAddPersonToAC",
        "value":{
            "personId": "string",
            "personName": "string",
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
