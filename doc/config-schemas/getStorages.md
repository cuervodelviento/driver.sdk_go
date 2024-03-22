# Get storages config

## Descripcion

Obtener todas las unidades de almacenamiento disponibles en el dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getStorages",
        "value": "{}"  // No necesita
    }
}
```



## Response Message
```json
{
    ...,
    "data": [ {
    "id": "42b0f590-6e86-4b7a-8dc0-44056b28c018",
    "deviceName":"DeviceTest",
    "recording": true,
    "freeSize":5000, 
    "totalSize": 4000,
    "status":0
    }]
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | string | id del item  |
| deviceName | string | nombre del store  |
| recording | boolean | está o no grabando |
| freeSize | string | Almacenamiento libre en MB |
| totalSize | string | Almacenamiento total en MB |
| status | string | Estado del almacenamiento, 0: OK, 1: Error, 2: Warning, 3: Unknown |

