package pessoa

import (
	"github.com/itispx/goimobiliar/endereco"
	"github.com/itispx/goimobiliar/endereco/tipo_endereco"
	"github.com/itispx/goimobiliar/pessoa/estado_civil"
	"github.com/itispx/goimobiliar/pessoa/sexo"
	"github.com/itispx/goimobiliar/pessoa/tipo_pessoa"
	"github.com/itispx/goimobiliar/tabela"
)

type Classificacao string

var (
	G Classificacao = "G" // GESTAO
	P Classificacao = "P" // PADRÃO
	A Classificacao = "A" // PRIME
	S Classificacao = "S" // SUPER VIP
	U Classificacao = "U" // TAXA UNICA
	V Classificacao = "V" // VIP
)

type Pessoa struct {
	CodPessoa         int                        `json:"CodPessoa,omitempty"`         // Código da pessoa.
	Nome              string                     `json:"Nome,omitempty"`              // Nome da pessoa.
	EstadoCivil       estado_civil.EstadoCivil   `json:"EstadoCivil,omitempty"`       // Estado civil da pessoa.
	Sexo              sexo.Sexo                  `json:"Sexo,omitempty"`              // Sexo/gênero da pessoa.
	TipoPessoa        tipo_pessoa.TipoPessoa     `json:"TipoPessoa,omitempty"`        // Tipo da pessoa.
	CpfCnpj           int                        `json:"CpfCnpj,omitempty"`           // Se for tipo de pessoa física o valor é um CPF. Se for tipo de pessoa jurídica o valor é um CNPJ. Se o tipo de pessoa não for informado então este campo é vazio.
	RG                string                     `json:"RG,omitempty"`                // Número do documento de identificação da pessoa física. Não preencher se for pessoa jurídica.
	OrgaoExpedidor    string                     `json:"OrgaoExpedidor,omitempty"`    // Órgão que expediu o documento de identificação informado.
	DataNascimento    string                     `json:"DataNascimento,omitempty"`    // Data de nascimento da pessoa física ou de criação da pessoa jurídica.
	Nacionalidade     string                     `json:"Nacionalidade,omitempty"`     // Nacionalidade da pessoa no padrão do e-Social.
	CodNacionalidade  int                        `json:"CodNacionalidade,omitempty"`  // Código de nacionalidade da pessoa no e-Social.
	Naturalidade      string                     `json:"Naturalidade,omitempty"`      // Naturalidade da pessoa no padrão do DIMOB.
	CodNaturalidade   int                        `json:"CodNaturalidade,omitempty"`   // Naturalidade da pessoa no DIMOB.
	Celular           string                     `json:"Celular,omitempty"`           // Número de celular.
	Email             string                     `json:"Email,omitempty"`             // E-mail da pessoa.
	Contato           string                     `json:"Contato,omitempty"`           // Informações de pessoa de contato.
	Ativo             tabela.SimNao              `json:"Ativo,omitempty"`             // Indica se está ativo.
	TipoEnderCobr     tipo_endereco.TipoEndereco `json:"TipoEnderCobr,omitempty"`     // Tipo de endereço de cobrança que deve existir no array 'Enderecos'.
	TipoEnderCorresp  tipo_endereco.TipoEndereco `json:"TipoEnderCorresp,omitempty"`  // Tipo de endereço de correpondência que deve existir no array 'Enderecos'.
	DataInclusao      string                     `json:"DataInclusao,omitempty"`      // Data de inclusão no sistema.
	NomePai           string                     `json:"NomePai,omitempty"`           // Nome do pai da pessoa física.
	NomeMae           string                     `json:"NomeMae,omitempty"`           // Nome da mãe da pessoa física.
	CodConjuge        int                        `json:"CodConjuge,omitempty"`        // Código de pessoa do cônjuge.
	PIS               string                     `json:"PIS,omitempty"`               // PIS da pessoa da pessoa física.
	CodBanco          int                        `json:"CodBanco,omitempty"`          // Código do banco.
	CodAgencia        int                        `json:"CodAgencia,omitempty"`        // Código da agência bancária.
	ContaCorrente     string                     `json:"ContaCorrente,omitempty"`     // Número da conta corrente desta pessoa.
	TipoConta         string                     `json:"TipoConta,omitempty"`         // Tipo da conta bancária desta pessoa.
	Passaporte        string                     `json:"Passaporte,omitempty"`        // Número do passaporte da pessoa física.
	SenhaInternetMD5  string                     `json:"SenhaInternetMD5,omitempty"`  // Valor MD5 da senha de acesso ao site/internet.
	CodIntegracaoSist string                     `json:"CodIntegracaoSist,omitempty"` // Código de integração/migração de sistema.
	CodProfissao      int                        `json:"CodProfissao,omitempty"`      // Código da profissão desta pessoa.
	Classificacao     Classificacao              `json:"Classificacao,omitempty"`     // Código de classificacão desta pessoa.
	DataAlteracao     string                     `json:"DataAlteracao,omitempty"`     // Data da última alteração no sistema.
	Enderecos         []*endereco.Endereco       `json:"Enderecos,omitempty"`         // Array no qual cada elemento é um endereço.
	Locatario         tabela.SimNao              `json:"Locatario,omitempty"`         // Indica se é locatário.
	Proprietario      tabela.SimNao              `json:"Proprietario,omitempty"`      // Indica se é proprietário.
	Fiador            tabela.SimNao              `json:"Fiador,omitempty"`            // Indica se é fiador.
	Sindico           tabela.SimNao              `json:"Sindico,omitempty"`           // Indica se é síndico.
	Condomino         tabela.SimNao              `json:"Condomino,omitempty"`         // Indica se é condômino.
	Beneficiario      tabela.SimNao              `json:"Beneficiario,omitempty"`      // Indica se é beneficiário.
	Procurador        tabela.SimNao              `json:"Procurador,omitempty"`        // Indica se é procurador.
	Assessor          string                     `json:"Assessor,omitempty"`          // Código de usuário do assessor responsável.
	AssessorTelefone  string                     `json:"AssessorTelefone,omitempty"`  // Telefone do assessor responsável.
	AssessorEmail     string                     `json:"AssessorEmail,omitempty"`     // Email do assessor responsável.
	EmailAutomatico   tabela.SimNao              `json:"EmailAutomatico,omitempty"`   // Avisos automáticos por e-mail.
	EmailNfse         string                     `json:"EmailNfse,omitempty"`         // Utilizado na emissão na NFSe.
	WhatsPrioritario  tabela.SimNao              `json:"WhatsPrioritario,omitempty"`  // Campanhas ativas por WhatsApp.
}
