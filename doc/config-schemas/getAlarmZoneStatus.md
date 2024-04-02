# Get gets the status of a zone

## Descripcion

Obtener el status actual de una zona de alarma

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAlarmZoneStatus",
        "value":{
            "zoneId": "1",
            "partitionId":"1"
        }
    }
}
```

| Campo       | Tipo   | Descripcion        |
| ----------- | ------ | ------------------ |
| partitionId | string | id de la particion |
| zoneId      | string | id de la zona      |

## Response Message

```json
{
    ...,
    "data": {
       "alarmState": 1
        "byPassed": false,
        "openClosed": true
    }
}
```

| Campo      | Tipo    | Descripcion                                   |
| ---------- | ------- | --------------------------------------------- |
| alarmState | int     | Estado de alarma: 0:sin activar, 1: en alarma |
| byPassed   | boolean | si la zona se encuentra en bypass             |
| openClosed | boolean | si se encuentra activo el openclosed          |
