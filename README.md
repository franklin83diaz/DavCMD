# DavCMD

DavCMD es un cliente de línea de comandos ligero y potente diseñado para interactuar con servidores WebDAV. Permite a los usuarios subir, descargar, sincronizar y gestionar archivos en un servidor WebDAV de manera eficiente y automatizada.

### Características

Subir archivos: Envía archivos desde tu sistema local al servidor WebDAV.
Descargar archivos: Recupera archivos del servidor WebDAV a tu sistema local.
Sincronización: Mantiene tus archivos locales y remotos sincronizados.
Gestión de archivos: Permite crear, mover, copiar y eliminar archivos y directorios en el servidor WebDAV.
Soporte de bloqueo de archivos: Evita conflictos de edición con el bloqueo de archivos.
Interfaz de línea de comandos: Facilita la integración con scripts y automatizaciones.
Requisitos previos
Sistema operativo compatible (Windows, macOS, Linux).
Acceso a un servidor WebDAV.
Credenciales de acceso al servidor WebDAV (si es necesario).
Instalación
sh
Copy code

### Clonar el repositorio (reemplazar con la URL del repositorio real)
```
git clone https://github.com/yourusername/DavCMD.git
```

### Cambiar al directorio del proyecto
```
cd DavCMD
```

### Ejecutar el script de instalación (si es aplicable)
```
./install.sh
```

Uso
Para comenzar a usar DavCMD, puedes utilizar los siguientes comandos:

### Subir un archivo al servidor WebDAV
```
davcmd -t /path/to/local/file.ext /path/to/remote/directory/
```

### Subir un archivo recursivo al servidor WebDAV
```
davcmd -r -t /path/to/local/file.ext /path/to/remote/directory/
```

### Descargar un archivo del servidor WebDAV
```
davcmd -d /path/to/remote/file.ext /path/to/local/directory/
```

### Listar archivos en el directorio remoto
```
davcmd list /path/to/remote/directory/
```
### Expecificar server
davcmd -d -u your_username -p your_password -s https://example.com/webdav/ /path/to/remote/file.ext /path/to/local/directory/

### Expecificar Perfil de .davcmd
davcmd -d -P personal /path/to/remote/file.ext /path/to/local/directory/

### Config ~/.davcmd
```
[default]
url = https://example.com/webdav/
username = your_username
password = your_password

[personal]
url = https://example.com/webdav/
username = your_username
password = your_password

```
## Contribuciones

Las contribuciones son bienvenidas. Por favor, envía tus pull requests al repositorio principal.

Licencia
DavCMD está licenciado bajo la MIT License.
