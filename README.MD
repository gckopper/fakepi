# FAKEPI

The idea here is to easily create prototypes for APIs. Just create a config file (with yml) and add some sample data.

## Usage

Write a config file like this:

```yml
_path:
  _method:
    file: <file path>
    body: <request body>
    headers:
        _header: <header content>
```

- `_path` should be the url path for that request.
- `_method` should be the http method for that request.
- `<file path>` should be the file path of the response body.
- `<request body>` **OPTIONAL** can be the body (as a string) for that request.
- `_header` **OPTIONAL** (you can have as many as these you want) specifies a required header.
- `<header content>` the required content of the specified header.

Then you just have to run it with your desired flags.

`tester [OPTIONS]`

***

OPTIONS:

- -ip

Change the ip the program will bind to. Default: localhost

- -port

Change the port the program will bind to. Default: 8000

- -config

Especify the path for your config file. All relative file paths specified in the config file are relative to the program current working directory. Default: config.yml

***

## Examples

Example config file in [config.yml](config.yml)
