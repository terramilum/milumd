# Setting up Environment

## Installing git

```bash
sudo apt update && sudo apt upgrade && sudo apt install git -y
```

## install go

Get latest version of golang and install it to the machine with below command.

```bash
sudo curl -OL https://golang.org/dl/go1.21.4.linux-amd64.tar.gz && sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.21.4.linux-amd64.tar.gz
```

## setting up environement

1. Open bashrc file

```bash
vim ~/.bashrc
```

2. Set environement varibales copy and past it to the file.

Press insert and paste below parameter then ctrl+c and write :wq for saving and exit.

```bash
# go environment
export GOROOT=/usr/local/go
export PATH=$PATH:/usr/local/go/bin

export GOPATH=/home/code
export PATH=$PATH:/home/code/bin
export GO111MODULE=auto
```

3. Update parameter 

```bash
source ~/.bashrc
```

4. Check go version

```bash
go version
```
- Result : go version go1.21.4 linux/amd64

5. Create required paths.

```bash
sudo mkdir $GOPATH && cd $GOPATH && sudo mkdir src && sudo mkdir bin && sudo mkdir pkg
```

6. Grant user for code directory to write 

```bash
sudo chown -R <your_username> /home/code
```