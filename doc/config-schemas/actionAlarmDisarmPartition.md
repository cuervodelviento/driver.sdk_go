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
            "partitionId": "xxx-xxxx",

        }
    }
}

```

| Campo       | Tipo   | Descripcion        |
| ----------- | ------ | ------------------ |
| partitionId | string | Id de la particion |

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
