package netatmo

import (
	"errors"
	"strings"
	"testing"
)

func TestAuthError(t *testing.T) {
	t.Run("Error() returns message", func(t *testing.T) {
		err := NewAuthError("test message", "test hint", ErrNotConfigured)
		if err.Error() != "test message" {
			t.Errorf("Error() = %q, want %q", err.Error(), "test message")
		}
	})

	t.Run("Unwrap() returns underlying error", func(t *testing.T) {
		err := NewAuthError("test message", "test hint", ErrNotConfigured)
		if !errors.Is(err, ErrNotConfigured) {
			t.Error("Unwrap() should return ErrNotConfigured")
		}
	})

	t.Run("UserMessage() with hint", func(t *testing.T) {
		err := NewAuthError("test message", "test hint", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "test message") {
			t.Errorf("UserMessage() = %q, want to contain message", result)
		}
		if !strings.Contains(result, "test hint") {
			t.Errorf("UserMessage() = %q, want to contain hint", result)
		}
		if !strings.Contains(result, "💡") {
			t.Errorf("UserMessage() = %q, want to contain hint emoji", result)
		}
	})

	t.Run("UserMessage() without hint", func(t *testing.T) {
		err := NewAuthError("test message", "", nil)
		result := err.UserMessage()
		if result != "test message" {
			t.Errorf("UserMessage() = %q, want %q", result, "test message")
		}
	})
}

func TestAPIError(t *testing.T) {
	t.Run("Error() returns formatted message", func(t *testing.T) {
		err := NewAPIError(401, "unauthorized", ErrAPIError)
		result := err.Error()
		if !strings.Contains(result, "401") {
			t.Errorf("Error() = %q, want to contain status code", result)
		}
		if !strings.Contains(result, "unauthorized") {
			t.Errorf("Error() = %q, want to contain message", result)
		}
	})

	t.Run("Unwrap() returns underlying error", func(t *testing.T) {
		err := NewAPIError(500, "server error", ErrAPIError)
		if !errors.Is(err, ErrAPIError) {
			t.Error("Unwrap() should return ErrAPIError")
		}
	})

	t.Run("UserMessage() for 401", func(t *testing.T) {
		err := NewAPIError(401, "unauthorized", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "Authentication failed") {
			t.Errorf("UserMessage() = %q, want auth failure message", result)
		}
		if !strings.Contains(result, "netatmo login") {
			t.Errorf("UserMessage() = %q, want login hint", result)
		}
	})

	t.Run("UserMessage() for 403", func(t *testing.T) {
		err := NewAPIError(403, "forbidden", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "Authentication failed") {
			t.Errorf("UserMessage() = %q, want auth failure message", result)
		}
	})

	t.Run("UserMessage() for 404", func(t *testing.T) {
		err := NewAPIError(404, "not found", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "not found") {
			t.Errorf("UserMessage() = %q, want not found message", result)
		}
	})

	t.Run("UserMessage() for 429", func(t *testing.T) {
		err := NewAPIError(429, "rate limited", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "Rate limit") {
			t.Errorf("UserMessage() = %q, want rate limit message", result)
		}
	})

	t.Run("UserMessage() for 500", func(t *testing.T) {
		err := NewAPIError(500, "server error", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "unavailable") {
			t.Errorf("UserMessage() = %q, want unavailable message", result)
		}
	})

	t.Run("UserMessage() for 502", func(t *testing.T) {
		err := NewAPIError(502, "bad gateway", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "unavailable") {
			t.Errorf("UserMessage() = %q, want unavailable message", result)
		}
	})

	t.Run("UserMessage() for 503", func(t *testing.T) {
		err := NewAPIError(503, "service unavailable", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "unavailable") {
			t.Errorf("UserMessage() = %q, want unavailable message", result)
		}
	})

	t.Run("UserMessage() for unknown status", func(t *testing.T) {
		err := NewAPIError(418, "I'm a teapot", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "418") {
			t.Errorf("UserMessage() = %q, want to contain status code", result)
		}
	})
}

func TestNetworkError(t *testing.T) {
	t.Run("Error() returns message", func(t *testing.T) {
		err := NewNetworkError("connection failed", ErrNetworkError)
		if err.Error() != "connection failed" {
			t.Errorf("Error() = %q, want %q", err.Error(), "connection failed")
		}
	})

	t.Run("Unwrap() returns underlying error", func(t *testing.T) {
		err := NewNetworkError("connection failed", ErrNetworkError)
		if !errors.Is(err, ErrNetworkError) {
			t.Error("Unwrap() should return ErrNetworkError")
		}
	})

	t.Run("UserMessage() includes hint", func(t *testing.T) {
		err := NewNetworkError("connection failed", nil)
		result := err.UserMessage()
		if !strings.Contains(result, "connection failed") {
			t.Errorf("UserMessage() = %q, want to contain message", result)
		}
		if !strings.Contains(result, "internet connection") {
			t.Errorf("UserMessage() = %q, want to contain hint", result)
		}
	})
}

func TestFormatError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		result := FormatError(nil)
		if result != "" {
			t.Errorf("FormatError(nil) = %q, want empty string", result)
		}
	})

	t.Run("AuthError", func(t *testing.T) {
		err := NewAuthError("auth failed", "run login", nil)
		result := FormatError(err)
		if !strings.Contains(result, "auth failed") {
			t.Errorf("FormatError() = %q, want to contain message", result)
		}
	})

	t.Run("APIError", func(t *testing.T) {
		err := NewAPIError(401, "unauthorized", nil)
		result := FormatError(err)
		if !strings.Contains(result, "Authentication") {
			t.Errorf("FormatError() = %q, want to contain user message", result)
		}
	})

	t.Run("NetworkError", func(t *testing.T) {
		err := NewNetworkError("timeout", nil)
		result := FormatError(err)
		if !strings.Contains(result, "timeout") {
			t.Errorf("FormatError() = %q, want to contain message", result)
		}
	})

	t.Run("generic error", func(t *testing.T) {
		err := errors.New("generic error")
		result := FormatError(err)
		if result != "generic error" {
			t.Errorf("FormatError() = %q, want %q", result, "generic error")
		}
	})
}

func TestSentinelErrors(t *testing.T) {
	// Verify sentinel errors are properly defined
	sentinels := []struct {
		name string
		err  error
	}{
		{"ErrNotConfigured", ErrNotConfigured},
		{"ErrNotAuthenticated", ErrNotAuthenticated},
		{"ErrTokenExpired", ErrTokenExpired},
		{"ErrNoDevices", ErrNoDevices},
		{"ErrNetworkError", ErrNetworkError},
		{"ErrAPIError", ErrAPIError},
	}

	for _, tt := range sentinels {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Errorf("%s should not be nil", tt.name)
			}
			if tt.err.Error() == "" {
				t.Errorf("%s.Error() should not be empty", tt.name)
			}
		})
	}
}
