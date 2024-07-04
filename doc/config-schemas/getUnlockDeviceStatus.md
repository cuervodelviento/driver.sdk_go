# Get Unlock Device Status

## Descripcion

Obtener el estado de desbloqueo del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setActionUnlockDevice",
        "value":{
            "deviceId": "bbc5203d-7f06-4806-8793-2026904738c8"
        } 
    }
}
```



## Response Message
```json
{
    ...,
    "data": {
        "unlockStatus": "UNLOCKING"
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| status | string | Puede ser alguno de estos posibles status: UNLOCKING, DEVICE_OFFLINE, DEVICE_BUSY, DEVICE_FORBIDDEN, DEVICE_SAY_ERROR |
