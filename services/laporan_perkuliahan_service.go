package services

import (
	"bytes"
	"errors"
	"path/filepath"
	"strconv"
	"time"

	"html/template"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"skripsi.id/obfuss/entities"
	"skripsi.id/obfuss/middlewares"
	"skripsi.id/obfuss/models"
	"skripsi.id/obfuss/queries"
	"skripsi.id/obfuss/utils"
)

type LaporanPerkuliahanService struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewLaporanPerkuliahanService(db *gorm.DB, logger *zap.Logger) *LaporanPerkuliahanService {
	return &LaporanPerkuliahanService{
		db,
		logger.With(zap.String("type", "LaporanPerkuliahanService")),
	}
}

func (s *LaporanPerkuliahanService) SaveTrans(payload *models.SaveTransLaporanPerkuliahan, user *middlewares.User) error {
	var (
		internalServer = errors.New("internal server error")
	)

	var angketDosen *entities.AngketDosen

	err := s.db.Model(&angketDosen).First(&angketDosen, map[string]any{
		"id_pertemuan": &payload.IDPertemuan,
	}).Error

	if err != nil {
		s.logger.Error("error when find angket dosen :", zap.Any("error", err))

		return internalServer
	}

	currentTime := time.Now()

	angketDosen.ModifiedDate = &currentTime
	angketDosen.ModifiedUser = &user.ActorID
	angketDosen.IDDosen = payload.IDDosen
	angketDosen.PertemuanKe = &payload.PertemuanKe
	angketDosen.JamKuliah = &payload.JamKuliah
	angketDosen.Keterangan = &payload.Keterangan
	angketDosen.RingkasanMateri = &payload.RingkasanMateri

	if payload.FileName != "" {
		angketDosen.FileGambar = &payload.FileName
	}

	if err := s.db.Model(&angketDosen).Save(angketDosen).Error; err != nil {
		s.logger.Error("error when update angket dosen :", zap.Any("error", err))

		return internalServer
	}

	return nil
}

func (s *LaporanPerkuliahanService) GetAngketDosenByPertemuan(idPertemuan int) (*entities.AngketDosen, error) {
	var angketDosen *entities.AngketDosen

	err := s.db.Model(&angketDosen).First(&angketDosen).Error

	if err != nil {
		s.logger.Error("error when find angket dosen :", zap.Any("error", err))
		return nil, errors.New("internal server error")
	}

	return angketDosen, nil
}

func (s *LaporanPerkuliahanService) GetAllAngeketDosen(
	pageParams models.PaginationParams,
	user *middlewares.User,
) (
	[]map[string]any,
	*models.PaginationMeta,
	error,
) {
	var kelasMhs []map[string]interface{}

	query := queries.GetKelasIDFromTPK(s.db, user)
	query.Find(&kelasMhs).Limit(1)
	idKelas := kelasMhs[0]["id_kelas"].(int32)

	query = queries.GetAllLaporanKuliahByUsertype(s.db, int(idKelas), user)
	meta, data, err := utils.Paginate[map[string]interface{}](pageParams, query, s.logger)

	if err != nil {
		s.logger.Error("failed to get laporan perkuliahan", zap.Error(err))
		return nil, nil, errors.New("failed to get laporan perkuliahan")
	}

	return data, meta, err
}

func (s *LaporanPerkuliahanService) ToPDF(idKelas int, idPertemuan int, idDosen int, idMk int, isPreview bool, baseUrl string) ([]byte, error) {
	var mk *entities.MataKuliah
	var kelas *entities.Kelas
	var angket *entities.AngketDosen
	var dosen *entities.Dosen
	var tanggalPertemuan *entities.JadwalPertemuan

	if err := s.db.First(&mk, map[string]any{"id": idMk}).Error; err != nil {
		s.logger.Error("error while find mata kuliah:", zap.Any("find-mata-kuliah", err.Error()))
		return nil, errors.New("failed when find data")
	}

	if err := s.db.First(&kelas, map[string]any{"id": idKelas}).Error; err != nil {
		s.logger.Error("error while find kelas:", zap.Any("find-kelas", err.Error()))
	}

	if err := s.db.First(&dosen, map[string]any{"id": idDosen}).Error; err != nil {
		s.logger.Error("error while find dosen:", zap.Any("find-dosen", err.Error()))
	}

	if err := s.db.First(&tanggalPertemuan, map[string]any{"id": idPertemuan}).Error; err != nil {
		s.logger.Error("error while find tanggal pertemuan:", zap.Any("find-tanggal-pertemuan", err.Error()))
	}

	if idPertemuan != 0 {
		if err := s.db.First(&angket, map[string]any{"id_pertemuan": idPertemuan}).Error; err != nil {
			s.logger.Error("error while find angket dosen:", zap.Any("find-angket-dosen", err.Error()))
		}
	}

	parsedDate := tanggalPertemuan.TanggalPertemuan.Format("2006-01-02")

	tglPertemuan := utils.ToDateIndo(parsedDate)
	day := utils.TimeToIDDay(tanggalPertemuan.TanggalPertemuan.Format("Mon"))

	templateFile := filepath.Join("templates/", "laporan-kuliah-pdf.html")

	tmpl, err := template.ParseFiles(templateFile)

	if err != nil {
		s.logger.Error("error while parse template:", zap.Any("parse-template", err.Error()))

		return nil, errors.New("failed when parse template")
	}

	data := map[string]interface{}{
		"nama_lengkap":       dosen.NamaLengkap,
		"nama_mk":            mk.NamaMk,
		"inisial_kelas":      kelas.InisialKelas,
		"hari":               day,
		"tanggal_pertemuan":  tglPertemuan,
		"jam_kuliah":         angket.JamKuliah,
		"jumlah_luring":      angket.JumlahLuring,
		"jumlah_daring":      angket.JumlahDaring,
		"jumlah_hadir":       angket.JumlahHadir,
		"jumlah_tidak_hadir": angket.JumlahTidakHadir,
		"pertemuan_ke":       angket.PertemuanKe,
		"ringkasan_materi":   angket.RingkasanMateri,
		"keterangan":         angket.Keterangan,
		"is_preview":         isPreview,
		"base_url":           baseUrl,
		"qr":                 "qrlaporan_" + strconv.Itoa(idPertemuan) + ".svg",
	}

	var body bytes.Buffer

	err = tmpl.Execute(&body, data)

	if err != nil {
		s.logger.Error("error while execute template:", zap.Any("execute-template", err.Error()))
		return nil, errors.New("failed when execute template")
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return nil, err
	}

	// read the HTML page as a PDF page
	page := wkhtmltopdf.NewPageReader(bytes.NewReader(body.Bytes()))

	// enable this if the HTML file contains local references such as images, CSS, etc.
	page.EnableLocalFileAccess.Set(true)

	// add the page to your generator
	pdfg.AddPage(page)

	// manipulate page attributes as needed
	pdfg.MarginLeft.Set(0)
	pdfg.MarginRight.Set(0)
	pdfg.Dpi.Set(300)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	err = pdfg.Create()
	if err != nil {
		return nil, err
	}

	return pdfg.Bytes(), nil
}

func (s *LaporanPerkuliahanService) Delete(idPertemuan int) error {
	var angketDosen *entities.AngketDosen

	if err := s.db.Delete(&angketDosen, map[string]any{"id_pertemuan": idPertemuan}).Error; err != nil {
		s.logger.Error("error while deletee angket dosen:", zap.Any("delete-angket-dosn", err.Error()))
		return errors.New("failed when delete data")
	}

	return nil
}
