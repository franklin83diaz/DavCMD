# DavCMD

DavCMD es un cliente de línea de comandos ligero y potente diseñado para interactuar con servidores WebDAV. Permite a los usuarios subir, descargar y subir archivos en un servidor WebDAV de manera eficiente y automatizada.

### Características

Subir archivos: Envía archivos desde tu sistema local al servidor WebDAV.
Descargar archivos: Recupera archivos del servidor WebDAV a tu sistema local.
Interfaz de línea de comandos: Facilita la integración con scripts y automatizaciones.
Requisitos previos
Sistema operativo compatible (Windows, macOS, Linux).
Acceso a un servidor WebDAV.

## Instalación

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
go build -o davcmd main.go
```

Uso
Para comenzar a usar DavCMD, puedes utilizar los siguientes comandos:

### Subir un archivo al servidor WebDAV
```
davcmd upload -u 5A8iKe53Ta1QnVr -l http://storage.adaptivecomputing.com/public.php/webdav/ test test
```

### Subir un archivo recursivo al servidor WebDAV
```
peding...
```

### Descargar un archivo del servidor WebDAV
```
davcmd download -u 5A8i -l https://s.com/index.php/s/7czpKX2VYuNG test
```

### Listar archivos en el directorio remoto
```
davcmd list /path/to/remote/directory/
```
### Expecificar server
davcmd -d -u your_username -p your_password -s https://example.com/webdav/ /path/to/remote/file.ext /path/to/local/directory/

### Expecificar Perfil de .davcmd
davcmd -d -P personal /path/to/remote/file.ext /path/to/local/directory/


## Contribuciones

Las contribuciones son bienvenidas. Por favor, envía tus pull requests al repositorio principal.

Licencia
DavCMD está licenciado bajo la MIT License.
