# Sets a zone name

## Descripcion

Setea el nombre de una zona

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAlarmZone",
        "value": {
            "partitionId" :"1",
            "zoneId": "1",
            "name": "Ventana 3, Sala 4"
        }
    }
}
```

| Campo       | Tipo   | Descripcion        |
| ----------- | ------ | ------------------ |
| partitionId | string | id de la particion |
| zoneId      | string | id de la zona      |
| name        | string | nombre de la zona  |

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
