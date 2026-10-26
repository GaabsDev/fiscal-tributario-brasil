// Code generated from captured official XSDs. DO NOT EDIT.
// Decimal, date and identifier values retain their exact XML lexical representation.
package nfeabi_v1_00

import "encoding/xml"

type RawElement struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr"`
	InnerXML string     `xml:",innerxml"`
}

type GTorna struct {
	VTorna      TDec1302  `xml:"http://www.portalfiscal.inf.br/nfeabi vTorna" json:"vTorna"`
	ImovelTorna []TImovel `xml:"http://www.portalfiscal.inf.br/nfeabi imovelTorna,omitempty" json:"imovelTorna,omitempty"`
}

type TAdquirente struct {
	NAdquir         string    `xml:"nAdquir,attr" json:"nAdquir"`
	IndDeclarante   *string   `xml:"http://www.portalfiscal.inf.br/nfeabi indDeclarante,omitempty" json:"indDeclarante,omitempty"`
	CNPJ            *TCNPJ    `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF             *TCPF     `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
	XNome           string    `xml:"http://www.portalfiscal.inf.br/nfeabi xNome" json:"xNome"`
	EnderAdquirente TEndereco `xml:"http://www.portalfiscal.inf.br/nfeabi enderAdquirente" json:"enderAdquirente"`
	PAquisicao      *TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAquisicao,omitempty" json:"pAquisicao,omitempty"`
	IndContrib      *string   `xml:"http://www.portalfiscal.inf.br/nfeabi indContrib,omitempty" json:"indContrib,omitempty"`
}

type TAmb string

type TAutXML struct {
	NAutXML String `xml:"nAutXML,attr" json:"nAutXML"`
	CNPJ    *TCNPJ `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF     *TCPF  `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
}

type TCOrgaoIBGE string

type TCadastroImovel struct {
	CCIB     *string `xml:"http://www.portalfiscal.inf.br/nfeabi cCIB,omitempty" json:"cCIB,omitempty"`
	IndIPTU  *string `xml:"http://www.portalfiscal.inf.br/nfeabi indIPTU,omitempty" json:"indIPTU,omitempty"`
	NIptu    *string `xml:"http://www.portalfiscal.inf.br/nfeabi nIptu,omitempty" json:"nIptu,omitempty"`
	IndIncra *string `xml:"http://www.portalfiscal.inf.br/nfeabi indIncra,omitempty" json:"indIncra,omitempty"`
	NIncra   *string `xml:"http://www.portalfiscal.inf.br/nfeabi nIncra,omitempty" json:"nIncra,omitempty"`
}

type TChDFe string

type TChNFeABI string

type TCNPJ string

type TCNPJOpc string

type TCNPJVar string

type TCodMunIBGE string

type TCodUFIBGE string

type TCodUFIBGEEX string

type TCPF string

type TCPFVar string

type TData string

type TDate string

type TDateTimeUTC string

type TDec0302 string

type TDec0302Opc string

type TDec03020303 string

type TDec030204 string

type TDec0302a04 string

type TDec0302a04Max100 string

type TDec0302a04Opc string

type TDec0303 string

type TDec0303Opc string

type TDec0304 string

type TDec04020408 string

type TDec0503 string

type TDec0702 string

type TDec0803 string

type TDec0803Opc string

type TDec0804 string

type TDec0804Opc string

type TDec0906 string

type TDec0906Opc string

type TDec0908 string

type TDec0908Opc string

type TDec11001104 string

type TDec1104 string

type TDec1104Opc string

type TDec1203 string

type TDec1203Opc string

type TDec1204 string

type TDec1204Opc string

type TDec1302 string

type TDec1302Opc string

type TDec13021304 string

type TDec13021306 string

type TDec13021308 string

type TDetPagImovel struct {
	NDetPag string   `xml:"nDetPag,attr" json:"nDetPag"`
	TPag    string   `xml:"http://www.portalfiscal.inf.br/nfeabi tPag" json:"tPag"`
	XPag    *string  `xml:"http://www.portalfiscal.inf.br/nfeabi xPag,omitempty" json:"xPag,omitempty"`
	VPag    TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vPag" json:"vPag"`
}

