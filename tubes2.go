package main

import "fmt"

const NMAX = 1000
const MK = 1000
const q = 15

type mahasiswa struct {
	nim, nama, kelas string
	matakuliah       [MK]matkul
	transkrip        hasil
	tot_nilai        float64
	tot_sks          int
}

type matkul struct {
	kode, nama string
	sks        int
	nilai      point
}

type point struct {
	uas, uts hasil
	quiz     [q]hasil
}

type hasil struct {
	grade string
	angka float64
}

type tabArr [NMAX]mahasiswa

func main() {
	var n, nData, nMatkul, nQuiz int
	var data tabArr
	var pw string
	fmt.Println("     Selamat Datang! Silahkan masukkan password ")
	fmt.Print("     Password:")
	fmt.Scan(&pw)
	for pw != "TelkomUniversityDatabase" {
		fmt.Println("     Password yang anda masukkan salah.")
		fmt.Print("     Silahkan masukkan password yang benar:")
		fmt.Scan(&pw)
	}
	for n != 7 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Selamat Datang di Pusat Informasi Mahasiswa   ")
		fmt.Println("       Silahkan pilih sesuai kebutuhan Anda       ")
		fmt.Println("                                                  ")
		fmt.Println("    1. Data Mahasiswa                             ")
		fmt.Println("    2. Pengurutan data                            ")
		fmt.Println("    3. Data seluruh mahasiswa                     ")
		fmt.Println("    4. Data mata kuliah yang diambil mahasiswa    ")
		fmt.Println("    5. Pencarian data mahasiswa tertentu          ")
		fmt.Println("    6. Transkrip seluruh mahasiswa                ")
		fmt.Println("    7. Logout                                     ")
		fmt.Println("       Masukkan angka 1/2/3/4/5/6                 ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Input: ")
		fmt.Scan(&n)
		if n == 1 {
			dataMahasiswa(&data, &nData, &nMatkul, &nQuiz)
		} else if n == 2 {
			pencarianData(&data, nData, nMatkul, nQuiz)
		} else if n == 3 {
			tampilmhs(&data, nData)
		} else if n == 4 {
			tampilmkdiambilmhs(data, nData, nMatkul)
		} else if n == 5 {
			cariDataMahasiswa(data, nData, nMatkul, nQuiz)
		} else if n == 6 {
			transkrip(&data, nData, nMatkul, nQuiz)
		}
	}
	fmt.Println("     Anda sudah logout")
	fmt.Print("     Terima kasih sudah menggunakan aplikasi ini!")
}

