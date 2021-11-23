import json
import os
import sys
import requests
import hashlib
import base64
import bech32

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

    # Transform the gentxs into validators
    genesis_validators = []
    genesis_delegations = []
    tendermint_validators = []
    for gentx in genesis['app_state']['genutil']['gen_txs']:
        message = gentx['body']['messages'][0]
        genesis_validators.append({
            'commission': {
                'commission_rates': message['commission'],
                'update_time': genesis['genesis_time']
            },
            'consensus_pubkey': message['pubkey'],
            'delegator_shares': message['value']['amount'],
            'description': message['description'],
            'jailed': False,
            'min_self_delegation': message['min_self_delegation'],
            'operator_address': message['validator_address'],
            'status': 'BOND_STATUS_BONDED',
            'tokens': message['value']['amount'],
            'unbonding_height': '0',
            'unbonding_time': '1970-01-01T00:00:00Z'
        })
        genesis_delegations.append({
            'delegator_address': message['delegator_address'],
            'validator_address': message['validator_address'],
            'shares': message['value']['amount']
        })

        pubkey = base64.b64decode(message['pubkey']['key'].encode('ascii'))
        alg = hashlib.sha256()
        alg.update(pubkey)

        tendermint_validators.append({
            'address': alg.digest()[:20].hex(),
            'name': message['description']['moniker'],
            'power': message['value']['amount'],
            'pub_key': {
                'type': 'tendermint/PubKeyEd25519',
                'value': message['pubkey']['key']
            }
        })

    del (genesis['app_state']['genutil'])

    # -------------------------------
    # --- Update the staking state
    for validator in chain_state['app_state']['staking']['validators']:
        validator['status'] = 'BOND_STATUS_UNBONDED'

    chain_state['app_state']['staking']['validators'] = genesis_validators

    # Change all the delegations so that they are delegated to the first validator
    changed_validators = []
    added_shares = 0.0
    validator_address = genesis_validators[0]['operator_address']
    for delegation in chain_state['app_state']['staking']['delegations']:
        added_shares += float(delegation['shares'])
        changed_validators.append(delegation['validator_address'])

        delegation['validator_address'] = validator_address

    # Update the delegator shares:
    # - set the old validators to 0
    # - add the added_shares amount to the first validator
    for changed_validator in changed_validators:
        for validator in chain_state['app_state']['staking']['validators']:
            if validator['operator_address'] == validator_address:
                validator['delegator_shares'] = '20017455415180.588050795495252068'

            if validator['operator_address'] == changed_validator:
                validator['delegator_shares'] = '0'

    chain_state['app_state']['staking']['delegations'] += genesis_delegations
    del (chain_state['app_state']['staking']['redelegations'])

    last_validator_powers = []
    for validator in genesis_validators:
        last_validator_powers.append({
            'address': validator['operator_address'],
            'power': validator['tokens']
        })
    chain_state['app_state']['staking']['last_validator_powers'] = last_validator_powers

    # -------------------------------
    # --- Update the distribution state
    genesis_initial_distribution = []
    genesis_validator_historical_rewards = []
    for delegation in genesis_delegations:
        genesis_initial_distribution.append({
            'delegator_address': delegation['delegator_address'],
            'starting_info': {
                'height': genesis['initial_height'],
                'previous_period': '1',
                'stake': delegation['shares']
            },
            'validator_address': delegation['validator_address']
        })
        genesis_validator_historical_rewards.append({
            'period': '0',
            'rewards': {
                'cumulative_reward_ratio': [],
                'reference_count': 1,
            },
            'validator_address': delegation['validator_address']
        })

    chain_state['app_state']['distribution']['delegator_starting_infos'] = genesis_initial_distribution
    chain_state['app_state']['distribution']['validator_historical_rewards'] = genesis_validator_historical_rewards
    chain_state['app_state']['distribution']['outstanding_rewards'] = []
    chain_state['app_state']['distribution']['validator_current_rewards'] = []
    chain_state['app_state']['distribution']['validator_accumulated_commissions'] = []

    # Update the distribution starting info
    for changed_validator in changed_validators:
        for distribution_info in chain_state['app_state']['distribution']['delegator_starting_infos']:
            if distribution_info['validator_address'] == changed_validator:
                distribution_info['validator_address'] = validator_address

    # -------------------------------
    # --- Update the bank state
    delegated_amount = 0
    for delegation in genesis_delegations:
        delegated_amount += int(delegation['shares'])

    chain_state['app_state']['bank']['balances'] += genesis['app_state']['bank']['balances']
    for balance in chain_state['app_state']['bank']['balances']:
        # Remove the distribution balance
        if balance['address'] == 'desmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8n8fv78':
            balance['coins'][0]['amount'] = '1665300627184'

        # Remove the bonded tokens balance
        if balance['address'] == 'desmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3prylw0':
            balance['coins'][0]['amount'] = str(delegated_amount)

        # Remove the unbonded tokens balance
        elif balance['address'] == 'desmos1tygms3xhhs3yv487phx3dw4a95jn7t7l4rcwcm':
            balance['coins'][0]['amount'] = '630839623447'

    # -------------------------------
    # --- Update the slashing state

    signing_infos = []
    for idx, validator in enumerate(tendermint_validators):
        signing_info = chain_state['app_state']['slashing']['signing_infos'][idx]
        data =  bytearray.fromhex(validator['address'])
        cons_addr = bech32.bech32_encode('desmosvalcons', bytes(data))
        signing_infos.append({
            'address': cons_addr,
            'validator_signing_info': {
                "address": cons_addr,
                'index_offset': signing_info['index_offset'],
                'jailed_until': signing_info['jailed_until'],
                'missed_blocks_counter': signing_info['missed_blocks_counter'],
                'start_height': signing_info['start_height'],
                'tombstoned': signing_info['tombstoned']
            }
        })

    chain_state['app_state']['slashing'] = genesis['app_state']['slashing']
    chain_state['app_state']['slashing']['signing_infos'] = signing_infos

    # -------------------------------
    # --- Fix modules state

    chain_state['app_state']['bank']['supply'] = []

    chain_state['app_state']['gov']['deposit_params']['max_deposit_period'] = '120s'
    chain_state['app_state']['gov']['voting_params']['voting_period'] = '120s'
    chain_state['app_state']['gov']['deposit_params']['min_deposit'] = [{'amount': '10000000', 'denom': 'udaric'}]

    # -------------------------------
    # --- Clear the validators list

    chain_state['validators'] = tendermint_validators

    # -------------------------------
    # --- Write the file

    out.write(json.dumps(chain_state))
    os.system(f"sed -i 's/udsm/udaric/g' {output_file}")

nodes_amount = args[1]
for i in range(0, int(nodes_amount)):
    genesis_path = f"{build_dir}/node{i}/desmos/config/genesis.json"
    with open(genesis_path, 'w') as file:
        os.system(f"cp {output_file} {genesis_path}")
