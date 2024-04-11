# Set Card to Person AC


## Descripcion

Agregar tarjeta(s) a una persona en un control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setCardToPersonAC",
        "value":{
            "personId": "string",
            "cards": ["string"],
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
