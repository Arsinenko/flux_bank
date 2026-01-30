package bank

import (
	"fmt"
	"orch-go/internal/domain/account"
	"orch-go/internal/domain/customer"
	"orch-go/internal/domain/customer_address"
	"orch-go/internal/domain/user_credential"
	"orch-go/internal/services"
	simcontext "orch-go/internal/simulation/context"
	"slices"
	"strings"
	"sync"
	"time"
)

type Registrable interface {
	SetCustomerID(id int32)
	SetAccountID(id int32)
	GetName() string
}

func RegisterAgent(ctx simcontext.AgentContext, svcs *services.ServiceContainer, agent Registrable) error {
	errChan := make(chan error, 3)
	var wg sync.WaitGroup

	fakeCustomer := customer.FakeCustomer(time.Now())
	createCustomer, err := svcs.CustomerService.CreateCustomer(ctx, fakeCustomer)
	if err != nil {
		return fmt.Errorf("agent %s failed to create createCustomer: %w", agent.GetName(), err)
	}
	agent.SetCustomerID(createCustomer.Id)

	wg.Add(1)
	go func() {
		defer wg.Done()
		fakeUserCreds := user_credential.FakeUserCreds()
		_, err = svcs.UserCredentialService.CreateUserCredential(ctx, &fakeUserCreds)
		if err != nil {
			errChan <- fmt.Errorf("agent %s failed to create user credential: %w", agent.GetName(), err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fakeAddress := customer_address.FakeCustomerAddress(createCustomer.Id)
		_, err = svcs.CustomerAddressService.CreateCustomerAddress(ctx, fakeAddress)
		if err != nil {
			errChan <- fmt.Errorf("agent %s failed to create customer address: %w", agent.GetName(), err)
		}
	}()

	wg.Wait()
	close(errChan)

	for err := range errChan {
		if err != nil {
			return err
		}
	}

	accountTypes, err := ctx.Services().AccountTypeService.GetAllAccountTypes(ctx, 0, 0, "", false)
	if err != nil {
		return fmt.Errorf("agent %s failed to get account types: %w", agent.GetName(), err)
	}
	pos, found := slices.BinarySearchFunc(accountTypes, "individual",
		func(accountType account.AccountType, s string) int {
			return strings.Compare(accountType.Name, s)
		},
	)
	if !found {
		return fmt.Errorf("agent %s failed to find account type 'individual'", agent.GetName())
	}

	// Assuming default account type, currency etc.
	acc, err := svcs.AccountService.CreateAccount(ctx, account.FakeAccount(createCustomer.Id, *accountTypes[pos].Id))
	if err != nil {
		return fmt.Errorf("agent %s failed to create account: %w", agent.GetName(), err)
	}
	agent.SetAccountID(*acc.Id)
	fmt.Printf("Agent %s registered in bank. CustomerID: %d, AccountID: %d\n", agent.GetName(), createCustomer.Id, *acc.Id)
	return nil
}
