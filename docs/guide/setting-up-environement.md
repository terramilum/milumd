# Setting up Environment

## Installing Git

```bash
sudo apt update && sudo apt upgrade && sudo apt install git -y
```

## Install Go

Get the latest version of Golang and install it on the machine with the following command:

```bash
sudo curl -OL https://golang.org/dl/go1.21.4.linux-amd64.tar.gz && sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
```

## Setting up Environment

1. Open the bashrc file:

```bash
vim ~/.bashrc
```

2. Set environment variables by copying and pasting the following lines into the file. Press Insert, paste the parameters, then press Ctrl+C and write :wq to save and exit.

```bash
# go environment
export GOROOT=/usr/local/go
export PATH=$PATH:/usr/local/go/bin

export GOPATH=/home/code
export PATH=$PATH:/home/code/bin
export GO111MODULE=auto
```

3. Update parameters to be effective on the system:

```bash
source ~/.bashrc
```

4. Check the Go version:

```bash
go version
```
- Result: go version go1.21.4 linux/amd64

5. Create the required paths:

```bash
sudo mkdir $GOPATH && cd $GOPATH && sudo mkdir src && sudo mkdir bin && sudo mkdir pkg
```

6. Grant user permissions for the code directory to write:

```bash
sudo chown -R <your_username> /home/code
```