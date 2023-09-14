# Mapper

> Mapper based on Laravel-style encapsulation

#### install
```shell
go get -u github.com/goextension/mapper@version
```

#### make mapper
```go
import (
	 mapper2 `github.com/goextension/mapper`
	`github.com/goextension/mapper/contacts`
    )

// make instance a
var a contacts.Mappable[string, string]

a = mapper2.MakeMapper[string, string]()

// make instance b
b := mapper2.MakeMapper[string,string]()
```
