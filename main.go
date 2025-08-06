package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "srv-eazle-advise-mock/pkg/gen/proto/outlet"
	"srv-eazle-advise-mock/pkg/mock"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

const (
	SECRET_KEY = "eazle-secret-2025"
	PORT       = "8080"
)

func main() {
	http.HandleFunc("/outlets", handleOutletDetails)
	http.HandleFunc("/health", handleHealth)

	fmt.Printf("Server starting on port %s\n", PORT)
	fmt.Printf("Secret key required: %s\n", SECRET_KEY)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func handleOutletDetails(w http.ResponseWriter, r *http.Request) {
	// Check secret key
	if !validateSecretKey(r) {
		http.Error(w, "Unauthorized - Invalid secret key", http.StatusUnauthorized)
		return
	}

	// Handle delay if specified
	handleDelay(r)

	// Get outlet ID from query parameter
	outletID := r.URL.Query().Get("outlet_id")
	if outletID == "" {
		outletID = "outlet-001" // Default outlet
	}

	// Generate mock outlet data with configurable settings
	settings := mock.MockSettings{
		AverageNotesList:               30,
		AverageVisitHistory:            96,
		AverageNumberOfOrders:          90,
		AverageOrderItemsPerOrder:      20,
		AverageTopProductsInStatistics: 6,
		AverageOutletsNearby:           10,
		AverageAssetList:               6,
		AverageChecklist:               18,
		AverageNews:                    22,
	}

	numberOfOutletsString := r.Header.Get("X-Outlet-Num")
	if numberOfOutletsString == "" {
		numberOfOutletsString = "100"
	}
	numberOfOutlets, err := strconv.Atoi(numberOfOutletsString)
	if err != nil {
		http.Error(w, "Invalid X-Outlet-Num header", http.StatusBadRequest)
		return
	}

	outlets := &pb.OutletDetailsResponse{
		Details: []*pb.OutletDetails{},
	}

	outletChan := make(chan *pb.OutletDetails, numberOfOutlets)
	errChan := make(chan error, numberOfOutlets)

	for i := 0; i < numberOfOutlets; i++ {
		go func() {
			outlet := mock.GenerateMockedOutlet(outletID, settings)
			outletChan <- outlet
			errChan <- nil // Signal successful generation
		}()
	}

	for i := 0; i < numberOfOutlets; i++ {
		outlet := <-outletChan
		err := <-errChan
		if err != nil {
			http.Error(w, "Error generating outlet data", http.StatusInternalServerError)
			return
		}
		outlets.Details = append(outlets.Details, outlet)
	}

	var data []byte

	if r.Header.Get("Accept") == "application/protobuf" {
		data, err = proto.Marshal(outlets)
		w.Header().Set("Content-Type", "application/protobuf")
	} else {
		data, err = protojson.Marshal(outlets)
		w.Header().Set("Content-Type", "application/json")
	}
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func validateSecretKey(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	apiKey := r.Header.Get("X-API-Key")

	// Check both Authorization header and X-API-Key header
	return authHeader == "Bearer "+SECRET_KEY || apiKey == SECRET_KEY
}

func handleDelay(r *http.Request) {
	delayHeader := r.Header.Get("X-Delay-Ms")
	if delayHeader != "" {
		if delayMs, err := strconv.Atoi(delayHeader); err == nil && delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
	}
}
