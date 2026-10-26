// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package retenvimdfe_v3_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TAmb string

type TCOrgaoIBGE string

type TChCTe string

type TChMDFe string

type TChNFe string

type TCNPJ string

type TCNPJOpc string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TContainer string

type TCPF string

type TCPFVar string

type TData string

type TDateTimeUTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec0303 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906Opc string

type TDec1104 string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TEmail string

type TEmit string

type TEndOrg struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/mdfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/mdfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/mdfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/mdfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/mdfe xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/mdfe CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/mdfe cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/mdfe xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/mdfe fone,omitempty" json:"fone,omitempty"`
}

type TEndReEnt struct {
	CNPJ    *TCNPJ      `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TCPF       `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	XNome   string      `xml:"http://www.portalfiscal.inf.br/mdfe xNome" json:"xNome"`
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/mdfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/mdfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/mdfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/mdfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/mdfe xMun" json:"xMun"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
}

type TEndeEmi struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/mdfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/mdfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/mdfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/mdfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/mdfe xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/mdfe CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
	Fone    *string     `xml:"http://www.portalfiscal.inf.br/mdfe fone,omitempty" json:"fone,omitempty"`
	Email   *TEmail     `xml:"http://www.portalfiscal.inf.br/mdfe email,omitempty" json:"email,omitempty"`
}

type TEnderFer struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/mdfe xLgr" json:"xLgr"`
	Nro     *string     `xml:"http://www.portalfiscal.inf.br/mdfe nro,omitempty" json:"nro,omitempty"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/mdfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro *string     `xml:"http://www.portalfiscal.inf.br/mdfe xBairro,omitempty" json:"xBairro,omitempty"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/mdfe xMun" json:"xMun"`
	CEP     string      `xml:"http://www.portalfiscal.inf.br/mdfe CEP" json:"CEP"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
}

type TEndereco struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/mdfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/mdfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/mdfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/mdfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/mdfe xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/mdfe CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
	CPais   *string     `xml:"http://www.portalfiscal.inf.br/mdfe cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string     `xml:"http://www.portalfiscal.inf.br/mdfe xPais,omitempty" json:"xPais,omitempty"`
}

type TEndernac struct {
	XLgr    string      `xml:"http://www.portalfiscal.inf.br/mdfe xLgr" json:"xLgr"`
	Nro     string      `xml:"http://www.portalfiscal.inf.br/mdfe nro" json:"nro"`
	XCpl    *string     `xml:"http://www.portalfiscal.inf.br/mdfe xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro string      `xml:"http://www.portalfiscal.inf.br/mdfe xBairro" json:"xBairro"`
	CMun    TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMun" json:"cMun"`
	XMun    string      `xml:"http://www.portalfiscal.inf.br/mdfe xMun" json:"xMun"`
	CEP     *string     `xml:"http://www.portalfiscal.inf.br/mdfe CEP,omitempty" json:"CEP,omitempty"`
	UF      TUF         `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
}

type TEnviMDFe struct {
	Versao TVerMDe `xml:"versao,attr" json:"versao"`
	IDLote TIDLote `xml:"http://www.portalfiscal.inf.br/mdfe idLote" json:"idLote"`
	MDFe   TMDFe   `xml:"http://www.portalfiscal.inf.br/mdfe MDFe" json:"MDFe"`
}

type TIPv4 string

type TIDLote string

type TIe string

type TIeDest string

type TJust string

type TLatitude string

type TLocal struct {
	CMun TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMun" json:"cMun"`
	XMun string      `xml:"http://www.portalfiscal.inf.br/mdfe xMun" json:"xMun"`
	UF   TUF         `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
}

type TLongitude string

type TMDFeInfMDFeIdeInfMunCarrega struct {
	CMunCarrega TCodMunIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cMunCarrega" json:"cMunCarrega"`
	XMunCarrega string      `xml:"http://www.portalfiscal.inf.br/mdfe xMunCarrega" json:"xMunCarrega"`
}

type TMDFeInfMDFeIdeInfPercurso struct {
	UFPer TUF `xml:"http://www.portalfiscal.inf.br/mdfe UFPer" json:"UFPer"`
}

