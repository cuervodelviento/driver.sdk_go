# Get storages config

## Descripcion

Obtener lista de estados de armado disponibles

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
        "armedStatus":"Away",
        "value": 0
    }
]
}
```

| Campo       | Tipo   | Descripcion      |
| ----------- | ------ | ---------------- |
| armedStatus | string | nombre de armado |
| value       | int    | valor del estado |
