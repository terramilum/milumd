# How to Install the `mirumd` Daemon and Full Node on Ubuntu

Follow these steps to install the `mirumd` executable on your Ubuntu system:

## 1. Install `mirumd` Daemon

Run the following command to install the `mirumd` application. Replace `0.1.3` and `1.2.3` with the desired versions of the application and WASM VM.

```bash
bash <(curl -s "https://raw.githubusercontent.com/terramirum/mirumd/main/scripts/install_mirumd.sh") 0.1.3 1.2.3
```
## 2. Verify Installation

After installation, run the following command to reload your environment and check the installed version of `mirumd`:

```bash
source "$HOME/.bash_profile" && mirumd version
```

This will confirm that the mirumd daemon is successfully installed and working.

# 3. Install the Mirum Network Full Node

To install the Mirum full node, use the Bash command below. Remember to replace node1 with your preferred moniker name, which identifies your node in the network.

```bash
bash <(curl -s "https://raw.githubusercontent.com/terramirum/mirumd/main/scripts/install_full_node_mainnet.sh") node1
```

Make sure to update node1 to a name that makes sense for your setup.

# 4. Create a Wallet and Transfer some validator coins

