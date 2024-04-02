# Get zones from an alarm

## Descripcion

Obtener las zonas de una alarma y su informacion

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
        "zoneId": "1",
        "name": "Ventana 3, Sala 4",
        "alarmState": 1
        "byPassed": false,
        "openClosed": true
    }
]
}
```

| Campo      | Tipo    | Descripcion                                   |
| ---------- | ------- | --------------------------------------------- |
| zoneId     | string  | id de la zona, puede ser el numero de la zona |
| name       | string  | nombre de la zona                             |
| alarmState | int     | Estado de alarma: 0:sin activar, 1: en alarma |
| byPassed   | boolean | si la zona se encuentra en bypass             |
| openClosed | boolean | si se encuentra activo el openclosed          |