type TDetPagIncorpLote struct {
	NDetPag  String  `xml:"nDetPag,attr" json:"nDetPag"`
	TParcela string  `xml:"http://www.portalfiscal.inf.br/nfeabi tParcela" json:"tParcela"`
	NParcela string  `xml:"http://www.portalfiscal.inf.br/nfeabi nParcela" json:"nParcela"`
	XParcela *string `xml:"http://www.portalfiscal.inf.br/nfeabi xParcela,omitempty" json:"xParcela,omitempty"`
}

type TEmail string

type TEmit struct {
	CNPJ      *TCNPJ    `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF       *TCPF     `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
	XNome     string    `xml:"http://www.portalfiscal.inf.br/nfeabi xNome" json:"xNome"`
	XFant     *string   `xml:"http://www.portalfiscal.inf.br/nfeabi xFant,omitempty" json:"xFant,omitempty"`
	EnderEmit TEndereco `xml:"http://www.portalfiscal.inf.br/nfeabi enderEmit" json:"enderEmit"`
	TpEmit    string    `xml:"http://www.portalfiscal.inf.br/nfeabi tpEmit" json:"tpEmit"`
}

type TEndereco struct {
	XLgr    string  `xml:"http://www.portalfiscal.inf.br/nfeabi xLgr" json:"xLgr"`
	Nro     string  `xml:"http://www.portalfiscal.inf.br/nfeabi nro" json:"nro"`
	XCpl    *string `xml:"http://www.portalfiscal.inf.br/nfeabi xCpl,omitempty" json:"xCpl,omitempty"`
	XBairro *string `xml:"http://www.portalfiscal.inf.br/nfeabi xBairro,omitempty" json:"xBairro,omitempty"`
	CMun    string  `xml:"http://www.portalfiscal.inf.br/nfeabi cMun" json:"cMun"`
	XMun    string  `xml:"http://www.portalfiscal.inf.br/nfeabi xMun" json:"xMun"`
	UF      string  `xml:"http://www.portalfiscal.inf.br/nfeabi UF" json:"UF"`
	CEP     *string `xml:"http://www.portalfiscal.inf.br/nfeabi CEP,omitempty" json:"CEP,omitempty"`
	CPais   *string `xml:"http://www.portalfiscal.inf.br/nfeabi cPais,omitempty" json:"cPais,omitempty"`
	XPais   *string `xml:"http://www.portalfiscal.inf.br/nfeabi xPais,omitempty" json:"xPais,omitempty"`
	Fone    *string `xml:"http://www.portalfiscal.inf.br/nfeabi fone,omitempty" json:"fone,omitempty"`
}

type TGCBS struct {
	PCBS TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pCBS" json:"pCBS"`
	GRed TGRed    `xml:"http://www.portalfiscal.inf.br/nfeabi gRed" json:"gRed"`
	VCBS TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vCBS" json:"vCBS"`
}

type TGCompraGov struct {
	TpEnteGov string     `xml:"http://www.portalfiscal.inf.br/nfeabi tpEnteGov" json:"tpEnteGov"`
	PRedutor  TDec030204 `xml:"http://www.portalfiscal.inf.br/nfeabi pRedutor" json:"pRedutor"`
	TpOperGov string     `xml:"http://www.portalfiscal.inf.br/nfeabi tpOperGov" json:"tpOperGov"`
}

type TGCont struct {
	DhCont TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfeabi dhCont" json:"dhCont"`
	XJust  string       `xml:"http://www.portalfiscal.inf.br/nfeabi xJust" json:"xJust"`
}

type TGEstornoCred struct {
	VIBSEstCred TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSEstCred" json:"vIBSEstCred"`
	VCBSEstCred TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vCBSEstCred" json:"vCBSEstCred"`
}

type TGIBSCBS struct {
	VOperacIndiv    TDec1302         `xml:"http://www.portalfiscal.inf.br/nfeabi vOperacIndiv" json:"vOperacIndiv"`
	VTornaIndiv     *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vTornaIndiv,omitempty" json:"vTornaIndiv,omitempty"`
	VRedAjusteIndiv *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vRedAjusteIndiv,omitempty" json:"vRedAjusteIndiv,omitempty"`
	VRedSocialIndiv *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi vRedSocialIndiv,omitempty" json:"vRedSocialIndiv,omitempty"`
	VBC             TDec1302         `xml:"http://www.portalfiscal.inf.br/nfeabi vBC" json:"vBC"`
	GIBSUF          TGIBSUF          `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSUF" json:"gIBSUF"`
	GIBSMun         TGIBSMun         `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSMun" json:"gIBSMun"`
	GCBS            TGCBS            `xml:"http://www.portalfiscal.inf.br/nfeabi gCBS" json:"gCBS"`
	GTribCompraGov  *TGTribCompraGov `xml:"http://www.portalfiscal.inf.br/nfeabi gTribCompraGov,omitempty" json:"gTribCompraGov,omitempty"`
}

