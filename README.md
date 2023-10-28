# Config2FS

## Description
Config2FS is a CLI tool written in Go that allows you to generate file system structures based on YAML configuration files. It provides an easy way to scaffold out new projects or generate commonly used directory and file structures.

## Features
- Generate folders and files from a simple YAML configuration
- Cross-platform support (Windows, macOS, Linux)
- Easy to extend and integrate into existing workflows

## Installation
1. Clone this repository.
2. Navigate to the project directory.
3. Run `go build` to compile the project.
4. (Optional) Add the compiled binary to your PATH.

## Usage
1. Create a YAML file describing your desired file system structure.
2. Run the tool using:
    ```
    go run main.go --config_path=/path/to/your/config.yaml
    ```

## Example Config
Here's an example YAML configuration:

```yaml
project_name: template-project
project_type: monorepo
project_path: /your/directory/path
project_structure:
  file:Dockerfile: []
  file:Makefile: []
  cmd:
    handlers:
        - file:email_collector_handler.go
        - file:sms_collector_handler.go
        - file:main.go
file:go.mod: []
```

## Contributing
Feel free to open issues or PRs if you have suggestions or bug reports.

## License
MIT
# config2fs
