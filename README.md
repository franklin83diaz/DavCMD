# DavCMD

DavCMD is a lightweight and powerful command line client designed to interact with WebDAV servers. It allows users to upload, download, and upload files to a WebDAV server in an efficient and automated manner.

### Characteristics

Upload files: Send files from your local system to the WebDAV server.
Download files: Recover files from the WebDAV server to your local system.
Command line interface: Facilitates integration with scripts and automations.
Previous requirements
Compatible operating system (Windows, macOS, Linux).
Access to a WebDAV server.

## Facility

### Clone the repository (replace with the actual repository URL)
```
git clone https://github.com/yourusername/DavCMD.git
```

### Change to the project directory
```
cd DavCMD
```

### Run the installation script (if applicable)
```
go build -o davcmd main.go
```

Use
To start using DavCMD, you can use the following commands:

### Upload a file to the WebDAV server
```
davcmd upload -u 5A8iKe53Ta1QnVr -l http://storage.adaptivecomputing.com/public.php/webdav/ test test
```

### Upload a recursive file to the WebDAV server
```
peding...
```

### Download a file from the WebDAV server
```
davcmd download -u 5A8i -l https://s.com/index.php/s/7czpKX2VYuNG test
```

### List files in remote directory
```
davcmd list /path/to/remote/directory/
```
### Specify server
davcmd -d -u your_username -p your_password -s https://example.com/webdav/ /path/to/remote/file.ext /path/to/local/directory/

### Specify Profile from .davcmd
davcmd -d -P personal /path/to/remote/file.ext /path/to/local/directory/


## Contributions

Contributions are welcome. Please submit your pull requests to the main repository.

License
DavCMD is licensed under the MIT License.
