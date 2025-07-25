package tipo_documento

type TipoDocumento string

var (
	B TipoDocumento = "B" // BLOQUETO BANCÁRIO
	C TipoDocumento = "C" // BLOQUETO CONCESSIONÁRIA
	H TipoDocumento = "H" // CONTRA CHEQUE
	U TipoDocumento = "U" // CUPOM FISCAL
	F TipoDocumento = "F" // DARF
	A TipoDocumento = "A" // DARF-CÓD.BARRA
	D TipoDocumento = "D" // DUPLICATA
	W TipoDocumento = "W" // FGTS-GRF/GRRF
	I TipoDocumento = "I" // INSS AUTÔNOMO
	T TipoDocumento = "T" // NOTA FISCAL PRODUTO
	N TipoDocumento = "N" // NOTA FISCAL SERVIÇO
	G TipoDocumento = "G" // OUTRA GUIA-CÓD.BARRA
	O TipoDocumento = "O" // OUTROS
	P TipoDocumento = "P" // PIS AUTÔNOMO
	R TipoDocumento = "R" // RECIBO SIMPLES(RPA)
	S TipoDocumento = "S" // SLIP LANÇAMENTO
)
