package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/trsouza/faculdade-pd/controllers"

	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")

	{
		aluno := main.Group("/aluno")
		{
			aluno.GET("/:id", controllers.BuscarAluno)
			aluno.GET("/", controllers.BuscarAlunos)
			aluno.GET("/curso/:id", controllers.FiltrarAlunosPorCurso)
			aluno.GET("/disciplina/:id", controllers.FiltrarAlunosPorDisciplina)
			aluno.GET("/faculdade/:id", controllers.FiltrarAlunosPorFaculdade)
			aluno.GET("/total/:id", controllers.FiltrarTotalAlunosPorFaculdade)
			aluno.POST("/", controllers.CriarAluno)
			aluno.PUT("/:id", controllers.EditarAluno)
			aluno.DELETE("/:id", controllers.DeletarAluno)
		}

		faculdade := main.Group("/faculdade")
		{
			faculdade.GET("/:id", controllers.BuscarFaculdade)
			faculdade.GET("/", controllers.BuscarFaculdades)
			faculdade.POST("/", controllers.CriarFaculdade)
			faculdade.PUT("/:id", controllers.EditarFaculdade)
		 	faculdade.DELETE("/:id", controllers.DeletarFaculdade)
		}

		professor := main.Group("/professor")
		{
			professor.GET("/:id", controllers.BuscarProfessor)
			professor.GET("/", controllers.BuscarProfessores)
			professor.POST("/", controllers.CriarProfessor)
			professor.PUT("/:id", controllers.EditarProfessor)
			professor.DELETE("/:id", controllers.DeletarProfessor)
		}

		curso := main.Group("/curso")
		{
			curso.GET("/:id", controllers.BuscarCurso)
			curso.GET("/", controllers.BuscarCursos)
			curso.POST("/", controllers.CriarCurso)
			curso.PUT("/:id", controllers.EditarCurso)
			curso.DELETE("/:id", controllers.DeletarCurso)
		}

		disciplina := main.Group("/disciplina")
		{
			disciplina.GET("/:id", controllers.BuscarDisciplina)
			disciplina.GET("/", controllers.BuscarDisciplinas)
			disciplina.GET("/aluno/:id", controllers.FiltrarDisciplinasPorAluno)
			disciplina.POST("/", controllers.CriarDisciplina)
			disciplina.PUT("/:id", controllers.EditarDisciplina)
			disciplina.DELETE("/:id", controllers.DeletarDisciplina)
		}

		disciplinaMatricula := main.Group("/disciplina-matricula")
		{
			disciplinaMatricula.GET("/:id", controllers.BuscarDisciplinaMatricula)
			disciplinaMatricula.GET("/", controllers.BuscarDisciplinaMatriculas)
			disciplinaMatricula.POST("/", controllers.CriarDisciplinaMatricula)
			disciplinaMatricula.PUT("/:id", controllers.EditarDisciplinaMatricula)
			disciplinaMatricula.DELETE("/:id", controllers.DeletarDisciplinaMatricula)
		}

		cursoDisciplina := main.Group("/curso-disciplina")
		{
			cursoDisciplina.GET("/:id", controllers.BuscarCursoDisciplina)
			cursoDisciplina.GET("/", controllers.BuscarCursoDisciplinas)
			cursoDisciplina.POST("/", controllers.CriarCursoDisciplina)
			cursoDisciplina.PUT("/:id", controllers.EditarCursoDisciplina)
			cursoDisciplina.DELETE("/:id", controllers.DeletarCursoDisciplina)
		}
	}

	router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	return router
}


