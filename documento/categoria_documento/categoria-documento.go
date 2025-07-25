package categoria_documento

type CategoriaDocumento string

var (
	CO            CategoriaDocumento = "CO" // Comprovante
	DT            CategoriaDocumento = "DT" // Documentos terceirizadas
	FB            CategoriaDocumento = "FB" // Fatura/Boleto
	IP            CategoriaDocumento = "IP" // Imobiliar Pay
	NF            CategoriaDocumento = "NF" // Nota Fiscal
	RE            CategoriaDocumento = "RE" // Recibo
	RP            CategoriaDocumento = "RP" // RPA
	SEM_CATEGORIA CategoriaDocumento = ""   // Sem Categoria
)
