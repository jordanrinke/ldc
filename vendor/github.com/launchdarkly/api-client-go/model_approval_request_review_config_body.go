/*
 * LaunchDarkly REST API
 *
 * Build custom integrations with the LaunchDarkly REST API
 *
 * API version: 5.3.0
 * Contact: support@launchdarkly.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package ldapi

type ApprovalRequestReviewConfigBody struct {
	// One of approve, decline, or comment.
	Kind string `json:"kind"`
	// comment will be included in audit log item for change.
	Comment string `json:"comment,omitempty"`
}