type TGIBSMun struct {
	PIBSMun TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pIBSMun" json:"pIBSMun"`
	GRed    TGRed    `xml:"http://www.portalfiscal.inf.br/nfeabi gRed" json:"gRed"`
	VIBSMun TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSMun" json:"vIBSMun"`
}

type TGIBSUF struct {
	PIBSUF TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pIBSUF" json:"pIBSUF"`
	GRed   TGRed    `xml:"http://www.portalfiscal.inf.br/nfeabi gRed" json:"gRed"`
	VIBSUF TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSUF" json:"vIBSUF"`
}

type TGIncorpLote struct {
	IndIncorpLote string  `xml:"http://www.portalfiscal.inf.br/nfeabi indIncorpLote" json:"indIncorpLote"`
	CNPJempImob   *TCNPJ  `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJempImob,omitempty" json:"CNPJempImob,omitempty"`
	CCIBObra      *string `xml:"http://www.portalfiscal.inf.br/nfeabi cCIBObra,omitempty" json:"cCIBObra,omitempty"`
}

type TGInfTrib struct {
	GCompraGov *TGCompraGov `xml:"http://www.portalfiscal.inf.br/nfeabi gCompraGov,omitempty" json:"gCompraGov,omitempty"`
	IBSCBS     TIBSCBS      `xml:"http://www.portalfiscal.inf.br/nfeabi IBSCBS" json:"IBSCBS"`
}

type TGIntermedCorret struct {
	NCorretagem string   `xml:"nCorretagem,attr" json:"nCorretagem"`
	CNPJ        *TCNPJ   `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF         *TCPF    `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
	XNome       string   `xml:"http://www.portalfiscal.inf.br/nfeabi xNome" json:"xNome"`
	VCorretagem TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vCorretagem" json:"vCorretagem"`
}

type TGModNat struct {
	ModOper string  `xml:"http://www.portalfiscal.inf.br/nfeabi modOper" json:"modOper"`
	NatOper string  `xml:"http://www.portalfiscal.inf.br/nfeabi natOper" json:"natOper"`
	DetOper *string `xml:"http://www.portalfiscal.inf.br/nfeabi detOper,omitempty" json:"detOper,omitempty"`
}

type TGRed struct {
	PRedAliq  TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pRedAliq" json:"pRedAliq"`
	PAliqEfet TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqEfet" json:"pAliqEfet"`
}

type TGRedAjusteImovel struct {
	VInicial         *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vInicial,omitempty" json:"vInicial,omitempty"`
	VITBI            *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vITBI,omitempty" json:"vITBI,omitempty"`
	VLaudemio        *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vLaudemio,omitempty" json:"vLaudemio,omitempty"`
	VContrap         *TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vContrap,omitempty" json:"vContrap,omitempty"`
	VRedAjusteImovel TDec1302  `xml:"http://www.portalfiscal.inf.br/nfeabi vRedAjusteImovel" json:"vRedAjusteImovel"`
}

type TGTribCompraGov struct {
	PAliqIBSUF  TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqIBSUF" json:"pAliqIBSUF"`
	VTribIBSUF  TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vTribIBSUF" json:"vTribIBSUF"`
	PAliqIBSMun TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqIBSMun" json:"pAliqIBSMun"`
	VTribIBSMun TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vTribIBSMun" json:"vTribIBSMun"`
	PAliqCBS    TDec0304 `xml:"http://www.portalfiscal.inf.br/nfeabi pAliqCBS" json:"pAliqCBS"`
	VTribCBS    TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vTribCBS" json:"vTribCBS"`
}

