# Get available outputs

## Descripcion

Obtener todas las salidas disponibles en el dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAvailableOutputs",
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
    "name": "output 1",
    "description": "output 1 description",
    "type": 0,
    "state": 0
    }]
}
```



| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | int | id de la salida |
| name | string | nombre de la salida |
| description | string | descripcion de la salida |
| type | int | tipo de salida, 0 = relay, 1 = led |
| state | int | estado de la salida, 0 = normalmente abierto, 1 = normalmente cerrado, 2 = encendido, 3 = apagado |
