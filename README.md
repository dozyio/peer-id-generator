# Peer ID Generator

Generates private/public key pair and associated peer id for use with libp2p.
Uses Ed25519 keys.
Outputs to json.

## Usage

```sh
go run . -c <number of peers>
```

## Example

```sh
go run . -c 3
```

```json
{
  "peers": [
    {
      "private_key": "080112405cebf3a51a9fb21cce71768b4ad4cc62e288abd43dea351d6480e0637228d5a0ecce5afb50dc611e18ec9b680e690b0199ef9e3941cc3d8a5c588110b3621aef",
      "public_key": "08011220ecce5afb50dc611e18ec9b680e690b0199ef9e3941cc3d8a5c588110b3621aef",
      "peer_id": "12D3KooWRkktQw7NTF1tWNdqzcDK6r9StRszP8tqD54Gid41ZY6J"
    },
    {
      "private_key": "08011240d5a99110e399de10e2721c9ab70869924258a12b7331eb6a3571d6ffcd527474aa08c2bfa113fede23eab09bd0d8bf4c9ece7f12b642d98e9b7099cfccef60ae",
      "public_key": "08011220aa08c2bfa113fede23eab09bd0d8bf4c9ece7f12b642d98e9b7099cfccef60ae",
      "peer_id": "12D3KooWMG7E2jUSXH4N73qXN8bmt71xFD9af9SezxR5Wn3zp9MX"
    },
    {
      "private_key": "080112405d14cb03ee5d5e2135294588e5dc8f460218375de1f5b69e63bb292c81c79eee25c69c9f57272340dcc8e08fa7a58a0b3b8b3acde6b6926f7899c7e4178e3178",
      "public_key": "0801122025c69c9f57272340dcc8e08fa7a58a0b3b8b3acde6b6926f7899c7e4178e3178",
      "peer_id": "12D3KooWCMpuZ7q2rzKW1mqAcQXw64178wfbSMBPHirM7Cae8fLB"
    }
  ]
}
```
