package busca

import (
	"CNPJ/controllers"

	"github.com/Nhanderu/brdoc"
	"github.com/gin-gonic/gin"
)

func CNPJ(c *gin.Context) {

	cnpj := c.Query("cnpj")

	if erro := brdoc.IsCNPJ(cnpj); !erro {
		c.JSON(200, gin.H{"data": "O CNPJ é inválido"})
		return
	}

	if len(cnpj) != 14 {
		c.AbortWithStatusJSON(400, gin.H{"Error": "CNPJ é inválido"})
		return
	}

	data := controllers.GetCNPJ(cnpj)
	if data == nil {
		c.AbortWithStatusJSON(400, gin.H{"Error": "CNPJ não encontrado"})
		return
	}

	c.JSON(200, data)

}
