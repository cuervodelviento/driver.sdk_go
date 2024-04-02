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
            "number":1,
            "systemId": "xxx-xxxx",
            "zoneNumber":14
        }
    }
}

```

| Campo      | Tipo   | Descripcion                   |
| ---------- | ------ | ----------------------------- |
| number     | int    | numero de particion           |
| systemId   | string | Id de sistema de la particion |
| zoneNumber | int    | numero de la zona             |

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
