package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)

var summaryFlag bool

var usageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Show Copilot premium request usage summary",
	Long: `Show the Copilot premium request usage summary for the authenticated user.

By default, outputs the full JSON response from the GitHub API.

With the --summary flag, only the total gross quantity is displayed.`,
	Args: cobra.NoArgs,
	RunE: runUsage,
}

func init() {
	rootCmd.AddCommand(usageCmd)
	usageCmd.Flags().BoolVarP(&summaryFlag, "summary", "s", false, "Show total gross quantity instead of full JSON")
}

// UsageResponse represents the billing usage response from the GitHub API
type UsageResponse struct {
	TimePeriod TimePeriod  `json:"timePeriod"`
	User       string      `json:"user"`
	UsageItems []UsageItem `json:"usageItems"`
}

// TimePeriod represents the billing period
type TimePeriod struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

// UsageItem represents a single usage item
type UsageItem struct {
	Product          string  `json:"product"`
	SKU              string  `json:"sku"`
	Model            string  `json:"model"`
	UnitType         string  `json:"unitType"`
	PricePerUnit     float64 `json:"pricePerUnit"`
	GrossQuantity    float64 `json:"grossQuantity"`
	GrossAmount      float64 `json:"grossAmount"`
	DiscountQuantity float64 `json:"discountQuantity"`
	DiscountAmount   float64 `json:"discountAmount"`
	NetQuantity      float64 `json:"netQuantity"`
	NetAmount        float64 `json:"netAmount"`
}

func runUsage(_ *cobra.Command, _ []string) error {
	client, err := api.DefaultRESTClient()
	if err != nil {
		return fmt.Errorf("failed to create REST client: %w", err)
	}

	// Get the authenticated user
	username, err := getAuthenticatedUser(client)
	if err != nil {
		return fmt.Errorf("failed to get authenticated user: %w", err)
	}

	// Fetch the usage data
	var response UsageResponse
	endpoint := fmt.Sprintf("users/%s/settings/billing/premium_request/usage", username)
	err = client.Get(endpoint, &response)
	if err != nil {
		return fmt.Errorf("failed to fetch usage data: %w", err)
	}

	if summaryFlag {
		// Calculate total gross quantity
		var totalGrossQuantity float64
		for _, item := range response.UsageItems {
			totalGrossQuantity += item.GrossQuantity
		}
		fmt.Printf("%.1f\n", totalGrossQuantity)
	} else {
		// Output full JSON response
		output, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal response: %w", err)
		}
		fmt.Println(string(output))
	}

	return nil
}

func getAuthenticatedUser(client *api.RESTClient) (string, error) {
	var user struct {
		Login string `json:"login"`
	}
	err := client.Get("user", &user)
	if err != nil {
		return "", err
	}
	return user.Login, nil
}
