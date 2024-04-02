# Adds a zone to a partition

## Descripcion

Añade una zona a una particion

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAddAlarmPartitionZone",
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
