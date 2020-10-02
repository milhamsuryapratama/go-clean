package handler

import (
	"go-clean/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

// SiswaHandler ...
type SiswaHandler struct {
	SiswaEntity domain.SiswaEntity
}

// SiswaHandlerFunc ...
func SiswaHandlerFunc(r *gin.RouterGroup, us domain.SiswaEntity) {
	handler := &SiswaHandler{
		SiswaEntity: us,
	}

	r.GET("/siswa", handler.GetSiswa)
	r.GET("/siswa/:id", handler.ShowSiswa)
	r.POST("/siswa", handler.CreateSiswa)
	r.PUT("/siswa/:id", handler.UpdateSiswa)
}

// GetSiswa ...
func (a *SiswaHandler) GetSiswa(c *gin.Context) {
	siswa, _ := a.SiswaEntity.Get(c.Request.Context())
	c.JSON(200, gin.H{"data": siswa})
}

// ShowSiswa ...
func (a *SiswaHandler) ShowSiswa(c *gin.Context) {
	siswa, _ := a.SiswaEntity.Show(c.Param("id"))
	c.JSON(200, gin.H{"data": siswa})
}

// CreateSiswa ...
func (a *SiswaHandler) CreateSiswa(c *gin.Context) {
	kelasID, _ := strconv.Atoi(c.PostForm("kelas_id"))
	s := domain.Siswa{
		Nama:    c.PostForm("nama"),
		Jk:      c.PostForm("jk"),
		Alamat:  c.PostForm("alamat"),
		KelasID: kelasID,
	}
	siswa, err := a.SiswaEntity.Create(s)
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	c.JSON(201, gin.H{"data": siswa})
}

// UpdateSiswa ..
func (a *SiswaHandler) UpdateSiswa(c *gin.Context) {
	kelasID, _ := strconv.Atoi(c.PostForm("kelas_id"))
	s := domain.Siswa{
		Nama:    c.PostForm("nama"),
		Jk:      c.PostForm("jk"),
		Alamat:  c.PostForm("alamat"),
		KelasID: kelasID,
	}
	siswa, err := a.SiswaEntity.Update(s, c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": siswa,
	})
}
