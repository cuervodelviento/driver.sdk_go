# Get User Group Permission Types Config

Esta configuracion permite informar a Netsocs cuales son los permisos que maneja el dispositivo a nivel de roles o grupos de usuarios. Esto permite identificar correctamente los permisos que se pueden asignar a los usuarios en un dispositivo determinado.

Es usada normalmente antes de crear un grupo

![](../../../img/first%20request%20permission%20type%20before%20create%20user%20group.png)

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getUserGroupPermissionTypes",
        "value": "{}" 
    }
}
```

## Response Message

```json
{
    ...,
    "data": [{
        "name": "canViewLive",
        "description": "Can view live video",
    }]
}
```
