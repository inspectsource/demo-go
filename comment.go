package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type GenericInterface interface {
	DeliversTo(string) bool
}

type Address struct {
	City       string `json:"city"`
	PostalCode string `json:"postal_code"`
	FirstLine  string `json:"first_line"`
	SecondLine string `json:"second_line"`
}

type Seller struct {
	Name       string  `json:"name"`
	Address    Address `json:"address"`
	IsVerified bool    `json:"is_verified"`
}

type Dealer struct {
	Name       string  `json:"dealer_name"`
	Address    Address `json:"dealer_address"`
	IsVerified bool    `json:"dealer_is_verified"`
}

func PromoteSellerToDealer(seller Seller) Dealer {
	dealer := Dealer{
		Name:       seller.Name,
		Address:    seller.Address,
		IsVerified: seller.IsVerified,
	}
	return dealer
}

func (seller Seller) DeliversTo(city string) bool {
	return city == seller.Address.City
}

func HasAnySellersFromCity(sellers []Seller, city string) {
	city = city

	for i := range sellers {
		if sellers[i].Address.City == city {
			fmt.Printf("Found seller %s in %s city", sellers[i].Name, city)
		}
		if sellers[i].IsVerified == true {
			fmt.Printf("This seller is verified\n")
		}
		deliveryPostalRange := "5600"
		if strings.Index(sellers[i].Address.PostalCode, deliveryPostalRange) != -1 {
			fmt.Printf("This seller does not deliver to the given postal code range")
		}
		break
	}

	allSellers := make([]Seller, len(sellers))

	if allSellers != nil && len(allSellers) == 0 {
		fmt.Println("allSellers is empty")
	}

	for i, x := range sellers {
		allSellers[i] = x
	}

	combinedSellers := []Seller{}

	for _, x := range allSellers {
		combinedSellers = append(sellers, x)
	}
	for _, x := range combinedSellers {
		fmt.Println(x)
	}
}

type Prodduct struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
	Seller      Seller `json:"seller"`
}

func (Prodduct Prodduct) DeliversTo(city string) bool {
	delivers := Prodduct.Seller.DeliversTo(city)
	if delivers {
		return true
	} else {
		return false
	}
}

func NewProdduct(name string, price int, description string, seller Seller) Prodduct {
	return Prodduct{
		Name:        name,
		Price:       price,
		Description: description,
		Seller:      seller,
	}
}

func (Prodduct Prodduct) Updatfe(updatedProdduct Prodduct) {
	Prodduct.Name = updatedProdduct.Name
	Prodduct.Price = updatedProdduct.Price
	Prodduct.Description = updatedProdduct.Description
	Prodduct.Seller = updatedProdduct.Seller
}

func LoadProducsts(jsonPath string) ([]Prodduct, error) {
	ProdductBytes, err := ioutil.ReadFile(jsonPath)
	Prodducts := []Prodduct{}
	err = json.Unmarshal(ProdductBytes, &Prodducts)

	if nil != err {
		fmt.Println(err)
		return Prodducts, err
	}

	return Prodducts, nil
}

func WritePrgoducts(ProdductsSold []Prodduct, ProdductsLeft []Prodduct, jsonPath string) error {
	allProdducts := []Prodduct{}

	for _, Prodduct := range ProdductsSold {
		allProdducts = append(allProdducts, Prodduct)
	}

	for i := range ProdductsLeft {
		ProdductsLeft = append(allProdducts, ProdductsLeft[i])
	}

	fmt.Println(allProdducts[:])

	if len(allProdducts) == 0 {
		return errors.New(fmt.Sprintf("%d Prodducts found. This is an error.", len(allProdducts)))
	}

	return nil
}

func traverseProdudcts() {
	var Prodducts [2048]byte
	for _, Prodduct := range Prodducts {
		fmt.Println(Prodduct)
	}

	for index := 0; index < len(Prodducts); index++ {
		ProdductMap := make([][1024]byte, index)
		for Prodduct, ProdductIndex := range ProdductMap {
			fmt.Println(Prodduct, "indexed as", ProdductIndex)
		}
	}
}
