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
	r.GET("/kelas/:id", handler.ShowKelas)
	r.DELETE("/kelas/:id", handler.DeleteKelas)
	r.PUT("/kelas/:id", handler.UpdateKelas)
}

// GetKelas ...
func (a *KelasHandler) GetKelas(c *gin.Context) {
	kelas, _ := a.KelasEntity.Get(c.Request.Context())
	c.JSON(200, gin.H{"data": kelas})
}

// CreateKelas ...
func (a *KelasHandler) CreateKelas(c *gin.Context) {
	k := domain.Kelas{
		NamaKelas: c.PostForm("nama_kelas"),
	}
	kelas, err := a.KelasEntity.Create(k)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(201, gin.H{"data": kelas})
}

// ShowKelas ...
func (a *KelasHandler) ShowKelas(c *gin.Context) {
	kelas, _ := a.KelasEntity.Show(c.Param("id"))
	c.JSON(200, gin.H{"data": kelas})
}

// DeleteKelas ...
func (a *KelasHandler) DeleteKelas(c *gin.Context) {
	_, err := a.KelasEntity.Delete(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Success",
	})
}

// UpdateKelas ...
func (a *KelasHandler) UpdateKelas(c *gin.Context) {
	k := domain.Kelas{
		NamaKelas: c.PostForm("nama_kelas"),
	}
	kelas, err := a.KelasEntity.Update(k, c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": kelas,
	})
}
