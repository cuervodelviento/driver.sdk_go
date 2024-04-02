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
            "number":1,
            "systemId": "xxx-xxxx",
            "value":1
        }
    }
}

```

| Campo    | Tipo   | Descripcion                         |
| -------- | ------ | ----------------------------------- |
| number   | int    | numero de particion                 |
| systemId | string | Id de sistema de la particion       |
| value    | int     | valor de armado 1:away, 2:night etc |

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
