// Code generated by "stringer -type=SignatureAlgorithm"; DO NOT EDIT.

package sema

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[SignatureAlgorithmECDSA_P256-0]
	_ = x[SignatureAlgorithmECDSA_Secp256k1-1]
}

const _SignatureAlgorithm_name = "SignatureAlgorithmECDSA_P256SignatureAlgorithmECDSA_Secp256k1"

var _SignatureAlgorithm_index = [...]uint8{0, 28, 61}

func (i SignatureAlgorithm) String() string {
	if i < 0 || i >= SignatureAlgorithm(len(_SignatureAlgorithm_index)-1) {
		return "SignatureAlgorithm(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SignatureAlgorithm_name[_SignatureAlgorithm_index[i]:_SignatureAlgorithm_index[i+1]]
}
