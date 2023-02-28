package main

func main() {
	cashier := &Cashier{}

	//Set next to medical departement
	medical := &Medical{}
	medical.setNext(cashier)

	//Set next for doctor departement
	doctor := &Doctor{}
	doctor.setNext(medical)

	//Set next for Reception
	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "abc"}

	//Patient visitng
	reception.execute(patient)

}
