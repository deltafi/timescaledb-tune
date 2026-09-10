package pgtune

import (
	"fmt"
	"testing"
)

func TestNullRecommender_Recommend(t *testing.T) {
	r := &NullRecommender{}
	// NullRecommender should ALWAYS return NoRecommendation
	for i := range 1000 {
		key := fmt.Sprintf("key%d", i)
		if val := r.Recommend(key); val != NoRecommendation {
			t.Errorf("Expected no recommendation for key %s but got %s", key, val)
		}
	}
}
