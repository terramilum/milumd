# Installing Full Node

After running the ./scripts/install_full_node_mainnet.sh bash file, the node is installed. Start the node by using the 'mirumd start' command. Next, create a wallet, transfer a certain amount of validator coins, and utilize the 'validator create' bash script.

1. **Add a wallet to validator node**

- List wallet in the server
  
```base
mirumd keys list
```

If no wallet exist, result is empty.

```bash
No records were found in keyring
```

After creating wallet like below.

```bash
Enter keyring passphrase (attempt 1/3):
- address: mirum19q32ge2795jthnvec3s6m93j8hk9hk0aafnfvv
  name: validatorwallet
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aq/yV9/jekluF/Py+W89eYRMEshz62a3/tbnQcCQi2qA"}'
  type: local
```

- Add a wallet

Create a new wallet with a robust password and ensure its secure storage.

```base
mirumd keys add validatorwallet
```

Result:

```base
Enter keyring passphrase (attempt 1/3):
Re-enter keyring passphrase:

- address: mirum19q32ge2795jthnvec3s6m93j8hk9hk0aafnfvv
  name: validatorwallet
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aq/yV9/jekluF/Py+W89eYRMEshz62a3/tbnQcCQi2qA"}'
  type: local


**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

position exit rabbit frequent point dinosaur cruel security curve mule provide profit dragon true tattoo absorb brass fatigue capable fever diary seek slow view

```

- Delete existing wallet

```bash
mirumd keys delete validatorwallet
```

Result

```bash
Enter keyring passphrase (attempt 1/3):
Key reference will be deleted. Continue? [y/N]: y
Key deleted forever (uh oh!)
```

**Note:** Before initiating the chain, make sure to securely record the mnemonic phrase of your validator wallet.

**PASSWORD:** The wallet's private key is stored in the system keyring and requires a password for access. Choose a robust password and store it securely.

1. **Deposit Trm to validador node**

Transfer a sum of money to your wallet from various sources such as an SD exchange, non-custodial wallet, and others.

Sample of a funds transfer by using below command.

```bash
mirumd tx bank send validatorwallet mirum1wlugx30qc4zmc32xy07tpc6hslectra8wghqgf 800000000000mirum --fees 2mirum --chain-id mirum-1
```

Retrieve the transaction hash (txhash) and verify its success or failure on the block explorer. A sample txhash looks like this: 1BA7CC62A6C4327FEE7295913013492ADBA25492ED50779DC7377364E95FFF61.
2. **Check Balance of money**

```bash
mirumd query bank balances mirum19q32ge2795jthnvec3s6m93j8hk9hk0aafnfvv
```

Result:

The amount is expressed with six decimal places. To obtain the actual balance, divide the amount by 10^6.

```bash
balances:
- amount: "1000000000000"
  denom: MIRUM
pagination:
  next_key: null
  total: "0"
```

3. **Create Validator**


After confirming the availability of the required validator coins and sufficient gas coins, use the following command to create a validator:

```bash
mirumd tx staking create-validator \
  --amount=1000000000000MIRUM \
  --pubkey=$(mirumd tendermint show-validator) \
  --moniker="validator" \
  --chain-id=mirum-1 \
  --commission-rate="0.10" \
  --commission-max-rate="0.20" \
  --commission-max-change-rate="0.01" \
  --min-self-delegation="1000000000000" \
  --gas="254246" \
  --gas-prices="0.01MIRUM" \
  --from=validatorwallet  
```

This command initiates the creation of a validator with specified parameters, such as the amount, public key, moniker, commission rates, self-delegation, gas, gas prices, and the wallet used for the transaction.

**Note:** Verify the transaction hash on the explorer to confirm its success or failure.