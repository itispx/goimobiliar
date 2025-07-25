package fornecedor

import (
	"github.com/itispx/goimobiliar/documento/tipo_documento"
	"github.com/itispx/goimobiliar/endereco/logradouro"
	"github.com/itispx/goimobiliar/endereco/uf"
	"github.com/itispx/goimobiliar/fornecedor/categoria"
	"github.com/itispx/goimobiliar/pagamento/forma_pagamento"
	"github.com/itispx/goimobiliar/pessoa/tipo_pessoa"
	"github.com/itispx/goimobiliar/tabela"
)

type Fornecedor struct {
	CodFornecedor       int                            `json:"CodFornecedor,omitempty"`       // Código do fornecedor.
	Nome                string                         `json:"Nome,omitempty"`                // Nome/Razão Social do fornecedor.
	NomeFantasia        string                         `json:"NomeFantasia,omitempty"`        // Nome de fantasia do fornecedor.
	TipoPessoa          tipo_pessoa.TipoPessoa         `json:"TipoPessoa,omitempty"`          // Tipo de pessoa do fornecedor.
	CpfCnpj             int                            `json:"CpfCnpj,omitempty"`             // CPF ou CNPJ do fornecedor.
	InscricaoInss       string                         `json:"InscricaoInss,omitempty"`       // CPF/CNPJ do fornecedor.
	InscricaoMunicipal  string                         `json:"InscricaoMunicipal,omitempty"`  // Inscrição municipal do fornecedor.
	Categoria           categoria.Categoria            `json:"Categoria,omitempty"`           // Categoria do fornecedor.
	PIS                 string                         `json:"PIS,omitempty"`                 // PIS do fornecedor.
	TipoConta           string                         `json:"TipoConta,omitempty"`           // Tipo da conta bancária do fornecedor.
	CodBanco            int                            `json:"CodBanco,omitempty"`            // Código do banco.
	CodAgencia          int                            `json:"CodAgencia,omitempty"`          // Código da agência bancária.
	ContaCorrente       string                         `json:"ContaCorrente,omitempty"`       // Número da conta corrente do fornecedor.
	Contato             string                         `json:"Contato,omitempty"`             // Contato no fornecedor.
	CargoContato        string                         `json:"CargoContato,omitempty"`        // Cargo do contato no fornecedor.
	CEP                 int                            `json:"CEP,omitempty"`                 // Número do CEP.
	TipoLograd          logradouro.TipoLograd          `json:"TipoLograd,omitempty"`          // Tipo de logradouro abreviado ou por extenso ('R' ou 'RUA', 'AV' ou 'AVENIDA', etc.).
	Logradouro          string                         `json:"Logradouro,omitempty"`          // Logradouro do endereço. Deve ser informado apenas o nome sem o tipo de logradouro.
	Numero              int                            `json:"Numero,omitempty"`              // Número do endereço.
	Complemento         string                         `json:"Complemento,omitempty"`         // Complemento do endereço.
	Bairro              string                         `json:"Bairro,omitempty"`              // Bairro do endereço.
	Cidade              string                         `json:"Cidade,omitempty"`              // Cidade do endereço.
	UF                  uf.UF                          `json:"UF,omitempty"`                  // Sigla da Unidade Federativa do endereço.
	Telefone1           string                         `json:"Telefone1,omitempty"`           // Número do telefone principal.
	Celular             string                         `json:"Celular,omitempty"`             // Número do celular do fornecedor.
	Email               string                         `json:"Email,omitempty"`               // E-mail do fornecedor.
	FormaPagamento      forma_pagamento.FormaPagamento `json:"FormaPagamento,omitempty"`      // Forma de pagamento do fornecedor.
	TipoDocumento       tipo_documento.TipoDocumento   `json:"TipoDocumento,omitempty"`       // Tipos de documentos.
	EmiteNFSE           string                         `json:"EmiteNFSE,omitempty"`           // Indica se fornecedor emite NFSe.
	Ativo               tabela.SimNao                  `json:"Ativo,omitempty"`               // Indica se está ativo.
	CodPessoaFavorecido int                            `json:"CodPessoaFavorecido,omitempty"` // Código da pessoa favorecida em pagamentos ao fornecedor.
	Favorecido          string                         `json:"Favorecido,omitempty"`          // Nome da pessoa favorecida.
	CodPessoaTitular    int                            `json:"CodPessoaTitular,omitempty"`    // Código da pessoa titular da empresa para fins previdenciários.
	Titular             string                         `json:"Titular,omitempty"`             // Nome da pessoa titular da empresa para fins previdenciários.
	MEI                 string                         `json:"MEI,omitempty"`                 // MEI do fornecedor.
	NIT                 string                         `json:"NIT,omitempty"`                 // NIT do fornecedor.
	ProdutorRural       string                         `json:"ProdutorRural,omitempty"`       // Indica se o fornecedor é produtor rural.
	CodigoCBO           string                         `json:"CodigoCBO,omitempty"`           // Código CBO (Classificação Brasileira de Ocupações).
}
