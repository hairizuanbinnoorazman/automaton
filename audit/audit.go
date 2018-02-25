// Package audit This package contains several subpackages to handle audits
package audit

// Auditor contains a few main functions.
// It needs to be able to do the following:
// - RunAudit
// - RenderOutput. The default provided by the package right now is markdown but this may change in the future.
type Auditor interface {
	RunAudit() error
}
