# Mapper
> Mapper based on Laravel-style encapsulation

#### install
```go
go get -u github.com/goextension/mapper
```

#### simple 
```go

var m contacts.Mappable[string, int]

m = MakeMapper[string, int]()

m.Keys()

// 
mapper := MakeMapper[string,any]()

mapper.Values()

```