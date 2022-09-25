package models

type Users struct {
	Id        string `form:"id" json:"id"`
	Nip string `form:"nip" json:"nip"`
	Nama string `form:"nama" json:"nama"`
	Password  string `form:"password" json:"password"`
	Role string `form:"role" json:"role"`
}

type Spp struct {
	Id        string `form:"id" json:"id"`
	NoSpp string `form:"no_spp" json:"no_spp"`
}

type Login struct {
	Id        string `form:"id" json:"id"`
	Nip string `form:"nip" json:"nip"`
	Nama string `form:"nama" json:"nama"`
	Password  string `form:"password" json:"password"`
	Role string `form:"role" json:"role"`
}


type Kelas struct {
	Id        string `form:"id" json:"id"`
	Nip string `form:"nip" json:"nip"`
	Nama string `form:"nama" json:"nama"`
	Password  string `form:"password" json:"password"`
	Role string `form:"role" json:"role"`
}


type Jurusan struct {
	Id        string `form:"id" json:"id"`
	Nip string `form:"nip" json:"nip"`
	Nama string `form:"nama" json:"nama"`
	Password  string `form:"password" json:"password"`
	Role string `form:"role" json:"role"`
}

