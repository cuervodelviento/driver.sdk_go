# Arms a partition

## Descripcion

Arma una particion con un valor de armado

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionAlarmArmPartition",
        "value":{

            "partitionId": "xxx-xxxx",
            "value":1
        }
    }
}

```

| Campo       | Tipo   | Descripcion                           |
| ----------- | ------ | ------------------------------------- |
| partitionId | string | Id de sistema de la particion         |
| value       | int    | valor de armado clear:0,away:1,stay:2,night:3 |

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
