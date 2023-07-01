

![GitHub repo file count](https://img.shields.io/github/directory-file-count/golangast/gentil) 
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/golangast/gentil)
![GitHub repo size](https://img.shields.io/github/repo-size/golangast/gentil)
![GitHub](https://img.shields.io/github/license/golangast/gentil)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/golangast/gentil)
![Go 100%](https://img.shields.io/badge/Go-100%25-blue)
![status beta](https://img.shields.io/badge/Status-Beta-red)

<h3 align="left">Languages and Tools:</h3>
<p align="left"> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> 

## Gentil
- [Gentil](#gentil)
- [General info](#general-info)
- [Why build this?](#why-build-this)
- [Repository overview](#repository-overview)
- [Setup](#setup)



## General info
This project is a clean little utility package for generating Go programs


## Why build this?
* It is cleaner to have these functions packaged away.





## Repository overview
```
├── ff (files and folders functions)
├── temp (template functions)
├── term (terminal functions)
├── text (text functions)
```

## Setup
Just import it's packages.
```
go get github.com/golangast/gentil
```
To run the functions import the package you want
```
. "github.com/golangast/gentil/utility/ff"
```
Then call the function
```
Filefolder("start", "starter")
```
