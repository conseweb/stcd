// Copyright (c) 2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"time"

	"github.com/conseweb/stcd/wire"
)

// genesisCoinbaseTx is the coinbase transaction for the genesis blocks for
// the main network, regression test network, and test network (version 3).
var genesisCoinbaseTx = wire.MsgTx{
	Version: 1,
	TxIn: []*wire.TxIn{
		{
			PreviousOutPoint: wire.OutPoint{
				Hash:  wire.ShaHash{},
				Index: 0xffffffff,
			},
			SignatureScript: []byte{
				0x04, 0xff, 0xff, 0x00, 0x1d, 0x01, 0x04, 0x55, /* |.......U| */
				0x32, 0x30, 0x31, 0x36, 0x2d, 0x31, 0x2d, 0x31,
				0x39, 0x3a, 0x20, 0x43, 0x6f, 0x6e, 0x73, 0x65,
				0x77, 0x65, 0x62, 0x20, 0x77, 0x69, 0x6c, 0x6c,
				0x20, 0x62, 0x65, 0x63, 0x6f, 0x6d, 0x65, 0x20,
				0x74, 0x68, 0x65, 0x20, 0x63, 0x6f, 0x72, 0x65,
				0x20, 0x74, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c,
				0x6f, 0x67, 0x79, 0x20, 0x74, 0x6f, 0x20, 0x62,
				0x75, 0x69, 0x6c, 0x64, 0x20, 0x74, 0x68, 0x65,
				0x20, 0x6e, 0x65, 0x78, 0x74, 0x2d, 0x67, 0x65,
				0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
				0x20, 0x57, 0x45, 0x42, 0x2e,
			},
			Sequence: 0xffffffff,
		},
	},
	TxOut: []*wire.TxOut{
		{
			Value: 0x12a05f200, // 50 coins or 5,000,000,000 satoshi
			PkScript: []byte{
				0x41, 0x04, 0xdb, 0x11, 0x09, 0x94, 0xc5, 0xcc,
				0xa3, 0x86, 0x08, 0x2f, 0x4e, 0x4e, 0xd0, 0x1f,
				0xc0, 0xab, 0xb3, 0x76, 0x1a, 0x92, 0x93, 0xc9,
				0x9e, 0x94, 0x59, 0xb4, 0xdb, 0xc7, 0x8f, 0x8f,
				0xbe, 0x14, 0x6f, 0x2c, 0x93, 0x1f, 0x6f, 0x4b,
				0x4d, 0xa6, 0xdc, 0x3e, 0x60, 0x74, 0x02, 0x45,
				0xa6, 0x24, 0x80, 0x72, 0xa8, 0xae, 0x7f, 0xc5,
				0xd7, 0x97, 0x51, 0xaa, 0x0d, 0xd4, 0x26, 0x50,
				0x77, 0x0a, 0xac, /* |._.| */
			},
		},
	},
	LockTime: 0,
}

// genesisHash is the hash of the first block in the block chain for the main
// network (genesis block).
// gh = '00000000e76f41a15308a34d446769703458a71ad609e1efdf842d24acfc36e9'
// ga = map(hex, map(ord, gh[::-1].decode('hex')))
var genesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0xe9, 0x36, 0xfc, 0xac, 0x24, 0x2d, 0x84, 0xdf,
	0xef, 0xe1, 0x09, 0xd6, 0x1a, 0xa7, 0x58, 0x34,
	0x70, 0x69, 0x67, 0x44, 0x4d, 0xa3, 0x08, 0x53,
	0xa1, 0x41, 0x6f, 0xe7, 0x00, 0x00, 0x00, 0x00,
})

// genesisMerkleRoot is the hash of the first transaction in the genesis block
// for the main network.
// mh = 'd2effb859267e598a4c916297836cfa6f93f4c4db5ba30a43768b4396385ac28'
// xf = lambda x: map(hex, map(ord, x[::-1].decode('hex')))
// ma = xf(mh)
var genesisMerkleRoot = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x28, 0xac, 0x85, 0x63, 0x39, 0xb4, 0x68, 0x37,
	0xa4, 0x30, 0xba, 0xb5, 0x4d, 0x4c, 0x3f, 0xf9,
	0xa6, 0xcf, 0x36, 0x78, 0x29, 0x16, 0xc9, 0xa4,
	0x98, 0xe5, 0x67, 0x92, 0x85, 0xfb, 0xef, 0xd2,
})

