package main

import (
	"time"
)

// BookingDetails - Main booking information (S1)
type BookingDetails struct {
	BookingId         *int       `db:"bookingId" json:"bookingId"`
	Refresa           *string    `db:"refresa" json:"refresa"`
	BookingRef        *string    `db:"bookingRef" json:"bookingRef"`
	BookingNumber     *string    `db:"bookingNumber" json:"bookingNumber"`
	VoucherNumber     *string    `db:"voucherNumber" json:"voucherNumber"`
	ConfirmNumber     *string    `db:"confirmNumber" json:"confirmNumber"`
	TrackingId        *string    `db:"trackingId" json:"trackingId"`
	ProcessId         *string    `db:"processId" json:"processId"`
	CustomerId        *int       `db:"customerId" json:"customerId"`
	PartnerId         *int       `db:"partnerId" json:"partnerId"`
	PartnerName       *string    `db:"partnerName" json:"partnerName"`
	AgencyId          *int       `db:"agencyId" json:"agencyId"`
	AgencyName        *string    `db:"agencyName" json:"agencyName"`
	LangId            *int       `db:"langId" json:"langId"`
	BookingDate       *time.Time `db:"bookingDate" json:"bookingDate"`
	PaymentDeadLine   *time.Time `db:"paymentDeadLine" json:"paymentDeadLine"`
	StatusId          *int       `db:"statusId" json:"statusId"`
	PaymentStatusId   *int       `db:"paymentStatusId" json:"paymentStatusId"`
	UserId            *int       `db:"userId" json:"userId"`
	AgentId           *int       `db:"agentId" json:"agentId"`
	Status            *string    `db:"Status" json:"status"`
	PaymentStatus     *string    `db:"paymentStatus" json:"paymentStatus"`
	UserName          *string    `db:"userName" json:"userName"`
	AgentName         *string    `db:"agentName" json:"agentName"`
	ProductName       *string    `db:"productName" json:"productName"`
	HotelId           *int       `db:"hotelId" json:"hotelId"`
	ProductId         *int       `db:"productId" json:"productId"`
	Destination       *string    `db:"Destination" json:"destination"`
	HotelAddress      *string    `db:"hotelAddress" json:"hotelAddress"`
	HotelPhoneNumber  *string    `db:"hotelPhoneNumber" json:"hotelPhoneNumber"`
	DepDate           *time.Time `db:"depDate" json:"depDate"`
	RetDate           *time.Time `db:"retDate" json:"retDate"`
	Days              *int       `db:"days" json:"days"`
	Nights            *int       `db:"nights" json:"nights"`
	DepCity           *string    `db:"depCity" json:"depCity"`
	Total             *float64   `db:"Total" json:"total"`
	TotalCost         *float64   `db:"totalCost" json:"totalCost"`
	ProfitMargin      *float64   `db:"profitMargin" json:"profitMargin"`
	ResellerMargin    *float64   `db:"resellerMargin" json:"resellerMargin"`
	DistributorMargin *float64   `db:"distributorMargin" json:"distributorMargin"`
	Commission        *float64   `db:"Commission" json:"commission"`
	Currency          *string    `db:"currency" json:"currency"`
	DisplayCurrency   *string    `db:"displayCurrency" json:"displayCurrency"`
	ExchangeRate      *float64   `db:"exchangeRate" json:"exchangeRate"`
	Title             *string    `db:"title" json:"title"`
	LastName          *string    `db:"lastName" json:"lastName"`
	FirstName         *string    `db:"firstName" json:"firstName"`
	CustomerAddress   *string    `db:"customerAddress" json:"customerAddress"`
	CustomerCity      *string    `db:"customerCity" json:"customerCity"`
	CustomerCountry   *string    `db:"customerCountry" json:"customerCountry"`
	Email             *string    `db:"email" json:"email"`
	Mobile            *string    `db:"mobile" json:"mobile"`
	Nationality       *string    `db:"nationality" json:"nationality"`
	ProfessionalId    *int       `db:"professionalId" json:"professionalId"`
	ProductTypeId     *int       `db:"productTypeId" json:"productTypeId"`
	ProductType       *string    `db:"productType" json:"productType"`
	CancellationFee   *float64   `db:"cancellationFee" json:"cancellationFee"`
	AmountToPay       *float64   `db:"amountToPay" json:"amountToPay"`
	SupplierId        *int       `db:"supplierId" json:"supplierId"`
	Supplier          *string    `db:"supplier" json:"supplier"`
}

