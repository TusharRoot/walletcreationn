package walletcreationn

import (
	"encoding/hex"
	"fmt"
)

func d() {
	mnemonic := "yard tissue turkey lucky unable cigar estate reflect manual food long uniform alley stumble error plate today dinner day render cabin artefact joy laptop"

	keygen0, err := generate()
	if err != nil {
		fmt.Println(err)
	}
	privatekey := hex.EncodeToString(keygen0.childkey.Key)
	fmt.Println("PrivateKey = ", privatekey, "", len(privatekey))
	fmt.Println("Public Address is ", keygen0.pubaddress)

	fmt.Println("\n", "Using mnemonic and Index from where you want to generate your Index")
	// for i := 1; i <= 10; i++ {
	keygen, err := generateWithIndex(mnemonic, uint32(1))
	if err != nil {
		fmt.Println(err)
	}
	privatekey1 := hex.EncodeToString(keygen.childkey.Key)
	fmt.Println("PrivateKey = ", privatekey1, "", len(privatekey1))
	fmt.Println("Public Address is ", keygen.pubaddress)

	fmt.Println("Generating Key from Parent Key")
	keygen2, err := generatefromkey(keygen.childkey, 1)
	if err != nil {
		fmt.Println(err)
	}
	privatekey2 := hex.EncodeToString(keygen.childkey.Key)
	fmt.Println("PrivateKey = ", privatekey2, "", len(privatekey2))
	fmt.Println("Public Address is ", keygen2.pubaddress)
	// }
}
