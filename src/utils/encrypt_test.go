package utils

import (
	"encoding/hex"
	"log"
	"os"
	"testing"

	"github.com/Bean-jun/Dlog/pkg"
)

func TestGenRsaKey(t *testing.T) {
	gotPrvkey, gotPubkey := GenRsaKey()
	// WriteFile("../cert/private_key.pem", gotPrvkey)
	// WriteFile("../cert/public_key.pem", gotPubkey)
	WriteFile("../cert/private_key_new.pem", gotPrvkey)
	WriteFile("../cert/public_key_new.pem", gotPubkey)
}

func TestRsaEncryptDecrypt(t *testing.T) {
	err := os.Chdir("../")
	if err != nil {
		panic(err)
	}

	content := "dlog"
	pkg.InitConfig("conf.yaml")

	encryTest := struct {
		data     []byte
		keyBytes []byte
	}{
		data:     []byte(content),
		keyBytes: pkg.Conf.Server.Cert.PublicKey,
	}

	encryData, err := RsaEncrypt(encryTest.data, encryTest.keyBytes)
	if err != nil {
		panic(err)
	}
	log.Println("加密后密文为:", hex.EncodeToString(encryData))

	decryTest := struct {
		data     []byte
		keyBytes []byte
	}{
		data:     encryData,
		keyBytes: pkg.Conf.Server.Cert.PrivateKey,
	}
	decryData, err := RsaDecrypt(decryTest.data, decryTest.keyBytes)
	if err != nil {
		panic(err)
	}
	log.Println("解密后明文为:", string(decryData))

	if content != string(decryData) {
		t.Errorf("RsaDecrypt() = %v, want %v", decryData, content)
	}
}

func TestRsaEncryptDecryptFactory(t *testing.T) {
	err := os.Chdir("../")
	if err != nil {
		panic(err)
	}

	content := "dlog"
	pkg.InitConfig("conf.yaml")

	encryTest := struct {
		data     []byte
		keyBytes []byte
	}{
		data:     []byte(content),
		keyBytes: pkg.Conf.Server.Cert.PublicKey,
	}

	encryData := RsaEncryptFactory(encryTest.data, encryTest.keyBytes)
	log.Println("加密后密文为:", encryData)

	decryTest := struct {
		ciphertext string
		keyBytes   []byte
	}{
		ciphertext: encryData,
		keyBytes:   pkg.Conf.Server.Cert.PrivateKey,
	}
	decryData := RsaDecryptFactory(decryTest.ciphertext, decryTest.keyBytes)
	log.Println("解密后明文为:", decryData)

	if content != decryData {
		t.Errorf("RsaDecrypt() = %v, want %v", decryData, content)
	}
}
