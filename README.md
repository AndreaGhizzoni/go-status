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
- `template` folder contains **at least** an `index.html` page to serve it as
  root

Go-status is configured to serve css and image folder ( into template folder )
to let every html template to use some css and image.

# Configuration
Go-status create it's home directory in `/home/andrea/`. Now that is my
configuration on my machines, to modify this condition edit the file
`constant/constant.go` and change the following part:
```go
const (
    ...
    AppHome = "/some/other/path"
    ...
)
```
I did this because panic error occurs when run `user.Current()` method on linux
arm.

Inside of it will be a file called `gos.json` that looks like this:
```json
{ "wp" : [] }
```
`wp` attribute stands for *Watched Paths* that will be displayed in the Disk
section of default template (index.html)

### How do I add a new path ?
Just add some path separated by comma, for example:
```json
{ "wp" : ["/some/path", "/other/path"] }
```

# Thanks to
[gopsutil](https://github.com/shirou/gopsutil)

[skeleton](http://getskeleton.com/)
