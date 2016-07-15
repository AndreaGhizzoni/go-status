# go-status
Simple project to monitoring computer status from browser.

# How to install
```bash
go get github.com/AndreaGhizzoni/go-status
cd github.com/AndreaGhizzoni/go-status
make
```

At this point will be create a directory called `deploy` that contains:
```
├── deploy/
│   ├── go-status*
│   ├── template/
│   │   ├── index.html
│   │   ├── css/
│   │   │   ├── normalize.css
│   │   │   ├── skeleton.css 
│   │   ├── image/
│   │   │   ├── favicon.png

```

Go-status needs:
- `template` folder is in the same directory of `go-status*` binary
- `template` folder contains *at least* an `index.html` page to serve it as root

Go-status is configured to serve css and image folder ( into template folder )
to let every html template to use some css and image.

# Thanks to
[gopsutil](https://github.com/shirou/gopsutil)

[skeleton](http://getskeleton.com/)
