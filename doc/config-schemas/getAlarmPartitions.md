# Get storages config

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
        "number": 1,
        "systemId": "xxxxx-xxxx",
        "name": "Perimetral Este",
        "enabled": true,
        "armedStatus": 0,
        "troubleState": 3,
        "alarmState": 1
    }
]
}
```

| Campo        | Tipo    | Descripcion                                                                      |
| ------------ | ------- | -------------------------------------------------------------------------------- |
| number       | int     | numero de partición                                                              |
| systemId     | string  | id interno del dispositivo                                                       |
| name         | string  | nombre de la partición                                                           |
| enabled      | boolean | particion activa                                                                 |
| armedStatus  | int     | Estado de la particion 0: desarmada,1:armada away, 2: armada night,3:armada stay |
| troubleState | int     | Estado de problema :0 ninguno, 1,2,3 POR DEFINIR                                 |
| alarmState   | int     | Estado de alarma: 0, ninguna, 1,2,3 POR DEFINIR                                  |