type TIBSCBS struct {
	CST          string         `xml:"http://www.portalfiscal.inf.br/nfeabi CST" json:"CST"`
	CClassTrib   string         `xml:"http://www.portalfiscal.inf.br/nfeabi cClassTrib" json:"cClassTrib"`
	GIBSCBS      *TGIBSCBS      `xml:"http://www.portalfiscal.inf.br/nfeabi gIBSCBS,omitempty" json:"gIBSCBS,omitempty"`
	GEstornoCred *TGEstornoCred `xml:"http://www.portalfiscal.inf.br/nfeabi gEstornoCred,omitempty" json:"gEstornoCred,omitempty"`
}

type TIPv4 string

type TIde struct {
	CUF          TCodUFIBGE   `xml:"http://www.portalfiscal.inf.br/nfeabi cUF" json:"cUF"`
	CNF          string       `xml:"http://www.portalfiscal.inf.br/nfeabi cNF" json:"cNF"`
	Mod          TModNFeABI   `xml:"http://www.portalfiscal.inf.br/nfeabi mod" json:"mod"`
	Serie        TSerie       `xml:"http://www.portalfiscal.inf.br/nfeabi serie" json:"serie"`
	NNF          TNF          `xml:"http://www.portalfiscal.inf.br/nfeabi nNF" json:"nNF"`
	DhEmi        TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfeabi dhEmi" json:"dhEmi"`
	TpNF         string       `xml:"http://www.portalfiscal.inf.br/nfeabi tpNF" json:"tpNF"`
	TpImp        string       `xml:"http://www.portalfiscal.inf.br/nfeabi tpImp" json:"tpImp"`
	GModNat      TGModNat     `xml:"http://www.portalfiscal.inf.br/nfeabi gModNat" json:"gModNat"`
	TpEmis       string       `xml:"http://www.portalfiscal.inf.br/nfeabi tpEmis" json:"tpEmis"`
	NSiteAutoriz string       `xml:"http://www.portalfiscal.inf.br/nfeabi nSiteAutoriz" json:"nSiteAutoriz"`
	CDV          string       `xml:"http://www.portalfiscal.inf.br/nfeabi cDV" json:"cDV"`
	TpAmb        TAmb         `xml:"http://www.portalfiscal.inf.br/nfeabi tpAmb" json:"tpAmb"`
	FinNFe       string       `xml:"http://www.portalfiscal.inf.br/nfeabi finNFe" json:"finNFe"`
	ProcEmi      string       `xml:"http://www.portalfiscal.inf.br/nfeabi procEmi" json:"procEmi"`
	VerProc      string       `xml:"http://www.portalfiscal.inf.br/nfeabi verProc" json:"verProc"`
	GCont        *TGCont      `xml:"http://www.portalfiscal.inf.br/nfeabi gCont,omitempty" json:"gCont,omitempty"`
}

type TIe string

type TIeDest string

type TImovel struct {
	Tipo             string           `xml:"http://www.portalfiscal.inf.br/nfeabi tipo" json:"tipo"`
	Cadastro         *TCadastroImovel `xml:"http://www.portalfiscal.inf.br/nfeabi cadastro,omitempty" json:"cadastro,omitempty"`
	EspImovel        string           `xml:"http://www.portalfiscal.inf.br/nfeabi espImovel" json:"espImovel"`
	XOutrosEspecies  *string          `xml:"http://www.portalfiscal.inf.br/nfeabi xOutrosEspecies,omitempty" json:"xOutrosEspecies,omitempty"`
	Enquadramento    *string          `xml:"http://www.portalfiscal.inf.br/nfeabi enquadramento,omitempty" json:"enquadramento,omitempty"`
	EnderImovel      TEndereco        `xml:"http://www.portalfiscal.inf.br/nfeabi enderImovel" json:"enderImovel"`
	ObsLocalImovel   *string          `xml:"http://www.portalfiscal.inf.br/nfeabi obsLocalImovel,omitempty" json:"obsLocalImovel,omitempty"`
	AreaTotal        *TDec1302        `xml:"http://www.portalfiscal.inf.br/nfeabi areaTotal,omitempty" json:"areaTotal,omitempty"`
	UAreaTotal       *string          `xml:"http://www.portalfiscal.inf.br/nfeabi uAreaTotal,omitempty" json:"uAreaTotal,omitempty"`
	CartorioRegistro *string          `xml:"http://www.portalfiscal.inf.br/nfeabi cartorioRegistro,omitempty" json:"cartorioRegistro,omitempty"`
	MatricTransc     *string          `xml:"http://www.portalfiscal.inf.br/nfeabi matricTransc,omitempty" json:"matricTransc,omitempty"`
}

