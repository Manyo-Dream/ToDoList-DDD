package encrypt

import (
	"github.com/manyodream/todolist-ddd/domain/user/repository"
	"golang.org/x/crypto/bcrypt"
)

type PwdEncryptSeevice struct {
}

func NewPwdEncryptService() repository.PwdEncrypt {
	return &PwdEncryptSeevice{}
}

const (
	PasswordCost = 12
)

func (p *PwdEncryptSeevice) Encrypt(data []byte) ([]byte, error) {
	bytes, err := bcrypt.GenerateFromPassword(data, PasswordCost)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (p *PwdEncryptSeevice) Decrypt(data []byte) ([]byte, error) {
	return []byte{}, nil
}

func (p *PwdEncryptSeevice) Check(data, src []byte) bool {
	return bcrypt.CompareHashAndPassword(data, src) == nil
}
