package fixtures

import "fmt"

type Chain struct{}

func (c *Chain) ChainCall(arg1 string, arg2 string, arg3 string) *Chain {
	return c
}

func (c *Chain) String() string {
	return ""
}

func NewChain() *Chain {
	return &Chain{}
}

func ChainedCalls() {
	c := Chain{}
	c.ChainCall("a long argument", "another long argument", "a third long argument").ChainCall("a long argument2", "another long argument2", "a third long argument2").ChainCall("a long argument3", "another long argument3", "a third long argument3")
	NewChain().ChainCall(
		"a really really really really really long argument4",
		"another really really really really really long argument4",
		fmt.Sprintf("%v", "this is a long method")).ChainCall("a really really really really really long argument5", "another really really really really really long argument5", "a third really really really really really long argument5").ChainCall("a", "b", fmt.Sprintf("%v", "this is a long method"))
	NewChain().ChainCall("a", "b", "c").ChainCall("d", "e", "f")
	NewChain().ChainCall("a", "b", "c").
		ChainCall("d", "e", "f").ChainCall("g", "h", "i")

	// Test case: pre-broken chain where middle line is still too long
	// See https://github.com/segmentio/golines/issues/142 discussion
	c.ChainCall("short", "args", "here").
		ChainCall("a]]really really really long argument", "another really really long argument", "a third really long arg").
		ChainCall("final", "short", "args")

	// Test case: pre-broken chain with nested function call that makes middle line too long
	c.ChainCall("ctx", "", "").
		ChainCall(fmt.Sprintf("%s %s %s", "event.Payload.Target.Tenant", "event.Payload.Target.UserID", "event.Payload.Target.VehicleID"), "", "").
		ChainCall("Refreshed forecast as a result of a schedule change", "", "")

	// Test case: simulating logger chain with EmbedObject pattern
	c.ChainCall("ctx", "", "").
		ChainCall(NewChain().ChainCall("event.Payload.Target.Tenant", "event.Payload.Target.UserID", "event.Payload.Target.VehicleID").String(), "", "").
		ChainCall("Refreshed forecast as a result of a schedule change", "", "")

	// Test case: long chain on single line - after breaking at dots, each segment should also be shortened
	c.ChainCall("ctx", "", "").ChainCall(NewChain().ChainCall("event.Payload.Target.Tenant", "event.Payload.Target.UserID", "event.Payload.Target.VehicleID").String(), "", "").ChainCall("msg", "", "")
}
