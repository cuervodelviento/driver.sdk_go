# Stop Listen Event config

## Acciones

- action

## Descripcion

Detener eventos del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionStopListenEvent",
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
| error | boolean | Error  |
| msg | string | Mensaje de error o log interno|
