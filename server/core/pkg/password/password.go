package password

import "golang.org/x/crypto/bcrypt"

// Encrypt 密码加密
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare 密码比对 (传入未加密的密码即可)
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Encrypted 判断密码是否加密过了
func Encrypted(pwd string) (status bool) {
  return len(pwd) == 60 // 长度等于 60 说明加密过了
}
