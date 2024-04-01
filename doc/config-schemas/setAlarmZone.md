# Get storages config

## Descripcion

Setea el nombre de una zona

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setAlarmZone",
        "value": {
            "number": 1,
            "name": "Ventana 3, Sala 4"
        }
    }
}
```

| Campo  | Tipo   | Descripcion       |
| ------ | ------ | ----------------- |
| number | int    | numero de la zona |
| name   | string | nombre de la zona |

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
