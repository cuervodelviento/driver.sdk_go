# Restart config

## Acciones

- action

## Descripcion

Reiniciar el dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionRestart",
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
| error | boolean | Error al hacer restart |
| msg | string | Mensaje de error o log interno|
