package mock

import (
	"fmt"
	"math/rand"
	"time"

	pb "srv-eazle-advise-mock/pkg/gen/proto/outlet"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type MockSettings struct {
	AverageNotesList               int `json:"averageNotesList,omitempty"`
	AverageVisitHistory            int `json:"averageVisitHistory,omitempty"`
	AverageNumberOfOrders          int `json:"averageNumberOfOrders,omitempty"`
	AverageOrderItemsPerOrder      int `json:"averageOrderItemsPerOrder,omitempty"`
	AverageTopProductsInStatistics int `json:"averageTopProductsInStatistics,omitempty"`
	AverageOutletsNearby           int `json:"averageOutletsNearby,omitempty"`
	AverageAssetList               int `json:"averageAssetList,omitempty"`
	AverageChecklist               int `json:"averageChecklist,omitempty"`
	AverageNews                    int `json:"averageNews,omitempty"`
}

type MockSettingsOptional struct {
	AverageNotesList               *int `json:"averageNotesList,omitempty"`
	AverageVisitHistory            *int `json:"averageVisitHistory,omitempty"`
	AverageNumberOfOrders          *int `json:"averageNumberOfOrders,omitempty"`
	AverageOrderItemsPerOrder      *int `json:"averageOrderItemsPerOrder,omitempty"`
	AverageTopProductsInStatistics *int `json:"averageTopProductsInStatistics,omitempty"`
	AverageOutletsNearby           *int `json:"averageOutletsNearby,omitempty"`
	AverageAssetList               *int `json:"averageAssetList,omitempty"`
	AverageChecklist               *int `json:"averageChecklist,omitempty"`
	AverageNews                    *int `json:"averageNews,omitempty"`
}

// Sample data pools for randomization
var (
	outletNames = []string{
		"SuperMart Downtown", "MegaStore Central", "QuickMart Express", "FamilyShop Plus",
		"GroceryWorld", "FreshMart Deluxe", "CityStore Premium", "NeighborhoodMart",
		"ValueMart Pro", "LocalShop Central", "MarketPlace Elite", "ShopRite Corner",
	}

	storeManagers = []string{
		"John Smith", "Sarah Johnson", "Mike Wilson", "Emily Davis",
		"Robert Brown", "Lisa Anderson", "David Miller", "Jennifer Garcia",
		"Michael Taylor", "Amanda Wilson", "Christopher Lee", "Jessica Martinez",
	}

	salesReps = []string{
		"Mike Wilson", "Sarah Thompson", "Alex Rodriguez", "Maria Garcia",
		"James Brown", "Nicole Taylor", "Kevin Davis", "Rachel Martinez",
		"Daniel Lee", "Amanda Clark", "Steven White", "Michelle Lopez",
	}

	productNames = []string{
		"Premium Cola 24-pack", "Organic Chips Variety Pack", "Energy Drink Mix",
		"Sparkling Water Cases", "Protein Bars Box", "Healthy Snack Mix",
		"Juice Bottle Set", "Coffee Bean Bags", "Tea Collection Box",
		"Vitamin Water Pack", "Sports Drink Cases", "Smoothie Mix Packets",
	}

	cities = []string{
		"Metro City", "Downtown Plaza", "Central District", "Uptown Area",
		"Riverside", "Hillside", "Lakeside", "Parkview", "Westside", "Eastgate",
	}

	noteTypes = []pb.NoteType{
		pb.NoteType_NOTE_TYPE_GENERAL,
		pb.NoteType_NOTE_TYPE_SALES,
		pb.NoteType_NOTE_TYPE_SUPPORT,
		pb.NoteType_NOTE_TYPE_COMPLAINT,
		pb.NoteType_NOTE_TYPE_OPPORTUNITY,
		pb.NoteType_NOTE_TYPE_REMINDER,
	}

	assetTypes = []pb.AssetType{
		pb.AssetType_ASSET_TYPE_FREEZER,
		pb.AssetType_ASSET_TYPE_REFRIGERATOR,
		pb.AssetType_ASSET_TYPE_DISPLAY_UNIT,
		pb.AssetType_ASSET_TYPE_SIGNAGE,
		pb.AssetType_ASSET_TYPE_POS_SYSTEM,
		pb.AssetType_ASSET_TYPE_SHELVING,
	}

	checklistCategories = []pb.ChecklistCategory{
		pb.ChecklistCategory_CHECKLIST_CATEGORY_COMPLIANCE,
		pb.ChecklistCategory_CHECKLIST_CATEGORY_SAFETY,
		pb.ChecklistCategory_CHECKLIST_CATEGORY_QUALITY,
		pb.ChecklistCategory_CHECKLIST_CATEGORY_INVENTORY,
		pb.ChecklistCategory_CHECKLIST_CATEGORY_MAINTENANCE,
		pb.ChecklistCategory_CHECKLIST_CATEGORY_MARKETING,
	}
)

func GenerateMockedOutlet(outletID string, settings MockSettings) *pb.OutletDetails {
	rand.Seed(time.Now().UnixNano())
	now := timestamppb.Now()

	outlet := &pb.OutletDetails{
		OutletId:  outletID,
		Name:      randomChoice(outletNames),
		Code:      fmt.Sprintf("ST-%s-%03d", randomString(2), rand.Intn(999)+1),
		Thumbnail: "https://picsum.photos/100",
		Type:      randomOutletType(),
		Status:    pb.OutletStatus_OUTLET_STATUS_ACTIVE,
		Location:  generateRandomLocation(),
		CreatedAt: timestamppb.New(time.Now().AddDate(-rand.Intn(3)-1, -rand.Intn(12), -rand.Intn(30))),
		UpdatedAt: now,
	}

	// Generate contact points (always 1-3)
	outlet.ContactPoints = generateContactPoints(rand.Intn(3) + 1)

	// Generate visit history
	if settings.AverageVisitHistory > 0 {
		visitCount := randomizeCount(settings.AverageVisitHistory)
		outlet.VisitHistory = generateVisitHistory(visitCount)
	}

	// Generate order history
	if settings.AverageNumberOfOrders > 0 {
		orderCount := randomizeCount(settings.AverageNumberOfOrders)
		outlet.OrderHistory = generateOrderHistory(orderCount, settings.AverageOrderItemsPerOrder)
	}

	// Generate statistics
	outlet.Statistics = generateStatistics(settings.AverageTopProductsInStatistics, outlet.OrderHistory, outlet.VisitHistory)

	// Generate nearby outlets
	if settings.AverageOutletsNearby > 0 {
		nearbyCount := randomizeCount(settings.AverageOutletsNearby)
		outlet.OutletsNearby = generateNearbyOutlets(nearbyCount, outlet.Location)
	}

	// Generate notes
	if settings.AverageNotesList > 0 {
		notesCount := randomizeCount(settings.AverageNotesList)
		outlet.Notes = generateNotes(notesCount)
	}

	// Generate assets
	if settings.AverageAssetList > 0 {
		assetCount := randomizeCount(settings.AverageAssetList)
		outlet.AssetList = generateAssets(assetCount)
	}

	// Generate checklist
	if settings.AverageChecklist > 0 {
		checklistCount := randomizeCount(settings.AverageChecklist)
		outlet.Checklist = generateChecklist(checklistCount)
	}

	// Generate news
	if settings.AverageNews > 0 {
		newsCount := randomizeCount(settings.AverageNews)
		outlet.News = generateNews(newsCount)
	}

	return outlet
}

func randomizeCount(average int) int {
	if average <= 1 {
		return average
	}
	// Generate count with ±50% variance from average
	min := average / 2
	max := average + (average / 2)
	return rand.Intn(max-min+1) + min
}

func randomChoice(slice []string) string {
	if len(slice) == 0 {
		return ""
	}
	return slice[rand.Intn(len(slice))]
}

func randomString(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

func randomOutletType() pb.OutletType {
	types := []pb.OutletType{
		pb.OutletType_OUTLET_TYPE_RETAIL,
		pb.OutletType_OUTLET_TYPE_WHOLESALE,
		pb.OutletType_OUTLET_TYPE_PHARMACY,
		pb.OutletType_OUTLET_TYPE_SUPERMARKET,
		pb.OutletType_OUTLET_TYPE_CONVENIENCE_STORE,
		pb.OutletType_OUTLET_TYPE_RESTAURANT,
	}
	return types[rand.Intn(len(types))]
}

func generateRandomLocation() *pb.Location {
	city := randomChoice(cities)
	return &pb.Location{
		Address:    fmt.Sprintf("%d %s Street, %s", rand.Intn(999)+1, randomChoice([]string{"Main", "Oak", "Pine", "Elm", "First", "Second"}), city),
		City:       city,
		State:      "Central State",
		PostalCode: fmt.Sprintf("%05d", rand.Intn(99999)),
		Country:    "Country Name",
		Latitude:   40.7128 + (rand.Float64()-0.5)*0.1,
		Longitude:  -74.0060 + (rand.Float64()-0.5)*0.1,
	}
}

func generateContactPoints(count int) []*pb.ContactPoint {
	contacts := make([]*pb.ContactPoint, count)
	for i := 0; i < count; i++ {
		name := randomChoice(storeManagers)
		contacts[i] = &pb.ContactPoint{
			ContactId: fmt.Sprintf("contact-%03d", i+1),
			Name:      name,
			Role:      randomChoice([]string{"Store Manager", "Assistant Manager", "Buyer", "Operations Manager"}),
			Phone:     fmt.Sprintf("+1-555-%04d", rand.Intn(9999)),
			Email:     fmt.Sprintf("%s@store.com", randomString(8)),
			Type:      pb.ContactType_CONTACT_TYPE_MANAGER,
			IsPrimary: i == 0, // First contact is primary
			CreatedAt: timestamppb.New(time.Now().AddDate(0, -rand.Intn(12), -rand.Intn(30))),
		}
	}
	return contacts
}

func generateVisitHistory(count int) []*pb.Visit {
	visits := make([]*pb.Visit, count)
	for i := 0; i < count; i++ {
		visitDate := timestamppb.New(time.Now().AddDate(0, 0, -(rand.Intn(365))))
		visits[i] = &pb.Visit{
			VisitId:           fmt.Sprintf("visit-%03d", i+1),
			SalesRepId:        fmt.Sprintf("rep-%03d", rand.Intn(10)+1),
			SalesRepName:      randomChoice(salesReps),
			VisitDate:         visitDate,
			VisitType:         randomVisitType(),
			VisitStatus:       pb.VisitStatus_VISIT_STATUS_COMPLETED,
			Purpose:           randomChoice([]string{"Product presentation", "Order discussion", "Customer support", "Inventory check", "Relationship building"}),
			Summary:           fmt.Sprintf("Visit completed successfully. %s", randomChoice([]string{"Client showed interest in new products.", "Discussed upcoming promotions.", "Resolved customer concerns.", "Planned next steps."})),
			ProductsDiscussed: []string{randomChoice(productNames), randomChoice(productNames)},
			ActionsTaken:      generateVisitActions(rand.Intn(2) + 1),
			Attachments:       []string{fmt.Sprintf("document_%d.pdf", i+1)},
			DurationSeconds:   int32(rand.Intn(3600) + 1800), // 30 minutes to 2 hours
		}
	}
	return visits
}

func randomVisitType() pb.VisitType {
	types := []pb.VisitType{
		pb.VisitType_VISIT_TYPE_SALES_CALL,
		pb.VisitType_VISIT_TYPE_DELIVERY,
		pb.VisitType_VISIT_TYPE_SUPPORT,
		pb.VisitType_VISIT_TYPE_AUDIT,
		pb.VisitType_VISIT_TYPE_TRAINING,
	}
	return types[rand.Intn(len(types))]
}

func generateVisitActions(count int) []*pb.VisitAction {
	actions := make([]*pb.VisitAction, count)
	for i := 0; i < count; i++ {
		actions[i] = &pb.VisitAction{
			ActionId:    fmt.Sprintf("action-%03d", i+1),
			Description: randomChoice([]string{"Follow up on pricing", "Schedule product demo", "Negotiate terms", "Arrange delivery", "Collect payment"}),
			Type:        randomActionType(),
			Status:      randomActionStatus(),
			DueDate:     timestamppb.New(time.Now().AddDate(0, 0, rand.Intn(30)+1)),
		}
	}
	return actions
}

func randomActionType() pb.ActionType {
	types := []pb.ActionType{
		pb.ActionType_ACTION_TYPE_FOLLOW_UP,
		pb.ActionType_ACTION_TYPE_PRODUCT_DEMO,
		pb.ActionType_ACTION_TYPE_PRICE_NEGOTIATION,
		pb.ActionType_ACTION_TYPE_DELIVERY_SCHEDULE,
		pb.ActionType_ACTION_TYPE_PAYMENT_COLLECTION,
	}
	return types[rand.Intn(len(types))]
}

func randomActionStatus() pb.ActionStatus {
	statuses := []pb.ActionStatus{
		pb.ActionStatus_ACTION_STATUS_PENDING,
		pb.ActionStatus_ACTION_STATUS_IN_PROGRESS,
		pb.ActionStatus_ACTION_STATUS_COMPLETED,
	}
	return statuses[rand.Intn(len(statuses))]
}

func generateOrderHistory(count int, avgItemsPerOrder int) []*pb.Order {
	orders := make([]*pb.Order, count)
	for i := 0; i < count; i++ {
		orderDate := timestamppb.New(time.Now().AddDate(0, 0, -(rand.Intn(365))))
		itemCount := randomizeCount(avgItemsPerOrder)
		if itemCount == 0 {
			itemCount = 1
		}

		items := generateOrderItems(itemCount)
		totalAmount := calculateOrderTotal(items)

		orders[i] = &pb.Order{
			OrderId:      fmt.Sprintf("order-%03d", i+1),
			OrderNumber:  fmt.Sprintf("ORD-2024-%06d", rand.Intn(999999)+1),
			OrderDate:    orderDate,
			Status:       randomOrderStatus(),
			TotalAmount:  totalAmount,
			Currency:     "USD",
			Items:        items,
			PaymentInfo:  generatePaymentInfo(totalAmount, orderDate),
			DeliveryInfo: generateDeliveryInfo(orderDate),
			SalesRepId:   fmt.Sprintf("rep-%03d", rand.Intn(10)+1),
			SalesRepName: randomChoice(salesReps),
			DeliveryDate: timestamppb.New(orderDate.AsTime().AddDate(0, 0, rand.Intn(7)+1)),
			Notes:        randomChoice([]string{"Standard delivery", "Express shipping", "Customer pickup", "Special instructions followed"}),
		}
	}
	return orders
}

func generateOrderItems(count int) []*pb.OrderItem {
	items := make([]*pb.OrderItem, count)
	for i := 0; i < count; i++ {
		unitPrice := float64(rand.Intn(50)+5) + rand.Float64()
		quantity := int32(rand.Intn(100) + 1)
		discountPct := float64(rand.Intn(20))
		totalPrice := float64(quantity) * unitPrice
		discountAmount := totalPrice * (discountPct / 100)

		items[i] = &pb.OrderItem{
			ProductId:          fmt.Sprintf("prod-%03d", rand.Intn(100)+1),
			ProductName:        randomChoice(productNames),
			Sku:                fmt.Sprintf("%s-%03d", randomString(3), rand.Intn(999)+1),
			Quantity:           quantity,
			UnitPrice:          unitPrice,
			TotalPrice:         totalPrice - discountAmount,
			DiscountPercentage: discountPct,
			DiscountAmount:     discountAmount,
		}
	}
	return items
}

func calculateOrderTotal(items []*pb.OrderItem) float64 {
	total := 0.0
	for _, item := range items {
		total += item.TotalPrice
	}
	return total
}

func randomOrderStatus() pb.OrderStatus {
	statuses := []pb.OrderStatus{
		pb.OrderStatus_ORDER_STATUS_DELIVERED,
		pb.OrderStatus_ORDER_STATUS_SHIPPED,
		pb.OrderStatus_ORDER_STATUS_PROCESSING,
		pb.OrderStatus_ORDER_STATUS_CONFIRMED,
	}
	return statuses[rand.Intn(len(statuses))]
}

func generatePaymentInfo(amount float64, orderDate *timestamppb.Timestamp) *pb.PaymentInfo {
	return &pb.PaymentInfo{
		Method:          randomPaymentMethod(),
		Status:          pb.PaymentStatus_PAYMENT_STATUS_PAID,
		PaymentDate:     orderDate,
		AmountPaid:      amount,
		AmountDue:       0.0,
		ReferenceNumber: fmt.Sprintf("PAY-2024-%06d", rand.Intn(999999)+1),
	}
}

func randomPaymentMethod() pb.PaymentMethod {
	methods := []pb.PaymentMethod{
		pb.PaymentMethod_PAYMENT_METHOD_CASH,
		pb.PaymentMethod_PAYMENT_METHOD_CREDIT_CARD,
		pb.PaymentMethod_PAYMENT_METHOD_BANK_TRANSFER,
		pb.PaymentMethod_PAYMENT_METHOD_CREDIT,
	}
	return methods[rand.Intn(len(methods))]
}

func generateDeliveryInfo(orderDate *timestamppb.Timestamp) *pb.DeliveryInfo {
	scheduledDate := timestamppb.New(orderDate.AsTime().AddDate(0, 0, rand.Intn(5)+1))
	return &pb.DeliveryInfo{
		DeliveryAddress: fmt.Sprintf("%d %s Street", rand.Intn(999)+1, randomChoice([]string{"Main", "Oak", "Pine", "Elm"})),
		ScheduledDate:   scheduledDate,
		ActualDate:      scheduledDate,
		Status:          pb.DeliveryStatus_DELIVERY_STATUS_DELIVERED,
		DeliveryNotes:   "Delivery completed successfully",
		TrackingNumber:  fmt.Sprintf("TRK-2024-%06d", rand.Intn(999999)+1),
	}
}

func generateStatistics(topProductsCount int, orders []*pb.Order, visits []*pb.Visit) *pb.OutletStatistics {
	totalRevenue := 0.0
	for _, order := range orders {
		totalRevenue += order.TotalAmount
	}

	avgOrderValue := 0.0
	if len(orders) > 0 {
		avgOrderValue = totalRevenue / float64(len(orders))
	}

	return &pb.OutletStatistics{
		TotalRevenueYtd:         totalRevenue,
		TotalRevenueLastYear:    totalRevenue * (0.7 + rand.Float64()*0.6), // 70-130% of current
		AverageOrderValue:       avgOrderValue,
		TotalOrdersYtd:          int32(len(orders)),
		TotalOrdersLastYear:     int32(float64(len(orders)) * (0.7 + rand.Float64()*0.6)),
		TotalVisitsYtd:          int32(len(visits)),
		RevenueGrowthPercentage: rand.Float64()*50 - 10, // -10% to +40%
		DaysSinceLastOrder:      int32(rand.Intn(30)),
		DaysSinceLastVisit:      int32(rand.Intn(30)),
		TopProducts:             generateTopProducts(topProductsCount),
		MonthlyRevenue:          generateMonthlyRevenue(),
		Segment:                 randomCustomerSegment(),
		CreditInfo:              generateCreditInfo(),
	}
}

func generateTopProducts(count int) []*pb.ProductStatistics {
	if count == 0 {
		return nil
	}

	products := make([]*pb.ProductStatistics, count)
	for i := 0; i < count; i++ {
		quantitySold := int32(rand.Intn(1000) + 100)
		revenue := float64(quantitySold) * (rand.Float64()*20 + 5)

		products[i] = &pb.ProductStatistics{
			ProductId:    fmt.Sprintf("prod-%03d", i+1),
			ProductName:  randomChoice(productNames),
			QuantitySold: quantitySold,
			Revenue:      revenue,
			OrdersCount:  int32(rand.Intn(20) + 5),
		}
	}
	return products
}

func generateMonthlyRevenue() []*pb.MonthlyRevenue {
	months := make([]*pb.MonthlyRevenue, 6)
	for i := 0; i < 6; i++ {
		months[i] = &pb.MonthlyRevenue{
			Year:        2024,
			Month:       int32(i + 1),
			Revenue:     rand.Float64()*5000 + 5000, // 5k-10k
			OrdersCount: int32(rand.Intn(10) + 3),
		}
	}
	return months
}

func randomCustomerSegment() pb.CustomerSegment {
	segments := []pb.CustomerSegment{
		pb.CustomerSegment_CUSTOMER_SEGMENT_BRONZE,
		pb.CustomerSegment_CUSTOMER_SEGMENT_SILVER,
		pb.CustomerSegment_CUSTOMER_SEGMENT_GOLD,
		pb.CustomerSegment_CUSTOMER_SEGMENT_PLATINUM,
	}
	return segments[rand.Intn(len(segments))]
}

func generateCreditInfo() *pb.CreditInfo {
	creditLimit := float64(rand.Intn(50000) + 10000)
	creditUsed := creditLimit * (rand.Float64() * 0.7) // Up to 70% used

	return &pb.CreditInfo{
		CreditLimit:      creditLimit,
		CreditUsed:       creditUsed,
		CreditAvailable:  creditLimit - creditUsed,
		PaymentTermsDays: int32(rand.Intn(60) + 15), // 15-75 days
		Status:           pb.CreditStatus_CREDIT_STATUS_GOOD,
	}
}

func generateNearbyOutlets(count int, baseLocation *pb.Location) []*pb.OutletNearby {
	outlets := make([]*pb.OutletNearby, count)
	for i := 0; i < count; i++ {
		outlets[i] = &pb.OutletNearby{
			OutletId:     fmt.Sprintf("nearby-outlet-%03d", i+1),
			Name:         randomChoice(outletNames),
			Type:         randomOutletType(),
			DistanceKm:   rand.Float64()*5 + 0.1, // 0.1 to 5.1 km
			Location:     generateNearbyLocation(baseLocation),
			IsCompetitor: rand.Float64() < 0.3, // 30% chance of competitor
			Relationship: randomChoice([]string{"Partner", "Competitor", "Neutral", "Supplier"}),
			Thumbnail:    "https://picsum.photos/100",
		}
	}
	return outlets
}

func generateNearbyLocation(base *pb.Location) *pb.Location {
	return &pb.Location{
		Address:   fmt.Sprintf("%d %s Avenue", rand.Intn(999)+1, randomChoice([]string{"Park", "Hill", "River", "Lake"})),
		City:      base.City,
		State:     base.State,
		Latitude:  base.Latitude + (rand.Float64()-0.5)*0.05,
		Longitude: base.Longitude + (rand.Float64()-0.5)*0.05,
	}
}

func generateNotes(count int) []*pb.Note {
	notes := make([]*pb.Note, count)
	for i := 0; i < count; i++ {
		createdAt := timestamppb.New(time.Now().AddDate(0, 0, -(rand.Intn(90))))

		notes[i] = &pb.Note{
			NoteId:    fmt.Sprintf("note-%03d", i+1),
			Title:     randomChoice([]string{"Customer Feedback", "Sales Opportunity", "Support Issue", "Follow-up Required", "Payment Discussion"}),
			Content:   fmt.Sprintf("Note content %d: %s", i+1, randomChoice([]string{"Customer showed interest in new products", "Discussed pricing options", "Resolved technical issue", "Scheduled follow-up meeting"})),
			Type:      noteTypes[rand.Intn(len(noteTypes))],
			CreatedBy: fmt.Sprintf("rep-%03d", rand.Intn(10)+1),
			CreatedAt: createdAt,
			UpdatedAt: createdAt,
			IsPrivate: rand.Float64() < 0.2, // 20% chance of private
			Tags:      []string{randomChoice([]string{"urgent", "follow-up", "opportunity", "issue", "pricing"})},
		}
	}
	return notes
}

func generateAssets(count int) []*pb.Asset {
	assets := make([]*pb.Asset, count)
	for i := 0; i < count; i++ {
		installDate := timestamppb.New(time.Now().AddDate(-rand.Intn(3)-1, 0, 0))
		lastMaintenance := timestamppb.New(time.Now().AddDate(0, -rand.Intn(6)-1, 0))

		assets[i] = &pb.Asset{
			AssetId:             fmt.Sprintf("asset-%03d", i+1),
			Name:                fmt.Sprintf("%s Unit #%d", randomChoice([]string{"Cooler", "Freezer", "Display", "POS System", "Shelving"}), i+1),
			Type:                assetTypes[rand.Intn(len(assetTypes))],
			Model:               fmt.Sprintf("Model-%s-%d", randomString(3), rand.Intn(999)+100),
			SerialNumber:        fmt.Sprintf("SN-%d-%06d", 2023+rand.Intn(2), rand.Intn(999999)+1),
			Status:              pb.AssetStatus_ASSET_STATUS_ACTIVE,
			InstallationDate:    installDate,
			LastMaintenanceDate: lastMaintenance,
			NextMaintenanceDate: timestamppb.New(time.Now().AddDate(0, rand.Intn(6)+1, 0)),
			LocationDetails:     fmt.Sprintf("Aisle %d, Section %s", rand.Intn(20)+1, randomChoice([]string{"A", "B", "C", "D"})),
			Condition:           randomChoice([]string{"Excellent", "Good", "Fair", "Needs Attention"}),
			MaintenanceHistory:  generateMaintenanceHistory(rand.Intn(3) + 1),
		}
	}
	return assets
}

func generateMaintenanceHistory(count int) []*pb.AssetMaintenance {
	history := make([]*pb.AssetMaintenance, count)
	for i := 0; i < count; i++ {
		history[i] = &pb.AssetMaintenance{
			Date:        timestamppb.New(time.Now().AddDate(0, -rand.Intn(12)-1, 0)),
			Type:        randomMaintenanceType(),
			Description: randomChoice([]string{"Routine cleaning", "Temperature calibration", "Parts replacement", "Software update"}),
			Technician:  randomChoice([]string{"Tech Services Inc", "Maintenance Pro", "Equipment Care Ltd"}),
			Cost:        rand.Float64()*500 + 50, // $50-$550
		}
	}
	return history
}

func randomMaintenanceType() pb.MaintenanceType {
	types := []pb.MaintenanceType{
		pb.MaintenanceType_MAINTENANCE_TYPE_ROUTINE,
		pb.MaintenanceType_MAINTENANCE_TYPE_REPAIR,
		pb.MaintenanceType_MAINTENANCE_TYPE_REPLACEMENT,
		pb.MaintenanceType_MAINTENANCE_TYPE_UPGRADE,
	}
	return types[rand.Intn(len(types))]
}

func generateChecklist(count int) []*pb.ChecklistItem {
	items := make([]*pb.ChecklistItem, count)
	for i := 0; i < count; i++ {
		dueDate := timestamppb.New(time.Now().AddDate(0, 0, rand.Intn(30)-15)) // -15 to +15 days
		isCompleted := rand.Float64() < 0.6                                    // 60% completion rate

		var completedDate *timestamppb.Timestamp
		var status pb.ChecklistStatus

		if isCompleted {
			completedDate = timestamppb.New(dueDate.AsTime().AddDate(0, 0, -rand.Intn(5)))
			status = pb.ChecklistStatus_CHECKLIST_STATUS_COMPLETED
		} else if dueDate.AsTime().Before(time.Now()) {
			status = pb.ChecklistStatus_CHECKLIST_STATUS_OVERDUE
		} else {
			status = pb.ChecklistStatus_CHECKLIST_STATUS_PENDING
		}

		items[i] = &pb.ChecklistItem{
			ItemId:        fmt.Sprintf("check-%03d", i+1),
			Title:         randomChoice([]string{"Display Compliance", "Inventory Check", "Safety Inspection", "Quality Review", "Marketing Setup"}),
			Description:   fmt.Sprintf("Checklist item %d description", i+1),
			Category:      checklistCategories[rand.Intn(len(checklistCategories))],
			Status:        status,
			Priority:      randomPriority(),
			DueDate:       dueDate,
			CompletedDate: completedDate,
			AssignedTo:    fmt.Sprintf("rep-%03d", rand.Intn(10)+1),
			CompletedBy: func() string {
				if isCompleted {
					return fmt.Sprintf("rep-%03d", rand.Intn(10)+1)
				} else {
					return ""
				}
			}(),
			Notes: randomChoice([]string{"Standard procedure", "Special attention required", "Follow brand guidelines", "Coordinate with manager"}),
		}
	}
	return items
}

func randomPriority() pb.Priority {
	priorities := []pb.Priority{
		pb.Priority_PRIORITY_LOW,
		pb.Priority_PRIORITY_MEDIUM,
		pb.Priority_PRIORITY_HIGH,
		pb.Priority_PRIORITY_CRITICAL,
	}
	return priorities[rand.Intn(len(priorities))]
}

func generateNews(count int) []*pb.News {
	news := make([]*pb.News, count)
	for i := 0; i < count; i++ {
		publishDate := timestamppb.New(time.Now().AddDate(0, 0, -(rand.Intn(30))))

		news[i] = &pb.News{
			NewsId:        fmt.Sprintf("news-%03d", i+1),
			Title:         randomChoice([]string{"Product Launch", "Market Update", "Policy Change", "Promotion Alert", "Training Available"}),
			Content:       fmt.Sprintf("News content %d: Important information about recent developments.", i+1),
			Type:          randomNewsType(),
			Source:        randomNewsSource(),
			PublishedDate: publishDate,
			Author:        randomChoice([]string{"Marketing Team", "Sales Department", "Management", "External Source"}),
			Url:           fmt.Sprintf("https://company.com/news/article-%d", i+1),
			Tags:          []string{randomChoice([]string{"product", "sales", "market", "policy", "training"})},
			IsImportant:   rand.Float64() < 0.3, // 30% chance of important
		}
	}
	return news
}

func randomNewsType() pb.NewsType {
	types := []pb.NewsType{
		pb.NewsType_NEWS_TYPE_GENERAL,
		pb.NewsType_NEWS_TYPE_PROMOTION,
		pb.NewsType_NEWS_TYPE_PRODUCT_LAUNCH,
		pb.NewsType_NEWS_TYPE_POLICY_CHANGE,
		pb.NewsType_NEWS_TYPE_MARKET_UPDATE,
		pb.NewsType_NEWS_TYPE_COMPETITOR,
	}
	return types[rand.Intn(len(types))]
}

func randomNewsSource() pb.NewsSource {
	sources := []pb.NewsSource{
		pb.NewsSource_NEWS_SOURCE_INTERNAL,
		pb.NewsSource_NEWS_SOURCE_EXTERNAL,
		pb.NewsSource_NEWS_SOURCE_OUTLET,
		pb.NewsSource_NEWS_SOURCE_MARKET_RESEARCH,
	}
	return sources[rand.Intn(len(sources))]
}

// Legacy function for backward compatibility
func generateMockOutletDetails(outletID string) *pb.OutletDetails {
	settings := MockSettings{
		AverageNotesList:               2,
		AverageVisitHistory:            3,
		AverageNumberOfOrders:          5,
		AverageOrderItemsPerOrder:      2,
		AverageTopProductsInStatistics: 5,
		AverageOutletsNearby:           3,
		AverageAssetList:               2,
		AverageChecklist:               4,
		AverageNews:                    3,
	}
	return GenerateMockedOutlet(outletID, settings)
}
