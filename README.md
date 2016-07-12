# go-status
Simple project to monitoring computer status from browser.

go-status creates a home directory under `~/.gos` in order to store `gos.log`
file.
Templates (css and images) must be under `~/.gos/template`

To sum up `.gos` folder structure:
```
├── ~/.gos/
│   ├── gos.log
│   ├── template/
│   │   ├── index.html
│   │   ├── css/
│   │   │   ├── normalize.css
│   │   │   ├── skeleton.css 
│   │   ├── image/
│   │   │   ├── favicon.png
```

Style and images can be customize as the user want, just follow the template 
parameters in `index.html`

# Thanks to
[gopsutil](https://github.com/shirou/gopsutil)
[skeleton](http://getskeleton.com/)
