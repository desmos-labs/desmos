import json
import os
import requests
import sys
import re

args = sys.argv[1:]

# Get the args
build_dir = args[0]
genesis_file = f"{build_dir}/node0/desmos/config/genesis.json"

genesis_url = args[2]
chain_state_file = f"{build_dir}/state.json"
output_file = f"{build_dir}/output_state.json"

request_pattern = r"""version https://git-lfs.github.com/spec/v1
oid sha256:(?P<hash>.*?)
size (?P<size>.*?)
"""
genesis_file_hash=""
genesis_file_size=0
# Get git lfs info
with requests.get(genesis_url) as r:
    request_match = re.match(request_pattern, r.text)
    genesis_file_hash = request_match.group('hash')
    genesis_file_size = int(request_match.group('size'))

repo_pattern = r"https://raw.githubusercontent.com/(?P<organization>.*?)/(?P<name>.*?)/master/morpheus-apollo-2-7199960.json"
repo_match =  re.match(repo_pattern, genesis_url)
repo_organization = repo_match.group('organization')
repo_name = repo_match.group('name')
chain_state_lfs_url = "https://github.com/{}/{}.git/info/lfs/objects/batch".format(repo_organization, repo_name)
chain_state_lfs_data = {"operation": "download", "transfer": ["basic"], "objects": [{"oid": genesis_file_hash, "size": genesis_file_size}]}
chain_state_lfs_headers = {'content-type': 'application/json', 'accept': 'application/vnd.git-lfs+json'}
chain_state_url=""
# Get chain state url from lfs api
with requests.post(chain_state_lfs_url, headers=chain_state_lfs_headers, data=json.dumps(chain_state_lfs_data)) as r:
    res = r.json()
    chain_state_url= res['objects'][0]['actions']['download']['href']

# Get the chain state inside the build dir
with requests.get(chain_state_url) as r, open(chain_state_file, 'w') as f:
    f.write(json.dumps(r.json()))

with open(chain_state_file, 'r') as chain_state_f, open(genesis_file, 'r') as genesis_f, open(output_file, 'w') as out:
    chain_state = json.load(chain_state_f)
    genesis = json.load(genesis_f)

    # -------------------------------
    # --- Update the bank state
    modules_addresses = [
        'desmos1jv65s3grqf6v6jl3dp4t6c9t9rk99cd8n8fv78',
        'desmos1fl48vsnmsdzcv85q5d2q4z5ajdha8yu3prylw0',
        'desmos1tygms3xhhs3yv487phx3dw4a95jn7t7l4rcwcm',
    ]

    genesis['app_state']['bank']['supply'] = []
    for balance in chain_state['app_state']['bank']['balances']:
        if balance['address'] in modules_addresses:
            balance['coins'] = []

        genesis['app_state']['bank']['balances'].append(balance)

    # -------------------------------
    # --- Fix modules state

    # x/auth
    genesis['app_state']['auth']['accounts'] += chain_state['app_state']['auth']['accounts']

    # x/gov
    genesis['app_state']['gov']['deposit_params']['max_deposit_period'] = '120s'
    genesis['app_state']['gov']['voting_params']['voting_period'] = '120s'
    genesis['app_state']['gov']['deposit_params']['min_deposit'] = [{'amount': '10000000', 'denom': 'udaric'}]

    # -------------------------------
    # --- Copy modules state over
    genesis['app_state']['crisis'] = chain_state['app_state']['crisis']
    genesis['app_state']['ibc'] = chain_state['app_state']['ibc']
    genesis['app_state']['profiles'] = chain_state['app_state']['profiles']

    custom_modules = ['profiles', 'relationships', 'subspaces', 'posts', 'reports', 'reactions', 'fees', 'supply', 'wasm']
    for module in custom_modules:
        if module in chain_state['app_state']:
            genesis['app_state'][module] = chain_state['app_state'][module]
        else:
            del (genesis['app_state'][module])

    # -------------------------------
    # --- Write the file

    out.write(json.dumps(genesis))
    os.system(f"sed -i 's/\"stake\"/\"udaric\"/g' {output_file}")
    os.system(f"sed -i 's/\"udsm\"/\"udaric\"/g' {output_file}")

nodes_amount = args[1]
for i in range(0, int(nodes_amount)):
    genesis_path = f"{build_dir}/node{i}/desmos/config/genesis.json"
    with open(genesis_path, 'w') as file:
        os.system(f"cp {output_file} {genesis_path}")
