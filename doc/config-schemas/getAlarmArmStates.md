# Get Alarm Arm states

## Descripcion

Obtener lista de estados de armado disponibles, para mostrar únicamente los que estan
disponibles en el sistema

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAlarmArmStates",
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
        "armedStatus":"Away",
        "value": 1
    }
]
}
```

| Campo       | Tipo   | Descripcion                            |
| ----------- | ------ | -------------------------------------- |
| armedStatus | string | nombre de armado (away,stay,night)     |
| value       | int    | valor del estado away:1,stay:2,night:3 |
