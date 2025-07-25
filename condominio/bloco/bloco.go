package bloco

import (
	"github.com/itispx/goimobiliar/condominio/bloco/fundo"
	"github.com/itispx/goimobiliar/condominio/conselho"
	"github.com/itispx/goimobiliar/endereco/logradouro"
)

type Bloco struct {
	CodBloco      string                `json:"CodBloco,omitempty"`      // Código do bloco do condomínio.
	TipoLograd    logradouro.TipoLograd `json:"TipoLograd,omitempty"`    // Tipo de logradouro do endereço.
	Descricao     string                `json:"Descricao,omitempty"`     // Descrição do bloco/conta.
	Fundo         fundo.Fundo           `json:"Fundo,omitempty"`         // Indica o tipo de fundo/conta.
	CEP           int                   `json:"CEP,omitempty"`           // CEP do condomínio.
	Endereco      string                `json:"Endereco,omitempty"`      // Endereço do condomínio.
	Bairro        string                `json:"Bairro,omitempty"`        // Bairro do endereço.
	QtdeEconomias int                   `json:"QtdeEconomias,omitempty"` // Total de economias do bloco.
	OrdemBloco    int                   `json:"OrdemBloco,omitempty"`    // Ordem de apresentação do bloco/conta.
	BlocoAtivo    string                `json:"BlocoAtivo,omitempty"`    // Informa se o Bloco/Conta está ativo.
	Conselho      []*conselho.Conselho  `json:"Conselho,omitempty"`      // Array no qual cada elemento é um objeto ("Membro" no XML) que possui os campos do objeto "conselho"
}
