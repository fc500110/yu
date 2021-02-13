package txn

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"time"
	. "yu/common"
)

type UnsignedTxn struct {
	id        Hash
	caller    Address
	ecall     *Ecall
	timestamp int64
}

func NewUnsignedTxn(caller Address, ecall *Ecall) (*UnsignedTxn, error) {
	UnsignedTxn := &UnsignedTxn{
		caller:    caller,
		ecall:     ecall,
		timestamp: time.Now().UnixNano(),
	}
	id, err := UnsignedTxn.Hash()
	if err != nil {
		return nil, err
	}
	UnsignedTxn.id = id
	return UnsignedTxn, nil
}

func (t *UnsignedTxn) Caller() Address {
	return t.caller
}

func (t *UnsignedTxn) Ecall() *Ecall {
	return t.ecall
}

func (t *UnsignedTxn) Timestamp() int64 {
	return t.timestamp
}

func (t *UnsignedTxn) Hash() (Hash, error) {
	var hash Hash
	byt, err := t.Encode()
	if err != nil {
		return NullHash, err
	}
	hash = sha256.Sum256(byt)
	return hash, nil
}

func (t *UnsignedTxn) Encode() ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)
	err := encoder.Encode(t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte) (*UnsignedTxn, error) {
	var UnsignedTxn UnsignedTxn
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&UnsignedTxn)
	if err != nil {
		return nil, err
	}
	return &UnsignedTxn, nil
}