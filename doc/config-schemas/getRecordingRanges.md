# Get channels config

## Descripcion

Devolver una lista de pares de fechas en los cuales el canal tiene video grabado

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getRecordingRanges",
        "value": {
            "channelId": "1-3",
            "utcStart":  "2024-07-12T14:40:10Z",
            "utcEnd":    "2024-07-12T16:30:25Z",
        }
    }
}
```

| Campo     | Tipo   | Descripción                                                   |
| --------- | ------ | ------------------------------------------------------------- |
| channelId | string | id del canal de video                                         |
| utStart   | string | debe seguir el formato UTC "YYYY-MM-DDTHH:mm:ssZ" ES opcional |
| utcEnd    | string | debe seguir el formato UTC "YYYY-MM-DDTHH:mm:ssZ" es opcional |

Nota sobre el formato UTC
YYYY son el año en 4 dígitos
MM son el mes en 2 dígitos (de 01 a 12)
DD son el día en 2 dígitos (de 01 a 31)
HH son la hora en 2 dígitos (de 00 a 23)
mm son los minutos en 2 dígitos (de 00 a 59)
ss son los segundos en 2 dígitos (de 00 a 59)
Z representa la zona horaria 0, es decir, timestamps locales ("America/Guyana" "America/Indianapolis" "America/Cancun" ) deben ajustarse a hora UTC
La granularidad es sobre segundos para garantizar la compatibilidad entre marcas (aún cuando algunas aceptan milésimas)

utcEnd sigue el mismo formato, ambas fechas son opcionales

## Response Message

```json
{
    ...,
    "data": [
        {
        "utcStart": "2024-07-12T14:40:10Z",
        "utcEnd":   "2024-07-12T15:40:10Z"
        },
        {
        "utcStart": "2024-08-12T14:40:10Z",
        "utcEnd":   "2024-08-12T15:40:10Z"
        },
        {
        "utcStart": "2024-09-12T14:40:10Z",
        "utcEnd":   "2024-09-12T15:40:10Z"
        },
    ]
}
```

| Campo    | Tipo   | Descripcion                        |
| -------- | ------ | ---------------------------------- |
| utcStart | string | Inicio de bloque con video grabado |
| utcEnd   | string | Fin de bloque con video grabado    |
