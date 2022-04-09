package util

const (
	No = "No"

	TableListFormat = "%-10d"
	TableList       = "DAFTAR MEJA MAKAN TERSEDIA"

	MenuList         = "DAFTAR MENU MAKANAN"
	MenuListFormat   = "%-5s%-20s%-30s\n"
	NamaMakanan      = "Nama Makanan"
	HargaMakanan     = "Harga"
	MenuListDbFormat = "%-5d%-20s%-30d\n"

	ConfirmationForNextOption = "Apakah anda sudah yakin dengan pilihan anda (y/n) ? "
	TableChoosen              = "Silahkan pilih meja anda : "
	ChooseMenu                = "Silahkan pilih makanan anda : "
	ChooseMenu2               = "Apakah ada lagi (y/n) ? "
	AskName                   = "Siapa nama anda? "
	PaymentStatus             = "Processing"
	PaymentStatus2            = "Finish"
)

func IfError(err error) {
	if err != nil {
		panic(err)
	}
}
