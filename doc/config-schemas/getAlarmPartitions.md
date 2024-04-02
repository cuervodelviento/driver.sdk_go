# Get alarm partitions

## Descripcion

Obtener las particiones de una alarma y su informacion

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAlarmPartitions",
        "value": "{}"  // No necesita
    }
}
```

## Response Message

```json
{
    ...,
    "data": [
    {
        "partitionId": "1",
        "systemId": "xxxxx-xxxx",
        "name": "Perimetral Este",
        "enabled": true,
        "armedStatus": 0,
        "troubleState": 0,
        "alarmState": 1
    }
]
}
```

| Campo        | Tipo    | Descripcion                                                                         |
| ------------ | ------- | ----------------------------------------------------------------------------------- |
| partitionId  | string  | id de la partición                                                                  |
| systemId     | string  | id interno del dispositivo                                                          |
| name         | string  | nombre de la partición                                                              |
| enabled      | boolean | particion activa                                                                    |
| armedStatus  | int     | Estado de la particion= 0: desarmada, 1:armada away, 2: armada stay, 3:armada night |
| troubleState | int     | Estado de problema= 0: ninguno, 1:problema en una o varias zonas                    |
| alarmState   | int     | Estado de alarma= 0: sin activar, 1: en alarma (se activo una o varias zonas)       |
