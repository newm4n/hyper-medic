package code

import (
	"fmt"
	"strings"
)

type BundleType string

const (
	BundleTypeDocument                 BundleType = "document"
	BundleTypeMessage                  BundleType = "message"
	BundleTypeTransaction              BundleType = "transaction"
	BundleTypeTransactionResponse      BundleType = "transaction-response"
	BundleTypeBatch                    BundleType = "batch"
	BundleTypeBatchResponse            BundleType = "batch-response"
	BundleTypeHistoryList              BundleType = "history"
	BundleTypeSearchResults            BundleType = "searchset"
	BundleTypeCollection               BundleType = "collection"
	BundleTypeSubscriptionNotification BundleType = "subscription-notification"
)

func AllBundleType() []BundleType {
	return []BundleType{
		BundleTypeDocument,
		BundleTypeMessage,
		BundleTypeTransaction,
		BundleTypeTransactionResponse,
		BundleTypeBatch,
		BundleTypeBatchResponse,
		BundleTypeHistoryList,
		BundleTypeSearchResults,
		BundleTypeCollection,
		BundleTypeSubscriptionNotification,
	}
}

func FindBundleType(filter string) []BundleType {
	ret := make([]BundleType, 0)
	for _, at := range AllBundleType() {
		if strings.ToLower(at.String())[0:len(filter)] == strings.ToLower(filter) {
			ret = append(ret, at)
		}
	}
	return ret
}

func (au BundleType) ToString() {
	fmt.Println(au.String())
}
func (au BundleType) String() string {
	switch au {
	case BundleTypeDocument:
		return "Document"
	case BundleTypeMessage:
		return "Message"
	case BundleTypeTransaction:
		return "Transaction"
	case BundleTypeTransactionResponse:
		return "Transaction Response"
	case BundleTypeBatch:
		return "Batch"
	case BundleTypeBatchResponse:
		return "Batch Response"
	case BundleTypeHistoryList:
		return "History List"
	case BundleTypeSearchResults:
		return "Search Results"
	case BundleTypeCollection:
		return "Collection"
	case BundleTypeSubscriptionNotification:
		return "Subscription Notification"
	default:
		return "Unknown Bundle Type"
	}
}
