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
