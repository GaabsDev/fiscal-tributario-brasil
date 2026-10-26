// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package schemas_1_01_cnc_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TBMISSQN string

type TCAtvEvento struct {
	XNome    TSDesc255          `xml:"http://www.sped.fazenda.gov.br/nfse xNome" json:"xNome"`
	DtIni    TSData             `xml:"http://www.sped.fazenda.gov.br/nfse dtIni" json:"dtIni"`
	DtFim    TSData             `xml:"http://www.sped.fazenda.gov.br/nfse dtFim" json:"dtFim"`
	IDAtvEvt *TSIdeEvento       `xml:"http://www.sped.fazenda.gov.br/nfse idAtvEvt,omitempty" json:"idAtvEvt,omitempty"`
	End      *TCEnderecoSimples `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
}

type TCBeneficioMunicipal struct {
	NBM      TSNumBeneficioMunicipal `xml:"http://www.sped.fazenda.gov.br/nfse nBM" json:"nBM"`
	VRedBCBM *TSDec15V2              `xml:"http://www.sped.fazenda.gov.br/nfse vRedBCBM,omitempty" json:"vRedBCBM,omitempty"`
	PRedBCBM *TSDec3V2               `xml:"http://www.sped.fazenda.gov.br/nfse pRedBCBM,omitempty" json:"pRedBCBM,omitempty"`
}

type TCCServ struct {
	CTribNac    TSCodTribNac                 `xml:"http://www.sped.fazenda.gov.br/nfse cTribNac" json:"cTribNac"`
	CTribMun    *TCCodTribMun                `xml:"http://www.sped.fazenda.gov.br/nfse cTribMun,omitempty" json:"cTribMun,omitempty"`
	XDescServ   TSDesc2000                   `xml:"http://www.sped.fazenda.gov.br/nfse xDescServ" json:"xDescServ"`
	CNBS        *TSCodNBS                    `xml:"http://www.sped.fazenda.gov.br/nfse cNBS,omitempty" json:"cNBS,omitempty"`
	CIntContrib *TSCodigoInternoContribuinte `xml:"http://www.sped.fazenda.gov.br/nfse cIntContrib,omitempty" json:"cIntContrib,omitempty"`
}

type TCCodTribMun string

type TCComExterior struct {
	MdPrestacao TSModoPrestacao   `xml:"http://www.sped.fazenda.gov.br/nfse mdPrestacao" json:"mdPrestacao"`
	VincPrest   TSVincPrest       `xml:"http://www.sped.fazenda.gov.br/nfse vincPrest" json:"vincPrest"`
	TpMoeda     TSCodMoeda        `xml:"http://www.sped.fazenda.gov.br/nfse tpMoeda" json:"tpMoeda"`
	VServMoeda  TSDec15V2         `xml:"http://www.sped.fazenda.gov.br/nfse vServMoeda" json:"vServMoeda"`
	MecAFComexP TSMecAFComExPrest `xml:"http://www.sped.fazenda.gov.br/nfse mecAFComexP" json:"mecAFComexP"`
	MecAFComexT TSMecAFComExToma  `xml:"http://www.sped.fazenda.gov.br/nfse mecAFComexT" json:"mecAFComexT"`
	MovTempBens TSMovTempBens     `xml:"http://www.sped.fazenda.gov.br/nfse movTempBens" json:"movTempBens"`
	NDI         *TSNumDocImport   `xml:"http://www.sped.fazenda.gov.br/nfse nDI,omitempty" json:"nDI,omitempty"`
	NRE         *TSNumRegExport   `xml:"http://www.sped.fazenda.gov.br/nfse nRE,omitempty" json:"nRE,omitempty"`
	Mdic        TSEnvMDIC         `xml:"http://www.sped.fazenda.gov.br/nfse mdic" json:"mdic"`
}

type TCDPS struct {
	Versao    TVerNFSe       `xml:"versao,attr" json:"versao"`
	InfDPS    TCInfDPS       `xml:"http://www.sped.fazenda.gov.br/nfse infDPS" json:"infDPS"`
	Signature *SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TCDocDedRed struct {
	ChNFSe              *TSChaveNFSe     `xml:"http://www.sped.fazenda.gov.br/nfse chNFSe,omitempty" json:"chNFSe,omitempty"`
	ChNFe               *TSChaveNFe      `xml:"http://www.sped.fazenda.gov.br/nfse chNFe,omitempty" json:"chNFe,omitempty"`
	NFSeMun             *TCDocOutNFSe    `xml:"http://www.sped.fazenda.gov.br/nfse NFSeMun,omitempty" json:"NFSeMun,omitempty"`
	NFNFS               *TCDocNFNFS      `xml:"http://www.sped.fazenda.gov.br/nfse NFNFS,omitempty" json:"NFNFS,omitempty"`
	NDocFisc            *TSDesc255       `xml:"http://www.sped.fazenda.gov.br/nfse nDocFisc,omitempty" json:"nDocFisc,omitempty"`
	NDoc                *TSDesc255       `xml:"http://www.sped.fazenda.gov.br/nfse nDoc,omitempty" json:"nDoc,omitempty"`
	TpDedRed            TSIdeDedRed      `xml:"http://www.sped.fazenda.gov.br/nfse tpDedRed" json:"tpDedRed"`
	XDescOutDed         *TSDescOutDedRed `xml:"http://www.sped.fazenda.gov.br/nfse xDescOutDed,omitempty" json:"xDescOutDed,omitempty"`
	DtEmiDoc            Date             `xml:"http://www.sped.fazenda.gov.br/nfse dtEmiDoc" json:"dtEmiDoc"`
	VDedutivelRedutivel TSDec15V2        `xml:"http://www.sped.fazenda.gov.br/nfse vDedutivelRedutivel" json:"vDedutivelRedutivel"`
	VDeducaoReducao     TSDec15V2        `xml:"http://www.sped.fazenda.gov.br/nfse vDeducaoReducao" json:"vDeducaoReducao"`
	Fornec              *TCInfoPessoa    `xml:"http://www.sped.fazenda.gov.br/nfse fornec,omitempty" json:"fornec,omitempty"`
}

type TCDocNFNFS struct {
	NNFS     TSNum7Dig    `xml:"http://www.sped.fazenda.gov.br/nfse nNFS" json:"nNFS"`
	ModNFS   TSNum15Dig   `xml:"http://www.sped.fazenda.gov.br/nfse modNFS" json:"modNFS"`
	SerieNFS TSSerieNFNFS `xml:"http://www.sped.fazenda.gov.br/nfse serieNFS" json:"serieNFS"`
}

type TCDocOutNFSe struct {
	CMunNFSeMun   TSCodMunIBGE     `xml:"http://www.sped.fazenda.gov.br/nfse cMunNFSeMun" json:"cMunNFSeMun"`
	NNFSeMun      TSNum15Dig       `xml:"http://www.sped.fazenda.gov.br/nfse nNFSeMun" json:"nNFSeMun"`
	CVerifNFSeMun TSCodVerificacao `xml:"http://www.sped.fazenda.gov.br/nfse cVerifNFSeMun" json:"cVerifNFSeMun"`
}

type TCEmitente struct {
	CNPJ     *TSCNPJ            `xml:"http://www.sped.fazenda.gov.br/nfse CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *TSCPF             `xml:"http://www.sped.fazenda.gov.br/nfse CPF,omitempty" json:"CPF,omitempty"`
	IM       *TSInscMun         `xml:"http://www.sped.fazenda.gov.br/nfse IM,omitempty" json:"IM,omitempty"`
	XNome    TSNomeRazaoSocial  `xml:"http://www.sped.fazenda.gov.br/nfse xNome" json:"xNome"`
	XFant    *TSNomeFantasia    `xml:"http://www.sped.fazenda.gov.br/nfse xFant,omitempty" json:"xFant,omitempty"`
	EnderNac TCEnderecoEmitente `xml:"http://www.sped.fazenda.gov.br/nfse enderNac" json:"enderNac"`
	Fone     *TSTelefone        `xml:"http://www.sped.fazenda.gov.br/nfse fone,omitempty" json:"fone,omitempty"`
	Email    *TSEmail           `xml:"http://www.sped.fazenda.gov.br/nfse email,omitempty" json:"email,omitempty"`
}

