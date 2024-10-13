# Palworld API Wrapper

<div align="center">


<img src="https://github.com/user-attachments/assets/c2693a70-0c25-4308-937f-e5f97137176c" width="240px" height="104px">



Easy to use wrapper for the Palworld REST API

[Set up a Palworld Server](https://tech.palworldgame.com/)

[Read the API Docs](https://tech.palworldgame.com/category/rest-api)
</div>

## Installation

You can install the package by using `go get` in your Go Project.

```
go get -u github.com/Ju0x/palworldapi
```
This will pull the latest release.

## Usage

First you need to initialize a new PalworldAPI instance with credentials.

```go
pal := palworldapi.New("http://localhost:8212", "admin", "admin password")
```

Typically the username is `admin`.

The `AdminPassword` is set in the Server Configuration. ([How to configure the Palworld Server](https://tech.palworldgame.com/settings-and-operation/configuration))


The REST API must be activated on the server with the `RESTAPIEnabled=True` option.


## Examples

You can find some examples [here](https://github.com/Ju0x/palworldapi/tree/main/examples)
