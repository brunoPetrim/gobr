# gobr

## pt-br
Módulo golang com validações diversas de documentos e especificidades do Brasil. Funcionalidades presentes:

- [x] Validação de CPF
- [x] Validação de CNPJ
- [ ] Formatação de CPF
- [ ] Formatação de CNPJ
- [ ] Desformatação de CPF
- [ ] Desformatação de CNPJ
- [ ] Geração de CPF válido para testes
- [ ] Geração de CNPJ válido para testes
- [ ] Validação de CNH
- [ ] Geração de CNH válida para testes
- [ ] Validação de RENAVAN
- [ ] Geração de RENAVAN válido para testes
- [ ] Validação de PIS/PASEP
- [ ] Gerador de PIS/PASEP válido para testes
- [ ] Validação de Título de Eleito
- [ ] Geração de Título de Eleitor válido para testes
- [ ] Validação de Placa
- [ ] Consulta CEP (fonte: BrasilAberto)
- [ ] Consulta CEP (fonte: ViaCEP)
- [ ] Consulta CEP (fonte: OpenCEP)
- [ ] Consulta CEP (fonte: BrasilAPI)
- [ ] Consulta CEP (fonte: APICep)
- [ ] Consulta RegistroBr (validação de domínios)
- [ ] Consulta ISBN

### Instalação

Use go get
```sh
go get github.com/brunoPetrim/gobr
```

### Como usar

```go
package main

import (
	"fmt"

	gobr "github.com/brunoPetrim/gobr/validation"
)

func main() {

	// Validações
	fmt.Println(gobr.IsCPFValid(""))
	fmt.Println(gobr.IsCNPJValid(""))
	fmt.Println(gobr.IsCNHValid(""))
	fmt.Println(gobr.IsRENAVANValid(""))
	fmt.Println(gobr.IsPlateValid(""))

}
```

## en
Golang module with various validations of documents and specificities of Brazil. Features included:

- [x] CPF Validation
- [x] CNPJ Validation
- [ ] CPF Formatting
- [ ] CNPJ Formatting
- [ ] CPF Deformatting
- [ ] CNPJ Deformatting
- [ ] CPF Generation (valid for tests)
- [ ] CNPJ Generation (valid for tests)
- [ ] CNH Validation
- [ ] CNH Generation (valid for tests)
- [ ] RENAVAN Validation
- [ ] RENAVAN Generation (valid for tests)
- [ ] PIS/PASEP Validation
- [ ] PIS/PASEP Generation (valid for tests)
- [ ] Voter ID Validation
- [ ] Voter ID Generation (valid for tests)
- [ ] Plate Validation
- [ ] Zip Code (CEP) Query (source: BrasilAberto)
- [ ] Zip Code (CEP) Query (source: ViaCEP)
- [ ] Zip Code (CEP) Query (source: OpenCEP)
- [ ] Zip Code (CEP) Query (source: BrasilAPI)
- [ ] Zip Code (CEP) Query (source: APICep)
- [ ] RegistroBr Query (domain validation)
- [ ] ISBN Query

### Installation

Use go get
```sh
go get github.com/brunoPetrim/gobr
```

### Usage

```go
package main

import (
	"fmt"

	gobr "github.com/brunoPetrim/gobr/validation"
)

func main() {

	// Validations
	fmt.Println(gobr.IsCPFValid(""))
	fmt.Println(gobr.IsCNPJValid(""))
	fmt.Println(gobr.IsCNHValid(""))
	fmt.Println(gobr.IsRENAVANValid(""))
	fmt.Println(gobr.IsPlateValid(""))

}
```

#### Publishing a module:
[https://go.dev/doc/modules/publishing](https://go.dev/doc/modules/publishing)