// HotelDetail - Hotel room details (S2)
type HotelDetail struct {
	IdDetailHotel *int       `db:"IdDetailHotel" json:"idDetailHotel"`
	HotelName     *string    `db:"hotelName" json:"hotelName"`
	Country       *string    `db:"country" json:"country"`
	City          *string    `db:"city" json:"city"`
	ArrDate       *time.Time `db:"arrDate" json:"arrDate"`
	DepDate       *time.Time `db:"depDate" json:"depDate"`
	RoomName      *string    `db:"roomName" json:"roomName"`
	RoomId        *int       `db:"roomId" json:"roomId"`
	PensionId     *int       `db:"pensionId" json:"pensionId"`
	Quantity      *int       `db:"quantity" json:"quantity"`
	Board         *string    `db:"Board" json:"board"`
	Adults        *int       `db:"adults" json:"adults"`
	Children      *int       `db:"children" json:"children"`
	Infant        *int       `db:"infant" json:"infant"`
	Ages          *string    `db:"Ages" json:"ages"`
	PurchasePrice *float64   `db:"PurchasePrice" json:"purchasePrice"`
	SalePrice     *float64   `db:"SalePrice" json:"salePrice"`
	HotelId       *int       `db:"hotelId" json:"hotelId"`
	ProductId     *int       `db:"productId" json:"productId"`
	NumChambre    *int       `db:"NumChambre" json:"numChambre"`
	RateKey       *string    `db:"rateKey" json:"rateKey"`
}

// DetailOption - Booking options (S3)
type DetailOption struct {
	IdDetailOption *int     `db:"IdDetailOption" json:"idDetailOption"`
	BookingId      *int     `db:"bookingId" json:"bookingId"`
	IdTypeProduit  *int     `db:"IdTypeProduit" json:"idTypeProduit"`
	Libelle        *string  `db:"Libelle" json:"libelle"`
	Description    *string  `db:"Description" json:"description"`
	Nombre         *int     `db:"Nombre" json:"nombre"`
	Quantity       *int     `db:"Quantity" json:"quantity"`
	PurchasePrice  *float64 `db:"PurchasePrice" json:"purchasePrice"`
	SalePrice      *float64 `db:"SalePrice" json:"salePrice"`
}

// BookingService - Services (S4)
type BookingService struct {
	BookingServiceId *int     `db:"bookingServiceId" json:"bookingServiceId"`
	BookingId        *int     `db:"bookingId" json:"bookingId"`
	ProductTypeId    *int     `db:"productTypeId" json:"productTypeId"`
	ProductType      *string  `db:"productType" json:"productType"`
	ServiceName      *string  `db:"serviceName" json:"serviceName"`
	Quantity         *int     `db:"quantity" json:"quantity"`
	Price            *float64 `db:"price" json:"price"`
}

// Miscellaneous - Misc items (S5)
type Miscellaneous struct {
	IdDetailOption *int    `db:"IdDetailOption" json:"idDetailOption"`
	Libelle        *string `db:"Libelle" json:"libelle"`
	Nombre         *int    `db:"Nombre" json:"nombre"`
	Quantity       *int    `db:"Quantity" json:"quantity"`
	Description    *string `db:"Description" json:"description"`
}

// Passenger - Passenger details (S6)
type Passenger struct {
	IdPassager *int       `db:"IdPassager" json:"idPassager"`
	RefResa    *string    `db:"refResa" json:"refResa"`
	BookingId  *int       `db:"bookingId" json:"bookingId"`
	Title      *string    `db:"title" json:"title"`
	LastName   *string    `db:"lastName" json:"lastName"`
	FirstName  *string    `db:"firstName" json:"firstName"`
	Age        *int       `db:"Age" json:"age"`
	BirthDate  *time.Time `db:"birthDate" json:"birthDate"`
	PaxType    *string    `db:"paxType" json:"paxType"`
	IdProduit  *int       `db:"IdProduit" json:"idProduit"`
	NumChambre *int       `db:"NumChambre" json:"numChambre"`
	RoomId     *int       `db:"roomId" json:"roomId"`
}

// BookingSummary - Room/passenger summary (S7)
type BookingSummary struct {
	NumberRooms    *int `db:"number_rooms" json:"numberRooms"`
	NumberAdults   *int `db:"number_adults" json:"numberAdults"`
	NumberChildren *int `db:"number_children" json:"numberChildren"`
	NumberInfant   *int `db:"number_infant" json:"numberInfant"`
}

// BookingHotel - Hotel information (S8)
type BookingHotel struct {
	BookingHotelID *int       `db:"bookingHotelID" json:"bookingHotelID"`
	HotelId        *int       `db:"hotelId" json:"hotelId"`
	HotelName      *string    `db:"hotelName" json:"hotelName"`
	Country        *string    `db:"country" json:"country"`
	City           *string    `db:"city" json:"city"`
	Address        *string    `db:"address" json:"address"`
	BookingId      *int       `db:"bookingId" json:"bookingId"`
	ArrDate        *time.Time `db:"arrDate" json:"arrDate"`
	DepDate        *time.Time `db:"depDate" json:"depDate"`
	Rating         *int       `db:"rating" json:"rating"`
}

