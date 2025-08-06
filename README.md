# Eazle Advise Mock Server

A mock HTTP server that provides outlet data for testing and development purposes. The server returns JSON-encoded protobuf data with configurable response delays and requires authentication.

## Features

- Configurable mock outlet data generation with randomized content
- Dynamic data sizing based on MockSettings parameters
- Configurable response delays via headers
- Secret key authentication
- JSON responses from protobuf structures
- Health check endpoint
- Comprehensive outlet data including visits, orders, assets, checklists, and more

## Getting Started

### Prerequisites

- Go 1.24.3 or later
- Generated protobuf files (already included)

### Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on port 8080.

## API Endpoints

### Authentication

All endpoints (except `/health`) require authentication via one of these methods:

- **Authorization header**: `Authorization: Bearer eazle-secret-2024`
- **API Key header**: `X-API-Key: eazle-secret-2024`

### Response Delay

You can simulate network latency by adding a delay header:
- **Delay header**: `X-Delay-Ms: 500` (delay in milliseconds)

### Endpoints

#### 1. Health Check
```
GET /health
```

**Response:**
```json
{
  "status": "healthy"
}
```

#### 2. Get Outlet Details
```
GET /outlet?outlet_id=outlet-001
```

**Headers:**
- `Authorization: Bearer eazle-secret-2024` (required)
- `X-Delay-Ms: 1000` (optional - delay response by 1000ms)

**Query Parameters:**
- `outlet_id` (optional) - Outlet ID to retrieve. Defaults to "outlet-001"

**Response:** Complete outlet details in JSON format including:
- Basic outlet information
- Location and contact points
- Visit history with sales rep details
- Order history with items and payment info
- Statistics (revenue, growth, top products)
- Nearby outlets
- Notes and reminders
- Asset list with maintenance history
- Checklist items
- News and updates

#### 3. Search Outlets
```
GET /outlets/search
```

**Headers:**
- `Authorization: Bearer eazle-secret-2024` (required)
- `X-Delay-Ms: 500` (optional)

**Response:** Array of outlet summaries with key metrics

## Example Usage

### Using curl

1. **Get outlet details with delay:**
```bash
curl -H "Authorization: Bearer eazle-secret-2024" \
     -H "X-Delay-Ms: 1000" \
     "http://localhost:8080/outlet?outlet_id=outlet-001"
```

2. **Search outlets:**
```bash
curl -H "X-API-Key: eazle-secret-2024" \
     "http://localhost:8080/outlets/search"
```

3. **Health check:**
```bash
curl "http://localhost:8080/health"
```

### Using JavaScript/Fetch

```javascript
// Get outlet details
const response = await fetch('http://localhost:8080/outlet?outlet_id=outlet-001', {
  headers: {
    'Authorization': 'Bearer eazle-secret-2024',
    'X-Delay-Ms': '500'
  }
});
const outletData = await response.json();
```

## Mock Data Structure

The server returns comprehensive outlet data including:

### Outlet Details
- **Basic Info**: ID, name, code, type, status
- **Location**: Address, GPS coordinates, region
- **Contacts**: Multiple contact points with roles
- **Statistics**: Revenue, orders, visits, growth metrics
- **History**: Complete visit and order history

### Visit History
- Sales rep information
- Visit types (sales call, delivery, support, audit, training)
- Actions taken and follow-ups
- Duration and attachments

### Order History
- Order items with pricing and discounts
- Payment information and status
- Delivery tracking and status
- Sales rep attribution

### Statistics
- Year-to-date and historical revenue
- Product performance metrics
- Monthly revenue trends
- Customer segmentation
- Credit information

### Additional Features
- **Nearby outlets** with distance and relationship info
- **Notes** with categories and privacy settings
- **Assets** with maintenance schedules
- **Checklist items** for compliance and tasks
- **News** and updates relevant to the outlet

## Configuration

### Environment Variables
- `PORT`: Server port (default: 8080)
- `SECRET_KEY`: Authentication key (default: eazle-secret-2024)

### Mock Data Configuration
The server uses `MockSettings` to configure the amount of data generated for each outlet:

```go
type MockSettings struct {
    AverageNotesList               int  // Average number of notes per outlet
    AverageVisitHistory            int  // Average number of visits per outlet
    AverageNumberOfOrders          int  // Average number of orders per outlet
    AverateOrderItemsPerOrder      int  // Average items per order
    AverateTopProductsInStatistics int  // Number of top products in statistics
    AverateOutletsNearby           int  // Number of nearby outlets
    AverateAssetList               int  // Number of assets per outlet
    AverateChecklist               int  // Number of checklist items
    AverateNews                    int  // Number of news items
}
```

Current default settings in the server:
- Notes: ~5 per outlet
- Visit History: ~10 visits per outlet
- Orders: ~15 orders per outlet with ~3 items each
- Top Products: 5 products in statistics
- Nearby Outlets: ~4 nearby outlets
- Assets: ~3 assets per outlet
- Checklist: ~6 checklist items
- News: ~4 news items

### Customization
You can modify the `MockSettings` in `main.go` or the data pools in `pkg/mock/mock.go` to customize:
- Names, locations, and product lists
- Random data ranges and probabilities
- Business logic for data relationships

## Development

### Project Structure
```
srv-eazle-advise-mock/
├── main.go                          # HTTP server implementation
├── go.mod                           # Go module definition
├── proto/                           # Protocol buffer definitions
│   ├── outlet.proto                 # Main outlet data structures
│   └── outlet_service.proto         # gRPC service definitions
├── pkg/
│   ├── mock/
│   │   └── mock.go                  # Mock data generation with configurable settings
│   └── gen/proto/proto/outlet/      # Generated Go code from protobuf
│       └── outlet.pb.go             # Generated protobuf Go structs
└── test_server.sh                   # Test script for server functionality
```

### Adding New Endpoints
1. Define the handler function
2. Add authentication check with `validateSecretKey()`
3. Add delay handling with `handleDelay()`
4. Generate or retrieve mock data
5. Convert to JSON and return

### Modifying Mock Data
The mock data generation is now modular and configurable:

1. **Adjust quantities**: Modify `MockSettings` in `main.go` to change data volume
2. **Customize data pools**: Edit arrays in `pkg/mock/mock.go` like:
   - `outletNames`: Store names
   - `storeManagers`: Contact names  
   - `salesReps`: Sales representative names
   - `productNames`: Available products
   - `cities`: Location options
3. **Modify business logic**: Update generation functions for custom data relationships
4. **Add new data types**: Extend `MockSettings` and add corresponding generation functions

### Randomization Features
- **Smart variance**: Each "average" setting generates ±50% variance
- **Realistic relationships**: Orders reference real visits, statistics match order history
- **Time-based data**: Dates and timestamps follow logical sequences
- **Geographic consistency**: Nearby outlets share location characteristics

## Error Responses

- **401 Unauthorized**: Invalid or missing secret key
- **500 Internal Server Error**: JSON encoding errors

## Notes

- All timestamps are in RFC3339 format
- Financial amounts are in USD
- Distance measurements are in kilometers
- The server generates deterministic mock data for consistent testing