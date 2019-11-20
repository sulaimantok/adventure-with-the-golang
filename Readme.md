# adventure with the golang

### sharing and learning with the Golang programming language

## Get To setup Go in Your Computer

### 1. OSX
#### Using Homebrew 
	```
	$ ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
	```
##### After that
	```
	$ brew install go --cross-compile-common
	```
#### Setup your Path
	```
	$ mkdir $HOME/go
	$ export GOPATH=$HOME/go
	$ open $HOME/.bash_profile
	export GOPATH=$HOME/go
	export PATH=$PATH:$GOPATH/bin
	```
#### Install mercurial and bazaar
	```
	$ brew install hg bzr
	```
### 2. Windows
#### Installer
You can download installer in this site [installer](https://golang.org/dl/)
#### Official Resource
besides the installer file, you can using file in [Official resource for Windows ](https://golang.org/doc/install#windows) to install Go
### 3. Linux
#### Official Package
Install Go from one of the [Official linux Package](https://golang.org/dl/)
For setup the path same with OSX
### 4. Other
#### Useful Tools
Installing a few tools Godoc, vet, and Golint
	```
	$ go get golang.org/x/tools/cmd/godoc
	$ go get golang.org/x/tools/cmd/vet
	$ go get github.com/golang/lint/golint
	```
#### Officail Resource
You can also download in [Official resource](https://golang.org/doc/go1.2#go_tools_godoc)