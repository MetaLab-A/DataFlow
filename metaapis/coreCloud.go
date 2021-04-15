package metaapis

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// ReadStock get data from cloud
func ReadCloud(field string, ctx context.Context, client *firestore.Client, readOnce bool) ([]map[string]interface{}, bool) {
	store := make([]map[string]interface{}, 0)
	iter := client.Collection(field).Documents(ctx)
	isEmpty := false

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed reading %s type: %v", field, err)
		}

		store = append(store, doc.Data())

		if readOnce {
			break
		}
	}

	if len(store) == 0 {
		isEmpty = true
		log.Printf("%s on cloud is empty\n", field)
	}

	return store, isEmpty
}
