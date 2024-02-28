# Listen Event config


## Descripcion

Bloquear persona en el control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setBlockPersonToAC",
        "value":{
            "personId": "em329",
        } 
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| personId | string | Id de la persona |



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
