# Set Add user To Alarm partition

## Descripcion

Agrega un usuario a una particion de alarma

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAddAlarmUser",
        "value":{
            "partitionId":"xxx-xxxx",
            "userName": "Em-32",
            "accessCode": "5125"
        }
    }
}

```

| Campo       | Tipo   | Descripcion        |
| ----------- | ------ | ------------------ |
| partitionId | string | Id de la particion |
| userName    | string | Nombre del usuario |
| accessCode  | string | codigo de acceso   |

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
