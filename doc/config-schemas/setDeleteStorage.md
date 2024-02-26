# Delete storages config

## Acciones

- Delete

## Descripcion

Eliminar item de Lista de dispositivos de almacenamiento (Video)

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setDeleteStorage",
        "value":{
           "id": [number], // requerido para saber el id a elimiar
         }  
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | int | Index dado o no por el driver |



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
