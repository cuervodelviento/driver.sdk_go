# Get recording source config

## Descripcion

Devuelve la url RTSP de una grabación incluyendo tiempo de inicio y tiempo final opcional

## Request Message

```json
{
    ...
    "data": {
        "deviceData": {...},
        "configKey": "getRecordingSource",
        "value": {
            "channelId" :"111-aaa",
            "utcStart" : "2024-07-12T14:40:10Z",
            "utcEnd" :"2024-07-12T16:30:25Z",
        }
    }
}
```

| Campo     | Tipo   | Descripción                                                   |
| --------- | ------ | ------------------------------------------------------------- |
| channelId | string | id del canal de video                                         |
| utStart   | string | debe seguir el formato UTC "YYYY-MM-DDTHH:mm:ssZ"             |
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

utcEnd sigue el mismo formato pero es opcional

## Response Message

```json
{
    ...,
    "data": {
        "rtspSource": "rtsp://user:password@10.1.1.23:553/playback/device/4/channel/2/media.smp?deviceQuery=23440t&__RangeHeader=clock=20240605T041500Z-20240605T041510Z"
    }
}
```

| Campo      | Tipo   | Descripcion              |
| ---------- | ------ | ------------------------ |
| rtspSource | string | La fuente rtsp del canal |

Nota sobre la fuente RTSP

en base a este ejemplo

```
rtsp://user:password@10.1.1.23:553/playback/device/4/channel/2/media.smp?deviceQuery=23440t&\_\_RangeHeader=clock=20240605T041500Z-20240605T041510Z
```

user:password son las credenciales para acceso al video por RTSP, pueden ser diferentes a las credenciales de acceso al dispositivo

ip y puerto del dispositivo: el puerto puede ser diferente al puerto de integración del dispositivo

path: varía entre dispositivos, puede llevar id de camara y canal

deviceQuery: son parametros propios del dispositivo, puede incluir inicio de video, formato, canal etc.

\_\_EngineQuery: son parametros ajenos al dispositivo y comienzan con \_\_ pero que son utilizados por el motor de video, son necesarios cuando el control de la reproducción se maneja a través del protocolo RTSP y no hay un reemplazo disponible a través de una deviceQuery. normalmente añaden valores al header de una petición RTSP (Puedes aprender más [aquí](https://www.rfc-editor.org/rfc/rfc2326.html#:~:text=835%0A%20%20%20%20%20%20%20%20%20%20%20Session%3A%2012345678-,Range%3A%20npt%3D10%2D15,-C%2D%3ES%3A%20PLAY) ) puede ser Range, Scale, Date según lo soporte el dispositivo y si ha sido integrado al servicio de video
