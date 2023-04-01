package main

import (
	"fmt"
	"projeto/controladores"
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Funcionando")
	inicializadores.CarregarVariaveisDeAmbiente()
	inicializadores.ConectarAoBD()

	inicializadores.BD.AutoMigrate(&modelos.Usuario{})
	inicializadores.BD.AutoMigrate(&modelos.Senhas{})
	inicializadores.BD.AutoMigrate(&modelos.Endereco{})
	inicializadores.BD.AutoMigrate(&modelos.Fornecedor{})
	inicializadores.BD.AutoMigrate(&modelos.Produto{})
	inicializadores.BD.AutoMigrate(&modelos.Carrinho{})
	inicializadores.BD.AutoMigrate(&modelos.Items_Carrinho{})
	inicializadores.BD.AutoMigrate(&modelos.Pedido{})
	inicializadores.BD.AutoMigrate(&modelos.Favoritos{})
}

func main() {
	r := gin.Default()

	r.POST("/usuario", controladores.CadastrarUsuario)
	r.GET("/usuario/:id", controladores.BuscarUsuario)
	r.GET("/usuario", controladores.BuscarUsuarios)
	r.PUT("/usuario/:id", controladores.AtualizarUsuario)
	r.PUT("/atualizarsenha/:id", controladores.AtualizarSenhaUsuario)
	r.DELETE("/usuario/:id", controladores.DeletarUsuario)
	r.PUT("/cadastrarTelefone/:id", controladores.AtualizarTelefone)
	r.POST("/cadastrarEndereco", controladores.CadastrarEndereco)

	r.POST("/usuario/login", controladores.Login)

	r.POST("/fornecedor", controladores.CadastrarFornecedor)
	r.GET("/fornecedor", controladores.BuscarFornecedores)
	r.GET("/fornecedor/:id", controladores.BuscarFornecedor)

	r.Run()
}
