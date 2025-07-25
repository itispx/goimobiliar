package arquivo

type Arquivo struct {
	ArquivoNome     string `json:"ArquivoNome,omitempty"`     // Nome do arquivo.
	ArquivoTamanho  string `json:"ArquivoTamanho,omitempty"`  // Tamanho do arquivo em bytes.
	ArquivoDataHora string `json:"ArquivoDataHora,omitempty"` // Data e hora da última modificação do arquivo no formato: AAAA-MM-DD-hh-mm-ss.
	URL             string `json:"URL,omitempty"`             // URL para download do arquivo.
}
