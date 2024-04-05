# Get supported cards types

## Descripcion

Obtiene una lista de los tipos de tarjeta soportadas por el dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getSupportedCardTypes",
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
        "cardType":"QR",
        "valueType": "string"
    },
    {
        "cardType":"MIFARE_1K",
        "valueType": "numeric"
    },
    {
        "cardType":"WIEGAND",
        "valueType": "hex"
    },
    {
        "cardType":"WIEGAND_V3",
        "valueType": "alphanumeric"
    }
]
}
```

| Campo     | Tipo   | Descripcion                                                                                                                                   |
| --------- | ------ | --------------------------------------------------------------------------------------------------------------------------------------------- |
| cardType  | string | Formato físico de la Tarjeta como "QR", proximidad "MIFARE_K1" "WIEGAND" y otros                                                              |
| valueType | string | formato aceptado del valor de la tarjeta como "string" entiendase, cualquier string, "alphanumeric" numeros y letras, "hex" valor hexadecimal |
