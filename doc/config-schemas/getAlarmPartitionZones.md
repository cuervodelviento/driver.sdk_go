# Get zones from an alarm partitions

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

            "partitionId":"1"
        }"
    }
}
```

| Campo       | Tipo   | Descripcion        |
| ----------- | ------ | ------------------ |
| partitionId | string | id de la particion |

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
