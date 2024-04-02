# Get storages config

## Descripcion

Obtener las zonas de una particion de alarma y su informacion

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAlarmPartitions",
        "value": "{
            "number":1,
            "systemId":"xxx-xxxx"
        }"  // No necesita
    }
}
```

| Campo    | Tipo   | Descripcion                |
| -------- | ------ | -------------------------- |
| number   | int    | numero de la particion     |
| systemId | string | id interno de la particion |

## Response Message

```json
{
    ...,
    "data": [
    {
        "number": 1,
        "name": "Ventana 3, Sala 4",
        "alarmState": 1
        "byPassed": false,
        "openClosed": true
    }
]
}
```

| Campo      | Tipo    | Descripcion                                     |
| ---------- | ------- | ----------------------------------------------- |
| number     | int     | numero de la zona                               |
| name       | string  | nombre de la zona                               |
| alarmState | int     | Estado de alarma: 0, ninguna, 1,2,3 POR DEFINIR |
| byPassed   | boolean | si la zona se encuentra en bypass               |
| openClosed | boolean | si se encuentra activo el openclosed            |
