# Delete users config

## Acciones

- delete

## Descripcion

eliminar item de la lista de usuarios del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setDeleteUser",
        "value": {
           "id" : [0] // requerido
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
