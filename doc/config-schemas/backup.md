# Backup config

## Acciones

- action

## Descripcion

Hacer un backup de respaldo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "backup",
        "value":{} // no needed
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
| error | boolean | Error al hacer backup |
| msg | string | Mensaje de error o log interno|
