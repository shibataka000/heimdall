package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, event json.RawMessage) error {
	// // Parse the input event
	// var order Order
	// if err := json.Unmarshal(event, &order); err != nil {
	// 	log.Printf("Failed to unmarshal event: %v", err)
	// 	return err
	// }

	// // Access environment variables
	// bucketName := os.Getenv("RECEIPT_BUCKET")
	// if bucketName == "" {
	// 	log.Printf("RECEIPT_BUCKET environment variable is not set")
	// 	return fmt.Errorf("missing required environment variable RECEIPT_BUCKET")
	// }

	// // Create the receipt content and key destination
	// receiptContent := fmt.Sprintf("OrderID: %s\nAmount: $%.2f\nItem: %s",
	// 	order.OrderID, order.Amount, order.Item)
	// key := "receipts/" + order.OrderID + ".txt"

	// // Upload the receipt to S3 using the helper method
	// if err := uploadReceiptToS3(ctx, bucketName, key, receiptContent); err != nil {
	// 	return err
	// }

	// log.Printf("Successfully processed order %s and stored receipt in S3 bucket %s", order.OrderID, bucketName)
	return nil
}

func main() {
	lambda.Start(handler)
}
