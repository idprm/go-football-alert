package handler

import (
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/services"
)

type ReportHandler struct {
	serviceService          services.IServiceService
	subscriptionService     services.ISubscriptionService
	transactionService      services.ITransactionService
	summaryDashboardService services.ISummaryDashboardService
	summaryRevenueService   services.ISummaryRevenueService
}

func NewReportHandler(
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	summaryDashboardService services.ISummaryDashboardService,
	summaryRevenueService services.ISummaryRevenueService,
) *ReportHandler {
	return &ReportHandler{
		serviceService:          serviceService,
		subscriptionService:     subscriptionService,
		transactionService:      transactionService,
		summaryDashboardService: summaryDashboardService,
		summaryRevenueService:   summaryRevenueService,
	}
}

func (h *ReportHandler) GetTotalActiveSub() {
	count, err := h.summaryDashboardService.GetTotalActiveSub()
	if err != nil {
		log.Println(err)
	}
	// summary save
	h.summaryDashboardService.Save(
		&entity.SummaryDashboard{
			TotalActiveSub: count,
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		},
	)
}

func (h *ReportHandler) GetTotalRevenue() {
	total, err := h.summaryDashboardService.GetTotalRevenue()
	if err != nil {
		log.Println(err)
	}
	// summary save
	h.summaryDashboardService.Save(
		&entity.SummaryDashboard{
			TotalRevenue: total,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	)
}

func (h *ReportHandler) PopulateRevenue() {
	summs, err := h.summaryRevenueService.SelectRevenue()
	if err != nil {
		log.Println(err.Error())
	}

	if len(*summs) > 0 {
		for _, s := range *summs {
			h.summaryRevenueService.Save(
				&entity.SummaryRevenue{
					Subject:   s.Subject,
					Status:    s.Status,
					Total:     s.Total,
					Revenue:   s.Revenue,
					CreatedAt: s.CreatedAt,
					UpdatedAt: time.Now(),
				},
			)
		}
	}
}