// BookingProduct - Product details (S9)
type BookingProduct struct {
	BookingProductId *int       `db:"bookingProductId" json:"bookingProductId"`
	ProductId        *int       `db:"productId" json:"productId"`
	ProductName      *string    `db:"productName" json:"productName"`
	Destination      *string    `db:"destination" json:"destination"`
	ProductTypeId    *int       `db:"productTypeId" json:"productTypeId"`
	ProductType      *string    `db:"productType" json:"productType"`
	Photo            *string    `db:"photo" json:"photo"`
	BookingId        *int       `db:"bookingId" json:"bookingId"`
	DepDate          *time.Time `db:"depDate" json:"depDate"`
	RetDate          *time.Time `db:"retDate" json:"retDate"`
	Days             *int       `db:"days" json:"days"`
	Nights           *int       `db:"nights" json:"nights"`
	Adults           *int       `db:"adults" json:"adults"`
	Children         *int       `db:"children" json:"children"`
	Infants          *int       `db:"infants" json:"infants"`
	Ages             *string    `db:"ages" json:"ages"`
	SalePrice        *float64   `db:"salePrice" json:"salePrice"`
	PurchasePrice    *float64   `db:"purchasePrice" json:"purchasePrice"`
	ConfirmNumber    *string    `db:"confirmNumber" json:"confirmNumber"`
	RefererBookingId *int       `db:"refererBookingId" json:"refererBookingId"`
}

// Transaction - Payment transactions (S10)
type Transaction struct {
	TransactionId *int       `db:"transactionId" json:"transactionId"`
	SId           *string    `db:"sId" json:"sId"`
	AffilieId     *int       `db:"affilieId" json:"affilieId"`
	NumCde        *string    `db:"NumCde" json:"numCde"`
	PaymentRef    *string    `db:"paymentRef" json:"paymentRef"`
	PaymentRef2   *string    `db:"paymentRef2" json:"paymentRef2"`
	Amount        *float64   `db:"amount" json:"amount"`
	BookingRef    *string    `db:"bookingRef" json:"bookingRef"`
	BookingId     *int       `db:"bookingId" json:"bookingId"`
	CreateDate    *time.Time `db:"createDate" json:"createDate"`
	UpdateDate    *time.Time `db:"updateDate" json:"updateDate"`
	PaymentMethod *string    `db:"paymentMethod" json:"paymentMethod"`
	Currency      *string    `db:"currency" json:"currency"`
	StatusId      *int       `db:"statusId" json:"statusId"`
	Terminal      *string    `db:"terminal" json:"terminal"`
	Action        *string    `db:"action" json:"action"`
	JsonConfirm   *string    `db:"jsonConfirm" json:"jsonConfirm"`
}

// CancellationPolicy - Cancellation policies (S11)
type CancellationPolicy struct {
	ID            *int       `db:"ID" json:"id"`
	DateFrom      *time.Time `db:"dateFrom" json:"dateFrom"`
	DateTo        *time.Time `db:"dateTo" json:"dateTo"`
	FeeType       *string    `db:"feeType" json:"feeType"`
	Value         *float64   `db:"value" json:"value"`
	Remarks       *string    `db:"remarks" json:"remarks"`
	BookingId     *int       `db:"bookingId" json:"bookingId"`
	Currency      *string    `db:"currency" json:"currency"`
	PurchasePrice *float64   `db:"purchasePrice" json:"purchasePrice"`
	SalePrice     *float64   `db:"salePrice" json:"salePrice"`
}

// BookingLog - Booking logs (S12)
type BookingLog struct {
	LogId     *int       `db:"logId" json:"logId"`
	Texte     *string    `db:"texte" json:"texte"`
	Date      *time.Time `db:"date" json:"date"`
	UserId    *int       `db:"userId" json:"userId"`
	AgentName *string    `db:"agentName" json:"agentName"`
	BookingId *int       `db:"bookingId" json:"bookingId"`
	Alias     *string    `db:"alias" json:"alias"`
}

// BookingDetailsResult - Contains all recordsets
type BookingDetailsResult struct {
	BookingDetails       []BookingDetails     `json:"bookingDetails"`
	HotelDetails         []HotelDetail        `json:"hotelDetails"`
	DetailOptions        []DetailOption       `json:"detailOptions"`
	BookingServices      []BookingService     `json:"bookingServices"`
	Miscellaneous        []Miscellaneous      `json:"miscellaneous"`
	Passengers           []Passenger          `json:"passengers"`
	BookingSummary       []BookingSummary     `json:"bookingSummary"`
	BookingHotels        []BookingHotel       `json:"bookingHotels"`
	BookingProducts      []BookingProduct     `json:"bookingProducts"`
	Transactions         []Transaction        `json:"transactions"`
	CancellationPolicies []CancellationPolicy `json:"cancellationPolicies"`
	BookingLogs          []BookingLog         `json:"bookingLogs"`
}
