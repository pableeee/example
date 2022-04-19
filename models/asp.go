package models

import "fmt"

type Usuario struct {
	ID                         uint   `gorm:"primary_key,id:id"`
	RazonSocial                string `gorm:"column:razon_social"`
	CUIT                       string `gorm:"column:cuit_empresa"`
	CalleDomicilioLegal        string `gorm:"column:calle_domicilio_legal"`
	NumeroDomicilioLegal       string `gorm:"column:numero_domicilio_legal"`
	CiudadDomicilioLegal       string `gorm:"column:ciudad_domicilio_legal"`
	CPDomicilioLegal           string `gorm:"column:cp_domicilio_legal"`
	CalleDomicilioConstituido  string `gorm:"column:calle_domicilio_constituido"`
	NumeroDomicilioConstituido string `gorm:"column:numero_domicilio_constituido"`
	CiudadDomicilioConstituido string `gorm:"column:ciudad_domicilio_constituido"`
	CPDomicilioConstituido     string `gorm:"column:cp_domicilio_constituido"`
}

func GetUsuarios() ([]Usuario, error) {
	var usuarios []Usuario

	result := msqlOrm.Find(&usuarios)
	if result.Error != nil {
		return nil, fmt.Errorf("unable to retieve usuarios %w", err)
	}

	return usuarios, nil
}
