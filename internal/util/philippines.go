package util

import "github.com/brianvoe/gofakeit/v7"

var provinces = []string{
	"Abra", "Agusan del Norte", "Agusan del Sur", "Aklan", "Albay",
	"Antique", "Apayao", "Aurora", "Basilan", "Bataan", "Batanes",
	"Batangas", "Benguet", "Biliran", "Bohol", "Bukidnon", "Bulacan",
	"Cagayan", "Camarines Norte", "Camarines Sur", "Camiguin", "Capiz",
	"Catanduanes", "Cavite", "Cebu", "Cotabato", "Davao de Oro",
	"Davao del Norte", "Davao del Sur", "Davao Occidental",
	"Davao Oriental", "Dinagat Islands", "Eastern Samar", "Guimaras",
	"Ifugao", "Ilocos Norte", "Ilocos Sur", "Iloilo", "Isabela", "Kalinga",
	"La Union", "Laguna", "Lanao del Norte", "Lanao del Sur", "Leyte",
	"Maguindanao del Norte", "Maguindanao del Sur", "Marinduque",
	"Masbate", "Metro Manila", "Misamis Occidental", "Misamis Oriental",
	"Mountain Province", "Negros Occidental", "Negros Oriental",
	"Northern Samar", "Nueva Ecija", "Nueva Vizcaya", "Occidental Mindoro",
	"Oriental Mindoro", "Palawan", "Pampanga", "Pangasinan", "Quezon",
	"Quirino", "Rizal", "Romblon", "Samar", "Sarangani", "Siquijor",
	"Sorsogon", "South Cotabato", "Southern Leyte", "Sultan Kudarat",
	"Sulu", "Surigao del Norte", "Surigao del Sur", "Tarlac", "Tawi-Tawi",
	"Zambales", "Zamboanga del Norte", "Zamboanga del Sur",
	"Zamboanga Sibugay",
}

type Province string

func (c *Province) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString(provinces), nil
}

var regionsNumerical = []string{
	"I", "II", "III", "IV-A", "IV-B", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "NCR", "CAR", "BARMM",
}

type RegionNumerical string

func (c *RegionNumerical) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString(regionsNumerical), nil
}

var regionsNames = []string{
	"Ilocos Region", "Cagayan Valley", "Central Luzon", "Calabarzon", "Mimaropa", "Bicol Region",
	"Western Visayas", "Central Visayas", "Eastern Visayas", "Zamboanga Peninsula", "Northern Mindanao",
	"Davao Region", "Soccsksargen", "Caraga", "National Capital Region", "Cordillera Administrative Region",
	"Bangsamoro Autonomous Region in Muslim Mindanao",
}

type Region string

func (c *Region) Fake(f *gofakeit.Faker) (any, error) {
	return f.RandomString(regionsNames), nil
}
