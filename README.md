# go-status
Simple project to monitoring computer status from browser.

# How to install
```bash
go get github.com/AndreaGhizzoni/go-status
cd github.com/AndreaGhizzoni/go-status
make
```

At this point will be create a directory `deploy` in the project directory that 
contains:
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

Go-status NEEDS that the folder `template` is in it's same directory in order to
serve html page.

It's important that THERE IS at least an html template (inde.html) into template
folder in order to let go-status to serve it as root.

Go-status is configured to serve css and image folder ( into template folder )
to let every html template to use some css and image.

# Thanks to
[gopsutil](https://github.com/shirou/gopsutil)

[skeleton](http://getskeleton.com/)
