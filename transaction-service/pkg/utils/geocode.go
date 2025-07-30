package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type GeocodingResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

func GetLatLong(address string) (string, string, error) {
	key := os.Getenv("GOOGLE_API_KEY")
	if key == "" {
		return "", "", fmt.Errorf("google API key is empty")
	}

	baseURL := "https://maps.googleapis.com/maps/api/geocode/json"
	params := url.Values{}
	params.Add("address", address)
	params.Add("key", key)

	fullURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println("Geocode URL:", fullURL)

	resp, err := http.Get(fullURL)
	if err != nil {
		return "", "", fmt.Errorf("HTTP error: %w", err)
	}
	defer resp.Body.Close()

	var geoRes GeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoRes); err != nil {
		return "", "", fmt.Errorf("decode error: %w", err)
	}

	if geoRes.Status != "OK" || len(geoRes.Results) == 0 {
		return "", "", fmt.Errorf("failed to geocode address: %s", geoRes.Status)
	}

	lat := fmt.Sprintf("%f", geoRes.Results[0].Geometry.Location.Lat)
	lng := fmt.Sprintf("%f", geoRes.Results[0].Geometry.Location.Lng)

	return lat, lng, nil
}
