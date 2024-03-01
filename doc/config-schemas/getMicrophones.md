# Get microphones config

## Descripcion

Devolver una lista de los microfonos que estan registrados en el dispositivo.

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getMicrophones",
        "value": "{}"
    }
}
```

El Request Message de `getMicrophones` no necesita value. El driver solo con el configKey identificara la operacion que debe realizar.

## Response Message

```json
{
    ...,
    "data": [{
        "name": "string",
        "endabled": true,
        "id": "string"
    }]
}
```

| Campo   | Tipo    | Descripcion              |
| ------- | ------- | ------------------------ |
| name    | string  | El nombre del micrófono  |
| enabled | boolean | está o no habilitado     |
| id      | string  | Id interno del micrófono |
