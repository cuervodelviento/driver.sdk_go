# Set user To Alarm partition

## Descripcion

Agrega un usuario a una particion de alarma

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAlarmUser",
        "value":{
            "systemId":"xxx-xxxx",
            "userName": "Em-32",
            "userId:"xxx-xxxx",
            "accessCode": "5125"
        }
    }
}

```

| Campo      | Tipo   | Descripcion                   |
| ---------- | ------ | ----------------------------- |
| systemId   | string | Id de sistema de la particion |
| userId     | string | Id de sistema de el usuario   |
| userName   | string | Nombre del usuario            |
| accessCode | string | codigo de acceso              |

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