type TCEnderContribuinteCNC struct {
	CEP  TSCEP                  `xml:"http://www.sped.fazenda.gov.br/nfse CEP" json:"CEP"`
	Nro  TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfse nro" json:"nro"`
	XCpl *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xCpl,omitempty" json:"xCpl,omitempty"`
}

type TCEnderExt struct {
	CPais       TSCodPaisISO       `xml:"http://www.sped.fazenda.gov.br/nfse cPais" json:"cPais"`
	CEndPost    TSCodigoEndPostal  `xml:"http://www.sped.fazenda.gov.br/nfse cEndPost" json:"cEndPost"`
	XCidade     TSCidade           `xml:"http://www.sped.fazenda.gov.br/nfse xCidade" json:"xCidade"`
	XEstProvReg TSEstadoProvRegiao `xml:"http://www.sped.fazenda.gov.br/nfse xEstProvReg" json:"xEstProvReg"`
}

type TCEnderExtSimples struct {
	CEndPost    TSCodigoEndPostal  `xml:"http://www.sped.fazenda.gov.br/nfse cEndPost" json:"cEndPost"`
	XCidade     TSCidade           `xml:"http://www.sped.fazenda.gov.br/nfse xCidade" json:"xCidade"`
	XEstProvReg TSEstadoProvRegiao `xml:"http://www.sped.fazenda.gov.br/nfse xEstProvReg" json:"xEstProvReg"`
}

type TCEnderNac struct {
	CMun TSCodMunIBGE `xml:"http://www.sped.fazenda.gov.br/nfse cMun" json:"cMun"`
	CEP  TSCEP        `xml:"http://www.sped.fazenda.gov.br/nfse CEP" json:"CEP"`
}

type TCEnderObraEvento struct {
	CEP     *TSCEP                 `xml:"http://www.sped.fazenda.gov.br/nfse CEP,omitempty" json:"CEP,omitempty"`
	EndExt  *TCEnderExtSimples     `xml:"http://www.sped.fazenda.gov.br/nfse endExt,omitempty" json:"endExt,omitempty"`
	XLgr    TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfse xLgr" json:"xLgr"`
	Nro     TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfse nro" json:"nro"`
	XCpl    *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfse xBairro" json:"xBairro"`
}

type TCEndereco struct {
	EndNac  *TCEnderNac            `xml:"http://www.sped.fazenda.gov.br/nfse endNac,omitempty" json:"endNac,omitempty"`
	EndExt  *TCEnderExt            `xml:"http://www.sped.fazenda.gov.br/nfse endExt,omitempty" json:"endExt,omitempty"`
	XLgr    TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfse xLgr" json:"xLgr"`
	Nro     TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfse nro" json:"nro"`
	XCpl    *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfse xBairro" json:"xBairro"`
}

type TCEnderecoEmitente struct {
	XLgr    TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfse xLgr" json:"xLgr"`
	Nro     TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfse nro" json:"nro"`
	XCpl    *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfse xBairro" json:"xBairro"`
	CMun    TSCodMunIBGE           `xml:"http://www.sped.fazenda.gov.br/nfse cMun" json:"cMun"`
	UF      TSUF                   `xml:"http://www.sped.fazenda.gov.br/nfse UF" json:"UF"`
	CEP     TSCEP                  `xml:"http://www.sped.fazenda.gov.br/nfse CEP" json:"CEP"`
}

type TCEnderecoSimples struct {
	CEP     *TSCEP                 `xml:"http://www.sped.fazenda.gov.br/nfse CEP,omitempty" json:"CEP,omitempty"`
	EndExt  *TCEnderExtSimples     `xml:"http://www.sped.fazenda.gov.br/nfse endExt,omitempty" json:"endExt,omitempty"`
	XLgr    TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfse xLgr" json:"xLgr"`
	Nro     TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfse nro" json:"nro"`
	XCpl    *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfse xBairro" json:"xBairro"`
}

type TCExigSuspensa struct {
	TpSusp    TSOpExigSuspensa      `xml:"http://www.sped.fazenda.gov.br/nfse tpSusp" json:"tpSusp"`
	NProcesso TSNumProcExigSuspensa `xml:"http://www.sped.fazenda.gov.br/nfse nProcesso" json:"nProcesso"`
}

type TCInfCNCContribuintesCnc struct {
	ContribuinteCnc []TCInfoContribuinteCNC `xml:"http://www.sped.fazenda.gov.br/nfse contribuinteCnc,omitempty" json:"contribuinteCnc,omitempty"`
}

type TCInfCNC struct {
	TpAmb            TSTipoAmbiente           `xml:"http://www.sped.fazenda.gov.br/nfse tpAmb" json:"tpAmb"`
	DhGeracaoArquivo TSDateTimeUTC            `xml:"http://www.sped.fazenda.gov.br/nfse dhGeracaoArquivo" json:"dhGeracaoArquivo"`
	VerAplic         TSVerAplic               `xml:"http://www.sped.fazenda.gov.br/nfse verAplic" json:"verAplic"`
	ContribuintesCnc TCInfCNCContribuintesCnc `xml:"http://www.sped.fazenda.gov.br/nfse contribuintesCnc" json:"contribuintesCnc"`
}

type TCInfDPS struct {
	ID            TSIDDPS          `xml:"Id,attr" json:"Id"`
	TpAmb         TSTipoAmbiente   `xml:"http://www.sped.fazenda.gov.br/nfse tpAmb" json:"tpAmb"`
	DhEmi         TSDateTimeUTC    `xml:"http://www.sped.fazenda.gov.br/nfse dhEmi" json:"dhEmi"`
	VerAplic      TSVerAplic       `xml:"http://www.sped.fazenda.gov.br/nfse verAplic" json:"verAplic"`
	Serie         TSSerieDPS       `xml:"http://www.sped.fazenda.gov.br/nfse serie" json:"serie"`
	NDPS          TSNumDPS         `xml:"http://www.sped.fazenda.gov.br/nfse nDPS" json:"nDPS"`
	DCompet       TSData           `xml:"http://www.sped.fazenda.gov.br/nfse dCompet" json:"dCompet"`
	TpEmit        TSEmitenteDPS    `xml:"http://www.sped.fazenda.gov.br/nfse tpEmit" json:"tpEmit"`
	CMotivoEmisTI *TSMotivoEmisTI  `xml:"http://www.sped.fazenda.gov.br/nfse cMotivoEmisTI,omitempty" json:"cMotivoEmisTI,omitempty"`
	ChNFSeRej     *TSChaveNFSe     `xml:"http://www.sped.fazenda.gov.br/nfse chNFSeRej,omitempty" json:"chNFSeRej,omitempty"`
	CLocEmi       TSCodMunIBGE     `xml:"http://www.sped.fazenda.gov.br/nfse cLocEmi" json:"cLocEmi"`
	Subst         *TCSubstituicao  `xml:"http://www.sped.fazenda.gov.br/nfse subst,omitempty" json:"subst,omitempty"`
	Prest         TCInfoPrestador  `xml:"http://www.sped.fazenda.gov.br/nfse prest" json:"prest"`
	Toma          *TCInfoPessoa    `xml:"http://www.sped.fazenda.gov.br/nfse toma,omitempty" json:"toma,omitempty"`
	Interm        *TCInfoPessoa    `xml:"http://www.sped.fazenda.gov.br/nfse interm,omitempty" json:"interm,omitempty"`
	Serv          TCServ           `xml:"http://www.sped.fazenda.gov.br/nfse serv" json:"serv"`
	Valores       TCInfoValores    `xml:"http://www.sped.fazenda.gov.br/nfse valores" json:"valores"`
	IBSCBS        *TCRTCInfoIBSCBS `xml:"http://www.sped.fazenda.gov.br/nfse IBSCBS,omitempty" json:"IBSCBS,omitempty"`
}

