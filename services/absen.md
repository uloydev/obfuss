## Query Absen

POST /mahasiswa/absen/save-trans

Flow:
Find jadwalKuliah by IdJawal
Tablename: trans_jadwal_kuliah

Delete All Absen By pertemuan

- Tablename: trans_absen_mahasiswa

Recreate Data:

- Tablename: trans_absen_mahasiswa
- payload: array of {
  IdPertemuan,
  IdMahasiswa,
  IsHadir,
  Keterangan,
  actorId,
  }

Get Jadwal Pertemuan

- Tablename: trans_jadwal_pertemuan
- Update Pertemuan and add flag isTerlaksanakan "yes"
- Count Absen
