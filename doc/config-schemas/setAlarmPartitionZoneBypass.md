# Sets a zone name

## Descripcion

Setea el bypass a una zona de una particion

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAlarmPartitionZoneBypass",
        "value": {
            "partitionId" :"1",
            "zoneId": "1",
            "bypass": false
        }
    }
}
```

| Campo       | Tipo    | Descripcion              |
| ----------- | ------- | ------------------------ |
| partitionId | string  | id de la particion       |
| zoneId      | string  | id de la zona            |
| bypass      | boolean | bypass activo o inactivo |

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
