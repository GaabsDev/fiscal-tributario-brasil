// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package nfeouro_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type TCCompradorOuro struct {
	TpComp TSTipoCompOuro     `xml:"http://www.sped.fazenda.gov.br/nfeouro tpComp" json:"tpComp"`
	CNPJ   *TSCNPJ            `xml:"http://www.sped.fazenda.gov.br/nfeouro CNPJ,omitempty" json:"CNPJ,omitempty"`
	IDCPF  *TCIDCPFCompVend   `xml:"http://www.sped.fazenda.gov.br/nfeouro idCPF,omitempty" json:"idCPF,omitempty"`
	NIF    *TSNIF             `xml:"http://www.sped.fazenda.gov.br/nfeouro NIF,omitempty" json:"NIF,omitempty"`
	XNome  *TSNomeFantasia    `xml:"http://www.sped.fazenda.gov.br/nfeouro xNome,omitempty" json:"xNome,omitempty"`
	End    TCEnderecoVendedor `xml:"http://www.sped.fazenda.gov.br/nfeouro end" json:"end"`
	Fone   *TSTelefone        `xml:"http://www.sped.fazenda.gov.br/nfeouro fone,omitempty" json:"fone,omitempty"`
	Email  *TSEmail           `xml:"http://www.sped.fazenda.gov.br/nfeouro email,omitempty" json:"email,omitempty"`
}