type TCInfNFSe struct {
	ID            TSIDNFSe         `xml:"Id,attr" json:"Id"`
	XLocEmi       TSDesc150        `xml:"http://www.sped.fazenda.gov.br/nfse xLocEmi" json:"xLocEmi"`
	XLocPrestacao TSDesc150        `xml:"http://www.sped.fazenda.gov.br/nfse xLocPrestacao" json:"xLocPrestacao"`
	NNFSe         TSNNFSe          `xml:"http://www.sped.fazenda.gov.br/nfse nNFSe" json:"nNFSe"`
	CLocIncid     *TSCodMunIBGE    `xml:"http://www.sped.fazenda.gov.br/nfse cLocIncid,omitempty" json:"cLocIncid,omitempty"`
	XLocIncid     *TSDesc150       `xml:"http://www.sped.fazenda.gov.br/nfse xLocIncid,omitempty" json:"xLocIncid,omitempty"`
	XTribNac      TSDesc600        `xml:"http://www.sped.fazenda.gov.br/nfse xTribNac" json:"xTribNac"`
	XTribMun      *TSDesc600       `xml:"http://www.sped.fazenda.gov.br/nfse xTribMun,omitempty" json:"xTribMun,omitempty"`
	XNBS          *TSDesc600       `xml:"http://www.sped.fazenda.gov.br/nfse xNBS,omitempty" json:"xNBS,omitempty"`
	VerAplic      TSVerAplic       `xml:"http://www.sped.fazenda.gov.br/nfse verAplic" json:"verAplic"`
	AmbGer        TSAmbGeradorNFSe `xml:"http://www.sped.fazenda.gov.br/nfse ambGer" json:"ambGer"`
	TpEmis        TSTipoEmissao    `xml:"http://www.sped.fazenda.gov.br/nfse tpEmis" json:"tpEmis"`
	ProcEmi       *TSProcEmissao   `xml:"http://www.sped.fazenda.gov.br/nfse procEmi,omitempty" json:"procEmi,omitempty"`
	CStat         TStat            `xml:"http://www.sped.fazenda.gov.br/nfse cStat" json:"cStat"`
	DhProc        TSDateTimeUTC    `xml:"http://www.sped.fazenda.gov.br/nfse dhProc" json:"dhProc"`
	NDFSe         TSNDFSe          `xml:"http://www.sped.fazenda.gov.br/nfse nDFSe" json:"nDFSe"`
	Emit          TCEmitente       `xml:"http://www.sped.fazenda.gov.br/nfse emit" json:"emit"`
	Valores       TCValoresNFSe    `xml:"http://www.sped.fazenda.gov.br/nfse valores" json:"valores"`
	XOutInf       *TSDesc2000      `xml:"http://www.sped.fazenda.gov.br/nfse xOutInf,omitempty" json:"xOutInf,omitempty"`
	IBSCBS        *TCRTCIBSCBS     `xml:"http://www.sped.fazenda.gov.br/nfse IBSCBS,omitempty" json:"IBSCBS,omitempty"`
	DPS           TCDPS            `xml:"http://www.sped.fazenda.gov.br/nfse DPS" json:"DPS"`
}

type TCInfoCompl struct {
	IDDocTec *TSDRT            `xml:"http://www.sped.fazenda.gov.br/nfse idDocTec,omitempty" json:"idDocTec,omitempty"`
	DocRef   *TSDesc255        `xml:"http://www.sped.fazenda.gov.br/nfse docRef,omitempty" json:"docRef,omitempty"`
	XPed     *TSNumeroEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xPed,omitempty" json:"xPed,omitempty"`
	GItemPed *TCInfoItemPed    `xml:"http://www.sped.fazenda.gov.br/nfse gItemPed,omitempty" json:"gItemPed,omitempty"`
	XInfComp *TSDescInfCompl   `xml:"http://www.sped.fazenda.gov.br/nfse xInfComp,omitempty" json:"xInfComp,omitempty"`
}

type TCInfoContribuinteCNC struct {
	CNPJ                               *TSCNPJ                               `xml:"http://www.sped.fazenda.gov.br/nfse CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF                                *TSCPF                                `xml:"http://www.sped.fazenda.gov.br/nfse CPF,omitempty" json:"CPF,omitempty"`
	IM                                 *TSInscMun                            `xml:"http://www.sped.fazenda.gov.br/nfse IM,omitempty" json:"IM,omitempty"`
	DInscricaoMunicipal                TSData                                `xml:"http://www.sped.fazenda.gov.br/nfse dInscricaoMunicipal" json:"dInscricaoMunicipal"`
	EnderContribuinteCNC               TCEnderContribuinteCNC                `xml:"http://www.sped.fazenda.gov.br/nfse enderContribuinteCNC" json:"enderContribuinteCNC"`
	Fone                               *TSTelefone                           `xml:"http://www.sped.fazenda.gov.br/nfse fone,omitempty" json:"fone,omitempty"`
	Email                              *TSEmail                              `xml:"http://www.sped.fazenda.gov.br/nfse email,omitempty" json:"email,omitempty"`
	RegEspTribContribuinteCNC          TSRegEspTrib                          `xml:"http://www.sped.fazenda.gov.br/nfse regEspTribContribuinteCNC" json:"regEspTribContribuinteCNC"`
	SituacaoCadastroContribuinte       *TSSituacaoCadastroContribuinte       `xml:"http://www.sped.fazenda.gov.br/nfse situacaoCadastroContribuinte,omitempty" json:"situacaoCadastroContribuinte,omitempty"`
	MotivoSituacaoCadastroContribuinte *TSMotivoSituacaoCadastroContribuinte `xml:"http://www.sped.fazenda.gov.br/nfse motivoSituacaoCadastroContribuinte,omitempty" json:"motivoSituacaoCadastroContribuinte,omitempty"`
	SituacaoEmissaoNFSe                TSSituacaoEmissaoNFSe                 `xml:"http://www.sped.fazenda.gov.br/nfse situacaoEmissaoNFSE" json:"situacaoEmissaoNFSE"`
}

type TCInfoDedRed struct {
	PDR        *TSDec3V2         `xml:"http://www.sped.fazenda.gov.br/nfse pDR,omitempty" json:"pDR,omitempty"`
	VDR        *TSDec15V2        `xml:"http://www.sped.fazenda.gov.br/nfse vDR,omitempty" json:"vDR,omitempty"`
	Documentos *TCListaDocDedRed `xml:"http://www.sped.fazenda.gov.br/nfse documentos,omitempty" json:"documentos,omitempty"`
}

type TCInfoItemPed struct {
	XItemPed []TSNumeroEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xItemPed,omitempty" json:"xItemPed,omitempty"`
}

