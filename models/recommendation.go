package models

type Recommendation struct {
	Address string `json:"address,omitempty"`
	Name    string `json:"name,omitempty"`
	Cost    int64 `json:"cost,omitempty"`
	PlaceId string `json:"placeId"`
	Rating  float64 `json:"rating,omitempty"`
	Types   []string `json:"types,omitempty"`
	Score   float64 `json:"-"`
}

type RecommendationList []Recommendation

func (r RecommendationList) Len() int {
	return len(r)
}

func (r RecommendationList) Less(i, j int) bool {
	return r[i].Score < r[j].Score
}

func (r RecommendationList) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
