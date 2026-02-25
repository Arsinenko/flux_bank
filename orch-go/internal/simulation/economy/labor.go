package economy

import (
	"errors"
	"sync"

	"github.com/google/uuid"
)

// JobVacancy represents a job opening.
type JobVacancy struct {
	ID          uuid.UUID
	EmployerID  uuid.UUID
	Title       string
	Salary      float64
	Description string
}

// EmploymentContract represents an active job contract.
type EmploymentContract struct {
	ID         uuid.UUID
	EmployeeID uuid.UUID
	EmployerID uuid.UUID
	Title      string
	Salary     float64
	StartDate  int64 // Tick or Time
}

type LaborMarket struct {
	mu        sync.RWMutex
	Vacancies map[uuid.UUID]*JobVacancy
	Contracts map[uuid.UUID]*EmploymentContract
}

func NewLaborMarket() *LaborMarket {
	return &LaborMarket{
		Vacancies: make(map[uuid.UUID]*JobVacancy),
		Contracts: make(map[uuid.UUID]*EmploymentContract),
	}
}

func (l *LaborMarket) PostVacancy(employerID uuid.UUID, title string, salary float64) uuid.UUID {
	l.mu.Lock()
	defer l.mu.Unlock()

	id := uuid.New()
	vacancy := &JobVacancy{
		ID:         id,
		EmployerID: employerID,
		Title:      title,
		Salary:     salary,
	}
	l.Vacancies[id] = vacancy
	return id
}

func (l *LaborMarket) GetVacancies() []*JobVacancy {
	l.mu.RLock()
	defer l.mu.RUnlock()

	result := make([]*JobVacancy, 0, len(l.Vacancies))
	for _, v := range l.Vacancies {
		result = append(result, v)
	}
	return result
}

func (l *LaborMarket) ApplyAndHire(vacancyID uuid.UUID, employeeID uuid.UUID) (*EmploymentContract, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	vacancy, ok := l.Vacancies[vacancyID]
	if !ok {
		return nil, errors.New("vacancy not found")
	}

	// Create contract
	contractID := uuid.New()
	contract := &EmploymentContract{
		ID:         contractID,
		EmployeeID: employeeID,
		EmployerID: vacancy.EmployerID,
		Title:      vacancy.Title,
		Salary:     vacancy.Salary,
	}
	l.Contracts[contractID] = contract

	// Remove vacancy (assuming 1 position per vacancy listing for simplicity for now)
	delete(l.Vacancies, vacancyID)

	return contract, nil
}

func (l *LaborMarket) FireEmployee(contractID uuid.UUID) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.Contracts, contractID)
}

func (l *LaborMarket) GetContracts() []*EmploymentContract {
	l.mu.RLock()
	defer l.mu.RUnlock()

	result := make([]*EmploymentContract, 0, len(l.Contracts))
	for _, c := range l.Contracts {
		result = append(result, c)
	}
	return result
}

func (l *LaborMarket) GetContract(contractID uuid.UUID) *EmploymentContract {
	l.mu.RLock()
	defer l.mu.RUnlock()

	return l.Contracts[contractID]
}

func (l *LaborMarket) GetContractsByEmployerIdD(employerID uuid.UUID) []*EmploymentContract {
	l.mu.RLock()
	defer l.mu.RUnlock()

	result := make([]*EmploymentContract, 0)
	for _, c := range l.Contracts {
		if c.EmployerID == employerID {
			result = append(result, c)
		}
	}

	return result
}
