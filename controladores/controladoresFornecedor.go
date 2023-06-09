package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func CadastrarFornecedor(c *gin.Context) {
	var fornecedorTemp struct {
		Nome      string `json:"nome"`
		Email     string `json:"email"`
		Telefone  string `json:"telefone"`
		TelefoneB string `json:"telefoneb"`
		Cpf       string `json:"cpf"`
		Cnpj      string `json:"cnpj"`
	}
	c.Bind(&fornecedorTemp)

	fornecedor := modelos.Fornecedor{
		Nome:      fornecedorTemp.Nome,
		Email:     fornecedorTemp.Email,
		Telefone:  fornecedorTemp.Telefone,
		TelefoneB: fornecedorTemp.TelefoneB,
		CPF:       fornecedorTemp.Cpf,
		CNPJ:      fornecedorTemp.Cnpj,
	}

	if err := inicializadores.BD.Create(&fornecedor).Error; err != nil {
		c.JSON(400, gin.H{"Error ao cadastrar o fornecedor ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Fornecedor cadastrado com sucesso"})
}

func DeletarFornecedor(c *gin.Context) {
	id := c.Param("id")

	if err := inicializadores.BD.Delete(&modelos.Fornecedor{}, id).Error; err != nil {
		c.JSON(400, gin.H{"Error ao excluir o fornecedor ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Fornecedor excluído com sucesso"})
}

func BuscarFornecedores(c *gin.Context) {
	var Fornecedores []modelos.Fornecedor

	if err := inicializadores.BD.Find(&Fornecedores).Error; err != nil {
		c.JSON(400, gin.H{"Error ao buscar os fornecedores ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"fornecedores": Fornecedores})
}

func BuscarFornecedor(c *gin.Context) {
	id := c.Param("id")

	var fornecedor modelos.Fornecedor

	if err := inicializadores.BD.First(&fornecedor, id).Error; err != nil {
		c.JSON(400, gin.H{"Error ao buscar o fornecedor ": err.Error()})
		return
	}

	c.JSON(200, gin.H{"fornecedor": fornecedor})
}

func AtualizarFornecedor(c *gin.Context) {
	id := c.Param("id")

	var fornecedorTemp struct {
		Nome      string `json:"nome"`
		Email     string `json:"email"`
		Telefone  string `json:"telefone"`
		TelefoneB string `json:"telefoneb"`
		CPF       string `json:"cpf"`
		CNPJ      string `json:"cnpj"`
	}

	c.Bind(&fornecedorTemp)

	var fornecedor modelos.Fornecedor

	if err := inicializadores.BD.First(&fornecedor, id).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao bucar o fornecedor": err.Error()})
		return
	}

	if err := inicializadores.BD.Model(&fornecedor).Updates(modelos.Fornecedor{
		Nome:      fornecedorTemp.Nome,
		Email:     fornecedorTemp.Email,
		Telefone:  fornecedorTemp.Telefone,
		TelefoneB: fornecedorTemp.TelefoneB,
		CPF:       fornecedorTemp.CPF,
		CNPJ:      fornecedorTemp.CNPJ,
	}).Error; err != nil {
		c.JSON(400, gin.H{"Erro ao atualizar o fornecedor": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Fornecedor atualizado com sucesso"})
}
