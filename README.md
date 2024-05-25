# Health Check CLI

## Overview
This is a simple command-line interface (CLI) application written in Go that checks the health of a system. It uses the `urfave/cli` package to handle command-line arguments and flags.

## Installation
To install this application, you need to have Go installed on your machine. Once Go is installed, you can clone this repository and build the application.

```bash
git clone <repository-url>
cd <repository-directory>
go build -o healthcheck
```

## Usage
The application requires the domain name to check. The port is optional and defaults to 80 if not provided.

```bash
./healthcheck --domain example.com --port 8080
```

or using short flags:

```bash
./healthcheck -d example.com -p 8080
```

The application will output the health status of the provided domain on the specified port.

## Flags
- `--domain` or `-d`: The domain name to check. This flag is required.
- `--port` or `-p`: The port to check. This flag is optional.

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

The available command-line flags for the Health Check CLI are:

- `--domain` or `-d`: This flag is used to specify the domain name to check. This flag is required.
- `--port` or `-p`: This flag is used to specify the port to check. This flag is optional. If not provided, the default value is "80".

## License
[MIT](https://choosealicense.com/licenses/mit/)