# Fiscal Tributário Brasil — schemas e serviços

Os pacotes `bindings/` são gerados dos XSDs capturados, incluindo os tipos anônimos e dependências. Valores numéricos e identificadores preservam a representação textual do XML.

Requer `xmllint` (libxml2) para validação XSD. A validação estrutural não verifica autorização, assinatura criptográfica nem todas as regras tributárias. Conteúdo misto ou escolhas repetidas intercaladas usam `InnerXML` para preservar a ordem, em vez de redistribuir elementos em listas e perder a sequência original.

| Família | Modelo | Schema | Importação Go |
| --- | --- | --- | --- |
| NFe | 55 | PL_010g_v1.03/Evento/e110001_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfe/e110001_v1_00` |
| NFe | 55 | PL_010g_v1.03/Evento/e110300_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfe/e110300_v1_00` |
| NFe | 55 | PL_010g_v1.03/NFe/nfe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfe/nfe_v4_00` |
| NFCe | 65 | PL_010g_v1.03/Evento/e110001_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfce/e110001_v1_00` |
| NFCe | 65 | PL_010g_v1.03/Evento/e110300_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfce/e110300_v1_00` |
| NFCe | 65 | PL_010g_v1.03/NFe/nfe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfce/nfe_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/GTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/gtve_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/consSitCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/conssitcte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/consStatServCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/consstatservcte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalAereo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctemodalaereo_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalAquaviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctemodalaquaviario_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalDutoviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctemodaldutoviario_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalFerroviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctemodalferroviario_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviarioOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctemodalrodoviarioos_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctemodalrodoviario_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteMultiModal_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctemultimodal_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/cteos_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cteSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/ctesimp_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/cte_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/cte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evApropriaCredPres_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evapropriacredpres_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evCCeCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evccecte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evCECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evcecte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evCancCECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evcanccecte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evCancCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evcanccte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evCancIECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evcanciecte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evCancPrestDesacordo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evcancprestdesacordo_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evCancVincPgto_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evcancvincpgto_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evEPECCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evepeccte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evGTV_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evgtv_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evIECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/eviecte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evPrestDesacordo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evprestdesacordo_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evRegMultimodal_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evregmultimodal_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/evVincPgto_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/evvincpgto_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/eventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/eventocte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/procCTeOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/proccteos_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/procCTeSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/procctesimp_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/procCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/proccte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/procEventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/proceventocte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/procGTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/procgtve_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/retCTeOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/retcteos_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/retCTeSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/retctesimp_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/retCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/retcte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/retConsSitCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/retconssitcte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/retConsStatServCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/retconsstatservcte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/retEventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/reteventocte_v4_00` |
| CTe | 57 | PL_CTe_400_NT2026.004 RTC_1.00/retGTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cte/retgtve_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/GTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/gtve_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/consSitCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/conssitcte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/consStatServCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/consstatservcte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalAereo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctemodalaereo_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalAquaviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctemodalaquaviario_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalDutoviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctemodaldutoviario_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalFerroviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctemodalferroviario_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviarioOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctemodalrodoviarioos_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctemodalrodoviario_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteMultiModal_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctemultimodal_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/cteos_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cteSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/ctesimp_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/cte_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/cte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evApropriaCredPres_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evapropriacredpres_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evCCeCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evccecte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evCECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evcecte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evCancCECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evcanccecte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evCancCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evcanccte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evCancIECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evcanciecte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evCancPrestDesacordo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evcancprestdesacordo_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evCancVincPgto_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evcancvincpgto_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evEPECCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evepeccte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evGTV_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evgtv_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evIECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/eviecte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evPrestDesacordo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evprestdesacordo_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evRegMultimodal_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evregmultimodal_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/evVincPgto_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/evvincpgto_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/eventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/eventocte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/procCTeOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/proccteos_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/procCTeSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/procctesimp_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/procCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/proccte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/procEventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/proceventocte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/procGTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/procgtve_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/retCTeOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/retcteos_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/retCTeSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/retctesimp_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/retCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/retcte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/retConsSitCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/retconssitcte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/retConsStatServCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/retconsstatservcte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/retEventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/reteventocte_v4_00` |
| CTe | 67 | PL_CTe_400_NT2026.004 RTC_1.00/retGTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/cteos/retgtve_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/GTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/gtve_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/consSitCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/conssitcte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/consStatServCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/consstatservcte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalAereo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctemodalaereo_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalAquaviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctemodalaquaviario_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalDutoviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctemodaldutoviario_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalFerroviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctemodalferroviario_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviarioOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctemodalrodoviarioos_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteModalRodoviario_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctemodalrodoviario_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteMultiModal_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctemultimodal_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/cteos_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cteSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/ctesimp_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/cte_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/cte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evApropriaCredPres_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evapropriacredpres_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evCCeCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evccecte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evCECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evcecte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evCancCECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evcanccecte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evCancCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evcanccte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evCancIECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evcanciecte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evCancPrestDesacordo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evcancprestdesacordo_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evCancVincPgto_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evcancvincpgto_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evEPECCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evepeccte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evGTV_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evgtv_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evIECTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/eviecte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evPrestDesacordo_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evprestdesacordo_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evRegMultimodal_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evregmultimodal_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/evVincPgto_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/evvincpgto_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/eventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/eventocte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/procCTeOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/proccteos_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/procCTeSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/procctesimp_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/procCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/proccte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/procEventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/proceventocte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/procGTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/procgtve_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/retCTeOS_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/retcteos_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/retCTeSimp_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/retctesimp_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/retCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/retcte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/retConsSitCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/retconssitcte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/retConsStatServCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/retconsstatservcte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/retEventoCTe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/reteventocte_v4_00` |
| GTVe | 64 | PL_CTe_400_NT2026.004 RTC_1.00/retGTVe_v4.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/gtve/retgtve_v4_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/consMDFeNaoEnc_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/consmdfenaoenc_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/consReciMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/consrecimdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/consSitMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/conssitmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/consStatServMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/consstatservmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/distMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/distmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/enviMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/envimdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/evAlteracaoPagtoServMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/evalteracaopagtoservmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/evCancMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/evcancmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/evConfirmaServMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/evconfirmaservmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/evEncMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/evencmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/evIncCondutorMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/evinccondutormdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/evInclusaoDFeMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/evinclusaodfemdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/evPagtoOperMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/evpagtoopermdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/eventoMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/eventomdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/mdfeConsultaDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/mdfeconsultadfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/mdfeModalAereo_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/mdfemodalaereo_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/mdfeModalAquaviario_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/mdfemodalaquaviario_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/mdfeModalFerroviario_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/mdfemodalferroviario_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/mdfeModalRodoviario_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/mdfemodalrodoviario_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/mdfe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/mdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/procEventoMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/proceventomdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/procMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/procmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retConsMDFeNaoEnc_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retconsmdfenaoenc_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retConsReciMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retconsrecimdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retConsSitMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retconssitmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retConsStatServMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retconsstatservmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retDistMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retdistmdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retEnviMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retenvimdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retEventoMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/reteventomdfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retMDFeConsultaDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retmdfeconsultadfe_v3_00` |
| MDFe | 58 | PL_MDFe_300b_NT012025_1.05/retMDFe_v3.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/mdfe/retmdfe_v3_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/bpeTA_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/bpeta_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/bpeTM_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/bpetm_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/bpe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/bpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/consSitBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/conssitbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/consStatServBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/consstatservbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/evAlteracaoPoltrona_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/evalteracaopoltrona_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/evCancBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/evcancbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/evCancVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/evcancvincpgto_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/evExcessoBagagem_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/evexcessobagagem_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/evNaoEmbBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/evnaoembbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/evVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/evvincpgto_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/eventoBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/eventobpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/procBPeTM_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/procbpetm_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/procBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/procbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/procEventoBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/proceventobpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/retBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/retbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/retConsSitBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/retconssitbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/retConsStatServBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/retconsstatservbpe_v1_00` |
| BPe | 63 | PL_BPe_100b_NT2026.004 RTC_1.00/retEventoBPe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/bpe/reteventobpe_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/consSitNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/conssitnfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/consStatServNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/consstatservnfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/evCancNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/evcancnfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/evCancVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/evcancvincpgto_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/evVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/evvincpgto_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/eventoNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/eventonfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/nfcom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/nfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/procEventoNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/proceventonfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/procNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/procnfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/retConsSitNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/retconssitnfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/retConsStatServNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/retconsstatservnfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/retEventoNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/reteventonfcom_v1_00` |
| NFCom | 62 | PL_NFCOM_1.00_NT2026.004 RTC_1.00/retNFCom_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfcom/retnfcom_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/consReciNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/consrecinf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/consSitNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/conssitnf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/consStatServNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/consstatservnf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/evCancNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/evcancnf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/evCancVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/evcancvincpgto_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/evVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/evvincpgto_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/eventoNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/eventonf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/nf3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/nf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/procEventoNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/proceventonf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/procNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/procnf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/retConsReciNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/retconsrecinf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/retConsSitNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/retconssitnf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/retConsStatServNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/retconsstatservnf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/retEventoNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/reteventonf3e_v1_00` |
| NF3e | 66 | PL_NF3E_1.00a_NT2026.004 RTC_1.00/retNF3e_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nf3e/retnf3e_v1_00` |
| NFSe | NFS | Schemas/1.00/CNC_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/schemas_1_00_cnc_v1_00` |
| NFSe | NFS | Schemas/1.00/DPS_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/dps_v1_00` |
| NFSe | NFS | Schemas/1.00/NFSe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/nfse_v1_00` |
| NFSe | NFS | Schemas/1.00/evento_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/evento_v1_00` |
| NFSe | NFS | Schemas/1.00/pedRegEvento_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/pedregevento_v1_00` |
| NFSe | NFS | Schemas/1.01/CNC_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/schemas_1_01_cnc_v1_00` |
| NFSe | NFS | Schemas/1.01/DPS_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/dps_v1_01` |
| NFSe | NFS | Schemas/1.01/NFSe_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/nfse_v1_01` |
| NFSe | NFS | Schemas/1.01/evento_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/evento_v1_01` |
| NFSe | NFS | Schemas/1.01/pedRegEvento_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfse/pedregevento_v1_01` |
| DCe | 99 | xsd/pl100/consulta/consSitDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/conssitdce_v1_00` |
| DCe | 99 | xsd/pl100/consulta/consStatServDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/consstatservdce_v1_00` |
| DCe | 99 | xsd/pl100/consulta/retConsSitDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/retconssitdce_v1_00` |
| DCe | 99 | xsd/pl100/consulta/retConsStatServDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/retconsstatservdce_v1_00` |
| DCe | 99 | xsd/pl100/dce/dce_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/dce_v1_00` |
| DCe | 99 | xsd/pl100/dce/procDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/procdce_v1_00` |
| DCe | 99 | xsd/pl100/dce/retDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/retdce_v1_00` |
| DCe | 99 | xsd/pl100/evento/evCancDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/evcancdce_v1_00` |
| DCe | 99 | xsd/pl100/evento/eventoDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/eventodce_v1_00` |
| DCe | 99 | xsd/pl100/evento/procEventoDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/proceventodce_v1_00` |
| DCe | 99 | xsd/pl100/evento/retEventoDCe_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dce/reteventodce_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/consSitNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/conssitnfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/consStatServNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/consstatservnfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/evCancNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/evcancnfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/evCancVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/evcancvincpgto_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/evVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/evvincpgto_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/eventoNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/eventonfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/nfag_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/nfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/procEventoNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/proceventonfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/procNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/procnfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/retConsSitNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/retconssitnfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/retConsStatServNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/retconsstatservnfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/retEventoNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/reteventonfag_v1_00` |
| NFAg | 75 | PL_NFAg_NT2026.004 RTC_1.00/retNFAg_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfag/retnfag_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/consSitNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/conssitnfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/consStatServNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/consstatservnfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/evCancNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/evcancnfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/evCancVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/evcancvincpgto_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/evVincPgto_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/evvincpgto_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/eventoNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/eventonfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/nfgas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/nfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/procEventoNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/proceventonfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/procNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/procnfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/retConsSitNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/retconssitnfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/retConsStatServNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/retconsstatservnfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/retEventoNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/reteventonfgas_v1_00` |
| NFGas | 76 | PL_NFGas_NT2026.004 RTC_1.00/retNFGas_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfgas/retnfgas_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/NFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/nfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/consSitNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/conssitnfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/consStatServNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/consstatservnfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/evApropCredIndNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/evapropcredindnfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/evCancNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/evcancnfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/evPagParcelaNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/evpagparcelanfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/eventoNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/eventonfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/procEventoNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/proceventonfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/procNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/procnfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/retConsSitNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/retconssitnfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/retConsStatServNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/retconsstatservnfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/retEventoNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/reteventonfeabi_v1_00` |
| NFeABI | 77 | PL_NFeABI_1.00/retNFeABI_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeabi/retnfeabi_v1_00` |
| NFSeVia | NFSVia | EventoVia_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfsevia/eventovia_v1_01` |
| NFSeVia | NFSVia | NFSeVia_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfsevia/nfsevia_v1_01` |
| NFSeVia | NFSVia | ProcEventoVia_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfsevia/proceventovia_v1_01` |
| NFSeVia | NFSVia | ProcNFSeVia_v1.01.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfsevia/procnfsevia_v1_01` |
| NFeOuro | Ouro | DAO_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeouro/dao_v1_00` |
| NFeOuro | Ouro | NFeOuro_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/nfeouro/nfeouro_v1_00` |
| DAO | DAO | DAO_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dao/dao_v1_00` |
| DAO | DAO | NFeOuro_v1.00.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dao/nfeouro_v1_00` |
| DeRE | DeRE | envioLoteDere-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/enviolotedere_v1_0_1` |
| DeRE | DeRE | evtAntecReceb-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtantecreceb_v0_0_2` |
| DeRE | DeRE | evtAplicResTec-v1_0_0.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtaplicrestec_v1_0_0` |
| DeRE | DeRE | evtArranjoCredDest-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtarranjocreddest_v0_0_2` |
| DeRE | DeRE | evtArranjoPartic-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtarranjopartic_v0_0_2` |
| DeRE | DeRE | evtArrendMerc-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtarrendmerc_v0_0_2` |
| DeRE | DeRE | evtBalancete-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtbalancete_v1_0_1` |
| DeRE | DeRE | evtDebOpOfPublic-v0_0_3.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtdebopofpublic_v0_0_3` |
| DeRE | DeRE | evtFechMensal-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtfechmensal_v0_0_2` |
| DeRE | DeRE | evtIdApostPrem-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtidapostprem_v0_0_2` |
| DeRE | DeRE | evtInfoContrib-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtinfocontrib_v1_0_1` |
| DeRE | DeRE | evtInfoCotistas-v0_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtinfocotistas_v0_0_1` |
| DeRE | DeRE | evtOperFinanc-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtoperfinanc_v0_0_2` |
| DeRE | DeRE | evtPGCC-v1_0_4.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtpgcc_v1_0_4` |
| DeRE | DeRE | evtPlAssistSaude-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtplassistsaude_v0_0_2` |
| DeRE | DeRE | evtReabertMensal-v0_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtreabertmensal_v0_0_1` |
| DeRE | DeRE | evtRelDeducoes-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtreldeducoes_v0_0_2` |
| DeRE | DeRE | evtRetornoAplicFin-v1_0_0.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornoaplicfin_v1_0_0` |
| DeRE | DeRE | evtRetornoBalan-v1_0_0.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornobalan_v1_0_0` |
| DeRE | DeRE | evtRetornoMensal-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornomensal_v0_0_2` |
| DeRE | DeRE | evtRetornoPMensal-v0_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornopmensal_v0_0_1` |
| DeRE | DeRE | evtRetornoRDed-v0_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornorded_v0_0_1` |
| DeRE | DeRE | evtRetornoReabert-v0_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornoreabert_v0_0_1` |
| DeRE | DeRE | evtRetornoTabela-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornotabela_v1_0_1` |
| DeRE | DeRE | evtRetornoTitPub-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornotitpub_v0_0_2` |
| DeRE | DeRE | evtRetornoTransac-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtretornotransac_v0_0_2` |
| DeRE | DeRE | evtSegElemSCob-v0_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtsegelemscob_v0_0_1` |
| DeRE | DeRE | evtSegPrevCap-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtsegprevcap_v0_0_2` |
| DeRE | DeRE | evtServRemPreco-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtservrempreco_v0_0_2` |
| DeRE | DeRE | evtServRemTarifa-v0_0_2.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/evtservremtarifa_v0_0_2` |
| DeRE | DeRE | retornoLoteDere-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_previa/retornolotedere_v1_0_1` |
| DeRE | DeRE | envioLoteDere-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_producao/enviolotedere_v1_0_1` |
| DeRE | DeRE | evtInfoContrib-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_producao/evtinfocontrib_v1_0_1` |
| DeRE | DeRE | evtPGCC-v1_0_3.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_producao/evtpgcc_v1_0_3` |
| DeRE | DeRE | evtRetornoTabela-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_producao/evtretornotabela_v1_0_1` |
| DeRE | DeRE | retornoLoteDere-v1_0_1.xsd | `github.com/GaabsDev/fiscal-tributario-brasil/bindings/dere_producao/retornolotedere_v1_0_1` |

## Uso

Escolha o pacote do serviço na tabela. `Parse(ctx, dadosXML, &documento)` valida o XML no schema antes de decodificar. `Marshal(ctx, documento)` serializa e valida antes de devolver o XML. Os tipos raiz disponíveis constam em `schema-profiles-manifest.json`.

Os arquivos `testdata/structural_*.xml` são dados sintéticos para testar a estrutura e a serialização. Não devem ser usados como documentos fiscais de produção. Extensões `xs:any` com validação `skip` ou `lax` seguem essa regra do schema; modais e eventos devem ser validados também pelo serviço específico correspondente.
