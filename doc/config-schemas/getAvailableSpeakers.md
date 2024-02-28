# Get available speakers

## Descripcion

Obtener todos los altavoces disponibles en el dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAvailableSpeakers",
        "value": "{}"  // No necesita
    }
}
```



## Response Message
```json
{
    ...,
    "data": [{
    "id": 1,
    "level": 0,
    "minLevel": 0,
    "maxLevel": 100,
    "enabled": true,
    "name": "speaker 1"
    }]
}
```



| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | int | id de la salida |
| level | int | nivel del altavoz |
| minLevel | int | nivel minimo del altavoz |
| maxLevel | int | nivel maximo del altavoz |
| enabled | bool | estado del altavoz |
| name | string | nombre del altavoz |
