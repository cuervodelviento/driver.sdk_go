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
        "value": [{
           "id" : 0,
           "used": true,
           "userName": "test",
           "password":"...",
        }]
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
