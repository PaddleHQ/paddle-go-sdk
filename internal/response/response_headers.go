package response

import (
	"net/http"
	"strconv"
	"time"

	"github.com/PaddleHQ/paddle-go-sdk/v5/pkg/paddleerr"
)

// parseRetry extracts Retry-After from response headers.
func parseRetry(headers http.Header) *paddleerr.RetryAfter {
	retryAfterValue := headers.Get("Retry-After")
	if retryAfterValue == "" {
		return nil
	}

	now := time.Now()

	// Try parsing as integer (seconds)
	if seconds, err := strconv.Atoi(retryAfterValue); err == nil {
		return &paddleerr.RetryAfter{
			IssuedAt: now,
			RetryAt:  now.Add(time.Duration(seconds) * time.Second),
		}
	}

	// Try parsing as HTTP date (RFC 7231)
	if parsed, err := time.Parse(time.RFC1123, retryAfterValue); err == nil {
		return &paddleerr.RetryAfter{
			IssuedAt: now,
			RetryAt:  parsed,
		}
	}

	// Try parsing as HTTP date with numeric timezone
	if parsed, err := time.Parse(time.RFC1123Z, retryAfterValue); err == nil {
		return &paddleerr.RetryAfter{
			IssuedAt: now,
			RetryAt:  parsed,
		}
	}

	return nil
}
