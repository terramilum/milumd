# Installing Nodes

Before starting the installation of a Node, ensure that your environment is set up for compiling Go code. Refer to the "setting-up-environment" documentation for detailed instructions.

## Getting the Source Code

- Obtain the code:

```bash  
go get github.com/terramirum/mirumd
```
If you encounter an error like "package github.com/terramirum/mirumd: no Go files in /home/code/src/github.com/terramirum/mirumd," navigate to its directory and use Git clone:

```bash
cd $GOPATH/src/github.com/terramirum/mirumd
```

If the code is not obtained using go get, use Git clone:

```bash
git clone https://github.com/terramirum/mirumd.git
```

## Compiling Source Code

```bash
sudo apt install make -y && sudo apt install gcc -y && make build
```

Check Terramirum chain application version to confirm a successful installation:

```bash
mirumd version
```
## Installing Node

- To begin the Node installation, navigate to the "scripts" folder to run the bash file.

```bash
cd $GOPATH/src/github.com/terramirum/mirumd/scripts
```

- Grant execute permission to the .sh file.

```bash
chmod 777 install_full_node_mainnet.sh
```

- Run the executable file to install the chain.

Before executing the chain, modify its parameters to prevent conflicts with other chains:

MONIKER: Node name visible on the block explorer if your validator approves a block.

```bash
./install_full_node_mainnet.sh
```

- Start Chain


```bash
mirumd start
```

To start it in the background, use nohup:

```bash
nohup mirumd start &
```
