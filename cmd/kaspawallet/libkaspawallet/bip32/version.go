package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// RustweaveMainnetPrivate is the version that is used for
// kaspa mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var RustweaveMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// RustweaveMainnetPublic is the version that is used for
// kaspa mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var RustweaveMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// RustweaveTestnetPrivate is the version that is used for
// kaspa testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var RustweaveTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// RustweaveTestnetPublic is the version that is used for
// kaspa testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var RustweaveTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// RustweaveDevnetPrivate is the version that is used for
// kaspa devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var RustweaveDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// RustweaveDevnetPublic is the version that is used for
// kaspa devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var RustweaveDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// RustweaveSimnetPrivate is the version that is used for
// kaspa simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var RustweaveSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// RustweaveSimnetPublic is the version that is used for
// kaspa simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var RustweaveSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case RustweaveMainnetPrivate:
		return RustweaveMainnetPublic, nil
	case RustweaveTestnetPrivate:
		return RustweaveTestnetPublic, nil
	case RustweaveDevnetPrivate:
		return RustweaveDevnetPublic, nil
	case RustweaveSimnetPrivate:
		return RustweaveSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case RustweaveMainnetPrivate:
		return true
	case RustweaveTestnetPrivate:
		return true
	case RustweaveDevnetPrivate:
		return true
	case RustweaveSimnetPrivate:
		return true
	}

	return false
}
