# Set QR to Person AC


## Descripcion

Agregar QR(s) a una persona en un control de acceso

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "setQRToPersonAC",
        "value":{
            "personId": "string",
            "qrCodes": [{
                "value": "string",
                "type": 1
            }],

        } // no needed
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| personId | string | Id de la persona |
| qrCodes | array | Arreglo de objetos QR |
| qrCodes.value | string | Valor del QR |
| qrCodes.type | int | Tipo de QR. Puedes usar el ID del [tipo de tecnologia](https://github.com/Netsocs-Team/driver.sdk_go/blob/main/doc/config-schemas/tech_types.md) donde `TipoMedidaBiometrica` sea QR. O en el SDK hay constantes de tipo `QRType` que contienen los tipos de QR |



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

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| error | boolean | Error |
| msg | string | Mensaje de error o log interno|
