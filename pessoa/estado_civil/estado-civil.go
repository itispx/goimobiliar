package estado_civil

type EstadoCivil string

var (
	S EstadoCivil = "S" // Solteiro(a)
	C EstadoCivil = "C" // Casado(a)
	D EstadoCivil = "D" // Divorciado(a)
	P EstadoCivil = "P" // seParado(a)
	V EstadoCivil = "V" // Viúvo(a)
	N EstadoCivil = "N" // uNião estável
)
