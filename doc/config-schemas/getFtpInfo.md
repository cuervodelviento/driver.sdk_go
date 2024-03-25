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

## Encryption

### Plain 
FTP sin cifrar, o FTP plano, es el modo más básico de FTP en el que no se utiliza ninguna forma de cifrado para proteger la información transmitida. En este modo, tanto los comandos como los datos, incluidos los nombres de usuario y las contraseñas, se transmiten en texto plano, lo que significa que son vulnerables a la interceptación y al robo de datos. Para este modo se utiliza el puerto 21.

### Explicit
En el modo explícito, la seguridad se establece después de que el cliente se conecte al servidor FTP, pero antes de enviar cualquier información de inicio de sesión o datos de archivo. Esto significa que el cliente y el servidor establecen primero una conexión no cifrada y, a continuación, el cliente debe enviar un comando específico (como "AUTH TLS" o "AUTH SSL") para solicitar seguridad. Una vez establecida la seguridad, toda la comunicación se cifra mediante SSL/TLS. El puerto estándar para FTP explícito es el puerto 21.

### Implicit
En modo implícito, la seguridad se establece automáticamente en cuanto se inicia la conexión con el servidor FTP. Esto se hace normalmente utilizando SSL/TLS (Secure Sockets Layer/Transport Layer Security) para encriptar toda la comunicación entre el cliente y el servidor. El puerto estándar para FTP implícito es el puerto 990. Este modo se utiliza menos que el explícito, pero algunos servidores aún lo admiten.

## Mode
### Active
En el modo activo, el cliente FTP abre un puerto de datos local (puerto efímero) y envía un comando PORT al servidor FTP para informarle del puerto de datos que ha abierto. A continuación, el servidor FTP inicia la conexión de datos desde su puerto de control (puerto 20 por defecto) al puerto de datos especificado por el cliente. En este modo, es el cliente el que escucha las conexiones entrantes del servidor.

El modo activo puede ser problemático en algunos entornos de red, ya que implica que el servidor intentará conectarse al cliente, lo que puede estar bloqueado por cortafuegos o routers que restrinjan las conexiones entrantes. Sin embargo, este modo es útil en situaciones en las que el cliente tiene un cortafuegos y no el servidor.

## Passive
En el modo pasivo, la iniciativa de la conexión de datos la toma el cliente FTP. El cliente envía un comando PASV al servidor, solicitándole que abra un puerto de datos y le informe de la dirección y el puerto en el que está escuchando. Una vez que el servidor responde con una dirección IP y un puerto, el cliente inicia la conexión de datos con el servidor.

El modo pasivo suele preferirse en entornos en los que existen restricciones de cortafuegos o NAT (Network Address Translation) que dificultan las conexiones entrantes al cliente. Además, el modo pasivo puede ser más seguro, ya que sólo el cliente inicia las conexiones de datos.

## Omitted
En este modo, el servidor FTP no especifica el modo de conexión, por lo que el cliente debe determinar el modo de conexión a utilizar. Este modo es menos común y no se utiliza en la mayoría de los servidores FTP.
