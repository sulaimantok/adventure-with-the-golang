# Adventure with the golang

### Sharing and learning with the Golang programming language

## Get To setup Go in Your Computer
- [OSX](#osx)
- [Windows](#windows)
- [Linux](#linux)
- [Other install(optional)](#other)

### OSX

- Using Homebrew 
	```
	$ ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	```
- After that
	```
	$ brew install go --cross-compile-common
	```
- Setup your Path
	```
	$ mkdir $HOME/go
	$ export GOPATH=$HOME/go
	$ open $HOME/.bash_profile
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
	```
- Install mercurial and bazaar
	```
	$ brew install hg bzr
	```
### Windows

- Installer
You can download installer in this site [installer](https://golang.org/dl/)
- Official Resource
besides the installer file, you can using file in [Official resource for Windows ](https://golang.org/doc/install#windows) to install Go

### Linux
- Official Package
Install Go from one of the [Official linux Package](https://golang.org/dl/)
For setup the path same with OSX

### Other

- Useful Tools
Installing a few tools Godoc, vet, and Golint

	```
	$ go get golang.org/x/tools/cmd/godoc
	$ go get golang.org/x/tools/cmd/vet
	$ go get github.com/golang/lint/golint
	```
- Officail Resource
You can also download in [Official resource](https://golang.org/doc/go1.2#go_tools_godoc)
