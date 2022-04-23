package poa

import (
	. "github.com/yu-org/yu/core/keypair"
)

type pair struct {
	pubkey  PubKey
	privkey PrivKey
}

func InitDefaultKeypairs(idx int) (PubKey, PrivKey, []ValidatorInfo) {
	pub0, priv0 := GenSrKey([]byte("node1"))
	println("pubkey0: ", pub0.String())

	pub1, priv1 := GenSrKey([]byte("node2"))
	println("pubkey1: ", pub1.String())

	pub2, priv2 := GenSrKey([]byte("node3"))
	println("pubkey2: ", pub2.String())

	pairArray := []pair{
		{
			pubkey:  pub0,
			privkey: priv0,
		},
		{
			pubkey:  pub1,
			privkey: priv1,
		},
		{
			pubkey:  pub2,
			privkey: priv2,
		},
	}

	myPubkey := pairArray[idx].pubkey
	myPrivkey := pairArray[idx].privkey
	validatorsAddrs := []ValidatorInfo{
		{Pubkey: pub0, P2pIP: "12D3KooWHHzSeKaY8xuZVzkLbKFfvNgPPeKhFBGrMbNzbm5akpqu"},
		{Pubkey: pub1, P2pIP: "12D3KooWSKPs95miv8wzj3fa5HkJ1tH7oEGumsEiD92n2MYwRtQG"},
		{Pubkey: pub2, P2pIP: "12D3KooWRuwP7nXaRhZrmoFJvPPGat2xPafVmGpQpZs5zKMtwqPH"},
	}

	return myPubkey, myPrivkey, validatorsAddrs
}