type TCDAO struct {
	Versao    TVerNFeOuro    `xml:"versao,attr" json:"versao"`
	InfDAO    TCInfDAO       `xml:"http://www.sped.fazenda.gov.br/nfeouro infDAO" json:"infDAO"`
	Signature *SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TCDetInfoOuro struct {
	TpOuro       TSTipologiaOuro    `xml:"http://www.sped.fazenda.gov.br/nfeouro tpOuro" json:"tpOuro"`
	TeorOuro     TSDec2V2           `xml:"http://www.sped.fazenda.gov.br/nfeouro teorOuro" json:"teorOuro"`
	Metodo       TSMetodoPurezaOuro `xml:"http://www.sped.fazenda.gov.br/nfeouro metodo" json:"metodo"`
	QtdOuroBruto TSDec15V2          `xml:"http://www.sped.fazenda.gov.br/nfeouro qtdOuroBruto" json:"qtdOuroBruto"`
	CNCM         TSNomComumMercosul `xml:"http://www.sped.fazenda.gov.br/nfeouro cNCM" json:"cNCM"`
}

type TCDetOuro struct {
	VUnGrOuro TSDec15V2       `xml:"http://www.sped.fazenda.gov.br/nfeouro vUnGrOuro" json:"vUnGrOuro"`
	InfOuro   []TCDetInfoOuro `xml:"http://www.sped.fazenda.gov.br/nfeouro infOuro,omitempty" json:"infOuro,omitempty"`
}

type TCEmitenteDAO struct {
	CNPJ  TSCNPJ          `xml:"http://www.sped.fazenda.gov.br/nfeouro CNPJ" json:"CNPJ"`
	IM    TSInscMun       `xml:"http://www.sped.fazenda.gov.br/nfeouro IM" json:"IM"`
	XNome *TSNomeFantasia `xml:"http://www.sped.fazenda.gov.br/nfeouro xNome,omitempty" json:"xNome,omitempty"`
	End   *TCEnderecoEmit `xml:"http://www.sped.fazenda.gov.br/nfeouro end,omitempty" json:"end,omitempty"`
	Fone  *TSTelefone     `xml:"http://www.sped.fazenda.gov.br/nfeouro fone,omitempty" json:"fone,omitempty"`
	Email *TSEmail        `xml:"http://www.sped.fazenda.gov.br/nfeouro email,omitempty" json:"email,omitempty"`
}

type TCEndExterior struct {
	CPais       TSCodPaisISO       `xml:"http://www.sped.fazenda.gov.br/nfeouro cPais" json:"cPais"`
	CEndPost    TSCodigoEndPostal  `xml:"http://www.sped.fazenda.gov.br/nfeouro cEndPost" json:"cEndPost"`
	XCidade     TSCidade           `xml:"http://www.sped.fazenda.gov.br/nfeouro xCidade" json:"xCidade"`
	XEstProvReg TSEstadoProvRegiao `xml:"http://www.sped.fazenda.gov.br/nfeouro xEstProvReg" json:"xEstProvReg"`
}

type TCEndNacional struct {
	CMun TSCodMunIBGE `xml:"http://www.sped.fazenda.gov.br/nfeouro cMun" json:"cMun"`
	CEP  TSCEP        `xml:"http://www.sped.fazenda.gov.br/nfeouro CEP" json:"CEP"`
}

type TCEnderecoEmit struct {
	CMun    TSCodMunIBGE           `xml:"http://www.sped.fazenda.gov.br/nfeouro cMun" json:"cMun"`
	CEP     TSCEP                  `xml:"http://www.sped.fazenda.gov.br/nfeouro CEP" json:"CEP"`
	XLgr    TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfeouro xLgr" json:"xLgr"`
	Nro     TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfeouro nro" json:"nro"`
	XCpl    *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfeouro xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfeouro xBairro" json:"xBairro"`
}

type TCEnderecoPostoAtend struct {
	CMun    TSCodMunIBGE           `xml:"http://www.sped.fazenda.gov.br/nfeouro cMun" json:"cMun"`
	CEP     *TSCEP                 `xml:"http://www.sped.fazenda.gov.br/nfeouro CEP,omitempty" json:"CEP,omitempty"`
	XLgr    TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfeouro xLgr" json:"xLgr"`
	Nro     TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfeouro nro" json:"nro"`
	XCpl    *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfeouro xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfeouro xBairro" json:"xBairro"`
}

type TCEnderecoVendedor struct {
	EndNac  *TCEndNacional         `xml:"http://www.sped.fazenda.gov.br/nfeouro endNac,omitempty" json:"endNac,omitempty"`
	EndExt  *TCEndExterior         `xml:"http://www.sped.fazenda.gov.br/nfeouro endExt,omitempty" json:"endExt,omitempty"`
	XLgr    TSLogradouro           `xml:"http://www.sped.fazenda.gov.br/nfeouro xLgr" json:"xLgr"`
	Nro     TSNumeroEndereco       `xml:"http://www.sped.fazenda.gov.br/nfeouro nro" json:"nro"`
	XCpl    *TSComplementoEndereco `xml:"http://www.sped.fazenda.gov.br/nfeouro xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro TSBairro               `xml:"http://www.sped.fazenda.gov.br/nfeouro xBairro" json:"xBairro"`
}

type TCIDCNPJVendedor struct {
	CNPJ     TSCNPJ `xml:"http://www.sped.fazenda.gov.br/nfeouro CNPJ" json:"CNPJ"`
	CPFRepPJ TSCPF  `xml:"http://www.sped.fazenda.gov.br/nfeouro cpfRepPJ" json:"cpfRepPJ"`
}

type TCIDCPFCompVend struct {
	CPF    TSCPF         `xml:"http://www.sped.fazenda.gov.br/nfeouro CPF" json:"CPF"`
	NDI    TSNumDocIdent `xml:"http://www.sped.fazenda.gov.br/nfeouro nDI" json:"nDI"`
	OrgExp TSDesc60      `xml:"http://www.sped.fazenda.gov.br/nfeouro orgExp" json:"orgExp"`
}

type TCIDTranspOuro struct {
	CNPJ     *TSCNPJ        `xml:"http://www.sped.fazenda.gov.br/nfeouro CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF      *TSCPF         `xml:"http://www.sped.fazenda.gov.br/nfeouro CPF,omitempty" json:"CPF,omitempty"`
	XNome    TSNomeFantasia `xml:"http://www.sped.fazenda.gov.br/nfeouro xNome" json:"xNome"`
	CMunOr   TSCodMunIBGE   `xml:"http://www.sped.fazenda.gov.br/nfeouro cMunOr" json:"cMunOr"`
	CMunDest TSCodMunIBGE   `xml:"http://www.sped.fazenda.gov.br/nfeouro cMunDest" json:"cMunDest"`
}

type TCInfAquisOuro struct {
	NTitLavra      TSNumTitLavra      `xml:"http://www.sped.fazenda.gov.br/nfeouro nTitLavra" json:"nTitLavra"`
	CNPJ           *TSCNPJ            `xml:"http://www.sped.fazenda.gov.br/nfeouro CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF            *TSCPF             `xml:"http://www.sped.fazenda.gov.br/nfeouro CPF,omitempty" json:"CPF,omitempty"`
	NProcPermLavra TSNumProcPermLavra `xml:"http://www.sped.fazenda.gov.br/nfeouro nProcPermLavra" json:"nProcPermLavra"`
	CMun           TSCodMunIBGE       `xml:"http://www.sped.fazenda.gov.br/nfeouro cMun" json:"cMun"`
	AreaLavra      TSDesc255          `xml:"http://www.sped.fazenda.gov.br/nfeouro areaLavra" json:"areaLavra"`
	FaseProcPerm   TSTipoFaseProcPerm `xml:"http://www.sped.fazenda.gov.br/nfeouro faseProcPerm" json:"faseProcPerm"`
	SitPLG         TSTipoSituacoPLG   `xml:"http://www.sped.fazenda.gov.br/nfeouro sitPLG" json:"sitPLG"`
	DtVencPLG      TSData             `xml:"http://www.sped.fazenda.gov.br/nfeouro dtVencPLG" json:"dtVencPLG"`
	NLicAmb        TSNumLicAmbiental  `xml:"http://www.sped.fazenda.gov.br/nfeouro nLicAmb" json:"nLicAmb"`
	OrgLicAmb      TSDesc255          `xml:"http://www.sped.fazenda.gov.br/nfeouro orgLicAmb" json:"orgLicAmb"`
	NGuia          *TSNumGuia         `xml:"http://www.sped.fazenda.gov.br/nfeouro nGuia,omitempty" json:"nGuia,omitempty"`
	DEmisGuia      *TSData            `xml:"http://www.sped.fazenda.gov.br/nfeouro dEmisGuia,omitempty" json:"dEmisGuia,omitempty"`
	DFimTrans      *TSData            `xml:"http://www.sped.fazenda.gov.br/nfeouro dFimTrans,omitempty" json:"dFimTrans,omitempty"`
	LocalAq        TCInfLocalAquis    `xml:"http://www.sped.fazenda.gov.br/nfeouro localAq" json:"localAq"`
	PgtoOp         *TCInfPgtoAquis    `xml:"http://www.sped.fazenda.gov.br/nfeouro pgtoOp,omitempty" json:"pgtoOp,omitempty"`
}

type TCInfComercioExtOuro struct {
	NEdmov TSNumEDmov    `xml:"http://www.sped.fazenda.gov.br/nfeouro nEdmov" json:"nEdmov"`
	CMun   *TSCodMunIBGE `xml:"http://www.sped.fazenda.gov.br/nfeouro cMun,omitempty" json:"cMun,omitempty"`
}

type TCInfDAO struct {
	ID        TSIDDAO             `xml:"Id,attr" json:"Id"`
	TpAmb     TSTipoAmbiente      `xml:"http://www.sped.fazenda.gov.br/nfeouro tpAmb" json:"tpAmb"`
	DhEmi     TSDateTimeUTC       `xml:"http://www.sped.fazenda.gov.br/nfeouro dhEmi" json:"dhEmi"`
	VerAplic  TSVerAplic          `xml:"http://www.sped.fazenda.gov.br/nfeouro verAplic" json:"verAplic"`
	SerieDAO  TSSerieDAO          `xml:"http://www.sped.fazenda.gov.br/nfeouro serieDAO" json:"serieDAO"`
	NDAO      TSNumDAO            `xml:"http://www.sped.fazenda.gov.br/nfeouro nDAO" json:"nDAO"`
	OpDAO     TSTipoOperDAO       `xml:"http://www.sped.fazenda.gov.br/nfeouro opDAO" json:"opDAO"`
	TpOp      TSTipoPapelEmitente `xml:"http://www.sped.fazenda.gov.br/nfeouro tpOp" json:"tpOp"`
	DCompet   TSData              `xml:"http://www.sped.fazenda.gov.br/nfeouro dCompet" json:"dCompet"`
	CMunEmi   TSCodMunIBGE        `xml:"http://www.sped.fazenda.gov.br/nfeouro cMunEmi" json:"cMunEmi"`
	Subst     *TCSubstituicao     `xml:"http://www.sped.fazenda.gov.br/nfeouro subst,omitempty" json:"subst,omitempty"`
	Emit      TCEmitenteDAO       `xml:"http://www.sped.fazenda.gov.br/nfeouro emit" json:"emit"`
	Vendedor  *TCVendedorOuro     `xml:"http://www.sped.fazenda.gov.br/nfeouro vendedor,omitempty" json:"vendedor,omitempty"`
	Comprador *TCCompradorOuro    `xml:"http://www.sped.fazenda.gov.br/nfeouro comprador,omitempty" json:"comprador,omitempty"`
	Transp    *TCTranspOuro       `xml:"http://www.sped.fazenda.gov.br/nfeouro transp,omitempty" json:"transp,omitempty"`
	InfOp     TCInfOperOuro       `xml:"http://www.sped.fazenda.gov.br/nfeouro infOp" json:"infOp"`
	NotasRef  *TCNotasReferenc    `xml:"http://www.sped.fazenda.gov.br/nfeouro notasRef,omitempty" json:"notasRef,omitempty"`
	Ouro      TCDetOuro           `xml:"http://www.sped.fazenda.gov.br/nfeouro ouro" json:"ouro"`
	XInfComp  *TSDesc255          `xml:"http://www.sped.fazenda.gov.br/nfeouro xInfComp,omitempty" json:"xInfComp,omitempty"`
}

type TCInfLocalAquis struct {
	CPF     TSCPF                `xml:"http://www.sped.fazenda.gov.br/nfeouro CPF" json:"CPF"`
	NUnicad *TSNumDoisDigitos    `xml:"http://www.sped.fazenda.gov.br/nfeouro nUnicad,omitempty" json:"nUnicad,omitempty"`
	XNome   TSNomeFantasia       `xml:"http://www.sped.fazenda.gov.br/nfeouro xNome" json:"xNome"`
	End     TCEnderecoPostoAtend `xml:"http://www.sped.fazenda.gov.br/nfeouro end" json:"end"`
	Fone    *TSTelefone          `xml:"http://www.sped.fazenda.gov.br/nfeouro fone,omitempty" json:"fone,omitempty"`
	Email   *TSEmail             `xml:"http://www.sped.fazenda.gov.br/nfeouro email,omitempty" json:"email,omitempty"`
}

type TCInfNFeOuro struct {
}

type TCInfNegocOuro struct {
	NatOp      TSTipoNatNegocOuro `xml:"http://www.sped.fazenda.gov.br/nfeouro natOp" json:"natOp"`
	CustodCNPJ *TSCNPJ            `xml:"http://www.sped.fazenda.gov.br/nfeouro custodCNPJ,omitempty" json:"custodCNPJ,omitempty"`
}

type TCInfOperOuro struct {
	Aquis  *TCInfAquisOuro       `xml:"http://www.sped.fazenda.gov.br/nfeouro aquis,omitempty" json:"aquis,omitempty"`
	Regmin *TCRegAprMineral      `xml:"http://www.sped.fazenda.gov.br/nfeouro regmin,omitempty" json:"regmin,omitempty"`
	Remes  *TCInfRemessaOuro     `xml:"http://www.sped.fazenda.gov.br/nfeouro remes,omitempty" json:"remes,omitempty"`
	Negoc  *TCInfNegocOuro       `xml:"http://www.sped.fazenda.gov.br/nfeouro negoc,omitempty" json:"negoc,omitempty"`
	Comex  *TCInfComercioExtOuro `xml:"http://www.sped.fazenda.gov.br/nfeouro comex,omitempty" json:"comex,omitempty"`
}

type TCInfPgtoAquis struct {
	NBanco   TSNumBanco   `xml:"http://www.sped.fazenda.gov.br/nfeouro nBanco" json:"nBanco"`
	NAgencia TSNumAgencia `xml:"http://www.sped.fazenda.gov.br/nfeouro nAgencia" json:"nAgencia"`
	NConta   TSNumConta   `xml:"http://www.sped.fazenda.gov.br/nfeouro nConta" json:"nConta"`
}

type TCInfRemessaOuro struct {
	NatRem TSTipoNatRemessa `xml:"http://www.sped.fazenda.gov.br/nfeouro natRem" json:"natRem"`
	CNPJ   *TSCNPJ          `xml:"http://www.sped.fazenda.gov.br/nfeouro CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF    *TSCPF           `xml:"http://www.sped.fazenda.gov.br/nfeouro CPF,omitempty" json:"CPF,omitempty"`
	XNome  *TSNomeFantasia  `xml:"http://www.sped.fazenda.gov.br/nfeouro xNome,omitempty" json:"xNome,omitempty"`
	End    *TCEnderecoEmit  `xml:"http://www.sped.fazenda.gov.br/nfeouro end,omitempty" json:"end,omitempty"`
	Fone   *TSTelefone      `xml:"http://www.sped.fazenda.gov.br/nfeouro fone,omitempty" json:"fone,omitempty"`
	Email  *TSEmail         `xml:"http://www.sped.fazenda.gov.br/nfeouro email,omitempty" json:"email,omitempty"`
}

type TCNFeOuro struct {
	Versao     TVerNFeOuro   `xml:"versao,attr" json:"versao"`
	InfNFeOuro TCInfNFeOuro  `xml:"http://www.sped.fazenda.gov.br/nfeouro infNFeOuro" json:"infNFeOuro"`
	Signature  SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TCNotasReferenc struct {
	ChaveAcesso []TSChaveNFeOuro `xml:"http://www.sped.fazenda.gov.br/nfeouro chaveAcesso,omitempty" json:"chaveAcesso,omitempty"`
}

type TCRegAprMineral struct {
	XInfo TSDescPrimVendaOuro `xml:"http://www.sped.fazenda.gov.br/nfeouro xInfo" json:"xInfo"`
}

type TCSubstituicao struct {
	ChSubstda TSChaveNFeOuro  `xml:"http://www.sped.fazenda.gov.br/nfeouro chSubstda" json:"chSubstda"`
	CMotivo   TSCodMotivSubst `xml:"http://www.sped.fazenda.gov.br/nfeouro cMotivo" json:"cMotivo"`
	XMotivo   *TSMotivo       `xml:"http://www.sped.fazenda.gov.br/nfeouro xMotivo,omitempty" json:"xMotivo,omitempty"`
}

type TCTranspOuro struct {
	IDTransp []TCIDTranspOuro `xml:"http://www.sped.fazenda.gov.br/nfeouro idTransp,omitempty" json:"idTransp,omitempty"`
}

type TCVendedorOuro struct {
	TpVend   TSTipoVendOuro      `xml:"http://www.sped.fazenda.gov.br/nfeouro tpVend" json:"tpVend"`
	IDCNPJ   *TCIDCNPJVendedor   `xml:"http://www.sped.fazenda.gov.br/nfeouro idCNPJ,omitempty" json:"idCNPJ,omitempty"`
	IDCPF    *TCIDCPFCompVend    `xml:"http://www.sped.fazenda.gov.br/nfeouro idCPF,omitempty" json:"idCPF,omitempty"`
	NIF      *TSNIF              `xml:"http://www.sped.fazenda.gov.br/nfeouro NIF,omitempty" json:"NIF,omitempty"`
	XNome    *TSNomeFantasia     `xml:"http://www.sped.fazenda.gov.br/nfeouro xNome,omitempty" json:"xNome,omitempty"`
	NRegVend *TSNumRegComercVend `xml:"http://www.sped.fazenda.gov.br/nfeouro nRegVend,omitempty" json:"nRegVend,omitempty"`
	End      TCEnderecoVendedor  `xml:"http://www.sped.fazenda.gov.br/nfeouro end" json:"end"`
	Fone     *TSTelefone         `xml:"http://www.sped.fazenda.gov.br/nfeouro fone,omitempty" json:"fone,omitempty"`
	Email    *TSEmail            `xml:"http://www.sped.fazenda.gov.br/nfeouro email,omitempty" json:"email,omitempty"`
}

type TSBairro string

type TSCEP string

type TSCNPJ string

type TSCPF string

type TSChaveNFeOuro string

type TSCidade string

type TSCodMotivSubst string

type TSCodMunIBGE string

type TSCodPaisISO string

type TSCodigoEndPostal string

type TSComplementoEndereco string

type TSData string

type TSDateTimeUTC string

type TSDec15V2 string

type TSDec1V2 string

type TSDec2V2 string

type TSDesc255 string

type TSDesc60 string

type TSDescPrimVendaOuro string

type TSEmail string

type TSEstadoProvRegiao string

type TSIDDAO string

type TSInscMun string

type TSLogradouro string

type TSMetodoPurezaOuro string

type TSMotivo string

type TSNIF string

type TSNomComumMercosul string

type TSNomeFantasia string

type TSNumAgencia string

type TSNumBanco string

type TSNumConta string

type TSNumDAO string

type TSNumDocIdent string

type TSNumDoisDigitos string

type TSNumEDmov string

type TSNumGuia string

type TSNumLicAmbiental string

type TSNumProcPermLavra string

type TSNumRegComercVend string

type TSNumTitLavra string

type TSNumeroEndereco string

type TSSerieDAO string

type TSString string

type TSStringComQuebraDeLinha string

type TSTelefone string

type TSTipoAmbiente string

type TSTipoCompOuro string

type TSTipoFaseProcPerm string

type TSTipoInscFederal string

type TSTipoNatNegocOuro string

type TSTipoNatRemessa string

type TSTipoOperDAO string

type TSTipoPapelEmitente string

type TSTipoSituacoPLG string

type TSTipoVendOuro string

type TSTipologiaOuro string

type TSVerAplic string

type TVerNFeOuro string

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

type String string

// NFeOuro represents the global element {http://www.sped.fazenda.gov.br/nfeouro}NFeOuro.
type NFeOuro struct {
	XMLName xml.Name `xml:"http://www.sped.fazenda.gov.br/nfeouro NFeOuro" json:"-"`
	TCNFeOuro
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
