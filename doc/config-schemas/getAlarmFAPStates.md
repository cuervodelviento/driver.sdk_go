# Get Alarm FAP states

## Descripcion

Obtener lista de estados de FAP disponibles en el sistema

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
        "value": 1
    },
    {
        "FAPStatus":"Alarm",
        "value": 2
    },
    {
        "FAPStatus":"Panic",
        "value": 3
    },
]
}
```

| Campo     | Tipo   | Descripcion                                                                  |
| --------- | ------ | ---------------------------------------------------------------------------- |
| FAPStatus | string | nombre de status Fire, Alarm, Panic, Assistante, Police, Medical             |
| value     | int    | valor del status Fire:1, Alarm:2, Panic:3, Assistance:4, Police:5, Medical:6 |

![FAP](https://github.com/cuervodelviento/driver.sdk_go/assets/60302418/ab922cf2-7208-4d5e-ab0d-351adf881db3)


