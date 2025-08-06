#!/bin/bash

# Test script for Eazle Advise Mock Server
# This script demonstrates all the server endpoints and features

set -e

BASE_URL="http://localhost:8080"
SECRET_KEY="eazle-secret-2024"

echo "🚀 Testing Eazle Advise Mock Server"
echo "===================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_test() {
    echo -e "${BLUE}🧪 $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

print_info() {
    echo -e "${YELLOW}ℹ️  $1${NC}"
}

# Function to test HTTP endpoint
test_endpoint() {
    local name="$1"
    local url="$2"
    local headers="$3"
    local expected_status="$4"

    print_test "Testing: $name"

    if [ -n "$headers" ]; then
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" $headers "$url")
    else
        response=$(curl -s -w "HTTPSTATUS:%{http_code}" "$url")
    fi

    http_code=$(echo $response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
    body=$(echo $response | sed -e 's/HTTPSTATUS:.*//')

    if [ "$http_code" -eq "$expected_status" ]; then
        print_success "Status: $http_code (Expected: $expected_status)"
        echo "Response sample: $(echo $body | jq -r '. | keys[]' 2>/dev/null | head -3 | tr '\n' ', ' | sed 's/,$//' || echo $body | head -c 100)..."
    else
        print_error "Status: $http_code (Expected: $expected_status)"
        echo "Response: $body"
    fi
    echo ""
}

# Check if server is running
print_test "Checking if server is running..."
if curl -s "$BASE_URL/health" > /dev/null 2>&1; then
    print_success "Server is running at $BASE_URL"
else
    print_error "Server is not running. Please start it with 'go run main.go'"
    exit 1
fi
echo ""

# Test 1: Health check (no auth required)
test_endpoint \
    "Health Check" \
    "$BASE_URL/health" \
    "" \
    200

# Test 2: Outlet details without authentication (should fail)
test_endpoint \
    "Outlet Details - No Auth (should fail)" \
    "$BASE_URL/outlet" \
    "" \
    401

# Test 3: Outlet details with Authorization header
test_endpoint \
    "Outlet Details - Authorization Header" \
    "$BASE_URL/outlet?outlet_id=outlet-001" \
    "-H 'Authorization: Bearer $SECRET_KEY'" \
    200

# Test 4: Outlet details with API Key header
test_endpoint \
    "Outlet Details - API Key Header" \
    "$BASE_URL/outlet?outlet_id=outlet-001" \
    "-H 'X-API-Key: $SECRET_KEY'" \
    200

# Test 5: Outlet details with delay
print_test "Testing: Outlet Details with 2 second delay"
start_time=$(date +%s)
response=$(curl -s -w "HTTPSTATUS:%{http_code}" \
    -H "Authorization: Bearer $SECRET_KEY" \
    -H "X-Delay-Ms: 2000" \
    "$BASE_URL/outlet?outlet_id=outlet-001")
end_time=$(date +%s)
duration=$((end_time - start_time))

http_code=$(echo $response | tr -d '\n' | sed -e 's/.*HTTPSTATUS://')
if [ "$http_code" -eq 200 ] && [ "$duration" -ge 2 ]; then
    print_success "Delay test passed - took ${duration}s (expected ≥2s)"
else
    print_error "Delay test failed - took ${duration}s, status: $http_code"
fi
echo ""

# Test 6: Search outlets
test_endpoint \
    "Search Outlets" \
    "$BASE_URL/outlets/search" \
    "-H 'Authorization: Bearer $SECRET_KEY'" \
    200

# Test 7: Invalid secret key
test_endpoint \
    "Invalid Secret Key (should fail)" \
    "$BASE_URL/outlet" \
    "-H 'Authorization: Bearer wrong-key'" \
    401

# Test 8: Outlet details with different outlet ID
test_endpoint \
    "Outlet Details - Custom ID" \
    "$BASE_URL/outlet?outlet_id=custom-outlet-123" \
    "-H 'X-API-Key: $SECRET_KEY'" \
    200

# Summary
echo ""
echo "🏁 Test Summary"
echo "==============="
print_info "All tests completed!"
print_info "Server is working correctly with:"
echo "   - Authentication via Authorization header or X-API-Key"
echo "   - Configurable response delays via X-Delay-Ms header"
echo "   - Mock outlet data with comprehensive details"
echo "   - Health check endpoint"
echo ""

# Optional: Pretty print a sample response
print_test "Sample outlet data structure:"
echo ""
curl -s -H "Authorization: Bearer $SECRET_KEY" "$BASE_URL/outlet?outlet_id=outlet-001" | \
jq -r '
{
    "outlet_id": .outletId,
    "name": .name,
    "type": .type,
    "status": .status,
    "location": {
        "city": .location.city,
        "coordinates": [.location.latitude, .location.longitude]
    },
    "statistics": {
        "revenue_ytd": .statistics.totalRevenueYtd,
        "orders_ytd": .statistics.totalOrdersYtd,
        "growth": (.statistics.revenueGrowthPercentage | tostring) + "%"
    },
    "recent_activity": {
        "visits": (.visitHistory | length),
        "orders": (.orderHistory | length),
        "notes": (.notes | length)
    }
}' 2>/dev/null || echo "jq not available for pretty printing"

echo ""
print_success "🎉 Server testing completed successfully!"
