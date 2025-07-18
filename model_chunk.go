/*
 * Woolball AI Network API
 *
 * **Transform idle browsers into a powerful distributed AI inference network**  For detailed examples and model lists, visit our [GitHub repository](https://github.com/woolball-xyz/woolball-server).
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Chunk struct {
	Timestamp []float64 `json:"timestamp,omitempty"`
	Text string `json:"text,omitempty"`
}