type TCInfoObra struct {
	InscImobFisc *TSInscImobFisc    `xml:"http://www.sped.fazenda.gov.br/nfse inscImobFisc,omitempty" json:"inscImobFisc,omitempty"`
	CObra        *TSCodObra         `xml:"http://www.sped.fazenda.gov.br/nfse cObra,omitempty" json:"cObra,omitempty"`
	CCIB         *TSCodCIB          `xml:"http://www.sped.fazenda.gov.br/nfse cCIB,omitempty" json:"cCIB,omitempty"`
	End          *TCEnderObraEvento `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
}

type TCInfoPessoa struct {
	CNPJ    *TSCNPJ           `xml:"http://www.sped.fazenda.gov.br/nfse CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TSCPF            `xml:"http://www.sped.fazenda.gov.br/nfse CPF,omitempty" json:"CPF,omitempty"`
	NIF     *TSNIF            `xml:"http://www.sped.fazenda.gov.br/nfse NIF,omitempty" json:"NIF,omitempty"`
	CNaoNIF *TSCodNaoNIF      `xml:"http://www.sped.fazenda.gov.br/nfse cNaoNIF,omitempty" json:"cNaoNIF,omitempty"`
	CAEPF   *TSCAEPF          `xml:"http://www.sped.fazenda.gov.br/nfse CAEPF,omitempty" json:"CAEPF,omitempty"`
	IM      *TSInscMun        `xml:"http://www.sped.fazenda.gov.br/nfse IM,omitempty" json:"IM,omitempty"`
	XNome   TSNomeRazaoSocial `xml:"http://www.sped.fazenda.gov.br/nfse xNome" json:"xNome"`
	End     *TCEndereco       `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
	Fone    *TSTelefone       `xml:"http://www.sped.fazenda.gov.br/nfse fone,omitempty" json:"fone,omitempty"`
	Email   *TSEmail          `xml:"http://www.sped.fazenda.gov.br/nfse email,omitempty" json:"email,omitempty"`
}

type TCInfoPrestador struct {
	CNPJ    *TSCNPJ            `xml:"http://www.sped.fazenda.gov.br/nfse CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TSCPF             `xml:"http://www.sped.fazenda.gov.br/nfse CPF,omitempty" json:"CPF,omitempty"`
	NIF     *TSNIF             `xml:"http://www.sped.fazenda.gov.br/nfse NIF,omitempty" json:"NIF,omitempty"`
	CNaoNIF *TSCodNaoNIF       `xml:"http://www.sped.fazenda.gov.br/nfse cNaoNIF,omitempty" json:"cNaoNIF,omitempty"`
	CAEPF   *TSCAEPF           `xml:"http://www.sped.fazenda.gov.br/nfse CAEPF,omitempty" json:"CAEPF,omitempty"`
	IM      *TSInscMun         `xml:"http://www.sped.fazenda.gov.br/nfse IM,omitempty" json:"IM,omitempty"`
	XNome   *TSNomeRazaoSocial `xml:"http://www.sped.fazenda.gov.br/nfse xNome,omitempty" json:"xNome,omitempty"`
	End     *TCEndereco        `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
	Fone    *TSTelefone        `xml:"http://www.sped.fazenda.gov.br/nfse fone,omitempty" json:"fone,omitempty"`
	Email   *TSEmail           `xml:"http://www.sped.fazenda.gov.br/nfse email,omitempty" json:"email,omitempty"`
	RegTrib TCRegTrib          `xml:"http://www.sped.fazenda.gov.br/nfse regTrib" json:"regTrib"`
}

type TCInfoRefNFSe struct {
	RefNFSe []TSChaveNFSe `xml:"http://www.sped.fazenda.gov.br/nfse refNFSe,omitempty" json:"refNFSe,omitempty"`
}

type TCInfoTributacao struct {
	TribMun TCTribMunicipal `xml:"http://www.sped.fazenda.gov.br/nfse tribMun" json:"tribMun"`
	TribFed *TCTribFederal  `xml:"http://www.sped.fazenda.gov.br/nfse tribFed,omitempty" json:"tribFed,omitempty"`
	TotTrib TCTribTotal     `xml:"http://www.sped.fazenda.gov.br/nfse totTrib" json:"totTrib"`
}

type TCInfoValores struct {
	VServPrest      TCVServPrest       `xml:"http://www.sped.fazenda.gov.br/nfse vServPrest" json:"vServPrest"`
	VDescCondIncond *TCVDescCondIncond `xml:"http://www.sped.fazenda.gov.br/nfse vDescCondIncond,omitempty" json:"vDescCondIncond,omitempty"`
	VDedRed         *TCInfoDedRed      `xml:"http://www.sped.fazenda.gov.br/nfse vDedRed,omitempty" json:"vDedRed,omitempty"`
	Trib            TCInfoTributacao   `xml:"http://www.sped.fazenda.gov.br/nfse trib" json:"trib"`
}

type TCListaDocDedRed struct {
	DocDedRed []TCDocDedRed `xml:"http://www.sped.fazenda.gov.br/nfse docDedRed,omitempty" json:"docDedRed,omitempty"`
}

type TCLocPrest struct {
	CLocPrestacao  *TSCodMunIBGE `xml:"http://www.sped.fazenda.gov.br/nfse cLocPrestacao,omitempty" json:"cLocPrestacao,omitempty"`
	CPaisPrestacao *TSCodPaisISO `xml:"http://www.sped.fazenda.gov.br/nfse cPaisPrestacao,omitempty" json:"cPaisPrestacao,omitempty"`
}

type TCNC struct {
	Versao TVerCNC  `xml:"versao,attr" json:"versao"`
	InfCNC TCInfCNC `xml:"http://www.sped.fazenda.gov.br/nfse infCNC" json:"infCNC"`
}

type TCNFSe struct {
	Versao    TVerNFSe      `xml:"versao,attr" json:"versao"`
	InfNFSe   TCInfNFSe     `xml:"http://www.sped.fazenda.gov.br/nfse infNFSe" json:"infNFSe"`
	Signature SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCObjetoLocacao string

type TCRTCIBSCBS struct {
	CLocalidadeIncid TSCodMunIBGE       `xml:"http://www.sped.fazenda.gov.br/nfse cLocalidadeIncid" json:"cLocalidadeIncid"`
	XLocalidadeIncid TSDesc600          `xml:"http://www.sped.fazenda.gov.br/nfse xLocalidadeIncid" json:"xLocalidadeIncid"`
	PRedutor         *TSDec2V2          `xml:"http://www.sped.fazenda.gov.br/nfse pRedutor,omitempty" json:"pRedutor,omitempty"`
	Valores          TCRTCValoresIBSCBS `xml:"http://www.sped.fazenda.gov.br/nfse valores" json:"valores"`
	TotCIBS          TCRTCTotalCIBS     `xml:"http://www.sped.fazenda.gov.br/nfse totCIBS" json:"totCIBS"`
}

type TCRTCInfoDest struct {
	CNPJ    *TSCNPJ      `xml:"http://www.sped.fazenda.gov.br/nfse CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TSCPF       `xml:"http://www.sped.fazenda.gov.br/nfse CPF,omitempty" json:"CPF,omitempty"`
	NIF     *TSNIF       `xml:"http://www.sped.fazenda.gov.br/nfse NIF,omitempty" json:"NIF,omitempty"`
	CNaoNIF *TSCodNaoNIF `xml:"http://www.sped.fazenda.gov.br/nfse cNaoNIF,omitempty" json:"cNaoNIF,omitempty"`
	XNome   TSDesc150    `xml:"http://www.sped.fazenda.gov.br/nfse xNome" json:"xNome"`
	End     *TCEndereco  `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
	Fone    *TSTelefone  `xml:"http://www.sped.fazenda.gov.br/nfse fone,omitempty" json:"fone,omitempty"`
	Email   *TSEmail     `xml:"http://www.sped.fazenda.gov.br/nfse email,omitempty" json:"email,omitempty"`
}

type TCRTCInfoIBSCBS struct {
	FinNFSe   TSRTCFinNFSe           `xml:"http://www.sped.fazenda.gov.br/nfse finNFSe" json:"finNFSe"`
	IndFinal  *TSRTCIndFinal         `xml:"http://www.sped.fazenda.gov.br/nfse indFinal,omitempty" json:"indFinal,omitempty"`
	CIndOp    TSRTCCodIndOp          `xml:"http://www.sped.fazenda.gov.br/nfse cIndOp" json:"cIndOp"`
	TpOper    *TSRTCTpOper           `xml:"http://www.sped.fazenda.gov.br/nfse tpOper,omitempty" json:"tpOper,omitempty"`
	GRefNFSe  *TCInfoRefNFSe         `xml:"http://www.sped.fazenda.gov.br/nfse gRefNFSe,omitempty" json:"gRefNFSe,omitempty"`
	TpEnteGov *TSRTCTpEnteGov        `xml:"http://www.sped.fazenda.gov.br/nfse tpEnteGov,omitempty" json:"tpEnteGov,omitempty"`
	IndDest   TSRTCIndDest           `xml:"http://www.sped.fazenda.gov.br/nfse indDest" json:"indDest"`
	Dest      *TCRTCInfoDest         `xml:"http://www.sped.fazenda.gov.br/nfse dest,omitempty" json:"dest,omitempty"`
	Imovel    *TCRTCInfoImovel       `xml:"http://www.sped.fazenda.gov.br/nfse imovel,omitempty" json:"imovel,omitempty"`
	Valores   TCRTCInfoValoresIBSCBS `xml:"http://www.sped.fazenda.gov.br/nfse valores" json:"valores"`
}

type TCRTCInfoImovel struct {
	InscImobFisc *TSInscImobFisc    `xml:"http://www.sped.fazenda.gov.br/nfse inscImobFisc,omitempty" json:"inscImobFisc,omitempty"`
	CCIB         *TSCodCIB          `xml:"http://www.sped.fazenda.gov.br/nfse cCIB,omitempty" json:"cCIB,omitempty"`
	End          *TCEnderObraEvento `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
}

type TCRTCInfoReeRepRes struct {
	Documentos []TCRTCListaDoc `xml:"http://www.sped.fazenda.gov.br/nfse documentos,omitempty" json:"documentos,omitempty"`
}

