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
bash <(curl -s "https://raw.githubusercontent.com/terramirum/mirumd/main/scripts/install_full_node_mainnet.sh") <moniker>
```

Make sure to update '<moniker>' to a name that makes sense for your setup.

# 4. Create a Wallet and Transfer validator coins

## Step 1: Create the Validator Wallet

To create a new validator wallet, use the following Bash script. This script checks if a wallet already exists, and if not, creates one using a password that you pass as an argument or environment variable.


```bash
bash <(curl -s "https://raw.githubusercontent.com/terramirum/mirumd/main/scripts/create_wallet.sh") <password>
```

Not: Replace '<password>' with secure saved password.

## Step 2: Transfer Coins to the Validator Wallet
Once you have created the validator wallet, you'll need to transfer some coins to it from an exchange or another wallet. You can use the following command to check your wallet address:
```bash
mirumd keys show validator
```

## Step 3: Check Your Wallet Balance
To check your balance, use the following command:

```bash
mirumd query bank balances <wallet_address>
```

## Next Steps: Becoming a Validator

```bash
bash <(curl -s "https://raw.githubusercontent.com/terramirum/mirumd/main/scripts/create_validator.sh") <moniker>
```
After validator is created, Please check it block explorer execute this transaction

Not: Replace '<moniker>' with secure saved password.

### Key Points:
- This Markdown document outlines the process of creating the wallet, transferring coins, checking the balance, and securing the mnemonic phrase.
- The Bash script provided ensures secure wallet creation with proper password handling.
- The document provides clear instructions on how to execute the script and proceed to the next steps of becoming a validator.
