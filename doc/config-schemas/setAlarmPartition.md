# Get storages config

## Descripcion

Setea el nombre de una partición

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAlarmPartition",
        "value": {
            "partitionId": "1",
            "systemId": "xxxxx-xxxx",
            "name": "Perimetral Este"
        }
    }
}
```

## Response Message

```json
{
    ...,
    "data": {
        "error" : false,
        "msg" : ""
    }
}
```

| Campo | Tipo    | Descripcion                    |
| ----- | ------- | ------------------------------ |
| error | boolean | Error                          |
| msg   | string  | Mensaje de error o log interno |
