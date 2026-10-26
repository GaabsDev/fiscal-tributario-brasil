package protocolos

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/GaabsDev/fiscal-tributario-brasil/nfe"
)

func TestProcNFe_Serialization(t *testing.T) {
	proc := ProcNFe{
		Versao: "4.00",
		NFe: nfe.NFe{
			InfNFe: nfe.InfNFe{
				Versao: "4.00",
				Id:     "NFe35260112345678000195550010000000011000000010",
			},
		},
		ProtNFe: ProtNFe{
			Versao: "4.00",
			InfProt: InfProtNFe{
				TpAmb:    "1",
				VerAplic: "SP_NFE_PL009_V4",
				ChNFe:    "35260112345678000195550010000000011000000010",
				NProt:    "135260000000001",
				DigVal:   "abcde12345=",
				CStat:    "100",
				XMotivo:  "Autorizado o uso da NF-e",
			},
		},
	}

	xmlBytes, err := xml.MarshalIndent(proc, "", "  ")
	if err != nil {
		t.Fatalf("falha ao serializar nfeProc: %v", err)
	}
	xmlStr := string(xmlBytes)
	if !strings.Contains(xmlStr, "<nfeProc") || !strings.Contains(xmlStr, "<nProt>135260000000001</nProt>") {
		t.Errorf("XML do nfeProc inválido: %s", xmlStr)
	}
}
