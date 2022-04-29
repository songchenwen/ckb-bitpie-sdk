package client

import (
	"testing"

	"github.com/shaojunda/ckb-bitpie-sdk/config"
)

func TestIsOldAcpAddress(t *testing.T) {
	conf, err := config.Load("../config-example.yaml")
	if err != nil {
		t.Error(err)
	}
	cases := []struct {
		Name     string
		Addr     string
		Conf     *config.Config
		Expected bool
	}{
		{"old acp address", "ckb1qg8mxsu48mncexvxkzgaa7mz2g25uza4zpz062relhjmyuc52ps3rg0ey4hrkr6sws7wtjsuv8qmnu7kmmy7u9ut8lp", conf, true},
		{"new acp address", "ckb1qyp2r7f9dcas75r58nju58rpcxul84k7e8hqvma9q8", conf, false},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans, _ := IsOldAcpAddress(c.Addr, c.Conf); ans != c.Expected {
				t.Fatalf("should return %t, but got %t", c.Expected, ans)
			}
		})
	}
}

func TestPubkey2Address(t *testing.T) {
	conf, err := config.Load("../config-example.yaml")
	if err != nil {
		t.Error(err)
	}
	t.Run("generate new acp address", func(t *testing.T) {
		addr, err := Pubkey2Address("0x024de021442989d81798d1eeedc983ab70b815949115b94b1b2b00b32f872e4ace", true, false, conf)
		if err != nil {
			t.Error(err)
		}
		expectedAcpAddr := "ckb1qyp2r7f9dcas75r58nju58rpcxul84k7e8hqvma9q8"
		if addr != expectedAcpAddr {
			t.Fatalf("should return %s, but got %s", expectedAcpAddr, addr)
		}
		t.Log("generate new acp address success")
	})

	t.Run("generate none acp address", func(t *testing.T) {
		addr, err := Pubkey2Address("0x024de021442989d81798d1eeedc983ab70b815949115b94b1b2b00b32f872e4ace", false, false, conf)
		if err != nil {
			t.Error(err)
		}
		expectedOldAcpAddr := "ckb1qyq2r7f9dcas75r58nju58rpcxul84k7e8hqzt8jm9"
		if addr != expectedOldAcpAddr {
			t.Fatalf("should return %s, but got %s", expectedOldAcpAddr, addr)
		}
	})

	t.Run("generate old acp address", func(t *testing.T) {
		addr, err := Pubkey2Address("0x024de021442989d81798d1eeedc983ab70b815949115b94b1b2b00b32f872e4ace", true, true, conf)
		if err != nil {
			t.Error(err)
		}
		expectedOldAcpAddr := "ckb1qg8mxsu48mncexvxkzgaa7mz2g25uza4zpz062relhjmyuc52ps3rg0ey4hrkr6sws7wtjsuv8qmnu7kmmy7u9ut8lp"
		if addr != expectedOldAcpAddr {
			t.Fatalf("should return %s, but got %s", expectedOldAcpAddr, addr)
		}
	})
}
