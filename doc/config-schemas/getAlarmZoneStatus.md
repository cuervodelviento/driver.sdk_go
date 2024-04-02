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
            "number": 1
        }
    }
}
```

| Campo  | Tipo | Descripcion       |
| ------ | ---- | ----------------- |
| number | int  | Numero de la zona |

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

| Campo      | Tipo    | Descripcion                                     |
| ---------- | ------- | ----------------------------------------------- |
| alarmState | int     | Estado de alarma: 0, ninguna, 1,2,3 POR DEFINIR |
| byPassed   | boolean | si la zona se encuentra en bypass               |
| openClosed | boolean | si se encuentra activo el openclosed            |
