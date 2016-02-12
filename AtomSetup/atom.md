Just to list out the steps to configure Atom editor for Go development: 
===

1. Download the .dep package for the deb pacakge from the https://atom.io/download/deb url. 
2. Install the package using the software manger or dpkg command line utility. 
3. You would need to install some packages. Edit > Preferences > Install > "Seach for go-plus". 
4. Once the package is installed you need to set the GOPATH Env. variable, Formating tool and check the option "Run Format tool when file is saved".
5. You would also, need to install the following package from the command line: 
   go get -u github.com/golang/lint/golint

