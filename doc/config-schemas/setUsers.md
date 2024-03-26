# Set users config

## Acciones

- Editar 
- Crear

## Descripcion

Editar lista de usuarios del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setUsers",
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
    "data": {
        "error" : false,
        "msg" : ""
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| error | boolean | Error al actualizar |
| msg | string | Mensaje de error o log interno|
