# Deletes a zone from a partition

## Descripcion

Remueve una zona de una particion

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setDeleteAlarmPartitionZone",
        "value":{
            "partitionId": "xxx-xxxx",
            "zoneId":"14"
        }
    }
}

```

| Campo       | Tipo   | Descripcion        |
| ----------- | ------ | ------------------ |
| partitionId | string | Id de la particion |
| zoneId      | string | Id de la zona      |

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
