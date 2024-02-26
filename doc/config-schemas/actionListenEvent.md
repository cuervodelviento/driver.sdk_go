# Listen Event config

## Acciones

- action

## Descripcion

Activar escucha de eventos

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionListenEvent",
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
| error | boolean | Error |
| msg | string | Mensaje de error o log interno|
