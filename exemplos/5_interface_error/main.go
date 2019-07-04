// criando um tipo de erro usando interface de error
// type error interface {
// 		Error() string
// }

package main

import (
	"fmt"
)

type InvalidArgumentError struct {
	Message string
	InnerError error
}

func New(message string, innerError error) *InvalidArgumentError {
	return &InvalidArgumentError{Message:message, InnerError:innerError}
}

func (e *InvalidArgumentError) Error() string {
	erro := fmt.Sprint(" ** Exception -> ", e.Message)

	if e.InnerError != nil {
		erro = fmt.Sprint(erro, "\n\tInner: ", e.InnerError)
	}

	return erro
}

func main() {
	total, err := somaPositivo(-1, 2)

	if err != nil {
		e := NewInvalidArgumentWithInnerError("Erro 1 na soma", err)
		e2 := NewInvalidArgumentWithInnerError("Erro 2 na soma", e)

		fmt.Println(e2)

		return
	}

	fmt.Println("total = ", total)

}

func somaPositivo(v1, v2 int) (int, error) {

	if v1 < 0 || v2 < 0 {
		return 0, NewInvalidArgumentError(fmt.Sprintf("valor 1 (%d) ou Valor 2 (%d) devem ser maior que zero", v1, v2))
	}

	return v1 + v2, nil

}

func NewInvalidArgumentError(message string) error {
	return New(message, nil)
}

func NewInvalidArgumentWithInnerError(message string, inner error) error {
	return New(message, inner)
}
