package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func CadastrarUsuario(c *gin.Context) {
	var usuarioTemp struct {
		Nome       string
		Telefone   string
		TelefoneB  string
		Email      string
		CPF        string
		Privilegio string
	}

	c.Bind(&usuarioTemp)

	usuario := modelos.Usuario{
		Nome:       usuarioTemp.Nome,
		Telefone:   usuarioTemp.Telefone,
		TelefoneB:  usuarioTemp.TelefoneB,
		Email:      usuarioTemp.Email,
		CPF:        usuarioTemp.CPF,
		Privilegio: usuarioTemp.Privilegio,
	}

	result := inicializadores.BD.Create(&usuario)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"usuarioz": usuario,
	})
}

func DeletarUsuario(c *gin.Context) {
	id := c.Param("id")

	inicializadores.BD.Delete(&modelos.Usuario{}, id)

	c.Status(200)
}

func BuscarUsuarios(c *gin.Context) {
	var usuarios []modelos.Usuario
	inicializadores.BD.Find(&usuarios)

	c.JSON(200, gin.H{
		"usuarios": usuarios,
	})
}

func BuscarUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuario modelos.Usuario
	inicializadores.BD.First(&usuario, id)

	c.JSON(200, gin.H{
		"usuario": usuario,
	})
}

func AtualizarUsuario(c *gin.Context) {
	id := c.Param("id")

	var usuarioTempAtt struct {
		Nome       string
		Telefone   string
		TelefoneB  string
		Email      string
		CPF        string
		Privilegio string
	}

	c.Bind(&usuarioTempAtt)

	var usuarioTemp modelos.Usuario
	inicializadores.BD.First(&usuarioTemp, id)

	inicializadores.BD.Model(&usuarioTemp).Updates(modelos.Usuario{
		Nome:       usuarioTemp.Nome,
		Telefone:   usuarioTemp.Telefone,
		TelefoneB:  usuarioTemp.TelefoneB,
		Email:      usuarioTemp.Email,
		CPF:        usuarioTemp.CPF,
		Privilegio: usuarioTemp.Privilegio,
	})

	c.JSON(200, gin.H{
		"post": usuarioTemp,
	})
}
