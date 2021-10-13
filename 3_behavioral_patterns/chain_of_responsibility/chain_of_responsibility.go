package main

import "fmt"

// Department the Handler.
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

// Patient the object.
type Patient struct {
    Name              string
    RegistrationDone  bool
    DoctorCheckUpDone bool
    MedicineDone      bool
    PaymentDone       bool
}

// Reception an concrete handler.
type Reception struct {
	Next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.RegistrationDone {
		fmt.Printf("Patient %s has already registerd.\n", p.Name)
		r.Next.Execute(p)
		return
	}

	fmt.Printf("Reception is registering patient %s.\n", p.Name)
	p.RegistrationDone = true
	fmt.Printf("Patient %s has already registerd.\n", p.Name)
	r.Next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.Next = next
}

// Doctor an concrete handler.
type Doctor struct {
	Next Department
}

func (r *Doctor) Execute(p *Patient) {
	if p.DoctorCheckUpDone {
		fmt.Printf("Patient %s has already checkuped by doctor.\n", p.Name)
		r.Next.Execute(p)
		return
	}

	fmt.Printf("Doctor is checkuping patient %s.\n", p.Name)
	p.DoctorCheckUpDone = true
	fmt.Printf("Patient %s has already checkuped by doctor.\n", p.Name)
	r.Next.Execute(p)
}

func (r *Doctor) SetNext(next Department) {
	r.Next = next
}

// Medicine an concrete handler.
type Medicine struct {
	Next Department
}

func (r *Medicine) Execute(p *Patient) {
	if p.MedicineDone {
		fmt.Printf("Patient %s has been recieved medicine.\n", p.Name)
		r.Next.Execute(p)
		return
	}

	fmt.Printf("Medicine is giving to patient %s.\n", p.Name)
	p.MedicineDone = true
	fmt.Printf("Patient %s has been recieved medicine.\n", p.Name)
	r.Next.Execute(p)
}

func (r *Medicine) SetNext(next Department) {
	r.Next = next
}

// Cashier an concrete handler.
type Cashier struct {
	Next Department
}

func (r *Cashier) Execute(p *Patient) {
	if p.PaymentDone {
		fmt.Printf("Patient %s has paid.\n", p.Name)
		return
	}

	fmt.Printf("Cashier is getting money from patient %s.\n", p.Name)
	p.PaymentDone = true
	fmt.Printf("Patient %s has paid.\n", p.Name)
}

func (r *Cashier) SetNext(next Department) {
	r.Next = next
}

func main() {
	cashier := &Cashier{}
    //Set next for medical department
    medical := &Medicine{}
    medical.SetNext(cashier)
    //Set next for doctor department
    doctor := &Doctor{}
    doctor.SetNext(medical)
    //Set next for reception department
    reception := &Reception{}
    reception.SetNext(doctor)
    patient := &Patient{Name: "CR7"}
    //Patient visiting
    reception.Execute(patient)
}