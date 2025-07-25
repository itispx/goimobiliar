package erros

import (
	"errors"
	"fmt"
)

var (
	ErrBaseInvalida   = errors.New("base inválida")
	ErrSessaoInvalida = errors.New("sessão inválida")
)

func ErrCampoVazio(f string) error {
	return fmt.Errorf("campo '%s' vazio", f)
}
