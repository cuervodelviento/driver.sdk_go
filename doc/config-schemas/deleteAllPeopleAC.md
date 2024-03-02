# Delete people AC


## Descripcion

Eliminar todas las personas de la lista de acceso del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "deleteAllPeopleAC",
        "value":{} // no actionOutputPulse 
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
| error | boolean | Error en el proceso de la accion |
| msg | string | Mensaje de error o log interno|
