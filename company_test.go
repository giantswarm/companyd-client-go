// +build integration

package client

import "fmt"

func ExampleManageCompany() {
	client, err := Dial("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	var initialCompanies ListCompaniesResult
	if initialCompanies, err = client.ListCompanies(); err != nil {
		panic(err)
	}

	if err := client.CreateCompany("giantswarm.io", CompanyFields{DefaultCluster: "foo"}); err != nil {
		panic(err)
	}

	if err := client.AddMembers("giantswarm.io", []string{"stephan", "tim1", "timo"}); err != nil {
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

	noMemberships, err := client.FindCompaniesByUser("stephan")
	if err != nil {
		panic(err)
	}

	if err := client.CreateCompany("test1.giantswarm.io", CompanyFields{DefaultCluster: "foo"}); err != nil {
		panic(err)
	}
	if err := client.CreateCompany("test2.giantswarm.io", CompanyFields{DefaultCluster: "foo"}); err != nil {
		panic(err)
	}
	if err := client.AddMembers("test1.giantswarm.io", []string{"chris"}); err != nil {
		panic(err)
	}
	if err := client.AddMembers("test2.giantswarm.io", []string{"chris"}); err != nil {
		panic(err)
	}
	chrisCompanies, err := client.FindCompaniesByUser("chris")
	if err != nil {
		panic(err)
	}
	if err := client.RemoveUserFromAllCompanies("chris"); err != nil {
		panic(err)
	}
	chrisCompaniesAfter, err := client.FindCompaniesByUser("chris")
	if err != nil {
		panic(err)
	}

	var afterCompanies ListCompaniesResult
	if afterCompanies, err = client.ListCompanies(); err != nil {
		panic(err)
	}

	fmt.Println(initialCompanies)
	fmt.Println(initialCompany.Members)
	fmt.Println(companies)
	fmt.Println(laterCompany.Members)
	fmt.Println(noMemberships)
	fmt.Println(chrisCompanies)
	fmt.Println(chrisCompaniesAfter)
	fmt.Println(afterCompanies)
	// Output:
	// {[] 0 false}
	// [stephan tim1 timo dennis]
	// [giantswarm.io]
	// [tim1 timo dennis]
	// []
	// [test1.giantswarm.io test2.giantswarm.io]
	// []
	// {[test1.giantswarm.io test2.giantswarm.io] 2 false}
}
