package proxy

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/rdwr-kumaresans/EdgeGo/internal/logger"
)

func NewReverseProxy(
	target string,
	log logger.Logger,
) (http.Handler, error) {
	targetURL, err := url.Parse(target)
	if err != nil {
		return nil, fmt.Errorf("parse proxy target: %w", err)
	}

	if (targetURL.Scheme != "http" &&
		targetURL.Scheme != "https")	||
		targetURL.Host == "" {
			return nil, fmt.Errorf(
				"invalid proxy target %q: expected an HTTP or HTTPS URL",
				target,
			)
		}
	
		reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

		reverseProxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
			log.Error(
				"Reverse proxy request failed",
				"method", r.Method,
				"path", r.URL.Path,
				"target", targetURL.String(),
				"error", err,
			)

			http.Error(
				w,
				"Bad Gateway",
				http.StatusBadGateway,
			)
		}

		return http.StripPrefix("/api", reverseProxy), nil
}