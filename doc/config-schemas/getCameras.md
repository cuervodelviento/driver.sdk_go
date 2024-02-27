# Get channels config

## Descripcion

Devolver una lista de las camaras que estan conectadas al dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getCameras",
        "value": "{}"
    }
}
```

El Request Message de `getCameras` no necesita value. El driver solo con el configKey identificara la operacion que debe realizar.

## Response Message

```json
{
    ...,
    "data": [{
        "name": "string",
        "channelNumber": 1,
        "id":"string"
    }]
}
```

| Campo         | Tipo   | Descripcion               |
| ------------- | ------ | --------------------------|
| name          | string | El nombre de la camara    | 
| channelNumber | number | El numero del canal       |
| id            | string | El id interno de la camara|
