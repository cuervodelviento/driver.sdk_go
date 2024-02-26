# Upgrade Firmware config

## Acciones

- action

## Descripcion

Actualizar firmware 

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "upgradeFirmware",
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
| error | boolean | Error al actualizar |
| msg | string | Mensaje de error o log interno|
