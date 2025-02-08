package services

import (
	"github.com/aliwert/go-hospital-management/internal/models"
	"github.com/aliwert/go-hospital-management/internal/repositories"
)

type AppointmentService struct {
	appointmentRepo *repositories.AppointmentRepository
}

func NewAppointmentService(appointmentRepo *repositories.AppointmentRepository) *AppointmentService {
	return &AppointmentService{
		appointmentRepo: appointmentRepo,
	}
}

func (s *AppointmentService) CreateAppointment(req *models.AppointmentCreateRequest) (*models.Appointment, error) {
	appointment := &models.Appointment{
		PatientID:       req.PatientID,
		DoctorID:        req.DoctorID,
		AppointmentDate: req.AppointmentDate,
		Description:     req.Description,
		Fee:             req.Fee,
		Status:          models.StatusPending,
		PaymentStatus:   models.PaymentStatusUnpaid,
	}

	if err := s.appointmentRepo.Create(appointment); err != nil {
		return nil, err
	}

	return appointment, nil
}

func (s *AppointmentService) GetAppointmentById(id uint) (*models.Appointment, error) {
	return s.appointmentRepo.FindById(id)
}

func (s *AppointmentService) GetAllAppointments() ([]models.Appointment, error) {
	return s.appointmentRepo.FindAll()
}

func (s *AppointmentService) GetPatientAppointments(patientId uint) ([]models.Appointment, error) {
	return s.appointmentRepo.FindByPatientId(patientId)
}

func (s *AppointmentService) GetDoctorAppointments(doctorId uint) ([]models.Appointment, error) {
	return s.appointmentRepo.FindByDoctorId(doctorId)
}

func (s *AppointmentService) UpdateAppointment(id uint, req *models.AppointmentUpdateRequest) error {
	appointment, err := s.appointmentRepo.FindById(id)
	if err != nil {
		return err
	}

	if req.Status != "" {
		appointment.Status = req.Status
	}
	if req.Description != "" {
		appointment.Description = req.Description
	}
	if req.PaymentStatus != "" {
		appointment.PaymentStatus = req.PaymentStatus
	}
	if req.CancelReason != "" {
		appointment.CancelReason = req.CancelReason
	}
	if req.Notes != "" {
		appointment.Notes = req.Notes
	}

	return s.appointmentRepo.Update(appointment)
}

func (s *AppointmentService) DeleteAppointment(id uint) error {
	return s.appointmentRepo.Delete(id)
}
