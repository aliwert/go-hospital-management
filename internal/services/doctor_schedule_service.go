package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type DoctorScheduleService struct {
	scheduleRepo *repositories.DoctorScheduleRepository
}

func NewDoctorScheduleService(scheduleRepo *repositories.DoctorScheduleRepository) *DoctorScheduleService {
	return &DoctorScheduleService{
		scheduleRepo: scheduleRepo,
	}
}

func (s *DoctorScheduleService) CreateSchedule(req *models.ScheduleCreateRequest) (*models.DoctorSchedule, error) {
	schedule := &models.DoctorSchedule{
		DoctorID:        req.DoctorID,
		WeekDay:         req.WeekDay,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		BreakStartTime:  req.BreakStartTime,
		BreakEndTime:    req.BreakEndTime,
		SlotDuration:    req.SlotDuration,
		MaxAppointments: req.MaxAppointments,
		IsActive:        true,
	}

	if err := s.scheduleRepo.Create(schedule); err != nil {
		return nil, err
	}

	return schedule, nil
}

func (s *DoctorScheduleService) GetScheduleById(id uint) (*models.DoctorSchedule, error) {
	return s.scheduleRepo.FindById(id)
}

func (s *DoctorScheduleService) GetDoctorSchedules(doctorId uint) ([]models.DoctorSchedule, error) {
	return s.scheduleRepo.FindByDoctorId(doctorId)
}

func (s *DoctorScheduleService) UpdateSchedule(id uint, req *models.ScheduleUpdateRequest) error {
	schedule, err := s.scheduleRepo.FindById(id)
	if err != nil {
		return err
	}

	if req.StartTime != "" {
		schedule.StartTime = req.StartTime
	}
	if req.EndTime != "" {
		schedule.EndTime = req.EndTime
	}
	if req.BreakStartTime != "" {
		schedule.BreakStartTime = req.BreakStartTime
	}
	if req.BreakEndTime != "" {
		schedule.BreakEndTime = req.BreakEndTime
	}
	if req.SlotDuration > 0 {
		schedule.SlotDuration = req.SlotDuration
	}
	if req.MaxAppointments > 0 {
		schedule.MaxAppointments = req.MaxAppointments
	}
	if req.IsActive != nil {
		schedule.IsActive = *req.IsActive
	}

	return s.scheduleRepo.Update(schedule)
}

func (s *DoctorScheduleService) DeleteSchedule(id uint) error {
	return s.scheduleRepo.Delete(id)
}
