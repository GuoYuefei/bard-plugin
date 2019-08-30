# Bard's subprojects ---- bard-plugin
- [Geting Startd](#getting-started)
    + [Install](#install)
    + [Usage](#usage)
- [License](#license)
- [Contributing](#contributing)
- [Preface](#preface)

## Getting Started

### Install

```go
go get -u install github.com/GuoYuefei/bard-plugin
```
If you want to use it as a sub-module of bard, go to the Bard project and run the command to get the sub-module of GIT.
```git
git clone git@github.com:GuoYuefei/bard.git // or git clone https://github.com/GuoYuefei/bard.git

git submodule init
git submodule update
```
If the sub-module has been modified after the last bard submission, the following command is needed to update the sub-module
```git
git submodule update
```

### Usage
If you are in Linux or mac, you should use plugin's construction to get. so files

For Example
```go
go build -buildmode=plugin base/server/*.go
```

If you use it under Windows, you should compile it with the program in the Bard project. Usually only the client needs to do this.


## License
MIT
If the AGPL protocol conflicts with the Bard project, the AGPL protocol is followed.

## Contributing
Welcome contributions.  
Need your support to increase the variability of bard's network traffic.  
I will post the writing specification for the plugin file after the project has stabilized.  
If you are interested in this project now, please contact me.  
email: guoyuefei@protonmail.com  

## Preface
The main project has not yet released version 1.0.0, and there may be major changes in the follow-up.  

