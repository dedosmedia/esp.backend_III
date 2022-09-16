package main

type PaymentMethod interface {
	Pay(purchase *Purchase)
}

type Purchase struct {
	Name          string
	Precio        float32
	metodosDePago []PaymentMethod
}

type TarjetaCredito struct {
}

func (tc TarjetaCredito) Pay(purchase *Purchase) {
}

type TransferenciaBancaria struct {
}

func (tb TransferenciaBancaria) Pay(purchase *Purchase) {
}

type Efectivo struct {
}

func (e Efectivo) Pay(purchase *Purchase) {

}

func main() {

	agua := CrearCompra("agua", 100, TarjetaCredito{}, Efectivo{})
	Process(TarjetaCredito{})

}

func CrearCompra(nombre string, precio float32, metodos ...PaymentMethod) Purchase {

	return Purchase{
		Name:          nombre,
		Precio:        precio,
		metodosDePago: metodos,
	}
}

func Process(metodo PaymentMethod, purchase Purchase) {

	metodo.Pay(&purchase)

}
