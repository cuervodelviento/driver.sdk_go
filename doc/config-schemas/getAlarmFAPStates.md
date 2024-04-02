# Get storages config

## Descripcion

Obtener lista de estados de FAP disponibles

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getAlarmFAPStates",
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
        "FAPStatus":"Fire",
        "value": 0
    }
]
}
```

| Campo     | Tipo   | Descripcion      |
| --------- | ------ | ---------------- |
| FAPStatus | string | nombre de status |
| value     | int    | valor del status |