type TInfAdic struct {
	InfAdFisco *string `xml:"http://www.portalfiscal.inf.br/nfeabi infAdFisco,omitempty" json:"infAdFisco,omitempty"`
	InfCpl     *string `xml:"http://www.portalfiscal.inf.br/nfeabi infCpl,omitempty" json:"infCpl,omitempty"`
	ObsCont    []TObs  `xml:"http://www.portalfiscal.inf.br/nfeabi obsCont,omitempty" json:"obsCont,omitempty"`
	ObsFisco   []TObs  `xml:"http://www.portalfiscal.inf.br/nfeabi obsFisco,omitempty" json:"obsFisco,omitempty"`
}

type TInfNFeABI struct {
	Versao     TVerNFeABI    `xml:"versao,attr" json:"versao"`
	ID         ID            `xml:"Id,attr" json:"Id"`
	Ide        TIde          `xml:"http://www.portalfiscal.inf.br/nfeabi ide" json:"ide"`
	NFref      []TNFref      `xml:"http://www.portalfiscal.inf.br/nfeabi NFref,omitempty" json:"NFref,omitempty"`
	Emit       TEmit         `xml:"http://www.portalfiscal.inf.br/nfeabi emit" json:"emit"`
	Transmit   []TTransmit   `xml:"http://www.portalfiscal.inf.br/nfeabi transmit,omitempty" json:"transmit,omitempty"`
	Adquirente []TAdquirente `xml:"http://www.portalfiscal.inf.br/nfeabi adquirente,omitempty" json:"adquirente,omitempty"`
	Imovel     TImovel       `xml:"http://www.portalfiscal.inf.br/nfeabi imovel" json:"imovel"`
	AutXML     []TAutXML     `xml:"http://www.portalfiscal.inf.br/nfeabi autXML,omitempty" json:"autXML,omitempty"`
	InfOper    TInfOper      `xml:"http://www.portalfiscal.inf.br/nfeabi infOper" json:"infOper"`
	GInfTrib   TGInfTrib     `xml:"http://www.portalfiscal.inf.br/nfeabi gInfTrib" json:"gInfTrib"`
	Pag        *TPag         `xml:"http://www.portalfiscal.inf.br/nfeabi pag,omitempty" json:"pag,omitempty"`
	Total      *TTotal       `xml:"http://www.portalfiscal.inf.br/nfeabi total,omitempty" json:"total,omitempty"`
	InfAdic    *TInfAdic     `xml:"http://www.portalfiscal.inf.br/nfeabi infAdic,omitempty" json:"infAdic,omitempty"`
	InfRespTec TInfRespTec   `xml:"http://www.portalfiscal.inf.br/nfeabi infRespTec" json:"infRespTec"`
}

type TInfNFeSupl struct {
	QRCode   string `xml:"http://www.portalfiscal.inf.br/nfeabi qrCode" json:"qrCode"`
	URLChave string `xml:"http://www.portalfiscal.inf.br/nfeabi urlChave" json:"urlChave"`
}

type TInfOper struct {
	PTransImovel      *TDec0304          `xml:"http://www.portalfiscal.inf.br/nfeabi pTransImovel,omitempty" json:"pTransImovel,omitempty"`
	VTotalOperacao    TDec1302           `xml:"http://www.portalfiscal.inf.br/nfeabi vTotalOperacao" json:"vTotalOperacao"`
	IndTorna          *string            `xml:"http://www.portalfiscal.inf.br/nfeabi indTorna,omitempty" json:"indTorna,omitempty"`
	GTorna            *GTorna            `xml:"http://www.portalfiscal.inf.br/nfeabi gTorna,omitempty" json:"gTorna,omitempty"`
	GIncorpLote       *TGIncorpLote      `xml:"http://www.portalfiscal.inf.br/nfeabi gIncorpLote,omitempty" json:"gIncorpLote,omitempty"`
	GRedAjusteImovel  *TGRedAjusteImovel `xml:"http://www.portalfiscal.inf.br/nfeabi gRedAjusteImovel,omitempty" json:"gRedAjusteImovel,omitempty"`
	IndRedSocial      *string            `xml:"http://www.portalfiscal.inf.br/nfeabi indRedSocial,omitempty" json:"indRedSocial,omitempty"`
	VRedSocial        *TDec0702          `xml:"http://www.portalfiscal.inf.br/nfeabi vRedSocial,omitempty" json:"vRedSocial,omitempty"`
	Instrumento       TInstrumento       `xml:"http://www.portalfiscal.inf.br/nfeabi instrumento" json:"instrumento"`
	IndIntermedCorret *string            `xml:"http://www.portalfiscal.inf.br/nfeabi indIntermedCorret,omitempty" json:"indIntermedCorret,omitempty"`
	GIntermedCorret   []TGIntermedCorret `xml:"http://www.portalfiscal.inf.br/nfeabi gIntermedCorret,omitempty" json:"gIntermedCorret,omitempty"`
}

