# Set admin user config

## Descripcion

Modificar credenciales del usuario admin 

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAdminUser",
        "value": {
            "id": "68e5cc48-e23f-46b5-a491-2a623fccf658",
            "username": "username",
            "prevPassword": "***",
            "newPassword": "***"
        } 
    }
}
```



## Response Message
```json
{
    ...,
     "data": "{
        "error" : "false",
        "msg" : ""
    }"
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| error | boolean | Error al actualizar o error del password anterior |
| msg | string | Mensaje de error o log interno|
