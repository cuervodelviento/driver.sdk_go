# Set QR to Person AC

## Descripcion

Agregar códigos QR(s) a una persona en un control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setQRToPersonAC",
        "value":{
            "personId": "string",
            "qrs": ["string"],
        } // no needed
    }
}
```

| Campo    | Tipo   | Descripcion                                    |
| -------- | ------ | ---------------------------------------------- |
| personId | string | id del usuario como string, ejemplo em-73145   |
| qrs      | string | valor del código QR, ejemplo "841141295241124" |

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
