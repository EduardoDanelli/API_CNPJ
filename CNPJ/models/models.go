package models

type CNPJ struct {
	AtividadePrincipal []struct {
		Text string `json:"text"`
		Code string `json:"code"`
	} `json:"atividade_principal"`
	DataSituacao          string `json:"data_situacao"`
	Tipo                  string `json:"tipo"`
	Nome                  string `json:"nome"`
	Uf                    string `json:"uf"`
	Telefone              string `json:"telefone"`
	Email                 string `json:"email"`
	AtividadesSecundarias []struct {
		Text string `json:"text"`
		Code string `json:"code"`
	} `json:"atividades_secundarias"`
	Qsa []struct {
		Qual         string
		QualReplegal string
		NomeRepLegal string
		Nome         string
	} `json:"qsa"`
	Situacao             string      `json:"situacao"`
	Bairro               string      `json:"bairro"`
	Logradouro           string      `json:"logradouro"`
	Numero               string      `json:"numero"`
	Cep                  string      `json:"cep"`
	Municipio            string      `json:"municipio"`
	Porte                string      `json:"porte"`
	Abertura             string      `json:"abertura"`
	NaturezaJuridica     string      `json:"natureza_juridica"`
	Fantasia             string      `json:"fantasia"`
	Cnpj                 string      `json:"cnpj"`
	UltimaAtualizacao    string      `json:"ultima_atualizacao"`
	Status               string      `json:"status"`
	Complemento          string      `json:"complemento"`
	Efr                  string      `json:"efr"`
	MotivoSituacao       string      `json:"motivo_situacao"`
	SituacaoEspecial     string      `json:"situacao_especial"`
	DataSituacaoEspecial string      `json:"data_situacao_especial"`
	CapitalSocial        string      `json:"capital_social"`
	Extra                interface{} `json:"extra"`
	Biling               struct {
		Free     bool `json:"free"`
		Database bool `json:"database"`
	}
}
