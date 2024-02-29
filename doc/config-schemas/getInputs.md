# Get inputs config

## Descripcion

Devolver una lista de las entradas disponibles del dispositivo.

## Request Message

```json
{
    ...,
    "data": {
        "deviceData": {...},
        "configKey": "getInputs",
        "value": "{}"
    }
}
```

El Request Message de `getInputs` no necesita value. El driver solo con el configKey identificara la operacion que debe realizar.

## Response Message

```json
{
    ...,
    "data": [{
        "name": "string",
        "id": "string",
        "description": "string"
    }]
}
```

| Campo       | Tipo   | Descripcion               |
| ----------- | ------ | ------------------------- |
| name        | string | El nombre de la entrada   |
| id          | string | Id interno de la entrada  |
| description | string | Descripción de la entrada |
