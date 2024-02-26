# Get storages config

## Descripcion

Lista de dispositivos de almacenamiento disponibles (Video)

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
    "id": 0,
    "deviceName":"DeviceTest",
    "recording": true,
    "freeSize":"30mb",
    "totalSize":"20mb",
    "status":""
    }]
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | int | index del item  |
| deviceName | string | nombre del store  |
| recording | boolean | está o no grabando |
| freeSize | string | Almacenamiento libre |
| totalSize | string | Almacenamiento total |
| status | string | Estado actual del dispositivo dependiendo del driver, ("online","offline","disconnected",etc) |