type TInfRespTec struct {
	CNPJ     TCNPJ  `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ" json:"CNPJ"`
	XContato string `xml:"http://www.portalfiscal.inf.br/nfeabi xContato" json:"xContato"`
	Email    TEmail `xml:"http://www.portalfiscal.inf.br/nfeabi email" json:"email"`
	Fone     string `xml:"http://www.portalfiscal.inf.br/nfeabi fone" json:"fone"`
}

type TInstrumento struct {
	TpInstrumento string  `xml:"http://www.portalfiscal.inf.br/nfeabi tpInstrumento" json:"tpInstrumento"`
	XOutrosInst   *string `xml:"http://www.portalfiscal.inf.br/nfeabi xOutrosInst,omitempty" json:"xOutrosInst,omitempty"`
	DtInstrumento TDate   `xml:"http://www.portalfiscal.inf.br/nfeabi dtInstrumento" json:"dtInstrumento"`
}

type TJust string

type TLatitude string

type TLongitude string

type TMed string

type TModNFeABI string

type TMotivo string

type TNF string

type TNFeABI struct {
	InfNFeABI  TInfNFeABI    `xml:"http://www.portalfiscal.inf.br/nfeabi infNFeABI" json:"infNFeABI"`
	InfNFeSupl TInfNFeSupl   `xml:"http://www.portalfiscal.inf.br/nfeabi infNFeSupl" json:"infNFeSupl"`
	Signature  SignatureType `xml:"http://www.w3.org/2000/09/xmldsig# Signature" json:"Signature"`
}

type TNFref struct {
	RefNFeABI TChNFeABI `xml:"http://www.portalfiscal.inf.br/nfeabi refNFeABI" json:"refNFeABI"`
}

type TNSU string

type TObs struct {
	XCampo String `xml:"xCampo,attr" json:"xCampo"`
	XTexto string `xml:"http://www.portalfiscal.inf.br/nfeabi xTexto" json:"xTexto"`
}

type TPag struct {
	DetPagImovel     []TDetPagImovel     `xml:"http://www.portalfiscal.inf.br/nfeabi detPagImovel,omitempty" json:"detPagImovel,omitempty"`
	DetPagIncorpLote []TDetPagIncorpLote `xml:"http://www.portalfiscal.inf.br/nfeabi detPagIncorpLote,omitempty" json:"detPagIncorpLote,omitempty"`
}

type TProt string

type TProtNFeABIInfProt struct {
	ID       *ID              `xml:"Id,attr,omitempty" json:"Id,omitempty"`
	TpAmb    TAmb             `xml:"http://www.portalfiscal.inf.br/nfeabi tpAmb" json:"tpAmb"`
	VerAplic string           `xml:"http://www.portalfiscal.inf.br/nfeabi verAplic" json:"verAplic"`
	ChNFeABI TChNFeABI        `xml:"http://www.portalfiscal.inf.br/nfeabi chNFeABI" json:"chNFeABI"`
	DhRecbto TDateTimeUTC     `xml:"http://www.portalfiscal.inf.br/nfeabi dhRecbto" json:"dhRecbto"`
	NProt    *string          `xml:"http://www.portalfiscal.inf.br/nfeabi nProt,omitempty" json:"nProt,omitempty"`
	DigVal   *DigestValueType `xml:"http://www.portalfiscal.inf.br/nfeabi digVal,omitempty" json:"digVal,omitempty"`
	CStat    string           `xml:"http://www.portalfiscal.inf.br/nfeabi cStat" json:"cStat"`
	XMotivo  string           `xml:"http://www.portalfiscal.inf.br/nfeabi xMotivo" json:"xMotivo"`
}

