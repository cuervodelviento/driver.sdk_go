# Set Action Unlock Device

## Descripcion

Desbloquear el dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setActionUnlockDevice",
        "value":{
            "unlock": true,
            "unlockCode": "1234",
            "unlockMode": "unlockMode",
            "aditionalData": "aditionalData",
            "deviceId": "bbc5203d-7f06-4806-8793-2026904738c8"
        } // no needed
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| unlock | boolean | Desbloquear el dispositivo |
| unlockCode | string | Codigo de desbloqueo |
| unlockMode | string | Modo de desbloqueo |
| aditionalData | string | Datos adicionales |
| deviceId | string | Id del dispositivo |



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
| error | boolean | Error al hacer restart |
| msg | string | Mensaje de error o log interno|
