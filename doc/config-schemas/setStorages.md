# Set storages config

## Acciones

- Editar

## Descripcion

Editar Lista de dispositivos de almacenamiento (Video)

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setStorages",
        "value":[{
           "id": 0,
           "recording": true,
         }]  
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | int | Index a actualizar |
| recording | boolean | si el dispositivo está o no grabando|



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
