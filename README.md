# Palworld API Wrapper
Easy to use wrapper for the Palworld REST API

---

If you don't have a Palworld Server, follow the instructions [here](https://tech.palworldgame.com/)

You can find Palworlds REST API documentation [here](https://tech.palworldgame.com/category/rest-api)

---

## Usage

First you need to initialize a new PalworldAPI instance with credentials.

```go
pal := palworldapi.New("http://localhost:8212", "admin", "admin password")
```

Typically the username is `admin`.

The `AdminPassword` is set in the Server Configuration. ([How to configure the Palworld Server](https://tech.palworldgame.com/settings-and-operation/configuration))


## Examples

You can find some examples [here](https://github.com/Ju0x/palworldapi/tree/main/examples)
