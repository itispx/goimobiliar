package conselho

import (
	"github.com/itispx/goimobiliar/condominio/conselho/cargo"
	"github.com/itispx/goimobiliar/condominio/conselho/sindico_profissional"
)

type Conselho struct {
	CodPessoa           int                                      `json:"CodPessoa,omitempty"`           // Código da pessoa.
	Cargo               cargo.Cargo                              `json:"Cargo,omitempty"`               // Cargo no conselho de condomínio.
	InicioMandato       string                                   `json:"InicioMandato,omitempty"`       // Data do início do mandato.
	FinalMandato        string                                   `json:"FinalMandato,omitempty"`        // Data do final de mandato.
	SindicoProfissional sindico_profissional.SindicoProfissional `json:"SindicoProfissional,omitempty"` // Indicação de síndico profissional.
	CodFornecedor       int                                      `json:"CodFornecedor,omitempty"`       // Código de fornecedor (se for o caso).
}
