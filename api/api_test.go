package api

import (
	"testing"
)

func TestGetAPOD(t *testing.T) {
	got := GetAPOD("1997-08-08")
	want := ResponseData{
		Copyright:   "",
		Date:        "1997-08-08",
		Title:       "White Oval Clouds on Jupiter",
		URL:         "https://apod.nasa.gov/apod/image/9708/jupiterovals_gal.jpg",
		Explanation: "What are those white ovals all over Jupiter? Storms!  Jupiter's clouds can swirl rapidly in raised high-pressure storm systems that circle the planet. The above pictured white ovals are located near the Great Red Spot, and have persisted on Jupiter since the 1930s. The Great Red Spot has persisted for at least 300 years.  Currently, no one knows why ovals last as long as they do. White ovals are confined to circular belts around Jupiter, but can interact to cause nearby chaotic cloud regions.",
	}

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
