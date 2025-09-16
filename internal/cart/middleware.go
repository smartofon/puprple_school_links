package cart

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type responseLogger struct {
	w      http.ResponseWriter
	status int
}

func (rl *responseLogger) Header() http.Header {
	return rl.w.Header()
}

func (rl *responseLogger) Write(b []byte) (int, error) {
	if rl.status == 0 {
		rl.status = http.StatusOK // По умолчанию считаем, что статус OK
	}
	return rl.w.Write(b)
}

func (rl *responseLogger) WriteHeader(statusCode int) {
	rl.status = statusCode
	rl.w.WriteHeader(statusCode)
}

// Middleware для логирования запросов
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Начальное время запроса
		startTime := time.Now()

		// Создаем обертку для ResponseWriter, чтобы захватить статус-код
		logResponseWriter := &responseLogger{w: w}

		logrus.SetFormatter(&logrus.JSONFormatter{})

		// Вызываем следующий обработчик
		next.ServeHTTP(logResponseWriter, r)

		// Логируем информацию о запросе
		logrus.WithFields(logrus.Fields{
			"method":     r.Method,
			"path":       r.URL.Path,
			"status":     logResponseWriter.status,
			"duration":   time.Since(startTime).String(),
			"remote_ip":  r.RemoteAddr,
			"user_agent": r.UserAgent(),
		}).Info("HTTP request")
	})
}