type TCRTCInfoTributosDif struct {
	PDifUF  TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pDifUF" json:"pDifUF"`
	PDifMun TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pDifMun" json:"pDifMun"`
	PDifCBS TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pDifCBS" json:"pDifCBS"`
}

type TCRTCInfoTributosIBSCBS struct {
	GIBSCBS TCRTCInfoTributosSitClas `xml:"http://www.sped.fazenda.gov.br/nfse gIBSCBS" json:"gIBSCBS"`
}

type TCRTCInfoTributosSitClas struct {
	CST          TSRTCCodSitTrib               `xml:"http://www.sped.fazenda.gov.br/nfse CST" json:"CST"`
	CClassTrib   TSRTCCodClassTrib             `xml:"http://www.sped.fazenda.gov.br/nfse cClassTrib" json:"cClassTrib"`
	CCredPres    *TSRTCCodCredPres             `xml:"http://www.sped.fazenda.gov.br/nfse cCredPres,omitempty" json:"cCredPres,omitempty"`
	GTribRegular *TCRTCInfoTributosTribRegular `xml:"http://www.sped.fazenda.gov.br/nfse gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GDif         *TCRTCInfoTributosDif         `xml:"http://www.sped.fazenda.gov.br/nfse gDif,omitempty" json:"gDif,omitempty"`
}

type TCRTCInfoTributosTribRegular struct {
	CSTReg        TSRTCCodSitTrib   `xml:"http://www.sped.fazenda.gov.br/nfse CSTReg" json:"CSTReg"`
	CClassTribReg TSRTCCodClassTrib `xml:"http://www.sped.fazenda.gov.br/nfse cClassTribReg" json:"cClassTribReg"`
}

type TCRTCInfoValoresIBSCBS struct {
	GReeRepRes *TCRTCInfoReeRepRes     `xml:"http://www.sped.fazenda.gov.br/nfse gReeRepRes,omitempty" json:"gReeRepRes,omitempty"`
	Trib       TCRTCInfoTributosIBSCBS `xml:"http://www.sped.fazenda.gov.br/nfse trib" json:"trib"`
}

type TCRTCListaDoc struct {
	DFeNacional    *TCRTCListaDocDFe         `xml:"http://www.sped.fazenda.gov.br/nfse dFeNacional,omitempty" json:"dFeNacional,omitempty"`
	DocFiscalOutro *TCRTCListaDocFiscalOutro `xml:"http://www.sped.fazenda.gov.br/nfse docFiscalOutro,omitempty" json:"docFiscalOutro,omitempty"`
	DocOutro       *TCRTCListaDocOutro       `xml:"http://www.sped.fazenda.gov.br/nfse docOutro,omitempty" json:"docOutro,omitempty"`
	Fornec         *TCRTCListaDocFornec      `xml:"http://www.sped.fazenda.gov.br/nfse fornec,omitempty" json:"fornec,omitempty"`
	DtEmiDoc       TSData                    `xml:"http://www.sped.fazenda.gov.br/nfse dtEmiDoc" json:"dtEmiDoc"`
	DtCompDoc      TSData                    `xml:"http://www.sped.fazenda.gov.br/nfse dtCompDoc" json:"dtCompDoc"`
	TpReeRepRes    TSRTCTpReeRepRes          `xml:"http://www.sped.fazenda.gov.br/nfse tpReeRepRes" json:"tpReeRepRes"`
	XTpReeRepRes   *TSDesc150                `xml:"http://www.sped.fazenda.gov.br/nfse xTpReeRepRes,omitempty" json:"xTpReeRepRes,omitempty"`
	VlrReeRepRes   TSDec15V2                 `xml:"http://www.sped.fazenda.gov.br/nfse vlrReeRepRes" json:"vlrReeRepRes"`
}

type TCRTCListaDocDFe struct {
	TipoChaveDFe  TSRTCTipoChaveDFe `xml:"http://www.sped.fazenda.gov.br/nfse tipoChaveDFe" json:"tipoChaveDFe"`
	XTipoChaveDFe *TSDesc255        `xml:"http://www.sped.fazenda.gov.br/nfse xTipoChaveDFe,omitempty" json:"xTipoChaveDFe,omitempty"`
	ChaveDFe      TSRTCChaveDFe     `xml:"http://www.sped.fazenda.gov.br/nfse chaveDFe" json:"chaveDFe"`
}

type TCRTCListaDocFiscalOutro struct {
	CMunDocFiscal TSNum7Dig `xml:"http://www.sped.fazenda.gov.br/nfse cMunDocFiscal" json:"cMunDocFiscal"`
	NDocFiscal    TSDesc255 `xml:"http://www.sped.fazenda.gov.br/nfse nDocFiscal" json:"nDocFiscal"`
	XDocFiscal    TSDesc255 `xml:"http://www.sped.fazenda.gov.br/nfse xDocFiscal" json:"xDocFiscal"`
}

type TCRTCListaDocFornec struct {
	CNPJ    *TSCNPJ      `xml:"http://www.sped.fazenda.gov.br/nfse CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TSCPF       `xml:"http://www.sped.fazenda.gov.br/nfse CPF,omitempty" json:"CPF,omitempty"`
	NIF     *TSNIF       `xml:"http://www.sped.fazenda.gov.br/nfse NIF,omitempty" json:"NIF,omitempty"`
	CNaoNIF *TSCodNaoNIF `xml:"http://www.sped.fazenda.gov.br/nfse cNaoNIF,omitempty" json:"cNaoNIF,omitempty"`
	XNome   TSDesc150    `xml:"http://www.sped.fazenda.gov.br/nfse xNome" json:"xNome"`
}

type TCRTCListaDocOutro struct {
	NDoc TSDesc255 `xml:"http://www.sped.fazenda.gov.br/nfse nDoc" json:"nDoc"`
	XDoc TSDesc255 `xml:"http://www.sped.fazenda.gov.br/nfse xDoc" json:"xDoc"`
}

type TCRTCTotalCBS struct {
	GCBSCredPres *TCRTCTotalCBSCredPres `xml:"http://www.sped.fazenda.gov.br/nfse gCBSCredPres,omitempty" json:"gCBSCredPres,omitempty"`
	VDifCBS      *TSDec15V2             `xml:"http://www.sped.fazenda.gov.br/nfse vDifCBS,omitempty" json:"vDifCBS,omitempty"`
	VCBS         TSDec15V2              `xml:"http://www.sped.fazenda.gov.br/nfse vCBS" json:"vCBS"`
}

type TCRTCTotalCBSCredPres struct {
	PCredPresCBS TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pCredPresCBS" json:"pCredPresCBS"`
	VCredPresCBS TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vCredPresCBS" json:"vCredPresCBS"`
}

type TCRTCTotalCIBS struct {
	VTotNF         TSDec15V2                `xml:"http://www.sped.fazenda.gov.br/nfse vTotNF" json:"vTotNF"`
	GIBS           TCRTCTotalIBS            `xml:"http://www.sped.fazenda.gov.br/nfse gIBS" json:"gIBS"`
	GCBS           TCRTCTotalCBS            `xml:"http://www.sped.fazenda.gov.br/nfse gCBS" json:"gCBS"`
	GTribRegular   *TCRTCTotalTribRegular   `xml:"http://www.sped.fazenda.gov.br/nfse gTribRegular,omitempty" json:"gTribRegular,omitempty"`
	GTribCompraGov *TCRTCTotalTribCompraGov `xml:"http://www.sped.fazenda.gov.br/nfse gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TCRTCTotalIBS struct {
	VIBSTot      TSDec15V2              `xml:"http://www.sped.fazenda.gov.br/nfse vIBSTot" json:"vIBSTot"`
	GIBSCredPres *TCRTCTotalIBSCredPres `xml:"http://www.sped.fazenda.gov.br/nfse gIBSCredPres,omitempty" json:"gIBSCredPres,omitempty"`
	GIBSUFTot    TCRTCTotalIBSUF        `xml:"http://www.sped.fazenda.gov.br/nfse gIBSUFTot" json:"gIBSUFTot"`
	GIBSMunTot   TCRTCTotalIBSMun       `xml:"http://www.sped.fazenda.gov.br/nfse gIBSMunTot" json:"gIBSMunTot"`
}

type TCRTCTotalIBSCredPres struct {
	PCredPresIBS TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pCredPresIBS" json:"pCredPresIBS"`
	VCredPresIBS TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vCredPresIBS" json:"vCredPresIBS"`
}

