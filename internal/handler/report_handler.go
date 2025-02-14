package handler

import (
	"log"
	"time"

	"github.com/idprm/go-football-alert/internal/domain/entity"
	"github.com/idprm/go-football-alert/internal/services"
)

type ReportHandler struct {
	serviceService      services.IServiceService
	subscriptionService services.ISubscriptionService
	transactionService  services.ITransactionService
	summaryService      services.ISummaryService
}

func NewReportHandler(
	serviceService services.IServiceService,
	subscriptionService services.ISubscriptionService,
	transactionService services.ITransactionService,
	summaryService services.ISummaryService,
) *ReportHandler {
	return &ReportHandler{
		serviceService:      serviceService,
		subscriptionService: subscriptionService,
		transactionService:  transactionService,
		summaryService:      summaryService,
	}
}

func (h *ReportHandler) TotalActiveSub() {

	services, err := h.serviceService.GetAll()
	if err != nil {
		log.Println(err)
	}

	if len(services) > 0 {

		for _, service := range services {
			count, err := h.subscriptionService.CountActiveSub(service.GetId())
			if err != nil {
				log.Println(err)
			}

			summary := &entity.Summary{
				ServiceID:      service.GetId(),
				TotalActiveSub: count,
				CreatedAt:      time.Now(),
			}

			// summary save
			h.summaryService.Save(summary)
		}
	}

}

func (h *ReportHandler) TotalReg() {

	services, err := h.serviceService.GetAll()
	if err != nil {
		log.Println(err)
	}

	if len(services) > 0 {

		for _, service := range services {
			count, err := h.transactionService.CountSubByDay(service.GetId())
			if err != nil {
				log.Println(err)
			}

			summary := &entity.Summary{
				ServiceID: service.GetId(),
				TotalSub:  count,
				CreatedAt: time.Now(),
			}

			// summary save
			h.summaryService.Save(summary)
		}
	}

}

func (h *ReportHandler) TotalUnreg() {

	services, err := h.serviceService.GetAll()
	if err != nil {
		log.Println(err)
	}

	if len(services) > 0 {

		for _, service := range services {
			count, err := h.transactionService.CountUnSubByDay(service.GetId())
			if err != nil {
				log.Println(err)
			}

			summary := &entity.Summary{
				ServiceID:  service.GetId(),
				TotalUnsub: count,
				CreatedAt:  time.Now(),
			}

			// summary save
			h.summaryService.Save(summary)
		}
	}

}

func (h *ReportHandler) TotalRenewal() {

	services, err := h.serviceService.GetAll()
	if err != nil {
		log.Println(err)
	}

	if len(services) > 0 {

		for _, service := range services {
			count, err := h.transactionService.CountRenewalByDay(service.GetId())
			if err != nil {
				log.Println(err)
			}

			summary := &entity.Summary{
				ServiceID:    service.GetId(),
				TotalRenewal: count,
				CreatedAt:    time.Now(),
			}

			// summary save
			h.summaryService.Save(summary)
		}
	}

}

func (h *ReportHandler) TotalRevenue() {

	services, err := h.serviceService.GetAll()
	if err != nil {
		log.Println(err)
	}

	if len(services) > 0 {

		for _, service := range services {
			count, err := h.transactionService.TotalRevenueByDay(service.GetId())
			if err != nil {
				log.Println(err)
			}

			summary := &entity.Summary{
				ServiceID:    service.GetId(),
				TotalRevenue: count,
				CreatedAt:    time.Now(),
			}

			// summary save
			h.summaryService.Save(summary)
		}
	}
}

func (h *ReportHandler) TotalSuccess() {
	services, err := h.serviceService.GetAll()
	if err != nil {
		log.Println(err)
	}

	if len(services) > 0 {

		for _, service := range services {
			count, err := h.transactionService.CountFailedByDay(service.GetId())
			if err != nil {
				log.Println(err)
			}

			summary := &entity.Summary{
				ServiceID:          service.GetId(),
				TotalChargeSuccess: count,
				CreatedAt:          time.Now(),
			}

			// summary save
			h.summaryService.Save(summary)
		}
	}
}

func (h *ReportHandler) TotalFailed() {
	services, err := h.serviceService.GetAll()
	if err != nil {
		log.Println(err)
	}

	if len(services) > 0 {

		for _, service := range services {
			count, err := h.transactionService.CountFailedByDay(service.GetId())
			if err != nil {
				log.Println(err)
			}

			summary := &entity.Summary{
				ServiceID:         service.GetId(),
				TotalChargeFailed: count,
				CreatedAt:         time.Now(),
			}

			// summary save
			h.summaryService.Save(summary)
		}
	}
}
