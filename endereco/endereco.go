package endereco

import (
	"github.com/itispx/goimobiliar/endereco/logradouro"
	"github.com/itispx/goimobiliar/endereco/tipo_endereco"
	"github.com/itispx/goimobiliar/endereco/uf"
)

type Endereco struct {
	TipoEnder   tipo_endereco.TipoEndereco `json:"TipoEnder"`   // Tipo de endereço.
	CEP         int                        `json:"CEP"`         // Número do CEP.
	TipoLograd  logradouro.TipoLograd      `json:"TipoLograd"`  // Tipo de logradouro abreviado ou por extenso ('R' ou 'RUA', 'AV' ou 'AVENIDA', etc.).
	Logradouro  string                     `json:"Logradouro"`  // Logradouro do endereço. Deve ser informado apenas o nome sem o tipo de logradouro.
	Numero      int                        `json:"Numero"`      // Número do endereço.
	Complemento string                     `json:"Complemento"` // Complemento do endereço.
	Bairro      string                     `json:"Bairro"`      // Bairro do endereço.
	Cidade      string                     `json:"Cidade"`      // Cidade do endereço.
	UF          uf.UF                      `json:"UF"`          // Sigla da Unidade Federativa do endereço.
	Telefone1   string                     `json:"Telefone1"`   // Número de telefone principal.
	Telefone2   string                     `json:"Telefone2"`   // Número de telefone alternativo.
}
