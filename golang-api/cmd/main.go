package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt" 
	"log"
	"net/http"
	"os" 
)

type Record struct {
	PartnerId                  string
	PartnerName                string
	CustomerId                 string
	CustomerName               string
	CustomerDomainName         string
	CustomerCountry            string
	MpnId                      string
	Tier2MpnId                 string
	InvoiceNumber              string
	ProductId                  string
	SkuId                      string
	AvailabilityId             string
	SkuName                    string
	ProductName                string
	PublisherName              string
	PublisherId                string
	SubscriptionDescription    string
	SubscriptionId             string
	ChargeStartDate            string
	ChargeEndDate              string
	UsageDate                  string
	MeterType                  string
	MeterCategory              string
	MeterId                    string
	MeterSubCategory           string
	MeterName                  string
	MeterRegion                string
	Unit                       string
	ResourceLocation           string
	ConsumedService            string
	ResourceGroup              string
	ResourceURI                string
	ChargeType                 string
	UnitPrice                  string
	Quantity                   string
	UnitType                   string
	BillingPreTaxTotal         string
	BillingCurrency            string
	PricingPreTaxTotal         string
	PricingCurrency            string
	ServiceInfo1               string
	ServiceInfo2               string
	Tags                       string
	AdditionalInfo             string
	EffectiveUnitPrice         string
	PCToBCExchangeRate         string
	PCToBCExchangeRateDate     string
	EntitlementId              string
	EntitlementDescription     string
	PartnerEarnedCreditPercentage string
	CreditPercentage           string
	CreditType                 string
	BenefitOrderId             string
	BenefitId                  string
	BenefitType                string
}

func readCSV(filename string) ([]Record, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []Record
	for _, record := range records[1:] { // Skip header row
		data = append(data, Record{
			PartnerId:                  record[0],
			PartnerName:                record[1],
			CustomerId:                 record[2],
			CustomerName:               record[3],
			CustomerDomainName:         record[4],
			CustomerCountry:            record[5],
			MpnId:                      record[6],
			Tier2MpnId:                 record[7],
			InvoiceNumber:              record[8],
			ProductId:                  record[9],
			SkuId:                      record[10],
			AvailabilityId:             record[11],
			SkuName:                    record[12],
			ProductName:                record[13],
			PublisherName:              record[14],
			PublisherId:                record[15],
			SubscriptionDescription:    record[16],
			SubscriptionId:             record[17],
			ChargeStartDate:            record[18],
			ChargeEndDate:              record[19],
			UsageDate:                  record[20],
			MeterType:                  record[21],
			MeterCategory:              record[22],
			MeterId:                    record[23],
			MeterSubCategory:           record[24],
			MeterName:                  record[25],
			MeterRegion:                record[26],
			Unit:                       record[27],
			ResourceLocation:           record[28],
			ConsumedService:            record[29],
			ResourceGroup:              record[30],
			ResourceURI:                record[31],
			ChargeType:                 record[32],
			UnitPrice:                  record[33],
			Quantity:                   record[34],
			UnitType:                   record[35],
			BillingPreTaxTotal:         record[36],
			BillingCurrency:            record[37],
			PricingPreTaxTotal:         record[38],
			PricingCurrency:            record[39],
			ServiceInfo1:               record[40],
			ServiceInfo2:               record[41],
			Tags:                       record[42],
			AdditionalInfo:             record[43],
			EffectiveUnitPrice:         record[44],
			PCToBCExchangeRate:         record[45],
			PCToBCExchangeRateDate:     record[46],
			EntitlementId:              record[47],
			EntitlementDescription:     record[48],
			PartnerEarnedCreditPercentage: record[49],
			CreditPercentage:           record[50],
			CreditType:                 record[51],
			BenefitOrderId:             record[52],
			BenefitId:                  record[53],
			BenefitType:                record[54],
		})
	}

	return data, nil
}

func writeCSV(filename string, records []Record) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	header := []string{"PartnerId", "PartnerName", "CustomerId", "CustomerName", "CustomerDomainName", "CustomerCountry", "MpnId", "Tier2MpnId", "InvoiceNumber", "ProductId", "SkuId", "AvailabilityId", "SkuName", "ProductName", "PublisherName", "PublisherId", "SubscriptionDescription", "SubscriptionId", "ChargeStartDate", "ChargeEndDate", "UsageDate", "MeterType", "MeterCategory", "MeterId", "MeterSubCategory", "MeterName", "MeterRegion", "Unit", "ResourceLocation", "ConsumedService", "ResourceGroup", "ResourceURI", "ChargeType", "UnitPrice", "Quantity", "UnitType", "BillingPreTaxTotal", "BillingCurrency", "PricingPreTaxTotal", "PricingCurrency", "ServiceInfo1", "ServiceInfo2", "Tags", "AdditionalInfo", "EffectiveUnitPrice", "PCToBCExchangeRate", "PCToBCExchangeRateDate", "EntitlementId", "EntitlementDescription", "PartnerEarnedCreditPercentage", "CreditPercentage", "CreditType", "BenefitOrderId", "BenefitId", "BenefitType"}
	writer.Write(header)

	for _, record := range records {
		writer.Write([]string{record.PartnerId, record.PartnerName, record.CustomerId, record.CustomerName, record.CustomerDomainName, record.CustomerCountry, record.MpnId, record.Tier2MpnId, record.InvoiceNumber, record.ProductId, record.SkuId, record.AvailabilityId, record.SkuName, record.ProductName, record.PublisherName, record.PublisherId, record.SubscriptionDescription, record.SubscriptionId, record.ChargeStartDate, record.ChargeEndDate, record.UsageDate, record.MeterType, record.MeterCategory, record.MeterId, record.MeterSubCategory, record.MeterName, record.MeterRegion, record.Unit, record.ResourceLocation, record.ConsumedService, record.ResourceGroup, record.ResourceURI, record.ChargeType, record.UnitPrice, record.Quantity, record.UnitType, record.BillingPreTaxTotal, record.BillingCurrency, record.PricingPreTaxTotal, record.PricingCurrency, record.ServiceInfo1, record.ServiceInfo2, record.Tags, record.AdditionalInfo, record.EffectiveUnitPrice, record.PCToBCExchangeRate, record.PCToBCExchangeRateDate, record.EntitlementId, record.EntitlementDescription, record.PartnerEarnedCreditPercentage, record.CreditPercentage, record.CreditType, record.BenefitOrderId, record.BenefitId, record.BenefitType})
	}

	return nil
}

func readHandler(w http.ResponseWriter, r *http.Request) {
	records, err := readCSV("data.csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonData, err := json.Marshal(records)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	var newRecord Record
	if err := json.NewDecoder(r.Body).Decode(&newRecord); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	records, err := readCSV("data.csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	records = append(records, newRecord)

	if err := writeCSV("data.csv", records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var updatedRecord Record
	if err := json.NewDecoder(r.Body).Decode(&updatedRecord); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	records, err := readCSV("data.csv")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	found := false
	for i, record := range records {
		if record.PartnerId == updatedRecord.PartnerId {
			records[i] = updatedRecord
			found = true
			break
		}
	}

	if !found {
		http.Error(w, "Record not found", http.StatusNotFound)
		return
	}

	if err := writeCSV("data.csv", records); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/read", readHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/update", updateHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