type TMDFeInfMDFeIde struct {
	CUF                 TCodUFIBGE                     `xml:"http://www.portalfiscal.inf.br/mdfe cUF" json:"cUF"`
	TpAmb               TAmb                           `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	TpEmit              TEmit                          `xml:"http://www.portalfiscal.inf.br/mdfe tpEmit" json:"tpEmit"`
	TpTransp            *TTransp                       `xml:"http://www.portalfiscal.inf.br/mdfe tpTransp,omitempty" json:"tpTransp,omitempty"`
	Mod                 TModMD                         `xml:"http://www.portalfiscal.inf.br/mdfe mod" json:"mod"`
	Serie               TSerie                         `xml:"http://www.portalfiscal.inf.br/mdfe serie" json:"serie"`
	NMDF                TNF                            `xml:"http://www.portalfiscal.inf.br/mdfe nMDF" json:"nMDF"`
	CMDF                string                         `xml:"http://www.portalfiscal.inf.br/mdfe cMDF" json:"cMDF"`
	CDV                 string                         `xml:"http://www.portalfiscal.inf.br/mdfe cDV" json:"cDV"`
	Modal               TModalMD                       `xml:"http://www.portalfiscal.inf.br/mdfe modal" json:"modal"`
	DhEmi               TDateTimeUTC                   `xml:"http://www.portalfiscal.inf.br/mdfe dhEmi" json:"dhEmi"`
	TpEmis              string                         `xml:"http://www.portalfiscal.inf.br/mdfe tpEmis" json:"tpEmis"`
	ProcEmi             string                         `xml:"http://www.portalfiscal.inf.br/mdfe procEmi" json:"procEmi"`
	VerProc             string                         `xml:"http://www.portalfiscal.inf.br/mdfe verProc" json:"verProc"`
	UFIni               TUF                            `xml:"http://www.portalfiscal.inf.br/mdfe UFIni" json:"UFIni"`
	UFFim               TUF                            `xml:"http://www.portalfiscal.inf.br/mdfe UFFim" json:"UFFim"`
	InfMunCarrega       []TMDFeInfMDFeIdeInfMunCarrega `xml:"http://www.portalfiscal.inf.br/mdfe infMunCarrega,omitempty" json:"infMunCarrega,omitempty"`
	InfPercurso         []TMDFeInfMDFeIdeInfPercurso   `xml:"http://www.portalfiscal.inf.br/mdfe infPercurso,omitempty" json:"infPercurso,omitempty"`
	DhIniViagem         *TDateTimeUTC                  `xml:"http://www.portalfiscal.inf.br/mdfe dhIniViagem,omitempty" json:"dhIniViagem,omitempty"`
	IndCanalVerde       *string                        `xml:"http://www.portalfiscal.inf.br/mdfe indCanalVerde,omitempty" json:"indCanalVerde,omitempty"`
	IndCarregaPosterior *string                        `xml:"http://www.portalfiscal.inf.br/mdfe indCarregaPosterior,omitempty" json:"indCarregaPosterior,omitempty"`
}

type TMDFeInfMDFeEmit struct {
	CNPJ      *TCNPJ   `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF    `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	IE        *string  `xml:"http://www.portalfiscal.inf.br/mdfe IE,omitempty" json:"IE,omitempty"`
	XNome     string   `xml:"http://www.portalfiscal.inf.br/mdfe xNome" json:"xNome"`
	XFant     *string  `xml:"http://www.portalfiscal.inf.br/mdfe xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndeEmi `xml:"http://www.portalfiscal.inf.br/mdfe enderEmit" json:"enderEmit"`
}

