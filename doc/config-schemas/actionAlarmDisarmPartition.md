# Disarms a partition

## Descripcion

Desarma una particion

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionAlarmDisarmPartition",
        "value":{
            "number":1,
            "systemId": "xxx-xxxx",

        }
    }
}

```

| Campo    | Tipo   | Descripcion                   |
| -------- | ------ | ----------------------------- |
| number   | int    | numero de particion           |
| systemId | string | Id de sistema de la particion |

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
