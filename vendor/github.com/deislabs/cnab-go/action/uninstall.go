package action

import (
	"io"

	"github.com/deislabs/cnab-go/claim"
	"github.com/deislabs/cnab-go/credentials"
	"github.com/deislabs/cnab-go/driver"
)

// Uninstall runs an uninstall action
type Uninstall struct {
	Driver driver.Driver
}

// Run performs the uninstall steps and updates the Claim
func (u *Uninstall) Run(c *claim.Claim, creds credentials.Set, w io.Writer) error {
	invocImage, err := selectInvocationImage(u.Driver, c)
	if err != nil {
		return err
	}

	op, err := opFromClaim(claim.ActionUninstall, stateful, c, invocImage, creds, w)
	if err != nil {
		return err
	}

	opResult, err := u.Driver.Run(op)
	c.Outputs = map[string]string{}
	for outputName, v := range c.Bundle.Outputs.Fields {
		if opResult.Outputs[v.Path] != "" {
			c.Outputs[outputName] = opResult.Outputs[v.Path]
		}
	}

	if err != nil {
		c.Update(claim.ActionUninstall, claim.StatusFailure)
		c.Result.Message = err.Error()
		return err
	}

	c.Update(claim.ActionUninstall, claim.StatusSuccess)
	return nil
}
