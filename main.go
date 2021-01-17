package main

import (
	"fmt"

	"github.com/JulianDavidGamboa/composition/pkg/customer"
	"github.com/JulianDavidGamboa/composition/pkg/invoice"
	"github.com/JulianDavidGamboa/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Colombia",
		"Cali",
		customer.New("Julian", "Cl 5 # 66 C - 14", "1235464"),
		invoiceitem.NewItems(

			invoiceitem.New(1, "Curso de Go", 12.34),
			invoiceitem.New(2, "Curso de POO con Go", 54.23),
			invoiceitem.New(3, "Curso de testing con Go", 90.00),
		),
	)
	i.SetTotal()
	fmt.Printf("%+v\n", i)
}
