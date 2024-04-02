# Activates a FAP Alarm

## Descripcion

Activa alarma de Fuego Alarma o Panico

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionAlarmFapPartition",
        "value":{
            "number":1,
            "systemId": "xxx-xxxx",
            "value":1
        }
    }
}

```

| Campo    | Tipo   | Descripcion                          |
| -------- | ------ | ------------------------------------ |
| number   | int    | numero de particion                  |
| systemId | string | Id de sistema de la particion        |
| value    | int    | valor de FAP 1:fire, 2:alarm,3:panic |

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

| Campo | Tipo    | Descripcion                    |
| ----- | ------- | ------------------------------ |
| error | boolean | Error                          |
| msg   | string  | Mensaje de error o log interno |