func dataMahasiswa(A *tabArr, n, m *int, qz *int) {
	var pilihan int
	for pilihan != 5 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("          Apa yang sedang anda butuhkan?          ")
		fmt.Println("                                                  ")
		fmt.Println("       1. Input Data Mahasiswa                    ")
		fmt.Println("       2. Hapus Data Mahasiswa                    ")
		fmt.Println("       3. Edit Data Mahasiswa                     ")
		fmt.Println("       4. Data Matakuliah Mahasiswa               ")
		fmt.Println("       5. Kembali ke menu utama                   ")
		fmt.Println("       Masukkan angka 1/2/3/4/5                   ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Input: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			inputDataMahasiswa(A, n)
		} else if pilihan == 2 {
			hapusDataMahasiswa(A, n)
		} else if pilihan == 3 {
			editDataMahasiswa(A, n)
		} else if pilihan == 4 {
			DataMatakuliahMahasiswa(A, n, m, qz)
		} else if pilihan == 5 {
			fmt.Println("       Baik, anda akan segera menuju ke menu utama")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputDataMahasiswa(A *tabArr, n *int) {
	var jawaban string
	var nilai int
	fmt.Println()
	fmt.Print("Berapa yang ingin anda masukkan: ")
	fmt.Scan(&nilai)
	fmt.Println("Silahkan masukkan NIM, nama dengan format namaDepan_namaBelakang, dan kelas: ")
	for i := 0; i < nilai; i++ {
		fmt.Scan(&A[*n].nim)
		fmt.Scan(&A[*n].nama)
		fmt.Scan(&A[*n].kelas)
		*n += 1
	}
	fmt.Println("Data telah diinputkan!")
	fmt.Println("Apakah anda ingin input data lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputDataMahasiswa(A, n)
	}
}

func hapusDataMahasiswa(A *tabArr, n *int) {
	var nim, jawaban string
	fmt.Println()
	fmt.Println("Masukkan data NIM mahasiswa yang ingin dihapus:")
	fmt.Scan(&nim)
	a := findDataNim(*A, *n, nim)
	fmt.Println(a)
	if a != -1 {
		fmt.Println("Data NIM Mahasiswa ditemukan")
		fmt.Println("Apakah anda ingin menghapus data mahasiswa tersebut?")
		fmt.Print("Ya/Tidak:")
		fmt.Scan(&jawaban)
		if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
			for a <= *n-1 {
				A[a] = A[a+1]
				a++
			}
		}
		fmt.Println("Data telah dihapus!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin hapus data lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusDataMahasiswa(A, n)
	}
}

func editDataMahasiswa(A *tabArr, n *int) {
	var nim, jawaban string
	var a int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda edit?")
	fmt.Print("Masukkan NIM Mahasiswa yang ingin di edit: ")
	fmt.Scan(&nim)
	a = findDataNim(*A, *n, nim)
	if a != -1 {
		fmt.Println("Data Mahasiswa ditemukan")
		fmt.Println("Silahkan masukkan NIM, nama, dan kelas mahasiswa: ")
		fmt.Scan(&A[a].nim)
		fmt.Scan(&A[a].nama)
		fmt.Scan(&A[a].kelas)
		fmt.Println("Data telah diedit!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin edit data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		editDataMahasiswa(A, n)
	}
}

func DataMatakuliahMahasiswa(A *tabArr, n, m *int, qz *int) {
	var pilihan int
	for pilihan != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("         Data apa yang ingin anda edit?           ")
		fmt.Println("                                                  ")
		fmt.Println("         1. Input / Edit Mata Kuliah              ")
		fmt.Println("         2. Nilai UTS dan UAS                     ")
		fmt.Println("         3. Nilai Quiz                            ")
		fmt.Println("         4. Kembali ke menu sebelumnya            ")
		fmt.Println("         Masukkan angka 1/2/3/4                   ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Input: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			mataKuliah(A, *n, m)
		} else if pilihan == 2 {
			UtsUas(A, *n, *m)
		} else if pilihan == 3 {
			Quiz(A, qz, *n, *m)
		} else if pilihan == 4 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func mataKuliah(A *tabArr, n int, m *int) {
	var jawaban int
	for jawaban != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("           Menu apa yang anda butuhkan?          ")
		fmt.Println("           1. Input Data Mata Kuliah              ")
		fmt.Println("           2. Edit Data Mata Kuliah               ")
		fmt.Println("           3. Hapus Data Mata Kuliah              ")
		fmt.Println("           4. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("    Input: ")
		fmt.Scan(&jawaban)
		if jawaban == 1 {
			inputMataKuliah(A, n, m)
		} else if jawaban == 2 {
			editMataKuliah(A, n, *m)
		} else if jawaban == 3 {
			hapusDataMatkul(A, n, m)
		} else if jawaban == 4 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputMataKuliah(A *tabArr, n int, m *int) {
	var nim, namkul, kodkul, jawaban string
	var a, skskul, nilai int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda input?")
	fmt.Print("Masukkan NIM Mahasiswa: ")
	fmt.Scan(&nim)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		fmt.Println("Data Mahasiswa ditemukan")
		fmt.Println("Berapa data mata kuliah yang ingin anda masukkan?")
		fmt.Print("Masukkan: ")
		fmt.Scan(&nilai)
		fmt.Println("Silahkan masukkan kode, nama, dan sks mata kuliah: ")
		for i := 0; i < nilai; i++ {
			fmt.Scan(&kodkul, &namkul, &skskul)
			A[a].matakuliah[*m].kode = kodkul
			A[a].matakuliah[*m].nama = namkul
			A[a].matakuliah[*m].sks = skskul
			*m++
		}
		fmt.Println("Data telah diinputkan!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin input data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputMataKuliah(A, n, m)
	}
}

func editMataKuliah(A *tabArr, n, m int) {
	var nim, namkul, kodkul, jawaban, kode string
	var a, b, skskul int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda edit?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode mata kuliah yang ingin di edit: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], m, kode)
	}
	if a != -1 && b != -1 {
		fmt.Println("Data Mahasiswa ditemukan")
		fmt.Println("Silahkan masukkan kode, nama, dan sks mata kuliah: ")
		fmt.Scan(&kodkul, &namkul, &skskul)
		A[a].matakuliah[b].kode = kodkul
		A[a].matakuliah[b].nama = namkul
		A[a].matakuliah[b].sks = skskul
		fmt.Println("Data telah diedit!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin edit data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		editMataKuliah(A, n, m)
	}
}

func hapusDataMatkul(A *tabArr, n int, m *int) {
	var nim, jawaban, kode string
	var a, b int
	fmt.Println()
	fmt.Println("Masukkan data NIM dan KODE mata kuliah mahasiswa yang ingin dihapus:")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], *m, kode)
	}
	if a != -1 && b != -1 {
		fmt.Println("Data NIM dan mata kuliah Mahasiswa ditemukan")
		fmt.Println("Apakah anda ingin menghapus data mata kuliah mahasiswa tersebut?")
		fmt.Print("Ya/Tidak:")
		fmt.Scan(&jawaban)
		if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
			for b < *m-1 {
				A[a].matakuliah[b] = A[a].matakuliah[b+1]
				b++
			}
		}
		fmt.Println("Data telah dihapus!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin hapus data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusDataMatkul(A, n, m)
	}
}

func UtsUas(A *tabArr, n, m int) {
	var jawaban int
	for jawaban != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("          Menu apa yang anda butuhkan?           ")
		fmt.Println("          1. Input / Edit Data UTS dan UAS        ")
		fmt.Println("          2. Hapus Data UTS dan UAS               ")
		fmt.Println("          3. Kembali ke menu sebelumnya           ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("    Input: ")
		fmt.Scan(&jawaban)
		if jawaban == 1 {
			inputUtsUas(A, n, m)
		} else if jawaban == 2 {
			hapusUtsUas(A, n, m)
		} else if jawaban == 3 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputUtsUas(A *tabArr, n, m int) {
	var nim, kode, jawaban, uas, uts string
	var UAS, UTS float64
	var a, b int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda input?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], m, kode)
	}
	if a != -1 && b != -1 {
		fmt.Println("Data Mahasiswa ditemukan")
		fmt.Println("Silahkan masukkan nilai UTS dan UAS beserta grade masing-masing terurut: ")
		fmt.Scan(&UTS, &UAS, &uts, &uas)
		*&A[a].matakuliah[b].nilai.uts.angka = UTS
		*&A[a].matakuliah[b].nilai.uas.angka = UAS
		*&A[a].matakuliah[b].nilai.uts.grade = uts
		*&A[a].matakuliah[b].nilai.uas.grade = uas
		fmt.Println("Data telah diinputkan!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin input/edit data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputUtsUas(A, n, m)
	}
}

func hapusUtsUas(A *tabArr, n, m int) {
	var nim, kode, jawaban string
	var a, b int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda hapus?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], m, kode)
	}
	if a != -1 && b != -1 {
		fmt.Println("Data NIM dan mata kuliah Mahasiswa ditemukan")
		fmt.Println("Apakah anda ingin menghapus data nilai UTS dan UAS mahasiswa tersebut?")
		fmt.Print("Ya/Tidak:")
		fmt.Scan(&jawaban)
		if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
			A[a].matakuliah[b].nilai.uts.angka = 0
			A[a].matakuliah[b].nilai.uts.grade = "-"
			A[a].matakuliah[b].nilai.uas.angka = 0
			A[a].matakuliah[b].nilai.uas.grade = "-"
		}
		fmt.Println("Data telah dihapus!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin hapus data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusUtsUas(A, n, m)
	}
}