// genesisBlock defines the genesis block of the block chain which serves as the
// public transaction ledger for the main network.
var genesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: genesisMerkleRoot,        // d2effb859267e598a4c916297836cfa6f93f4c4db5ba30a43768b4396385ac28
		Timestamp:  time.Unix(0x569e0ebb, 0), // 1453199035, 2016/1/19 下午6:23:55 GMT+8:00
		Bits:       0x1d00ffff,               // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0xed1dc177,               // 3978150263
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// regTestGenesisHash is the hash of the first block in the block chain for the
// regression test network (genesis block).
var regTestGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x06, 0x22, 0x6e, 0x46, 0x11, 0x1a, 0x0b, 0x59,
	0xca, 0xaf, 0x12, 0x60, 0x43, 0xeb, 0x5b, 0xbf,
	0x28, 0xc3, 0x4f, 0x3a, 0x5e, 0x33, 0x2a, 0x1f,
	0xc7, 0xb2, 0xb7, 0x3c, 0xf1, 0x88, 0x91, 0x0f,
})

// regTestGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the regression test network.  It is the same as the merkle root for
// the main network.
var regTestGenesisMerkleRoot = genesisMerkleRoot

// regTestGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the regression test network.
var regTestGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: regTestGenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0), // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// testNet3GenesisHash is the hash of the first block in the block chain for the
// test network (version 3).
var testNet3GenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0x43, 0x49, 0x7f, 0xd7, 0xf8, 0x26, 0x95, 0x71,
	0x08, 0xf4, 0xa3, 0x0f, 0xd9, 0xce, 0xc3, 0xae,
	0xba, 0x79, 0x97, 0x20, 0x84, 0xe9, 0x0e, 0xad,
	0x01, 0xea, 0x33, 0x09, 0x00, 0x00, 0x00, 0x00,
})

// testNet3GenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the test network (version 3).  It is the same as the merkle root
// for the main network.
var testNet3GenesisMerkleRoot = genesisMerkleRoot

// testNet3GenesisBlock defines the genesis block of the block chain which
// serves as the public transaction ledger for the test network (version 3).
var testNet3GenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},            // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: testNet3GenesisMerkleRoot, // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1296688602, 0),  // 2011-02-02 23:16:42 +0000 UTC
		Bits:       0x1d00ffff,                // 486604799 [00000000ffff0000000000000000000000000000000000000000000000000000]
		Nonce:      0x18aea41a,                // 414098458
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}

// simNetGenesisHash is the hash of the first block in the block chain for the
// simulation test network.
var simNetGenesisHash = wire.ShaHash([wire.HashSize]byte{ // Make go vet happy.
	0xf6, 0x7a, 0xd7, 0x69, 0x5d, 0x9b, 0x66, 0x2a,
	0x72, 0xff, 0x3d, 0x8e, 0xdb, 0xbb, 0x2d, 0xe0,
	0xbf, 0xa6, 0x7b, 0x13, 0x97, 0x4b, 0xb9, 0x91,
	0x0d, 0x11, 0x6d, 0x5c, 0xbd, 0x86, 0x3e, 0x68,
})

// simNetGenesisMerkleRoot is the hash of the first transaction in the genesis
// block for the simulation test network.  It is the same as the merkle root for
// the main network.
var simNetGenesisMerkleRoot = genesisMerkleRoot

// simNetGenesisBlock defines the genesis block of the block chain which serves
// as the public transaction ledger for the simulation test network.
var simNetGenesisBlock = wire.MsgBlock{
	Header: wire.BlockHeader{
		Version:    1,
		PrevBlock:  wire.ShaHash{},           // 0000000000000000000000000000000000000000000000000000000000000000
		MerkleRoot: simNetGenesisMerkleRoot,  // 4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b
		Timestamp:  time.Unix(1401292357, 0), // 2014-05-28 15:52:37 +0000 UTC
		Bits:       0x207fffff,               // 545259519 [7fffff0000000000000000000000000000000000000000000000000000000000]
		Nonce:      2,
	},
	Transactions: []*wire.MsgTx{&genesisCoinbaseTx},
}
