package ipdata_test

import (
	"context"
	"testing"

	"github.com/vnkot/piklnk/pkg/ipdata"
)

func TestGetIPWho_Success(t *testing.T) {
	testIPList := map[string]ipdata.IPData{
		"8.8.8.8": ipdata.IPData{
			City:    "Mountain View",
			Country: "United States",
			Region:  "California",
		},
		"1.1.1.1": ipdata.IPData{
			City:    "South Brisbane",
			Country: "Australia",
			Region:  "Queensland",
		},
		"77.88.55.66": ipdata.IPData{
			City:    "Moscow",
			Country: "Russia",
			Region:  "Moscow",
		},
	}

	for ip, expectedIpData := range testIPList {
		receivedIpData, err := ipdata.GetIPWho(context.Background(), ip)

		if err != nil {
			t.Fatalf("Error get IP data %v", err)
		}

		if expectedIpData.City != receivedIpData.City {
			t.Fatalf("Expected %s, received %s", expectedIpData.City, receivedIpData.City)
		}

		if expectedIpData.Region != receivedIpData.Region {
			t.Fatalf("Expected %s, received %s", expectedIpData.Region, receivedIpData.Region)
		}

		if expectedIpData.Country != receivedIpData.Country {
			t.Fatalf("Expected %s, received %s", expectedIpData.Country, receivedIpData.Country)
		}
	}
}
