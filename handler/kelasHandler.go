package handler

import (
	"go-clean/domain"

	"github.com/gin-gonic/gin"
)

// KelasHandler ...
type KelasHandler struct {
	KelasEntity domain.KelasEntity
}

// KelasHandlerFunc ...
func KelasHandlerFunc(r *gin.RouterGroup, us domain.KelasEntity) {
	handler := &KelasHandler{
		KelasEntity: us,
	}

	r.GET("/kelas", handler.GetKelas)
	r.POST("/kelas", handler.CreateKelas)
}

// GetKelas ...
func (a *KelasHandler) GetKelas(c *gin.Context) {
	kelas, _ := a.KelasEntity.Get(c.Request.Context())
	c.JSON(200, gin.H{"data": kelas})
}

// CreateKelas ...
func (a *KelasHandler) CreateKelas(c *gin.Context) {
	k := domain.Kelas{
		ID:        24,
		NamaKelas: c.PostForm("nama_kelas"),
	}
	kelas, err := a.KelasEntity.Create(k)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(201, gin.H{"data": kelas})
}
