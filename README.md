# Palworld API Wrapper
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/96c7fd4725fa49c88121c7e1af158899)](https://app.codacy.com/gh/Ju0x/palworldapi/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)


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


### First Hello World

With the palworldapi instance you are now able to use the wrapper functions.

```go
package main

import "github.com/Ju0x/palworldapi"

func main() {
    pal := palworldapi.New("http://localhost:8212", os.Getenv("USERNAME"), os.Getenv("ADMIN_PASSWORD"))

    // Sends the message to the server globally
    pal.Announce("Hello World!")
}
```

Then run it
```
USERNAME=admin ADMIN_PASSWORD=your_admin_password go run .
```

The result should look like this:

<img src="https://github.com/user-attachments/assets/d2d986da-5295-4a3d-ad83-57ea1cd058d4" height="48px">


## Examples

You can find some examples [here](https://github.com/Ju0x/palworldapi/tree/main/examples)
