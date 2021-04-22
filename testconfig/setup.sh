rly config init

rly chains add -f ibc0.json
rly chains add -f ibc1.json

rly keys add ibc0
rly keys add ibc1

desmos tx bank send test0-0 $(rly chains address ibc0) 10000desmos --chain-id ibc0 --keyring-backend test --home ../build/ibc/ibc0 --node tcp://localhost:26657 --yes
desmos tx bank send test1-0 $(rly chains address ibc1) 10000desmos --chain-id ibc1 --keyring-backend test --home ../build/ibc/ibc1 --node tcp://localhost:26667 --yes

rly light init ibc0 -f
rly light init ibc1 -f

rly paths generate ibc0 ibc1 links --port=links --version=links-1
rly paths generate ibc0 ibc1 transfer --port=transfer

rly tx link links -d -o 3s
rly tx link transfer -d -o 3s

rly start links