type TProtNFeABI struct {
	Versao    TVerNFeABI         `xml:"versao,attr" json:"versao"`
	InfProt   TProtNFeABIInfProt `xml:"http://www.portalfiscal.inf.br/nfeabi infProt" json:"infProt"`
	Signature *SignatureType     `xml:"http://www.w3.org/2000/09/xmldsig# Signature,omitempty" json:"Signature,omitempty"`
}

type TRec string

type TRetNFeABI struct {
	Versao     TVerNFeABI   `xml:"versao,attr" json:"versao"`
	TpAmb      TAmb         `xml:"http://www.portalfiscal.inf.br/nfeabi tpAmb" json:"tpAmb"`
	VerAplic   string       `xml:"http://www.portalfiscal.inf.br/nfeabi verAplic" json:"verAplic"`
	CStat      string       `xml:"http://www.portalfiscal.inf.br/nfeabi cStat" json:"cStat"`
	XMotivo    string       `xml:"http://www.portalfiscal.inf.br/nfeabi xMotivo" json:"xMotivo"`
	CUF        TCodUFIBGE   `xml:"http://www.portalfiscal.inf.br/nfeabi cUF" json:"cUF"`
	DhRecbto   TDateTimeUTC `xml:"http://www.portalfiscal.inf.br/nfeabi dhRecbto" json:"dhRecbto"`
	ProtNFeABI *TProtNFeABI `xml:"http://www.portalfiscal.inf.br/nfeabi protNFeABI,omitempty" json:"protNFeABI,omitempty"`
}

type TSerie string

type TServ string

type TStat string

type TString string

type TTotal struct {
	VTotalOperacao TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vTotalOperacao" json:"vTotalOperacao"`
	VOperacIndiv   TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vOperacIndiv" json:"vOperacIndiv"`
	TribTot        TTribTot `xml:"http://www.portalfiscal.inf.br/nfeabi tribTot" json:"tribTot"`
	VNF            TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vNF" json:"vNF"`
}

type TTransmit struct {
	NTransmit     string     `xml:"nTransmit,attr" json:"nTransmit"`
	IndDeclarante *string    `xml:"http://www.portalfiscal.inf.br/nfeabi indDeclarante,omitempty" json:"indDeclarante,omitempty"`
	CNPJ          *TCNPJ     `xml:"http://www.portalfiscal.inf.br/nfeabi CNPJ,omitempty" json:"CNPJ,omitempty"`
	CPF           *TCPF      `xml:"http://www.portalfiscal.inf.br/nfeabi CPF,omitempty" json:"CPF,omitempty"`
	XNome         string     `xml:"http://www.portalfiscal.inf.br/nfeabi xNome" json:"xNome"`
	EnderTransmit *TEndereco `xml:"http://www.portalfiscal.inf.br/nfeabi enderTransmit,omitempty" json:"enderTransmit,omitempty"`
	PTransIndiv   *TDec0304  `xml:"http://www.portalfiscal.inf.br/nfeabi pTransIndiv,omitempty" json:"pTransIndiv,omitempty"`
	IndContrib    *string    `xml:"http://www.portalfiscal.inf.br/nfeabi indContrib,omitempty" json:"indContrib,omitempty"`
}

type TTribTot struct {
	VRedAjusteIndiv TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vRedAjusteIndiv" json:"vRedAjusteIndiv"`
	VRedSocialIndiv TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vRedSocialIndiv" json:"vRedSocialIndiv"`
	VBC             TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vBC" json:"vBC"`
	VIBSUF          TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSUF" json:"vIBSUF"`
	VIBSMun         TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vIBSMun" json:"vIBSMun"`
	VCBS            TDec1302 `xml:"http://www.portalfiscal.inf.br/nfeabi vCBS" json:"vCBS"`
}

type TUF string

type TUFSemEX string

type TVerAplic string

type TVerNFeABI string

type Tano string

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

// NFeABI represents the global element {http://www.portalfiscal.inf.br/nfeabi}NFeABI.
type NFeABI struct {
	XMLName xml.Name `xml:"http://www.portalfiscal.inf.br/nfeabi NFeABI" json:"-"`
	TNFeABI
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
