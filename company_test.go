package client

import "fmt"

func ExampleManageCompany() {
	client, err := Dial("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	if err := client.CreateCompany("giantswarm.io"); err != nil {
		panic(err)
	}

	if err := client.AddMembers("giantswarm.io", []string{"stephan", "tim", "timo"}); err != nil {
		panic(err)
	}

	if err := client.AddMembers("giantswarm.io", []string{"dennis"}); err != nil {
		panic(err)
	}

	var initialCompany Company
	if err := client.GetCompany("giantswarm.io", &initialCompany); err != nil {
		panic(err)
	}

	companies, err := client.FindCompaniesByUser("dennis")
	if err != nil {
		panic(err)
	}

	if err := client.RemoveMembers("giantswarm.io", []string{"stephan"}); err != nil {
		panic(err)
	}

	var laterCompany Company
	if err := client.GetCompany("giantswarm.io", &laterCompany); err != nil {
		panic(err)
	}

	if err := client.DeleteCompany("giantswarm.io"); err != nil {
		panic(err)
	}

	var closedCompany Company
	if err := client.GetCompany("giantswarm.io", &closedCompany); err == nil {
		panic("Expected not-found-error")
	}

	noMembersShips, err := client.FindCompaniesByUser("stephan")
	if err != nil {
		panic(err)
	}

	fmt.Println(initialCompany.Members)
	fmt.Println(companies)
	fmt.Println(laterCompany.Members)
	fmt.Println(noMembersShips)
	// Output:
	// [stephan tim timo dennis]
	// [giantswarm.io]
	// [tim timo dennis]
	// []
}
