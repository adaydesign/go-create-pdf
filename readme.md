# Go create PDF file

### Library
- gofpdf
- gofpdi

link : github.com/phpdave11/gofpdf

- [Fix bug] file size increase - 4/8/2021

### Run
```
$ go run main.go
```

### Test by API

- server run on port :3030
- test url : (can run REST-Client in test.http)

```
GET http://localhost:3030/1
```

### Information
- change pdf template and font in pdf-template folder
- when the program has been run finished, pdf file output will create in pdf-output folder