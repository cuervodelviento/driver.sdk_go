# Ping Device action

## Acciones

- action

## Descripcion

Prueba de conexion al dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionPingDevice",
        "value":{} // no needed
    }
}
```



## Response Message
```json
{
    ...,
    "data": {
        "status" : true,
        "error" : false,
        "msg" : "",
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| status | boolean | conectado/desconectado |
| error | boolean | Error al hacer restart |
| msg | string | Mensaje de error o log interno|
