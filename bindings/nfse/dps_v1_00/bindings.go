// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package dps_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TCAtvEvento struct {
	Desc  TSDesc255          `xml:"http://www.sped.fazenda.gov.br/nfse desc" json:"desc"`
	DtIni Date               `xml:"http://www.sped.fazenda.gov.br/nfse dtIni" json:"dtIni"`
	DtFim Date               `xml:"http://www.sped.fazenda.gov.br/nfse dtFim" json:"dtFim"`
	ID    *TSIdeEvento       `xml:"http://www.sped.fazenda.gov.br/nfse id,omitempty" json:"id,omitempty"`
	End   *TCEnderecoSimples `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
}

type TCBeneficioMunicipal struct {
	TpBM     TSOpTipoBM              `xml:"http://www.sped.fazenda.gov.br/nfse tpBM" json:"tpBM"`
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
	CEndPost    TSCodigoEndPostal      `xml:"http://www.sped.fazenda.gov.br/nfse cEndPost" json:"cEndPost"`
	XLgr        TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfse xLgr" json:"xLgr"`
	Nro         TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfse nro" json:"nro"`
	XCpl        *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfse xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro     TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfse xBairro" json:"xBairro"`
	XCidade     TSCidade               `xml:"http://www.sped.fazenda.gov.br/nfse xCidade" json:"xCidade"`
	XEstProvReg TSEstadoProvRegiao     `xml:"http://www.sped.fazenda.gov.br/nfse xEstProvReg" json:"xEstProvReg"`
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

type TCExploracaoRodoviaria struct {
	CategVeic    TSCategVeic    `xml:"http://www.sped.fazenda.gov.br/nfse categVeic" json:"categVeic"`
	NEixos       TSNumEixos     `xml:"http://www.sped.fazenda.gov.br/nfse nEixos" json:"nEixos"`
	Rodagem      TSRodagem      `xml:"http://www.sped.fazenda.gov.br/nfse rodagem" json:"rodagem"`
	Sentido      TSSentido      `xml:"http://www.sped.fazenda.gov.br/nfse sentido" json:"sentido"`
	Placa        TSPlaca        `xml:"http://www.sped.fazenda.gov.br/nfse placa" json:"placa"`
	CodAcessoPed TSCodAcessoPed `xml:"http://www.sped.fazenda.gov.br/nfse codAcessoPed" json:"codAcessoPed"`
	CodContrato  TSCodContrato  `xml:"http://www.sped.fazenda.gov.br/nfse codContrato" json:"codContrato"`
}

type TCInfDPS struct {
	ID       TSIDDPS         `xml:"Id,attr" json:"Id"`
	TpAmb    TSTipoAmbiente  `xml:"http://www.sped.fazenda.gov.br/nfse tpAmb" json:"tpAmb"`
	DhEmi    TSDateTimeUTC   `xml:"http://www.sped.fazenda.gov.br/nfse dhEmi" json:"dhEmi"`
	VerAplic TSVerAplic      `xml:"http://www.sped.fazenda.gov.br/nfse verAplic" json:"verAplic"`
	Serie    TSSerieDPS      `xml:"http://www.sped.fazenda.gov.br/nfse serie" json:"serie"`
	NDPS     TSNumDPS        `xml:"http://www.sped.fazenda.gov.br/nfse nDPS" json:"nDPS"`
	DCompet  TSData          `xml:"http://www.sped.fazenda.gov.br/nfse dCompet" json:"dCompet"`
	TpEmit   TSEmitenteDPS   `xml:"http://www.sped.fazenda.gov.br/nfse tpEmit" json:"tpEmit"`
	CLocEmi  TSCodMunIBGE    `xml:"http://www.sped.fazenda.gov.br/nfse cLocEmi" json:"cLocEmi"`
	Subst    *TCSubstituicao `xml:"http://www.sped.fazenda.gov.br/nfse subst,omitempty" json:"subst,omitempty"`
	Prest    TCInfoPrestador `xml:"http://www.sped.fazenda.gov.br/nfse prest" json:"prest"`
	Toma     *TCInfoPessoa   `xml:"http://www.sped.fazenda.gov.br/nfse toma,omitempty" json:"toma,omitempty"`
	Interm   *TCInfoPessoa   `xml:"http://www.sped.fazenda.gov.br/nfse interm,omitempty" json:"interm,omitempty"`
	Serv     TCServ          `xml:"http://www.sped.fazenda.gov.br/nfse serv" json:"serv"`
	Valores  TCInfoValores   `xml:"http://www.sped.fazenda.gov.br/nfse valores" json:"valores"`
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
	DPS           TCDPS            `xml:"http://www.sped.fazenda.gov.br/nfse DPS" json:"DPS"`
}

