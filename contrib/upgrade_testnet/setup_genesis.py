import json
import os
import sys
import requests

args = sys.argv[1:]

# Get the args
build_dir = args[0]
genesis_file = f"{build_dir}/node0/desmos/config/genesis.json"

chain_state_url = args[2]
chain_state_file = f"{build_dir}/state.json"
output_file = f"{build_dir}/output_state.json"

# Get the chain state inside the build dir
with requests.get(chain_state_url) as r, open(chain_state_file, 'w') as f:
    f.write(json.dumps(r.json()))

with open(chain_state_file, 'r') as chain_state_f, open(genesis_file, 'r') as genesis_f, open(output_file, 'w') as out:
    chain_state = json.load(chain_state_f)
    genesis = json.load(genesis_f)

    chain_state['genesis_time'] = genesis['genesis_time']
    chain_state['chain_id'] = genesis['chain_id']
    chain_state['initial_height'] = genesis['initial_height']
    chain_state['app_state']['auth']['accounts'] += genesis['app_state']['auth']['accounts']

    # Update the balances coin denom
    for balance in genesis['app_state']['bank']['balances']:
        for coin in balance['coins']:
            if coin['denom'] == 'stake':
                coin['denom'] = 'udaric'

    chain_state['app_state']['bank']['balances'] += genesis['app_state']['bank']['balances']
    for balance in chain_state['app_state']['bank']['balances']:
        # Remove the distribution balance
        if balance['address'] == 'desmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8n8fv78':
            chain_state['app_state']['bank']['balances'].remove(balance)

        # Remove the bonded tokens balance
        elif balance['address'] == 'desmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3prylw0':
            chain_state['app_state']['bank']['balances'].remove(balance)

        # Remove the unbonded tokens balance
        elif balance['address'] == 'desmos1tygms3xhhs3yv487phx3dw4a95jn7t7l4rcwcm':
            chain_state['app_state']['bank']['balances'].remove(balance)

    # -------------------------------
    # --- Fix modules state

    chain_state['app_state']['bank']['supply'] = []

    chain_state['app_state']['distribution'] = genesis['app_state']['distribution']
    chain_state['app_state']['genutil'] = genesis['app_state']['genutil']

    chain_state['app_state']['staking'] = genesis['app_state']['staking']
    chain_state['app_state']['staking']['params']['bond_denom'] = 'udaric'

    chain_state['app_state']['slashing'] = genesis['app_state']['slashing']

    chain_state['app_state']['gov']['deposit_params']['max_deposit_period'] = '120s'
    chain_state['app_state']['gov']['voting_params']['voting_period'] = '120s'
    chain_state['app_state']['gov']['deposit_params']['min_deposit'] = [{'amount': '10000000', 'denom': 'udaric'}]

    # -------------------------------
    # --- Clear the validators list

    del (chain_state['validators'])

    # -------------------------------
    # --- Write the file

    out.write(json.dumps(chain_state))

nodes_amount = args[1]
for i in range(0, int(nodes_amount)):
    genesis_path = f"{build_dir}/node{i}/desmos/config/genesis.json"
    with open(genesis_path, 'w') as file:
        os.system(f"cp {output_file} {genesis_path}")
