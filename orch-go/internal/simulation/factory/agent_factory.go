package factory

import (
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/agents/business"
)

type AgentFactory struct{}

func NewAgentFactory() *AgentFactory {
	return &AgentFactory{}
}

func (f *AgentFactory) CreateIndividual(name string) agents.Agent {
	return agents.NewIndividual(name)
}

func (f *AgentFactory) CreateCompany(name string, targetEmployees int) agents.Agent {
	return business.NewCompany(name, targetEmployees)
}

func (f *AgentFactory) CreateShop(name string) agents.Agent {
	return business.NewShop(name)
}

func (f *AgentFactory) CreateEntrepreneur(name string) agents.Agent {
	return business.NewEntrepreneur(name)
}

func (f *AgentFactory) CreateSelfEmployed(name string) agents.Agent {
	return business.NewSelfEmployed(name)
}
