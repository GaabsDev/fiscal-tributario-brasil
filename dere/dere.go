// Package dere implementa os contratos e tipos normativos da Declaração de Regimes Específicos (DeRE).
//
// AVISO TÉCNICO NORMATIVO (Data de Referência: 24/10/2026):
// A Declaração de Regimes Específicos (DeRE) foi instituída pela Reforma Tributária do Consumo (Lei Complementar nº 214/2025,
// arts. 165 a 240) para apuração de regimes diferenciados, monofásicos e favorecidos (combustíveis, serviços financeiros,
// bens imóveis, planos de saúde, hotelaria e parques temáticos).
// O cronograma normativo publicado pela Receita Federal do Brasil (RFB) estabelece a entrega obrigatória a partir de Novembro/2026.
// O pacote oficial de Schemas XML (XSDs) ainda não foi disponibilizado para download público no Portal da RFB até a presente data.
// Este pacote implementa os tipos de dados fundamentais, estruturas canônicas e contratos normativos definidos na Lei Complementar,
// permitindo a interoperabilidade e conformidade antecipada com as regras de apuração do IBS e da CBS.
package dere

import (
	"encoding/json"
	"time"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

const (
	VersaoDeRE           = "1.00"
	StatusRegulamentacao = "Cronograma RFB previsto para Nov/2026 - Aguardando publicação pública do pacote XSD"
)

// TipoRegimeEspecifico identifica a categoria do regime tributário diferenciado.
type TipoRegimeEspecifico string

const (
	RegimeCombustiveisMono    TipoRegimeEspecifico = "COMBUSTIVEIS_MONOFASICO"
	RegimeServicosFinanceiros TipoRegimeEspecifico = "SERVICOS_FINANCEIROS"
	RegimeBensImoveis         TipoRegimeEspecifico = "BENS_IMOVEIS"
	RegimePlanosSaude         TipoRegimeEspecifico = "PLANOS_SAUDE"
	RegimeHotelariaParques    TipoRegimeEspecifico = "HOTELARIA_PARQUES"
)

// DeclaracaoRegimesEspecifcos estrutura o contrato canônico da declaração de apuração periódica da DeRE.
type DeclaracaoRegimesEspecifcos struct {
	Versao          string           `json:"versao"`
	CNPJDeclarante  nfe.CNPJ         `json:"cnpj_declarante"`
	PeriodoApuracao string           `json:"periodo_apuracao"` // AAAA-MM
	DataEnvio       time.Time        `json:"data_envio"`
	Regimes         []ApuracaoRegime `json:"regimes"`
	Totalizacao     TotalizacaoDeRE  `json:"totalizacao"`
	AvisoTecnico    string           `json:"aviso_tecnico"`
}

// ApuracaoRegime detalha os cálculos específicos de um regime particular.
type ApuracaoRegime struct {
	Tipo             TipoRegimeEspecifico `json:"tipo"`
	CClassTrib       nfe.CClassTrib       `json:"cClassTrib"`
	BaseCalculo      nfe.Decimal          `json:"base_calculo"`
	ValorApuradoIBS  nfe.Decimal          `json:"valor_apurado_ibs"`
	ValorApuradoCBS  nfe.Decimal          `json:"valor_apurado_cbs"`
	Deducoes         *nfe.Decimal         `json:"deducoes,omitempty"`
	CreditoPresumido *nfe.Decimal         `json:"credito_presumido,omitempty"`
}

// TotalizacaoDeRE consolida os valores devidos apurados em todos os regimes declarados.
type TotalizacaoDeRE struct {
	TotalIBS nfe.Decimal `json:"total_ibs"`
	TotalCBS nfe.Decimal `json:"total_cbs"`
}

// NewDeclaracao cria uma nova instância da DeRE com metadados normativos.
func NewDeclaracao(cnpj nfe.CNPJ, periodo string) *DeclaracaoRegimesEspecifcos {
	return &DeclaracaoRegimesEspecifcos{
		Versao:          VersaoDeRE,
		CNPJDeclarante:  cnpj,
		PeriodoApuracao: periodo,
		DataEnvio:       time.Now().UTC(),
		Regimes:         make([]ApuracaoRegime, 0),
		AvisoTecnico:    StatusRegulamentacao,
	}
}

func (d *DeclaracaoRegimesEspecifcos) ToJSON() ([]byte, error) {
	return json.MarshalIndent(d, "", "  ")
}
