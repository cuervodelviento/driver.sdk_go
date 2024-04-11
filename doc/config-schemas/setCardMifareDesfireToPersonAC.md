# Set Smart Card to Person AC

## Descripcion

Agregar tarjeta(s) a una persona en un control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setCardMifareDesfireToPersonAC",
        "value":{
            "personId": "string",
            "cards": ["string"],
        } // no needed
    }
}
```

| Campo    | Tipo   | Descripcion                                                                                        |
| -------- | ------ | -------------------------------------------------------------------------------------------------- |
| personId | string | id del usuario ejemplo: em-17415                                                                   |
| cards    | string | valor numérico de la tarjeta expresado como string (mejor compatibilidad para enteros muy grandes) |

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
