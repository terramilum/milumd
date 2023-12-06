# Installing Nodes

Before starting to install node, Please check and make ready your environment for compiling go code.
Check the setting-up-environment documentation.

## Getting source code

- Get codes

```bash  
go get github.com/terramirum/mirumd
```
if getting error like package github.com/terramirum/mirumd: no Go files in /home/code/src/github.com/terramirum/mirumd
go to its directory and user git clone.

```bash
cd $GOPATH/src/github.com/terramirum/mirumd
```

if codes is not get by go get, then use git clone

```bash
git clone https://github.com/terramirum/mirumd.git
```

## Compiling Source Code

```bash
sudo apt install make -y && sudo apt install gcc -y && make build
```

Checking version of Terramirum chain code to be sure successfully installed.

```bash
mirumd version
```