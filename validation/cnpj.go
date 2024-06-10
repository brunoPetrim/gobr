package validation

// IsCNPJValid verifies if the given string is a valid CNPJ document.
func IsCNPJValid(cnpj string) bool {
	const (
		size = 12
		pos  = 5
	)

	return isCPFOrCNPJ(cnpj, CNPJRegexp, size, pos)
}