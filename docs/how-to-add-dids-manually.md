# How to add DIDs and keys manually

### 1. Create DID using Cheqd
1.1. Follow this [cheqd doc](https://github.com/cheqd/identity-docs/blob/main/tutorials/dids/cheqd-cosmos-cli/create-did-and-did-document.md) to create dids with corresponding private keys

### 2. Add DID methods' private keys using Aries Web Wallet API
2.1. Make sure have running aries rest agents
 - Use the following command to run issuer/holder/verifier aries agents locally
    ```bash
    make run-aries-agents
    ```
2.2 Create a wallet profile
- Example using curl
    ```bash
    curl --location --request POST 'http://localhost:8082/vcwallet/create-profile' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "userId": "d0001",
        "localKMSPassphrase": "Hp6GF4Wg6PPL"
    }'
    ```
- `userId` - user id of profile being created
- `localKMSPassphrase` - passphrase

2.3 Open a wallet using credentials from section (2.2) and copy `token` from response
- Example using curl
    ```bash
    curl --location --request POST 'http://localhost:8082/vcwallet/open' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "userId": "d0001",
        "localKMSPassphrase": "Hp6GF4Wg6PPL"
    }'
    ```
- Example curl response
    ```bash
    {"token":"b111b945fb843412fe2a81889563ff75753f5c9010fdeafa87e83751c809595e"}
    ```

- `userId` - user id of profile being created
- `localKMSPassphrase` - passphrase

2.4 Add DID method private key to a wallet
- Here is an example did document:
    ```json
    {
        "@context": [
            "https://www.w3.org/ns/did/v1"
        ],
        "id": "did:cheqd:testnet:zChwoA37NUMV4ZsURH1kzQ5UzJbPGn2L",
        "verificationMethod": [
            {
                "controller": "did:cheqd:testnet:zChwoA37NUMV4ZsURH1kzQ5UzJbPGn2L",
                "id": "did:cheqd:testnet:zChwoA37NUMV4ZsURH1kzQ5UzJbPGn2L#key-73rtyj8ebx",
                "publicKeyMultibase": "zChwoA37NUMV4ZsURH1kzQ5UzJbPGn2LzNvHry2YuJmS9",
                "type": "Ed25519VerificationKey2020"
            }
        ],
        "authentication": [
            "did:cheqd:testnet:zChwoA37NUMV4ZsURH1kzQ5UzJbPGn2L#key-73rtyj8ebx"
        ],
        "assertionMethod": [
            "did:cheqd:testnet:zChwoA37NUMV4ZsURH1kzQ5UzJbPGn2L#key-73rtyj8ebx"
        ]
    }
    ```
- Example using curl
    ```bash
    curl --location --request POST 'http://localhost:8082/vcwallet/add' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "auth": "<token received after opening wallet>",
        "content": {
            "id": "<verification method id from DID document>",
            "type": "ed25519verificationkey2018",
            "privateKeyBase58": "<base58 encoded private key corresponding to verification method public key>"
        },
        "contentType": "key",
        "userID": "<wallet user id>"
    }'
    ```

2.5 Repeat the steps above for all public private key pairs