type TCRTCTotalIBSMun struct {
	VDifMun *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vDifMun,omitempty" json:"vDifMun,omitempty"`
	VIBSMun TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vIBSMun" json:"vIBSMun"`
}

type TCRTCTotalIBSUF struct {
	VDifUF *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vDifUF,omitempty" json:"vDifUF,omitempty"`
	VIBSUF TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vIBSUF" json:"vIBSUF"`
}

type TCRTCTotalTribCompraGov struct {
	PIBSUF  TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pIBSUF" json:"pIBSUF"`
	VIBSUF  TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vIBSUF" json:"vIBSUF"`
	PIBSMun TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pIBSMun" json:"pIBSMun"`
	VIBSMun TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vIBSMun" json:"vIBSMun"`
	PCBS    TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pCBS" json:"pCBS"`
	VCBS    TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vCBS" json:"vCBS"`
}

type TCRTCTotalTribRegular struct {
	PAliqEfeRegIBSUF  TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pAliqEfeRegIBSUF" json:"pAliqEfeRegIBSUF"`
	VTribRegIBSUF     TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vTribRegIBSUF" json:"vTribRegIBSUF"`
	PAliqEfeRegIBSMun TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pAliqEfeRegIBSMun" json:"pAliqEfeRegIBSMun"`
	VTribRegIBSMun    TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vTribRegIBSMun" json:"vTribRegIBSMun"`
	PAliqEfeRegCBS    TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pAliqEfeRegCBS" json:"pAliqEfeRegCBS"`
	VTribRegCBS       TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vTribRegCBS" json:"vTribRegCBS"`
}

type TCRTCValoresIBSCBS struct {
	VBC            TSDec15V2             `xml:"http://www.sped.fazenda.gov.br/nfse vBC" json:"vBC"`
	VCalcReeRepRes *TSDec15V2            `xml:"http://www.sped.fazenda.gov.br/nfse vCalcReeRepRes,omitempty" json:"vCalcReeRepRes,omitempty"`
	UF             TCRTCValoresIBSCBSUF  `xml:"http://www.sped.fazenda.gov.br/nfse uf" json:"uf"`
	Mun            TCRTCValoresIBSCBSMun `xml:"http://www.sped.fazenda.gov.br/nfse mun" json:"mun"`
	Fed            TCRTCValoresIBSCBSFed `xml:"http://www.sped.fazenda.gov.br/nfse fed" json:"fed"`
}

type TCRTCValoresIBSCBSFed struct {
	PCBS         TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pCBS" json:"pCBS"`
	PRedAliqCBS  *TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pRedAliqCBS,omitempty" json:"pRedAliqCBS,omitempty"`
	PAliqEfetCBS TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pAliqEfetCBS" json:"pAliqEfetCBS"`
}

type TCRTCValoresIBSCBSMun struct {
	PIBSMun      TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pIBSMun" json:"pIBSMun"`
	PRedAliqMun  *TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pRedAliqMun,omitempty" json:"pRedAliqMun,omitempty"`
	PAliqEfetMun TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pAliqEfetMun" json:"pAliqEfetMun"`
}

type TCRTCValoresIBSCBSUF struct {
	PIBSUF      TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pIBSUF" json:"pIBSUF"`
	PRedAliqUF  *TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pRedAliqUF,omitempty" json:"pRedAliqUF,omitempty"`
	PAliqEfetUF TSDec2V2  `xml:"http://www.sped.fazenda.gov.br/nfse pAliqEfetUF" json:"pAliqEfetUF"`
}

type TCRegTrib struct {
	OpSimpNac   TSOpSimpNac              `xml:"http://www.sped.fazenda.gov.br/nfse opSimpNac" json:"opSimpNac"`
	RegApTribSN *TSRegimeApuracaoSimpNac `xml:"http://www.sped.fazenda.gov.br/nfse regApTribSN,omitempty" json:"regApTribSN,omitempty"`
	RegEspTrib  TSRegEspTrib             `xml:"http://www.sped.fazenda.gov.br/nfse regEspTrib" json:"regEspTrib"`
}

type TCServ struct {
	LocPrest  TCLocPrest     `xml:"http://www.sped.fazenda.gov.br/nfse locPrest" json:"locPrest"`
	CServ     TCCServ        `xml:"http://www.sped.fazenda.gov.br/nfse cServ" json:"cServ"`
	ComExt    *TCComExterior `xml:"http://www.sped.fazenda.gov.br/nfse comExt,omitempty" json:"comExt,omitempty"`
	Obra      *TCInfoObra    `xml:"http://www.sped.fazenda.gov.br/nfse obra,omitempty" json:"obra,omitempty"`
	AtvEvento *TCAtvEvento   `xml:"http://www.sped.fazenda.gov.br/nfse atvEvento,omitempty" json:"atvEvento,omitempty"`
	InfoCompl *TCInfoCompl   `xml:"http://www.sped.fazenda.gov.br/nfse infoCompl,omitempty" json:"infoCompl,omitempty"`
}

type TCSubstituicao struct {
	ChSubstda TSChaveNFSe    `xml:"http://www.sped.fazenda.gov.br/nfse chSubstda" json:"chSubstda"`
	CMotivo   TSCodJustSubst `xml:"http://www.sped.fazenda.gov.br/nfse cMotivo" json:"cMotivo"`
	XMotivo   *TSMotivo      `xml:"http://www.sped.fazenda.gov.br/nfse xMotivo,omitempty" json:"xMotivo,omitempty"`
}

type TCTribFederal struct {
	Piscofins *TCTribOutrosPISCOFINS `xml:"http://www.sped.fazenda.gov.br/nfse piscofins,omitempty" json:"piscofins,omitempty"`
	VRetCP    *TSDec15V2             `xml:"http://www.sped.fazenda.gov.br/nfse vRetCP,omitempty" json:"vRetCP,omitempty"`
	VRetIRRF  *TSDec15V2             `xml:"http://www.sped.fazenda.gov.br/nfse vRetIRRF,omitempty" json:"vRetIRRF,omitempty"`
	VRetCSLL  *TSDec15V2             `xml:"http://www.sped.fazenda.gov.br/nfse vRetCSLL,omitempty" json:"vRetCSLL,omitempty"`
}

type TCTribMunicipal struct {
	TribISSQN   TSTribISSQN           `xml:"http://www.sped.fazenda.gov.br/nfse tribISSQN" json:"tribISSQN"`
	CPaisResult *TSCodPaisISO         `xml:"http://www.sped.fazenda.gov.br/nfse cPaisResult,omitempty" json:"cPaisResult,omitempty"`
	TpImunidade *TSTipoImunidadeISSQN `xml:"http://www.sped.fazenda.gov.br/nfse tpImunidade,omitempty" json:"tpImunidade,omitempty"`
	ExigSusp    *TCExigSuspensa       `xml:"http://www.sped.fazenda.gov.br/nfse exigSusp,omitempty" json:"exigSusp,omitempty"`
	BM          *TCBeneficioMunicipal `xml:"http://www.sped.fazenda.gov.br/nfse BM,omitempty" json:"BM,omitempty"`
	TpRetISSQN  TSTipoRetISSQN        `xml:"http://www.sped.fazenda.gov.br/nfse tpRetISSQN" json:"tpRetISSQN"`
	PAliq       *TSDec1V2             `xml:"http://www.sped.fazenda.gov.br/nfse pAliq,omitempty" json:"pAliq,omitempty"`
}

type TCTribOutrosPISCOFINS struct {
	CST            TSTipoCST           `xml:"http://www.sped.fazenda.gov.br/nfse CST" json:"CST"`
	VBCPISCOFINS   *TSDec15V2          `xml:"http://www.sped.fazenda.gov.br/nfse vBCPisCofins,omitempty" json:"vBCPisCofins,omitempty"`
	PAliqPIS       *TSDec2V2           `xml:"http://www.sped.fazenda.gov.br/nfse pAliqPis,omitempty" json:"pAliqPis,omitempty"`
	PAliqCOFINS    *TSDec2V2           `xml:"http://www.sped.fazenda.gov.br/nfse pAliqCofins,omitempty" json:"pAliqCofins,omitempty"`
	VPIS           *TSDec15V2          `xml:"http://www.sped.fazenda.gov.br/nfse vPis,omitempty" json:"vPis,omitempty"`
	VCOFINS        *TSDec15V2          `xml:"http://www.sped.fazenda.gov.br/nfse vCofins,omitempty" json:"vCofins,omitempty"`
	TpRetPISCOFINS *TSTipoRetPISCOFINS `xml:"http://www.sped.fazenda.gov.br/nfse tpRetPisCofins,omitempty" json:"tpRetPisCofins,omitempty"`
}

