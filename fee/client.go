// Package fee provides the /application_fees APIs
package fee

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke application_fees APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an application fee.
// For more details see https://stripe.com/docs/api#retrieve_application_fee.
func Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.ApplicationFeeParams) (*stripe.ApplicationFee, error) {
	path := stripe.FormatURLPath("/application_fees/%s", id)
	fee := &stripe.ApplicationFee{}
	err := c.B.Call2("GET", path, c.Key, params, fee)
	return fee, err
}

// List returns a list of application fees.
// For more details see https://stripe.com/docs/api#list_application_fees.
func List(params *stripe.ApplicationFeeListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.ApplicationFeeListParams) *Iter {
	return &Iter{stripe.GetIter2(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ApplicationFeeList{}
		err := c.B.CallRaw("GET", "/application_fees", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of ApplicationFees.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// ApplicationFee returns the most recent ApplicationFee
// visited by a call to Next.
func (i *Iter) ApplicationFee() *stripe.ApplicationFee {
	return i.Current().(*stripe.ApplicationFee)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
