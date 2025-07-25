package pesquisar_por

type PesquisarPor string

var (
	Nome    PesquisarPor = "NOME"    // Nome/Nome Fantasia
	CpfCnpj PesquisarPor = "CPFCNPJ" // CPF/CNPJ
	Ender   PesquisarPor = "ENDER"   // Endereço
)
