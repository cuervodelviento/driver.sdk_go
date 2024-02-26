# Pulse action

## Acciones

- action

## Descripcion

Enviar pulsdo para una accion dinamica al driver, ej: 
 
- Cambiar el estado de un rele  
- Activar/Desactivar una alarma
- Levantar una pluma
- Etc.

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "actionOutputPulse",
        "value":{} // no actionOutputPulse 
    }
}
```



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
| error | boolean | Error en el proceso de la accion |
| msg | string | Mensaje de error o log interno|
