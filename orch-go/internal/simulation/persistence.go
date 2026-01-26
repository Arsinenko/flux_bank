package simulation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"orch-go/internal/simulation/agents"
	"orch-go/internal/simulation/agents/business"
	"os"
)

const storageFile = "simulation_state.json"

// AgentWrapper is used for unmarshalling into the correct agent type.
type AgentWrapper struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// SaveAgents serializes the list of agents to a JSON file.
func SaveAgents(agentsToSave []agents.Agent) error {
	var wrappers []AgentWrapper
	for _, agent := range agentsToSave {
		data, err := json.Marshal(agent)
		if err != nil {
			return fmt.Errorf("failed to marshal agent %s: %w", agent.ID(), err)
		}
		wrappers = append(wrappers, AgentWrapper{
			Type: agent.Type(),
			Data: data,
		})
	}

	fileData, err := json.MarshalIndent(wrappers, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal agent wrappers: %w", err)
	}

	return ioutil.WriteFile(storageFile, fileData, 0644)
}

// LoadAgents deserializes agents from a JSON file.
func LoadAgents() ([]agents.Agent, error) {
	if _, err := os.Stat(storageFile); os.IsNotExist(err) {
		return nil, nil // File doesn't exist, no state to load
	}

	fileData, err := ioutil.ReadFile(storageFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read state file: %w", err)
	}

	var wrappers []AgentWrapper
	if err := json.Unmarshal(fileData, &wrappers); err != nil {
		return nil, fmt.Errorf("failed to unmarshal agent wrappers: %w", err)
	}

	var loadedAgents []agents.Agent
	for _, w := range wrappers {
		var agent agents.Agent
		switch w.Type {
		case "Individual":
			var individual agents.Individual
			if err := json.Unmarshal(w.Data, &individual); err != nil {
				return nil, fmt.Errorf("failed to unmarshal individual: %w", err)
			}
			agent = &individual
		case "Company":
			var company business.Company
			if err := json.Unmarshal(w.Data, &company); err != nil {
				return nil, fmt.Errorf("failed to unmarshal company: %w", err)
			}
			agent = &company
		case "Shop":
			var shop business.Shop
			if err := json.Unmarshal(w.Data, &shop); err != nil {
				return nil, fmt.Errorf("failed to unmarshal shop: %w", err)
			}
			agent = &shop
		// Add other agent types here as they become serializable
		default:
			fmt.Printf("Unknown agent type '%s' during loading, skipping.\n", w.Type)
			continue
		}
		loadedAgents = append(loadedAgents, agent)
	}

	return loadedAgents, nil
}
