package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HitRequest struct {
	Hit     int    `json:"hit"`
	Request string `json:"request"`
}

type RequestCounter struct {
	counts         map[string]int
	maxHits        int
	maxHitsRequest string
}

func NewRequestCounter() *RequestCounter {
	return &RequestCounter{
		counts: make(map[string]int),
	}
}

func (r *RequestCounter) Count(req string) *RequestCounter {
	if _, ok := r.counts[req]; !ok {
		r.counts[req] = 0
	}
	if req == r.maxHitsRequest {
		r.maxHits++
	}
	r.counts[req]++
	if r.counts[req] > r.maxHits {
		r.maxHits = r.counts[req]
		r.maxHitsRequest = req
	}
	return r
}

func (r *RequestCounter) ToHitRequest() *HitRequest {
	return &HitRequest{
		Hit:     r.maxHits,
		Request: r.maxHitsRequest,
	}
}

func CreateStatisticsRoute(cache *RequestCounter) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.JSON(http.StatusOK, cache.ToHitRequest())
	}
}
