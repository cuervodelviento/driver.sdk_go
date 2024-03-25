# Get ftp info

## Descripcion

Obtener la informacion de la configuracion ftp del dispositivo

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getFtpInfo",
        "value": "{}"  // No necesita
    }
}
```



## Response Message
```json
{
    ...,
    "data": [{
    "id": "abc123",
    "host": "1.1.1.1",
    "user": "admin",
    "directory": "/directory",
    "port": 21,
    "encryption": "plain",
    "mode": "passive"
    }]
}
```



| Campo | Tipo | Descripcion |
| --- | --- | --- |
| id | string | uuid |
| host | string | ip del servidor ftp |
| user | string | usuario ftp |
| directory | string | directorio en cual se desea almacenar guardar |
| port | int | puerto del servidor ftp |
| encryption | select | tipo de encriptacion: [plain, explicit, implicit] |
| mode | select | modo de conexion: [ommited, active, passive] |


