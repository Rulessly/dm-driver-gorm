package security

type Cipher interface {
	Encrypt(plaintext []byte, genDigest bool) []byte
	Decrypt(ciphertext []byte, checkDigest bool) ([]byte, error)
}
