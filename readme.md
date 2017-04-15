# hkbridge

Bridge for group all devices under single HC process

## Using

1. Create configuration file in `main` directory or directory that you're going to run main.go
```json
{
  "name": "MyBridge",
  "manufacturer": "Your own manufacturer",
  "serial": "Magic Serial Number",
  "model": "Super Awesome HomeControl bridge",
  "pin": "12345678",
  "port": "51826",
  "debug": true,
  "accessories": [
    {
      "type": "github.com/llun/hksoundtouch"
    },
    {
      "type": "github.com/llun/hkwifioccupancy",
      "option": {
        "file": "/tmp/presence.wifi",
        "addresses": [
          "mac address"
        ]
      }
    },
    {
      "type": "github.com/llun/hksensibo",
      "option": {
        "key": "sensibo-api-key"
      }
    }
  ]
}

```
2. Run main.go
```sh
go run main/main.go
```

## Building

```sh
make && ./main/homekit
```

Currently target to ARMv5 for using with WPQ864

## License
MIT
