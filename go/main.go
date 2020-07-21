package main

import "fmt"

import "github.com/nfjBill/yk5/utils"

func main() {
	cipherG := utils.Cipher.General
	gKeyPair := cipherG.Asymmetric.KeyPair()
	fmt.Printf("publicKey：%v\n", gKeyPair.Pub)
	fmt.Printf("privateKey：%v\n", gKeyPair.Pri)

	gRsaEnc := cipherG.Asymmetric.Encrypt("aaa", gKeyPair.Pub)
	fmt.Printf("GeneralRsaEnc：%v\n", gRsaEnc)

	gRsaDec := cipherG.Asymmetric.Decrypt(gRsaEnc, gKeyPair.Pri)
	fmt.Printf("GeneralRsaEDec：%v\n", gRsaDec)

	vc := "中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文中文"
	gKey := cipherG.Symmetry.GetKey(vc)
	fmt.Printf("GeneralKey：%v\n", gKey)

	gEcbEnc := cipherG.Symmetry.ECB.Encrypt(vc, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Printf("GeneralEcbEnc：%v\n", gEcbEnc)

	gEcbDec := cipherG.Symmetry.ECB.Decrypt(gEcbEnc, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	fmt.Printf("GeneralEcbEDec：%v\n", gEcbDec)

	gCbcEnc := cipherG.Symmetry.CBC.Encrypt(vc, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaa")
	fmt.Printf("GeneralCbcEnc：%v\n", gCbcEnc)

	gCbcDec := cipherG.Symmetry.CBC.Decrypt(gCbcEnc, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaa")
	fmt.Printf("GeneralCbcEDec：%v\n", gCbcDec)
}
