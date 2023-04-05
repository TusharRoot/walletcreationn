package walletgo

import (
	"crypto/sha256"
	"fmt"
	"strconv"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
)

type key struct {
	childkey   *bip32.Key
	pubaddress string
}
type customerrors struct {
	Message string
	Code    int
}

func generate() (key, error) {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "")
	fmt.Println("Your Mnemonic:->", mnemonic)
	masterkey, err := bip32.NewMasterKey(seed)
	Error(err)
	childkey, err := masterkey.NewChildKey(0)
	Error(err)
	childpub := childkey.PublicKey()
	_, pubaddress := pubkeyhash(childpub.Key)
	//Return ChildPublic Address and Child Public Key
	return key{childkey, pubaddress}, nil
}

func (e customerrors) Error() string {
	return e.Message + " Error Code: " + strconv.Itoa(e.Code)
}
func generateWithIndex(mnemonic string, index uint32) (key, error) {
	if index > 10 {
		return key{}, customerrors{"Index Must be less than 10", 10}
	} else {
		fmt.Println("Your Mnemonic:->", mnemonic)
		seed := bip39.NewSeed(mnemonic, "")
		masterkey, err := bip32.NewMasterKey(seed)
		Error(err)
		return generatefromkey(masterkey, index)
	}
}
func generatefromkey(masterkey *bip32.Key, index uint32) (key, error) {
	if index > 10 {
		return key{}, customerrors{"Index Must be less than 10", 10}
	} else {
		childkey, err := masterkey.NewChildKey(index)
		Error(err)
		childpub := childkey.PublicKey()
		_, pubaddress := pubkeyhash(childpub.Key)
		//Return ChildPublic Address and Child Public Key
		return key{childkey, pubaddress}, nil
	}
}
func pubkeyhash(key []byte) (string, string) {
	versionByte := byte(0x00)
	shahash := sha256.Sum256(key)
	hasher := ripemd160.New()
	hasher.Write(shahash[:])
	hashBytes := hasher.Sum(nil)
	hashString := fmt.Sprintf("%x", hashBytes)
	versionedHash := append([]byte{versionByte}, hashBytes...)
	return hashString, pubkeyaddress(versionedHash)
}
func pubkeyaddress(versionedHash []byte) string {
	checksum := Checksum(versionedHash)
	fullHash := append(checksum, versionedHash...)
	pubaddress := "TS" + bip32.BitcoinBase58Encoding.EncodeToString(fullHash)
	return pubaddress
}
func Checksum(masterkey *bip32.Key, index uint32) (key, error) {
	// firstHash := sha256.Sum256(payload)
	// secondHash := sha256.Sum256(firstHash[:])
	// return secondHash[:4]

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	seed := bip39.NewSeed(mnemonic, "")
	fmt.Println("Your Mnemonic:->", mnemonic)
	masterkey, err := bip32.NewMasterKey(seed)
	Error(err)
	childkey, err := masterkey.NewChildKey(0)
	Error(err)
	childpub := childkey.PublicKey()
	_, pubaddress := pubkeyhash(childpub.Key)
	//Return ChildPublic Address and Child Public Key
	return key{childkey, pubaddress}, nil
}
func Error(e error) {
	if e != nil {
		panic(e)
	}
}