type TCTribTotal struct {
	VTotTrib   *TCTribTotalMonet   `xml:"http://www.sped.fazenda.gov.br/nfse vTotTrib,omitempty" json:"vTotTrib,omitempty"`
	PTotTrib   *TCTribTotalPercent `xml:"http://www.sped.fazenda.gov.br/nfse pTotTrib,omitempty" json:"pTotTrib,omitempty"`
	IndTotTrib *TSTipoIndTotTrib   `xml:"http://www.sped.fazenda.gov.br/nfse indTotTrib,omitempty" json:"indTotTrib,omitempty"`
	PTotTribSN *TSDec2V2           `xml:"http://www.sped.fazenda.gov.br/nfse pTotTribSN,omitempty" json:"pTotTribSN,omitempty"`
}

type TCTribTotalMonet struct {
	VTotTribFed TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vTotTribFed" json:"vTotTribFed"`
	VTotTribEst TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vTotTribEst" json:"vTotTribEst"`
	VTotTribMun TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vTotTribMun" json:"vTotTribMun"`
}

type TCTribTotalPercent struct {
	PTotTribFed TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pTotTribFed" json:"pTotTribFed"`
	PTotTribEst TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pTotTribEst" json:"pTotTribEst"`
	PTotTribMun TSDec3V2 `xml:"http://www.sped.fazenda.gov.br/nfse pTotTribMun" json:"pTotTribMun"`
}

type TCVDescCondIncond struct {
	VDescIncond *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vDescIncond,omitempty" json:"vDescIncond,omitempty"`
	VDescCond   *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vDescCond,omitempty" json:"vDescCond,omitempty"`
}

type TCVServPrest struct {
	VReceb *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vReceb,omitempty" json:"vReceb,omitempty"`
	VServ  TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vServ" json:"vServ"`
}

type TCValoresNFSe struct {
	VCalcDR    *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vCalcDR,omitempty" json:"vCalcDR,omitempty"`
	TpBM       *TBMISSQN  `xml:"http://www.sped.fazenda.gov.br/nfse tpBM,omitempty" json:"tpBM,omitempty"`
	VCalcBM    *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vCalcBM,omitempty" json:"vCalcBM,omitempty"`
	VBC        *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vBC,omitempty" json:"vBC,omitempty"`
	PAliqAplic *TSDec1V2  `xml:"http://www.sped.fazenda.gov.br/nfse pAliqAplic,omitempty" json:"pAliqAplic,omitempty"`
	VISSQN     *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vISSQN,omitempty" json:"vISSQN,omitempty"`
	VTotalRet  *TSDec15V2 `xml:"http://www.sped.fazenda.gov.br/nfse vTotalRet,omitempty" json:"vTotalRet,omitempty"`
	VLiq       TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vLiq" json:"vLiq"`
}

type TSAmbGeradorEvt string

type TSAmbGeradorNFSe string

type TSBairro string

type TSCAEPF string

type TSCEP string

type TSCNPJ string

type TSCPF string

type TSCategVeic string

type TSCategoriaServico string

type TSChaveNFSe string

type TSChaveNFe string

type TSCidade string

type TSCodAcessoPed string

type TSCodAutorManifestacao string

type TSCodCIB string

type TSCodContrato string

type TSCodJustAnaliseFiscalCanc string

type TSCodJustAnaliseFiscalCancDef string

type TSCodJustAnaliseFiscalCancIndef string

type TSCodJustCanc string

type TSCodJustSubst string

type TSCodMoeda string

type TSCodMotivoRejeicao string

type TSCodMunIBGE string

type TSCodNBS string

type TSCodNaoNIF string

type TSCodObra string

type TSCodPaisISO string

type TSCodTribNac string

type TSCodVerificacao string

type TSCodigoEndPostal string

type TSCodigoEventoNFSe string

type TSCodigoInternoContribuinte string

type TSComplementoEndereco string

type TSDRT string

type TSData string

type TSDateTimeUTC string

type TSDec15V2 string

type TSDec1V2 string

type TSDec2V2 string

type TSDec3V2 string

type TSDesc1000 string

type TSDesc150 string

type TSDesc2000 string

type TSDesc255 string

type TSDesc40 string

type TSDesc500 string

type TSDesc600 string

type TSDescInfCompl string

type TSDescOutDedRed string

type TSEmail string

type TSEmitenteDPS string

type TSEnderCompletoExt string

type TSEnvMDIC string

type TSEstadoProvRegiao string

type TSExtensaoTotal string

type TSIDDPS string

type TSIDEvento string

type TSIDNFSe string

type TSIDNumEvento string

type TSIDPedRegEvt string

type TSIdeDedRed string

type TSIdeEvento string

type TSInscImobFisc string

type TSInscMun string

type TSLogradouro string

type TSMecAFComExPrest string

type TSMecAFComExToma string

type TSModoPrestacao string

type TSMotivo string

type TSMotivoEmisTI string

type TSMotivoSituacaoCadastroContribuinte string

type TSMovTempBens string

type TSNDFSe string

type TSNIF string

type TSNNFSe string

type TSNomeFantasia string

type TSNomeRazaoSocial string

type TSNum14Dig string

type TSNum15Dig string

type TSNum3Dig string

type TSNum7Dig string

type TSNumBeneficioMunicipal string

type TSNumDFe string

type TSNumDPS string

type TSNumDocImport string

type TSNumEixos string

type TSNumProcAdmAnaliseFiscalCanc string

type TSNumProcExigSuspensa string

type TSNumProcesso string

type TSNumRegExport string

type TSNumeroEndereco string

type TSNumeroPostes string

type TSOpConsumServ string

type TSOpExigSuspensa string

type TSOpSNLimUltrap string

type TSOpSimpNac string

type TSPlaca string

type TSProcEmissao string

type TSRTCChaveDFe string

type TSRTCCodClassTrib string

type TSRTCCodCredPres string

type TSRTCCodIndOp string

type TSRTCCodSitTrib string

type TSRTCFinNFSe string

type TSRTCIndDest string

type TSRTCIndFinal string

type TSRTCTipoChaveDFe string

type TSRTCTpEnteGov string

type TSRTCTpOper string

type TSRTCTpReeRepRes string

type TSRegEspTrib string

type TSRegimeApuracaoSimpNac string

type TSRodagem string

type TSSentido string

type TSSerieDPS string

type TSSerieNFNFS string

type TSSituacaoCadastroContribuinte string

type TSSituacaoEmissaoNFSe string

type TSString string

type TSStringComQuebraDeLinha string

type TSTelefone string

type TSTipoAmbiente string

type TSTipoCST string

type TSTipoEmissao string

type TSTipoImunidadeISSQN string

type TSTipoIndTotTrib string

type TSTipoRetISSQN string

type TSTipoRetPISCOFINS string

type TSTribISSQN string

type TSUF string

type TSVerAplic string

type TSVincPrest string

type TStat string

type TVerCNC string

type TVerNFSe string

type CanonicalizationMethodType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type CryptoBinary string

type DSAKeyValueType struct {
	P           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# P,omitempty" json:"P,omitempty"`
	Q           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Q,omitempty" json:"Q,omitempty"`
	G           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# G,omitempty" json:"G,omitempty"`
	Y           CryptoBinary  `xml:"http://www.w3.org/2000/09/xmldsig# Y" json:"Y"`
	J           *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# J,omitempty" json:"J,omitempty"`
	Seed        *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Seed,omitempty" json:"Seed,omitempty"`
	PgenCounter *CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# PgenCounter,omitempty" json:"PgenCounter,omitempty"`
}

type DigestMethodType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type DigestValueType string

type HMACOutputLengthType string

type KeyInfoType struct {
	ID *ID `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type KeyValueType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type ManifestType struct {
	ID        *ID             `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	Reference []ReferenceType `xml:"http://www.w3.org/2000/09/xmldsig# Reference,omitempty" json:"Reference,omitempty"`
}

