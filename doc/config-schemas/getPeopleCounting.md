# Get people counting

## Descripcion

Devolver una lista de lineas de conteo de personas

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getPeopleCounting",
        "value": {
            "channelNumber": "111-aaa",
            "utcStart":  "2024-07-12T14:40:10Z",
            "utcEnd":    "2024-07-12T16:30:25Z",
        }
    }
}
```

| Campo         | Tipo   | Descripción                                                                                                                                                   |
| ------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| channelNumber | string | id del canal de video                                                                                                                                         |
| utStart       | string | debe seguir el formato [UTC](https://agileappscloud.info/aawiki/UTC_Format#:~:text=A%20date/time%20looks%20like%20this%3A) "YYYY-MM-DDTHH:mm:ssZ" es opcional |
| utcEnd        | string | debe seguir el formato [UTC](https://agileappscloud.info/aawiki/UTC_Format#:~:text=A%20date/time%20looks%20like%20this%3A) "YYYY-MM-DDTHH:mm:ssZ" es opcional |

## Response Message

```json
{
    ...,
    "data": [
        {
        "lineId": "111-aaa",
        "name": "Linea 1",
        "inCount": 10,
        "outCount": 5
        },
        {
        "lineId": "111-aaa",
        "name": "Linea 2",
        "inCount": 5,
        "outCount": 2
        },
        {
        "lineId": "111-aaa",
        "name": "Linea 3",
        "inCount": 3,
        "outCount": 1
        },
    ]
}
```

| Campo    | Tipo   | Descripcion                              |
| -------- | ------ | ---------------------------------------- |
| lineId   | string | id de la linea de conteo de personas     |
| name     | string | nombre de la linea de conteo de personas |
| inCount  | int    | cantidad de personas que entran          |
| outCount | int    | cantidad de personas que salen           |