type TMDFeInfMDFeInfModal struct {
	VersaoModal string       `xml:"versaoModal,attr" json:"versaoModal"`
	Any         []RawElement `xml:",any" json:"any,omitempty"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfCTePeri struct {
	NONU      string  `xml:"http://www.portalfiscal.inf.br/mdfe nONU" json:"nONU"`
	XNomeAE   *string `xml:"http://www.portalfiscal.inf.br/mdfe xNomeAE,omitempty" json:"xNomeAE,omitempty"`
	XClaRisco *string `xml:"http://www.portalfiscal.inf.br/mdfe xClaRisco,omitempty" json:"xClaRisco,omitempty"`
	GrEmb     *string `xml:"http://www.portalfiscal.inf.br/mdfe grEmb,omitempty" json:"grEmb,omitempty"`
	QTotProd  string  `xml:"http://www.portalfiscal.inf.br/mdfe qTotProd" json:"qTotProd"`
	QVolTipo  *string `xml:"http://www.portalfiscal.inf.br/mdfe qVolTipo,omitempty" json:"qVolTipo,omitempty"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfCTeInfEntregaParcial struct {
	QtdTotal   TDec1104 `xml:"http://www.portalfiscal.inf.br/mdfe qtdTotal" json:"qtdTotal"`
	QtdParcial TDec1104 `xml:"http://www.portalfiscal.inf.br/mdfe qtdParcial" json:"qtdParcial"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfCTeInfNFePrestParcial struct {
	ChNFe TChCTe `xml:"http://www.portalfiscal.inf.br/mdfe chNFe" json:"chNFe"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfCTe struct {
	ChCTe               TChCTe                                                     `xml:"http://www.portalfiscal.inf.br/mdfe chCTe" json:"chCTe"`
	SegCodBarra         *TSegCodBarra                                              `xml:"http://www.portalfiscal.inf.br/mdfe SegCodBarra,omitempty" json:"SegCodBarra,omitempty"`
	IndReentrega        *string                                                    `xml:"http://www.portalfiscal.inf.br/mdfe indReentrega,omitempty" json:"indReentrega,omitempty"`
	InfUnidTransp       []TUnidadeTransp                                           `xml:"http://www.portalfiscal.inf.br/mdfe infUnidTransp,omitempty" json:"infUnidTransp,omitempty"`
	Peri                []TMDFeInfMDFeInfDocInfMunDescargaInfCTePeri               `xml:"http://www.portalfiscal.inf.br/mdfe peri,omitempty" json:"peri,omitempty"`
	InfEntregaParcial   *TMDFeInfMDFeInfDocInfMunDescargaInfCTeInfEntregaParcial   `xml:"http://www.portalfiscal.inf.br/mdfe infEntregaParcial,omitempty" json:"infEntregaParcial,omitempty"`
	IndPrestacaoParcial *string                                                    `xml:"http://www.portalfiscal.inf.br/mdfe indPrestacaoParcial,omitempty" json:"indPrestacaoParcial,omitempty"`
	InfNFePrestParcial  []TMDFeInfMDFeInfDocInfMunDescargaInfCTeInfNFePrestParcial `xml:"http://www.portalfiscal.inf.br/mdfe infNFePrestParcial,omitempty" json:"infNFePrestParcial,omitempty"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfNFePeri struct {
	NONU      string  `xml:"http://www.portalfiscal.inf.br/mdfe nONU" json:"nONU"`
	XNomeAE   *string `xml:"http://www.portalfiscal.inf.br/mdfe xNomeAE,omitempty" json:"xNomeAE,omitempty"`
	XClaRisco *string `xml:"http://www.portalfiscal.inf.br/mdfe xClaRisco,omitempty" json:"xClaRisco,omitempty"`
	GrEmb     *string `xml:"http://www.portalfiscal.inf.br/mdfe grEmb,omitempty" json:"grEmb,omitempty"`
	QTotProd  string  `xml:"http://www.portalfiscal.inf.br/mdfe qTotProd" json:"qTotProd"`
	QVolTipo  *string `xml:"http://www.portalfiscal.inf.br/mdfe qVolTipo,omitempty" json:"qVolTipo,omitempty"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfNFe struct {
	ChNFe         TChNFe                                       `xml:"http://www.portalfiscal.inf.br/mdfe chNFe" json:"chNFe"`
	SegCodBarra   *TSegCodBarra                                `xml:"http://www.portalfiscal.inf.br/mdfe SegCodBarra,omitempty" json:"SegCodBarra,omitempty"`
	IndReentrega  *string                                      `xml:"http://www.portalfiscal.inf.br/mdfe indReentrega,omitempty" json:"indReentrega,omitempty"`
	InfUnidTransp []TUnidadeTransp                             `xml:"http://www.portalfiscal.inf.br/mdfe infUnidTransp,omitempty" json:"infUnidTransp,omitempty"`
	Peri          []TMDFeInfMDFeInfDocInfMunDescargaInfNFePeri `xml:"http://www.portalfiscal.inf.br/mdfe peri,omitempty" json:"peri,omitempty"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfMDFeTranspPeri struct {
	NONU      string  `xml:"http://www.portalfiscal.inf.br/mdfe nONU" json:"nONU"`
	XNomeAE   *string `xml:"http://www.portalfiscal.inf.br/mdfe xNomeAE,omitempty" json:"xNomeAE,omitempty"`
	XClaRisco *string `xml:"http://www.portalfiscal.inf.br/mdfe xClaRisco,omitempty" json:"xClaRisco,omitempty"`
	GrEmb     *string `xml:"http://www.portalfiscal.inf.br/mdfe grEmb,omitempty" json:"grEmb,omitempty"`
	QTotProd  string  `xml:"http://www.portalfiscal.inf.br/mdfe qTotProd" json:"qTotProd"`
	QVolTipo  *string `xml:"http://www.portalfiscal.inf.br/mdfe qVolTipo,omitempty" json:"qVolTipo,omitempty"`
}

type TMDFeInfMDFeInfDocInfMunDescargaInfMDFeTransp struct {
	ChMDFe        TChNFe                                              `xml:"http://www.portalfiscal.inf.br/mdfe chMDFe" json:"chMDFe"`
	IndReentrega  *string                                             `xml:"http://www.portalfiscal.inf.br/mdfe indReentrega,omitempty" json:"indReentrega,omitempty"`
	InfUnidTransp []TUnidadeTransp                                    `xml:"http://www.portalfiscal.inf.br/mdfe infUnidTransp,omitempty" json:"infUnidTransp,omitempty"`
	Peri          []TMDFeInfMDFeInfDocInfMunDescargaInfMDFeTranspPeri `xml:"http://www.portalfiscal.inf.br/mdfe peri,omitempty" json:"peri,omitempty"`
}

type TMDFeInfMDFeInfDocInfMunDescarga struct {
	CMunDescarga  TCodMunIBGE                                     `xml:"http://www.portalfiscal.inf.br/mdfe cMunDescarga" json:"cMunDescarga"`
	XMunDescarga  string                                          `xml:"http://www.portalfiscal.inf.br/mdfe xMunDescarga" json:"xMunDescarga"`
	InfCTe        []TMDFeInfMDFeInfDocInfMunDescargaInfCTe        `xml:"http://www.portalfiscal.inf.br/mdfe infCTe,omitempty" json:"infCTe,omitempty"`
	InfNFe        []TMDFeInfMDFeInfDocInfMunDescargaInfNFe        `xml:"http://www.portalfiscal.inf.br/mdfe infNFe,omitempty" json:"infNFe,omitempty"`
	InfMDFeTransp []TMDFeInfMDFeInfDocInfMunDescargaInfMDFeTransp `xml:"http://www.portalfiscal.inf.br/mdfe infMDFeTransp,omitempty" json:"infMDFeTransp,omitempty"`
}

type TMDFeInfMDFeInfDoc struct {
	InfMunDescarga []TMDFeInfMDFeInfDocInfMunDescarga `xml:"http://www.portalfiscal.inf.br/mdfe infMunDescarga,omitempty" json:"infMunDescarga,omitempty"`
}

type TMDFeInfMDFeSegInfResp struct {
	RespSeg string `xml:"http://www.portalfiscal.inf.br/mdfe respSeg" json:"respSeg"`
	CNPJ    *TCNPJ `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TCPF  `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
}

type TMDFeInfMDFeSegInfSeg struct {
	XSeg string   `xml:"http://www.portalfiscal.inf.br/mdfe xSeg" json:"xSeg"`
	CNPJ TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ" json:"CNPJ"`
}

type TMDFeInfMDFeSeg struct {
	InfResp TMDFeInfMDFeSegInfResp `xml:"http://www.portalfiscal.inf.br/mdfe infResp" json:"infResp"`
	InfSeg  *TMDFeInfMDFeSegInfSeg `xml:"http://www.portalfiscal.inf.br/mdfe infSeg,omitempty" json:"infSeg,omitempty"`
	NApol   *string                `xml:"http://www.portalfiscal.inf.br/mdfe nApol,omitempty" json:"nApol,omitempty"`
	NAver   []string               `xml:"http://www.portalfiscal.inf.br/mdfe nAver,omitempty" json:"nAver,omitempty"`
}

type TMDFeInfMDFeProdPredInfLotacaoInfLocalCarrega struct {
	CEP       *string     `xml:"http://www.portalfiscal.inf.br/mdfe CEP,omitempty" json:"CEP,omitempty"`
	Latitude  *TLatitude  `xml:"http://www.portalfiscal.inf.br/mdfe latitude,omitempty" json:"latitude,omitempty"`
	Longitude *TLongitude `xml:"http://www.portalfiscal.inf.br/mdfe longitude,omitempty" json:"longitude,omitempty"`
}

type TMDFeInfMDFeProdPredInfLotacaoInfLocalDescarrega struct {
	CEP       *string     `xml:"http://www.portalfiscal.inf.br/mdfe CEP,omitempty" json:"CEP,omitempty"`
	Latitude  *TLatitude  `xml:"http://www.portalfiscal.inf.br/mdfe latitude,omitempty" json:"latitude,omitempty"`
	Longitude *TLongitude `xml:"http://www.portalfiscal.inf.br/mdfe longitude,omitempty" json:"longitude,omitempty"`
}

type TMDFeInfMDFeProdPredInfLotacao struct {
	InfLocalCarrega    TMDFeInfMDFeProdPredInfLotacaoInfLocalCarrega    `xml:"http://www.portalfiscal.inf.br/mdfe infLocalCarrega" json:"infLocalCarrega"`
	InfLocalDescarrega TMDFeInfMDFeProdPredInfLotacaoInfLocalDescarrega `xml:"http://www.portalfiscal.inf.br/mdfe infLocalDescarrega" json:"infLocalDescarrega"`
}

type TMDFeInfMDFeProdPred struct {
	TpCarga    string                          `xml:"http://www.portalfiscal.inf.br/mdfe tpCarga" json:"tpCarga"`
	XProd      string                          `xml:"http://www.portalfiscal.inf.br/mdfe xProd" json:"xProd"`
	CEAN       *string                         `xml:"http://www.portalfiscal.inf.br/mdfe cEAN,omitempty" json:"cEAN,omitempty"`
	NCM        *string                         `xml:"http://www.portalfiscal.inf.br/mdfe NCM,omitempty" json:"NCM,omitempty"`
	InfLotacao *TMDFeInfMDFeProdPredInfLotacao `xml:"http://www.portalfiscal.inf.br/mdfe infLotacao,omitempty" json:"infLotacao,omitempty"`
}

type TMDFeInfMDFeTot struct {
	QCTe   *string  `xml:"http://www.portalfiscal.inf.br/mdfe qCTe,omitempty" json:"qCTe,omitempty"`
	QNFe   *string  `xml:"http://www.portalfiscal.inf.br/mdfe qNFe,omitempty" json:"qNFe,omitempty"`
	QMDFe  *string  `xml:"http://www.portalfiscal.inf.br/mdfe qMDFe,omitempty" json:"qMDFe,omitempty"`
	VCarga TDec1302 `xml:"http://www.portalfiscal.inf.br/mdfe vCarga" json:"vCarga"`
	CUnid  string   `xml:"http://www.portalfiscal.inf.br/mdfe cUnid" json:"cUnid"`
	QCarga TDec1104 `xml:"http://www.portalfiscal.inf.br/mdfe qCarga" json:"qCarga"`
}

type TMDFeInfMDFeLacres struct {
	NLacre string `xml:"http://www.portalfiscal.inf.br/mdfe nLacre" json:"nLacre"`
}

type TMDFeInfMDFeAutXML struct {
	CNPJ *TCNPJ `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF  *TCPF  `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
}

type TMDFeInfMDFeInfAdic struct {
	InfAdFisco *string `xml:"http://www.portalfiscal.inf.br/mdfe infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     *string `xml:"http://www.portalfiscal.inf.br/mdfe infCpl,omitempty" json:"infCpl,omitempty"`
}

type TMDFeInfMDFeInfSolicNFF struct {
	XSolic string `xml:"http://www.portalfiscal.inf.br/mdfe xSolic" json:"xSolic"`
}

type TMDFeInfMDFeInfPAAPAASignature struct {
	SignatureValue Base64Binary     `xml:"http://www.portalfiscal.inf.br/mdfe SignatureValue" json:"SignatureValue"`
	RSAKeyValue    TRSAKeyValueType `xml:"http://www.portalfiscal.inf.br/mdfe RSAKeyValue" json:"RSAKeyValue"`
}

type TMDFeInfMDFeInfPAA struct {
	CNPJPAA      TCNPJ                          `xml:"http://www.portalfiscal.inf.br/mdfe CNPJPAA" json:"CNPJPAA"`
	PAASignature TMDFeInfMDFeInfPAAPAASignature `xml:"http://www.portalfiscal.inf.br/mdfe PAASignature" json:"PAASignature"`
}

type TMDFeInfMDFe struct {
	Versao      TVerMDe                  `xml:"versao,attr" json:"versao"`
	ID          string                   `xml:"Id,attr" json:"Id"`
	Ide         TMDFeInfMDFeIde          `xml:"http://www.portalfiscal.inf.br/mdfe ide" json:"ide"`
	Emit        TMDFeInfMDFeEmit         `xml:"http://www.portalfiscal.inf.br/mdfe emit" json:"emit"`
	InfModal    TMDFeInfMDFeInfModal     `xml:"http://www.portalfiscal.inf.br/mdfe infModal" json:"infModal"`
	InfDoc      TMDFeInfMDFeInfDoc       `xml:"http://www.portalfiscal.inf.br/mdfe infDoc" json:"infDoc"`
	Seg         []TMDFeInfMDFeSeg        `xml:"http://www.portalfiscal.inf.br/mdfe seg,omitempty" json:"seg,omitempty"`
	ProdPred    *TMDFeInfMDFeProdPred    `xml:"http://www.portalfiscal.inf.br/mdfe prodPred,omitempty" json:"prodPred,omitempty"`
	Tot         TMDFeInfMDFeTot          `xml:"http://www.portalfiscal.inf.br/mdfe tot" json:"tot"`
	Lacres      []TMDFeInfMDFeLacres     `xml:"http://www.portalfiscal.inf.br/mdfe lacres,omitempty" json:"lacres,omitempty"`
	AutXML      []TMDFeInfMDFeAutXML     `xml:"http://www.portalfiscal.inf.br/mdfe autXML,omitempty" json:"autXML,omitempty"`
	InfAdic     *TMDFeInfMDFeInfAdic     `xml:"http://www.portalfiscal.inf.br/mdfe infAdic,omitempty" json:"infAdic,omitempty"`
	InfRespTec  *TRespTec                `xml:"http://www.portalfiscal.inf.br/mdfe infRespTec,omitempty" json:"infRespTec,omitempty"`
	InfSolicNFF *TMDFeInfMDFeInfSolicNFF `xml:"http://www.portalfiscal.inf.br/mdfe infSolicNFF,omitempty" json:"infSolicNFF,omitempty"`
	InfPAA      *TMDFeInfMDFeInfPAA      `xml:"http://www.portalfiscal.inf.br/mdfe infPAA,omitempty" json:"infPAA,omitempty"`
}

type TMDFeInfMDFeSupl struct {
	QRCodMDFe string `xml:"http://www.portalfiscal.inf.br/mdfe qrCodMDFe" json:"qrCodMDFe"`
}

type TMDFe struct {
	InfMDFe     TMDFeInfMDFe      `xml:"http://www.portalfiscal.inf.br/mdfe infMDFe" json:"infMDFe"`
	InfMDFeSupl *TMDFeInfMDFeSupl `xml:"http://www.portalfiscal.inf.br/mdfe infMDFeSupl,omitempty" json:"infMDFeSupl,omitempty"`
	Signature   SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TMed string

type TModDoc string

type TModMD string

type TModalMD string

type TMotivo string

type TNF string

type TNFeNFInfNFe struct {
	ChNFe TChNFe `xml:"http://www.portalfiscal.inf.br/mdfe chNFe" json:"chNFe"`
	PIN   *TPIN  `xml:"http://www.portalfiscal.inf.br/mdfe PIN,omitempty" json:"PIN,omitempty"`
}

type TNFeNFInfNFEmi struct {
	CNPJ  TCNPJ  `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ" json:"CNPJ"`
	XNome string `xml:"http://www.portalfiscal.inf.br/mdfe xNome" json:"xNome"`
	UF    TUF    `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
}

type TNFeNFInfNFDest struct {
	CNPJ  *TCNPJOpc `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF   *TCPF     `xml:"http://www.portalfiscal.inf.br/mdfe CPF,omitempty" json:"CPF,omitempty"`
	XNome string    `xml:"http://www.portalfiscal.inf.br/mdfe xNome" json:"xNome"`
	UF    TUF       `xml:"http://www.portalfiscal.inf.br/mdfe UF" json:"UF"`
}

type TNFeNFInfNF struct {
	Emi   TNFeNFInfNFEmi  `xml:"http://www.portalfiscal.inf.br/mdfe emi" json:"emi"`
	Dest  TNFeNFInfNFDest `xml:"http://www.portalfiscal.inf.br/mdfe dest" json:"dest"`
	Serie string          `xml:"http://www.portalfiscal.inf.br/mdfe serie" json:"serie"`
	NNF   string          `xml:"http://www.portalfiscal.inf.br/mdfe nNF" json:"nNF"`
	DEmi  TData           `xml:"http://www.portalfiscal.inf.br/mdfe dEmi" json:"dEmi"`
	VNF   TDec1302        `xml:"http://www.portalfiscal.inf.br/mdfe vNF" json:"vNF"`
	PIN   *TPIN           `xml:"http://www.portalfiscal.inf.br/mdfe PIN,omitempty" json:"PIN,omitempty"`
}

type TNFeNF struct {
	InfNFe *TNFeNFInfNFe `xml:"http://www.portalfiscal.inf.br/mdfe infNFe,omitempty" json:"infNFe,omitempty"`
	InfNF  *TNFeNFInfNF  `xml:"http://www.portalfiscal.inf.br/mdfe infNF,omitempty" json:"infNF,omitempty"`
}

type TNSU string

type TPIN string

type TPlaca string

type TProcEmi string

type TProt string

type TProtMDFeInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	VerAplic TVerAplic        `xml:"http://www.portalfiscal.inf.br/mdfe verAplic" json:"verAplic"`
	ChMDFe   TChMDFe          `xml:"http://www.portalfiscal.inf.br/mdfe chMDFe" json:"chMDFe"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/mdfe dhRecbto" json:"dhRecbto"`
	NProt    *TProt           `xml:"http://www.portalfiscal.inf.br/mdfe nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/mdfe digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/mdfe cStat" json:"cStat"`
	XMotivo  TMotivo          `xml:"http://www.portalfiscal.inf.br/mdfe xMotivo" json:"xMotivo"`
}

type TProtMDFeInfFisco struct {
	CMsg string  `xml:"http://www.portalfiscal.inf.br/mdfe cMsg" json:"cMsg"`
	XMsg TMotivo `xml:"http://www.portalfiscal.inf.br/mdfe xMsg" json:"xMsg"`
}

type TProtMDFe struct {
	Versao    TVerMDe            `xml:"versao,attr" json:"versao"`
	InfProt   TProtMDFeInfProt   `xml:"http://www.portalfiscal.inf.br/mdfe infProt" json:"infProt"`
	InfFisco  *TProtMDFeInfFisco `xml:"http://www.portalfiscal.inf.br/mdfe infFisco,omitempty" json:"infFisco,omitempty"`
	Signature *SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRSAKeyValueType struct {
	Modulus  Base64Binary `xml:"http://www.portalfiscal.inf.br/mdfe Modulus" json:"Modulus"`
	Exponent Base64Binary `xml:"http://www.portalfiscal.inf.br/mdfe Exponent" json:"Exponent"`
}

type TRec string

type TRespTec struct {
	CNPJ     TCNPJ   `xml:"http://www.portalfiscal.inf.br/mdfe CNPJ" json:"CNPJ"`
	XContato string  `xml:"http://www.portalfiscal.inf.br/mdfe xContato" json:"xContato"`
	Email    TEmail  `xml:"http://www.portalfiscal.inf.br/mdfe email" json:"email"`
	Fone     string  `xml:"http://www.portalfiscal.inf.br/mdfe fone" json:"fone"`
	IDCSRT   *string `xml:"http://www.portalfiscal.inf.br/mdfe idCSRT,omitempty" json:"idCSRT,omitempty"`
	HashCSRT *string `xml:"http://www.portalfiscal.inf.br/mdfe hashCSRT,omitempty" json:"hashCSRT,omitempty"`
}

type TRetEnviMDFeInfRec struct {
	NRec     TRec     `xml:"http://www.portalfiscal.inf.br/mdfe nRec" json:"nRec"`
	DhRecbto DateTime `xml:"http://www.portalfiscal.inf.br/mdfe dhRecbto" json:"dhRecbto"`
	TMed     string   `xml:"http://www.portalfiscal.inf.br/mdfe tMed" json:"tMed"`
}

type TRetEnviMDFe struct {
	Versao   TVerMDe             `xml:"versao,attr" json:"versao"`
	TpAmb    RawElement          `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE          `xml:"http://www.portalfiscal.inf.br/mdfe cUF" json:"cUF"`
	VerAplic TVerAplic           `xml:"http://www.portalfiscal.inf.br/mdfe verAplic" json:"verAplic"`
	CStat    TStat               `xml:"http://www.portalfiscal.inf.br/mdfe cStat" json:"cStat"`
	XMotivo  TMotivo             `xml:"http://www.portalfiscal.inf.br/mdfe xMotivo" json:"xMotivo"`
	InfRec   *TRetEnviMDFeInfRec `xml:"http://www.portalfiscal.inf.br/mdfe infRec,omitempty" json:"infRec,omitempty"`
}

type TRetMDFe struct {
	Versao   TVerMDe    `xml:"versao,attr" json:"versao"`
	TpAmb    RawElement `xml:"http://www.portalfiscal.inf.br/mdfe tpAmb" json:"tpAmb"`
	CUF      TCodUFIBGE `xml:"http://www.portalfiscal.inf.br/mdfe cUF" json:"cUF"`
	VerAplic TVerAplic  `xml:"http://www.portalfiscal.inf.br/mdfe verAplic" json:"verAplic"`
	CStat    TStat      `xml:"http://www.portalfiscal.inf.br/mdfe cStat" json:"cStat"`
	XMotivo  TMotivo    `xml:"http://www.portalfiscal.inf.br/mdfe xMotivo" json:"xMotivo"`
	ProtMDFe *TProtMDFe `xml:"http://www.portalfiscal.inf.br/mdfe protMDFe,omitempty" json:"protMDFe,omitempty"`
}

type TSegCodBarra string

type TSerie string

type TServ string

type TStat string

type TString string

type TTime string

type TTransp string

type TUF string

type TUnidCargaLacUnidCarga struct {
	NLacre string `xml:"http://www.portalfiscal.inf.br/mdfe nLacre" json:"nLacre"`
}

type TUnidCarga struct {
	TpUnidCarga  TtipoUnidCarga           `xml:"http://www.portalfiscal.inf.br/mdfe tpUnidCarga" json:"tpUnidCarga"`
	IDUnidCarga  TContainer               `xml:"http://www.portalfiscal.inf.br/mdfe idUnidCarga" json:"idUnidCarga"`
	LacUnidCarga []TUnidCargaLacUnidCarga `xml:"http://www.portalfiscal.inf.br/mdfe lacUnidCarga,omitempty" json:"lacUnidCarga,omitempty"`
	QtdRat       *string                  `xml:"http://www.portalfiscal.inf.br/mdfe qtdRat,omitempty" json:"qtdRat,omitempty"`
}

type TUnidadeTranspLacUnidTransp struct {
	NLacre string `xml:"http://www.portalfiscal.inf.br/mdfe nLacre" json:"nLacre"`
}

type TUnidadeTransp struct {
	TpUnidTransp  TtipoUnidTransp               `xml:"http://www.portalfiscal.inf.br/mdfe tpUnidTransp" json:"tpUnidTransp"`
	IDUnidTransp  TContainer                    `xml:"http://www.portalfiscal.inf.br/mdfe idUnidTransp" json:"idUnidTransp"`
	LacUnidTransp []TUnidadeTranspLacUnidTransp `xml:"http://www.portalfiscal.inf.br/mdfe lacUnidTransp,omitempty" json:"lacUnidTransp,omitempty"`
	InfUnidCarga  []TUnidCarga                  `xml:"http://www.portalfiscal.inf.br/mdfe infUnidCarga,omitempty" json:"infUnidCarga,omitempty"`
	QtdRat        *TDec03020303                 `xml:"http://www.portalfiscal.inf.br/mdfe qtdRat,omitempty" json:"qtdRat,omitempty"`
}

type TVerAplic string

type TVerMDe string

type Tano string

type TtipoUnidCarga string

type TtipoUnidTransp string

type DigestValueType string

type KeyInfoType struct {
	ID       *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	X509Data X509DataType `xml:"http://www.w3.org/2000/09/xmldsig# X509Data" json:"X509Data"`
}

type ReferenceTypeDigestMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type ReferenceType struct {
	ID           *ID                       `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	URI          string                    `xml:"URI,attr" json:"URI"`
	TypeField    *AnyURI                   `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	Transforms   TransformsType            `xml:"http://www.w3.org/2000/09/xmldsig# Transforms" json:"Transforms"`
	DigestMethod ReferenceTypeDigestMethod `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod" json:"DigestMethod"`
	DigestValue  DigestValueType           `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue" json:"DigestValue"`
}

type SignatureType struct {
	ID             *ID                `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	SignedInfo     SignedInfoType     `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo" json:"SignedInfo"`
	SignatureValue SignatureValueType `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue" json:"SignatureValue"`
	KeyInfo        KeyInfoType        `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo" json:"KeyInfo"`
}

type SignatureValueType struct {
	Value Base64Binary `xml:",chardata" json:"value"`
	ID    *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
}

type SignedInfoTypeCanonicalizationMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type SignedInfoTypeSignatureMethod struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
}

type SignedInfoType struct {
	ID                     *ID                                  `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	CanonicalizationMethod SignedInfoTypeCanonicalizationMethod `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod" json:"CanonicalizationMethod"`
	SignatureMethod        SignedInfoTypeSignatureMethod        `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod" json:"SignatureMethod"`
	Reference              ReferenceType                        `xml:"http://www.w3.org/2000/09/xmldsig# Reference" json:"Reference"`
}

type TTransformURI string

type TransformType struct {
	Algorithm TTransformURI `xml:"Algorithm,attr" json:"Algorithm"`
	XPath     []String      `xml:"http://www.w3.org/2000/09/xmldsig# XPath,omitempty" json:"XPath,omitempty"`
}

type TransformsType struct {
	Transform []TransformType `xml:"http://www.w3.org/2000/09/xmldsig# Transform,omitempty" json:"Transform,omitempty"`
}

type X509DataType struct {
	X509Certificate Base64Binary `xml:"http://www.w3.org/2000/09/xmldsig# X509Certificate" json:"X509Certificate"`
}

type ID string

type AnyURI string

type Base64Binary string

type DateTime string

type Integer string

type String string

// RetEnviMDFe represents the global element {http://www.portalfiscal.inf.br/mdfe}retEnviMDFe.
type RetEnviMDFe struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/mdfe retEnviMDFe" json:"-"`
	TRetEnviMDFe
}

// Signature represents the global element {http://www.w3.org/2000/09/xmldsig#}Signature.
type Signature struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"-"`
	SignatureType
}

// MarshalXML avoids re-emitting a default xmlns already supplied by encoding/xml.
func (value RawElement) MarshalXML(encoder *xml.Encoder, start xml.StartElement) error {
	if value.XMLName.Local != "" {
		start.Name = value.XMLName
	}
	for _, attr := range value.Attrs {
		if attr.Name.Space == "" && attr.Name.Local == "xmlns" && attr.Value == start.Name.Space {
			continue
		}
		if attr.Name.Space == "xmlns" {
			attr.Name = xml.Name{Local: "xmlns:" + attr.Name.Local}
		}
		start.Attr = append(start.Attr, attr)
	}
	body := struct {
		InnerXML string `xml:",innerxml"`
	}{value.InnerXML}
	return encoder.EncodeElement(body, start)
}
