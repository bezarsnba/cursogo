package operadora

import "strconv"
import "../prefixo"

// NomeOperadora representa o nome da operadora
var NomeOperadora = "XPTO"

//PrefixoDaCapitalOperadora + o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora