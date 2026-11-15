package main

import (
	"testing"
	"time"
)

func TestThrottleEventDoublesInterval(t *testing.T) {
	l := NewLimiter(10)

	l.wait(429)

	if l.Interval != 200*time.Millisecond {
		t.Errorf("Expected interval is 200ms, but i got %v", l.Interval)
	}
}

func TestDropPercent(t *testing.T) {
	l := Limiter{
		Interval:            200 * time.Millisecond,
		BaseInterval:        100 * time.Millisecond,
		GoodStatusCodeCount: 4,
	}

	l.nextDelay(200)

	if l.Interval < 165*time.Millisecond || l.Interval > 170*time.Millisecond {
		t.Errorf("Expected interval to drop to ~168ms, got %v", l.Interval)
	}
}

func TestRecoveryToBaseline(t *testing.T) {
	l := NewLimiter(10)

	type record struct {
		statusCode       int
		expectedInterval time.Duration
	}

	var records []record

	// 24 initial clean requests at baseline (100ms)
	for range 24 {
		records = append(records, record{200, 100 * time.Millisecond})
	}

	records = append(records,
		record{429, 200 * time.Millisecond}, // Throttled -> doubles
		record{503, 400 * time.Millisecond}, // Throttled again -> doubles

		// 4 clean requests holding steady before threshold (5) triggers decay
		record{200, 400 * time.Millisecond},
		record{200, 400 * time.Millisecond},
		record{200, 400 * time.Millisecond},
		record{200, 400 * time.Millisecond},

		// Steady descent (rounded to nearest millisecond)
		record{200, 336 * time.Millisecond}, // 5th good: 400 / 1.189 = 336.41ms
		record{200, 283 * time.Millisecond}, // 6th good: 336.41 / 1.189 = 282.94ms
		record{200, 238 * time.Millisecond}, // 7th good: 282.94 / 1.189 = 237.96ms
		record{200, 200 * time.Millisecond}, // 8th good: 237.96 / 1.189 = 200.14ms
		record{200, 168 * time.Millisecond}, // 9th good: 200.14 / 1.189 = 168.32ms
		record{200, 142 * time.Millisecond}, // 10th good: 168.32 / 1.189 = 141.57ms
		record{200, 119 * time.Millisecond}, // 11th good: 141.57 / 1.189 = 119.06ms
		record{200, 100 * time.Millisecond}, // 12th good: 119.06 / 1.189 = 100.14ms (rounds to 100ms)
	)

	// 39 trailing clean requests clamped at baseline (100ms)
	for range 39 {
		records = append(records, record{200, 100 * time.Millisecond})
	}

	for index, rec := range records {
		l.wait(rec.statusCode)
		actual := l.Interval.Round(time.Millisecond)

		t.Logf("Record #%d: sent status %d, interval is %v (rounded: %v)", index, rec.statusCode, l.Interval, actual)

		if actual != rec.expectedInterval {
			t.Errorf("Record #%d FAILED: expected %v, but got %v (raw: %v)",
				index, rec.expectedInterval, actual, l.Interval)
		}
	}
}