# Arms a partition

## Descripcion

Arma una particion con un valor de armado

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
| value    | id     | valor de FAP 1:fire, 2:alarm,3:panic |

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