type TCInfoCompl struct {
	IDDocTec *TSDRT          `xml:"http://www.sped.fazenda.gov.br/nfse idDocTec,omitempty" json:"idDocTec,omitempty"`
	DocRef   *TSDesc255      `xml:"http://www.sped.fazenda.gov.br/nfse docRef,omitempty" json:"docRef,omitempty"`
	XInfComp *TSDescInfCompl `xml:"http://www.sped.fazenda.gov.br/nfse xInfComp,omitempty" json:"xInfComp,omitempty"`
}

type TCInfoDedRed struct {
	PDR        *TSDec3V2         `xml:"http://www.sped.fazenda.gov.br/nfse pDR,omitempty" json:"pDR,omitempty"`
	VDR        *TSDec15V2        `xml:"http://www.sped.fazenda.gov.br/nfse vDR,omitempty" json:"vDR,omitempty"`
	Documentos *TCListaDocDedRed `xml:"http://www.sped.fazenda.gov.br/nfse documentos,omitempty" json:"documentos,omitempty"`
}

type TCInfoObra struct {
	CObra        *TSCodObra         `xml:"http://www.sped.fazenda.gov.br/nfse cObra,omitempty" json:"cObra,omitempty"`
	InscImobFisc *TSInscImobFisc    `xml:"http://www.sped.fazenda.gov.br/nfse inscImobFisc,omitempty" json:"inscImobFisc,omitempty"`
	End          *TCEnderecoSimples `xml:"http://www.sped.fazenda.gov.br/nfse end,omitempty" json:"end,omitempty"`
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

type TCLocacaoSublocacao struct {
	Categ    TSCategoriaServico `xml:"http://www.sped.fazenda.gov.br/nfse categ" json:"categ"`
	Objeto   TCObjetoLocacao    `xml:"http://www.sped.fazenda.gov.br/nfse objeto" json:"objeto"`
	Extensao TSExtensaoTotal    `xml:"http://www.sped.fazenda.gov.br/nfse extensao" json:"extensao"`
	NPostes  TSNumeroPostes     `xml:"http://www.sped.fazenda.gov.br/nfse nPostes" json:"nPostes"`
}

type TCNFSe struct {
	Versao    TVerNFSe      `xml:"versao,attr" json:"versao"`
	InfNFSe   TCInfNFSe     `xml:"http://www.sped.fazenda.gov.br/nfse infNFSe" json:"infNFSe"`
	Signature SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCObjetoLocacao string

type TCRegTrib struct {
	OpSimpNac   TSOpSimpNac              `xml:"http://www.sped.fazenda.gov.br/nfse opSimpNac" json:"opSimpNac"`
	RegApTribSN *TSRegimeApuracaoSimpNac `xml:"http://www.sped.fazenda.gov.br/nfse regApTribSN,omitempty" json:"regApTribSN,omitempty"`
	RegEspTrib  TSRegEspTrib             `xml:"http://www.sped.fazenda.gov.br/nfse regEspTrib" json:"regEspTrib"`
}

type TCServ struct {
	LocPrest  TCLocPrest              `xml:"http://www.sped.fazenda.gov.br/nfse locPrest" json:"locPrest"`
	CServ     TCCServ                 `xml:"http://www.sped.fazenda.gov.br/nfse cServ" json:"cServ"`
	ComExt    *TCComExterior          `xml:"http://www.sped.fazenda.gov.br/nfse comExt,omitempty" json:"comExt,omitempty"`
	Lsadppu   *TCLocacaoSublocacao    `xml:"http://www.sped.fazenda.gov.br/nfse lsadppu,omitempty" json:"lsadppu,omitempty"`
	Obra      *TCInfoObra             `xml:"http://www.sped.fazenda.gov.br/nfse obra,omitempty" json:"obra,omitempty"`
	AtvEvento *TCAtvEvento            `xml:"http://www.sped.fazenda.gov.br/nfse atvEvento,omitempty" json:"atvEvento,omitempty"`
	ExplRod   *TCExploracaoRodoviaria `xml:"http://www.sped.fazenda.gov.br/nfse explRod,omitempty" json:"explRod,omitempty"`
	InfoCompl *TCInfoCompl            `xml:"http://www.sped.fazenda.gov.br/nfse infoCompl,omitempty" json:"infoCompl,omitempty"`
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
	BM          *TCBeneficioMunicipal `xml:"http://www.sped.fazenda.gov.br/nfse BM,omitempty" json:"BM,omitempty"`
	ExigSusp    *TCExigSuspensa       `xml:"http://www.sped.fazenda.gov.br/nfse exigSusp,omitempty" json:"exigSusp,omitempty"`
	TpImunidade *TSTipoImunidadeISSQN `xml:"http://www.sped.fazenda.gov.br/nfse tpImunidade,omitempty" json:"tpImunidade,omitempty"`
	PAliq       *TSDec1V2             `xml:"http://www.sped.fazenda.gov.br/nfse pAliq,omitempty" json:"pAliq,omitempty"`
	TpRetISSQN  TSTipoRetISSQN        `xml:"http://www.sped.fazenda.gov.br/nfse tpRetISSQN" json:"tpRetISSQN"`
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
	VCalcDR    *TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vCalcDR,omitempty" json:"vCalcDR,omitempty"`
	TpBM       *TSDesc40   `xml:"http://www.sped.fazenda.gov.br/nfse tpBM,omitempty" json:"tpBM,omitempty"`
	VCalcBM    *TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vCalcBM,omitempty" json:"vCalcBM,omitempty"`
	VBC        *TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vBC,omitempty" json:"vBC,omitempty"`
	PAliqAplic *TSDec1V2   `xml:"http://www.sped.fazenda.gov.br/nfse pAliqAplic,omitempty" json:"pAliqAplic,omitempty"`
	VISSQN     *TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vISSQN,omitempty" json:"vISSQN,omitempty"`
	VTotalRet  *TSDec15V2  `xml:"http://www.sped.fazenda.gov.br/nfse vTotalRet,omitempty" json:"vTotalRet,omitempty"`
	VLiq       TSDec15V2   `xml:"http://www.sped.fazenda.gov.br/nfse vLiq" json:"vLiq"`
	XOutInf    *TSDesc2000 `xml:"http://www.sped.fazenda.gov.br/nfse xOutInf,omitempty" json:"xOutInf,omitempty"`
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

type TSDataCNC string

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

type TSDesc600 string

type TSDescInfCompl string

type TSDescOutDedRed string

type TSEmail string

type TSEmitenteDPS string

type TSEnderCompletoExt string

type TSEnvMDIC string

type TSEstadoProvRegiao string

type TSExtensaoTotal string

type TSIDCNC string

type TSIDDPS string

type TSIDEvento string

type TSIDNFSe string

type TSIDNumEvento string

type TSIDPedRefEvt string

type TSIdeDedRed string

type TSIdeEvento string

type TSInscImobFisc string

type TSInscMun string

type TSLogradouro string

type TSMecAFComExPrest string

type TSMecAFComExToma string

type TSModoPrestacao string

type TSMotivo string

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

type TSOpTipoBM string

type TSPlaca string

type TSProcEmissao string

type TSRegEspTrib string

type TSRegimeApuracaoSimpNac string

type TSRodagem string

type TSSentido string

type TSSerieDPS string

type TSSerieNFNFS string

type TSSituacaoCadastroContribuinte string

type TSSituacaoCadastroContribuinteCNC string

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

type TVerNFSe string

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

type Date string

type String string

// DPS represents the global element {http://www.sped.fazenda.gov.br/nfse}DPS.
type DPS struct {
	XMLName xml.Name `xml:"http://www.sped.fazenda.gov.br/nfse DPS" json:"-"`
	TCDPS
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
