package forma_pagamento

type FormaPagamento string

var (
	B FormaPagamento = "B" // CHEQUE (BANCO)
	C FormaPagamento = "C" // CHEQUE (CAIXA)
	R FormaPagamento = "R" // CRÉDITO CONTA
	A FormaPagamento = "A" // DÉBITO AUTOMÁTICO
	V FormaPagamento = "V" // DINHEIRO
	G FormaPagamento = "G" // LIQ. TIT. AGRUPAMENTO
	L FormaPagamento = "L" // LIQ. TÍTULOS
	N FormaPagamento = "N" // NÃO PAGAR
	O FormaPagamento = "O" // ORDEM PAGAMENTO
	Q FormaPagamento = "Q" // PIX (QRCODE)
	T FormaPagamento = "T" // PIX (TRANSFERÊNCIA)
)
