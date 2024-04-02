# Activates a FAP Alarm

## Descripcion

Activa alarma de Fuego Alarma o Panico

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionAlarmFapPartition",
        "value":{
            "partitionId": "xxx-xxxx",
            "value":1
        }
    }
}

```

| Campo       | Tipo   | Descripcion                                                              |
| ----------- | ------ | ------------------------------------------------------------------------ |
| partitionId | string | Id de la particion                                                       |
| value       | int    | valor de FAP Fire:1, Alarm:2, Panic:3, Assistance:4, Police:5, Medical:6 |

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
