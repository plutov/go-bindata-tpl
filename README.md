#### What is go-bindata and why do we need it?

[go-bindata](https://github.com/jteeuwen/go-bindata) converts any text or binary file into Go source code, which is useful for embedding data into Go programs. So you can build your whole project into 1 binary file for easier delivery.

#### html/template

[html/template](https://github.com/arschles/go-bindata-html-template/blob/master/godoc.org/html/template)'s functions `Parse`, `ParseFiles` works only with files on the filesystem, so we need to implement a port to work with both approaches: files or go-bindata. Files:

```
go build && ./go-bindata-tpl

<!DOCTYPE html>
<html lang="en">
<body>
Hello
</body>
</html>
```

#### Generate templates with go-bindata

We need to install go-bindata CLI and generate a .go file from our templates:

```
go get -u github.com/jteeuwen/go-bindata/...
go-bindata -o tpl.go tpl
```

I prefer to add last command to `go:generate`:

```
//go:generate go-bindata -o tpl.go tpl
```

#### Use go-bindata templates

I made it by providing a flag `-go-bindata`:

```
./go-bindata-tpl -go-bindata
<!DOCTYPE html>
<html lang="en">
<body>
Hello
</body>
</html>
```

#### Conclusion