type ObjectType struct {
	ID       *ID     `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	MimeType *String `xml:"MimeType,attr,omitempty" json:"MimeType,omitempty"`
	Encoding *AnyURI `xml:"Encoding,attr,omitempty" json:"Encoding,omitempty"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type PGPDataType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type RSAKeyValueType struct {
	Modulus  CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Modulus" json:"Modulus"`
	Exponent CryptoBinary `xml:"http://www.w3.org/2000/09/xmldsig# Exponent" json:"Exponent"`
}

type ReferenceType struct {
	ID           *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	URI          *AnyURI          `xml:"URI,attr,omitempty" json:"URI,omitempty"`
	TypeField    *AnyURI          `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	Transforms   *TransformsType  `xml:"http://www.w3.org/2000/09/xmldsig# Transforms,omitempty" json:"Transforms,omitempty"`
	DigestMethod DigestMethodType `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod" json:"DigestMethod"`
	DigestValue  DigestValueType  `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue" json:"DigestValue"`
}

type RetrievalMethodType struct {
	URI        *AnyURI         `xml:"URI,attr,omitempty" json:"URI,omitempty"`
	TypeField  *AnyURI         `xml:"Type,attr,omitempty" json:"Type,omitempty"`
	Transforms *TransformsType `xml:"http://www.w3.org/2000/09/xmldsig# Transforms,omitempty" json:"Transforms,omitempty"`
}

type SPKIDataType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type SignatureMethodType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type SignaturePropertiesType struct {
	ID                *ID                     `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	SignatureProperty []SignaturePropertyType `xml:"http://www.w3.org/2000/09/xmldsig# SignatureProperty,omitempty" json:"SignatureProperty,omitempty"`
}

type SignaturePropertyType struct {
	Target AnyURI `xml:"Target,attr" json:"Target"`
	ID     *ID    `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type SignatureType struct {
	ID             *ID                `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	SignedInfo     SignedInfoType     `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo" json:"SignedInfo"`
	SignatureValue SignatureValueType `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue" json:"SignatureValue"`
	KeyInfo        *KeyInfoType       `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo,omitempty" json:"KeyInfo,omitempty"`
	Object         []ObjectType       `xml:"http://www.w3.org/2000/09/xmldsig# Object,omitempty" json:"Object,omitempty"`
}

type SignatureValueType struct {
	Value Base64Binary `xml:",chardata" json:"value"`
	ID    *ID          `xml:"Id,attr,omitempty" json:"Id,omitempty"`
}

type SignedInfoType struct {
	ID                     *ID                        `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	CanonicalizationMethod CanonicalizationMethodType `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod" json:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethodType        `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod" json:"SignatureMethod"`
	Reference              []ReferenceType            `xml:"http://www.w3.org/2000/09/xmldsig# Reference,omitempty" json:"Reference,omitempty"`
}

type TransformType struct {
	Algorithm AnyURI `xml:"Algorithm,attr" json:"Algorithm"`
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type TransformsType struct {
	Transform []TransformType `xml:"http://www.w3.org/2000/09/xmldsig# Transform,omitempty" json:"Transform,omitempty"`
}

type X509DataType struct {
	// InnerXML preserves mixed or interleaved repeated content in schema order.
	InnerXML string `xml:",innerxml" json:"innerXML"`
}

type X509IssuerSerialType struct {
	X509IssuerName   String  `xml:"http://www.w3.org/2000/09/xmldsig# X509IssuerName" json:"X509IssuerName"`
	X509SerialNumber Integer `xml:"http://www.w3.org/2000/09/xmldsig# X509SerialNumber" json:"X509SerialNumber"`
}

type ID string

type AnyURI string

type Base64Binary string

type Date string

type Integer string

type String string

// CNC represents the global element {http://www.sped.fazenda.gov.br/nfse}CNC.
type CNC struct {
	XMLName xml.Name `xml:"http://www.sped.fazenda.gov.br/nfse CNC" json:"-"`
	TCNC
}

// CanonicalizationMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}CanonicalizationMethod.
type CanonicalizationMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# CanonicalizationMethod" json:"-"`
	CanonicalizationMethodType
}

// DSAKeyValue represents the global element {http://www.w3.org/2000/09/xmldsig#}DSAKeyValue.
type DSAKeyValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DSAKeyValue" json:"-"`
	DSAKeyValueType
}

// DigestMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}DigestMethod.
type DigestMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# DigestMethod" json:"-"`
	DigestMethodType
}

// DigestValue represents the global element {http://www.w3.org/2000/09/xmldsig#}DigestValue.
type DigestValue struct {
	XMLName xml.Name        `xml:"http://www.w3.org/2000/09/xmldsig# DigestValue" json:"-"`
	Value   DigestValueType `xml:",chardata" json:"value"`
}

// KeyInfo represents the global element {http://www.w3.org/2000/09/xmldsig#}KeyInfo.
type KeyInfo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyInfo" json:"-"`
	KeyInfoType
}

// KeyName represents the global element {http://www.w3.org/2000/09/xmldsig#}KeyName.
type KeyName struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyName" json:"-"`
	Value   String   `xml:",chardata" json:"value"`
}

// KeyValue represents the global element {http://www.w3.org/2000/09/xmldsig#}KeyValue.
type KeyValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# KeyValue" json:"-"`
	KeyValueType
}

// Manifest represents the global element {http://www.w3.org/2000/09/xmldsig#}Manifest.
type Manifest struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Manifest" json:"-"`
	ManifestType
}

// MgmtData represents the global element {http://www.w3.org/2000/09/xmldsig#}MgmtData.
type MgmtData struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# MgmtData" json:"-"`
	Value   String   `xml:",chardata" json:"value"`
}

// Object represents the global element {http://www.w3.org/2000/09/xmldsig#}Object.
type Object struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Object" json:"-"`
	ObjectType
}

// PGPData represents the global element {http://www.w3.org/2000/09/xmldsig#}PGPData.
type PGPData struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# PGPData" json:"-"`
	PGPDataType
}

// RSAKeyValue represents the global element {http://www.w3.org/2000/09/xmldsig#}RSAKeyValue.
type RSAKeyValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# RSAKeyValue" json:"-"`
	RSAKeyValueType
}

// Reference represents the global element {http://www.w3.org/2000/09/xmldsig#}Reference.
type Reference struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Reference" json:"-"`
	ReferenceType
}

// RetrievalMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}RetrievalMethod.
type RetrievalMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# RetrievalMethod" json:"-"`
	RetrievalMethodType
}

// SPKIData represents the global element {http://www.w3.org/2000/09/xmldsig#}SPKIData.
type SPKIData struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SPKIData" json:"-"`
	SPKIDataType
}

// Signature represents the global element {http://www.w3.org/2000/09/xmldsig#}Signature.
type Signature struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"-"`
	SignatureType
}

// SignatureMethod represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureMethod.
type SignatureMethod struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureMethod" json:"-"`
	SignatureMethodType
}

// SignatureProperties represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureProperties.
type SignatureProperties struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureProperties" json:"-"`
	SignaturePropertiesType
}

// SignatureProperty represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureProperty.
type SignatureProperty struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureProperty" json:"-"`
	SignaturePropertyType
}

// SignatureValue represents the global element {http://www.w3.org/2000/09/xmldsig#}SignatureValue.
type SignatureValue struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignatureValue" json:"-"`
	SignatureValueType
}

// SignedInfo represents the global element {http://www.w3.org/2000/09/xmldsig#}SignedInfo.
type SignedInfo struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# SignedInfo" json:"-"`
	SignedInfoType
}

// Transform represents the global element {http://www.w3.org/2000/09/xmldsig#}Transform.
type Transform struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Transform" json:"-"`
	TransformType
}

// Transforms represents the global element {http://www.w3.org/2000/09/xmldsig#}Transforms.
type Transforms struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# Transforms" json:"-"`
	TransformsType
}

// X509Data represents the global element {http://www.w3.org/2000/09/xmldsig#}X509Data.
type X509Data struct {
	XMLName xml.Name `xml:"http://www.w3.org/2000/09/xmldsig# X509Data" json:"-"`
	X509DataType
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
