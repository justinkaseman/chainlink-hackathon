# chainlink-hackathon

Proof of concept for a modified Chainlink node that will use Uniswap v1 to turn itself into a Goerli validator.

Once staked as a validator it will return itself to normal node functionality.

### Starting local GETH

```
geth console --ropsten --syncmode light --ws --ipcdisable --wsaddr 0.0.0.0 --wsorigins="\*"
```

### Run chainlink node

- make install
- sudo ./chainlink node start
