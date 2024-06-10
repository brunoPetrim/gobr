package validation

// IsCPFValid verifies if the given string is a valid CPF document.
func IsCPFValid(cpf string) bool {
	const (
		size = 9
		pos  = 10
	)

	return isCPFOrCNPJ(cpf, CPFRegexp, size, pos)
}
