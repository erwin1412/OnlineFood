package scheduler

import (
	"context"
	"log"
	"merchant-service/internal/infra"
	"time"

	"github.com/robfig/cron/v3"
)

func StartMerchantStatusScheduler(repo *infra.PgMerchantRepository) {
	c := cron.New()

	// Contoh: setiap menit jalankan update
	c.AddFunc("@every 1m", func() {
		ctx := context.Background()
		now := time.Now()

		merchants, err := repo.GetAll(ctx)
		if err != nil {
			log.Printf("scheduler error: %v", err)
			return
		}

		for _, m := range merchants {
			open, _ := time.Parse("15:04", m.OpenHour)
			close, _ := time.Parse("15:04", m.CloseHour)

			// Hari ini + jam buka/tutup
			openTime := time.Date(now.Year(), now.Month(), now.Day(), open.Hour(), open.Minute(), 0, 0, now.Location())
			closeTime := time.Date(now.Year(), now.Month(), now.Day(), close.Hour(), close.Minute(), 0, 0, now.Location())

			status := "closed"
			if now.After(openTime) && now.Before(closeTime) {
				status = "open"
			}

			// Update status hanya kalau berubah
			if m.Status != status {
				m.Status = status
				m.UpdatedAt = time.Now()
				_, err := repo.Update(ctx, m)
				if err != nil {
					log.Printf("failed to update merchant %s: %v", m.ID, err)
				} else {
					log.Printf("updated merchant %s status to %s", m.ID, status)
				}
			}
		}
	})

	c.Start()
	log.Println("Merchant status scheduler started")
}
