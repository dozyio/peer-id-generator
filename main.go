package main

import (
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/peer"
)

type KeysPeerID struct {
	PrivKey string `json:"private_key"`
	PubKey  string `json:"public_key"`
	PeerID  string `json:"peer_id"`
}

type KeysPeers struct {
	Peers []KeysPeerID `json:"peers"`
}

func main() {
	c := flag.Int("c", 1, "number of keys to generate")

	flag.Parse()

	keysPeers := KeysPeers{
		Peers: []KeysPeerID{},
	}

	for i := 0; i < *c; i++ {
		keysPeers.Peers = append(keysPeers.Peers, generate())
	}

	out, err := json.MarshalIndent(keysPeers, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", out)
}

func generate() KeysPeerID {
	priv, pub, err := crypto.GenerateKeyPair(crypto.Ed25519, -1)
	if err != nil {
		panic(err)
	}

	mPriv, err := crypto.MarshalPrivateKey(priv)
	if err != nil {
		panic(err)
	}

	mPub, err := crypto.MarshalPublicKey(pub)
	if err != nil {
		panic(err)
	}

	id, err := peer.IDFromPublicKey(pub)
	if err != nil {
		panic(err)
	}

	keysPeerID := KeysPeerID{
		PubKey:  hex.EncodeToString(mPub),
		PrivKey: hex.EncodeToString(mPriv),
		PeerID:  id.String(),
	}

	return keysPeerID
}
