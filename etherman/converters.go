package etherman

import (
	"fmt"
	"math/big"

	"github.com/0xPolygonHermez/zkevm-node/aggregator/pb"
	"github.com/0xPolygonHermez/zkevm-node/encoding"
)

const (
	minProofLen = 2
	maxProofLen = 3
)

func strSliceToBigIntArray(data []string) ([2]*big.Int, error) {
	if len(data) < minProofLen || len(data) > maxProofLen {
		return [2]*big.Int{}, fmt.Errorf("wrong slice length, current %d, expected between %d or %d", len(data), minProofLen, maxProofLen)
	}
	var res [2]*big.Int
	for i, v := range data {
		if i < minProofLen {
			bigInt, ok := new(big.Int).SetString(v, encoding.Base10)
			if !ok {
				return [2]*big.Int{}, fmt.Errorf("failed to convert string to big int, str: %s", v)
			}
			res[i] = bigInt
		}
	}
	return res, nil
}

func proofSlcToIntArray(proofs []*pb.ProofB) ([2][2]*big.Int, error) {
	if len(proofs) < minProofLen || len(proofs) > maxProofLen {
		return [2][2]*big.Int{}, fmt.Errorf("wrong slice length, current %d, expected between %d or %d", len(proofs), minProofLen, maxProofLen)
	}

	var res [2][2]*big.Int
	for i, v := range proofs {
		if i < minProofLen {
			for j, b := range proofs[i].Proofs {
				if j < minProofLen {
					bigInt, ok := new(big.Int).SetString(b, encoding.Base10)
					if !ok {
						return [2][2]*big.Int{}, fmt.Errorf("failed to convert string to big int, str: %s", v)
					}
					res[i][1-j] = bigInt
				}
			}
		}
	}

	return res, nil
}
