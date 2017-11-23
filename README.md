# install

## gvm

- version managemt tool: https://github.com/moovweb/gvm
- install go via gvm: https://gist.github.com/d2s/6503f815431d1587c28bc37bfd715dbf
- set package env
```
mkdir -p $HOME/hoge/{pkg,bin,src}
gvm pkgset create hoge

gvm pkgenv hoge
# line12
export GOPATH; GOPATH="/Users/nghinv/.gvm/pkgsets/go1.7/hoge:$HOME/hoge:$GOPATH"
# line16
export PATH; PATH="/Users/nghinv/.gvm/pkgsets/go1.7/hoge/bin:${GVM_OVERLAY_PREFIX}/bin:$HOME/hoge/bin:${PATH}"

gvm pkgset use hoge
```

## glide

-package management tool
- https://github.com/Masterminds/glide
- warning: you should edit install shell about $GOPATH
```
curl https://glide.sh/get > glide.sh
vim glide_install.sh

GOPATH='~/hoge'
$GOPATH -> ${GOPATH}

sh glide_install.sh
```

## revel

- install
```
$ go get github.com/revel/cmd/revel
```
- setup $ run
```
$ revel new github.com/*user*/*myproj*
$ revel run github.com/*user*/*myproj*
```

- glide.yaml
```
$ cd src/github.com/*user*/*myproj*
$ glide create
$ glide install
```

## GORM

- install
```
go get -u github.com/jinzhu/gorm
'''

