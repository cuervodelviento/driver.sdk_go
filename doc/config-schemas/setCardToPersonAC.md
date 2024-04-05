# Set Card to Person AC

## Descripcion

Agregar tarjeta(s) a una persona en un control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setCardToPersonAC",
        "value":{
            "personId": "string",
            "cards": [
                {
                    "cardValue":"23118291915511151",
                    "cardType":"QR",
                },
                {
                    "cardValue":"9181283181851",
                    "cardType":"MIFARE_1K",
                }
            ],
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