func Quiz(A *tabArr, qz *int, n, m int) {
	var jawaban int
	for jawaban != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("         Menu apa yang anda butuhkan?            ")
		fmt.Println("         1. Input Data Quiz                       ")
		fmt.Println("         2. Edit Data Quiz                        ")
		fmt.Println("         3. Hapus Data Quiz                       ")
		fmt.Println("         4. Kembali ke menu sebelumnya            ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("    Input: ")
		fmt.Scan(&jawaban)
		if jawaban == 1 {
			inputQuiz(A, qz, n, m)
		} else if jawaban == 2 {
			editQuiz(A, qz, n, m)
		} else if jawaban == 3 {
			hapusQuiz(A, qz, n, m)
		} else if jawaban == 4 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func inputQuiz(A *tabArr, qz *int, n, m int) {
	var a, b, c int
	var Q float64
	var quiz, nim, kode, jawaban string
	fmt.Println()
	fmt.Println("Data Mahasiswa mana yang ingin anda input nilai quiznya?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], m, kode)
	}
	if a != -1 && b != -1 {
		fmt.Println("Data Mahasiswa ditemukan")
		fmt.Print("Ada berapa quiz yang ingin di input nilainya?:")
		fmt.Scan(&c)
		fmt.Println("Silahkan masukkan nilai quiz beserta grade: ")
		if c > q {
			c = q
		}
		for i := 0; i < c; i++ {
			fmt.Scan(&Q, &quiz)
			A[a].matakuliah[b].nilai.quiz[*qz].angka = Q
			A[a].matakuliah[b].nilai.quiz[*qz].grade = quiz
			*qz++
		}
		fmt.Println("Data telah diinputkan!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println(A[a].matakuliah[b].nilai.quiz[*qz].angka, A[a].matakuliah[b].nilai.quiz[*qz].grade)
	fmt.Println("Apakah anda ingin input data quiz lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		inputQuiz(A, qz, n, m)
	}
}

func editQuiz(A *tabArr, qz *int, n, m int) {
	var Q float64
	var quiz, nim, kode, jawaban string
	var nilai, a, b int
	fmt.Println()
	fmt.Println("Data Mahasiswa yang mana yang ingin anda edit?")
	fmt.Print("Masukkan NIM Mahasiswa dan kode Mata kuliah: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], m, kode)
	}
	if a != -1 && b != -1 {
		fmt.Println("Data Mahasiswa ditemukan")
		fmt.Print("Silahkan masukkan quiz ke berapa yang ingin anda edit: ")
		fmt.Scan(&nilai)
		fmt.Println("Silahkan masukkan nilai quiz beserta grade: ")
		fmt.Scan(&Q, &quiz)
		A[a].matakuliah[b].nilai.quiz[nilai-1].angka = Q
		A[a].matakuliah[b].nilai.quiz[nilai-1].grade = quiz
		fmt.Println("Data telah diedit!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin edit data quiz lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		editQuiz(A, qz, n, m)
	}
}

func hapusQuiz(A *tabArr, qz *int, n, m int) {
	var nim, kode, jawaban string
	var a, b, quiz int
	fmt.Println()
	fmt.Println("Masukkan NIM Mahasiswa dan kode Mata kuliah yang ingin anda hapus nilai quiznya: ")
	fmt.Scan(&nim, &kode)
	a = findDataNim(*A, n, nim)
	if a != -1 {
		b = findDataMatkul(A[a], m, kode)
	}
	if a != -1 && b != -1 {
		fmt.Println("Data NIM dan mata kuliah Mahasiswa ditemukan")
		fmt.Print("Masukkan quiz ke berapa yang ingin anda hapus?: ")
		fmt.Scan(&quiz)
		if quiz-1 < *qz {
			fmt.Println("Apakah anda ingin menghapus data mata kuliah mahasiswa tersebut?")
			fmt.Print("Ya/Tidak:")
			fmt.Scan(&jawaban)
			if jawaban == "ya" || jawaban == "Ya" || jawaban == "YA" {
				A[a].matakuliah[b].nilai.quiz[quiz-1].angka = 0
				A[a].matakuliah[b].nilai.quiz[quiz-1].grade = "-"
			}
		}
		fmt.Println("Data telah dihapus!")
	} else {
		fmt.Println("Mohon maaf, data yang anda cari tidak ditemukan")
	}
	fmt.Println("Apakah anda ingin hapus data mata kuliah lainnya?")
	fmt.Print("Ya/Tidak:")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		hapusQuiz(A, qz, n, m)
	}
}

func findDataNim(A tabArr, n int, x string) int {
	var idx int
	idx = -1
	i := 0
	for i < n && idx == -1 {
		if x == A[i].nim {
			idx = i
		}
		i++
	}
	return idx
}

func findDataMatkul(mhs mahasiswa, n int, x string) int {
	var idx int
	idx = -1
	i := 0
	for i < n && idx == -1 {
		if x == mhs.matakuliah[i].kode {
			idx = i
		}
		i++
	}
	return idx
}

func rataQuiz(A tabArr, n, m, qz int) float64 {
	var rerataQuiz, total float64
	var i int
	i = 0
	for i = 0; i < qz; i++ {
		total += A[n].matakuliah[m].nilai.quiz[i].angka
	}
	rerataQuiz = total / float64(i)
	return rerataQuiz
}

func totalNilai(A *tabArr, n, m, qz int) {
	var rerataQuiz, total float64
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rerataQuiz = rataQuiz(*A, i, m, qz)
			total += A[i].matakuliah[j].nilai.uas.angka + A[i].matakuliah[j].nilai.uts.angka + rerataQuiz/3*25
		}
		A[i].tot_nilai = total
	}
}

func transkrip(A *tabArr, n, m, qz int) {
	totalNilai(A, n, m, qz)
	fmt.Println()
	fmt.Printf("%15s %40s %10s %6s", "NIM", "NAMA", "NILAI", "GRADE")
	for i := 0; i < n; i++ {
		a := totalSks(A, i, m)
		A[i].transkrip.angka = (A[i].tot_nilai * float64(a)) / float64(a)
		if A[i].transkrip.angka == 4 {
			A[i].transkrip.grade = "A"
		} else if A[i].transkrip.angka >= 3.5 && A[i].transkrip.angka < 4 {
			A[i].transkrip.grade = "AB"
		} else if A[i].transkrip.angka >= 3 && A[i].transkrip.angka < 3.5 {
			A[i].transkrip.grade = "B"
		} else if A[i].transkrip.angka >= 2.5 && A[i].transkrip.angka < 3 {
			A[i].transkrip.grade = "BC"
		} else if A[i].transkrip.angka >= 2 && A[i].transkrip.angka < 2.5 {
			A[i].transkrip.grade = "C"
		} else if A[i].transkrip.angka >= 1 && A[i].transkrip.angka < 2 {
			A[i].transkrip.grade = "D"
		} else if A[i].transkrip.angka >= 0 && A[i].transkrip.angka < 1 {
			A[i].transkrip.grade = "E"
		}
		fmt.Printf("%15s %40s %3.2f %2s", A[i].nim, A[i].nama, A[i].transkrip.angka, A[i].transkrip.grade)
	}
	fmt.Println()
}

func totalSks(A *tabArr, n, m int) int {
	var total int
	for j := 0; j < m; j++ {
		total += A[n].matakuliah[m].sks
	}
	return total
}

func pencarianData(A *tabArr, n, m, qz int) {
	var pilihan int
	for pilihan != 4 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("           Menu apa yang anda butuhkan?           ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Data menurut Nilai                  ")
		fmt.Println("           2. Data menurut SKS                    ")
		fmt.Println("           3. Data menurut NIM                    ")
		fmt.Println("           4. Kembali ke menu sebelumnya          ")
		fmt.Println("           Masukkan angka 1/2/3/4                 ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("      Input:")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			dataMenurutNilai(A, n)
		} else if pilihan == 2 {
			dataMenurutSks(A, n, m)
		} else if pilihan == 3 {
			dataMenurutNim(A, n)
		} else if pilihan == 4 {
			fmt.Println("       Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("       Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func dataMenurutNilai(A *tabArr, n int) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Jenis data terurut apa yang anda inginkan?    ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Terurut membesar                    ")
		fmt.Println("           2. Terurut mengecil                    ")
		fmt.Println("           3. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("      Input:")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			nilaiUrutMembesar(A, n)
		} else if pilihan == 2 {
			nilaiUrutMengecil(A, n)
		} else if pilihan == 3 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func nilaiUrutMembesar(A *tabArr, n int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		for j < n {
			if (A[idx].tot_nilai > A[j].tot_nilai) || (A[idx].tot_nilai == A[j].tot_nilai && A[idx].transkrip.angka > A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilnilai(A, n)
}

func nilaiUrutMengecil(A *tabArr, n int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		for j < n {
			if (A[idx].tot_nilai < A[j].tot_nilai) || (A[idx].tot_nilai == A[j].tot_nilai && A[idx].transkrip.angka < A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilnilai(A, n)
}

func dataMenurutSks(A *tabArr, n, m int) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Jenis data terurut apa yang anda inginkan?    ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Terurut membesar                    ")
		fmt.Println("           2. Terurut mengecil                    ")
		fmt.Println("           3. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("      Input:")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			sksUrutMembesar(A, n, m)
		} else if pilihan == 2 {
			sksUrutMengecil(A, n, m)
		} else if pilihan == 3 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func sksUrutMembesar(A *tabArr, n, m int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		A[idx].tot_sks = totalSks(A, idx, m)
		A[j].tot_sks = totalSks(A, j, m)
		for j < n {
			if (A[idx].tot_sks > A[j].tot_sks) || (A[idx].tot_sks == A[j].tot_sks && A[idx].transkrip.angka > A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilsks(A, n)
}

func sksUrutMengecil(A *tabArr, n, m int) {
	var i, j, idx int
	var t mahasiswa
	i = 1
	for i <= n-1 {
		idx = i - 1
		j = 1 + 1
		A[idx].tot_sks = totalSks(A, idx, m)
		A[j].tot_sks = totalSks(A, j, m)
		for j < n {
			if (A[idx].tot_sks < A[j].tot_sks) || (A[idx].tot_sks == A[j].tot_sks && A[idx].transkrip.angka < A[j].transkrip.angka) {
				idx = j
			}
			j++
		}
		t = A[idx]
		A[idx] = A[i-1]
		A[i-1] = t
		i++
	}
	tampilsks(A, n)
}

func dataMenurutNim(A *tabArr, n int) {
	var pilihan int
	for pilihan != 3 {
		fmt.Println()
		fmt.Println("--------------------------------------------------")
		fmt.Println("    Jenis data terurut apa yang anda inginkan?    ")
		fmt.Println("                                                  ")
		fmt.Println("           1. Terurut membesar                    ")
		fmt.Println("           2. Terurut mengecil                    ")
		fmt.Println("           3. Kembali ke menu sebelumnya          ")
		fmt.Println("--------------------------------------------------")
		fmt.Print("       Masukkan: ")
		fmt.Scan(&pilihan)
		if pilihan == 1 {
			nimUrutMembesar(A, n)
		} else if pilihan == 2 {
			nimUrutMengecil(A, n)
		} else if pilihan == 3 {
			fmt.Println("Baik, anda akan segera menuju ke menu sebelumnya")
			fmt.Println("Mohon tunggu sebentar")
			fmt.Println()
		}
	}
}

func nimUrutMembesar(A *tabArr, n int) {
	var i, j int
	var temp mahasiswa
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.nim < A[j-1].nim {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i++
	}
	tampilmhs(A, n)
}

func nimUrutMengecil(A *tabArr, n int) {
	var i, j int
	var temp mahasiswa
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.nim > A[j-1].nim {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i++
	}
	tampilmhs(A, n)
}

func pokoknyaUrutbdamatlahcapek(A *tabArr, n int) {
	var i, j int
	var temp mahasiswa
	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]
		for j > 0 && temp.nim > A[j-1].nim {
			A[j] = A[j-1]
			j = j - 1
		}
		A[j] = temp
		i++
	}
}

func cariDataMahasiswa(A tabArr, n, m, qz int) {
	var nim string
	fmt.Println()
	fmt.Println("Silahkan masukkan NIM mahasiswa yang ingin anda cari")
	fmt.Print("Masukkan :")
	fmt.Scan(&nim)
	fmt.Println()
	pokoknyaUrutbdamatlahcapek(&A, n)
	hasil := binarySearch(A, n, nim)
	if hasil != -1 {
		fmt.Println("Data mahasiswa ditemukan!")
		fmt.Println("Jika tidak ada mata kuliah atau nilai yang diinputkan,")
		fmt.Println("tampilan akan kosong")
		fmt.Println("Berikut adalah data mahasiswa tersebut:")
		fmt.Printf("%15s %40s %10s %15s %20s %3s %10s %10s %10s\n", "NIM ", " NAMA", "KELAS", "KODE MATKUL", "NAMA MATKUL", "SKS", "UTS", "UAS", "QUIZ")
		fmt.Printf("%15s %40s %10s\n", A[hasil].nim, A[hasil].nama, A[hasil].kelas)
		for j := 0; j < m; j++ {
			for k := 0; k < qz; k++ {
				fmt.Printf("%15s %40s %10s %15s %20s %1d %3.2f %2s %3.2f %2s %3.2f %2s\n", A[hasil].nim, A[hasil].nama, A[hasil].kelas, A[hasil].matakuliah[j].kode, A[hasil].matakuliah[j].nama, A[hasil].matakuliah[j].sks, A[hasil].matakuliah[j].nilai.uts.angka, A[hasil].matakuliah[j].nilai.uts.grade, A[hasil].matakuliah[j].nilai.uas.angka, A[hasil].matakuliah[j].nilai.uas.grade, A[hasil].matakuliah[j].nilai.quiz[k].angka, A[hasil].matakuliah[j].nilai.quiz[k].grade)
			}
		}
	} else {
		fmt.Println("Mohon maaf, data tersebut tidak ditemukan")
	}
}

func binarySearch(A tabArr, n int, nim string) int {
	var found int
	var kiri, kanan, tengah int
	found = -1
	kiri = 0
	kanan = n - 1
	for kiri <= kanan && found == -1 {
		tengah = (kiri + kanan) / 2
		if nim > A[tengah].nim {
			kanan = tengah - 1
		} else if nim < A[tengah].nim {
			kiri = tengah + 1
		} else {
			found = tengah
		}
	}
	return found
}

func datanim(A tabArr, n int) {
	var nim string
	var hasil, i int
	hasil = -1
	i = 0
	for i < n && hasil == -1 {
		if nim == A[i].nim {
			hasil = i
		}
	}
}

func tampilmhs(A *tabArr, n int) {
	fmt.Printf("%15s %40s %10s\n", "NIM ", " NAMA", "KELAS")
	for i := 0; i < n; i++ {
		fmt.Printf("%15s %40s %10s\n", A[i].nim, A[i].nama, A[i].kelas)
	}
}

func tampilmkdiambilmhs(A tabArr, n, m int) {
	var kode, jawaban string
	fmt.Println()
	fmt.Print("Masukkan kode mata kuliah yang ingin anda tampilkan datanya: ")
	fmt.Scan(&kode)
	fmt.Println("Jika mata kuliah tidak ada atau tidak ada mahasiswa yang")
	fmt.Println("mengambil mata kuliah ini, maka hasilnya akan kosong")
	fmt.Println("Berikut adalah data mahasiswa yang mengambil mata kuliah ini")
	fmt.Println()

	fmt.Printf("%15s %40s %10s\n", "NIM", "NAMA", "KELAS")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if kode == A[i].matakuliah[m].kode {
				fmt.Printf("%15s %40s %10s\n", A[i].nim, A[i].nama, A[i].kelas)
			}
		}
	}
	fmt.Println("Apakah anda ingin menampilkan data mata kuliah yang lain?")
	fmt.Print("Ya/Tidak: ")
	fmt.Scan(&jawaban)
	if jawaban == "Ya" || jawaban == "ya" || jawaban == "YA" {
		tampilmkdiambilmhs(A, n, m)
	}
}

func tampilsks(A *tabArr, n int) {
	fmt.Printf("%15s %40s %10s %3s\n", "NIM ", " NAMA", "KELAS", "SKS")
	for i := 0; i < n; i++ {
		fmt.Printf("%15s %40s %10s %3d\n", A[i].nim, A[i].nama, A[i].kelas, A[i].tot_sks)
	}
}

func tampilnilai(A *tabArr, n int) {
	fmt.Printf("%15s %40s %10s %12s\n", "NIM ", " NAMA", "KELAS", "TOTAL NILAI")
	for i := 0; i < n; i++ {
		fmt.Printf("%15s %40s %10s %7.2f\n", A[i].nim, A[i].nama, A[i].kelas, A[i].tot_nilai)
	}
}