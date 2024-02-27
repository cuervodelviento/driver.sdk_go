# Set Example Config

## Acciones

- set

## Descripcion

Configuracion de ejemplo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "exampleConfig",
        "value": {
            "macAddress": "00:00:00:00:00:00"
        }
    }
}
```



## Response Message
```json
{
    ...,
    "data": {
        "macAddress": "00:00:00:00:00:00",
    }
}
```

| Campo | Tipo | Descripcion |
| --- | --- | --- |
| macAddress | string | La direccion mac del dispositivo |
